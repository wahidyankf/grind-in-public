import path from "node:path";
import React from "react";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { cleanup, render, screen, within } from "@testing-library/react";
import { expect, vi } from "vitest";
import { HomeContent } from "@/features/home/shell/home-content";

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: vi.fn() }),
  useSearchParams: () => new URLSearchParams(),
}));

vi.mock("@/features/app-shell/shell/navigation", () => ({
  Navigation: () =>
    React.createElement("div", { "data-testid": "navigation" }, "Navigation"),
}));

// Vitest Cucumber registers each step as a test inside one scenario describe.
// The behaviour setup deliberately defers DOM cleanup to AfterEachScenario so
// When can invoke the component and Then can observe that exact render.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/home.feature",
  ),
);

describeFeature(feature, ({ Scenario, Background, AfterEachScenario }) => {
  AfterEachScenario(cleanup);
  Background(({ Given }) => {
    Given("the app is running", () => {
      window.history.replaceState({}, "", "/");
    });
  });

  Scenario("Home renders the welcome heading", ({ When, Then }) => {
    When("a visitor opens the home page", () => {
      window.history.replaceState({}, "", "/");
      render(React.createElement(HomeContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/home.feature:Home renders the welcome heading
    Then('the H1 shows "Welcome to My Portfolio"', () => {
      expect(
        screen.getByRole("heading", {
          level: 1,
          name: "Welcome to My Portfolio",
        }),
      ).toBeInTheDocument();
    });
  });

  Scenario("Home renders the About Me card", ({ When, Then }) => {
    When("a visitor opens the home page", () => {
      window.history.replaceState({}, "", "/");
      render(React.createElement(HomeContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/home.feature:Home renders the About Me card
    Then("an About Me card is visible", () => {
      expect(
        screen.getByRole("heading", { name: "About Me" }),
      ).toBeInTheDocument();
    });
  });

  Scenario(
    "Home renders the Skills & Expertise card with three subsections",
    ({ When, Then, And }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
        render(React.createElement(HomeContent));
      });

      Then("a Skills & Expertise card is visible", () => {
        expect(
          screen.getByRole("heading", { name: "Skills & Expertise" }),
        ).toBeInTheDocument();
      });

      And(
        'the card has a "Top Skills Used in The Last 5 Years" subsection',
        () => {
          expect(
            screen.getByText("Top Skills Used in The Last 5 Years"),
          ).toBeInTheDocument();
        },
      );

      And(
        'the card has a "Top Programming Languages Used in The Last 5 Years" subsection',
        () => {
          expect(
            screen.getByText(
              "Top Programming Languages Used in The Last 5 Years",
            ),
          ).toBeInTheDocument();
        },
      );

      // @covers specs/apps/wahidyankf-www/behaviours/home.feature:Home renders the Skills & Expertise card with three subsections
      And(
        'the card has a "Top Frameworks & Libraries Used in The Last 5 Years" subsection',
        () => {
          expect(
            screen.getByText(
              "Top Frameworks & Libraries Used in The Last 5 Years",
            ),
          ).toBeInTheDocument();
        },
      );
    },
  );

  Scenario(
    "Home renders the Quick Links card with two internal links",
    ({ When, Then, And }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
        render(React.createElement(HomeContent));
      });

      Then("a Quick Links card is visible", () => {
        expect(
          screen.getByRole("heading", { name: "Quick Links" }),
        ).toBeInTheDocument();
      });

      And('the card contains a "View My CV" link to /cv', () => {
        expect(
          screen.getByRole("link", { name: "View My CV" }),
        ).toHaveAttribute("href", "/cv");
      });

      // @covers specs/apps/wahidyankf-www/behaviours/home.feature:Home renders the Quick Links card with two internal links
      And(
        'the card contains a "Browse My Independent Projects" link to /personal-projects',
        () => {
          expect(
            screen.getByRole("link", {
              name: "Browse My Independent Projects",
            }),
          ).toHaveAttribute("href", "/personal-projects");
        },
      );
    },
  );

  Scenario(
    "Home renders the Connect With Me card with five external links",
    ({ When, Then, And }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
        render(React.createElement(HomeContent));
      });

      Then("a Connect With Me card is visible", () => {
        expect(
          screen.getByRole("heading", { name: "Connect With Me" }),
        ).toBeInTheDocument();
      });

      // @covers specs/apps/wahidyankf-www/behaviours/home.feature:Home renders the Connect With Me card with five external links
      And(
        "the card has Github, GithubOrg, Linkedin, Website, and Email links",
        () => {
          const heading = screen.getByRole("heading", {
            name: "Connect With Me",
          });
          const section = heading.closest("section");
          expect(section).not.toBeNull();
          for (const name of [
            "Github",
            "GithubOrg",
            "Linkedin",
            "Website",
            "Email",
          ]) {
            expect(
              within(section as HTMLElement).getByRole("link", { name }),
            ).toBeInTheDocument();
          }
        },
      );
    },
  );
});
