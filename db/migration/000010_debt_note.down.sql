DROP INDEX IF EXISTS "debt_repayment_debt_money_idx";

ALTER TABLE "debt_note" DROP CONSTRAINT IF EXISTS "debt_note_user_created_fkey";
ALTER TABLE "debt_note" DROP CONSTRAINT IF EXISTS "debt_note_status_fkey";
ALTER TABLE "debt_note" DROP CONSTRAINT IF EXISTS "debt_note_type_fkey";
ALTER TABLE "debt_note" DROP CONSTRAINT IF EXISTS "debt_note_company_fkey";
ALTER TABLE "debt_repayment" DROP CONSTRAINT IF EXISTS "debt_repayment_debt_fkey";
ALTER TABLE "debt_repayment" DROP CONSTRAINT IF EXISTS "debt_repayment_user_created_fkey";

DROP TABLE IF EXISTS "debt_note_type";
DROP TABLE IF EXISTS "debt_note_status";
DROP TABLE IF EXISTS "debt_note";
DROP TABLE IF EXISTS "debt_repayment";