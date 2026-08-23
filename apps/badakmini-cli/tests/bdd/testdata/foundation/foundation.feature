Feature: BDD foundation

  Scenario: A fixture executes through the shared suite
    Given a foundation fixture
    When the fixture runs
    Then the fixture succeeds
