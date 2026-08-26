Feature: Automatic rule-change workflow triggers

  Scenario: A staged rule path automatically triggers the workflow
    Given a repository with a staged rule-bearing file
    When Badak Mini runs staged rule-change detection
    Then the command succeeds with the automatically triggered rules-propagation workflow

  Scenario: An ordinary staged path stays silent
    Given a repository with only an ordinary staged file
    When Badak Mini runs staged rule-change detection
    Then the command succeeds without output

  Scenario: A harness edit automatically triggers both workflows
    Given a pre-edit payload for a harness instruction file
    When Badak Mini runs hook rule-change detection
    Then the command succeeds with both workflow notices
