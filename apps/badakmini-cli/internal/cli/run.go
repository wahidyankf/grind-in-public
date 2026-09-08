// Package cli owns Badak Mini command parsing, dispatch, and observable process behaviour.
package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/wahidyankf/grind-in-public/apps/badakmini-cli/internal/rulechange"
)

// Runtime supplies the process and repository boundaries used by command orchestration.
type Runtime struct {
	Stdin              io.Reader
	Stdout             io.Writer
	Stderr             io.Writer
	FindRepositoryRoot func() (string, error)
	ListStagedPaths    func(string) ([]string, error)
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

	// Neither failure may block the edit, and a response of plain strings
	// always marshals -- so the two are one path rather than two.
	encoded, err := json.Marshal(response)
	if err != nil || writef(stdout, "%s\n", encoded) != nil {
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

// writef propagates output failures so commands do not report success when a
// caller cannot receive their validation result.
func writef(writer io.Writer, format string, arguments ...any) error {
	_, err := fmt.Fprintf(writer, format, arguments...)
	if err != nil {
		return fmt.Errorf("write formatted output: %w", err)
	}

	return nil
}
