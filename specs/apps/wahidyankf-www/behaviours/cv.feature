Feature: CV page

  As a visitor to wahidyankf-web
  I want the CV page to show my career and education history
  So that I can browse my professional background

  Background:
    Given the app is running

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / CV renders the Curriculum Vitae heading
  @integration-exempt
  Scenario: CV renders the Curriculum Vitae heading
    When a visitor opens the CV page
    Then the H1 shows "Curriculum Vitae"

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / CV renders a search input
  @integration-exempt
  Scenario: CV renders a search input
    When a visitor opens the CV page
    Then a search input with placeholder "Search CV entries..." is visible

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / CV renders the Highlights section header
  @integration-exempt
  Scenario: CV renders the Highlights section header
    When a visitor opens the CV page
    Then a "Highlights" section header is visible

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / CV cross-linked via scrollTop query scrolls into the entries
  @integration-exempt
  Scenario: CV cross-linked via scrollTop query scrolls into the entries
    When a visitor opens the CV page with search term "TypeScript" and scrollTop true
    Then the page scrolls past Highlights into the matching entries

  # Exemption(integration): Browser-rendered behaviour has no local resource boundary; alternative-proof: wahidyankf-www-e2e:test:e2e / CV offers a downloadable PDF
  @integration-exempt
  Scenario: CV offers a downloadable PDF
    When a visitor opens the CV page
    Then a "Download CV (PDF)" link pointing at the generated PDF is visible
