import Fastify, { type FastifyInstance } from "fastify";
import type { Config } from "./config.js";
import { registerRoutes } from "./routes.js";
import type { AuthStore } from "./store.js";

export type AppDeps = {
  store: AuthStore;
  config: Config;
  logger?: boolean;
};

// Builds the app without listening, so tests can call it with app.inject.
export function buildApp({
  store,
  config,
  logger = false,
}: AppDeps): FastifyInstance {
  const app = Fastify({ logger });
  registerRoutes(app, { store, config });
  return app;
}
