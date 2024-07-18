-- Xóa các khóa ngoại trên bảng appointment_schedule_service
ALTER TABLE "appointment_schedule_service" DROP CONSTRAINT "appointment_schedule_service_as_uuid_fkey";
ALTER TABLE "appointment_schedule_service" DROP CONSTRAINT "appointment_schedule_service_service_fkey";
ALTER TABLE "appointment_schedule_service" DROP CONSTRAINT "appointment_schedule_service_order_service_fkey";

-- Xóa các khóa ngoại trên bảng appointment_schedule_url
ALTER TABLE "appointment_schedule_url" DROP CONSTRAINT "appointment_schedule_url_as_uuid_fkey";

-- Xóa các khóa ngoại trên bảng appointment_schedule_drug
ALTER TABLE "appointment_schedule_drug" DROP CONSTRAINT "appointment_schedule_drug_as_uuid_fkey";
ALTER TABLE "appointment_schedule_drug" DROP CONSTRAINT "appointment_schedule_drug_variant_fkey";

-- Xóa các khóa ngoại trên bảng appointment_schedules
ALTER TABLE "appointment_schedules" DROP CONSTRAINT "appointment_schedules_customer_fkey";
ALTER TABLE "appointment_schedules" DROP CONSTRAINT "appointment_schedules_company_fkey";
ALTER TABLE "appointment_schedules" DROP CONSTRAINT "appointment_schedules_doctor_fkey";
ALTER TABLE "appointment_schedules" DROP CONSTRAINT "appointment_schedules_user_created_fkey";
ALTER TABLE "appointment_schedules" DROP CONSTRAINT "appointment_schedules_user_updated_fkey";

-- Xóa bảng appointment_schedule_service
DROP TABLE IF EXISTS "appointment_schedule_service";

-- Xóa bảng appointment_schedule_url
DROP TABLE IF EXISTS "appointment_schedule_url";

-- Xóa bảng appointment_schedule_drug
DROP TABLE IF EXISTS "appointment_schedule_drug";

-- Xóa bảng appointment_schedules
DROP TABLE IF EXISTS "appointment_schedules";
