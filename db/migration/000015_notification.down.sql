ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_type_fkey";
ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_company_fkey";

DROP TABLE IF EXISTS "notification";
DROP TABLE IF EXISTS "noti_type";