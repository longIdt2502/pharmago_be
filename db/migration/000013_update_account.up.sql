CREATE TYPE "gender" AS ENUM (
  'nam',
  'nữ',
  'khác'
);

ALTER TABLE "accounts" ADD COLUMN IF NOT EXISTS "gender" gender;
ALTER TABLE "accounts" ADD COLUMN IF NOT EXISTS "licence" varchar;
ALTER TABLE "accounts" ADD COLUMN IF NOT EXISTS "dob" timestamp;
ALTER TABLE "accounts" ADD COLUMN IF NOT EXISTS "address" serial;
ALTER TABLE "accounts" ALTER COLUMN "address" DROP NOT NULL;
UPDATE "accounts" SET "address" = NULL;
ALTER TABLE "accounts" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;
