Feature: Accessibility

  As a visitor with accessibility needs using wahidyankf-web
  I want the site to follow WCAG 2.1 AA guidelines
  So that I can navigate and read content using assistive technologies

  Background:
    Given the app is running

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Home page has zero axe-core WCAG 2.1 AA violations
  @integration-exempt
  Scenario: Home page has zero axe-core WCAG 2.1 AA violations
    When a visitor opens the home page
    Then an axe-core scan against WCAG 2.1 AA reports zero violations

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / CV page has zero axe-core WCAG 2.1 AA violations
  @integration-exempt
  Scenario: CV page has zero axe-core WCAG 2.1 AA violations
    When a visitor opens the CV page
    Then an axe-core scan against WCAG 2.1 AA reports zero violations

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Every page has exactly one H1
  @integration-exempt
  Scenario: Every page has exactly one H1
    When a visitor opens any of the home, CV, or personal-projects pages
    Then each of those pages has exactly one H1 element

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / Interactive controls expose accessible names
  @integration-exempt
  Scenario: Interactive controls expose accessible names
    When a visitor opens the home page
    Then the theme toggle button exposes an aria-label
    And every navigation link exposes link text or an aria-label
