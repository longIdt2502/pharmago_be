-- Xóa các foreign key trước khi xóa bảng chính
ALTER TABLE "promotions" DROP CONSTRAINT IF EXISTS "promotions_user_created_fkey";
ALTER TABLE "promotions" DROP CONSTRAINT IF EXISTS "promotions_company_fkey";
ALTER TABLE "promotions" DROP CONSTRAINT IF EXISTS "promotions_user_updated_fkey";
ALTER TABLE "promotions" DROP CONSTRAINT IF EXISTS "promotions_type_fkey";
ALTER TABLE "promotion_item" DROP CONSTRAINT IF EXISTS "promotion_item_promotions_fkey";
ALTER TABLE "promotion_item" DROP CONSTRAINT IF EXISTS "promotion_item_applicable_variant_fkey";
ALTER TABLE "promotion_item" DROP CONSTRAINT IF EXISTS "promotion_item_applicable_service_fkey";
ALTER TABLE "promotion_item" DROP CONSTRAINT IF EXISTS "promotion_item_variant_fkey";
ALTER TABLE "promotion_item" DROP CONSTRAINT IF EXISTS "promotion_item_service_fkey";

-- Xóa các comment nếu chúng tồn tại
COMMENT ON COLUMN "promotions"."time_apply" IS NULL;
COMMENT ON COLUMN "promotions"."status" IS NULL;
COMMENT ON COLUMN "promotion_type"."title" IS NULL;

-- Xóa các bảng nếu chúng tồn tại
DROP TABLE IF EXISTS "promotion_item" CASCADE;
DROP TABLE IF EXISTS "promotions" CASCADE;
DROP TABLE IF EXISTS "promotion_type" CASCADE;


