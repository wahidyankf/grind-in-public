Feature: Harness capability parity

  Scenario: Canonical harness content and adapters match
    Given a repository whose canonical harness contract matches
    When Badak Mini runs capability-parity validation
    Then the command succeeds with canonical counts and a digest

  Scenario: A canonical agent adapter is missing
    Given a canonical harness contract with a missing Codex agent adapter
    When Badak Mini runs capability-parity validation
    Then the command fails with a missing-agent-adapter diagnostic

  Scenario: Another instruction source competes with the canonical rules
    Given a canonical harness contract with an instruction overlay
    When Badak Mini runs capability-parity validation
    Then the command fails with an unexpected-instruction-source diagnostic

  Scenario: A thin skill adapter diverges from its canonical bundle
    Given a canonical harness contract with a stale Claude skill adapter
    When Badak Mini runs capability-parity validation
    Then the command fails with a skill-content-divergence diagnostic

  Scenario: A native agent adapter weakens a canonical denial
    Given a canonical harness contract with weakened opencode permissions
    When Badak Mini runs capability-parity validation
    Then the command fails with an agent-semantic-divergence diagnostic
