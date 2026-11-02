-- Up Migration
-- gen_random_uuid() is built in from Postgres 13, so no extension is needed.
create table users (
  id uuid primary key default gen_random_uuid(),
  email text not null unique,
  password_hash text not null,
  created_at timestamptz not null default now()
);

-- Down Migration
drop table users;
