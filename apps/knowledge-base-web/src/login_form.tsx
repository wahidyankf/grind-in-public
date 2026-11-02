import { useMutation } from "@tanstack/react-query";
import { useState, type FormEvent } from "react";
import { loginUser } from "./api";
import type { Session } from "./auth";
import { loginSchema } from "./validation";

type Props = { onLoggedIn: (session: Session) => void };

export function LoginForm({ onLoggedIn }: Props) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [validationError, setValidationError] = useState<string | null>(null);

  const mutation = useMutation({
    mutationFn: (input: { email: string; password: string }) =>
      loginUser(input.email, input.password),
    onSuccess: (data) =>
      onLoggedIn({ accessToken: data.accessToken, user: data.user }),
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const parsed = loginSchema.safeParse({ email, password });
    if (!parsed.success) {
      setValidationError(parsed.error.issues[0]?.message ?? "Invalid input");
      return;
    }
    setValidationError(null);
    mutation.mutate(parsed.data);
  }

  const errorMessage = validationError ?? mutation.error?.message ?? null;

  return (
    <form onSubmit={handleSubmit} noValidate>
      <div>
        <label htmlFor="login-email">Email</label>
        <input
          id="login-email"
          type="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
        />
      </div>
      <div>
        <label htmlFor="login-password">Password</label>
        <input
          id="login-password"
          type="password"
          value={password}
          onChange={(event) => setPassword(event.target.value)}
        />
      </div>
      {errorMessage && <p role="alert">{errorMessage}</p>}
      <button type="submit" disabled={mutation.isPending}>
        Log in
      </button>
    </form>
  );
}
