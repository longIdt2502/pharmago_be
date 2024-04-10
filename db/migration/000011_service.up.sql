ALTER TABLE "companies" ADD COLUMN IF NOT EXISTS "type" varchar NOT NULL DEFAULT 'DRUGSTORE';

CREATE TABLE "company_type" (
  "code" varchar UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar UNIQUE NOT NULL
);

CREATE TABLE "services" (
  "id" serial PRIMARY KEY,
  "image" serial,
  "code" varchar UNIQUE NOT NULL,
  "title" varchar NOT NULL,
  "entity" varchar,
  "staff" serial NOT NULL,
  "frequency" varchar,
  "unit" varchar NOT NULL,
  "price" float NOT NULL DEFAULT 0,
  "description" varchar,
  "company" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "service_variant" (
  "id" serial PRIMARY KEY,
  "service" serial,
  "variant" serial
);

CREATE UNIQUE INDEX ON "service_variant" ("service", "variant");

COMMENT ON COLUMN "company_type"."code" IS '
💸 1 = CLINIC,
✔️ 2 = DRUGSTORE
';

INSERT INTO "company_type" (code, title) VALUES ('CLINIC', 'Phòng khám');
INSERT INTO "company_type" (code, title) VALUES ('DRUGSTORE', 'Nhà thuốc');

ALTER TABLE "companies" ADD FOREIGN KEY ("type") REFERENCES "company_type" ("code") ON DELETE SET NULL;

ALTER TABLE "services" ADD FOREIGN KEY ("staff") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "services" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "services" ALTER COLUMN "image" DROP NOT NULL;
ALTER TABLE "services" ADD FOREIGN KEY ("image") REFERENCES "medias" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "services" ALTER COLUMN "user_updated" DROP NOT NULL;
ALTER TABLE "services" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "service_variant" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE CASCADE;

ALTER TABLE "service_variant" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id");
