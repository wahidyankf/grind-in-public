Feature: Runtime listener port resolution
  As an operator starting a service
  I want one precedence rule for the listener port across every app in the repository
  So that the same flag and the same prefixed variable move the port in local development and in a container alike

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / The CLI flag outranks every other source
  @e2e-exempt
  Scenario: The CLI flag outranks every other source
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment sets "WAHIDYANKF_WWW_PORT" to "4000"
    When the port resolves with a "--port" flag of "5000"
    Then the resolved port is 5000

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / The prefixed variable outranks the fallback
  @e2e-exempt
  Scenario: The prefixed variable outranks the fallback
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment sets "WAHIDYANKF_WWW_PORT" to "4000"
    When the port resolves with no "--port" flag
    Then the resolved port is 4000

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / The fallback applies when nothing else supplies a port
  @e2e-exempt
  Scenario: The fallback applies when nothing else supplies a port
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment does not set "WAHIDYANKF_WWW_PORT"
    When the port resolves with no "--port" flag
    Then the resolved port is 3100

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / A bare PORT variable never moves the listener
  @e2e-exempt
  Scenario: A bare PORT variable never moves the listener
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment sets "PORT" to "4000"
    And the environment does not set "WAHIDYANKF_WWW_PORT"
    When the port resolves with no "--port" flag
    Then the resolved port is 3100

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / A blank value at a tier falls through to the next tier
  @e2e-exempt
  Scenario Outline: A blank value at a tier falls through to the next tier
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment sets "WAHIDYANKF_WWW_PORT" to "<envValue>"
    When the port resolves with a "--port" flag of "<flagValue>"
    Then the resolved port is <expected>

    Examples:
      | flagValue | envValue | expected |
      |           | 4000     | 4000     |
      |           |          | 3100     |
      | 5000      |          | 5000     |

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / A present but malformed port fails loudly instead of falling through
  @e2e-exempt
  Scenario Outline: A present but malformed port fails loudly instead of falling through
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment does not set "WAHIDYANKF_WWW_PORT"
    When the port resolves with a "--port" flag of "<flagValue>"
    Then resolution throws, naming "--port" and the valid range

    Examples:
      | flagValue |
      | 0         |
      | 65536     |
      | abc       |
      | 3100abc   |
      | 31.5      |
      | -1        |
      | 0x10      |
      | 0b1010    |
      | 1e3       |
      | +3100     |

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / A malformed prefixed variable names that variable in the error
  @e2e-exempt
  Scenario: A malformed prefixed variable names that variable in the error
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 3100
    And the environment sets "WAHIDYANKF_WWW_PORT" to "not-a-port"
    When the port resolves with no "--port" flag
    Then resolution throws, naming "WAHIDYANKF_WWW_PORT" and the valid range

  # Exemption(e2e): Local configuration behaviour is not observable through the browser boundary; alternative-proof: wahidyankf-www:test:integration / An out-of-range compiled-in fallback is caught at startup
  @e2e-exempt
  Scenario: An out-of-range compiled-in fallback is caught at startup
    Given the app declares the prefixed variable "WAHIDYANKF_WWW_PORT" with fallback 70000
    And the environment does not set "WAHIDYANKF_WWW_PORT"
    When the port resolves with no "--port" flag
    Then resolution throws, naming "WAHIDYANKF_WWW_PORT" and the valid range
