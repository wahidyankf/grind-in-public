import { afterEach, describe, expect, it, vi } from "vitest";
import { ApiError, loginUser, registerUser } from "./api";

function stubFetch(status: number, body: unknown) {
  const fetchMock = vi.fn().mockResolvedValue(
    new Response(JSON.stringify(body), {
      status,
      headers: { "Content-Type": "application/json" },
    }),
  );
  vi.stubGlobal("fetch", fetchMock);
  return fetchMock;
}

afterEach(() => {
  vi.unstubAllGlobals();
});

describe("registerUser", () => {
  it("posts JSON to the register route and returns the user", async () => {
    const fetchMock = stubFetch(201, { id: "1", email: "user@example.com" });

    const user = await registerUser("user@example.com", "correct horse");

    expect(user).toEqual({ id: "1", email: "user@example.com" });
    expect(fetchMock).toHaveBeenCalledWith("/api/auth/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: "user@example.com",
        password: "correct horse",
      }),
    });
  });

  it("throws an ApiError with the server message on failure", async () => {
    stubFetch(409, { message: "Email is already registered" });

    const error = await registerUser("user@example.com", "correct horse").catch(
      (caught: unknown) => caught,
    );

    expect(error).toBeInstanceOf(ApiError);
    expect((error as ApiError).status).toBe(409);
    expect((error as ApiError).message).toBe("Email is already registered");
  });
});

describe("loginUser", () => {
  it("returns the access token and user", async () => {
    stubFetch(200, {
      accessToken: "token",
      tokenType: "bearer",
      expiresIn: 900,
      user: { id: "1", email: "user@example.com" },
    });

    const session = await loginUser("user@example.com", "correct horse");

    expect(session.accessToken).toBe("token");
    expect(session.user.email).toBe("user@example.com");
  });
});
