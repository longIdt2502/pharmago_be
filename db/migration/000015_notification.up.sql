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

INSERT INTO "noti_type" (code) VALUES ('ORDER');
INSERT INTO "noti_type" (code) VALUES ('SERVICE');

COMMENT ON COLUMN "notification"."data" IS 'save code entity';

ALTER TABLE "notification" ADD FOREIGN KEY ("type") REFERENCES "noti_type" ("code") ON DELETE SET NULL;

ALTER TABLE "notification" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;