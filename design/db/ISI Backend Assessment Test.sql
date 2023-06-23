CREATE TABLE "customers" (
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "id" bigserial PRIMARY KEY,
  "nama" varchar NOT NULL,
  "nik" varchar UNIQUE NOT NULL,
  "no_hp" varchar UNIQUE NOT NULL
);

CREATE TABLE "accounts" (
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "customer_id" bigint NOT NULL,
  "no_rekening" varchar PRIMARY KEY,
  "saldo" bigint NOT NULL
);

CREATE TABLE "entries" (
  "code" varchar(1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "id" bigserial PRIMARY KEY,
  "nominal" bigint NOT NULL,
  "no_rekening" varchar NOT NULL
);

COMMENT ON COLUMN "entries"."nominal" IS 'can be negative or positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("no_rekening") REFERENCES "accounts" ("no_rekening");
