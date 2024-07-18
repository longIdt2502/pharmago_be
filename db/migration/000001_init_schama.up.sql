-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-12-25T02:57:54.550Z
CREATE TABLE "accounts" (
    "id" serial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "hashed_password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "type" serial NOT NULL,
    "is_verify" boolean NOT NULL DEFAULT false,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_media" (
    "id" serial PRIMARY KEY,
    "account" serial NOT NULL,
    "media" serial NOT NULL
);

CREATE TABLE "account_company" (
    "id" serial PRIMARY KEY,
    "account" serial NOT NULL,
    "company" serial NOT NULL
);

CREATE TABLE "account_type" (
    "id" serial PRIMARY KEY,
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
    "id" serial PRIMARY KEY,
    "username" varchar NOT NULL,
    "email" varchar NOT NULL,
    "secret_code" varchar NOT NULL,
    "is_used" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "expired_at" timestamptz NOT NULL DEFAULT ((now() + interval '15 minutes'))
);

CREATE TABLE "companies" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "code" varchar UNIQUE NOT NULL,
    "tax_code" varchar,
    "phone" varchar,
    "description" varchar,
    "address" serial,
    "oa_id" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "owner" serial NOT NULL
);

CREATE TABLE "address" (
    "id" serial PRIMARY KEY,
    "lat" float NOT NULL,
    "lng" float NOT NULL,
    "province" varchar,
    "district" varchar,
    "ward" varchar,
    "title" varchar NOT NULL,
    "user_created" serial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "administrative_regions" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "code_name" varchar NOT NULL,
    "code_name_en" varchar NOT NULL
);

CREATE TABLE "administrative_units" (
    "id" serial PRIMARY KEY,
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
    "administrative_unit_id" serial,
    "administrative_region_id" serial
);

CREATE TABLE "districts" (
    "code" varchar UNIQUE PRIMARY KEY,
    "name" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "full_name_en" varchar NOT NULL,
    "code_name" varchar NOT NULL,
    "province_code" varchar,
    "administrative_unit_id" serial
);

CREATE TABLE "wards" (
    "code" varchar UNIQUE PRIMARY KEY,
    "name" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "full_name_en" varchar NOT NULL,
    "code_name" varchar NOT NULL,
    "district_code" varchar,
    "administrative_unit_id" serial
);

CREATE TABLE "products" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "code" varchar NOT NULL,
    "product_category" serial,
    "type" serial,
    "brand" serial NOT NULL,
    "unit" serial NOT NULL,
    "ta_duoc" varchar(255),
    "nong_do" varchar(255),
    "lieu_dung" varchar(255) NOT NULL,
    "chi_dinh" varchar(255) NOT NULL,
    "chong_chi_dinh" varchar(255),
    "cong_dung" varchar(255) NOT NULL,
    "tac_dung_phu" varchar(255) NOT NULL,
    "than_trong" varchar(255) NOT NULL,
    "tuong_tac" varchar(255),
    "bao_quan" varchar(255) NOT NULL,
    "dong_goi" varchar(255) NOT NULL,
    "phan_loai" varchar DEFAULT null,
    "dang_bao_che" varchar NOT NULL,
    "tieu_chuan_sx" varchar NOT NULL,
    "cong_ty_sx" serial NOT NULL,
    "cong_ty_dk" serial NOT NULL,
    "active" boolean NOT NULL DEFAULT true,
    "company" serial NOT NULL,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE
    "products"
ALTER COLUMN
    "brand" DROP NOT NULL;

ALTER TABLE
    "products"
ALTER COLUMN
    "product_category" DROP NOT NULL;

ALTER TABLE
    "products"
ALTER COLUMN
    "type" DROP NOT NULL;

CREATE TABLE "products_bank" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "code" varchar NOT NULL,
    "ta_duoc" varchar(255),
    "nong_do" varchar(255),
    "lieu_dung" varchar(255) NOT NULL,
    "chi_dinh" varchar(255) NOT NULL,
    "chong_chi_dinh" varchar(255),
    "cong_dung" varchar(255) NOT NULL,
    "tac_dung_phu" varchar(255) NOT NULL,
    "than_trong" varchar(255) NOT NULL,
    "tuong_tac" varchar(255),
    "bao_quan" varchar(255) NOT NULL,
    "dong_goi" varchar(255) NOT NULL,
    "phan_loai" varchar DEFAULT null,
    "dang_bao_che" varchar NOT NULL,
    "tieu_chuan_sx" varchar NOT NULL,
    "cong_ty_sx" serial NOT NULL,
    "cong_ty_dk" serial NOT NULL
);

CREATE TABLE "price_list" (
    "id" serial PRIMARY KEY,
    "variant_code" varchar UNIQUE NOT NULL,
    "variant_name" varchar NOT NULL,
    "price_import" float NOT NULL,
    "price_sell" float NOT NULL,
    "unit" serial NOT NULL,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "price_list_log" (
    "id" serial PRIMARY KEY,
    "old_price_import" float NOT NULL,
    "new_price_import" float NOT NULL,
    "old_price_sell" float NOT NULL,
    "new_price_sell" float NOT NULL,
    "price_list" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company_pharma" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "code" varchar(255),
    "country" varchar(255),
    "address" varchar(255),
    "company_pharma_type" varchar DEFAULT null,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company_pharma_type" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "title" varchar(255) NOT NULL
);

CREATE TABLE "product_media" (
    "id" serial PRIMARY KEY,
    "product" serial NOT NULL,
    "media" serial NOT NULL
);

CREATE TABLE "product_categories" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "user_created" serial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "company" serial NOT NULL
);

CREATE TABLE "product_brand" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "user_created" serial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "company" serial NOT NULL
);

CREATE TABLE "product_type" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "user_created" serial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "company" serial NOT NULL
);

CREATE TABLE "classify" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "preparation_type" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "production_standard" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "ingredient" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "weight" float NOT NULL DEFAULT 0,
    "unit" varchar NOT NULL,
    "product" serial NOT NULL
);

CREATE TABLE "units" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "sell_price" float NOT NULL DEFAULT 0,
    "import_price" float NOT NULL DEFAULT 0,
    "weight" float,
    "weight_unit" varchar,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "unit_changes" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "value" bigint NOT NULL DEFAULT 0,
    "sell_price" float NOT NULL DEFAULT 0,
    "unit" serial NOT NULL,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variants" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "code" varchar UNIQUE NOT NULL,
    "barcode" varchar NOT NULL,
    "decision_number" varchar(255) UNIQUE NOT NULL,
    "register_number" varchar(255) UNIQUE NOT NULL,
    "longevity" varchar(255) NOT NULL,
    "vat" float NOT NULL DEFAULT 0,
    "product" serial NOT NULL,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variant_media" (
    "id" serial PRIMARY KEY,
    "variant" serial NOT NULL,
    "media" serial NOT NULL
);

CREATE TABLE "orders" (
    "id" serial PRIMARY KEY,
    "vat" float NOT NULL DEFAULT 0,
    "total_price" float NOT NULL DEFAULT 0,
    "description" varchar,
    "customer" serial,
    "status" serial,
    "type" serial,
    "ticket" serial,
    "qr" serial
);

CREATE TABLE "order_type" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "title" varchar NOT NULL
);

CREATE TABLE "order_status" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "title" varchar NOT NULL
);

CREATE TABLE "order_items" (
    "id" serial PRIMARY KEY,
    "order" serial NOT NULL,
    "variant" serial NOT NULL,
    "value" int NOT NULL DEFAULT 0,
    "expired_at" timestamptz,
    "manufactured_at" timestamptz
);

CREATE TABLE "customers" (
    "id" serial PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "code" varchar NOT NULL,
    "company" serial NOT NULL,
    "address" serial,
    "email" varchar,
    "birthday" timestamptz,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE
    "customers"
ALTER COLUMN
    "address" DROP NOT NULL;

CREATE TABLE "medias" (
    "id" serial PRIMARY KEY,
    "media_url" varchar NOT NULL
);

CREATE TABLE "warehouses" (
    "id" serial PRIMARY KEY,
    "address" serial,
    "companies" serial
);

CREATE TABLE "tickets" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "type" serial,
    "status" serial,
    "note" varchar,
    "qr" serial,
    "export_to" serial,
    "import_from" serial,
    "total_price" float NOT NULL DEFAULT 0,
    "warehouse" serial NOT NULL,
    "user_created" serial NOT NULL,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE
    "tickets"
ALTER COLUMN
    "export_to" DROP NOT NULL;

CREATE TABLE "ticket_type" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "title" varchar NOT NULL
);

CREATE TABLE "ticket_status" (
    "id" serial PRIMARY KEY,
    "code" varchar NOT NULL,
    "title" varchar NOT NULL
);

CREATE TABLE "consignment" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "quantity" int NOT NULL DEFAULT 0,
    "inventory" int NOT NULL DEFAULT 0,
    "ticket" serial,
    "product" serial,
    "expired_at" timestamptz NOT NULL,
    "producted_at" timestamptz NOT NULL,
    "is_available" bool NOT NULL DEFAULT false,
    "user_created" serial,
    "user_updated" serial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "consignment_log" (
    "id" serial PRIMARY KEY,
    "consignment" serial NOT NULL,
    "inventory" int NOT NULL DEFAULT 0,
    "amount_change" int NOT NULL DEFAULT 0,
    "user_created" serial,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "suplier" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar(255) NOT NULL,
    "deputy_name" varchar(255) NOT NULL,
    "phone" varchar(255) NOT NULL,
    "email" varchar(255),
    "address" serial,
    "company" serial
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

CREATE INDEX ON "products" ("unit");

CREATE INDEX ON "products" ("id", "unit");

CREATE INDEX ON "company_pharma" ("company_pharma_type");

CREATE UNIQUE INDEX ON "company_pharma" ("name", "company_pharma_type");

CREATE INDEX ON "product_media" ("product");

CREATE INDEX ON "product_media" ("media");

CREATE UNIQUE INDEX ON "product_media" ("product", "media");

CREATE INDEX ON "variant_media" ("variant");

CREATE INDEX ON "variant_media" ("media");

CREATE UNIQUE INDEX ON "variant_media" ("variant", "media");

CREATE INDEX ON "orders" ("qr");

CREATE INDEX ON "orders" ("id", "qr");

CREATE INDEX ON "customers" ("address");

CREATE UNIQUE INDEX ON "customers" ("id", "address");

CREATE INDEX ON "tickets" ("qr");

CREATE UNIQUE INDEX ON "tickets" ("id", "qr");

ALTER TABLE
    "accounts"
ADD
    FOREIGN KEY ("type") REFERENCES "account_type" ("id") ON DELETE CASCADE;

ALTER TABLE
    "account_media"
ADD
    FOREIGN KEY ("account") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "account_media"
ADD
    FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE CASCADE;

ALTER TABLE
    "account_company"
ADD
    FOREIGN KEY ("account") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "account_company"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE
    "sessions"
ADD
    FOREIGN KEY ("username") REFERENCES "accounts" ("username") ON DELETE CASCADE;

ALTER TABLE
    "verifies"
ADD
    FOREIGN KEY ("username") REFERENCES "accounts" ("username") ON DELETE CASCADE;

ALTER TABLE
    "companies"
ADD
    FOREIGN KEY ("owner") REFERENCES "accounts" ("id");

ALTER TABLE
    "companies"
ADD
    FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "address"
ADD
    FOREIGN KEY ("province") REFERENCES "provinces" ("code");

ALTER TABLE
    "address"
ADD
    FOREIGN KEY ("district") REFERENCES "districts" ("code");

ALTER TABLE
    "address"
ADD
    FOREIGN KEY ("ward") REFERENCES "wards" ("code");

ALTER TABLE
    "address"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE
    "provinces"
ADD
    FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE
    "provinces"
ADD
    FOREIGN KEY ("administrative_region_id") REFERENCES "administrative_regions" ("id");

ALTER TABLE
    "districts"
ADD
    FOREIGN KEY ("province_code") REFERENCES "provinces" ("code");

ALTER TABLE
    "districts"
ADD
    FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE
    "wards"
ADD
    FOREIGN KEY ("district_code") REFERENCES "districts" ("code");

ALTER TABLE
    "wards"
ADD
    FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("product_category") REFERENCES "product_categories" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("type") REFERENCES "product_type" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("brand") REFERENCES "product_brand" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("unit") REFERENCES "units" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("phan_loai") REFERENCES "classify" ("code") ON DELETE
SET
    NULL;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("dang_bao_che") REFERENCES "preparation_type" ("code") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("tieu_chuan_sx") REFERENCES "production_standard" ("code") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("cong_ty_sx") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("cong_ty_dk") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "products_bank"
ADD
    FOREIGN KEY ("phan_loai") REFERENCES "classify" ("code") ON DELETE
SET
    NULL;

ALTER TABLE
    "products_bank"
ADD
    FOREIGN KEY ("dang_bao_che") REFERENCES "preparation_type" ("code") ON DELETE CASCADE;

ALTER TABLE
    "products_bank"
ADD
    FOREIGN KEY ("tieu_chuan_sx") REFERENCES "production_standard" ("code") ON DELETE CASCADE;

ALTER TABLE
    "products_bank"
ADD
    FOREIGN KEY ("cong_ty_sx") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE
    "products_bank"
ADD
    FOREIGN KEY ("cong_ty_dk") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE
    "price_list"
ADD
    FOREIGN KEY ("variant_code") REFERENCES "variants" ("code") ON DELETE CASCADE;

ALTER TABLE
    "price_list"
ADD
    FOREIGN KEY ("unit") REFERENCES "units" ("id") ON DELETE CASCADE;

ALTER TABLE
    "price_list"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "price_list"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "price_list_log"
ADD
    FOREIGN KEY ("price_list") REFERENCES "price_list" ("id") ON DELETE CASCADE;

ALTER TABLE
    "price_list_log"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "company_pharma"
ADD
    FOREIGN KEY ("company_pharma_type") REFERENCES "company_pharma_type" ("code") ON DELETE
SET
    NULL;

ALTER TABLE
    "product_media"
ADD
    FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE
    "product_media"
ADD
    FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE CASCADE;

ALTER TABLE
    "product_categories"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE
    "product_categories"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE
    "product_brand"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE
    "product_brand"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE
    "product_type"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE
    "product_type"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE
    "ingredient"
ADD
    FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE
    "units"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "units"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "unit_changes"
ADD
    FOREIGN KEY ("unit") REFERENCES "units" ("id") ON DELETE CASCADE;

ALTER TABLE
    "unit_changes"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "unit_changes"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "variants"
ADD
    FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE
    "variants"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "variants"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "variant_media"
ADD
    FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE CASCADE;

ALTER TABLE
    "variant_media"
ADD
    FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE CASCADE;

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("status") REFERENCES "order_status" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("type") REFERENCES "order_type" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("ticket") REFERENCES "tickets" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "orders"
ADD
    FOREIGN KEY ("qr") REFERENCES "medias" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "order_items"
ADD
    FOREIGN KEY ("order") REFERENCES "orders" ("id") ON DELETE CASCADE;

ALTER TABLE
    "order_items"
ADD
    FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE CASCADE;

ALTER TABLE
    "customers"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "customers" ALTER COLUMN "address" DROP NOT NULL;
ALTER TABLE
    "customers"
ADD
    FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "customers"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "customers"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "warehouses"
ADD
    FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "warehouses"
ADD
    FOREIGN KEY ("companies") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("type") REFERENCES "ticket_type" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("status") REFERENCES "ticket_status" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("qr") REFERENCES "medias" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("export_to") REFERENCES "address" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("import_from") REFERENCES "address" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("warehouse") REFERENCES "warehouses" ("id") ON DELETE CASCADE;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE
    "tickets"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "consignment"
ADD
    FOREIGN KEY ("ticket") REFERENCES "tickets" ("id") ON DELETE CASCADE;

ALTER TABLE
    "consignment"
ADD
    FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE
    "consignment"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "consignment"
ADD
    FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "consignment_log"
ADD
    FOREIGN KEY ("consignment") REFERENCES "consignment" ("id") ON DELETE CASCADE;

ALTER TABLE
    "consignment_log"
ADD
    FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "suplier"
ADD
    FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE
SET
    NULL;

ALTER TABLE
    "suplier"
ADD
    FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;