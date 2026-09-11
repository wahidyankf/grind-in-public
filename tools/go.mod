// This module exists only to pin two executables. `go tool` needs a module to
// resolve a pinned version from, and after Badak Mini's retirement there is no
// other Go module here to host these. Keeping them pinned by version, rather
// than running whatever is on PATH, is the discipline; the module is just
// where that pin lives.
//
// Actionlint validates this repository's GitHub Actions workflows.
// Govulncheck scans this module's own dependency tree, which matters because
// these two are executed rather than merely declared. Both versions are the
// ones Badak Mini pinned; the relocation moved them and changed neither.

module github.com/wahidyankf/grind-in-public/tools

go 1.26.6

tool (
	github.com/rhysd/actionlint/cmd/actionlint
	golang.org/x/vuln/cmd/govulncheck
)

require (
	github.com/rhysd/actionlint v1.7.12
	golang.org/x/vuln v1.3.0
)

require (
	github.com/bmatcuk/doublestar/v4 v4.10.0 // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/fatih/color v1.19.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.21 // indirect
	github.com/mattn/go-shellwords v1.0.12 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	go.yaml.in/yaml/v4 v4.0.0-rc.3 // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/sync v0.23.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/telemetry v0.0.0-20260908163034-4bcc4b2ee518 // indirect
	golang.org/x/tools v0.50.0 // indirect
)
