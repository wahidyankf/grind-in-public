# Product Requirements

## User Stories

As the repository owner, I want Badak Mini behaviour declared in canonical Gherkin so that changing a feature, scenario,
or step immediately revalidates every required test layer.

As a contributor, I want fast unit and static behaviour gates in `test:quick` so that most specification drift fails
before push without launching real processes.

As a maintainer, I want integration and process E2E adapters to consume the same corpus so that lower-level confidence
cannot hide a broken public command.

## Acceptance Scenarios

```gherkin
Feature: Command help

  Scenario: Help is available outside a repository
    Given repository discovery would fail
    When Badak Mini runs with "--help"
    Then the command succeeds and prints usage
```

```gherkin
Feature: Instruction size validation

  Scenario: Governance documents fit the word limit
    Given a repository whose governance documents fit the word limit
    When Badak Mini runs instruction-size validation
    Then the command succeeds with the word-limit confirmation

  Scenario: A governance document exceeds the word limit
    Given a repository with an oversized agent instruction file
    When Badak Mini runs instruction-size validation
    Then the command fails with the oversized document diagnostic
```

```gherkin
Feature: Markdown link validation

  Scenario: Repository Markdown links resolve
    Given a repository whose tracked Markdown links resolve
    When Badak Mini runs Markdown-link validation
    Then the command succeeds with the link confirmation

  Scenario: A tracked Markdown link is broken
    Given a repository with a broken tracked Markdown link
    When Badak Mini runs Markdown-link validation
    Then the command fails with the missing-target diagnostic
```

```gherkin
Feature: Harness capability parity

  Scenario: Harness capabilities match
    Given a repository whose harness capabilities match
    When Badak Mini runs capability-parity validation
    Then the command succeeds with the parity confirmation

  Scenario: A harness capability is missing
    Given a repository with a harness missing a shared subagent
    When Badak Mini runs capability-parity validation
    Then the command fails with the parity diagnostic
```

```gherkin
Feature: Rule change notices

  Scenario: A staged rule path announces the workflow
    Given a repository with a staged rule-bearing file
    When Badak Mini runs staged rule-change detection
    Then the command succeeds with the rules-propagation notice

  Scenario: An ordinary staged path stays silent
    Given a repository with only an ordinary staged file
    When Badak Mini runs staged rule-change detection
    Then the command succeeds without output

  Scenario: A harness edit announces both workflows
    Given a pre-edit payload for a harness instruction file
    When Badak Mini runs hook rule-change detection
    Then the command succeeds with both workflow notices
```

## In Scope

- Recursive executable Gherkin for the existing CLI contracts above.
- Shared bindings and driver contract with unit, integration, and E2E implementations.
- Static behaviour completeness and adapter parity.
- Role-based repository rules for applications, libraries, and dedicated E2E projects.

## Out of Scope

- Replacing focused package unit tests with Gherkin.
- Network, loopback servers, or external credentials.
- E2E scenarios owned by a library or by the E2E harness itself.
