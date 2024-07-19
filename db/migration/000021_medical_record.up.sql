CREATE TYPE "medical_record_link_type" AS ENUM (
  'test',
  'patient',
  'diagnostic'
);

CREATE TABLE "medical_record_link" (
  "id" serial PRIMARY KEY,
  "uuid" uuid NOT NULL,
  "type" medical_record_link_type NOT NULL,
  "title" varchar,
  "url" varchar NOT NULL,
  "customer" serial,
  "appointment_schedule" uuid,
  "user_created" serial,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE CASCADE;

ALTER TABLE "medical_record_link" ALTER COLUMN "appointment_schedule" DROP NOT NULL;
UPDATE "medical_record_link" SET "appointment_schedule" = NULL;
ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("appointment_schedule") REFERENCES "appointment_schedules" ("uuid") ON DELETE SET NULL;

ALTER TABLE "medical_record_link" ALTER COLUMN "user_created" DROP NOT NULL;
UPDATE "medical_record_link" SET "user_created" = NULL;
ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;