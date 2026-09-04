Feature: Governance instruction size

  Scenario: Governance documents fit the word limit
    Given a repository whose governance documents fit the word limit
    When Badak Mini runs instruction-size validation
    Then the command succeeds with the word-limit confirmation

  Scenario: A governance document exceeds the word limit
    Given a repository with an oversized agent instruction file
    When Badak Mini runs instruction-size validation
    Then the command fails with the oversized document diagnostic
