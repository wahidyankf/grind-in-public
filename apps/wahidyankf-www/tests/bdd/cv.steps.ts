import path from "node:path";
import React from "react";
import { loadFeature, describeFeature } from "@amiceli/vitest-cucumber";
import { cleanup, render, screen } from "@testing-library/react";
import { expect, vi } from "vitest";
import { CvContent } from "@/features/cv/shell/cv-content";

const mockPush = vi.fn();
const mockReplace = vi.fn();
let mockSearchParams = new URLSearchParams();

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: mockPush, replace: mockReplace }),
  useSearchParams: () => mockSearchParams,
}));

vi.mock("@/features/app-shell/shell/navigation", () => ({
  Navigation: () =>
    React.createElement("div", { "data-testid": "navigation" }, "Navigation"),
}));

vi.stubGlobal("scrollTo", vi.fn());

// DOM cleanup runs after each scenario, allowing Then to inspect the component
// instance mounted by When.
const feature = await loadFeature(
  path.resolve(
    process.cwd(),
    "../../specs/apps/wahidyankf-www/behaviours/cv.feature",
  ),
);

describeFeature(feature, ({ Scenario, Background, AfterEachScenario }) => {
  AfterEachScenario(cleanup);
  Background(({ Given }) => {
    Given("the app is running", () => {
      vi.mocked(window.scrollTo).mockClear();
      mockSearchParams = new URLSearchParams();
      window.history.replaceState({}, "", "/cv");
    });
  });

  Scenario("CV renders the Curriculum Vitae heading", ({ When, Then }) => {
    When("a visitor opens the CV page", () => {
      window.history.replaceState({}, "", "/cv");
      render(React.createElement(CvContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/cv.feature:CV renders the Curriculum Vitae heading
    Then('the H1 shows "Curriculum Vitae"', () => {
      expect(
        screen.getByRole("heading", { level: 1, name: "Curriculum Vitae" }),
      ).toBeInTheDocument();
    });
  });

  Scenario("CV renders a search input", ({ When, Then }) => {
    When("a visitor opens the CV page", () => {
      window.history.replaceState({}, "", "/cv");
      render(React.createElement(CvContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/cv.feature:CV renders a search input
    Then(
      'a search input with placeholder "Search CV entries..." is visible',
      () => {
        expect(
          screen.getByPlaceholderText("Search CV entries..."),
        ).toBeInTheDocument();
      },
    );
  });

  Scenario("CV renders the Highlights section header", ({ When, Then }) => {
    When("a visitor opens the CV page", () => {
      window.history.replaceState({}, "", "/cv");
      render(React.createElement(CvContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/cv.feature:CV renders the Highlights section header
    Then('a "Highlights" section header is visible', () => {
      expect(
        screen.getByRole("heading", { name: "Highlights" }),
      ).toBeInTheDocument();
    });
  });

  Scenario(
    "CV cross-linked via scrollTop query scrolls into the entries",
    ({ When, Then }) => {
      When(
        'a visitor opens the CV page with search term "TypeScript" and scrollTop true',
        () => {
          window.history.replaceState(
            {},
            "",
            "/cv?search=TypeScript&scrollTop=true",
          );
          mockSearchParams = new URLSearchParams(
            "search=TypeScript&scrollTop=true",
          );
          render(React.createElement(CvContent));
        },
      );

      // @covers specs/apps/wahidyankf-www/behaviours/cv.feature:CV cross-linked via scrollTop query scrolls into the entries
      Then("the page scrolls past Highlights into the matching entries", () => {
        expect(window.scrollTo).toHaveBeenCalledWith(0, 0);
        expect(
          screen.getByRole("heading", { level: 1, name: "Curriculum Vitae" }),
        ).toBeInTheDocument();
        expect(
          screen.getByPlaceholderText("Search CV entries..."),
        ).toBeInTheDocument();
      });
    },
  );

  Scenario("CV offers a downloadable PDF", ({ When, Then }) => {
    When("a visitor opens the CV page", () => {
      window.history.replaceState({}, "", "/cv");
      render(React.createElement(CvContent));
    });

    // @covers specs/apps/wahidyankf-www/behaviours/cv.feature:CV offers a downloadable PDF
    Then(
      'a "Download CV (PDF)" link pointing at the generated PDF is visible',
      () => {
        const downloadLink = screen.getByRole("link", {
          name: /Download CV \(PDF\)/,
        });
        expect(downloadLink).toHaveAttribute(
          "href",
          "/wahidyankf-kresna-fridayoka-cv.pdf",
        );
        expect(downloadLink).toHaveAttribute("download");
      },
    );
  });
});
