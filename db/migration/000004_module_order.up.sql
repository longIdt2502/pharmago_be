DROP INDEX IF EXISTS "orders_qr_idx";
DROP INDEX IF EXISTS "orders_id_qr_idx";
DROP INDEX IF EXISTS "customers_address_idx";
DROP INDEX IF EXISTS "customers_id_address_idx";

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

DROP TABLE IF EXISTS "orders" CASCADE;
DROP TABLE IF EXISTS "order_items" CASCADE;
DROP TABLE IF EXISTS "customers" CASCADE;
DROP TABLE IF EXISTS "order_type" CASCADE;
DROP TABLE IF EXISTS "order_status" CASCADE;

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

CREATE TABLE "order_items" (
    "id" serial PRIMARY KEY,
    "order" serial NOT NULL,
    "variant" serial NOT NULL,
    "value" int NOT NULL DEFAULT 0,
    "total_price" float NOT NULL DEFAULT 0,
    "consignment" serial,
    "consignment_log" serial
);

CREATE TABLE "customers" (
     "id" serial PRIMARY KEY,
     "full_name" varchar NOT NULL,
     "code" varchar NOT NULL,
     "company" serial NOT NULL,
     "address" serial,
     "email" varchar,
     "phone" varchar(20),
     "license" varchar(20),
     "birthday" timestamptz,
     "user_created" serial NOT NULL,
     "user_updated" serial,
     "updated_at" timestamptz,
     "created_at" timestamptz NOT NULL DEFAULT (now())
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

ALTER TABLE "customers" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE CASCADE;

ALTER TABLE "customers" ADD FOREIGN KEY ("address") REFERENCES "address" ("id") ON DELETE SET NULL;

ALTER TABLE "customers" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "customers" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "payment_items" ADD FOREIGN KEY ("payment") REFERENCES "payments" ("id") ON DELETE CASCADE;

ALTER TABLE "payment_items" ADD FOREIGN KEY ("type") REFERENCES "payment_item_types" ("code") ON DELETE CASCADE;