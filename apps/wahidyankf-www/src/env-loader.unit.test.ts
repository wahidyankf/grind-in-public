import { describe, it, expect, vi } from "vitest";

// The loader's own behavior — tier resolution, the stray-file guard, and
// process-env-wins — is specified by `tier-env-loading.feature` and bound under
// `tests/bdd/`. This test is about the composition root instead: that importing
// this module calls `loadTierEnv()` exactly once, and that it re-exports the
// three names `next.config.ts` reaches for. The mock is what makes the first
// assertion possible, since the call happens at import time and cannot be
// observed after the fact.
const loadTierEnv = vi.fn();

vi.mock("./features/env/core/tier-env", () => ({
  loadTierEnv,
  resolveTier: vi.fn(),
  tierEnvFilePath: vi.fn(),
}));

describe("env-loader", () => {
  it("calls loadTierEnv exactly once when imported", async () => {
    await import("./env-loader");
    expect(loadTierEnv).toHaveBeenCalledTimes(1);
  });

  it("re-exports the three names the config reaches for", async () => {
    const module = await import("./env-loader");
    expect(Object.keys(module).sort()).toEqual([
      "loadTierEnv",
      "resolveTier",
      "tierEnvFilePath",
    ]);
  });
});
