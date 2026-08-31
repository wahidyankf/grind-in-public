import { describe, it, expect } from "vitest";
import sitemap from "./sitemap";

describe("sitemap", () => {
  it("publishes the three public routes and nothing else", () => {
    expect(sitemap().map((entry) => entry.url)).toEqual([
      "https://www.wahidyankf.com",
      "https://www.wahidyankf.com/cv",
      "https://www.wahidyankf.com/personal-projects",
    ]);
  });

  it("ranks the root above the two sections", () => {
    const [root, ...sections] = sitemap();
    expect(root.priority).toBe(1);
    expect(sections.map((entry) => entry.priority)).toEqual([0.8, 0.8]);
  });

  it("declares every route as monthly", () => {
    expect(
      sitemap().every((entry) => entry.changeFrequency === "monthly"),
    ).toBe(true);
  });
});
