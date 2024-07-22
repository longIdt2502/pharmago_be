CREATE TABLE "medical_bills" (
  "id" serial PRIMARY KEY,
  "uuid" uuid UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "customer" serial,
  "company" serial,
  "doctor" serial,
  "symptoms" varchar,
  "diagnostic" varchar,
  "qr_code_url" varchar,
  "is_done" bool NOT NULL,
  "meeting_at" timestamp NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_service" ALTER COLUMN "as_uuid" DROP NOT NULL;
ALTER TABLE "appointment_schedule_service" ADD COLUMN IF NOT EXISTS "mb_uuid" UUID;
ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("mb_uuid") REFERENCES "medical_bills" ("uuid");

ALTER TABLE "appointment_schedule_url" ALTER COLUMN "as_uuid" DROP NOT NULL;
ALTER TABLE "appointment_schedule_url" ADD COLUMN IF NOT EXISTS "mb_uuid" UUID;
ALTER TABLE "appointment_schedule_url" ADD FOREIGN KEY ("mb_uuid") REFERENCES "medical_bills" ("uuid");

ALTER TABLE "appointment_schedule_drug" ALTER COLUMN "as_uuid" DROP NOT NULL;
ALTER TABLE "appointment_schedule_drug" ADD COLUMN IF NOT EXISTS "mb_uuid" UUID;
ALTER TABLE "appointment_schedule_drug" ADD FOREIGN KEY ("mb_uuid") REFERENCES "medical_bills" ("uuid");
