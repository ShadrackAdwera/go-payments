CREATE TYPE "payment_types" AS ENUM (
  'master_card',
  'visa',
  'mpesa'
);

CREATE TYPE "approval_status" AS ENUM (
  'pending',
  'approved',
  'rejected'
);

CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar NOT NULL
);

CREATE TABLE "clients" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "account_number" varchar,
  "preferred_payment_type" payment_types NOT NULL
);

CREATE TABLE "requests" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "status" approval_status NOT NULL,
  "amount" bigint NOT NULL,
  "paid_to_id" bigint NOT NULL,
  "createdby_id" varchar NOT NULL,
  "approvedby_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "approved_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_payments" (
  "id" bigserial PRIMARY KEY,
  "request_id" bigint,
  "client_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "permissions" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "role_id" bigint NOT NULL
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "users_roles" (
  "id" bigserial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "role_id" bigint NOT NULL
);

CREATE UNIQUE INDEX ON "requests" ("createdby_id", "approvedby_id");

CREATE UNIQUE INDEX ON "user_payments" ("client_id", "request_id");

COMMENT ON COLUMN "requests"."status" IS 'Payment Status can be PENDING or RESOLVED';

ALTER TABLE "requests" ADD FOREIGN KEY ("paid_to_id") REFERENCES "clients" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("createdby_id") REFERENCES "users" ("id");

ALTER TABLE "requests" ADD FOREIGN KEY ("approvedby_id") REFERENCES "users" ("id");

ALTER TABLE "user_payments" ADD FOREIGN KEY ("request_id") REFERENCES "requests" ("id");

ALTER TABLE "user_payments" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "users_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");