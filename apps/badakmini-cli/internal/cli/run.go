// Package cli owns Badak Mini command parsing, dispatch, and observable process behavior.
package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/rulechange"
)

// Each supported invocation is spelled out, so an unknown command fails with
// usage rather than falling through to a check the caller did not ask for.
const (
	instructionSizeCommand    = "harness instruction-size validate"
	markdownLinksCommand      = "harness markdown-links validate"
	ruleChangeValidateCommand = "harness rule-change validate"
	ruleChangeHookCommand     = "harness rule-change hook"
	capabilityParityCommand   = "harness capability-parity validate"
	invalidInvocationExitCode = 2
)

var supportedCommands = []string{
	instructionSizeCommand,
	markdownLinksCommand,
	ruleChangeValidateCommand,
	ruleChangeHookCommand,
	capabilityParityCommand,
}

const usage = `Usage:
  badak-mini harness instruction-size validate
  badak-mini harness markdown-links validate
  badak-mini harness rule-change validate
  badak-mini harness rule-change hook
  badak-mini harness capability-parity validate

Validate governance Markdown word limits, repository-local Markdown links, the
subagents, skills, and commands every harness exposes, or announce the
rules-propagation workflow when a rule changes. The rule-change validate form
reads the staged paths; its hook form reads a pre-edit payload on stdin.
`

// Runtime supplies the process and repository boundaries used by command orchestration.
type Runtime struct {
	Stdin              io.Reader
	Stdout             io.Writer
	Stderr             io.Writer
	FindRepositoryRoot func() (string, error)
	CheckGovernance    func(string) ([]governance.Finding, error)
	CheckMarkdownLinks func(string) ([]markdownlinks.Finding, error)
	ListStagedPaths    func(string) ([]string, error)
	CheckParity        func(string) ([]parity.Finding, error)
}

// Run parses one invocation and returns the public process exit code.
func Run(ctx context.Context, runtime Runtime, args []string) int {
	if exitCode, handled := handleUsage(args, runtime.Stdout, runtime.Stderr); handled {
		return exitCode
	}

	root, err := runtime.FindRepositoryRoot()
	if err != nil {
		writeErr := writef(runtime.Stderr, "ERROR: could not find the Git repository root: %v\n", err)
		if writeErr != nil {
			return 1
		}
		return 1
	}

	return executeCommand(ctx, runtime, strings.Join(args, " "), root)
}

// handleUsage separates invocations that need no repository from validation.
func handleUsage(args []string, stdout, stderr io.Writer) (int, bool) {
	// Help is successful even outside a repository because it does not inspect
	// files; invalid commands fail before doing an unnecessary root lookup.
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		if err := writef(stdout, usage); err != nil {
			return 1, true
		}

		return 0, true
	}
	if matchesCommand(args) {
		return 0, false
	}
	if err := writef(stderr, usage); err != nil {
		return 1, true
	}

	// Exit status 2 distinguishes invalid input from a failed validation.
	return invalidInvocationExitCode, true
}

// executeCommand keeps command routing separate from repository discovery and
// usage handling, so each stage has one failure domain.
func executeCommand(
	_ context.Context,
	runtime Runtime,
	command, root string,
) int {
	switch command {
	case instructionSizeCommand:
		return validateInstructionSize(runtime, root)
	case ruleChangeValidateCommand:
		return announceStagedRuleChange(runtime, root)
	case ruleChangeHookCommand:
		return announceHookRuleChange(root, runtime.Stdin, runtime.Stdout)
	case capabilityParityCommand:
		return validateCapabilityParity(runtime, root)
	default:
		return validateMarkdownLinks(runtime, root)
	}
}

// announceStagedRuleChange reports a staged rule change to a contributor. It
// succeeds either way: the workflow is the author's next step, not a gate that
// a hook can decide has been satisfied.
func announceStagedRuleChange(runtime Runtime, root string) int {
	staged, err := runtime.ListStagedPaths(root)
	if err != nil {
		writeErr := writef(runtime.Stderr, "ERROR: %v\n", err)
		if writeErr != nil {
			return 1
		}
		return 1
	}

	paths := rulechange.RulePaths(staged)
	if len(paths) == 0 {
		return 0
	}

	if err := writef(runtime.Stdout, "%s\n", rulechange.Notice(paths)); err != nil {
		return 1
	}
	return 0
}

// announceHookRuleChange answers a harness pre-edit hook. It returns the notice
// as additional context so the agent loads the workflow before it edits, and it
// stays silent for every other file so ordinary work is never interrupted.
func announceHookRuleChange(root string, stdin io.Reader, stdout io.Writer) int {
	payload, err := io.ReadAll(stdin)
	if err != nil {
		// A hook that cannot read its payload must not block the edit.
		return 0
	}

	paths := rulechange.RulePaths(rulechange.HookPaths(payload, root))
	if len(paths) == 0 {
		return 0
	}

	response := hookResponse{}
	response.HookSpecificOutput.HookEventName = "PreToolUse"
	response.HookSpecificOutput.AdditionalContext = rulechange.Notice(paths)
	response.SystemMessage = rulechange.Notice(paths)

	encoded, err := json.Marshal(response)
	if err != nil {
		return 0
	}
	if err := writef(stdout, "%s\n", encoded); err != nil {
		return 0
	}
	return 0
}

// hookResponse is the pre-edit hook reply shape: additional context informs the
// agent without denying the edit, and the system message tells the human.
type hookResponse struct {
	HookSpecificOutput struct {
		HookEventName     string `json:"hookEventName"`
		AdditionalContext string `json:"additionalContext"`
	} `json:"hookSpecificOutput"`
	SystemMessage string `json:"systemMessage"`
}

func validateInstructionSize(runtime Runtime, root string) int {
	findings, err := runtime.CheckGovernance(root)
	if err != nil {
		writeErr := writef(runtime.Stderr, "ERROR: %v\n", err)
		if writeErr != nil {
			return 1
		}
		return 1
	}
	if len(findings) == 0 {
		err := writef(runtime.Stdout, "Governance word counts are within the %d-word limit.\n", governance.MaxWords)
		if err != nil {
			return 1
		}
		return 0
	}
	for _, finding := range findings {
		// Report every over-limit document in one run so a contributor can repair
		// the complete guidance set before retrying the hook.
		err := writef(
			runtime.Stderr,
			"ERROR: %s contains %d words; the limit is %d.\n",
			finding.Path,
			finding.WordCount,
			governance.MaxWords,
		)
		if err != nil {
			return 1
		}
	}
	guidance := "Use progressive disclosure: split detailed guidance into focused files.\n"
	if err := writef(runtime.Stderr, "%s", guidance); err != nil {
		return 1
	}
	return 1
}

// validateCapabilityParity fails when one harness exposes a capability another
// supporting harness lacks, because an agent that exists for one tool and not
// the next makes the repository behave differently depending on who runs it.
func validateCapabilityParity(runtime Runtime, root string) int {
	findings, err := runtime.CheckParity(root)
	if err != nil {
		writeErr := writef(runtime.Stderr, "ERROR: %v\n", err)
		if writeErr != nil {
			return 1
		}
		return 1
	}
	if len(findings) == 0 {
		err := writef(runtime.Stdout, "Every harness exposes the same subagents, skills, and commands.\n")
		if err != nil {
			return 1
		}
		for _, note := range parity.UnsupportedNotes() {
			// Naming the exemption keeps an absent harness from reading as an
			// oversight the next time someone compares the directories.
			err := writef(runtime.Stdout, "Exempt, %s.\n", note)
			if err != nil {
				return 1
			}
		}
		return 0
	}
	for _, finding := range findings {
		err := writef(runtime.Stderr, "ERROR: %s\n", finding.Message())
		if err != nil {
			return 1
		}
	}
	policy := "See repo-governance/conventions/harness-capability-parity-policy.md.\n"
	if err := writef(runtime.Stderr, "%s", policy); err != nil {
		return 1
	}
	return 1
}

func validateMarkdownLinks(runtime Runtime, root string) int {
	findings, err := runtime.CheckMarkdownLinks(root)
	if err != nil {
		writeErr := writef(runtime.Stderr, "ERROR: %v\n", err)
		if writeErr != nil {
			return 1
		}
		return 1
	}
	if len(findings) == 0 {
		err := writef(runtime.Stdout, "Repository-local Markdown links are valid.\n")
		if err != nil {
			return 1
		}
		return 0
	}
	for _, finding := range findings {
		// Source path and one-based line number make a hook failure actionable in
		// a terminal or CI log without requiring a separate report file.
		err := writef(
			runtime.Stderr,
			"ERROR: %s:%d: %q %s.\n",
			finding.Path,
			finding.Line,
			finding.Destination,
			finding.Problem,
		)
		if err != nil {
			return 1
		}
	}
	return 1
}

// writef propagates output failures so commands do not report success when a
// caller cannot receive their validation result.
func writef(writer io.Writer, format string, arguments ...any) error {
	_, err := fmt.Fprintf(writer, format, arguments...)
	if err != nil {
		return fmt.Errorf("write formatted output: %w", err)
	}

	return nil
}

func matchesCommand(args []string) bool {
	return slices.Contains(supportedCommands, strings.Join(args, " "))
}
