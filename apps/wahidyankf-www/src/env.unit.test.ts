import { describe, it, expect } from "vitest";
import { env } from "./env";

// This application declares no environment variables of its own: everything it
// reads at runtime comes from Next.js or from the deployment platform. The
// `createEnv()` call is still here rather than deleted, because it is the
// validation seam a future variable gets added to, and it runs at import time —
// so importing this module at all is most of what the test does. What is
// asserted is the emptiness, which is the claim `.env.example` also makes.
describe("env", () => {
  it("validates successfully with no declared server variables", () => {
    expect(env).toBeDefined();
  });

  it("declares no variables of its own", () => {
    // `createEnv` returns a proxy carrying only what was declared. Anything
    // appearing here would be a variable added without a matching entry in
    // `.env.example`.
    expect(Object.keys(env)).toEqual([]);
  });
});
