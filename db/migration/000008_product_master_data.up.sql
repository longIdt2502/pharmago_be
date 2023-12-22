-- Xoá ràng buộc trên bảng "product_categories"
ALTER TABLE IF EXISTS "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_created_fkey";
ALTER TABLE IF EXISTS "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_updated_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_product_category_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_type_fkey";

DROP TABLE IF EXISTS "product_categories";
DROP TABLE IF EXISTS "product_type";

CREATE TABLE "products_bank" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "code" varchar NOT NULL,
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
    "phanLoai" varchar,
    "dangBaoche" varchar NOT NULL,
    "tieuChuanSx" varchar NOT NULL,
    "congTySx" bigserial NOT NULL,
    "congTyDk" bigserial NOT NULL
);

CREATE TABLE "product_categories" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "user_created" bigserial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_type" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "user_created" bigserial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_brand" (
    "id" serial PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "name" varchar NOT NULL,
    "user_created" bigserial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
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
    "unit" varchar NOT NULL
);

CREATE TABLE "product_ingredient" (
    "id" bigserial PRIMARY KEY,
    "product" bigserial NOT NULL,
    "ingredient" serial NOT NULL
);

ALTER TABLE "products" ADD COLUMN IF NOT EXISTS "brand" serial NOT NULL;

ALTER TABLE "products" ADD FOREIGN KEY ("product_category") REFERENCES "product_categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("type") REFERENCES "product_type" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("brand") REFERENCES "product_brand" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_brand" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_type" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_ingredient" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "product_ingredient" ADD FOREIGN KEY ("ingredient") REFERENCES "ingredient" ("id");

ALTER TABLE "products_bank" ADD FOREIGN KEY ("phanLoai") REFERENCES "classify" ("code");

ALTER TABLE "products_bank" ADD FOREIGN KEY ("dangBaoche") REFERENCES "preparation_type" ("code");

ALTER TABLE "products_bank" ADD FOREIGN KEY ("tieuChuanSx") REFERENCES "production_standard" ("code");

ALTER TABLE "products_bank" ADD FOREIGN KEY ("congTySx") REFERENCES "company_pharma" ("id");

ALTER TABLE "products_bank" ADD FOREIGN KEY ("congTyDk") REFERENCES "company_pharma" ("id");