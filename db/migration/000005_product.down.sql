-- Rollback changes to the "products" table
ALTER TABLE products DROP COLUMN IF EXISTS taDuoc;
ALTER TABLE products DROP COLUMN IF EXISTS nongDo;
ALTER TABLE products DROP COLUMN IF EXISTS lieuDung;
ALTER TABLE products DROP COLUMN IF EXISTS chiDinh;
ALTER TABLE products DROP COLUMN IF EXISTS chongChiDinh;
ALTER TABLE products DROP COLUMN IF EXISTS congDung;
ALTER TABLE products DROP COLUMN IF EXISTS tacDungPhu;
ALTER TABLE products DROP COLUMN IF EXISTS thanTrong;
ALTER TABLE products DROP COLUMN IF EXISTS tuongTac;
ALTER TABLE products DROP COLUMN IF EXISTS baoQuan;
ALTER TABLE products DROP COLUMN IF EXISTS dongGoi;
ALTER TABLE products DROP COLUMN IF EXISTS noiSx;
ALTER TABLE products DROP COLUMN IF EXISTS congTySx;
ALTER TABLE products DROP COLUMN IF EXISTS congTyDk;

-- Rollback changes to the "variants" table
ALTER TABLE variants DROP CONSTRAINT IF EXISTS decision_number_unique;
ALTER TABLE variants ALTER COLUMN decision_number DROP NOT NULL;
ALTER TABLE variants DROP COLUMN IF EXISTS "decision_number";
ALTER TABLE variants ADD COLUMN decision_number bigint ;

ALTER TABLE variants DROP CONSTRAINT IF EXISTS register_number_unique;
ALTER TABLE variants ALTER COLUMN register_number DROP NOT NULL;
ALTER TABLE variants DROP COLUMN IF EXISTS "register_number";
ALTER TABLE variants ADD COLUMN register_number bigint ;

ALTER TABLE variants DROP COLUMN IF EXISTS longevity;
ALTER TABLE variants ADD COLUMN IF NOT EXISTS "discount" numeric;
ALTER TABLE variants DROP CONSTRAINT IF EXISTS code_unique;
ALTER TABLE variants DROP COLUMN IF EXISTS code;
ALTER TABLE variants ADD COLUMN code numeric;
