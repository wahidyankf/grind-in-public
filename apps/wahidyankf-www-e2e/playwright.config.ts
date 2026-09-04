import { createServer } from "node:net";
import { defineConfig, devices } from "@playwright/test";
import { defineBddConfig } from "playwright-bdd";

// Pin the tier deterministically for e2e runs. `loadTierEnv` in
// apps/wahidyankf-www/src/features/env/core/tier-env.ts reads APP_ENV and falls back to "local"
// when it is unset, and "local" is the one tier whose stray-file guard is skipped — so an unset
// APP_ENV here would let Next.js auto-load a developer's real .env.local and hand the suite their
// own environment instead of test fixtures. The contract those two sentences describe is specified
// in specs/apps/wahidyankf-www/behaviours/tier-env-loading.feature.
process.env.APP_ENV = "test";

async function allocateRunPort(): Promise<number> {
  const server = createServer();
  await new Promise<void>((resolve, reject) => {
    server.once("error", reject);
    server.listen(0, "127.0.0.1", resolve);
  });
  const address = server.address();
  if (address === null || typeof address === "string") {
    server.close();
    throw new Error("Playwright could not allocate an isolated IPv4 port");
  }
  await new Promise<void>((resolve, reject) => {
    server.close((error) => (error ? reject(error) : resolve()));
  });
  return address.port;
}

const inheritedRunPort = process.env.WAHIDYANKF_WWW_E2E_RUN_PORT;
const isPlaywrightWorker = process.env.TEST_WORKER_INDEX !== undefined;
if (
  isPlaywrightWorker &&
  (inheritedRunPort === undefined || !/^\d+$/.test(inheritedRunPort))
) {
  throw new Error("Playwright worker did not inherit its controller run port");
}
const e2ePort = isPlaywrightWorker
  ? Number(inheritedRunPort)
  : await allocateRunPort();
if (!Number.isInteger(e2ePort) || e2ePort < 1 || e2ePort > 65_535) {
  throw new Error("Playwright received an invalid inherited E2E run port");
}
const e2eOrigin = `http://127.0.0.1:${e2ePort}`;
process.env.WAHIDYANKF_WWW_E2E_RUN_PORT = String(e2ePort);
process.env.WAHIDYANKF_WWW_PORT = String(e2ePort);

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
    baseURL: e2eOrigin,
    trace: "on-first-retry",
    screenshot: "only-on-failure",
  },
  // The suite drives a real `next start`. The source project built a container image, ran it, and
  // read the published port back out; this repository runs the application process directly, so
  // the port is allocated per run and passed through the application's prefixed port contract.
  // This owned process builds before starting so the forced test tier applies to the build as well
  // as the server; an outer Nx prerequisite would start before this config can sanitize the tier.
  //
  // Never reuse an existing listener. A process already bound to this run's selected port is
  // unverified external state, so Playwright must fail closed instead of treating it as this run's
  // freshly built application.
  webServer: {
    command:
      "npm exec nx -- run wahidyankf-www:build --skip-nx-cache && npm exec nx -- run wahidyankf-www:start",
    url: e2eOrigin,
    cwd: "../..",
    reuseExistingServer: false,
    timeout: 120000,
  },
  projects: [
    {
      name: "chromium",
      use: { ...devices["Desktop Chrome"] },
    },
  ],
});
