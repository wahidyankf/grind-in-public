Feature: Rule-change workflow notices

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
