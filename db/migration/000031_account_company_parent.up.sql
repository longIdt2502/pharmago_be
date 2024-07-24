ALTER TABLE account_company ADD COLUMN IF NOT EXISTS company_parent serial;
ALTER TABLE account_company ALTER COLUMN company_parent DROP NOT NULL;
UPDATE account_company SET company_parent = NULL;
ALTER TABLE account_company ADD FOREIGN KEY ("company_parent") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE account_company ALTER COLUMN company DROP NOT NULL;