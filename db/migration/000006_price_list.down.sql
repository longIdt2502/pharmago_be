-- Drop foreign keys in price_list_log
ALTER TABLE price_list_log DROP CONSTRAINT price_list_log_user_updated_fkey;
ALTER TABLE price_list_log DROP CONSTRAINT price_list_log_price_list_fkey;

-- Drop foreign keys in price_list
ALTER TABLE price_list DROP CONSTRAINT price_list_variant_code_fkey;
ALTER TABLE price_list DROP CONSTRAINT price_list_unit_fkey;
ALTER TABLE price_list DROP CONSTRAINT price_list_user_created_fkey;
ALTER TABLE price_list DROP CONSTRAINT price_list_user_updated_fkey;

-- Drop tables
DROP TABLE IF EXISTS price_list_log;
DROP TABLE IF EXISTS price_list;
