CREATE TABLE IF NOT EXISTS "roles" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "title" varchar NOT NULL,
    "note" varchar,
    "company" serial,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "role_item" (
    "id" serial PRIMARY KEY,
    "roles" serial NOT NULL,
    "app" varchar NOT NULL,
    "value" bool DEFAULT false
);

CREATE TABLE IF NOT EXISTS "apps" (
    "id" serial PRIMARY KEY,
    "title" varchar NOT NULL,
    "code" varchar NOT NULL UNIQUE,
    "parent" varchar,
    "level" int DEFAULT 1
);

CREATE UNIQUE INDEX ON "role_item" ("roles", "app");

ALTER TABLE "accounts" ADD COLUMN IF NOT EXISTS "role" INTEGER;

ALTER TABLE "accounts" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id") ON DELETE SET NULL;

ALTER TABLE "roles" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "role_item" ADD FOREIGN KEY ("roles") REFERENCES "roles" ("id") ON DELETE CASCADE;

ALTER TABLE "role_item" ADD FOREIGN KEY ("app") REFERENCES "apps" ("code") ON DELETE CASCADE;

