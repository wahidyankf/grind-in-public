import { buildApp } from "./app.js";
import { config } from "./config.js";
import { createPool } from "./db.js";
import { createSqlAuthStore } from "./sql_store.js";

const pool = createPool(config.databaseUrl);
const app = buildApp({
  store: createSqlAuthStore(pool),
  config,
  logger: true,
});

await app.listen({ port: config.port });
