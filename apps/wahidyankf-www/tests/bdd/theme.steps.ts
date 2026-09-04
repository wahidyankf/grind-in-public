import path from "node:path";
import React from "react";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import { expect } from "vitest";
import { ThemeToggle } from "@/features/ui/shell";

// DOM cleanup runs after each scenario. Navigation and reload actions perform
// any remount they imply; Then only observes the resulting DOM or persistence.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/theme.feature",
  ),
);

describeFeature(feature, ({ Scenario, Background, AfterEachScenario }) => {
  AfterEachScenario(cleanup);
  Background(({ Given }) => {
    Given("the app is running", () => {
      // Every scenario starts from a clean slate: no persisted theme choice
      // and no "light-theme" class on the document root.
      window.localStorage.clear();
      document.documentElement.classList.remove("light-theme");
    });
  });

  Scenario("Default theme is dark", ({ When, Then, And }) => {
    When("a visitor opens the home page for the first time", () => {
      window.history.replaceState({}, "", "/");
      render(React.createElement(ThemeToggle));
    });

    Then('the html element has no "light-theme" class', () => {
      expect(document.documentElement.classList.contains("light-theme")).toBe(
        false,
      );
    });

    // @covers specs/apps/wahidyankf-www/behaviours/theme.feature:Default theme is dark
    And('the theme toggle aria-label is "Switch to light theme"', () => {
      expect(
        screen.getByRole("button", { name: "Switch to light theme" }),
      ).toBeInTheDocument();
    });
  });

  Scenario(
    "Clicking the toggle switches to light theme",
    ({ When, And, Then }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
      });

      And("the visitor clicks the theme toggle", () => {
        render(React.createElement(ThemeToggle));
        fireEvent.click(
          screen.getByRole("button", { name: "Switch to light theme" }),
        );
      });

      Then('the html element has the "light-theme" class', () => {
        expect(document.documentElement.classList.contains("light-theme")).toBe(
          true,
        );
      });

      // @covers specs/apps/wahidyankf-www/behaviours/theme.feature:Clicking the toggle switches to light theme
      And('the theme toggle aria-label is "Switch to dark theme"', () => {
        expect(
          screen.getByRole("button", { name: "Switch to dark theme" }),
        ).toBeInTheDocument();
      });
    },
  );

  Scenario("Theme persists across navigation", ({ When, And, Then }) => {
    When("a visitor opens the home page", () => {
      window.history.replaceState({}, "", "/");
    });

    And("the visitor clicks the theme toggle", () => {
      render(React.createElement(ThemeToggle));
      fireEvent.click(
        screen.getByRole("button", { name: "Switch to light theme" }),
      );
    });

    And("the visitor navigates to the CV page", () => {
      cleanup();
      window.history.replaceState({}, "", "/cv");
      render(React.createElement(ThemeToggle));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/theme.feature:Theme persists across navigation
    Then('the html element still has the "light-theme" class', () => {
      expect(document.documentElement.classList.contains("light-theme")).toBe(
        true,
      );
    });
  });

  Scenario("Theme choice persists across reloads", ({ When, And, Then }) => {
    When("a visitor opens the home page", () => {
      window.history.replaceState({}, "", "/");
    });

    And("the visitor clicks the theme toggle", () => {
      render(React.createElement(ThemeToggle));
      fireEvent.click(
        screen.getByRole("button", { name: "Switch to light theme" }),
      );
    });

    And("the visitor reloads the page", () => {
      // A real reload tears down the entire document and re-runs the app
      // from scratch — simulate that by clearing the DOM class the toggle
      // set. Only the localStorage-backed choice should survive.
      cleanup();
      document.documentElement.classList.remove("light-theme");
      render(React.createElement(ThemeToggle));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/theme.feature:Theme choice persists across reloads
    Then('the html element still has the "light-theme" class', () => {
      expect(document.documentElement.classList.contains("light-theme")).toBe(
        true,
      );
    });
  });
});
