-- Drop các INDEX đã tạo
DROP INDEX IF EXISTS "address_province_idx";
DROP INDEX IF EXISTS "address_district_idx";
DROP INDEX IF EXISTS "address_ward_idx";
DROP INDEX IF EXISTS "provinces_administrative_unit_id_idx";
DROP INDEX IF EXISTS "provinces_administrative_region_id_idx";
DROP INDEX IF EXISTS "districts_province_code_idx";
DROP INDEX IF EXISTS "districts_administrative_unit_id_idx";
DROP INDEX IF EXISTS "wards_district_code_idx";
DROP INDEX IF EXISTS "wards_administrative_unit_id_idx";
DROP INDEX IF EXISTS "products_unit_idx";
DROP INDEX IF EXISTS "products_id_unit_idx";
DROP INDEX IF EXISTS "company_pharma_company_pharma_type_idx";
DROP INDEX IF EXISTS "company_pharma_name_company_pharma_type_idx";
DROP INDEX IF EXISTS "product_media_product_idx";
DROP INDEX IF EXISTS "product_media_media_idx";
DROP INDEX IF EXISTS "product_media_product_media_idx";
DROP INDEX IF EXISTS "variant_media_variant_idx";
DROP INDEX IF EXISTS "variant_media_media_idx";
DROP INDEX IF EXISTS "variant_media_variant_media_idx";
DROP INDEX IF EXISTS "orders_qr_idx";
DROP INDEX IF EXISTS "orders_id_qr_idx";
DROP INDEX IF EXISTS "customers_address_idx";
DROP INDEX IF EXISTS "customers_id_address_idx";
DROP INDEX IF EXISTS "tickets_qr_idx";
DROP INDEX IF EXISTS "tickets_id_qr_idx";

-- Drop các FOREIGN KEY CONSTRAINT đã thêm
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_type_fkey";
ALTER TABLE "account_media" DROP CONSTRAINT IF EXISTS "account_media_account_fkey";
ALTER TABLE "account_media" DROP CONSTRAINT IF EXISTS "account_media_media_fkey";
ALTER TABLE "account_company" DROP CONSTRAINT IF EXISTS "account_company_account_fkey";
ALTER TABLE "account_company" DROP CONSTRAINT IF EXISTS "account_company_company_fkey";
ALTER TABLE "sessions" DROP CONSTRAINT IF EXISTS "sessions_username_fkey";
ALTER TABLE "verifies" DROP CONSTRAINT IF EXISTS "verifies_username_fkey";
ALTER TABLE "companies" DROP CONSTRAINT IF EXISTS "companies_owner_fkey";
ALTER TABLE "companies" DROP CONSTRAINT IF EXISTS "companies_address_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_province_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_district_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_ward_fkey";
ALTER TABLE "address" DROP CONSTRAINT IF EXISTS "address_user_created_fkey";
ALTER TABLE "provinces" DROP CONSTRAINT IF EXISTS "provinces_administrative_unit_id_fkey";
ALTER TABLE "provinces" DROP CONSTRAINT IF EXISTS "provinces_administrative_region_id_fkey";
ALTER TABLE "districts" DROP CONSTRAINT IF EXISTS "districts_province_code_fkey";
ALTER TABLE "districts" DROP CONSTRAINT IF EXISTS "districts_administrative_unit_id_fkey";
ALTER TABLE "wards" DROP CONSTRAINT IF EXISTS "wards_district_code_fkey";
ALTER TABLE "wards" DROP CONSTRAINT IF EXISTS "wards_administrative_unit_id_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_product_category_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_type_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_brand_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_unit_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_phan_loai_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_dang_bao_che_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_tieu_chuan_sx_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_cong_ty_sx_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_cong_ty_dk_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_company_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_user_created_fkey";
ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_user_updated_fkey";
ALTER TABLE "products_bank" DROP CONSTRAINT IF EXISTS "products_bank_phan_loai_fkey";
ALTER TABLE "products_bank" DROP CONSTRAINT IF EXISTS "products_bank_dang_bao_che_fkey";
ALTER TABLE "products_bank" DROP CONSTRAINT IF EXISTS "products_bank_tieu_chuan_sx_fkey";
ALTER TABLE "products_bank" DROP CONSTRAINT IF EXISTS "products_bank_cong_ty_sx_fkey";
ALTER TABLE "products_bank" DROP CONSTRAINT IF EXISTS "products_bank_cong_ty_dk_fkey";
ALTER TABLE "price_list" DROP CONSTRAINT IF EXISTS "price_list_variant_code_fkey";
ALTER TABLE "price_list" DROP CONSTRAINT IF EXISTS "price_list_unit_fkey";
ALTER TABLE "price_list" DROP CONSTRAINT IF EXISTS "price_list_user_created_fkey";
ALTER TABLE "price_list" DROP CONSTRAINT IF EXISTS "price_list_user_updated_fkey";

ALTER TABLE "suplier" DROP CONSTRAINT IF EXISTS "suplier_warehouses_fkey";
ALTER TABLE "suplier" DROP CONSTRAINT IF EXISTS "suplier_address_fkey";
ALTER TABLE "suplier" DROP CONSTRAINT IF EXISTS "suplier_company_fkey";

ALTER TABLE "consignment_log" DROP CONSTRAINT IF EXISTS "consignment_log_consignment_fkey";
ALTER TABLE "consignment_log" DROP CONSTRAINT IF EXISTS "consignment_log_user_created_fkey";

ALTER TABLE "consignment" DROP CONSTRAINT IF EXISTS "consignment_ticket_fkey";
ALTER TABLE "consignment" DROP CONSTRAINT IF EXISTS "consignment_product_fkey";
ALTER TABLE "consignment" DROP CONSTRAINT IF EXISTS "consignment_user_created_fkey";
ALTER TABLE "consignment" DROP CONSTRAINT IF EXISTS "consignment_user_updated_fkey";

ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_type_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_status_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_qr_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_export_to_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_import_from_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_warehouse_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_user_created_fkey";
ALTER TABLE "tickets" DROP CONSTRAINT IF EXISTS "tickets_user_updated_fkey";

ALTER TABLE "variants" DROP CONSTRAINT IF EXISTS "variants_product_fkey";
ALTER TABLE "variants" DROP CONSTRAINT IF EXISTS "variants_user_created_fkey";
ALTER TABLE "variants" DROP CONSTRAINT IF EXISTS "variants_user_updated_fkey";

ALTER TABLE "variant_media" DROP CONSTRAINT IF EXISTS "variant_media_variant_fkey";
ALTER TABLE "variant_media" DROP CONSTRAINT IF EXISTS "variant_media_media_fkey";

ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_customer_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_status_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_type_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_ticket_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_qr_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_company_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_user_created_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_user_updated_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_payment_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_address_fkey";

ALTER TABLE "order_items" DROP CONSTRAINT IF EXISTS "order_items_order_fkey";
ALTER TABLE "order_items" DROP CONSTRAINT IF EXISTS "order_items_variant_fkey";
ALTER TABLE "order_items" DROP CONSTRAINT IF EXISTS "order_items_consignment_fkey";
ALTER TABLE "order_items" DROP CONSTRAINT IF EXISTS "order_items_consignment_log_fkey";

ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "customers_company_fkey";
ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "customers_address_fkey";
ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "customers_user_created_fkey";
ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "customers_user_updated_fkey";

ALTER TABLE "units" DROP CONSTRAINT IF EXISTS "units_user_created_fkey";
ALTER TABLE "units" DROP CONSTRAINT IF EXISTS "units_user_updated_fkey";

ALTER TABLE "unit_changes" DROP CONSTRAINT IF EXISTS "unit_changes_unit_fkey";
ALTER TABLE "unit_changes" DROP CONSTRAINT IF EXISTS "unit_changes_user_created_fkey";
ALTER TABLE "unit_changes" DROP CONSTRAINT IF EXISTS "unit_changes_user_updated_fkey";

ALTER TABLE "ingredient" DROP CONSTRAINT IF EXISTS "ingredient_product_fkey";

ALTER TABLE "product_type" DROP CONSTRAINT IF EXISTS "product_type_user_created_fkey";
ALTER TABLE "product_type" DROP CONSTRAINT IF EXISTS "product_type_company_fkey";

ALTER TABLE "product_brand" DROP CONSTRAINT IF EXISTS "product_brand_user_created_fkey";
ALTER TABLE "product_brand" DROP CONSTRAINT IF EXISTS "product_brand_company_fkey";

ALTER TABLE "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_user_created_fkey";
ALTER TABLE "product_categories" DROP CONSTRAINT IF EXISTS "product_categories_company_fkey";

ALTER TABLE "product_media" DROP CONSTRAINT IF EXISTS "product_media_product_fkey";
ALTER TABLE "product_media" DROP CONSTRAINT IF EXISTS "product_media_media_fkey";

ALTER TABLE "company_pharma" DROP CONSTRAINT IF EXISTS "company_pharma_company_pharma_type_fkey";

ALTER TABLE "price_list_log" DROP CONSTRAINT IF EXISTS "price_list_log_price_list_fkey";
ALTER TABLE "price_list_log" DROP CONSTRAINT IF EXISTS "price_list_log_user_updated_fkey";

ALTER TABLE "price_list" DROP CONSTRAINT IF EXISTS "price_list_user_created_fkey";
ALTER TABLE "price_list" DROP CONSTRAINT IF EXISTS "price_list_user_updated_fkey";

-- Drop các bảng
DROP TABLE IF EXISTS "suplier" CASCADE;
DROP TABLE IF EXISTS "consignment_log" CASCADE;
DROP TABLE IF EXISTS "consignment" CASCADE;
DROP TABLE IF EXISTS "ticket_status" CASCADE;
DROP TABLE IF EXISTS "ticket_type" CASCADE;
DROP TABLE IF EXISTS "tickets" CASCADE;
DROP TABLE IF EXISTS "warehouses" CASCADE;
DROP TABLE IF EXISTS "medias" CASCADE;
DROP TABLE IF EXISTS "customers" CASCADE;
DROP TABLE IF EXISTS "order_items" CASCADE;
DROP TABLE IF EXISTS "order_status" CASCADE;
DROP TABLE IF EXISTS "order_type" CASCADE;
DROP TABLE IF EXISTS "orders" CASCADE;
DROP TABLE IF EXISTS "variant_media" CASCADE;
DROP TABLE IF EXISTS "variants" CASCADE;
DROP TABLE IF EXISTS "unit_changes" CASCADE;
DROP TABLE IF EXISTS "units" CASCADE;
DROP TABLE IF EXISTS "ingredient" CASCADE;
DROP TABLE IF EXISTS "production_standard" CASCADE;
DROP TABLE IF EXISTS "preparation_type" CASCADE;
DROP TABLE IF EXISTS "classify" CASCADE;
DROP TABLE IF EXISTS "product_type" CASCADE;
DROP TABLE IF EXISTS "product_brand" CASCADE;
DROP TABLE IF EXISTS "product_categories" CASCADE;
DROP TABLE IF EXISTS "product_media" CASCADE;
DROP TABLE IF EXISTS "company_pharma_type" CASCADE;
DROP TABLE IF EXISTS "company_pharma" CASCADE;
DROP TABLE IF EXISTS "price_list_log" CASCADE;
DROP TABLE IF EXISTS "price_list" CASCADE;
DROP TABLE IF EXISTS "products_bank" CASCADE;
DROP TABLE IF EXISTS "products" CASCADE;
DROP TABLE IF EXISTS "wards" CASCADE;
DROP TABLE IF EXISTS "districts" CASCADE;
DROP TABLE IF EXISTS "provinces" CASCADE;
DROP TABLE IF EXISTS "administrative_units" CASCADE;
DROP TABLE IF EXISTS "administrative_regions" CASCADE;
DROP TABLE IF EXISTS "address" CASCADE;
DROP TABLE IF EXISTS "companies" CASCADE;
DROP TABLE IF EXISTS "verifies" CASCADE;
DROP TABLE IF EXISTS "sessions" CASCADE;
DROP TABLE IF EXISTS "account_type" CASCADE;
DROP TABLE IF EXISTS "account_company" CASCADE;
DROP TABLE IF EXISTS "account_media" CASCADE;
DROP TABLE IF EXISTS "accounts" CASCADE;
