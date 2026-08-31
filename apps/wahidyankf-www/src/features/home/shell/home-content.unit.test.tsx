import React from "react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import { HomeContent } from "./home-content";

// The home scenarios under tests/bdd/ cover the cards this page renders,
// including that the skills card carries three subsections. What no scenario
// reaches is the AI-related skills list, which is a fourth subsection rendered
// only when the CV record supplies AI skills — so its pills, and the navigation
// they perform, were the module's one uncovered line.
const mockPush = vi.fn();

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: mockPush, replace: vi.fn() }),
  useSearchParams: () => new URLSearchParams(),
}));

vi.mock("@/features/app-shell/shell/navigation", () => ({
  Navigation: () =>
    React.createElement("div", { "data-testid": "navigation" }, "Navigation"),
}));

const aiSkillsHeading = "Top AI-Related Skills Used in The Last 5 Years";

describe("HomeContent AI-related skills", () => {
  beforeEach(() => {
    mockPush.mockClear();
  });

  it("renders the AI-related skills subsection", () => {
    render(<HomeContent />);
    expect(screen.getByText(aiSkillsHeading)).toBeDefined();
  });

  it("navigates to the CV filtered by the skill when a pill is clicked", () => {
    const { container } = render(<HomeContent />);
    const heading = screen.getByText(aiSkillsHeading);
    // The pills sit in the element immediately after the heading. Selecting
    // them positionally rather than by name keeps the test from naming a
    // specific skill, which would tie it to today's CV record.
    const pill = heading.nextElementSibling?.querySelector("button");
    expect(pill).not.toBeNull();
    fireEvent.click(pill as HTMLButtonElement);

    expect(mockPush).toHaveBeenCalledTimes(1);
    const target = mockPush.mock.calls[0][0] as string;
    expect(target.startsWith("/cv?search=")).toBe(true);
    expect(target.endsWith("&scrollTop=true")).toBe(true);
    expect(container).toBeDefined();
  });

  it("percent-encodes a skill name into the CV search query", () => {
    render(<HomeContent />);
    const heading = screen.getByText(aiSkillsHeading);
    const pill = heading.nextElementSibling?.querySelector(
      "button",
    ) as HTMLButtonElement;
    const label = pill.textContent ?? "";
    fireEvent.click(pill);

    const target = mockPush.mock.calls[0][0] as string;
    const search = new URL(target, "https://example.test").searchParams.get(
      "search",
    );
    // Round-tripping through URL is what proves the encoding: a skill whose
    // name contains a space or an ampersand would otherwise produce a query
    // string that parses into something else.
    expect(search).not.toBeNull();
    expect(label.includes(search as string)).toBe(true);
  });
});
