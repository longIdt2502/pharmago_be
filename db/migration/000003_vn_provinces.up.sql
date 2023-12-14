CREATE TABLE "administrative_regions" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "code_name" varchar NOT NULL,
    "code_name_en" varchar NOT NULL
);

CREATE TABLE "administrative_units" (
    "id" bigserial PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "full_name_en" varchar NOT NULL,
    "short_name" varchar NOT NULL,
    "short_name_en" varchar NOT NULL,
    "code_name" varchar NOT NULL,
    "code_name_en" varchar NOT NULL
);

CREATE TABLE "provinces" (
     "code" varchar(20) UNIQUE PRIMARY KEY,
     "name" varchar NOT NULL,
     "name_en" varchar NOT NULL,
     "full_name" varchar NOT NULL,
     "full_name_en" varchar NOT NULL,
     "code_name" varchar NOT NULL,
     "administrative_unit_id" bigserial,
     "administrative_region_id" bigserial
);

CREATE TABLE "districts" (
     "code" varchar UNIQUE PRIMARY KEY,
     "name" varchar NOT NULL,
     "name_en" varchar NOT NULL,
     "full_name" varchar NOT NULL,
     "full_name_en" varchar NOT NULL,
     "code_name" varchar NOT NULL,
     "province_code" varchar,
     "administrative_unit_id" bigserial
);

CREATE TABLE "wards" (
     "code" varchar UNIQUE PRIMARY KEY,
     "name" varchar NOT NULL,
     "name_en" varchar NOT NULL,
     "full_name" varchar NOT NULL,
     "full_name_en" varchar NOT NULL,
     "code_name" varchar NOT NULL,
     "district_code" varchar,
     "administrative_unit_id" bigserial
);

CREATE INDEX ON "provinces" ("administrative_unit_id");

CREATE INDEX ON "provinces" ("administrative_region_id");

CREATE INDEX ON "districts" ("province_code");

CREATE INDEX ON "districts" ("administrative_unit_id");

CREATE INDEX ON "wards" ("district_code");

CREATE INDEX ON "wards" ("administrative_unit_id");

ALTER TABLE "provinces" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "provinces" ADD FOREIGN KEY ("administrative_region_id") REFERENCES "administrative_regions" ("id");

ALTER TABLE "districts" ADD FOREIGN KEY ("province_code") REFERENCES "provinces" ("code");

ALTER TABLE "districts" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "wards" ADD FOREIGN KEY ("district_code") REFERENCES "districts" ("code");

ALTER TABLE "wards" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "address" ADD COLUMN "province" varchar NULL DEFAULT NULL;

ALTER TABLE "address" ADD COLUMN "district" varchar NULL DEFAULT NULL;

ALTER TABLE "address" ADD COLUMN "ward" varchar NULL DEFAULT NULL;

CREATE INDEX ON "address" ("province");

CREATE INDEX ON "address" ("district");

CREATE INDEX ON "address" ("ward");

ALTER TABLE "address" ADD FOREIGN KEY ("province") REFERENCES "provinces" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("district") REFERENCES "districts" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("ward") REFERENCES "wards" ("code");

ALTER TABLE "companies" ADD COLUMN "address" bigserial NOT NULL;

ALTER TABLE "companies" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");