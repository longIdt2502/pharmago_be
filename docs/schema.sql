-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-08-01T03:19:03.546Z

CREATE TYPE "gender" AS ENUM (
  'nam',
  'nữ',
  'khác'
);

CREATE TYPE "medical_record_link_type" AS ENUM (
  'test',
  'patient',
  'diagnostic'
);

CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "type" serial NOT NULL,
  "is_verify" boolean NOT NULL DEFAULT false,
  "role" serial,
  "gender" gender,
  "licence" varchar,
  "dob" timestamp,
  "address" serial,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_media" (
  "id" serial PRIMARY KEY,
  "account" serial NOT NULL,
  "media" serial NOT NULL
);

CREATE TABLE "account_company" (
  "id" serial PRIMARY KEY,
  "account" serial NOT NULL,
  "company" serial NOT NULL
);

CREATE TABLE "account_type" (
  "id" serial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "roles" (
  "id" serial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL,
  "note" varchar,
  "company" serial,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "role_item" (
  "id" serial PRIMARY KEY,
  "roles" serial NOT NULL,
  "app" varchar NOT NULL,
  "value" bool DEFAULT false
);

CREATE TABLE "apps" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "code" varchar NOT NULL,
  "parent" varchar,
  "level" int DEFAULT 1
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verifies" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT ((now() + interval '15 minutes'))
);

CREATE TABLE "companies" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "tax_code" varchar,
  "phone" varchar,
  "description" varchar,
  "address" serial,
  "oa_id" varchar,
  "time_open" time,
  "time_close" time,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "owner" serial NOT NULL,
  "type" varchar NOT NULL
);

CREATE TABLE "company_type" (
  "code" varchar UNIQUE PRIMARY KEY NOT NULL,
  "title" varchar UNIQUE NOT NULL
);

CREATE TABLE "address" (
  "id" serial PRIMARY KEY,
  "lat" float NOT NULL,
  "lng" float NOT NULL,
  "province" varchar,
  "district" varchar,
  "ward" varchar,
  "title" varchar NOT NULL,
  "user_created" serial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "administrative_regions" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "code_name" varchar NOT NULL,
  "code_name_en" varchar NOT NULL
);

CREATE TABLE "administrative_units" (
  "id" serial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "full_name_en" varchar NOT NULL,
  "short_name" varchar NOT NULL,
  "short_name_en" varchar NOT NULL,
  "code_name" varchar NOT NULL,
  "code_name_en" varchar NOT NULL
);

CREATE TABLE "provinces" (
  "code" varchar(20) UNIQUE PRIMARY KEY,
  "name" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "full_name_en" varchar NOT NULL,
  "code_name" varchar NOT NULL,
  "administrative_unit_id" serial,
  "administrative_region_id" serial
);

CREATE TABLE "districts" (
  "code" varchar UNIQUE PRIMARY KEY,
  "name" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "full_name_en" varchar NOT NULL,
  "code_name" varchar NOT NULL,
  "province_code" varchar,
  "administrative_unit_id" serial
);

CREATE TABLE "wards" (
  "code" varchar UNIQUE PRIMARY KEY,
  "name" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "full_name_en" varchar NOT NULL,
  "code_name" varchar NOT NULL,
  "district_code" varchar,
  "administrative_unit_id" serial
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "product_category" serial,
  "type" serial,
  "brand" serial NOT NULL,
  "unit" serial NOT NULL,
  "ta_duoc" varchar(255),
  "nong_do" varchar(255),
  "lieu_dung" varchar(255) NOT NULL,
  "chi_dinh" varchar(255) NOT NULL,
  "chong_chi_dinh" varchar(255),
  "cong_dung" varchar(255) NOT NULL,
  "tac_dung_phu" varchar(255) NOT NULL,
  "than_trong" varchar(255) NOT NULL,
  "tuong_tac" varchar(255),
  "bao_quan" varchar(255) NOT NULL,
  "dong_goi" varchar(255) NOT NULL,
  "phan_loai" varchar DEFAULT null,
  "dang_bao_che" varchar NOT NULL,
  "tieu_chuan_sx" varchar NOT NULL,
  "cong_ty_sx" serial NOT NULL,
  "cong_ty_dk" serial NOT NULL,
  "active" boolean NOT NULL DEFAULT true,
  "company" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products_bank" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "ta_duoc" varchar(255),
  "nong_do" varchar(255),
  "lieu_dung" varchar(255) NOT NULL,
  "chi_dinh" varchar(255) NOT NULL,
  "chong_chi_dinh" varchar(255),
  "cong_dung" varchar(255) NOT NULL,
  "tac_dung_phu" varchar(255) NOT NULL,
  "than_trong" varchar(255) NOT NULL,
  "tuong_tac" varchar(255),
  "bao_quan" varchar(255) NOT NULL,
  "dong_goi" varchar(255) NOT NULL,
  "phan_loai" varchar DEFAULT null,
  "dang_bao_che" varchar NOT NULL,
  "tieu_chuan_sx" varchar NOT NULL,
  "cong_ty_sx" serial NOT NULL,
  "cong_ty_dk" serial NOT NULL
);

CREATE TABLE "price_list" (
  "id" serial PRIMARY KEY,
  "variant_code" varchar UNIQUE NOT NULL,
  "variant_name" varchar NOT NULL,
  "price_import" float NOT NULL,
  "price_sell" float NOT NULL,
  "unit" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "price_list_log" (
  "id" serial PRIMARY KEY,
  "old_price_import" float NOT NULL,
  "new_price_import" float NOT NULL,
  "old_price_sell" float NOT NULL,
  "new_price_sell" float NOT NULL,
  "price_list" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company_pharma" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "code" varchar(255),
  "country" varchar(255),
  "address" varchar(255),
  "company_pharma_type" varchar DEFAULT null,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company_pharma_type" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "title" varchar(255) NOT NULL
);

CREATE TABLE "product_media" (
  "id" serial PRIMARY KEY,
  "product" serial NOT NULL,
  "media" serial NOT NULL
);

CREATE TABLE "product_categories" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "user_created" serial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "company" serial NOT NULL
);

CREATE TABLE "product_brand" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "user_created" serial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "company" serial NOT NULL
);

CREATE TABLE "product_type" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "user_created" serial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "company" serial NOT NULL
);

CREATE TABLE "classify" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "preparation_type" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "production_standard" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "ingredient" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "weight" float NOT NULL DEFAULT 0,
  "unit" varchar NOT NULL,
  "product" serial NOT NULL
);

CREATE TABLE "units" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "sell_price" float NOT NULL DEFAULT 0,
  "import_price" float NOT NULL DEFAULT 0,
  "weight" float,
  "weight_unit" varchar,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "unit_changes" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "value" bigint NOT NULL DEFAULT 0,
  "sell_price" float NOT NULL DEFAULT 0,
  "unit" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variants" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "barcode" varchar NOT NULL,
  "decision_number" varchar(255) UNIQUE NOT NULL,
  "register_number" varchar(255) UNIQUE NOT NULL,
  "longevity" varchar(255) NOT NULL,
  "vat" float NOT NULL DEFAULT 0,
  "product" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variant_media" (
  "id" serial PRIMARY KEY,
  "variant" serial NOT NULL,
  "media" serial NOT NULL
);

CREATE TABLE "promotions" (
  "id" uuid PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "type" varchar,
  "title" varchar,
  "conditions_text" varchar,
  "conditions_point_customer" int,
  "min_value" float DEFAULT 0,
  "is_discount_percent" bool NOT NULL DEFAULT false,
  "value_discount" float NOT NULL DEFAULT 0,
  "max_discount" float DEFAULT 0,
  "time_apply" int,
  "date_start" timestamp,
  "date_end" timestamp,
  "apply_multiple_times" bool NOT NULL DEFAULT false,
  "apply_simultaneously" bool NOT NULL DEFAULT false,
  "status" bool NOT NULL,
  "company" serial,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "promotion_item" (
  "id" uuid PRIMARY KEY,
  "min_buy" int DEFAULT 0,
  "amount_gift" int DEFAULT 0,
  "promotions" uuid NOT NULL,
  "applicable_variant" serial,
  "applicable_service" serial,
  "variant" serial,
  "service" serial
);

CREATE TABLE "promotion_type" (
  "code" varchar PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "code" varchar(255) UNIQUE NOT NULL,
  "total_price" float NOT NULL DEFAULT 0,
  "description" varchar(255),
  "vat" float NOT NULL DEFAULT 0,
  "discount" varchar(255) NOT NULL DEFAULT '0',
  "service_price" float NOT NULL DEFAULT 0,
  "must_paid" float NOT NULL DEFAULT 0,
  "customer" serial,
  "address" serial,
  "status" varchar,
  "type" varchar,
  "ticket" serial,
  "qr" serial,
  "company" serial NOT NULL,
  "payment" serial NOT NULL,
  "user_created" serial,
  "user_updated" serial,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "order_type" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "order_status" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "order_items" (
  "id" serial PRIMARY KEY,
  "order" serial NOT NULL,
  "variant" serial NOT NULL,
  "value" int NOT NULL DEFAULT 0,
  "total_price" float NOT NULL DEFAULT 0,
  "consignment" serial,
  "consignment_log" serial
);

CREATE TABLE "service_order_item" (
  "id" serial PRIMARY KEY,
  "order" serial NOT NULL,
  "service" serial,
  "unit_price" float NOT NULL DEFAULT 0,
  "discount" float NOT NULL DEFAULT 0,
  "total_price" float NOT NULL DEFAULT 0
);

CREATE TABLE "customers" (
  "id" serial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "company" serial NOT NULL,
  "address" serial,
  "email" varchar,
  "title" varchar,
  "phone" varchar(20),
  "gender" gender,
  "license" varchar(20),
  "license_date" timestamp,
  "birthday" timestamptz,
  "group" serial,
  "contact_name" varchar,
  "contact_title" varchar,
  "contact_phone" varchar,
  "contact_email" varchar,
  "contact_address" serial,
  "account_number" varchar,
  "bank_name" varchar,
  "bank_branch" varchar,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "customer_group" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "company" serial NOT NULL,
  "note" varchar,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "medical_records" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "customer" serial NOT NULL,
  "weight" float,
  "long" float,
  "symptom" varchar NOT NULL,
  "diagnostic" varchar NOT NULL,
  "result" varchar NOT NULL,
  "doctor" serial,
  "re_examination" int NOT NULL DEFAULT 0,
  "note" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  "user_created" serial,
  "user_updated" serial
);

CREATE TABLE "medical_record_variant" (
  "id" serial PRIMARY KEY,
  "medical_record" serial NOT NULL,
  "variant" serial NOT NULL
);

CREATE TABLE "payment_item_types" (
  "id" serial PRIMARY KEY,
  "code" varchar(255) UNIQUE NOT NULL,
  "title" varchar(255) NOT NULL
);

CREATE TABLE "payments" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "must_paid" float NOT NULL DEFAULT 0,
  "had_paid" float NOT NULL DEFAULT 0,
  "need_pay" float NOT NULL DEFAULT 0
);

CREATE TABLE "payment_items" (
  "id" serial PRIMARY KEY,
  "type" varchar NOT NULL,
  "value" float NOT NULL DEFAULT 0,
  "is_paid" bool NOT NULL DEFAULT false,
  "payment" serial NOT NULL,
  "extra_note" varchar
);

CREATE TABLE "medias" (
  "id" serial PRIMARY KEY,
  "media_url" varchar NOT NULL
);

CREATE TABLE "warehouses" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "code" varchar(255) UNIQUE NOT NULL,
  "address" serial,
  "companies" serial
);

CREATE TABLE "tickets" (
  "id" serial PRIMARY KEY,
  "code" varchar NOT NULL,
  "type" serial,
  "status" serial,
  "note" varchar,
  "qr" serial,
  "export_to" serial,
  "import_from" serial,
  "total_price" float NOT NULL DEFAULT 0,
  "warehouse" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ticket_type" (
  "id" serial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "ticket_status" (
  "id" serial PRIMARY KEY,
  "code" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "consignment" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "quantity" int NOT NULL DEFAULT 0,
  "inventory" int NOT NULL DEFAULT 0,
  "ticket" serial,
  "variant" serial,
  "expired_at" timestamptz NOT NULL,
  "producted_at" timestamptz NOT NULL,
  "is_available" bool NOT NULL DEFAULT false,
  "user_created" serial,
  "user_updated" serial,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "consignment_log" (
  "id" serial PRIMARY KEY,
  "consignment" serial NOT NULL,
  "inventory" int NOT NULL DEFAULT 0,
  "amount_change" int NOT NULL DEFAULT 0,
  "user_created" serial,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "suplier" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar(255) NOT NULL,
  "deputy_name" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL,
  "email" varchar(255),
  "address" serial,
  "warehouses" serial,
  "company" serial
);

CREATE TABLE "debt_note_type" (
  "code" varchar UNIQUE PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "debt_note_status" (
  "code" varchar UNIQUE PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "debt_note" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "title" varchar,
  "entity" varchar NOT NULL,
  "money" float NOT NULL DEFAULT 0,
  "paymented" float NOT NULL DEFAULT 0,
  "note" varchar,
  "type" varchar NOT NULL,
  "status" varchar NOT NULL,
  "company" serial NOT NULL,
  "user_created" serial NOT NULL,
  "exprise" timestamp NOT NULL,
  "dabt_note_at" timestamp DEFAULT (now())
);

CREATE TABLE "debt_repayment" (
  "id" serial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "money" float NOT NULL DEFAULT 0,
  "created_at" timestamp DEFAULT (now()),
  "debt" serial NOT NULL,
  "user_created" serial NOT NULL
);

CREATE TABLE "services" (
  "id" serial PRIMARY KEY,
  "active" bool NOT NULL DEFAULT true,
  "image" serial,
  "code" varchar UNIQUE NOT NULL,
  "title" varchar NOT NULL,
  "entity" varchar,
  "brand" serial,
  "staff" serial NOT NULL,
  "frequency" varchar,
  "reminder_time" int,
  "action_time" varchar(255),
  "unit" varchar NOT NULL,
  "price" float NOT NULL DEFAULT 0,
  "description" varchar,
  "chi_dinh" varchar(255),
  "chong_chi_dinh" varchar(255),
  "cong_dung" varchar(255),
  "caution" varchar(255),
  "hinh_thuc" varchar(255),
  "tac_dung_phu" varchar(255),
  "number_register" varchar(255),
  "number_decision" varchar(255),
  "cong_ty_dk" varchar(255),
  "message" varchar(255),
  "company" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "service_variant" (
  "id" serial PRIMARY KEY,
  "service" serial,
  "variant" serial
);

CREATE TABLE "notification" (
  "id" serial PRIMARY KEY,
  "type" varchar NOT NULL,
  "topic" varchar NOT NULL,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "is_read" bool NOT NULL DEFAULT false,
  "data" varchar,
  "company" serial,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "noti_type" (
  "code" varchar UNIQUE NOT NULL
);

CREATE TABLE "appointment_schedules" (
  "id" serial PRIMARY KEY,
  "uuid" uuid UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "customer" serial,
  "company" serial,
  "doctor" serial,
  "symptoms" varchar,
  "diagnostic" varchar,
  "qr_code_url" varchar,
  "is_done" bool NOT NULL,
  "meeting_at" timestamp NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "appointment_schedule_service" (
  "id" serial PRIMARY KEY,
  "as_uuid" uuid,
  "mb_uuid" uuid,
  "service" serial,
  "order_service" serial
);

CREATE TABLE "appointment_schedule_url" (
  "id" serial PRIMARY KEY,
  "as_uuid" uuid,
  "mb_uuid" uuid,
  "url" varchar,
  "name_doc" varchar
);

CREATE TABLE "medical_bills" (
  "id" serial PRIMARY KEY,
  "uuid" uuid UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "customer" serial,
  "company" serial,
  "doctor" serial,
  "symptoms" varchar,
  "diagnostic" varchar,
  "qr_code_url" varchar,
  "is_done" bool NOT NULL,
  "prescription" uuid,
  "meeting_at" timestamp NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "medical_bill_order_sell" (
  "uuid" uuid,
  "order" serial
);

CREATE TABLE "medical_record_link" (
  "id" serial PRIMARY KEY,
  "uuid" uuid NOT NULL,
  "type" medical_record_type NOT NULL,
  "title" varchar,
  "url" varchar NOT NULL,
  "customer" serial,
  "appointment_schedule" uuid,
  "medical_bill" uuid,
  "user_created" serial,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "prescriptions" (
  "id" serial PRIMARY KEY,
  "uuid" uuid UNIQUE NOT NULL,
  "code" varchar UNIQUE NOT NULL,
  "symptoms" varchar,
  "diagnostic" varchar,
  "customer" serial,
  "doctor" serial,
  "company" serial NOT NULL,
  "user_created" serial NOT NULL,
  "user_updated" serial,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "prescription_item" (
  "id" serial PRIMARY KEY,
  "prescription_uuid" uuid,
  "variant" serial,
  "lieu_dung" varchar,
  "quantity" int NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX ON "role_item" ("roles", "app");

CREATE INDEX ON "address" ("province");

CREATE INDEX ON "address" ("district");

CREATE INDEX ON "address" ("ward");

CREATE INDEX ON "provinces" ("administrative_unit_id");

CREATE INDEX ON "provinces" ("administrative_region_id");

CREATE INDEX ON "districts" ("province_code");

CREATE INDEX ON "districts" ("administrative_unit_id");

CREATE INDEX ON "wards" ("district_code");

CREATE INDEX ON "wards" ("administrative_unit_id");

CREATE INDEX ON "products" ("unit");

CREATE INDEX ON "products" ("id", "unit");

CREATE INDEX ON "company_pharma" ("company_pharma_type");

CREATE UNIQUE INDEX ON "company_pharma" ("name", "company_pharma_type");

CREATE INDEX ON "product_media" ("product");

CREATE INDEX ON "product_media" ("media");

CREATE UNIQUE INDEX ON "product_media" ("product", "media");

CREATE INDEX ON "variant_media" ("variant");

CREATE INDEX ON "variant_media" ("media");

CREATE UNIQUE INDEX ON "variant_media" ("variant", "media");

CREATE INDEX ON "orders" ("qr");

CREATE INDEX ON "orders" ("id", "qr");

CREATE INDEX ON "customers" ("address");

CREATE UNIQUE INDEX ON "customers" ("id", "address");

CREATE INDEX ON "tickets" ("qr");

CREATE UNIQUE INDEX ON "tickets" ("id", "qr");

CREATE INDEX ON "debt_repayment" ("debt", "money");

CREATE UNIQUE INDEX ON "service_variant" ("service", "variant");

COMMENT ON COLUMN "company_type"."code" IS '
💸 1 = CLINIC,
✔️ 2 = DRUGSTORE
';

COMMENT ON COLUMN "promotions"."time_apply" IS 'null: vố số, int: số lần cụ thể';

COMMENT ON COLUMN "promotions"."status" IS 'true: áp dụng, false: vô hiệu';

COMMENT ON COLUMN "promotion_type"."title" IS '1: khuyến mãi giảm giá, 2: khuyến mãi tặng sp';

COMMENT ON COLUMN "order_items"."consignment" IS 'Không truyền lúc tạo tự chọn theo BE';

COMMENT ON COLUMN "customers"."title" IS 'Chức danh';

COMMENT ON COLUMN "medical_records"."symptom" IS 'Triệu chứng';

COMMENT ON COLUMN "medical_records"."diagnostic" IS 'Chuẩn đoán';

COMMENT ON COLUMN "medical_records"."result" IS 'Kết luận';

COMMENT ON COLUMN "services"."reminder_time" IS 'Thời gian nhắc hẹn (seconds)';

COMMENT ON COLUMN "notification"."data" IS 'save code entity';

COMMENT ON COLUMN "appointment_schedules"."symptoms" IS 'Triệu chứng bệnh';

COMMENT ON COLUMN "appointment_schedules"."diagnostic" IS 'Chuẩn đoán bệnh';

COMMENT ON COLUMN "appointment_schedules"."is_done" IS 'true: Đã xong, false: Chưa diễn ra';

COMMENT ON COLUMN "medical_bills"."symptoms" IS 'Triệu chứng bệnh';

COMMENT ON COLUMN "medical_bills"."diagnostic" IS 'Chuẩn đoán bệnh';

COMMENT ON COLUMN "medical_bills"."is_done" IS 'true: Đã xong, false: Chưa diễn ra';

COMMENT ON COLUMN "prescriptions"."symptoms" IS 'Triệu chứng bệnh';

COMMENT ON COLUMN "prescriptions"."diagnostic" IS 'Chuẩn đoán bệnh';

ALTER TABLE "accounts" ADD FOREIGN KEY ("type") REFERENCES "account_type" ("id") ON DELETE CASCADE;

ALTER TABLE "accounts" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id") ON DELETE SET NULL;

ALTER TABLE "accounts" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "account_media" ADD FOREIGN KEY ("account") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "account_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE CASCADE;

ALTER TABLE "account_company" ADD FOREIGN KEY ("account") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "account_company" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "roles" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "role_item" ADD FOREIGN KEY ("roles") REFERENCES "roles" ("id") ON DELETE CASCADE;

ALTER TABLE "role_item" ADD FOREIGN KEY ("app") REFERENCES "apps" ("code") ON DELETE CASCADE;

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "accounts" ("username") ON DELETE CASCADE;

ALTER TABLE "verifies" ADD FOREIGN KEY ("username") REFERENCES "accounts" ("username") ON DELETE CASCADE;

ALTER TABLE "companies" ADD FOREIGN KEY ("owner") REFERENCES "accounts" ("id");

ALTER TABLE "companies" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "companies" ADD FOREIGN KEY ("type") REFERENCES "company_type" ("code") ON DELETE SET NULL;

ALTER TABLE "address" ADD FOREIGN KEY ("province") REFERENCES "provinces" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("district") REFERENCES "districts" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("ward") REFERENCES "wards" ("code");

ALTER TABLE "address" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "provinces" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "provinces" ADD FOREIGN KEY ("administrative_region_id") REFERENCES "administrative_regions" ("id");

ALTER TABLE "districts" ADD FOREIGN KEY ("province_code") REFERENCES "provinces" ("code");

ALTER TABLE "districts" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "wards" ADD FOREIGN KEY ("district_code") REFERENCES "districts" ("code");

ALTER TABLE "wards" ADD FOREIGN KEY ("administrative_unit_id") REFERENCES "administrative_units" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("product_category") REFERENCES "product_categories" ("id") ON DELETE SET NULL;

ALTER TABLE "products" ADD FOREIGN KEY ("type") REFERENCES "product_type" ("id") ON DELETE SET NULL;

ALTER TABLE "products" ADD FOREIGN KEY ("brand") REFERENCES "product_brand" ("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("phan_loai") REFERENCES "classify" ("code") ON DELETE SET NULL;

ALTER TABLE "products" ADD FOREIGN KEY ("dang_bao_che") REFERENCES "preparation_type" ("code") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("tieu_chuan_sx") REFERENCES "production_standard" ("code") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("cong_ty_sx") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("cong_ty_dk") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "products" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "products_bank" ADD FOREIGN KEY ("phan_loai") REFERENCES "classify" ("code") ON DELETE SET NULL;

ALTER TABLE "products_bank" ADD FOREIGN KEY ("dang_bao_che") REFERENCES "preparation_type" ("code") ON DELETE CASCADE;

ALTER TABLE "products_bank" ADD FOREIGN KEY ("tieu_chuan_sx") REFERENCES "production_standard" ("code") ON DELETE CASCADE;

ALTER TABLE "products_bank" ADD FOREIGN KEY ("cong_ty_sx") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE "products_bank" ADD FOREIGN KEY ("cong_ty_dk") REFERENCES "company_pharma" ("id") ON DELETE CASCADE;

ALTER TABLE "price_list" ADD FOREIGN KEY ("variant_code") REFERENCES "variants" ("code") ON DELETE CASCADE;

ALTER TABLE "price_list" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id") ON DELETE CASCADE;

ALTER TABLE "price_list" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "price_list" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "price_list_log" ADD FOREIGN KEY ("price_list") REFERENCES "price_list" ("id") ON DELETE CASCADE;

ALTER TABLE "price_list_log" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "company_pharma" ADD FOREIGN KEY ("company_pharma_type") REFERENCES "company_pharma_type" ("code") ON DELETE SET NULL;

ALTER TABLE "product_media" ADD FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "product_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE CASCADE;

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "product_brand" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_brand" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "product_type" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_type" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "ingredient" ADD FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "units" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "units" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id") ON DELETE CASCADE;

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "variants" ADD FOREIGN KEY ("product") REFERENCES "products" ("id") ON DELETE CASCADE;

ALTER TABLE "variants" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "variants" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "variant_media" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE CASCADE;

ALTER TABLE "variant_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE CASCADE;

ALTER TABLE "promotions" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "promotions" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "promotions" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "promotions" ADD FOREIGN KEY ("type") REFERENCES "promotion_type" ("code") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ADD FOREIGN KEY ("promotions") REFERENCES "promotions" ("id") ON DELETE CASCADE;

ALTER TABLE "promotion_item" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ADD FOREIGN KEY ("applicable_variant") REFERENCES "variants" ("id") ON DELETE SET NULL;

ALTER TABLE "promotion_item" ADD FOREIGN KEY ("applicable_service") REFERENCES "services" ("id") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("status") REFERENCES "order_status" ("code") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("type") REFERENCES "order_type" ("code") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("ticket") REFERENCES "tickets" ("id") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("qr") REFERENCES "medias" ("id") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "orders" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "orders" ADD FOREIGN KEY ("payment") REFERENCES "payments" ("id") ON DELETE CASCADE;

ALTER TABLE "orders" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "order_items" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id") ON DELETE CASCADE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE CASCADE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("consignment") REFERENCES "consignment" ("id") ON DELETE SET NULL;

ALTER TABLE "order_items" ADD FOREIGN KEY ("consignment_log") REFERENCES "consignment_log" ("id") ON DELETE SET NULL;

ALTER TABLE "service_order_item" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id") ON DELETE CASCADE;

ALTER TABLE "service_order_item" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE SET NULL;

ALTER TABLE "customers" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "customers" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "customers" ADD FOREIGN KEY ("contact_address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "customers" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "customers" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "customers" ADD FOREIGN KEY ("group") REFERENCES "customer_group" ("id") ON DELETE SET NULL;

ALTER TABLE "customer_group" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "customer_group" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "customer_group" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_records" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_record_variant" ADD FOREIGN KEY ("medical_record") REFERENCES "medical_records" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_record_variant" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;

ALTER TABLE "payment_items" ADD FOREIGN KEY ("type") REFERENCES "payment_item_types" ("code") ON DELETE CASCADE;

ALTER TABLE "payment_items" ADD FOREIGN KEY ("payment") REFERENCES "payments" ("id") ON DELETE CASCADE;

ALTER TABLE "warehouses" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "warehouses" ADD FOREIGN KEY ("companies") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("type") REFERENCES "ticket_type" ("id") ON DELETE SET NULL;

ALTER TABLE "tickets" ADD FOREIGN KEY ("status") REFERENCES "ticket_status" ("id") ON DELETE SET NULL;

ALTER TABLE "tickets" ADD FOREIGN KEY ("qr") REFERENCES "medias" ("id") ON DELETE SET NULL;

ALTER TABLE "tickets" ADD FOREIGN KEY ("export_to") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "tickets" ADD FOREIGN KEY ("import_from") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "tickets" ADD FOREIGN KEY ("warehouse") REFERENCES "warehouses" ("id") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "consignment" ADD FOREIGN KEY ("ticket") REFERENCES "tickets" ("id") ON DELETE CASCADE;

ALTER TABLE "consignment" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE CASCADE;

ALTER TABLE "consignment" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "consignment" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "consignment_log" ADD FOREIGN KEY ("consignment") REFERENCES "consignment" ("id") ON DELETE CASCADE;

ALTER TABLE "consignment_log" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "suplier" ADD FOREIGN KEY ("warehouses") REFERENCES "warehouses" ("id") ON DELETE SET NULL;

ALTER TABLE "suplier" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "suplier" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("type") REFERENCES "debt_note_type" ("code") ON DELETE SET NULL;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("status") REFERENCES "debt_note_status" ("code") ON DELETE SET NULL;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "debt_repayment" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "debt_repayment" ADD FOREIGN KEY ("debt") REFERENCES "debt_note" ("id") ON DELETE SET NULL;

ALTER TABLE "services" ADD FOREIGN KEY ("staff") REFERENCES "accounts" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("image") REFERENCES "medias" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("brand") REFERENCES "product_brand" ("id") ON DELETE SET NULL;

ALTER TABLE "services" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "services" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "service_variant" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE CASCADE;

ALTER TABLE "service_variant" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("type") REFERENCES "noti_type" ("code") ON DELETE SET NULL;

ALTER TABLE "notification" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "appointment_schedules" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("as_uuid") REFERENCES "appointment_schedules" ("uuid");

ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("mb_uuid") REFERENCES "medical_bills" ("uuid");

ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_service" ADD FOREIGN KEY ("order_service") REFERENCES "orders" ("id") ON DELETE SET NULL;

ALTER TABLE "appointment_schedule_url" ADD FOREIGN KEY ("as_uuid") REFERENCES "appointment_schedules" ("uuid");

ALTER TABLE "appointment_schedule_url" ADD FOREIGN KEY ("mb_uuid") REFERENCES "medical_bills" ("uuid");

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_bills" ADD FOREIGN KEY ("prescription") REFERENCES "prescriptions" ("uuid") ON DELETE SET NULL;

ALTER TABLE "medical_bill_order_sell" ADD FOREIGN KEY ("uuid") REFERENCES "medical_bills" ("uuid");

ALTER TABLE "medical_bill_order_sell" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id") ON DELETE SET NULL;

ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE CASCADE;

ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("appointment_schedule") REFERENCES "appointment_schedules" ("uuid") ON DELETE SET NULL;

ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("medical_bill") REFERENCES "medical_bills" ("uuid") ON DELETE SET NULL;

ALTER TABLE "medical_record_link" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("doctor") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id") ON DELETE SET NULL;

ALTER TABLE "prescriptions" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "prescription_item" ADD FOREIGN KEY ("prescription_uuid") REFERENCES "prescriptions" ("uuid");

ALTER TABLE "prescription_item" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id") ON DELETE SET NULL;
