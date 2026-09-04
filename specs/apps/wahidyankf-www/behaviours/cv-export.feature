Feature: CV export

  As the owner of this site
  I want the CV export to write a real PDF and to fail loudly when it cannot
  So that a published CV is never silently missing or truncated

  # Exemption(e2e): Local configuration and filesystem behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / Generating the CV writes a PDF to the local filesystem
  @e2e-exempt
  Scenario: Generating the CV writes a PDF to the local filesystem
    Given the application CV record contains at least one entry
    When the CV export runs against a writable output directory
    Then a readable PDF file exists at the configured output path
    And the file begins with the PDF header bytes

  # Exemption(e2e): Local configuration and filesystem behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / Generating the CV reports an unwritable output location
  @e2e-exempt
  Scenario: Generating the CV reports an unwritable output location
    Given the configured CV output directory does not exist
    When the CV export runs
    Then the export fails with a message naming the output path
    And no partial file is left behind
