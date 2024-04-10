DROP INDEX IF EXISTS "service_variant_service_variant_idx";

ALTER TABLE "companies" DROP CONSTRAINT IF EXISTS "companies_type_fkey";
ALTER TABLE "services" DROP CONSTRAINT IF EXISTS "services_staff_fkey";
ALTER TABLE "services" DROP CONSTRAINT IF EXISTS "services_image_fkey";
ALTER TABLE "services" DROP CONSTRAINT IF EXISTS "services_company_fkey";
ALTER TABLE "services" DROP CONSTRAINT IF EXISTS "services_user_created_fkey";
ALTER TABLE "services" DROP CONSTRAINT IF EXISTS "services_user_updated_fkey";
ALTER TABLE "service_variant" DROP CONSTRAINT IF EXISTS "service_variant_service_fkey";
ALTER TABLE "service_variant" DROP CONSTRAINT IF EXISTS "service_variant_variant_fkey";

ALTER TABLE "companies" DROP COLUMN IF EXISTS "type";

DROP TABLE IF EXISTS "services";
DROP TABLE IF EXISTS "service_variant";
DROP TABLE IF EXISTS "company_type";