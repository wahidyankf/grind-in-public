// Package cli owns Badak Mini command parsing, dispatch, and observable process behaviour.
package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/governance"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/markdownlinks"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/parity"
	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/rulechange"
)

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
