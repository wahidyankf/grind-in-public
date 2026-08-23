Feature: Harness capability parity

  Scenario: Harness capabilities match
    Given a repository whose harness capabilities match
    When Badak Mini runs capability-parity validation
    Then the command succeeds with the parity confirmation

  Scenario: A harness capability is missing
    Given a repository with a harness missing a shared subagent
    When Badak Mini runs capability-parity validation
    Then the command fails with the parity diagnostic
