import path from "node:path";
import React from "react";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import { expect, vi } from "vitest";
import { HomeContent } from "@/features/home/shell/home-content";

const mockPush = vi.fn();
let mockSearchParams = new URLSearchParams();

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: mockPush }),
  useSearchParams: () => mockSearchParams,
}));

vi.mock("@/features/app-shell/shell/navigation", () => ({
  Navigation: () =>
    React.createElement("div", { "data-testid": "navigation" }, "Navigation"),
}));

// DOM cleanup runs after each scenario, so assertion steps inspect the state
// produced by the preceding action instead of remounting their subject.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/search.feature",
  ),
);

describeFeature(feature, ({ Scenario, Background, AfterEachScenario }) => {
  AfterEachScenario(cleanup);
  Background(({ Given }) => {
    Given("the app is running", () => {
      mockPush.mockClear();
      mockSearchParams = new URLSearchParams();
      window.history.replaceState({}, "", "/");
    });
  });

  Scenario(
    "Typing a term updates the URL query string",
    ({ When, And, Then }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
      });

      And('the visitor types "TypeScript" in the search input', () => {
        render(React.createElement(HomeContent));
        const input = screen.getByPlaceholderText(
          "Search skills, languages, or frameworks...",
        );
        fireEvent.change(input, { target: { value: "TypeScript" } });
      });

      // @covers specs/apps/wahidyankf-www/behaviours/search.feature:Typing a term updates the URL query string
      Then("the URL becomes /?search=TypeScript", () => {
        expect(mockPush).toHaveBeenCalledWith("/?search=TypeScript", {
          scroll: false,
        });
      });
    },
  );

  Scenario(
    "Matching content is highlighted with a yellow mark",
    ({ When, Then }) => {
      When(
        'a visitor opens the home page with search term "TypeScript"',
        () => {
          window.history.replaceState({}, "", "/?search=TypeScript");
          mockSearchParams = new URLSearchParams("search=TypeScript");
          render(React.createElement(HomeContent));
        },
      );

      // @covers specs/apps/wahidyankf-www/behaviours/search.feature:Matching content is highlighted with a yellow mark
      Then('the matching pill wraps "TypeScript" in a mark element', () => {
        const marks = Array.from(document.querySelectorAll("mark"));
        expect(
          marks.some((mark) => /TypeScript/i.test(mark.textContent ?? "")),
        ).toBe(true);
      });
    },
  );

  Scenario("Non-matching About Me shows a placeholder", ({ When, Then }) => {
    When('a visitor opens the home page with search term "NoSuchTerm"', () => {
      window.history.replaceState({}, "", "/?search=NoSuchTerm");
      mockSearchParams = new URLSearchParams("search=NoSuchTerm");
      render(React.createElement(HomeContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/search.feature:Non-matching About Me shows a placeholder
    Then(
      'the About Me card shows "No matching content in the About Me section."',
      () => {
        expect(
          screen.getByText("No matching content in the About Me section."),
        ).toBeInTheDocument();
      },
    );
  });

  Scenario(
    "Clicking a skill pill navigates to the CV with scrollTop",
    ({ When, And, Then }) => {
      When("a visitor opens the home page", () => {
        window.history.replaceState({}, "", "/");
      });

      And('the visitor clicks the "TypeScript" skill pill', () => {
        render(React.createElement(HomeContent));
        fireEvent.click(screen.getByRole("button", { name: "TypeScript" }));
      });

      // @covers specs/apps/wahidyankf-www/behaviours/search.feature:Clicking a skill pill navigates to the CV with scrollTop
      Then("the URL becomes /cv?search=TypeScript&scrollTop=true", () => {
        expect(mockPush).toHaveBeenCalledWith(
          "/cv?search=TypeScript&scrollTop=true",
        );
      });
    },
  );
});
