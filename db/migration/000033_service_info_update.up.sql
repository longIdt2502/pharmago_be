ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "brand" serial;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "action_time" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "chi_dinh" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "chong_chi_dinh" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "cong_dung" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "caution" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "hinh_thuc" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "tac_dung_phu" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "number_register" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "number_decision" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "cong_ty_dk" varchar;
ALTER TABLE "services" ADD COLUMN IF NOT EXISTS "message" varchar;

ALTER TABLE "services" ALTER COLUMN "brand" DROP NOT NULL;
UPDATE "services" SET "brand" = NULL;
ALTER TABLE "services" ADD FOREIGN KEY ("brand") REFERENCES "product_brand" ("id") ON DELETE SET NULL;