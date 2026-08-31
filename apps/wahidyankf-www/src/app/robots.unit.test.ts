import { describe, it, expect } from "vitest";
import robots from "./robots";

// This module is already reached by the static-filterable-routes scenarios,
// which assert that crawlers receive discovery directives at all. These tests
// assert what the directives say, so a change to the allow rule or to the
// sitemap URL fails here rather than passing a scenario that only checks the
// route responds.
describe("robots", () => {
  it("allows every crawler to reach every path", () => {
    expect(robots().rules).toEqual([{ userAgent: "*", allow: "/" }]);
  });

  it("names the sitemap by absolute URL", () => {
    // A crawler resolves this field against the origin, not against the file it
    // read, so a relative path here would point somewhere that does not exist.
    expect(robots().sitemap).toBe("https://www.wahidyankf.com/sitemap.xml");
  });
});
