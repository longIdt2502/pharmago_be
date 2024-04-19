ALTER TABLE "medical_records" DROP CONSTRAINT "medical_records_customer_fkey";
ALTER TABLE "medical_records" DROP CONSTRAINT "medical_records_doctor_fkey";
ALTER TABLE "medical_records" DROP CONSTRAINT "medical_records_user_created_fkey";
ALTER TABLE "medical_records" DROP CONSTRAINT "medical_records_user_updated_fkey";

ALTER TABLE "medical_record_variant" DROP CONSTRAINT "medical_record_variant_medical_record_fkey";
ALTER TABLE "medical_record_variant" DROP CONSTRAINT "medical_record_variant_variant_fkey";

DROP TABLE IF EXISTS "medical_record_variant";
DROP TABLE IF EXISTS "medical_records";