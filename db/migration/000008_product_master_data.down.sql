-- Migration Down for "products_bank" table
DROP TABLE IF EXISTS "products_bank";

-- Migration Down for "product_ingredient" table
ALTER TABLE "product_ingredient" DROP CONSTRAINT IF EXISTS "product_ingredient_product_fkey";
ALTER TABLE "product_ingredient" DROP CONSTRAINT IF EXISTS "product_ingredient_ingredient_fkey";

ALTER TABLE "product_brand" DROP CONSTRAINT IF EXISTS "product_brand_user_created_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_brand_fkey";

ALTER TABLE "product_type" DROP CONSTRAINT IF EXISTS "product_type_user_created_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_type_fkey";

ALTER TABLE "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_created_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_product_category_fkey";

DROP TABLE IF EXISTS "product_ingredient";

-- Migration Down for "ingredient" table
DROP TABLE IF EXISTS "ingredient";

-- Migration Down for "production_standard" table
DROP TABLE IF EXISTS "production_standard";

-- Migration Down for "preparation_type" table
DROP TABLE IF EXISTS "preparation_type";

-- Migration Down for "classify" table
DROP TABLE IF EXISTS "classify";

-- Migration Down for "product_brand" table
DROP TABLE IF EXISTS "product_brand";

-- Migration Down for "product_type" table
DROP TABLE IF EXISTS "product_type";

-- Migration Down for "product_categories" table
DROP TABLE IF EXISTS "product_categories";