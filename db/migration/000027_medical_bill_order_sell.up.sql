CREATE TABLE "medical_bill_order_sell" (
  "uuid" uuid,
  "order" serial
);

ALTER TABLE "medical_bill_order_sell" ADD FOREIGN KEY ("uuid") REFERENCES "medical_bills" ("uuid");

ALTER TABLE "medical_bill_order_sell" ADD FOREIGN KEY ("order") REFERENCES "orders" ("id") ON DELETE SET NULL;