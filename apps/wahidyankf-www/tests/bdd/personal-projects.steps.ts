import path from "node:path";
import React from "react";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import {
  cleanup,
  render,
  screen,
  fireEvent,
  within,
} from "@testing-library/react";
import { expect, vi } from "vitest";
import { PersonalProjectsContent } from "@/features/personal-projects/shell/personal-projects-content";
import { projects } from "@/features/personal-projects/core/projects";

const mockPush = vi.fn();

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: mockPush }),
  useSearchParams: () => new URLSearchParams(),
}));

vi.mock("@/features/app-shell/shell/navigation", () => ({
  Navigation: () =>
    React.createElement("div", { "data-testid": "navigation" }, "Navigation"),
}));

// DOM cleanup runs after each scenario, allowing Then to inspect the component
// instance mounted by When while keeping scenarios isolated.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/personal-projects.feature",
  ),
);

describeFeature(feature, ({ Scenario, Background, AfterEachScenario }) => {
  AfterEachScenario(cleanup);
  Background(({ Given }) => {
    Given("the app is running", () => {
      mockPush.mockClear();
      window.history.replaceState({}, "", "/personal-projects");
    });
  });

  Scenario("Personal projects page renders the heading", ({ When, Then }) => {
    When("a visitor opens the personal projects page", () => {
      window.history.replaceState({}, "", "/personal-projects");
      render(React.createElement(PersonalProjectsContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Personal projects page renders the heading
    Then('the H1 shows "Independent Projects"', () => {
      expect(
        screen.getByRole("heading", { level: 1, name: "Independent Projects" }),
      ).toBeInTheDocument();
    });
  });

  Scenario(
    "Personal projects page renders a search input",
    ({ When, Then }) => {
      When("a visitor opens the personal projects page", () => {
        window.history.replaceState({}, "", "/personal-projects");
        render(React.createElement(PersonalProjectsContent));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Personal projects page renders a search input
      Then(
        'a search input with placeholder "Search projects..." is visible',
        () => {
          expect(
            screen.getByPlaceholderText("Search projects..."),
          ).toBeInTheDocument();
        },
      );
    },
  );

  Scenario(
    "Personal projects page lists at least one project card",
    ({ When, Then }) => {
      When("a visitor opens the personal projects page", () => {
        window.history.replaceState({}, "", "/personal-projects");
        render(React.createElement(PersonalProjectsContent));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Personal projects page lists at least one project card
      Then("at least one project card is visible", () => {
        const headings = screen.getAllByRole("heading", { level: 2 });
        expect(headings.length).toBeGreaterThan(0);
      });
    },
  );

  Scenario(
    "Each project card exposes external links where applicable",
    ({ When, Then }) => {
      When("a visitor opens the personal projects page", () => {
        window.history.replaceState({}, "", "/personal-projects");
        render(React.createElement(PersonalProjectsContent));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Each project card exposes external links where applicable
      Then(
        "every project card exposes a Repository, Website, or YouTube link where the project has that resource",
        () => {
          projects.forEach((project, index) => {
            const card = document.getElementById(`project-${index}`);
            expect(card).not.toBeNull();
            const cardScope = within(card as HTMLElement);
            for (const [resource, href] of Object.entries(project.links)) {
              const link = cardScope.getByRole("link", {
                name: new RegExp(`^${resource}$`, "i"),
              });
              expect(link).toHaveAttribute("href", href);
              expect(link).toHaveAttribute("target", "_blank");
              expect(link).toHaveAttribute("rel", "noopener noreferrer");
            }
          });
        },
      );
    },
  );

  Scenario(
    "Each project card shows how long the project has been running",
    ({ When, Then }) => {
      When("a visitor opens the personal projects page", () => {
        window.history.replaceState({}, "", "/personal-projects");
        render(React.createElement(PersonalProjectsContent));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Each project card shows how long the project has been running
      Then("every project card shows a duration next to its start date", () => {
        projects.forEach((_, index) => {
          const card = document.getElementById(`project-${index}`);
          expect(card).not.toBeNull();
          const cardScope = within(card as HTMLElement);
          const durationMatches = cardScope.getAllByText((_content, element) =>
            /\(\d+\s+(year|month)/i.test(element?.textContent ?? ""),
          );
          expect(durationMatches.length).toBeGreaterThan(0);
        });
      });
    },
  );

  Scenario(
    "Each project card exposes clickable skill tags",
    ({ When, Then }) => {
      When("a visitor opens the personal projects page", () => {
        window.history.replaceState({}, "", "/personal-projects");
        render(React.createElement(PersonalProjectsContent));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Each project card exposes clickable skill tags
      Then(
        "every project card exposes at least one clickable skill tag",
        () => {
          projects.forEach((_, index) => {
            const card = document.getElementById(`project-${index}`);
            expect(card).not.toBeNull();
            const cardScope = within(card as HTMLElement);
            expect(cardScope.getAllByRole("button").length).toBeGreaterThan(0);
          });
        },
      );
    },
  );

  Scenario(
    "Clicking a skill tag filters the project list",
    ({ When, Then }) => {
      When(
        'a visitor opens the personal projects page and clicks the "TypeScript" skill tag',
        () => {
          render(React.createElement(PersonalProjectsContent));
          fireEvent.click(
            screen.getAllByRole("button", { name: "TypeScript" })[0],
          );
        },
      );

      // @covers specs/apps/wahidyankf-www/behaviours/personal-projects.feature:Clicking a skill tag filters the project list
      Then("the URL becomes /personal-projects?search=TypeScript", () => {
        expect(mockPush).toHaveBeenCalledWith(
          "/personal-projects?search=TypeScript",
          { scroll: false },
        );
      });
    },
  );
});
