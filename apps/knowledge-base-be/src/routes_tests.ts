import type { FastifyInstance } from "fastify";
import { afterEach, beforeEach, describe, expect, it } from "vitest";
import { buildApp } from "./app.js";
import { createFakeAuthStore } from "./fake_store.js";

const config = {
  databaseUrl: "postgresql://unused",
  jwtSecret: "s".repeat(32),
  accessTokenTtlSeconds: 900,
  port: 8200,
};

let app: FastifyInstance;

beforeEach(() => {
  app = buildApp({ store: createFakeAuthStore(), config });
});

afterEach(async () => {
  await app.close();
});

describe("GET /api/health", () => {
  it("returns ok", async () => {
    const response = await app.inject({ method: "GET", url: "/api/health" });
    expect(response.statusCode).toBe(200);
    expect(response.json()).toEqual({ status: "ok" });
  });
});

describe("POST /api/auth/register", () => {
  it("returns 201 with the new user", async () => {
    const response = await app.inject({
      method: "POST",
      url: "/api/auth/register",
      payload: { email: "user@example.com", password: "correct horse" },
    });
    expect(response.statusCode).toBe(201);
    expect(response.json()).toMatchObject({ email: "user@example.com" });
  });

  it("returns 409 when the email is taken", async () => {
    const payload = { email: "user@example.com", password: "correct horse" };
    await app.inject({ method: "POST", url: "/api/auth/register", payload });
    const response = await app.inject({
      method: "POST",
      url: "/api/auth/register",
      payload,
    });
    expect(response.statusCode).toBe(409);
  });

  it("returns 422 for a weak password", async () => {
    const response = await app.inject({
      method: "POST",
      url: "/api/auth/register",
      payload: { email: "user@example.com", password: "password" },
    });
    expect(response.statusCode).toBe(422);
  });

  it("returns 422 when the body is missing fields", async () => {
    const response = await app.inject({
      method: "POST",
      url: "/api/auth/register",
      payload: {},
    });
    expect(response.statusCode).toBe(422);
  });
});

describe("POST /api/auth/login", () => {
  it("returns 200 with a bearer token and the user", async () => {
    const credentials = {
      email: "user@example.com",
      password: "correct horse",
    };
    await app.inject({
      method: "POST",
      url: "/api/auth/register",
      payload: credentials,
    });
    const response = await app.inject({
      method: "POST",
      url: "/api/auth/login",
      payload: credentials,
    });
    expect(response.statusCode).toBe(200);
    expect(response.json()).toMatchObject({
      tokenType: "bearer",
      expiresIn: 900,
      user: { email: "user@example.com" },
    });
  });

  it("returns 401 for a wrong password", async () => {
    await app.inject({
      method: "POST",
      url: "/api/auth/register",
      payload: { email: "user@example.com", password: "correct horse" },
    });
    const response = await app.inject({
      method: "POST",
      url: "/api/auth/login",
      payload: { email: "user@example.com", password: "wrong horse" },
    });
    expect(response.statusCode).toBe(401);
  });
});
