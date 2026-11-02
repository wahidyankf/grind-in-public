import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { cleanup, fireEvent, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it, vi } from "vitest";
import { LoginForm } from "./login_form";

afterEach(() => {
  cleanup();
  vi.unstubAllGlobals();
});

function renderForm(onLoggedIn = vi.fn()) {
  const client = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  render(
    <QueryClientProvider client={client}>
      <LoginForm onLoggedIn={onLoggedIn} />
    </QueryClientProvider>,
  );
  return onLoggedIn;
}

function fill(email: string, password: string) {
  fireEvent.change(screen.getByLabelText("Email"), {
    target: { value: email },
  });
  fireEvent.change(screen.getByLabelText("Password"), {
    target: { value: password },
  });
  fireEvent.click(screen.getByRole("button", { name: "Log in" }));
}

describe("LoginForm", () => {
  it("shows the server message for a wrong password", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        new Response(JSON.stringify({ message: "Invalid email or password" }), {
          status: 401,
          headers: { "Content-Type": "application/json" },
        }),
      ),
    );
    renderForm();

    fill("user@example.com", "wrong horse");

    expect((await screen.findByRole("alert")).textContent).toBe(
      "Invalid email or password",
    );
  });

  it("passes the session to onLoggedIn after a successful login", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        new Response(
          JSON.stringify({
            accessToken: "token",
            tokenType: "bearer",
            expiresIn: 900,
            user: { id: "1", email: "user@example.com" },
          }),
          { status: 200, headers: { "Content-Type": "application/json" } },
        ),
      ),
    );
    const onLoggedIn = renderForm();

    fill("user@example.com", "correct horse");

    await vi.waitFor(() => {
      expect(onLoggedIn).toHaveBeenCalledWith({
        accessToken: "token",
        user: { id: "1", email: "user@example.com" },
      });
    });
  });
});
