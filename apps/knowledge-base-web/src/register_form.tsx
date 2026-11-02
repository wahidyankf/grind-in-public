import { useMutation } from "@tanstack/react-query";
import { useState, type FormEvent } from "react";
import { registerUser, type PublicUser } from "./api";
import { registerSchema } from "./validation";

type Props = { onRegistered: (user: PublicUser) => void };

export function RegisterForm({ onRegistered }: Props) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [validationError, setValidationError] = useState<string | null>(null);

  const mutation = useMutation({
    mutationFn: (input: { email: string; password: string }) =>
      registerUser(input.email, input.password),
    // Pass only the user. TanStack Query also sends the variables and context.
    onSuccess: (user) => onRegistered(user),
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const parsed = registerSchema.safeParse({ email, password });
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
        <label htmlFor="register-email">Email</label>
        <input
          id="register-email"
          type="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
        />
      </div>
      <div>
        <label htmlFor="register-password">Password</label>
        <input
          id="register-password"
          type="password"
          value={password}
          onChange={(event) => setPassword(event.target.value)}
        />
      </div>
      {errorMessage && <p role="alert">{errorMessage}</p>}
      <button type="submit" disabled={mutation.isPending}>
        Register
      </button>
    </form>
  );
}
