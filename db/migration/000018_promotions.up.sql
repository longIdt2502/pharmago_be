CREATE TABLE "promotions" (
  "id" uuid PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "type" varchar,
  "title" varchar,
  "conditions_text" varchar,
  "conditions_point_customer" int,
  "min_value" float DEFAULT 0,
  "is_discount_percent" bool NOT NULL DEFAULT false,
  "value_discount" float NOT NULL DEFAULT 0,
  "max_discount" float DEFAULT 0,
  "time_apply" int,
  "date_start" timestamp,
  "date_end" timestamp,
  "apply_multiple_times" bool NOT NULL DEFAULT false,
  "apply_simultaneously" bool NOT NULL DEFAULT false,
  "status" bool NOT NULL,
  "company" serial,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "promotion_item" (
  "id" uuid PRIMARY KEY,
  "min_buy" int DEFAULT 0,
  "amount_gift" int DEFAULT 0,
  "promotions" uuid NOT NULL,
  "applicable_variant" serial,
  "applicable_service" serial,
  "variant" serial,
  "service" serial
);

CREATE TABLE "promotion_type" (
  "code" varchar PRIMARY KEY,
  "title" varchar
);

COMMENT ON COLUMN "promotions"."time_apply" IS 'null: vố số, int: số lần cụ thể';

COMMENT ON COLUMN "promotions"."status" IS 'true: áp dụng, false: vô hiệu';

COMMENT ON COLUMN "promotion_type"."title" IS '1: khuyến mãi giảm giá, 2: khuyến mãi tặng sp';

ALTER TABLE "promotions" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "promotions" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "promotions" ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE "promotions" SET "user_updated" = NULL;
ALTER TABLE "promotions" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "promotions" ADD FOREIGN KEY ("type") REFERENCES "promotion_type" ("code") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ADD FOREIGN KEY ("promotions") REFERENCES "promotions" ("id") ON DELETE CASCADE;

ALTER TABLE "promotion_item" ALTER COLUMN "variant" DROP NOT NULL;
UPDATE "promotion_item" SET "variant" = NULL;
ALTER TABLE "promotion_item" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ALTER COLUMN "service" DROP NOT NULL;
UPDATE "promotion_item" SET "service" = NULL;
ALTER TABLE "promotion_item" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ALTER COLUMN "applicable_variant" DROP NOT NULL;
UPDATE "promotion_item" SET "applicable_variant" = NULL;
ALTER TABLE "promotion_item" ADD FOREIGN KEY ("applicable_variant") REFERENCES "variants" ("id") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ALTER COLUMN "applicable_service" DROP NOT NULL;
UPDATE "promotion_item" SET "applicable_service" = NULL;
ALTER TABLE "promotion_item" ADD FOREIGN KEY ("applicable_service") REFERENCES "services" ("id") ON DELETE SET NULL;

INSERT INTO "promotion_type" (code, title) VALUES ('DISCOUNT', 'Khuyến mãi giảm giá');
INSERT INTO "promotion_type" (code, title) VALUES ('GIFT', 'Khuyến mãi sản phẩm');