import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it } from "vitest";
import Home from "./page";

// Vitest globals are off, so RTL does not clean up the DOM between tests.
afterEach(cleanup);

describe("Home", () => {
  it("shows the Next.js logo", () => {
    render(<Home />);

    expect(screen.getByAltText("Next.js logo")).toBeTruthy();
  });

  it("shows the main heading", () => {
    const { container } = render(<Home />);

    expect(container.querySelector("h1")?.textContent).toContain(
      "To get started",
    );
  });
});
