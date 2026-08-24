package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

const invalidInvocationExitCode = 2

type repositoryAction func(Runtime, string) int

var errCommandFailed = errors.New("command failed")

// Run executes one Cobra command tree and returns the public process exit code.
func Run(ctx context.Context, runtime Runtime, args []string) int {
	stdout := &errorTrackingWriter{target: runtime.Stdout}
	root := newRootCommand(runtime)
	root.SetArgs(args)
	root.SetIn(runtime.Stdin)
	root.SetOut(stdout)
	root.SetErr(runtime.Stderr)

	_, err := root.ExecuteContextC(ctx)
	if stdout.err != nil {
		return 1
	}
	if err == nil {
		return 0
	}

	var commandErr commandExitError
	if errors.As(err, &commandErr) {
		return commandErr.code
	}
	if writeErr := writef(runtime.Stderr, "%s", root.UsageString()); writeErr != nil {
		return 1
	}

	// Exit status 2 distinguishes invalid input from a failed validation.
	return invalidInvocationExitCode
}

func newRootCommand(runtime Runtime) *cobra.Command {
	root := &cobra.Command{
		Use:   "badak-mini",
		Short: "Validate repository-local governance",
		Long: "Validate governance Markdown word limits, repository-local Markdown links, " +
			"harness capability parity, and rule-change workflow notices.",
		Example: strings.Join([]string{
			"  badak-mini harness instruction-size validate",
			"  badak-mini harness markdown-links validate",
			"  badak-mini harness rule-change validate",
			"  badak-mini harness rule-change hook",
			"  badak-mini harness capability-parity validate",
		}, "\n"),
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	root.CompletionOptions.DisableDefaultCmd = true
	root.DisableSuggestions = true

	harness := commandGroup("harness", "Validate agent harness governance")
	harness.AddCommand(
		validationGroup(
			"instruction-size",
			"Check agent instruction word limits",
			runtime,
			validateInstructionSize,
		),
		validationGroup(
			"markdown-links",
			"Check repository-local Markdown links",
			runtime,
			validateMarkdownLinks,
		),
		validationGroup(
			"capability-parity",
			"Check capabilities exposed by every harness",
			runtime,
			validateCapabilityParity,
		),
		ruleChangeCommand(runtime),
	)
	root.AddCommand(harness)

	return root
}

func commandGroup(use, summary string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: summary,
	}
}

func validationGroup(
	use, summary string,
	runtime Runtime,
	action repositoryAction,
) *cobra.Command {
	group := commandGroup(use, summary)
	group.AddCommand(repositoryCommand("validate", summary, runtime, action))
	return group
}

func ruleChangeCommand(runtime Runtime) *cobra.Command {
	group := commandGroup("rule-change", "Announce workflows required by rule changes")
	group.AddCommand(
		repositoryCommand("validate", "Inspect staged paths", runtime, announceStagedRuleChange),
		repositoryCommand(
			"hook",
			"Inspect a harness pre-edit payload",
			runtime,
			func(runtime Runtime, root string) int {
				return announceHookRuleChange(root, runtime.Stdin, runtime.Stdout)
			},
		),
	)
	return group
}

func repositoryCommand(
	use, summary string,
	runtime Runtime,
	action repositoryAction,
) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: summary,
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			root, err := runtime.FindRepositoryRoot()
			if err != nil {
				writeErr := writef(
					runtime.Stderr,
					"ERROR: could not find the Git repository root: %v\n",
					err,
				)
				if writeErr != nil {
					return commandExitError{error: errCommandFailed, code: 1}
				}
				return commandExitError{error: errCommandFailed, code: 1}
			}

			exitCode := action(runtime, root)
			if exitCode != 0 {
				return commandExitError{error: errCommandFailed, code: exitCode}
			}
			return nil
		},
	}
}

type commandExitError struct {
	error

	code int
}

type errorTrackingWriter struct {
	target io.Writer
	err    error
}

func (writer *errorTrackingWriter) Write(message []byte) (int, error) {
	written, err := writer.target.Write(message)
	if err == nil {
		return written, nil
	}

	wrapped := fmt.Errorf("write Cobra output: %w", err)
	if writer.err == nil {
		writer.err = wrapped
	}
	return written, wrapped
}
