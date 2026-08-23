package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestProductionConstructionPreservesHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := execute([]string{"--help"}, strings.NewReader(""), &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("expected help success, got %d", exitCode)
	}
	if !strings.Contains(stdout.String(), "Usage:") {
		t.Fatalf("expected usage output, got %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no error output, got %q", stderr.String())
	}
}

func TestProductionRuntimeBindsAllAdapters(t *testing.T) {
	stdin := strings.NewReader("payload")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	runtime := productionRuntime(stdin, &stdout, &stderr)

	if runtime.Stdin != stdin || runtime.Stdout != &stdout || runtime.Stderr != &stderr {
		t.Fatal("expected production runtime to preserve process streams")
	}
	bindings := []struct {
		name     string
		actual   any
		expected any
	}{
		{name: "repository discovery", actual: runtime.FindRepositoryRoot, expected: findRepositoryRoot},
		{name: "governance", actual: runtime.CheckGovernance, expected: checkGovernance},
		{name: "Markdown links", actual: runtime.CheckMarkdownLinks, expected: checkMarkdownLinks},
		{name: "staged paths", actual: runtime.ListStagedPaths, expected: listStagedPaths},
		{name: "parity", actual: runtime.CheckParity, expected: checkParity},
	}
	for _, binding := range bindings {
		if reflect.ValueOf(binding.actual).Pointer() != reflect.ValueOf(binding.expected).Pointer() {
			t.Fatalf("expected %s production adapter binding", binding.name)
		}
	}
}
