CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestampz NOT NULL DEFAULT (0001-01-01 00:00:00Z),
  "created_at" timestampz NOT NULL DEFAULT (now())
);

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "account" ("owner", "currency"); 
-- ATAS BAWAH SAMA AJA, PILIH AJA SALAH SATU

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");