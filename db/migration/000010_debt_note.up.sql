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

CREATE INDEX ON "debt_repayment" ("debt", "money");

ALTER TABLE "debt_note" ADD FOREIGN KEY ("type") REFERENCES "debt_note_type" ("code") ON DELETE SET NULL;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("status") REFERENCES "debt_note_status" ("code") ON DELETE SET NULL;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "debt_note" ADD FOREIGN KEY ("company") REFERENCES "companies" ("id") ON DELETE SET NULL;

ALTER TABLE "debt_repayment" ADD FOREIGN KEY ("user_created") REFERENCES "accounts" ("id") ON DELETE SET NULL;

ALTER TABLE "debt_repayment" ADD FOREIGN KEY ("debt") REFERENCES "debt_note" ("id") ON DELETE SET NULL;


INSERT INTO debt_note_type (code, title) VALUES ('REVENUE', 'Khoản thu');
INSERT INTO debt_note_type (code, title) VALUES ('EXPENSE', 'Khoản chi');

INSERT INTO debt_note_status (code, title) VALUES ('OPEN', 'Chưa thanh toán');
INSERT INTO debt_note_status (code, title) VALUES ('REPAYING', 'Thanh toán 1 phần');
INSERT INTO debt_note_status (code, title) VALUES ('SETTLED', 'Hoàn thành');
INSERT INTO debt_note_status (code, title) VALUES ('OVERDUE', 'Quá hạn');