import type { Pool } from "pg";
import type { AuthStore, NewUser, PublicUser, StoredUser } from "./store.js";

// Each call is one statement, so a failed request leaves no partial rows.
export function createSqlAuthStore(pool: Pool): AuthStore {
  return {
    async createUser({
      email,
      passwordHash,
    }: NewUser): Promise<PublicUser | null> {
      // "on conflict do nothing" returns no row when the email is taken.
      const result = await pool.query<PublicUser>(
        "insert into users (email, password_hash) values ($1, $2) on conflict (email) do nothing returning id, email",
        [email, passwordHash],
      );
      return result.rows[0] ?? null;
    },
    async findUserByEmail(email: string): Promise<StoredUser | null> {
      const result = await pool.query<{
        id: string;
        email: string;
        password_hash: string;
      }>("select id, email, password_hash from users where email = $1", [
        email,
      ]);
      const row = result.rows[0];
      if (!row) {
        return null;
      }
      return { id: row.id, email: row.email, passwordHash: row.password_hash };
    },
  };
}
