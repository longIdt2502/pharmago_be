-- Drop foreign key related to order
ALTER TABLE "medical_bill_order_sell" DROP CONSTRAINT IF EXISTS "medical_bill_order_sell_order_fkey";

-- Drop foreign key related to uuid
ALTER TABLE "medical_bill_order_sell" DROP CONSTRAINT IF EXISTS "medical_bill_order_sell_uuid_fkey";

-- Drop table medical_bill_order_sell
DROP TABLE IF EXISTS "medical_bill_order_sell";
