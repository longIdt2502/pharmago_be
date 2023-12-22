CREATE TABLE "company_pharma" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "code" varchar(255),
    "country" varchar(255),
    "address" varchar(255),
    "company_pharma_type" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company_pharma_type" (
    "id" bigserial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "title" varchar(255) NOT NULL
);

CREATE INDEX ON "company_pharma" ("company_pharma_type");

CREATE UNIQUE INDEX ON "company_pharma" ("name", "company_pharma_type");

ALTER TABLE "products" DROP COLUMN IF EXISTS "congTySx";
ALTER TABLE "products" ADD COLUMN "congTySx" bigserial;

ALTER TABLE "products" DROP COLUMN IF EXISTS "congTyDk";
ALTER TABLE "products" ADD COLUMN "congTyDk" bigserial;

ALTER TABLE "products" ADD FOREIGN KEY ("congTySx") REFERENCES "company_pharma" ("id");
ALTER TABLE "products" ALTER COLUMN "congTySx" DROP NOT NULL;


ALTER TABLE "products" ADD FOREIGN KEY ("congTyDk") REFERENCES "company_pharma" ("id");
ALTER TABLE "products" ALTER COLUMN "congTyDk" DROP NOT NULL;


ALTER TABLE "company_pharma" ADD FOREIGN KEY ("company_pharma_type") REFERENCES "company_pharma_type" ("code");