CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "username" varchar UNIQUE NOT NULL,
                            "hashed_password" varchar NOT NULL,
                            "full_name" varchar NOT NULL,
                            "email" varchar UNIQUE NOT NULL,
                            "type" bigserial NOT NULL,
                            "media" bigserial,
                            "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_type" (
                                "id" bigserial PRIMARY KEY,
                                "code" varchar NOT NULL,
                                "title" varchar NOT NULL
);

CREATE TABLE "companies" (
                             "id" bigserial PRIMARY KEY,
                             "name" varchar NOT NULL,
                             "tax_code" varchar,
                             "phone" varchar,
                             "description" varchar,
                             "created_at" timestamptz NOT NULL DEFAULT (now()),
                             "owner" bigserial
);

CREATE TABLE "address" (
                           "id" bigserial PRIMARY KEY,
                           "lat" numeric NOT NULL,
                           "lng" numeric NOT NULL,
                           "title" varchar NOT NULL,
                           "user_created" bigserial NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "warehouses" (
                              "id" bigserial PRIMARY KEY,
                              "address" bigserial,
                              "companies" bigserial
);

CREATE TABLE "products" (
                            "id" bigserial PRIMARY KEY,
                            "name" varchar NOT NULL,
                            "code" varchar NOT NULL,
                            "product_category" bigserial,
                            "type" bigserial,
                            "unit" bigserial NOT NULL,
                            "company" bigserial,
                            "user_created" bigserial NOT NULL,
                            "user_updated" bigserial,
                            "updated_at" timestamptz,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_media" (
                                 "id" bigserial PRIMARY KEY,
                                 "product" bigserial,
                                 "media" bigserial
);

CREATE TABLE "product_categories" (
                                      "id" bigserial PRIMARY KEY,
                                      "name" varchar NOT NULL,
                                      "user_created" bigserial NOT NULL,
                                      "user_updated" bigserial,
                                      "updated_at" timestamptz,
                                      "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_type" (
                                "id" bigserial PRIMARY KEY,
                                "name" varchar NOT NULL
);

CREATE TABLE "units" (
                         "id" bigserial PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "sell_price" numeric NOT NULL DEFAULT 0,
                         "import_price" numeric NOT NULL DEFAULT 0,
                         "weight" numeric,
                         "weight_unit" varchar,
                         "user_created" bigserial NOT NULL,
                         "user_updated" bigserial,
                         "updated_at" timestamptz,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "unit_changes" (
                                "id" bigserial PRIMARY KEY,
                                "name" varchar NOT NULL,
                                "value" bigint NOT NULL DEFAULT 0,
                                "sell_price" numeric NOT NULL DEFAULT 0,
                                "unit" bigserial,
                                "user_created" bigserial NOT NULL,
                                "user_updated" bigserial,
                                "updated_at" timestamptz,
                                "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "variants" (
                            "id" bigserial PRIMARY KEY,
                            "name" varchar NOT NULL,
                            "code" varchar NOT NULL,
                            "barcode" varchar NOT NULL,
                            "decision_number" bigint NOT NULL,
                            "register_number" bigint NOT NULL,
                            "discount" numeric NOT NULL DEFAULT 0,
                            "vat" numeric NOT NULL DEFAULT 0,
                            "product" bigserial,
                            "user_created" bigserial NOT NULL,
                            "user_updated" bigserial,
                            "updated_at" timestamptz,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
                          "id" bigserial PRIMARY KEY,
                          "vat" numeric NOT NULL DEFAULT 0,
                          "total_price" numeric NOT NULL DEFAULT 0,
                          "description" varchar,
                          "customer" bigserial,
                          "status" bigserial NOT NULL,
                          "type" bigserial,
                          "ticket" bigserial,
                          "qr" bigserial
);

CREATE TABLE "order_type" (
                              "id" bigserial PRIMARY KEY,
                              "code" varchar NOT NULL,
                              "title" varchar NOT NULL
);

CREATE TABLE "order_status" (
                                "id" bigserial PRIMARY KEY,
                                "code" varchar NOT NULL,
                                "title" varchar NOT NULL
);

CREATE TABLE "order_items" (
                               "id" bigserial PRIMARY KEY,
                               "order" bigserial,
                               "variant" bigserial,
                               "value" int NOT NULL DEFAULT 0
);

CREATE TABLE "customers" (
                             "id" bigserial PRIMARY KEY,
                             "full_name" varchar NOT NULL,
                             "code" varchar NOT NULL,
                             "company" bigserial,
                             "address" bigserial,
                             "email" varchar,
                             "birthday" timestamptz,
                             "user_created" bigserial NOT NULL,
                             "user_updated" bigserial,
                             "updated_at" timestamptz,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tickets" (
                           "id" bigserial PRIMARY KEY,
                           "code" varchar NOT NULL,
                           "type" bigserial NOT NULL,
                           "status" bigserial NOT NULL,
                           "note" varchar,
                           "qr" bigserial,
                           "export_from" bigserial NOT NULL,
                           "import_to" bigserial NOT NULL,
                           "user_created" bigserial NOT NULL,
                           "user_updated" bigserial,
                           "updated_at" timestamptz,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ticket_type" (
                               "id" bigserial PRIMARY KEY,
                               "code" varchar NOT NULL,
                               "title" varchar NOT NULL
);

CREATE TABLE "ticket_status" (
                                 "id" bigserial PRIMARY KEY,
                                 "code" varchar NOT NULL,
                                 "title" varchar NOT NULL
);

CREATE TABLE "medias" (
                          "id" bigserial PRIMARY KEY,
                          "media_url" varchar NOT NULL
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("type") REFERENCES "account_type" ("id");

ALTER TABLE "accounts" ALTER COLUMN "media" DROP NOT NULL;
ALTER TABLE "accounts" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;

-- ALTER TABLE "accounts" ADD CONSTRAINT "fk_media" FOREIGN KEY ("media") REFERENCES "medias" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;

ALTER TABLE "companies" ADD FOREIGN KEY ("owner") REFERENCES "accounts" ("id");

ALTER TABLE "address" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "warehouses" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");

ALTER TABLE "warehouses" ADD FOREIGN KEY ("companies") REFERENCES "companies" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("product_category") REFERENCES "product_categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("type") REFERENCES "product_type" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "product_media" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "product_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "product_categories" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "units" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "units" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id");

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "unit_changes" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "variants" ADD FOREIGN KEY ("product") REFERENCES "products" ("id");

ALTER TABLE "variants" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "variants" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer") REFERENCES "customers" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("status") REFERENCES "order_status" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("type") REFERENCES "order_type" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("ticket") REFERENCES "tickets" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("qr") REFERENCES "medias" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("address") REFERENCES "address" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "customers" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("type") REFERENCES "ticket_type" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("status") REFERENCES "ticket_status" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("qr") REFERENCES "medias" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("export_from") REFERENCES "warehouses" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("import_to") REFERENCES "warehouses" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");
