ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "title" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "license_date" timestamp;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "contact_name" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "contact_title" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "contact_phone" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "contact_email" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "contact_address" SERIAL;
ALTER TABLE "customers" ALTER COLUMN "contact_address" DROP NOT NULL;
UPDATE "customers" SET "contact_address" = NULL;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "account_number" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "bank_name" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "bank_branch" varchar;
ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "issued_by" varchar;

COMMENT ON COLUMN "customers"."title" IS 'Chức danh';

ALTER TABLE "customers" ADD FOREIGN KEY ("contact_address") REFERENCES "address" ("id") ON DELETE SET NULL;
