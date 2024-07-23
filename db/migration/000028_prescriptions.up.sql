CREATE TABLE "prescriptions" (
  "id" serial PRIMARY KEY,
  "uuid" uuid UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "symptoms" varchar,
  "diagnostic" varchar,
  "customer" serial,
  "doctor" serial,
  "company" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "prescription_item" (
  "id" serial PRIMARY KEY,
  "prescription_uuid" uuid,
  "variant" serial,
  "lieu_dung" varchar,
  "quantity" int NOT NULL DEFAULT 0
);

COMMENT ON COLUMN "prescriptions"."symptoms" IS 'Triệu chứng bệnh';

COMMENT ON COLUMN "prescriptions"."diagnostic" IS 'Chuẩn đoán bệnh';

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE "prescriptions" SET "user_updated" = NULL;
ALTER TABLE "prescriptions" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ALTER COLUMN "customer" DROP NOT NULL;
UPDATE "prescriptions" SET "customer" = NULL;
ALTER TABLE "prescriptions" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "prescription_item" ADD FOREIGN KEY ("prescription_uuid") REFERENCES "prescriptions" ("uuid");

ALTER TABLE "prescription_item" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD COLUMN IF NOT EXISTS "prescription" uuid;
ALTER TABLE "medical_bills" ALTER COLUMN "prescription" DROP NOT NULL;
UPDATE "medical_bills" SET "prescription" = NULL;
ALTER TABLE "medical_bills" ADD FOREIGN KEY ("prescription") REFERENCES "prescriptions" ("uuid") ON DELETE SET NULL;

ALTER TABLE IF EXISTS "appointment_schedule_drug" DROP CONSTRAINT "appointment_schedule_drug_as_uuid_fkey";
ALTER TABLE IF EXISTS "appointment_schedule_drug" DROP CONSTRAINT "appointment_schedule_drug_mb_uuid_fkey";
ALTER TABLE IF EXISTS "appointment_schedule_drug" DROP CONSTRAINT "appointment_schedule_drug_variant_fkey";

DROP TABLE IF EXISTS "appointment_schedule_drug";