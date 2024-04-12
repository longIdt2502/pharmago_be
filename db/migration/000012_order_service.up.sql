CREATE TABLE "service_order_item" (
  "id" serial PRIMARY KEY,
  "order" serial NOT NULL,
  "service" serial,
  "unit_price" float NOT NULL DEFAULT 0,
  "discount" float NOT NULL DEFAULT 0,
  "total_price" float NOT NULL DEFAULT 0
);

ALTER TABLE "service_order_item" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id") ON DELETE CASCADE;

ALTER TABLE "service_order_item" ADD FOREIGN KEY ("service") REFERENCES "services" ("id") ON DELETE SET NULL;

