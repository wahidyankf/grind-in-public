import { expect, test } from "@playwright/test";

test("home page shows the logo and main heading", async ({ page }) => {
  await page.goto("/");

  await expect(page.getByAltText("Next.js logo")).toBeVisible();
  await expect(page.locator("h1")).toContainText("To get started");
});
