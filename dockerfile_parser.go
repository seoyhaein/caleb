package main

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"os"
	"strings"
)

// ParseDockerfile 읽어들인 Dockerfile을 AST로 파싱해서 반환합니다.
func ParseDockerfile(path string) (*parser.Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Dockerfile: %w", err)
	}
	defer f.Close()

	result, err := parser.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Dockerfile: %w", err)
	}
	return result.AST, nil
}

// ParseDockerfileString 문자열로 제공된 Dockerfile도 파싱할 수 있습니다.
func ParseDockerfileString(dockerfile string) (*parser.Node, error) {
	// parser.Parse는 (*parser.Result, error)를 반환합니다.
	result, err := parser.Parse(strings.NewReader(dockerfile))
	if err != nil {
		return nil, fmt.Errorf("failed to parse Dockerfile: %w", err)
	}
	// AST 필드에 실제 *parser.Node가 있습니다.
	return result.AST, nil
}
