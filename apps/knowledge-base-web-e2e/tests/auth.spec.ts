import { randomUUID } from "node:crypto";
import { expect, test, type Page } from "@playwright/test";

const PASSWORD = "correct horse battery";

// Each test uses its own synthetic email, so runs do not clash.
function newEmail(): string {
  return `user-${randomUUID()}@example.test`;
}

async function registerThroughUi(page: Page, email: string) {
  await page.goto("/register");
  await page.getByLabel("Email").fill(email);
  await page.getByLabel("Password").fill(PASSWORD);
  await page.getByRole("button", { name: "Register" }).click();
}

async function logInThroughUi(page: Page, email: string, password: string) {
  await page.goto("/login");
  await page.getByLabel("Email").fill(email);
  await page.getByLabel("Password").fill(password);
  await page.getByRole("button", { name: "Log in" }).click();
}

test("register sends the user to the log in page", async ({ page }) => {
  await registerThroughUi(page, newEmail());

  await expect(page.getByRole("heading", { name: "Log in" })).toBeVisible();
});

test("registering the same email twice shows an error", async ({ page }) => {
  const email = newEmail();
  await registerThroughUi(page, email);
  await expect(page.getByRole("heading", { name: "Log in" })).toBeVisible();

  await registerThroughUi(page, email);

  await expect(page.getByRole("alert")).toHaveText(
    "Email is already registered",
  );
});

test("logging in with a wrong password shows an error", async ({ page }) => {
  const email = newEmail();
  await registerThroughUi(page, email);
  await expect(page.getByRole("heading", { name: "Log in" })).toBeVisible();

  await logInThroughUi(page, email, "wrong horse battery");

  await expect(page.getByRole("alert")).toHaveText("Invalid email or password");
});

test("logging in with the right password shows the signed-in home page", async ({
  page,
}) => {
  const email = newEmail();
  await registerThroughUi(page, email);
  await expect(page.getByRole("heading", { name: "Log in" })).toBeVisible();

  await logInThroughUi(page, email, PASSWORD);

  await expect(page.getByText(`Signed in as ${email}`)).toBeVisible();
});
