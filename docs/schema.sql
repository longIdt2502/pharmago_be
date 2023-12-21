-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-12-21T08:00:08.737Z

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "type" bigserial NOT NULL,
  "is_verify" boolean NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_media" (
  "id" bigserial PRIMARY KEY,
  "account" bigserial NOT NULL,
  "media" bigserial NOT NULL
);

CREATE TABLE "account_company" (
  "id" bigserial PRIMARY KEY,
  "account" bigserial NOT NULL,
  "company" bigserial NOT NULL
);

CREATE TABLE "account_type" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verifies" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT ((now() + interval '15 minutes'))
);

CREATE TABLE "companies" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "tax_code" varchar,
  "phone" varchar,
  "description" varchar,
  "address" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "owner" bigserial NOT NULL
);

CREATE TABLE "address" (
  "id" bigserial PRIMARY KEY,
  "lat" numeric NOT NULL,
  "lng" numeric NOT NULL,
  "province" varchar,
  "district" varchar,
  "ward" varchar,
  "title" varchar NOT NULL,
  "user_created" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

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

CREATE TABLE "warehouses" (
  "id" bigserial PRIMARY KEY,
  "address" bigserial NOT NULL,
  "companies" bigserial NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "product_category" serial,
  "type" serial,
  "unit" bigserial NOT NULL,
  "taDuoc" varchar(255),
  "nongDo" varchar(255),
  "lieuDung" varchar(255) NOT NULL,
  "chiDinh" varchar(255) NOT NULL,
  "chongChiDinh" varchar(255),
  "congDung" varchar(255) NOT NULL,
  "tacDungPhu" varchar(255) NOT NULL,
  "thanTrong" varchar(255) NOT NULL,
  "tuongTac" varchar(255),
  "baoQuan" varchar(255) NOT NULL,
  "dongGoi" varchar(255) NOT NULL,
  "noiSx" varchar(255) NOT NULL,
  "congTySx" varchar(255) NOT NULL,
  "congTyDk" varchar(255) NOT NULL,
  "company" bigserial,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "price_list" (
  "id" bigserial PRIMARY KEY,
  "variant_code" varchar UNIQUE NOT NULL,
  "variant_name" varchar NOT NULL,
  "price_import" float NOT NULL,
  "price_sell" float NOT NULL,
  "unit" bigserial NOT NULL,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "price_list_log" (
  "id" bigserial PRIMARY KEY,
  "old_price_import" float NOT NULL,
  "new_price_import" float NOT NULL,
  "old_price_sell" float NOT NULL,
  "new_price_sell" float NOT NULL,
  "price_list" bigserial NOT NULL,
  "user_updated" bigserial NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_media" (
  "id" bigserial PRIMARY KEY,
  "product" bigserial,
  "media" bigserial
);

CREATE TABLE "product_categories" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_type" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "units" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "sell_price" numeric NOT NULL DEFAULT 0,
  "import_price" numeric NOT NULL DEFAULT 0,
  "weight" numeric,
  "weight_unit" varchar,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "unit_changes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "value" bigint NOT NULL DEFAULT 0,
  "sell_price" numeric NOT NULL DEFAULT 0,
  "unit" bigserial,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variants" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "barcode" varchar NOT NULL,
  "decision_number" varchar(255) UNIQUE NOT NULL,
  "register_number" varchar(255) UNIQUE NOT NULL,
  "longevity" varchar(255) NOT NULL,
  "vat" numeric NOT NULL DEFAULT 0,
  "product" bigserial,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variant_media" (
  "id" bigserial PRIMARY KEY,
  "variant" bigserial NOT NULL,
  "media" bigserial NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "vat" numeric NOT NULL DEFAULT 0,
  "total_price" numeric NOT NULL DEFAULT 0,
  "description" varchar,
  "customer" bigserial,
  "status" bigserial NOT NULL,
  "type" bigserial,
  "ticket" bigserial,
  "qr" bigserial
);

CREATE TABLE "order_type" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "order_status" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "order" bigserial,
  "variant" bigserial,
  "value" int NOT NULL DEFAULT 0,
  "expired_at" timestamptz,
  "manufactured_at" timestamptz
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "company" bigserial,
  "address" bigserial,
  "email" varchar,
  "birthday" timestamptz,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "type" bigserial NOT NULL,
  "status" bigserial NOT NULL,
  "note" varchar,
  "qr" bigserial,
  "export_from" bigserial NOT NULL,
  "import_to" bigserial NOT NULL,
  "user_created" bigserial NOT NULL,
  "user_updated" bigserial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ticket_type" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "ticket_status" (
  "id" bigserial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "medias" (
  "id" bigserial PRIMARY KEY,
  "media_url" varchar NOT NULL
);

CREATE INDEX ON "address" ("province");

CREATE INDEX ON "address" ("district");

CREATE INDEX ON "address" ("ward");

CREATE INDEX ON "provinces" ("administrative_unit_id");

CREATE INDEX ON "provinces" ("administrative_region_id");

CREATE INDEX ON "districts" ("province_code");

CREATE INDEX ON "districts" ("administrative_unit_id");

CREATE INDEX ON "wards" ("district_code");

CREATE INDEX ON "wards" ("administrative_unit_id");

CREATE INDEX ON "warehouses" ("address");

CREATE UNIQUE INDEX ON "warehouses" ("id", "address");

CREATE INDEX ON "products" ("unit");

CREATE UNIQUE INDEX ON "products" ("id", "unit");

CREATE INDEX ON "product_media" ("product");

CREATE INDEX ON "product_media" ("media");

CREATE UNIQUE INDEX ON "product_media" ("product", "media");

CREATE INDEX ON "variant_media" ("variant");

CREATE INDEX ON "variant_media" ("media");

CREATE UNIQUE INDEX ON "variant_media" ("variant", "media");

CREATE INDEX ON "orders" ("qr");

CREATE UNIQUE INDEX ON "orders" ("id", "qr");

CREATE INDEX ON "customers" ("address");

CREATE UNIQUE INDEX ON "customers" ("id", "address");

CREATE INDEX ON "tickets" ("qr");

CREATE UNIQUE INDEX ON "tickets" ("id", "qr");

ALTER TABLE "accounts" ADD FOREIGN KEY ("type") REFERENCES "account_type" ("code");

ALTER TABLE "account_media" ADD FOREIGN KEY ("account") REFERENCES "accounts" ("id");

ALTER TABLE "account_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id");

ALTER TABLE "account_company" ADD FOREIGN KEY ("account") REFERENCES "accounts" ("id");

ALTER TABLE "account_company" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "accounts" ("username");

ALTER TABLE "verifies" ADD FOREIGN KEY ("username") REFERENCES "accounts" ("username");

ALTER TABLE "companies" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");

ALTER TABLE "companies" ADD FOREIGN KEY ("owner") REFERENCES "accounts" ("id");

ALTER TABLE "address" ADD FOREIGN KEY ("province") REFERENCES "provinces" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("district") REFERENCES "districts" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("ward") REFERENCES "wards" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "provinces" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "provinces" ADD FOREIGN KEY ("administrative_region_id") REFERENCES "administrative_regions" ("id");

ALTER TABLE "districts" ADD FOREIGN KEY ("province_code") REFERENCES "provinces" ("code");

ALTER TABLE "districts" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "wards" ADD FOREIGN KEY ("district_code") REFERENCES "districts" ("code");

ALTER TABLE "wards" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "warehouses" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");

ALTER TABLE "warehouses" ADD FOREIGN KEY ("companies") REFERENCES "companies" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("product_category") REFERENCES "product_categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("type") REFERENCES "product_type" ("id");

ALTER TABLE "units" ADD FOREIGN KEY ("id") REFERENCES "products" ("unit");

ALTER TABLE "products" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "price_list" ADD FOREIGN KEY ("variant_code") REFERENCES "variants" ("code");

ALTER TABLE "price_list" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id");

ALTER TABLE "price_list" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "price_list" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "price_list_log" ADD FOREIGN KEY ("price_list") REFERENCES "price_list" ("id");

ALTER TABLE "price_list_log" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "product_media" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "product_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "units" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "units" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id");

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "variants" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "variants" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "variants" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "variant_media" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id");

ALTER TABLE "variant_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("status") REFERENCES "order_status" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("type") REFERENCES "order_type" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("ticket") REFERENCES "tickets" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("qr") REFERENCES "medias" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("type") REFERENCES "ticket_type" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("status") REFERENCES "ticket_status" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("qr") REFERENCES "medias" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("export_from") REFERENCES "warehouses" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("import_to") REFERENCES "warehouses" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");
