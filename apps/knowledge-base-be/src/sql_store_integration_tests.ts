import { randomUUID } from "node:crypto";
import { afterAll, describe, expect, it } from "vitest";
import { DEFAULT_DATABASE_URL } from "./config.js";
import { createPool } from "./db.js";
import { createSqlAuthStore } from "./sql_store.js";

// Needs Postgres from compose.knowledge-base.yaml with migrations applied (npx nx db:migrate knowledge-base-be).
const pool = createPool(process.env.DATABASE_URL ?? DEFAULT_DATABASE_URL);
const store = createSqlAuthStore(pool);

afterAll(async () => {
  await pool.end();
});

describe("SQL auth store", () => {
  it("stores a user and finds it by email", async () => {
    const email = `user-${randomUUID()}@example.test`;
    const created = await store.createUser({ email, passwordHash: "hash" });
    expect(created?.email).toBe(email);

    const found = await store.findUserByEmail(email);
    expect(found?.passwordHash).toBe("hash");
  });

  it("returns null when the email is already stored", async () => {
    const email = `user-${randomUUID()}@example.test`;
    await store.createUser({ email, passwordHash: "hash" });
    expect(await store.createUser({ email, passwordHash: "other" })).toBeNull();
  });
});
