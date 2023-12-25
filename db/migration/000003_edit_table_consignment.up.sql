ALTER TABLE "consignment" ADD COLUMN IF NOT EXISTS "variant" serial;

ALTER TABLE "consignment" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE CASCADE;

ALTER TABLE "consignment" DROP CONSTRAINT IF EXISTS "consignment_product_fkey";

ALTER TABLE "consignment" DROP COLUMN IF EXISTS "product";