-- Drop foreign keys
ALTER TABLE "medical_record_link" DROP CONSTRAINT "medical_record_link_customer_fkey";
ALTER TABLE "medical_record_link" DROP CONSTRAINT "medical_record_link_appointment_schedule_fkey";
ALTER TABLE "medical_record_link" DROP CONSTRAINT "medical_record_link_user_created_fkey";

-- If needed, set columns back to NOT NULL
ALTER TABLE "medical_record_link" ALTER COLUMN "appointment_schedule" SET NOT NULL;
ALTER TABLE "medical_record_link" ALTER COLUMN "user_created" SET NOT NULL;

-- Drop table
DROP TABLE IF EXISTS "medical_record_link";

-- Drop type
DROP TYPE IF EXISTS "medical_record_link_type";
