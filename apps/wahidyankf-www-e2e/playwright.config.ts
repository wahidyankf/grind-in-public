import { defineConfig, devices } from "@playwright/test";
import { defineBddConfig } from "playwright-bdd";

// Pin the tier deterministically for e2e runs. `loadTierEnv` in
// apps/wahidyankf-www/src/features/env/core/tier-env.ts reads APP_ENV and falls back to "local"
// when it is unset, and "local" is the one tier whose stray-file guard is skipped — so an unset
// APP_ENV here would let Next.js auto-load a developer's real .env.local and hand the suite their
// own environment instead of test fixtures. The contract those two sentences describe is specified
// in specs/apps/wahidyankf-www/behavior/tier-env-loading.feature.
process.env.APP_ENV ??= "test";

const testDir = defineBddConfig({
  featuresRoot: "../../specs/apps/wahidyankf-www/behavior",
  steps: "./steps/**/*.ts",
  // Default is 'fail-on-gen': bddgen refuses to generate ANY test file while ANY scenario in the
  // globbed features lacks a matching step def. This adapter binds eight of the corpus's twelve
  // feature files — accessibility, cv, home, personal-projects, responsive, search,
  // static-filterable-routes, and theme — and deliberately binds none of the other four, whose
  // nineteen scenarios have no browser equivalent to drive:
  //
  //   env-loader.feature       (4)  Node-process env loading, bound by the application's unit layer
  //   port-resolver.feature    (8)  the runtime port contract, likewise unit-bound
  //   tier-env-loading.feature (5)  the tier-file loader, likewise unit-bound
  //   cv-export.feature        (2)  PDF export at the filesystem boundary, bound by the
  //                                 application's integration layer
  //
  // 'skip-scenario' lets generation succeed and renders those nineteen as `test.fixme` instead of
  // hard-blocking the whole suite. `specs:e2e:baseline` is what holds that count to nineteen, so a
  // scenario silently falling out of this suite is a failure rather than a quieter run.
  //
  // `tags: "not @unit"` was tried and reverted, and that rejection is why the setting is glob-wide
  // rather than tag-scoped. This corpus tags many already-IMPLEMENTED e2e scenarios `@unit @e2e`,
  // not just the plain-`@unit` ones above. A tag filter excludes by tag regardless of co-tags, so
  // `not @unit` silently drops those real, already-bound scenarios from generation too.
  missingSteps: "skip-scenario",
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
  // it. `test:e2e` declares `dependsOn` on `wahidyankf-www:build`, because `next start` serves a
  // `.next` directory it will not build itself.
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
