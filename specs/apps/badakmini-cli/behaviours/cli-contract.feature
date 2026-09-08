Feature: Badak Mini CLI contract

  Scenario: Help is available outside a repository
    Given repository discovery would fail
    When Badak Mini runs with "--help"
    Then the command succeeds and prints usage

  Scenario: Command group help is available outside a repository
    Given repository discovery would fail
    When Badak Mini runs with "harness --help"
    Then the command succeeds and prints usage

  Scenario: A command name this CLI does not have is an invalid invocation
    Given repository discovery would fail
    When Badak Mini runs with "harness instruction-size validate"
    Then the command reports an invalid invocation
