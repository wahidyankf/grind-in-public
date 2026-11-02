import { defineConfig } from "vitest/config";

// Integration tests need the Postgres container from compose.knowledge-base.yaml.
export default defineConfig({
  test: {
    include: ["src/**/*_integration_tests.ts"],
    exclude: ["node_modules", "dist"],
  },
});
