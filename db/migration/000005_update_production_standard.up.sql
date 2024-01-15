ALTER TABLE "production_standard" ADD COLUMN IF NOT EXISTS "company" serial;
ALTER TABLE "production_standard" ALTER COLUMN "company" DROP NOT NULL;
UPDATE "production_standard" SET "company" = NULL;
ALTER TABLE "production_standard" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "production_standard" ADD COLUMN IF NOT EXISTS "user_created" serial;
ALTER TABLE "production_standard" ALTER COLUMN "user_created" DROP NOT NULL;
UPDATE "production_standard" SET "user_created" = NULL;
ALTER TABLE "production_standard" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "production_standard" ADD COLUMN IF NOT EXISTS "user_updated" serial;
ALTER TABLE "production_standard" ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE "production_standard" SET "user_updated" = NULL;
ALTER TABLE "production_standard" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "production_standard" ADD COLUMN IF NOT EXISTS "created_at" timestamptz NOT NULL DEFAULT (now());
ALTER TABLE "production_standard" ADD COLUMN IF NOT EXISTS "updated_at" timestamptz;

ALTER TABLE "production_standard" ADD COLUMN IF NOT EXISTS "description" varchar;
