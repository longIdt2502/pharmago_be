-- Xoá ràng buộc trên bảng "accounts"
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_type_fkey";
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_media_fkey";
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_user_created_fkey";

-- Xoá ràng buộc trên bảng "warehouses"
ALTER TABLE IF EXISTS "warehouses" DROP CONSTRAINT IF EXISTS "warehouses_address_fkey";
ALTER TABLE IF EXISTS "warehouses" DROP CONSTRAINT IF EXISTS "warehouses_companies_fkey";

-- Xoá ràng buộc trên bảng "products"
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_product_category_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_type_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_unit_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_company_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_user_created_fkey";
ALTER TABLE IF EXISTS "products" DROP CONSTRAINT IF EXISTS "products_user_updated_fkey";

-- Xoá ràng buộc trên bảng "product_media"
ALTER TABLE IF EXISTS "product_media" DROP CONSTRAINT IF EXISTS "product_media_product_fkey";
ALTER TABLE IF EXISTS "product_media" DROP CONSTRAINT IF EXISTS "product_media_media_fkey";

-- Xoá ràng buộc trên bảng "product_categories"
ALTER TABLE IF EXISTS "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_created_fkey";
ALTER TABLE IF EXISTS "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_updated_fkey";

-- Xoá ràng buộc trên bảng "units"
ALTER TABLE IF EXISTS "units" DROP CONSTRAINT IF EXISTS "units_user_created_fkey";
ALTER TABLE IF EXISTS "units" DROP CONSTRAINT IF EXISTS "units_user_updated_fkey";

-- Xoá ràng buộc trên bảng "unit_changes"
ALTER TABLE IF EXISTS "unit_changes" DROP CONSTRAINT IF EXISTS "unit_changes_unit_fkey";
ALTER TABLE IF EXISTS "unit_changes" DROP CONSTRAINT IF EXISTS "unit_changes_user_created_fkey";
ALTER TABLE IF EXISTS "unit_changes" DROP CONSTRAINT IF EXISTS "unit_changes_user_updated_fkey";

-- Xoá ràng buộc trên bảng "variants"
ALTER TABLE IF EXISTS "variants" DROP CONSTRAINT IF EXISTS "variants_product_fkey";
ALTER TABLE IF EXISTS "variants" DROP CONSTRAINT IF EXISTS "variants_user_created_fkey";
ALTER TABLE IF EXISTS "variants" DROP CONSTRAINT IF EXISTS "variants_user_updated_fkey";

-- Xoá ràng buộc trên bảng "orders"
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_customer_fkey";
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_status_fkey";
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_type_fkey";
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_ticket_fkey";
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_qr_fkey";

-- Xoá ràng buộc trên bảng "order_items"
ALTER TABLE IF EXISTS "order_items" DROP CONSTRAINT IF EXISTS "order_items_order_fkey";
ALTER TABLE IF EXISTS "order_items" DROP CONSTRAINT IF EXISTS "order_items_variant_fkey";

-- Xoá ràng buộc trên bảng "customers"
ALTER TABLE IF EXISTS "customers" DROP CONSTRAINT IF EXISTS "customers_company_fkey";
ALTER TABLE IF EXISTS "customers" DROP CONSTRAINT IF EXISTS "customers_address_fkey";
ALTER TABLE IF EXISTS "customers" DROP CONSTRAINT IF EXISTS "customers_user_created_fkey";
ALTER TABLE IF EXISTS "customers" DROP CONSTRAINT IF EXISTS "customers_user_updated_fkey";

-- Xoá ràng buộc trên bảng "tickets"
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_type_fkey";
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_status_fkey";
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_qr_fkey";
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_export_from_fkey";
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_import_to_fkey";
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_user_created_fkey";
ALTER TABLE IF EXISTS "tickets" DROP CONSTRAINT IF EXISTS "tickets_user_updated_fkey";

-- Xoá ràng buộc trên bảng "medias"
ALTER TABLE IF EXISTS "medias" DROP CONSTRAINT IF EXISTS "medias_accounts_media_fk";

-- Xoá ràng buộc trên bảng "account_media"
ALTER TABLE IF EXISTS "account_media" DROP CONSTRAINT IF EXISTS "account_media_account_fkey";
ALTER TABLE IF EXISTS "account_media" DROP CONSTRAINT IF EXISTS "account_media_media_fkey";

-- Xoá ràng buộc trên bảng "account_company"
ALTER TABLE IF EXISTS "account_company" DROP CONSTRAINT IF EXISTS "account_company_account_fkey";
ALTER TABLE IF EXISTS "account_company" DROP CONSTRAINT IF EXISTS "account_company_company_fkey";

-- Xoá ràng buộc trên bảng "sessions"
ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT IF EXISTS "sessions_username_fkey";

-- Xoá ràng buộc trên bảng "verifies"
ALTER TABLE IF EXISTS "verifies" DROP CONSTRAINT IF EXISTS "verifies_username_fkey";

DROP TABLE IF EXISTS order_type;
DROP TABLE IF EXISTS ticket_type;
DROP TABLE IF EXISTS product_type;
DROP TABLE IF EXISTS product_media;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS warehouses;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS product_categories;
DROP TABLE IF EXISTS units;
DROP TABLE IF EXISTS unit_changes;
DROP TABLE IF EXISTS variants;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_status;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS ticket_status;
DROP TABLE IF EXISTS address;
DROP TABLE IF EXISTS account_type;
DROP TABLE IF EXISTS account_media;
DROP TABLE IF EXISTS medias;
DROP TABLE IF EXISTS account_company;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS verifies;