import { describe, expect, it } from "vitest";
import { DEFAULT_DATABASE_URL, buildConfig } from "./config.js";

const LONG_SECRET = "s".repeat(32);

describe("buildConfig", () => {
  it("uses local defaults when no env is set", () => {
    const config = buildConfig({});
    expect(config.databaseUrl).toBe(DEFAULT_DATABASE_URL);
    expect(config.port).toBe(8200);
    expect(config.accessTokenTtlSeconds).toBe(900);
  });

  it("reads values from env", () => {
    const config = buildConfig({
      DATABASE_URL: "postgresql://example/db",
      JWT_SECRET: LONG_SECRET,
      PORT: "9000",
      ACCESS_TOKEN_TTL_SECONDS: "60",
    });
    expect(config).toEqual({
      databaseUrl: "postgresql://example/db",
      jwtSecret: LONG_SECRET,
      accessTokenTtlSeconds: 60,
      port: 9000,
    });
  });

  it("rejects a non-integer PORT", () => {
    expect(() => buildConfig({ PORT: "abc" })).toThrow("PORT");
  });

  it("rejects a zero token lifetime", () => {
    expect(() => buildConfig({ ACCESS_TOKEN_TTL_SECONDS: "0" })).toThrow(
      "ACCESS_TOKEN_TTL_SECONDS",
    );
  });

  it("rejects a missing JWT_SECRET in production", () => {
    expect(() => buildConfig({ NODE_ENV: "production" })).toThrow("JWT_SECRET");
  });

  it("rejects a short JWT_SECRET in production", () => {
    expect(() =>
      buildConfig({ NODE_ENV: "production", JWT_SECRET: "short" }),
    ).toThrow("JWT_SECRET");
  });
});
