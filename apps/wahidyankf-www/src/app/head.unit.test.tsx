import { describe, it, expect } from "vitest";
import { render } from "@testing-library/react";
import Head from "./head";

// React 19 hoists `<link>` elements into `document.head` itself, so the render
// container these assertions would normally read is empty and the elements have
// to be looked for where they actually land. That hoisting is also why the
// component works at all without Next.js around it. The links are not removed
// by hand between tests: they belong to React, and the setup file's
// `afterEach(cleanup)` unmounts the tree that owns them. Deleting them
// directly makes React's own removal fail on a node that is already gone.
describe("Head", () => {
  it("declares a favicon in both ICO and PNG form", () => {
    render(<Head />);
    const icons = Array.from(
      document.head.querySelectorAll<HTMLLinkElement>('link[rel="icon"]'),
    );
    expect(icons.map((link) => link.getAttribute("href"))).toEqual([
      "/favicon.ico",
      "/favicon.png",
    ]);
    expect(icons[1].getAttribute("type")).toBe("image/png");
  });

  it("declares an apple touch icon for the iOS home screen", () => {
    render(<Head />);
    const touch = document.head.querySelector<HTMLLinkElement>(
      'link[rel="apple-touch-icon"]',
    );
    expect(touch?.getAttribute("href")).toBe("/favicon.png");
  });
});
