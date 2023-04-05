CREATE TABLE "daraja_token" (
  "id" bigserial PRIMARY KEY,
  "access_token" varchar NOT NULL,
  "expires_at" timestamptz NOT NULL DEFAULT (now())
);