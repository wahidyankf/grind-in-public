import { z } from "zod";
import type { Config } from "./config.js";
import {
  hashPassword,
  isAcceptablePassword,
  verifyPassword,
} from "./passwords.js";
import type { AuthStore, PublicUser } from "./store.js";
import { signAccessToken } from "./tokens.js";

export class InvalidEmailError extends Error {}
export class WeakPasswordError extends Error {}
export class EmailTakenError extends Error {}
export class InvalidCredentialsError extends Error {}

const emailSchema = z.email();

function normalizeEmail(rawEmail: string): string {
  const email = rawEmail.trim().toLowerCase();
  if (!emailSchema.safeParse(email).success) {
    throw new InvalidEmailError("Invalid email");
  }
  return email;
}

export async function register(
  store: AuthStore,
  rawEmail: string,
  password: string,
): Promise<PublicUser> {
  const email = normalizeEmail(rawEmail);
  if (!isAcceptablePassword(password)) {
    throw new WeakPasswordError("Weak password");
  }
  const user = await store.createUser({
    email,
    passwordHash: await hashPassword(password),
  });
  if (!user) {
    throw new EmailTakenError("Email is already registered");
  }
  return user;
}

// Checked when the email is unknown, so login takes about as long either way.
let unknownUserHash: Promise<string> | undefined;
function getUnknownUserHash(): Promise<string> {
  unknownUserHash ??= hashPassword("unknown-user-timing-filler");
  return unknownUserHash;
}

export type LoginResult = {
  accessToken: string;
  tokenType: "bearer";
  expiresIn: number;
  user: PublicUser;
};

export async function login(
  store: AuthStore,
  config: Pick<Config, "jwtSecret" | "accessTokenTtlSeconds">,
  rawEmail: string,
  password: string,
): Promise<LoginResult> {
  const email = rawEmail.trim().toLowerCase();
  const user = await store.findUserByEmail(email);
  const hash = user?.passwordHash ?? (await getUnknownUserHash());
  const passwordMatches = await verifyPassword(hash, password);
  if (!user || !passwordMatches) {
    throw new InvalidCredentialsError("Invalid email or password");
  }
  const accessToken = await signAccessToken(
    user.id,
    config.jwtSecret,
    config.accessTokenTtlSeconds,
  );
  return {
    accessToken,
    tokenType: "bearer",
    expiresIn: config.accessTokenTtlSeconds,
    user: { id: user.id, email: user.email },
  };
}
