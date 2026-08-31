import { createBdd } from "playwright-bdd";
import { expect } from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";

const { When, Then } = createBdd();

// @covers specs/apps/wahidyankf-www/behavior/accessibility.feature:Home page has zero axe-core WCAG 2.1 AA violations
// @covers specs/apps/wahidyankf-www/behavior/accessibility.feature:CV page has zero axe-core WCAG 2.1 AA violations
Then(
  "an axe-core scan against WCAG 2.1 AA reports zero violations",
  async ({ page }) => {
    const results = await new AxeBuilder({ page })
      .withTags(["wcag2a", "wcag2aa"])
      .analyze();
    expect(results.violations).toEqual([]);
  },
);

When(
  "a visitor opens any of the home, CV, or personal-projects pages",
  async ({ page }) => {
    await page.goto("/");
    await page.waitForLoadState("load");
  },
);

// @covers specs/apps/wahidyankf-www/behavior/accessibility.feature:Every page has exactly one H1
Then("each of those pages has exactly one H1 element", async ({ page }) => {
  for (const path of ["/", "/cv", "/personal-projects"]) {
    await page.goto(path);
    await page.waitForLoadState("load");
    const h1Count = await page.locator("h1").count();
    expect(h1Count, `expected exactly one H1 on ${path}`).toBe(1);
  }
});

Then("the theme toggle button exposes an aria-label", async ({ page }) => {
  const toggle = page.getByRole("button", {
    name: /Switch to (light|dark) theme/,
  });
  await expect(toggle).toBeVisible();
});

// @covers specs/apps/wahidyankf-www/behavior/accessibility.feature:Interactive controls expose accessible names
Then(
  "every navigation link exposes link text or an aria-label",
  async ({ page }) => {
    // These are the three labels `Navigation` renders, not the three route
    // names. The third route is `/personal-projects` but its link reads
    // "Independent Projects", and the sibling binding of this same scenario in
    // `apps/wahidyankf-www/tests/bdd/accessibility.steps.ts` asserts the same
    // three. The scenario itself names no label, so the list belongs to the
    // binding and has to track what the component actually renders.
    const linkNames = ["Home", "CV", "Independent Projects"];
    for (const name of linkNames) {
      await expect(
        page.getByRole("link", { name: new RegExp(`^${name}$`) }).first(),
      ).toBeVisible();
    }
  },
);
