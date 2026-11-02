import { randomUUID } from "node:crypto";
import type { AuthStore, StoredUser } from "./store.js";

// In-memory store for unit tests. It keeps no data between runs.
export function createFakeAuthStore(): AuthStore {
  const usersByEmail = new Map<string, StoredUser>();
  return {
    async createUser({ email, passwordHash }) {
      if (usersByEmail.has(email)) {
        return null;
      }
      const user: StoredUser = { id: randomUUID(), email, passwordHash };
      usersByEmail.set(email, user);
      return { id: user.id, email: user.email };
    },
    async findUserByEmail(email) {
      return usersByEmail.get(email) ?? null;
    },
  };
}
