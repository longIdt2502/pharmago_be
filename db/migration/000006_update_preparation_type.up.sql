ALTER TABLE "preparation_type" ADD COLUMN IF NOT EXISTS "company" serial;
ALTER TABLE "preparation_type" ALTER COLUMN "company" DROP NOT NULL;
UPDATE "preparation_type" SET "company" = NULL;
ALTER TABLE "preparation_type" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "preparation_type" ADD COLUMN IF NOT EXISTS "user_created" serial;
ALTER TABLE "preparation_type" ALTER COLUMN "user_created" DROP NOT NULL;
UPDATE "preparation_type" SET "user_created" = NULL;
ALTER TABLE "preparation_type" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "preparation_type" ADD COLUMN IF NOT EXISTS "user_updated" serial;
ALTER TABLE "preparation_type" ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE "preparation_type" SET "user_updated" = NULL;
ALTER TABLE "preparation_type" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "preparation_type" ADD COLUMN IF NOT EXISTS "created_at" timestamptz NOT NULL DEFAULT (now());
ALTER TABLE "preparation_type" ADD COLUMN IF NOT EXISTS "updated_at" timestamptz;

ALTER TABLE "preparation_type" ADD COLUMN IF NOT EXISTS "description" varchar;
