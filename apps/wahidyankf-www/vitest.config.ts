import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";

// `react()` compiles the JSX every component test and step file renders, and
// `tsconfigPaths()` resolves the `@/features/...` specifiers they import. A
// project missing either fails to load its own subject, so all three below
// carry this list rather than inheriting it — Vitest projects do not inherit
// plugins from the root config.
const sharedPlugins = [react(), tsconfigPaths()];

export default defineConfig({
  plugins: sharedPlugins,
  test: {
    passWithNoTests: true,
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
          include: ["src/**/*.unit.test.{ts,tsx}"],
          exclude: ["node_modules"],
          environment: "jsdom",
          setupFiles: ["./src/test/setup.ts"],
        },
      },
      {
        plugins: sharedPlugins,
        test: {
          name: "behavior",
          include: ["tests/bdd/**/*.{ts,tsx}"],
          exclude: ["node_modules"],
          environment: "jsdom",
          setupFiles: ["./src/test/setup.ts"],
        },
      },
      {
        // No setup file here. Its `@testing-library/jest-dom/vitest` matchers
        // and `afterEach(cleanup)` act on a DOM, and this project runs under
        // `node` against the real filesystem.
        plugins: sharedPlugins,
        test: {
          name: "integration",
          include: ["tests/integration/**/*.{ts,tsx}"],
          exclude: ["node_modules"],
          environment: "node",
        },
      },
    ],
  },
});
