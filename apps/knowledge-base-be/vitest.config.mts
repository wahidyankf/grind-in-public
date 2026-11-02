import { defineConfig } from "vitest/config";

// Unit tests use the in-memory store, so they run without Postgres.
export default defineConfig({
  test: {
    include: ["src/**/*_tests.ts"],
    exclude: ["node_modules", "dist", "src/**/*_integration_tests.ts", "src/**/*_production_tests.ts"],
  },
});
