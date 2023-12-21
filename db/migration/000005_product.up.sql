ALTER TABLE products ADD COLUMN taDuoc varchar(255);
ALTER TABLE products ADD COLUMN nongDo varchar(255);
ALTER TABLE products ADD COLUMN lieuDung varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN chiDinh varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN chongChiDinh varchar(255);
ALTER TABLE products ADD COLUMN congDung varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN tacDungPhu varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN thanTrong varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN tuongTac varchar(255);
ALTER TABLE products ADD COLUMN baoQuan varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN dongGoi varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN noiSx varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN congTySx varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN congTyDk varchar(255) NOT NULL DEFAULT '';
ALTER TABLE products ADD COLUMN active boolean NOT NULL DEFAULT true;

-- ALTER TABLE variants ALTER COLUMN decision_number TYPE varchar(255) NOT NULL UNIQUE;
ALTER TABLE variants ALTER COLUMN decision_number SET DATA TYPE varchar(255),
ALTER COLUMN decision_number SET NOT NULL,
ADD CONSTRAINT decision_number_unique UNIQUE (decision_number);

-- ALTER TABLE variants ALTER COLUMN register_number TYPE varchar(255) NOT NULL UNIQUE;
ALTER TABLE variants ALTER COLUMN register_number SET DATA TYPE varchar(255),
ALTER COLUMN register_number SET NOT NULL,
ADD CONSTRAINT register_number_unique UNIQUE (register_number);

ALTER TABLE variants ADD COLUMN longevity varchar(255) NOT NULL DEFAULT '';
ALTER TABLE variants DROP COLUMN IF EXISTS "discount";
-- ALTER TABLE variants ALTER COLUMN code TYPE varchar(255) NOT NULL UNIQUE;
ALTER TABLE variants ALTER COLUMN code SET DATA TYPE varchar(255),
ALTER COLUMN code SET NOT NULL,
ADD CONSTRAINT code_unique UNIQUE (code);

-- Xoá ràng buộc trên bảng "products"
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_product_category_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_type_fkey";
-- Xoá ràng buộc trên bảng "product_categories"
ALTER TABLE IF EXISTS "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_updated_fkey";
ALTER TABLE IF EXISTS "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_created_fkey";
DROP TABLE IF EXISTS product_categories;
DROP TABLE IF EXISTS product_type;

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

ALTER TABLE "products" ADD FOREIGN KEY ("product_category") REFERENCES "product_categories" ("id");
ALTER TABLE products ALTER COLUMN product_category DROP NOT NULL;

ALTER TABLE "products" ADD FOREIGN KEY ("type") REFERENCES "product_type" ("id");
ALTER TABLE products ALTER COLUMN type DROP NOT NULL;

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");



