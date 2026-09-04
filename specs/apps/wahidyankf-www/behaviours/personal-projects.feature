Feature: Personal projects page

  As a visitor to wahidyankf-web
  I want to browse a list of personal projects
  So that I can learn what I have built outside of employed work

  Background:
    Given the app is running

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Personal projects page renders the heading
  @integration-exempt
  Scenario: Personal projects page renders the heading
    When a visitor opens the personal projects page
    Then the H1 shows "Independent Projects"

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Personal projects page renders a search input
  @integration-exempt
  Scenario: Personal projects page renders a search input
    When a visitor opens the personal projects page
    Then a search input with placeholder "Search projects..." is visible

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Personal projects page lists at least one project card
  @integration-exempt
  Scenario: Personal projects page lists at least one project card
    When a visitor opens the personal projects page
    Then at least one project card is visible

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Each project card exposes external links where applicable
  @integration-exempt
  Scenario: Each project card exposes external links where applicable
    When a visitor opens the personal projects page
    Then every project card exposes a Repository, Website, or YouTube link where the project has that resource

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Each project card shows how long the project has been running
  @integration-exempt
  Scenario: Each project card shows how long the project has been running
    When a visitor opens the personal projects page
    Then every project card shows a duration next to its start date

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Each project card exposes clickable skill tags
  @integration-exempt
  Scenario: Each project card exposes clickable skill tags
    When a visitor opens the personal projects page
    Then every project card exposes at least one clickable skill tag

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Clicking a skill tag filters the project list
  @integration-exempt
  Scenario: Clicking a skill tag filters the project list
    When a visitor opens the personal projects page and clicks the "TypeScript" skill tag
    Then the URL becomes /personal-projects?search=TypeScript
