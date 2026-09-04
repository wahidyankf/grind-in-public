Feature: Badak Mini CLI contract

  Scenario: Help is available outside a repository
    Given repository discovery would fail
    When Badak Mini runs with "--help"
    Then the command succeeds and prints usage

  Scenario: Command group help is available outside a repository
    Given repository discovery would fail
    When Badak Mini runs with "harness --help"
    Then the command succeeds and prints usage
