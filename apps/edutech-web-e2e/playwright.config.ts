import { defineConfig, devices } from "@playwright/test";

const port = 3100;

// Tests run against the production build, so they match what ships.
export default defineConfig({
  testDir: "./tests",
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  reporter: "list",
  use: {
    baseURL: `http://localhost:${port}`,
    trace: "on-first-retry",
  },
  projects: [
    {
      name: "chromium",
      use: { ...devices["Desktop Chrome"] },
    },
  ],
  webServer: {
    command:
      "npm --prefix ../edutech-web run build && npm --prefix ../edutech-web run start",
    // Wait for the health check, not just an open port, so tests start only when the app is ready.
    url: `http://localhost:${port}/api/health`,
    // Never reuse a server already on this port. A running dev server would pass the
    // health check and the tests would run against it instead of the production build.
    reuseExistingServer: false,
    timeout: 240_000,
  },
});
