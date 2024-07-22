-- Drop foreign keys for appointment_schedule_drug
ALTER TABLE "appointment_schedule_drug" DROP CONSTRAINT IF EXISTS "appointment_schedule_drug_mb_uuid_fkey";
ALTER TABLE "appointment_schedule_drug" DROP COLUMN IF EXISTS "mb_uuid";
ALTER TABLE "appointment_schedule_drug" ALTER COLUMN "as_uuid" SET NOT NULL;

-- Drop foreign keys for appointment_schedule_url
ALTER TABLE "appointment_schedule_url" DROP CONSTRAINT IF EXISTS "appointment_schedule_url_mb_uuid_fkey";
ALTER TABLE "appointment_schedule_url" DROP COLUMN IF EXISTS "mb_uuid";
ALTER TABLE "appointment_schedule_url" ALTER COLUMN "as_uuid" SET NOT NULL;

-- Drop foreign keys for appointment_schedule_service
ALTER TABLE "appointment_schedule_service" DROP CONSTRAINT IF EXISTS "appointment_schedule_service_mb_uuid_fkey";
ALTER TABLE "appointment_schedule_service" DROP COLUMN IF EXISTS "mb_uuid";
ALTER TABLE "appointment_schedule_service" ALTER COLUMN "as_uuid" SET NOT NULL;

-- Drop foreign keys from medical_bills
ALTER TABLE "medical_bills" DROP CONSTRAINT IF EXISTS "medical_bills_user_updated_fkey";
ALTER TABLE "medical_bills" DROP CONSTRAINT IF EXISTS "medical_bills_user_created_fkey";
ALTER TABLE "medical_bills" DROP CONSTRAINT IF EXISTS "medical_bills_doctor_fkey";
ALTER TABLE "medical_bills" DROP CONSTRAINT IF EXISTS "medical_bills_company_fkey";
ALTER TABLE "medical_bills" DROP CONSTRAINT IF EXISTS "medical_bills_customer_fkey";

-- Drop table medical_bills
DROP TABLE IF EXISTS "medical_bills";
