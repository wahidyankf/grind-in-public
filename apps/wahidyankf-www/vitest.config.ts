import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";

// `react()` compiles the JSX every component test and step file renders, and
// `tsconfigPaths()` resolves the `@/features/...` specifiers they import. A
// project missing either fails to load its own subject, so all three below
// carry this list rather than inheriting it — Vitest projects do not inherit
// plugins from the root config.
const sharedPlugins = [react(), tsconfigPaths()];

// Vitest's 5s default is measured from when a test starts, not from when its
// worker is ready, so the first test in a file absorbs that file's jsdom setup
// and module import alongside its own work. Rendering `CvContent` — the whole
// CV in one component — under the contention of a full parallel run crossed 5s
// intermittently: green in isolation, red roughly one run in four. A flaky gate
// is worse than a slow one, and this raises the limit rather than narrowing the
// test, because the render being measured is the behaviour under test.
//
// Spelled on each project below rather than once at the root, for the same
// reason `sharedPlugins` is: a project's own `test` block is what the project
// runs with, and relying on inheritance here would leave the fix in place only
// where it happened to reach.
const TEST_TIMEOUT_MS = 20_000;

export default defineConfig({
  plugins: sharedPlugins,
  test: {
    passWithNoTests: true,
    testTimeout: TEST_TIMEOUT_MS,
    coverage: {
      provider: "v8",
      // Vitest 4 reports only the files a test actually covered. Naming the
      // denominator explicitly is what makes an untested module count against
      // the 99% floor instead of vanishing from the report; without it the
      // gate measures what happened to be imported.
      include: ["src/**"],
      exclude: [
        "src/app/fonts/**",
        "src/app/**/*.css",
        "src/test/**",
        "**/*.config.*",
        "**/.next/**",
        "**/dist/**",
        "**/coverage/**",
      ],
      reporter: ["text", "json-summary", "lcov"],
    },
    projects: [
      {
        plugins: sharedPlugins,
        test: {
          name: "unit",
          testTimeout: TEST_TIMEOUT_MS,
          include: ["src/**/*.unit.test.{ts,tsx}"],
          exclude: ["node_modules"],
          environment: "jsdom",
          setupFiles: ["./src/test/setup.ts"],
        },
      },
      {
        plugins: sharedPlugins,
        test: {
          name: "behaviour-unit",
          testTimeout: TEST_TIMEOUT_MS,
          include: [
            "tests/bdd/**/*.{ts,tsx}",
            "tests/integration/cv-pdf.integration.test.ts",
          ],
          exclude: ["tests/bdd/test-layer.ts", "node_modules"],
          env: { WAHIDYANKF_WWW_BEHAVIOUR_TEST_LAYER: "unit" },
          environment: "jsdom",
          setupFiles: ["./src/test/behaviour-setup.ts"],
        },
      },
      {
        // No setup file here. Its `@testing-library/jest-dom/vitest` matchers
        // and `afterEach(cleanup)` act on a DOM, and this project runs under
        // `node` against the real filesystem.
        plugins: sharedPlugins,
        test: {
          name: "integration",
          testTimeout: TEST_TIMEOUT_MS,
          include: [
            "tests/integration/**/*.{ts,tsx}",
            "tests/bdd/env-loader.steps.ts",
            "tests/bdd/port-resolver.behaviour.test.ts",
            "tests/bdd/tier-env.behaviour.test.ts",
          ],
          exclude: ["node_modules"],
          env: { WAHIDYANKF_WWW_BEHAVIOUR_TEST_LAYER: "integration" },
          environment: "node",
        },
      },
    ],
  },
});
