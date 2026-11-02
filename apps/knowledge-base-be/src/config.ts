export type Config = {
  databaseUrl: string;
  jwtSecret: string;
  accessTokenTtlSeconds: number;
  port: number;
};

// Matches the defaults in compose.knowledge-base.yaml, so db:up works without setup.
export const DEFAULT_DATABASE_URL =
  "postgresql://knowledge_base_be:knowledge_base_be@localhost:5232/knowledge_base_be";

// Development only. Production must set JWT_SECRET.
const DEV_JWT_SECRET = "development-only-secret-do-not-use-in-production";
const MIN_PRODUCTION_SECRET_LENGTH = 32;

export function buildConfig(env: NodeJS.ProcessEnv): Config {
  const isProduction = env.NODE_ENV === "production";
  const jwtSecret = env.JWT_SECRET ?? (isProduction ? "" : DEV_JWT_SECRET);
  // Fail at startup, not on the first login, when production has a weak secret.
  if (isProduction && jwtSecret.length < MIN_PRODUCTION_SECRET_LENGTH) {
    throw new Error(
      `JWT_SECRET must be at least ${MIN_PRODUCTION_SECRET_LENGTH} characters in production`,
    );
  }
  return {
    databaseUrl: env.DATABASE_URL ?? DEFAULT_DATABASE_URL,
    jwtSecret,
    accessTokenTtlSeconds: parsePositiveInteger(
      "ACCESS_TOKEN_TTL_SECONDS",
      env.ACCESS_TOKEN_TTL_SECONDS,
      900,
    ),
    port: parsePositiveInteger("PORT", env.PORT, 8200),
  };
}

function parsePositiveInteger(
  name: string,
  value: string | undefined,
  fallback: number,
): number {
  if (value === undefined) {
    return fallback;
  }
  const parsed = Number(value);
  if (!Number.isInteger(parsed) || parsed <= 0) {
    throw new Error(`${name} must be a positive integer`);
  }
  return parsed;
}

// Built once at import time, so the rest of the app reads a single object.
export const config = buildConfig(process.env);
