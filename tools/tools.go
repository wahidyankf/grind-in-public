// Package tools exists so this module has something to scan.
//
// Nothing imports it and nothing runs it. The `tool` directives in go.mod pin
// the two executables; this file names their library roots so
// `govulncheck -scan module` has a package to start from. Without a Go file a
// tool-only module is invisible to it, and a vulnerability scanner that
// silently scans nothing is worse than one that is obviously absent.

package tools

import (
	_ "github.com/rhysd/actionlint"
	_ "golang.org/x/vuln/scan"
)
