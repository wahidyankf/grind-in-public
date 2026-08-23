package bdd

import (
	"fmt"
	"regexp"

	"github.com/cucumber/godog"
)

// Definition is one typed binding and its scenario-scoped implementation.
type Definition struct {
	Kind       StepKind
	Expression string
	Handler    func(*State) any
}

// Registry is the single binding catalog shared by every adapter.
type Registry struct {
	definitions []Definition
}

// NewRegistry creates a binding catalog from explicit definitions.
func NewRegistry(definitions ...Definition) Registry {
	return Registry{definitions: append([]Definition(nil), definitions...)}
}

// Definitions returns a defensive copy for static compliance checks.
func (registry Registry) Definitions() []Definition {
	return append([]Definition(nil), registry.definitions...)
}

// Register installs every binding against scenario-scoped state.
func (registry Registry) Register(context *godog.ScenarioContext, state **State) {
	for _, definition := range registry.definitions {
		current := definition
		context.Step(regexp.MustCompile(current.Expression), func() error {
			handler, ok := current.Handler(*state).(func() error)
			if !ok {
				return fmt.Errorf("binding %q has an unsupported handler", current.Expression)
			}

			return handler()
		})
	}
}

// Validate proves every canonical step resolves once and every binding is used.
func (registry Registry) Validate(steps []Step) []error {
	used := make([]bool, len(registry.definitions))
	var findings []error
	for _, step := range steps {
		matches, matchFindings := registry.matches(step)
		findings = append(findings, matchFindings...)
		switch len(matches) {
		case 0:
			findings = append(findings, fmt.Errorf("undefined %s step %q", step.Kind, step.Text))
		case 1:
			used[matches[0]] = true
		default:
			findings = append(findings, fmt.Errorf("ambiguous %s step %q", step.Kind, step.Text))
		}
	}
	for index, definition := range registry.definitions {
		if !used[index] {
			findings = append(findings, fmt.Errorf("unused %s binding %q", definition.Kind, definition.Expression))
		}
	}

	return findings
}

func (registry Registry) matches(step Step) ([]int, []error) {
	matches := make([]int, 0, 1)
	var findings []error
	for index, definition := range registry.definitions {
		if definition.Kind != step.Kind {
			continue
		}
		matched, err := regexp.MatchString("^(?:"+definition.Expression+")$", step.Text)
		if err != nil {
			findings = append(findings, fmt.Errorf("invalid binding %q: %w", definition.Expression, err))
			continue
		}
		if matched {
			matches = append(matches, index)
		}
	}

	return matches, findings
}
