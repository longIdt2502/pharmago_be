CREATE TABLE "customer_group" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "company" serial NOT NULL,
  "note" varchar,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "group" SERIAL;
ALTER TABLE "customers" ALTER COLUMN "group" DROP NOT NULL;
UPDATE "customers" SET "group" = NULL;

ALTER TABLE "customers" ADD FOREIGN KEY ("group") REFERENCES "customer_group" ("id") ON DELETE SET NULL;

ALTER TABLE "customer_group" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "customer_group" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "customer_group" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;