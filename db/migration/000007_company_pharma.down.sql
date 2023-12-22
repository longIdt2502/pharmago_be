-- Migration Down for "company_pharma" table
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_congTySx_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_congTyDk_fkey";
ALTER TABLE "company_pharma" DROP CONSTRAINT IF EXISTS "company_pharma_company_pharma_type_fkey";

ALTER TABLE "products" DROP COLUMN IF EXISTS "congTySx";
ALTER TABLE "products" ADD COLUMN "congTySx" varchar;

ALTER TABLE "products" DROP COLUMN IF EXISTS "congTyDk";
ALTER TABLE "products" ADD COLUMN "congTyDk" varchar;

DROP TABLE IF EXISTS "company_pharma";
DROP TABLE IF EXISTS "company_pharma_type";
