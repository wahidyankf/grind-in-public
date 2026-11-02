import { describe, expect, it } from "vitest";
import {
  hashPassword,
  isAcceptablePassword,
  verifyPassword,
} from "./passwords.js";

describe("isAcceptablePassword", () => {
  it("accepts a normal passphrase", () => {
    expect(isAcceptablePassword("correct horse")).toBe(true);
  });

  it("rejects passwords shorter than 8 characters", () => {
    expect(isAcceptablePassword("short1")).toBe(false);
  });

  it("rejects passwords longer than 64 characters", () => {
    expect(isAcceptablePassword("a".repeat(65))).toBe(false);
  });

  it("rejects common passwords in any letter case", () => {
    expect(isAcceptablePassword("password")).toBe(false);
    expect(isAcceptablePassword("PASSWORD")).toBe(false);
  });
});

describe("hashPassword and verifyPassword", () => {
  it("verifies the right password and rejects a wrong one", async () => {
    const hash = await hashPassword("correct horse");
    expect(await verifyPassword(hash, "correct horse")).toBe(true);
    expect(await verifyPassword(hash, "wrong horse")).toBe(false);
  });

  it("treats a malformed hash as a failed check", async () => {
    expect(await verifyPassword("not-a-hash", "correct horse")).toBe(false);
  });
});
