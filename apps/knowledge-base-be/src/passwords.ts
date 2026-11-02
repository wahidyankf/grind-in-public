import argon2 from "argon2";
import { COMMON_PASSWORDS } from "./common_passwords.js";

// NIST SP 800-63B: at least 8 characters, support at least 64.
export const MIN_LENGTH = 8;
export const MAX_LENGTH = 64;

const commonPasswords = new Set(
  COMMON_PASSWORDS.map((password) => password.toLowerCase()),
);

export function isAcceptablePassword(password: string): boolean {
  const length = [...password].length;
  if (length < MIN_LENGTH || length > MAX_LENGTH) {
    return false;
  }
  return !commonPasswords.has(password.toLowerCase());
}

export function hashPassword(password: string): Promise<string> {
  return argon2.hash(password);
}

export async function verifyPassword(
  hash: string,
  password: string,
): Promise<boolean> {
  try {
    return await argon2.verify(hash, password);
  } catch {
    // A malformed stored hash counts as a failed check, not a server error.
    return false;
  }
}
