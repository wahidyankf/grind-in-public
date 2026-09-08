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
		Short: "Announce workflows required by rule changes",
		Long: "Announce the workflows a staged or pending rule change requires. " +
			"Documentation hygiene is checked by RHINO, from repo-config.yml.",
		Example: strings.Join([]string{
			"  badak-mini harness rule-change validate",
			"  badak-mini harness rule-change hook",
		}, "\n"),
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	root.CompletionOptions.DisableDefaultCmd = true
	root.DisableSuggestions = true

	harness := commandGroup("harness", "Validate agent harness governance")
	harness.AddCommand(ruleChangeCommand(runtime))
	root.AddCommand(harness)

	return root
}

func commandGroup(use, summary string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: summary,
		// Cobra reaches a group's argument rule only once the group is
		// runnable; a group that merely holds subcommands answers any word
		// with its own help and exit 0. A name this CLI does not have -- a
		// retired check still wired into a stale hook -- has to be an invalid
		// invocation rather than a green run that checked nothing.
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error { return cmd.Help() },
	}
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
				// The diagnostic is best effort: a caller that cannot receive it
				// still gets the failing exit code, which is the same answer.
				_ = writef(
					runtime.Stderr,
					"ERROR: could not find the Git repository root: %v\n",
					err,
				)
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
