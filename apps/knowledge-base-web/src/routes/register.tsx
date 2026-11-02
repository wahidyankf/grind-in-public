import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { RegisterForm } from "../register_form";

export const Route = createFileRoute("/register")({
  component: RegisterPage,
});

function RegisterPage() {
  const navigate = useNavigate();
  return (
    <section>
      <h1>Register</h1>
      <RegisterForm onRegistered={() => navigate({ to: "/login" })} />
    </section>
  );
}
