CREATE TABLE "appointment_schedules" (
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

CREATE TABLE "appointment_schedule_service" (
  "id" serial PRIMARY KEY,
  "as_uuid" uuid NOT NULL,
  "service" serial,
  "order_service" serial
);

CREATE TABLE "appointment_schedule_url" (
  "id" serial PRIMARY KEY,
  "as_uuid" uuid NOT NULL,
  "url" varchar,
  "name_doc" varchar
);

CREATE TABLE "appointment_schedule_drug" (
  "id" serial PRIMARY KEY,
  "as_uuid" uuid NOT NULL,
  "variant" serial,
  "lieu_dung" varchar,
  "quantity" int NOT NULL DEFAULT 0
);

COMMENT ON COLUMN "appointment_schedules"."symptoms" IS 'Triệu chứng bệnh';

COMMENT ON COLUMN "appointment_schedules"."diagnostic" IS 'Chuẩn đoán bệnh';

COMMENT ON COLUMN "appointment_schedules"."is_done" IS 'true: Đã xong, false: Chưa diễn ra';

ALTER TABLE "appointment_schedules" ALTER COLUMN "customer" DROP NOT NULL;
UPDATE "appointment_schedules" SET "customer" = NULL;
ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "appointment_schedules" ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE "appointment_schedules" SET "user_updated" = NULL;
ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("as_uuid") REFERENCES "appointment_schedules" ("uuid");

ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_service" ALTER COLUMN "order_service" DROP NOT NULL;
UPDATE "appointment_schedule_service" SET "order_service" = NULL;
ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("order_service") REFERENCES "orders" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_url" ADD FOREIGN KEY ("as_uuid") REFERENCES "appointment_schedules" ("uuid");

ALTER TABLE "appointment_schedule_drug" ADD FOREIGN KEY ("as_uuid") REFERENCES "appointment_schedules" ("uuid");

ALTER TABLE "appointment_schedule_drug" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;
