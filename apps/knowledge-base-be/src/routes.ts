import type { FastifyInstance, FastifyReply } from "fastify";
import { z } from "zod";
import {
  EmailTakenError,
  InvalidCredentialsError,
  InvalidEmailError,
  WeakPasswordError,
  login,
  register,
} from "./auth.js";
import type { Config } from "./config.js";
import type { AuthStore } from "./store.js";

// Every HTTP route lives in this file, with its full path, so the list is easy to scan.
// Shape checks only here; email and password rules live in auth.ts.
const credentialsBody = z.object({
  email: z.string().max(254),
  password: z.string().max(1024),
});

type RouteDeps = { store: AuthStore; config: Config };

export function registerRoutes(
  app: FastifyInstance,
  { store, config }: RouteDeps,
): void {
  app.get("/api/health", async () => ({ status: "ok" }));

  app.post("/api/auth/register", async (request, reply) => {
    const body = credentialsBody.safeParse(request.body);
    if (!body.success) {
      return reply.code(422).send({ message: "Invalid request body" });
    }
    try {
      const user = await register(store, body.data.email, body.data.password);
      return reply.code(201).send(user);
    } catch (error) {
      return sendAuthError(reply, error);
    }
  });

  app.post("/api/auth/login", async (request, reply) => {
    const body = credentialsBody.safeParse(request.body);
    if (!body.success) {
      return reply.code(422).send({ message: "Invalid request body" });
    }
    try {
      const session = await login(
        store,
        config,
        body.data.email,
        body.data.password,
      );
      return reply.code(200).send(session);
    } catch (error) {
      return sendAuthError(reply, error);
    }
  });
}

function sendAuthError(reply: FastifyReply, error: unknown): FastifyReply {
  if (error instanceof InvalidEmailError) {
    return reply.code(422).send({ message: "Invalid email" });
  }
  if (error instanceof WeakPasswordError) {
    return reply.code(422).send({
      message: "Password must be 8 to 64 characters and not a common password",
    });
  }
  if (error instanceof EmailTakenError) {
    return reply.code(409).send({ message: "Email is already registered" });
  }
  if (error instanceof InvalidCredentialsError) {
    return reply.code(401).send({ message: "Invalid email or password" });
  }
  throw error;
}
