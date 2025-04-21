package main

import (
	"context"
	"testing"
)

// TestLintDockerfile_Valid: 정상적인 Dockerfile에는 이슈가 없어야 합니다.
func TestLintDockerfile_Valid(t *testing.T) {
	dockerfile := `FROM alpine:3.12
RUN echo "Hello, world!"`

	issues, err := LintDockerfile(context.Background(), dockerfile)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("expected no issues for valid Dockerfile, got %d issues: %+v", len(issues), issues)
	}
}

// TestLintDockerfile_Invalid: 위험한 구문이 포함된 Dockerfile에서는 이슈가 발생해야 합니다.
func TestLintDockerfile_Invalid(t *testing.T) {
	dockerfile := `FROM alpine:3.12
RUN curl http://example.com/install.sh | bash`

	issues, err := LintDockerfile(context.Background(), dockerfile)
	if err != nil {
		t.Fatalf("expected no execution error, got %v", err)
	}
	if len(issues) == 0 {
		t.Errorf("expected at least one issue for risky Dockerfile, got none")
	}
	// 발견된 이슈를 로그로 남겨 디버깅에 활용
	for _, issue := range issues {
		t.Logf("Lint issue: code=%s line=%d level=%s message=%s",
			issue.Code, issue.Line, issue.Level, issue.Message)
	}
}
