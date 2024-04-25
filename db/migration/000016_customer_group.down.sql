ALTER TABLE "customer_group" DROP CONSTRAINT IF EXISTS "customer_group_company_fkey";
ALTER TABLE "customer_group" DROP CONSTRAINT IF EXISTS "customer_group_user_updated_fkey";
ALTER TABLE "customer_group" DROP CONSTRAINT IF EXISTS "customer_group_user_created_fkey";

ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "customers_group_fkey";

ALTER TABLE "customers" DROP COLUMN IF EXISTS "group";

DROP TABLE IF EXISTS "customer_group";