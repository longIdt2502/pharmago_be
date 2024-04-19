CREATE TABLE "medical_records" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "customer" serial NOT NULL,
  "weight" float,
  "long" float,
  "symptom" varchar NOT NULL,
  "diagnostic" varchar NOT NULL,
  "result" varchar NOT NULL,
  "doctor" serial,
  "re_examination" int NOT NULL DEFAULT 0,
  "note" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "user_created" serial,
  "user_updated" serial
);

CREATE TABLE "medical_record_variant" (
  "id" serial PRIMARY KEY,
  "medical_record" serial NOT NULL,
  "variant" serial NOT NULL
);

COMMENT ON COLUMN "medical_records"."symptom" IS 'Triệu chứng';

COMMENT ON COLUMN "medical_records"."diagnostic" IS 'Chuẩn đoán';

COMMENT ON COLUMN "medical_records"."result" IS 'Kết luận';

ALTER TABLE "medical_records" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ALTER COLUMN "user_updated" DROP NOT NULL;
ALTER TABLE "medical_records" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_record_variant" ADD FOREIGN KEY ("medical_record") REFERENCES "medical_records" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_record_variant" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;
