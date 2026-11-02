import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { cleanup, fireEvent, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it, vi } from "vitest";
import { RegisterForm } from "./register_form";

afterEach(() => {
  cleanup();
  vi.unstubAllGlobals();
});

function renderForm(onRegistered = vi.fn()) {
  const client = new QueryClient({
    defaultOptions: { mutations: { retry: false } },
  });
  render(
    <QueryClientProvider client={client}>
      <RegisterForm onRegistered={onRegistered} />
    </QueryClientProvider>,
  );
  return onRegistered;
}

function fill(email: string, password: string) {
  fireEvent.change(screen.getByLabelText("Email"), {
    target: { value: email },
  });
  fireEvent.change(screen.getByLabelText("Password"), {
    target: { value: password },
  });
  fireEvent.click(screen.getByRole("button", { name: "Register" }));
}

describe("RegisterForm", () => {
  it("shows a validation error and does not call the API for a bad email", () => {
    const fetchMock = vi.fn();
    vi.stubGlobal("fetch", fetchMock);
    renderForm();

    fill("not-an-email", "correct horse");

    expect(screen.getByRole("alert").textContent).toBe("Enter a valid email");
    expect(fetchMock).not.toHaveBeenCalled();
  });

  it("shows the server message when the email is taken", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        new Response(
          JSON.stringify({ message: "Email is already registered" }),
          {
            status: 409,
            headers: { "Content-Type": "application/json" },
          },
        ),
      ),
    );
    renderForm();

    fill("user@example.com", "correct horse");

    expect((await screen.findByRole("alert")).textContent).toBe(
      "Email is already registered",
    );
  });

  it("calls onRegistered after a successful register", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        new Response(JSON.stringify({ id: "1", email: "user@example.com" }), {
          status: 201,
          headers: { "Content-Type": "application/json" },
        }),
      ),
    );
    const onRegistered = renderForm();

    fill("user@example.com", "correct horse");

    await vi.waitFor(() => {
      expect(onRegistered).toHaveBeenCalledWith({
        id: "1",
        email: "user@example.com",
      });
    });
  });
});
