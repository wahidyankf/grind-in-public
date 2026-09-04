import { createBdd } from "playwright-bdd";
import { expect, type Page } from "@playwright/test";

const { When, Then } = createBdd();
const observedScrollTopUrls = new WeakMap<Page, string>();

When('the visitor types "TypeScript" in the search input', async ({ page }) => {
  const input = page.getByPlaceholder(
    /Search skills, languages, or frameworks/i,
  );
  await input.fill("TypeScript");
});

// @covers specs/apps/wahidyankf-www/behaviours/search.feature:Typing a term updates the URL query string
Then("the URL becomes \\/?search=TypeScript", async ({ page }) => {
  await expect(page).toHaveURL(/\?search=TypeScript$/);
});

// @covers specs/apps/wahidyankf-www/behaviours/search.feature:Clicking a skill pill navigates to the CV with scrollTop
Then(
  "the URL becomes \\/cv?search=TypeScript&scrollTop=true",
  async ({ page }) => {
    expect(observedScrollTopUrls.get(page)).toMatch(
      /\/cv\?search=TypeScript&scrollTop=true$/,
    );
  },
);

When(
  'a visitor opens the home page with search term "TypeScript"',
  async ({ page }) => {
    await page.goto("/?search=TypeScript");
    await page.waitForLoadState("load");
  },
);

// @covers specs/apps/wahidyankf-www/behaviours/search.feature:Matching content is highlighted with a yellow mark
Then(
  'the matching pill wraps "TypeScript" in a mark element',
  async ({ page }) => {
    const mark = page
      .locator("mark")
      .filter({ hasText: /TypeScript/i })
      .first();
    await expect(mark).toBeVisible();
  },
);

When(
  'a visitor opens the home page with search term "NoSuchTerm"',
  async ({ page }) => {
    await page.goto("/?search=NoSuchTerm");
    await page.waitForLoadState("load");
  },
);

// @covers specs/apps/wahidyankf-www/behaviours/search.feature:Non-matching About Me shows a placeholder
Then(
  'the About Me card shows "No matching content in the About Me section."',
  async ({ page }) => {
    await expect(
      page.getByText(/No matching content in the About Me section\./i),
    ).toBeVisible();
  },
);

When('the visitor clicks the "TypeScript" skill pill', async ({ page }) => {
  await page.goto("/?search=TypeScript");
  await page.waitForLoadState("load");
  const pill = page
    .getByRole("button")
    .filter({ hasText: /^TypeScript$/ })
    .first();
  const navigatedWithScrollTop = page.waitForEvent(
    "framenavigated",
    (frame) =>
      frame === page.mainFrame() &&
      frame.url().endsWith("/cv?search=TypeScript&scrollTop=true"),
  );
  await pill.click();
  const frame = await navigatedWithScrollTop;
  observedScrollTopUrls.set(page, frame.url());
});
