import { defineConfig, devices } from "@playwright/test";
import { defineBddConfig } from "playwright-bdd";

// Pin the tier deterministically for e2e runs. `loadTierEnv` in
// apps/wahidyankf-www/src/features/env/core/tier-env.ts reads APP_ENV and falls back to "local"
// when it is unset, and "local" is the one tier whose stray-file guard is skipped — so an unset
// APP_ENV here would let Next.js auto-load a developer's real .env.local and hand the suite their
// own environment instead of test fixtures. The contract those two sentences describe is specified
// in specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature.
process.env.APP_ENV ??= "test";

const testDir = defineBddConfig({
  featuresRoot: "../../specs/apps/wahidyankf-www/behaviours",
  steps: "./tests/steps/**/*.ts",
  tags: "not @e2e-exempt",
});

export default defineConfig({
  testDir,
  timeout: 60000,
  fullyParallel: false,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,
  workers: 1,
  reporter: process.env.CI ? [["list"], ["html"]] : "list",
  use: {
    baseURL: process.env.BASE_URL || "http://localhost:3201",
    trace: "on-first-retry",
    screenshot: "only-on-failure",
  },
  // The suite drives a real `next start`. The source project built a container image, ran it, and
  // read the published port back out; this repository runs the application process directly, so
  // the port is the fixed default the `start` target resolves and `baseURL` above already names
  // it. `test:e2e` depends on the owner build, because `next start` serves a `.next` directory it
  // will not build itself.
  //
  // `reuseExistingServer` is off under CI so a stale process can never be mistaken for a fresh
  // build, and on locally so a developer iterating on steps does not pay a server restart per run.
  webServer: {
    command: "npx nx run wahidyankf-www:start",
    url: "http://localhost:3201",
    cwd: "../..",
    reuseExistingServer: !process.env.CI,
    timeout: 120000,
  },
  projects: [
    {
      name: "chromium",
      use: { ...devices["Desktop Chrome"] },
    },
  ],
});
