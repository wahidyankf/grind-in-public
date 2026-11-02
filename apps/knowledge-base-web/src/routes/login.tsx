import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useAuth } from "../auth";
import { LoginForm } from "../login_form";

export const Route = createFileRoute("/login")({
  component: LoginPage,
});

function LoginPage() {
  const navigate = useNavigate();
  const { signIn } = useAuth();
  return (
    <section>
      <h1>Log in</h1>
      <LoginForm
        onLoggedIn={(session) => {
          signIn(session);
          navigate({ to: "/" });
        }}
      />
    </section>
  );
}
