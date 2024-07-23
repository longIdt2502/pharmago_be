-- Drop foreign keys related to user_updated
ALTER TABLE companies DROP CONSTRAINT IF EXISTS "companies_user_updated_fkey";
ALTER TABLE companies DROP COLUMN IF EXISTS "user_updated";

-- Drop foreign keys related to user_created
ALTER TABLE companies DROP CONSTRAINT IF EXISTS "companies_user_created_fkey";
ALTER TABLE companies DROP COLUMN IF EXISTS "user_created";

-- Drop foreign keys related to manager
ALTER TABLE companies DROP CONSTRAINT IF EXISTS "companies_manager_fkey";
ALTER TABLE companies DROP COLUMN IF EXISTS "manager";

-- Drop foreign keys related to parent
ALTER TABLE companies DROP CONSTRAINT IF EXISTS "companies_parent_fkey";
ALTER TABLE companies DROP COLUMN IF EXISTS "parent";

-- Drop column is_active
ALTER TABLE companies DROP COLUMN IF EXISTS "is_active";

-- Drop column updated_at
ALTER TABLE companies DROP COLUMN IF EXISTS "updated_at";
