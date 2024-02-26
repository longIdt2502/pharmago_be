ALTER TABLE "product_categories" ADD COLUMN IF NOT EXISTS "user_updated" serial;
ALTER TABLE "product_categories" ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE "product_categories" SET "user_updated" = NULL;
ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "product_categories" ADD COLUMN IF NOT EXISTS "created_at" timestamptz NOT NULL DEFAULT (now());
ALTER TABLE "product_categories" ADD COLUMN IF NOT EXISTS "updated_at" timestamptz;

ALTER TABLE "product_categories" ADD COLUMN IF NOT EXISTS "description" varchar;