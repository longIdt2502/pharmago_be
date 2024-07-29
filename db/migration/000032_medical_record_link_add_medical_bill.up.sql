ALTER TABLE "medical_record_link" ADD COLUMN IF NOT EXISTS "medical_bill" uuid;
ALTER TABLE "medical_record_link" ALTER COLUMN "medical_bill" DROP NOT NULL;
ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("medical_bill") REFERENCES "medical_bills" ("uuid") ON DELETE SET NULL;
