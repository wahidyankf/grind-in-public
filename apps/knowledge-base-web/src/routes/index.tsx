import { createFileRoute } from "@tanstack/react-router";
import { useAuth } from "../auth";

export const Route = createFileRoute("/")({
  component: HomePage,
});

function HomePage() {
  const { session } = useAuth();
  if (!session) {
    return <p>You are not signed in.</p>;
  }
  return <p>Signed in as {session.user.email}</p>;
}
