CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar UNIQUE NOT NULL
);

CREATE TABLE "clients" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "account_number" varchar,
  "preferred_payment_type" varchar UNIQUE NOT NULL
);

CREATE TABLE "requests" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'PENDING',
  "amount" bigint NOT NULL,
  "paid_to" bigint,
  "createdby_id" uuid NOT NULL,
  "approvedby_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "approved_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_payments" (
  "id" bigserial PRIMARY KEY,
  "request_id" bigint,
  "client_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "requests" ("createdby_id", "approvedby_id");

CREATE UNIQUE INDEX ON "user_payments" ("client_id", "request_id");

COMMENT ON COLUMN "requests"."status" IS 'Payment Status can be PENDING or RESOLVED';

ALTER TABLE "requests" ADD FOREIGN KEY ("paid_to") REFERENCES "clients" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("createdby_id") REFERENCES "users" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("approvedby_id") REFERENCES "users" ("id");

ALTER TABLE "user_payments" ADD FOREIGN KEY ("request_id") REFERENCES "requests" ("id");

ALTER TABLE "user_payments" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");