ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "customers_contact_address_fkey";

ALTER TABLE "customers" DROP COLUMN IF EXISTS "title";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "license_date";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "contact_name";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "contact_title";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "contact_phone";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "contact_email";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "contact_address";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "account_number";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "bank_name";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "bank_branch";
ALTER TABLE "customers" DROP COLUMN IF EXISTS "issued_by";