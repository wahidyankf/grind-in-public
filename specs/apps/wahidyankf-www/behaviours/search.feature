Feature: Search

  As a visitor to wahidyankf-web
  I want to filter portfolio content by keyword
  So that I can find specific skills, languages, or frameworks quickly

  Background:
    Given the app is running

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Typing a term updates the URL query string
  @integration-exempt
  Scenario: Typing a term updates the URL query string
    When a visitor opens the home page
    And the visitor types "TypeScript" in the search input
    Then the URL becomes /?search=TypeScript

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Matching content is highlighted with a yellow mark
  @integration-exempt
  Scenario: Matching content is highlighted with a yellow mark
    When a visitor opens the home page with search term "TypeScript"
    Then the matching pill wraps "TypeScript" in a mark element

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Non-matching About Me shows a placeholder
  @integration-exempt
  Scenario: Non-matching About Me shows a placeholder
    When a visitor opens the home page with search term "NoSuchTerm"
    Then the About Me card shows "No matching content in the About Me section."

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Clicking a skill pill navigates to the CV with scrollTop
  @integration-exempt
  Scenario: Clicking a skill pill navigates to the CV with scrollTop
    When a visitor opens the home page
    And the visitor clicks the "TypeScript" skill pill
    Then the URL becomes /cv?search=TypeScript&scrollTop=true
