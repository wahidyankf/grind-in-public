import { describe, expect, it } from "vitest";
import {
  EmailTakenError,
  InvalidCredentialsError,
  InvalidEmailError,
  WeakPasswordError,
  login,
  register,
} from "./auth.js";
import { createFakeAuthStore } from "./fake_store.js";
import { verifyAccessToken } from "./tokens.js";

const config = { jwtSecret: "s".repeat(32), accessTokenTtlSeconds: 900 };

describe("register", () => {
  it("stores the user with a normalized email", async () => {
    const store = createFakeAuthStore();
    const user = await register(store, "  User@Example.COM ", "correct horse");
    expect(user.email).toBe("user@example.com");
    expect(user.id).toBeTruthy();
  });

  it("rejects an email that is already registered", async () => {
    const store = createFakeAuthStore();
    await register(store, "user@example.com", "correct horse");
    await expect(
      register(store, "USER@example.com", "another horse"),
    ).rejects.toThrow(EmailTakenError);
  });

  it("rejects an invalid email", async () => {
    await expect(
      register(createFakeAuthStore(), "not-an-email", "correct horse"),
    ).rejects.toThrow(InvalidEmailError);
  });

  it("rejects a weak password", async () => {
    await expect(
      register(createFakeAuthStore(), "user@example.com", "password"),
    ).rejects.toThrow(WeakPasswordError);
  });
});

describe("login", () => {
  it("returns a bearer token for the right password", async () => {
    const store = createFakeAuthStore();
    const registered = await register(
      store,
      "user@example.com",
      "correct horse",
    );
    const session = await login(
      store,
      config,
      "user@example.com",
      "correct horse",
    );
    expect(session.tokenType).toBe("bearer");
    expect(session.expiresIn).toBe(900);
    expect(session.user).toEqual(registered);
    expect(await verifyAccessToken(session.accessToken, config.jwtSecret)).toBe(
      registered.id,
    );
  });

  it("rejects a wrong password", async () => {
    const store = createFakeAuthStore();
    await register(store, "user@example.com", "correct horse");
    await expect(
      login(store, config, "user@example.com", "wrong horse"),
    ).rejects.toThrow(InvalidCredentialsError);
  });

  it("rejects an unknown email with the same error as a wrong password", async () => {
    await expect(
      login(
        createFakeAuthStore(),
        config,
        "nobody@example.com",
        "correct horse",
      ),
    ).rejects.toThrow(InvalidCredentialsError);
  });
});
