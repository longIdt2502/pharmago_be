CREATE TABLE "price_list" (
    "id" bigserial PRIMARY KEY,
    "variant_code" varchar UNIQUE NOT NULL,
    "variant_name" varchar NOT NULL,
    "price_import" float NOT NULL,
    "price_sell" float NOT NULL,
    "unit" bigserial NOT NULL,
    "user_created" bigserial NOT NULL,
    "user_updated" bigserial,
    "updated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "price_list_log" (
    "id" bigserial PRIMARY KEY,
    "old_price_import" float NOT NULL,
    "new_price_import" float NOT NULL,
    "old_price_sell" float NOT NULL,
    "new_price_sell" float NOT NULL,
    "price_list" bigserial NOT NULL,
    "user_updated" bigserial NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "price_list" ADD FOREIGN KEY ("variant_code") REFERENCES "variants" ("code");

ALTER TABLE "price_list" ADD FOREIGN KEY ("unit") REFERENCES "units" ("id");

ALTER TABLE "price_list" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id");

ALTER TABLE "price_list" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");

ALTER TABLE "price_list_log" ADD FOREIGN KEY ("price_list") REFERENCES "price_list" ("id");

ALTER TABLE "price_list_log" ADD FOREIGN KEY ("user_updated") REFERENCES "accounts" ("id");