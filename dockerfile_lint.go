package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

// LintIssue represents a single hadolint finding.
type LintIssue struct {
	Code    string `json:"code"`
	Line    int    `json:"line"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

// LintDockerfile runs hadolint in a Docker container against the given Dockerfile content.
// It returns the list of issues (warnings or errors) found.
func LintDockerfile(ctx context.Context, dockerfileContent string) ([]LintIssue, error) {
	// Prepare the hadolint Docker command
	cmd := exec.CommandContext(ctx,
		"podman", "run", "--rm", "-i",
		"hadolint/hadolint:latest", // 태그도 꼭 붙여 주세요
		"hadolint",                 // entrypoint 대신 직접 실행
		"--format", "json", "-",
	)

	// Pipe the Dockerfile content into the container's stdin
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start hadolint container: %w", err)
	}

	// Write Dockerfile content
	if _, err := io.WriteString(stdin, dockerfileContent); err != nil {
		stdin.Close()
		return nil, fmt.Errorf("failed to send Dockerfile to hadolint: %w", err)
	}
	stdin.Close()

	// Wait for command to complete
	err = cmd.Wait()
	exitCode := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode = exitErr.ExitCode()
	}
	// exitCode 1 means hadolint found issues; higher codes are execution errors
	if err != nil && exitCode > 1 {
		return nil, fmt.Errorf("hadolint execution failed (exit %d): %s", exitCode, output.String())
	}

	// Parse JSON output into issues slice
	var issues []LintIssue
	if output.Len() > 0 {
		if err := json.Unmarshal(output.Bytes(), &issues); err != nil {
			return nil, fmt.Errorf("failed to parse hadolint output: %w", err)
		}
	}
	return issues, nil
}
