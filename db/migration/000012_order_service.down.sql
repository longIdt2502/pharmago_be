ALTER TABLE "service_order_item" DROP CONSTRAINT IF EXISTS "service_order_item_order_fkey";
ALTER TABLE "service_order_item" DROP CONSTRAINT IF EXISTS "service_order_item_service_fkey";

DROP TABLE IF EXISTS "service_order_item";