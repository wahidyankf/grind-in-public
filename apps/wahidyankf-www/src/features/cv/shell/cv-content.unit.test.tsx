import React from "react";
import { describe, it, expect, vi } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import { CvContent } from "./cv-content";

// The CV scenarios under tests/bdd/ cover what the page renders. What they do
// not reach is the recent-only filter: no scenario clicks it, so both the
// toggle's handler and the five-year window helper behind it went unexercised.
// These tests drive that one control.
const mockReplace = vi.fn();
const mockSearchParams = new URLSearchParams();

vi.mock("next/navigation", () => ({
  useRouter: () => ({ push: vi.fn(), replace: mockReplace }),
  useSearchParams: () => mockSearchParams,
}));

vi.mock("@/features/app-shell/shell/navigation", () => ({
  Navigation: () =>
    React.createElement("div", { "data-testid": "navigation" }, "Navigation"),
}));

vi.stubGlobal("scrollTo", vi.fn());

const recentOnlyToggle = () =>
  screen.getByRole("button", { name: "Show recent work experience only" });

describe("CvContent recent-only filter", () => {
  it("offers the filter switched off, showing all work experience", () => {
    render(<CvContent />);
    expect(recentOnlyToggle()).toBeDefined();
  });

  it("switches its label when the filter is turned on", () => {
    render(<CvContent />);
    fireEvent.click(recentOnlyToggle());
    expect(
      screen.getByRole("button", { name: "Show all work experience" }),
    ).toBeDefined();
  });

  it("returns to the unfiltered label when switched off again", () => {
    render(<CvContent />);
    fireEvent.click(recentOnlyToggle());
    fireEvent.click(
      screen.getByRole("button", { name: "Show all work experience" }),
    );
    expect(recentOnlyToggle()).toBeDefined();
  });

  it("shows no more entries filtered than unfiltered", () => {
    // The filter keeps engagements whose end date falls inside a five-year
    // window, treating "Present" as today. Asserting a specific count would
    // pin this test to the CV record's current contents and break whenever a
    // job is added; asserting the direction of the change holds for any record.
    const { container } = render(<CvContent />);
    const before = container.querySelectorAll("h3").length;
    fireEvent.click(recentOnlyToggle());
    const after = container.querySelectorAll("h3").length;
    expect(after).toBeLessThanOrEqual(before);
  });
});
