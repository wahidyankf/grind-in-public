import { describe, expect, it } from "vitest";
import { signAccessToken, verifyAccessToken } from "./tokens.js";

const SECRET = "s".repeat(32);

describe("access tokens", () => {
  it("verifies a token signed with the same secret", async () => {
    const token = await signAccessToken("user-1", SECRET, 900);
    expect(await verifyAccessToken(token, SECRET)).toBe("user-1");
  });

  it("rejects a token signed with another secret", async () => {
    const token = await signAccessToken("user-1", SECRET, 900);
    await expect(verifyAccessToken(token, "x".repeat(32))).rejects.toThrow();
  });
});
