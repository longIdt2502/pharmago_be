ALTER TABLE companies ADD COLUMN IF NOT EXISTS "parent" serial;
ALTER TABLE companies ALTER COLUMN "parent" DROP NOT NULL;
UPDATE companies SET parent = NULL;
ALTER TABLE companies ADD FOREIGN KEY ("parent") REFERENCES companies ("id");

ALTER TABLE companies ADD COLUMN IF NOT EXISTS "is_active" bool NOT NULL DEFAULT FALSE;

ALTER TABLE companies ADD COLUMN IF NOT EXISTS "manager" serial;
ALTER TABLE companies ALTER COLUMN "manager" DROP NOT NULL;
UPDATE companies SET manager = NULL;
ALTER TABLE companies ADD FOREIGN KEY ("manager") REFERENCES accounts ("id");

ALTER TABLE companies ADD COLUMN IF NOT EXISTS "user_created" serial;
ALTER TABLE companies ALTER COLUMN "user_created" DROP NOT NULL;
UPDATE companies SET user_created = NULL;
ALTER TABLE companies ADD FOREIGN KEY ("user_created") REFERENCES accounts ("id");

ALTER TABLE companies ADD COLUMN IF NOT EXISTS "user_updated" serial;
ALTER TABLE companies ALTER COLUMN "user_updated" DROP NOT NULL;
UPDATE companies SET user_updated = NULL;
ALTER TABLE companies ADD FOREIGN KEY ("user_updated") REFERENCES accounts ("id");

ALTER TABLE companies ADD COLUMN IF NOT EXISTS "updated_at" timestamp;