-- Drop foreign key related to prescription in medical_bills
ALTER TABLE "medical_bills" DROP CONSTRAINT IF EXISTS "medical_bills_prescription_fkey";
ALTER TABLE "medical_bills" DROP COLUMN IF EXISTS "prescription";

-- Drop foreign key related to variant in prescription_item
ALTER TABLE "prescription_item" DROP CONSTRAINT IF EXISTS "prescription_item_variant_fkey";

-- Drop foreign key related to prescription_uuid in prescription_item
ALTER TABLE "prescription_item" DROP CONSTRAINT IF EXISTS "prescription_item_prescription_uuid_fkey";

-- Drop table prescription_item
DROP TABLE IF EXISTS "prescription_item";

-- Drop foreign key related to company in prescriptions
ALTER TABLE "prescriptions" DROP CONSTRAINT IF EXISTS "prescriptions_company_fkey";

-- Drop foreign key related to customer in prescriptions
ALTER TABLE "prescriptions" DROP CONSTRAINT IF EXISTS "prescriptions_customer_fkey";

-- Drop foreign key related to user_updated in prescriptions
ALTER TABLE "prescriptions" DROP CONSTRAINT IF EXISTS "prescriptions_user_updated_fkey";

-- Drop foreign key related to user_created in prescriptions
ALTER TABLE "prescriptions" DROP CONSTRAINT IF EXISTS "prescriptions_user_created_fkey";

-- Drop foreign key related to doctor in prescriptions
ALTER TABLE "prescriptions" DROP CONSTRAINT IF EXISTS "prescriptions_doctor_fkey";

-- Drop table prescriptions
DROP TABLE IF EXISTS "prescriptions";