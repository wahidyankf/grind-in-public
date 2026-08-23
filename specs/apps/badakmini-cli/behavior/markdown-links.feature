Feature: Repository Markdown links

  Scenario: Repository Markdown links resolve
    Given a repository whose tracked Markdown links resolve
    When Badak Mini runs Markdown-link validation
    Then the command succeeds with the link confirmation

  Scenario: A tracked Markdown link is broken
    Given a repository with a broken tracked Markdown link
    When Badak Mini runs Markdown-link validation
    Then the command fails with the missing-target diagnostic
