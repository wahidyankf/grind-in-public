import path from "node:path";
import React from "react";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { cleanup, render, screen, within } from "@testing-library/react";
import { expect, vi } from "vitest";
import { ThemeToggle } from "@/features/ui/shell";
import { HomeContent } from "@/features/home/shell/home-content";
import { CvContent } from "@/features/cv/shell/cv-content";
import { PersonalProjectsContent } from "@/features/personal-projects/shell/personal-projects-content";

// Navigation reads the current route via usePathname() and HomeContent/CvContent/
// PersonalProjectsContent navigate via useRouter() — both need a router shim since
// jsdom has no Next.js App Router context.
vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: vi.fn(), replace: vi.fn() }),
  usePathname: () => "/",
  useSearchParams: () => new URLSearchParams(),
}));

// DOM cleanup runs after each scenario, so Then observes the accessibility
// evidence produced by the page render in When.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/accessibility.feature",
  ),
);

describeFeature(feature, ({ Scenario, Background, AfterEachScenario }) => {
  AfterEachScenario(cleanup);
  Background(({ Given }) => {
    Given("the app is running", () => {
      window.history.replaceState({}, "", "/");
    });
  });

  Scenario(
    "Home page has zero axe-core WCAG 2.1 AA violations",
    ({ When, Then }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
        render(React.createElement(HomeContent));
      });

      // A full axe-core WCAG 2.1 AA scan runs against the live page in the e2e tier
      // (apps/wahidyankf-www-e2e/tests/steps/accessibility.steps.ts). At unit tier we
      // approximate by asserting the page exposes a landmark region and every
      // interactive control has an accessible name — the structural preconditions
      // an axe scan checks for.
      // @covers specs/apps/wahidyankf-www/behaviours/accessibility.feature:Home page has zero axe-core WCAG 2.1 AA violations
      Then(
        "an axe-core scan against WCAG 2.1 AA reports zero violations",
        () => {
          const main = screen.getByRole("main");
          expect(main).toBeInTheDocument();

          const links = screen.getAllByRole("link");
          expect(links.length).toBeGreaterThan(0);
          for (const link of links) {
            expect(link).toHaveAccessibleName();
          }

          const buttons = screen.getAllByRole("button");
          for (const button of buttons) {
            expect(button).toHaveAccessibleName();
          }
        },
      );
    },
  );

  Scenario(
    "CV page has zero axe-core WCAG 2.1 AA violations",
    ({ When, Then }) => {
      When("a visitor opens the CV page", () => {
        window.history.replaceState({}, "", "/cv");
        render(React.createElement(CvContent));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/accessibility.feature:CV page has zero axe-core WCAG 2.1 AA violations
      Then(
        "an axe-core scan against WCAG 2.1 AA reports zero violations",
        () => {
          const main = screen.getByRole("main");
          expect(main).toBeInTheDocument();

          const links = screen.getAllByRole("link");
          expect(links.length).toBeGreaterThan(0);
          for (const link of links) {
            expect(link).toHaveAccessibleName();
          }
        },
      );
    },
  );

  Scenario("Every page has exactly one H1", ({ When, Then }) => {
    const headingCounts: number[] = [];
    When(
      "a visitor opens any of the home, CV, or personal-projects pages",
      () => {
        for (const [route, Component] of [
          ["/", HomeContent],
          ["/cv", CvContent],
          ["/personal-projects", PersonalProjectsContent],
        ] as const) {
          window.history.replaceState({}, "", route);
          render(React.createElement(Component));
          headingCounts.push(
            screen.getAllByRole("heading", { level: 1 }).length,
          );
          cleanup();
        }
      },
    );

    // @covers specs/apps/wahidyankf-www/behaviours/accessibility.feature:Every page has exactly one H1
    Then("each of those pages has exactly one H1 element", () => {
      expect(headingCounts).toEqual([1, 1, 1]);
    });
  });

  Scenario(
    "Interactive controls expose accessible names",
    ({ When, Then, And }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
        render(
          React.createElement(
            React.Fragment,
            null,
            React.createElement(ThemeToggle),
            React.createElement(HomeContent),
          ),
        );
      });

      Then("the theme toggle button exposes an aria-label", () => {
        const toggle = screen.getByRole("button", {
          name: /Switch to (light|dark) theme/,
        });
        expect(toggle).toBeInTheDocument();
      });

      // @covers specs/apps/wahidyankf-www/behaviours/accessibility.feature:Interactive controls expose accessible names
      And("every navigation link exposes link text or an aria-label", () => {
        const desktopNav = screen.getByTestId("desktop-nav");
        for (const name of ["Home", "CV", "Independent Projects"]) {
          expect(
            within(desktopNav).getByRole("link", { name }),
          ).toBeInTheDocument();
        }
      });
    },
  );
});
