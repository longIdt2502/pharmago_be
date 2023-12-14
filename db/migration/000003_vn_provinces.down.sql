-- Drop foreign keys
ALTER TABLE "provinces" DROP CONSTRAINT IF EXISTS "provinces_administrative_unit_fkey";
ALTER TABLE "provinces" DROP CONSTRAINT IF EXISTS "provinces_administrative_region_fkey";
ALTER TABLE "districts" DROP CONSTRAINT IF EXISTS "districts_province_fkey";
ALTER TABLE "districts" DROP CONSTRAINT IF EXISTS "districts_administrative_unit_fkey";
ALTER TABLE "wards" DROP CONSTRAINT IF EXISTS "wards_district_fkey";
ALTER TABLE "wards" DROP CONSTRAINT IF EXISTS "wards_administrative_unit_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_province_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_district_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_ward_fkey";

-- Drop indexes
DROP INDEX IF EXISTS "provinces_administrative_unit_id_idx";
DROP INDEX IF EXISTS "provinces_administrative_region_id_idx";
DROP INDEX IF EXISTS "districts_province_code_idx";
DROP INDEX IF EXISTS "districts_administrative_unit_id_idx";
DROP INDEX IF EXISTS "wards_district_code_idx";
DROP INDEX IF EXISTS "wards_administrative_unit_id_idx";
DROP INDEX IF EXISTS "address_province_idx";
DROP INDEX IF EXISTS "address_district_idx";
DROP INDEX IF EXISTS "address_ward_idx";

-- Remove foreign keys
ALTER TABLE "provinces" DROP COLUMN IF EXISTS "administrative_unit_id";
ALTER TABLE "provinces" DROP COLUMN IF EXISTS "administrative_region_id";
ALTER TABLE "districts" DROP COLUMN IF EXISTS "province_code";
ALTER TABLE "districts" DROP COLUMN IF EXISTS "administrative_unit_id";
ALTER TABLE "wards" DROP COLUMN IF EXISTS "district_code";
ALTER TABLE "wards" DROP COLUMN IF EXISTS "administrative_unit_id";
ALTER TABLE "address" DROP COLUMN IF EXISTS "province";
ALTER TABLE "address" DROP COLUMN IF EXISTS "district";
ALTER TABLE "address" DROP COLUMN IF EXISTS "ward";

-- Drop tables
DROP TABLE IF EXISTS "provinces";
DROP TABLE IF EXISTS "districts";
DROP TABLE IF EXISTS "wards";
DROP TABLE IF EXISTS "administrative_units";
DROP TABLE IF EXISTS "administrative_regions";
