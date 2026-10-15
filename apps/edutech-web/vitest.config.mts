import react from "@vitejs/plugin-react";
import { defineConfig } from "vitest/config";

export default defineConfig({
  plugins: [react()],
  test: {
    environment: "jsdom",
    // Unit tests sit next to the source as *_tests.ts(x).
    include: ["**/*_tests.{ts,tsx}"],
    // *_production_tests are config checks for production mode only.
    exclude: ["node_modules", ".next", "**/*_production_tests.{ts,tsx}"],
  },
});
