// Package bdd provides the executable behavior contract shared by every Badak Mini test adapter.
package bdd

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

// StepKind identifies the primary Gherkin keyword a step inherits.
type StepKind string

const (
	// GivenStep identifies setup expressed with Given or an inherited And/But.
	GivenStep StepKind = "Given"
	// WhenStep identifies the primary action expressed with When or an inherited And/But.
	WhenStep StepKind = "When"
	// ThenStep identifies an observation expressed with Then or an inherited And/But.
	ThenStep StepKind = "Then"
)

// Step is one expanded scenario step used by static binding compliance.
type Step struct {
	Feature  string
	Scenario string
	Kind     StepKind
	Text     string
}

// Feature records the observable catalog of one canonical feature file.
type Feature struct {
	Path          string
	ScenarioCount int
	Steps         []Step
}

// Catalog is the sorted recursive canonical feature corpus.
type Catalog struct {
	Features []Feature
}

// Paths returns every canonical feature path in deterministic order.
func (catalog Catalog) Paths() []string {
	paths := make([]string, 0, len(catalog.Features))
	for _, feature := range catalog.Features {
		paths = append(paths, feature.Path)
	}

	return paths
}

// Steps returns every expanded scenario step in corpus order.
func (catalog Catalog) Steps() []Step {
	var steps []Step
	for _, feature := range catalog.Features {
		steps = append(steps, feature.Steps...)
	}

	return steps
}

// ValidateAdapterParity proves every required adapter consumes the exact same catalog.
func ValidateAdapterParity(adapters map[string]Catalog) error {
	var expectedName string
	var expected Catalog
	for name, catalog := range adapters {
		if expectedName == "" {
			expectedName = name
			expected = catalog
			continue
		}
		if !reflect.DeepEqual(expected, catalog) {
			return fmt.Errorf("adapter %q corpus differs from %q", name, expectedName)
		}
	}

	return nil
}

// Discover recursively loads every .feature file below root.
func Discover(root string) (Catalog, error) {
	catalog, err := DiscoverFS(os.DirFS(root), ".")
	if err != nil {
		return Catalog{}, err
	}
	for featureIndex := range catalog.Features {
		catalog.Features[featureIndex].Path = filepath.Join(root, catalog.Features[featureIndex].Path)
		for stepIndex := range catalog.Features[featureIndex].Steps {
			catalog.Features[featureIndex].Steps[stepIndex].Feature = catalog.Features[featureIndex].Path
		}
	}

	return catalog, nil
}

// DiscoverFS recursively loads a corpus from an injected filesystem.
func DiscoverFS(filesystem fs.FS, root string) (Catalog, error) {
	paths := make([]string, 0)
	err := fs.WalkDir(filesystem, root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".feature" {
			return nil
		}
		paths = append(paths, path)

		return nil
	})
	if err != nil {
		return Catalog{}, fmt.Errorf("discover behavior features: %w", err)
	}
	if len(paths) == 0 {
		return Catalog{}, fmt.Errorf("no feature files found below %q", root)
	}
	sort.Strings(paths)

	features := make([]Feature, 0, len(paths))
	for _, path := range paths {
		feature, parseErr := parseFeature(filesystem, path)
		if parseErr != nil {
			return Catalog{}, parseErr
		}
		features = append(features, feature)
	}

	return Catalog{Features: features}, nil
}

func parseFeature(filesystem fs.FS, path string) (Feature, error) {
	contents, err := fs.ReadFile(filesystem, path)
	if err != nil {
		return Feature{}, fmt.Errorf("open behavior feature %q: %w", path, err)
	}
	state := featureParser{path: path, feature: Feature{Path: path}}
	scanner := bufio.NewScanner(strings.NewReader(string(contents)))

	for scanner.Scan() {
		if err := state.consume(strings.TrimSpace(scanner.Text())); err != nil {
			return Feature{}, err
		}
	}
	if err := scanner.Err(); err != nil {
		return Feature{}, fmt.Errorf("read behavior feature %q: %w", path, err)
	}
	if err := state.finishScenario(); err != nil {
		return Feature{}, err
	}
	if !state.hasFeature {
		return Feature{}, fmt.Errorf("%s: missing non-empty Feature declaration", path)
	}
	if state.feature.ScenarioCount == 0 {
		return Feature{}, fmt.Errorf("%s: feature must contain at least one scenario", path)
	}

	return state.feature, nil
}

type featureParser struct {
	path            string
	feature         Feature
	scenario        string
	inherited       StepKind
	primary         map[StepKind]int
	scenarioSteps   []Step
	outline         bool
	insideExamples  bool
	exampleRows     int
	exampleHeader   bool
	insideDocString bool
	hasFeature      bool
}

func (state *featureParser) consume(line string) error {
	handled, err := state.consumeDirective(line)
	if handled || err != nil {
		return err
	}
	if state.scenario == "" {
		return nil
	}
	if strings.HasPrefix(line, "Examples:") {
		state.insideExamples = true
		state.exampleHeader = false
		return nil
	}
	if state.insideExamples && strings.HasPrefix(line, "|") {
		state.recordExampleRow()
		return nil
	}
	state.recordStep(line)
	return nil
}

func (state *featureParser) consumeDirective(line string) (bool, error) {
	if line == `"""` {
		state.insideDocString = !state.insideDocString
		return true, nil
	}
	if state.insideDocString || line == "" || strings.HasPrefix(line, "#") {
		return true, nil
	}
	if strings.HasPrefix(line, "@") {
		return true, state.validateTags(line)
	}
	if featureName, ok := strings.CutPrefix(line, "Feature:"); ok {
		state.hasFeature = strings.TrimSpace(featureName) != ""
		return true, nil
	}
	if name, isOutline, ok := scenarioName(line); ok {
		return true, state.startScenario(name, isOutline)
	}
	return false, nil
}

func (state *featureParser) validateTags(line string) error {
	for tag := range strings.FieldsSeq(line) {
		switch strings.ToLower(tag) {
		case "@unit", "@integration", "@e2e":
			return fmt.Errorf("%s: layer-filter tag %q is forbidden", state.path, tag)
		}
	}
	return nil
}

func (state *featureParser) startScenario(name string, outline bool) error {
	if err := state.finishScenario(); err != nil {
		return err
	}
	state.scenario = name
	state.outline = outline
	state.inherited = ""
	state.primary = map[StepKind]int{}
	state.scenarioSteps = nil
	state.insideExamples = false
	state.exampleRows = 0
	state.exampleHeader = false
	return nil
}

func (state *featureParser) recordExampleRow() {
	if state.exampleHeader {
		state.exampleRows++
		return
	}
	state.exampleHeader = true
}

func (state *featureParser) recordStep(line string) {
	kind, text, primaryStep := scenarioStep(line, state.inherited)
	if kind == "" {
		return
	}
	if primaryStep {
		state.primary[kind]++
		state.inherited = kind
	}
	state.scenarioSteps = append(
		state.scenarioSteps,
		Step{Feature: state.path, Scenario: state.scenario, Kind: kind, Text: text},
	)
}

func (state *featureParser) finishScenario() error {
	if state.scenario == "" {
		return nil
	}
	for _, kind := range []StepKind{GivenStep, WhenStep, ThenStep} {
		if state.primary[kind] != 1 {
			return fmt.Errorf(
				"%s: scenario %q requires exactly one primary %s step",
				state.path,
				state.scenario,
				kind,
			)
		}
	}
	expansions := 1
	if state.outline {
		if state.exampleRows == 0 {
			return fmt.Errorf("%s: scenario outline %q requires at least one example row", state.path, state.scenario)
		}
		expansions = state.exampleRows
	}
	state.feature.ScenarioCount += expansions
	for range expansions {
		state.feature.Steps = append(state.feature.Steps, state.scenarioSteps...)
	}
	return nil
}

func scenarioName(line string) (string, bool, bool) {
	if name, ok := strings.CutPrefix(line, "Scenario:"); ok {
		return strings.TrimSpace(name), false, true
	}
	for _, prefix := range []string{"Scenario Outline:", "Scenario Template:"} {
		if name, ok := strings.CutPrefix(line, prefix); ok {
			return strings.TrimSpace(name), true, true
		}
	}

	return "", false, false
}

func scenarioStep(line string, inherited StepKind) (StepKind, string, bool) {
	for _, kind := range []StepKind{GivenStep, WhenStep, ThenStep} {
		prefix := string(kind) + " "
		if text, ok := strings.CutPrefix(line, prefix); ok {
			return kind, strings.TrimSpace(text), true
		}
	}
	for _, prefix := range []string{"And ", "But "} {
		if text, ok := strings.CutPrefix(line, prefix); ok && inherited != "" {
			return inherited, strings.TrimSpace(text), false
		}
	}

	return "", "", false
}

// RepositoryRoot returns the nearest ancestor containing the workspace specs directory.
func RepositoryRoot(start string) (string, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("resolve repository start: %w", err)
	}
	for {
		info, statErr := os.Stat(filepath.Join(current, "specs"))
		if statErr == nil && info.IsDir() {
			return current, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", errors.New("could not find repository specs directory")
		}
		current = parent
	}
}
