-- ALTER TABLE "products" ALTER COLUMN "name" DROP NOT NULL;
-- ALTER TABLE "products" ALTER COLUMN "code" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "brand" DROP NOT NULL;
-- ALTER TABLE "products" ALTER COLUMN "unit" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "lieu_dung" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "chi_dinh" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "cong_dung" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "tac_dung_phu" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "than_trong" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "bao_quan" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "dong_goi" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "dang_bao_che" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "tieu_chuan_sx" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "cong_ty_sx" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "cong_ty_dk" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "cong_ty_dk" DROP NOT NULL;
-- ALTER TABLE "products" ALTER COLUMN "company" DROP NOT NULL;
ALTER TABLE "products" ALTER COLUMN "user_updated" DROP NOT NULL;
ALTER TABLE "products" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "variants" ALTER COLUMN "barcode" DROP NOT NULL;
ALTER TABLE "variants" ALTER COLUMN "decision_number" DROP NOT NULL;
ALTER TABLE "variants" ALTER COLUMN "register_number" DROP NOT NULL;
ALTER TABLE "variants" ALTER COLUMN "longevity" DROP NOT NULL;
ALTER TABLE "variants" ALTER COLUMN "vat" DROP NOT NULL;
ALTER TABLE "variants" ALTER COLUMN "user_updated" DROP NOT NULL;
ALTER TABLE "variants" ADD COLUMN IF NOT EXISTS "initial_inventory" INT NOT NULL DEFAULT 0;
ALTER TABLE "variants" ADD COLUMN IF NOT EXISTS "real_inventory" INT NOT NULL DEFAULT 0;

ALTER TABLE "order_items" ALTER COLUMN "consignment" DROP NOT NULL;
ALTER TABLE "order_items" ALTER COLUMN "consignment_log" DROP NOT NULL;
