import { defineConfig, devices } from "@playwright/test";

const webPort = 3200;
const backendPort = 8200;

// Tests run against the production builds, so they match what ships.
export default defineConfig({
  testDir: "./tests",
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  reporter: "list",
  use: {
    baseURL: `http://localhost:${webPort}`,
    trace: "on-first-retry",
  },
  projects: [
    {
      name: "chromium",
      use: { ...devices["Desktop Chrome"] },
    },
  ],
  webServer: [
    {
      command:
        "npm --prefix ../knowledge-base-be run build && npm --prefix ../knowledge-base-be run start",
      // Wait for the health check, not just an open port, so tests start only when the API is ready.
      url: `http://localhost:${backendPort}/api/health`,
      reuseExistingServer: !process.env.CI,
      timeout: 240_000,
    },
    {
      command:
        "npm --prefix ../knowledge-base-web run build && npm --prefix ../knowledge-base-web run start",
      url: `http://localhost:${webPort}/api/health`,
      reuseExistingServer: !process.env.CI,
      timeout: 240_000,
    },
  ],
});
