-- DROP foreign key constraint
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_address_fkey";

-- DROP columns added in the migration
ALTER TABLE "accounts" DROP COLUMN IF EXISTS "address";
ALTER TABLE "accounts" DROP COLUMN IF EXISTS "dob";
ALTER TABLE "accounts" DROP COLUMN IF EXISTS "licence";
ALTER TABLE "accounts" DROP COLUMN IF EXISTS "gender";

-- DROP ENUM type
DROP TYPE IF EXISTS "gender";
