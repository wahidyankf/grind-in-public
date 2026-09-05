// Package bdd provides the executable behaviour contract shared by every Badak Mini test adapter.
package bdd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/cucumber/godog"
)

// CanonicalCatalog discovers the single behaviour corpus every Badak Mini adapter must execute.
func CanonicalCatalog() (Catalog, error) {
	_, sourcePath, _, ok := runtime.Caller(0)
	if !ok {
		return Catalog{}, errors.New("locate canonical behaviour catalog source")
	}
	workspaceRoot := filepath.Clean(filepath.Join(filepath.Dir(sourcePath), "..", "..", "..", ".."))
	return Discover(filepath.Join(workspaceRoot, "specs", "apps", "badakmini-cli", "behaviours"))
}

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

// DiscoverFS asks Godog to recursively parse and compile a corpus from an injected filesystem.
func DiscoverFS(filesystem fs.FS, root string) (Catalog, error) {
	if err := validateFeaturePolicies(filesystem, root); err != nil {
		return Catalog{}, err
	}
	parsedFeatures, err := godog.TestSuite{
		Name: "badakmini-catalog",
		Options: &godog.Options{
			FS:    filesystem,
			Paths: []string{root},
		},
	}.RetrieveFeatures()
	if err != nil {
		return Catalog{}, fmt.Errorf("parse behaviour features with Godog: %w", err)
	}
	if len(parsedFeatures) == 0 {
		return Catalog{}, emptyCatalogError(filesystem, root)
	}

	features := make([]Feature, 0, len(parsedFeatures))
	for _, parsedFeature := range parsedFeatures {
		primarySteps := indexPrimarySteps(parsedFeature.GherkinDocument)
		feature, conversionErr := catalogFeature(parsedFeature.GherkinDocument, parsedFeature.Pickles, primarySteps)
		if conversionErr != nil {
			return Catalog{}, conversionErr
		}
		features = append(features, feature)
	}

	return Catalog{Features: features}, nil
}

var (
	scenarioDeclaration = regexp.MustCompile(`(?i)^Scenario(?: Outline| Template)?:`)
	exemptionComment    = regexp.MustCompile(`^# Exemption\((integration|e2e)\): (.+); alternative-proof: (.+)$`)
	invalidReason       = regexp.MustCompile(`(?i)\b(?:hard|slow|flaky|cost(?:ly)?|expensive|not yet implemented|todo)\b`)
	alternativeProof    = regexp.MustCompile(`(?i)^[a-z0-9-]+:test(?::[a-z0-9-]+)*\s+/\s+\S`)
)

const (
	featureFileExtension = ".feature"
	previousSourceLine   = 2
)

type sourceTag struct {
	line int
	name string
}

func validateFeaturePolicies(filesystem fs.FS, root string) error {
	var findings []string
	err := fs.WalkDir(filesystem, root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(entry.Name()) != featureFileExtension {
			return nil
		}
		source, readErr := fs.ReadFile(filesystem, path)
		if readErr != nil {
			return fmt.Errorf("read behaviour feature %s: %w", path, readErr)
		}
		findings = append(findings, validateExemptionPolicy(path, string(source))...)
		return nil
	})
	if err != nil {
		return fmt.Errorf("inspect behaviour feature policy: %w", err)
	}
	if len(findings) > 0 {
		return errors.New(strings.Join(findings, "\n"))
	}
	return nil
}

func validateExemptionPolicy(resourceName, source string) []string {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(source, "\r\n", "\n"), "\r", "\n"), "\n")
	var findings []string
	var pending []sourceTag
	for index, line := range lines {
		trimmed := strings.TrimSpace(line)
		lineNumber := index + 1
		if strings.HasPrefix(trimmed, "@") {
			lineFindings, tags := inspectTagLine(resourceName, trimmed, lineNumber)
			findings = append(findings, lineFindings...)
			pending = append(pending, tags...)
			continue
		}
		if isGherkinDeclaration(trimmed) {
			findings = append(findings, exemptionDeclarationFindings(
				resourceName, lines, pending, trimmed, lineNumber,
			)...)
			pending = nil
			continue
		}
		if trimmed != "" && !strings.HasPrefix(trimmed, "#") && len(pending) > 0 {
			findings = append(findings, fmt.Sprintf(
				"%s:%d: tags must be followed by their Gherkin declaration",
				resourceName, lineNumber,
			))
			pending = nil
		}
	}
	if len(pending) > 0 {
		findings = append(findings, resourceName+": dangling tags are not attached to a scenario")
	}
	return findings
}

func inspectTagLine(resourceName, line string, lineNumber int) ([]string, []sourceTag) {
	var findings []string
	var tags []sourceTag
	for field := range strings.FieldsSeq(line) {
		if !strings.HasPrefix(field, "@") {
			continue
		}
		name := strings.TrimPrefix(field, "@")
		switch strings.ToLower(name) {
		case "unit", "integration", "e2e", "unit-exempt", "no-unit", "no-integration", "no-e2e":
			findings = append(findings, fmt.Sprintf(
				"%s:%d: layer-filter tag @%s is forbidden; use canonical higher-layer exemptions",
				resourceName, lineNumber, name,
			))
		}
		tags = append(tags, sourceTag{line: lineNumber, name: strings.ToLower(name)})
	}
	return findings, tags
}

func exemptionDeclarationFindings(
	resourceName string,
	lines []string,
	pending []sourceTag,
	declaration string,
	lineNumber int,
) []string {
	var exemptions []sourceTag
	for _, tag := range pending {
		if tag.name == "integration-exempt" || tag.name == "e2e-exempt" {
			exemptions = append(exemptions, tag)
		}
	}
	if len(exemptions) == 0 {
		return nil
	}
	var findings []string
	if !scenarioDeclaration.MatchString(declaration) {
		findings = append(findings, fmt.Sprintf(
			"%s:%d: exemption tags may only annotate a Scenario or Scenario Outline",
			resourceName, lineNumber,
		))
	}
	for _, exemption := range exemptions {
		findings = append(findings, documentedExemptionFindings(resourceName, lines, exemption)...)
	}
	return findings
}

func documentedExemptionFindings(resourceName string, lines []string, exemption sourceTag) []string {
	layer := strings.TrimSuffix(exemption.name, "-exempt")
	comment := ""
	if exemption.line >= previousSourceLine {
		comment = strings.TrimSpace(lines[exemption.line-previousSourceLine])
	}
	match := exemptionComment.FindStringSubmatch(comment)
	if len(match) != 4 || match[1] != layer {
		return []string{fmt.Sprintf(
			"%s:%d: @%s requires the immediately preceding canonical comment",
			resourceName, exemption.line, exemption.name,
		)}
	}
	var findings []string
	if invalidReason.MatchString(match[2]) {
		findings = append(findings, fmt.Sprintf(
			"%s:%d: an exemption cannot be justified by difficulty, speed, cost, flakiness, or missing implementation",
			resourceName, exemption.line,
		))
	}
	if !alternativeProof.MatchString(match[3]) {
		findings = append(findings, fmt.Sprintf(
			"%s:%d: alternative proof must name an Nx test target and scenario after ' / '",
			resourceName, exemption.line,
		))
	}
	return findings
}

func isGherkinDeclaration(line string) bool {
	lower := strings.ToLower(line)
	prefixes := []string{
		"feature:", "rule:", "background:", "scenario:", "scenario outline:",
		"scenario template:", "examples:", "example:",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(lower, prefix) {
			return true
		}
	}
	return false
}

type primaryStepCounts struct {
	scenario string
	counts   map[StepKind]int
}

func catalogFeature(
	document *godog.GherkinDocument,
	pickles []*godog.Scenario,
	primarySteps map[string]primaryStepCounts,
) (Feature, error) {
	if document.Feature == nil || strings.TrimSpace(document.Feature.Name) == "" {
		return Feature{}, fmt.Errorf("%s: missing non-empty Feature declaration", document.Uri)
	}
	if len(pickles) == 0 {
		return Feature{}, fmt.Errorf("%s: feature must contain at least one scenario", document.Uri)
	}

	feature := Feature{Path: document.Uri, ScenarioCount: len(pickles)}
	for _, pickle := range pickles {
		if err := validateLayerTags(document.Uri, pickle); err != nil {
			return Feature{}, err
		}
		if len(pickle.AstNodeIds) == 0 {
			return Feature{}, fmt.Errorf("%s: locate scenario %q parsed by Godog", document.Uri, pickle.Name)
		}
		primary, present := primarySteps[pickle.AstNodeIds[0]]
		if !present {
			return Feature{}, fmt.Errorf("%s: locate scenario %q parsed by Godog", document.Uri, pickle.Name)
		}
		if err := validatePrimarySteps(document.Uri, primary); err != nil {
			return Feature{}, err
		}
		steps, err := catalogSteps(document.Uri, pickle)
		if err != nil {
			return Feature{}, err
		}
		feature.Steps = append(feature.Steps, steps...)
	}
	return feature, nil
}

func indexPrimarySteps(document *godog.GherkinDocument) map[string]primaryStepCounts {
	index := make(map[string]primaryStepCounts)
	if document.Feature == nil {
		return index
	}
	for _, child := range document.Feature.Children {
		if child.Scenario != nil {
			keywordTypes := make([]string, 0, len(child.Scenario.Steps))
			for _, step := range child.Scenario.Steps {
				keywordTypes = append(keywordTypes, string(step.KeywordType))
			}
			index[child.Scenario.Id] = countPrimarySteps(child.Scenario.Name, keywordTypes)
		}
		if child.Rule == nil {
			continue
		}
		for _, ruleChild := range child.Rule.Children {
			if ruleChild.Scenario != nil {
				keywordTypes := make([]string, 0, len(ruleChild.Scenario.Steps))
				for _, step := range ruleChild.Scenario.Steps {
					keywordTypes = append(keywordTypes, string(step.KeywordType))
				}
				index[ruleChild.Scenario.Id] = countPrimarySteps(
					ruleChild.Scenario.Name,
					keywordTypes,
				)
			}
		}
	}
	return index
}

func countPrimarySteps(scenario string, keywordTypes []string) primaryStepCounts {
	primary := primaryStepCounts{scenario: scenario, counts: map[StepKind]int{}}
	for _, keywordType := range keywordTypes {
		switch keywordType {
		case "Context":
			primary.counts[GivenStep]++
		case "Action":
			primary.counts[WhenStep]++
		case "Outcome":
			primary.counts[ThenStep]++
		}
	}
	return primary
}

func validateLayerTags(path string, pickle *godog.Scenario) error {
	for _, tag := range pickle.Tags {
		switch strings.ToLower(tag.Name) {
		case "@unit", "@integration", "@e2e":
			return fmt.Errorf("%s: layer-filter tag %q is forbidden", path, tag.Name)
		}
	}
	return nil
}

func validatePrimarySteps(path string, primary primaryStepCounts) error {
	for _, kind := range []StepKind{GivenStep, WhenStep, ThenStep} {
		if primary.counts[kind] != 1 {
			return fmt.Errorf("%s: scenario %q requires exactly one primary %s step", path, primary.scenario, kind)
		}
	}
	return nil
}

func catalogSteps(path string, pickle *godog.Scenario) ([]Step, error) {
	steps := make([]Step, 0, len(pickle.Steps))
	for _, pickleStep := range pickle.Steps {
		kind, err := pickleStepKind(string(pickleStep.Type))
		if err != nil {
			return nil, fmt.Errorf("%s: scenario %q: %w", path, pickle.Name, err)
		}
		steps = append(steps, Step{
			Feature:  path,
			Scenario: pickle.Name,
			Kind:     kind,
			Text:     pickleStep.Text,
		})
	}
	return steps, nil
}

func emptyCatalogError(filesystem fs.FS, root string) error {
	hasFeatureFile, err := containsFeatureFile(filesystem, root)
	if err != nil {
		return fmt.Errorf("discover behaviour features: %w", err)
	}
	if hasFeatureFile {
		return fmt.Errorf("feature below %q must contain at least one scenario", root)
	}
	return fmt.Errorf("no feature files found below %q", root)
}

func containsFeatureFile(filesystem fs.FS, root string) (bool, error) {
	found := false
	err := fs.WalkDir(filesystem, root, func(_ string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".feature" {
			found = true
		}
		return nil
	})
	if err != nil {
		return false, fmt.Errorf("walk %q: %w", root, err)
	}
	return found, nil
}

func pickleStepKind(stepType string) (StepKind, error) {
	switch stepType {
	case "Context":
		return GivenStep, nil
	case "Action":
		return WhenStep, nil
	case "Outcome":
		return ThenStep, nil
	default:
		return "", fmt.Errorf("unsupported Godog step type %q", stepType)
	}
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
