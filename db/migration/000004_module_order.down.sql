ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_payment_fkey";

ALTER TABLE "payment_items" DROP CONSTRAINT IF EXISTS "payment_items_payment_fkey";

ALTER TABLE "payment_items" DROP CONSTRAINT IF EXISTS "payment_items_type_fkey";

DROP TABLE IF EXISTS "payments";
DROP TABLE IF EXISTS "payment_items";
DROP TABLE IF EXISTS "payment_item_types";