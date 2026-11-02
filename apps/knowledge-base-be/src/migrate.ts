import { readFile, readdir } from "node:fs/promises";
import { join } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import { Pool } from "pg";
import { config } from "./config.js";

// Kept in-house so no migration library with an open advisory is added.
const migrationsDir = fileURLToPath(new URL("../migrations/", import.meta.url));

// Applies each migrations/*.sql file once, in filename order, each in its own transaction.
// Only the "-- Up Migration" part runs. The "-- Down Migration" part is kept for manual rollback.
export async function migrate(databaseUrl: string): Promise<string[]> {
  const pool = new Pool({ connectionString: databaseUrl });
  try {
    await pool.query(
      "create table if not exists schema_migrations (name text primary key, applied_at timestamptz not null default now())",
    );
    const applied = await pool.query<{ name: string }>(
      "select name from schema_migrations",
    );
    const appliedNames = new Set(applied.rows.map((row) => row.name));
    const files = (await readdir(migrationsDir))
      .filter((file) => file.endsWith(".sql"))
      .sort();

    const ran: string[] = [];
    for (const file of files) {
      if (appliedNames.has(file)) {
        continue;
      }
      const sql = await readFile(join(migrationsDir, file), "utf8");
      const up = sql
        .split("-- Down Migration")[0]
        .replace("-- Up Migration", "");
      const client = await pool.connect();
      try {
        await client.query("begin");
        await client.query(up);
        await client.query("insert into schema_migrations (name) values ($1)", [
          file,
        ]);
        await client.query("commit");
        ran.push(file);
      } catch (error) {
        await client.query("rollback");
        throw error;
      } finally {
        client.release();
      }
    }
    return ran;
  } finally {
    await pool.end();
  }
}

// Runs only when this file is started directly, for example by `npm run db:migrate`.
if (import.meta.url === pathToFileURL(process.argv[1] ?? "").href) {
  const ran = await migrate(config.databaseUrl);
  process.stdout.write(
    ran.length > 0 ? `Applied: ${ran.join(", ")}\n` : "No new migrations.\n",
  );
}
