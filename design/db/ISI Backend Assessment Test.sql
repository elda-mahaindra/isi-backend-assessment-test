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
  "no_rekening" bigserial PRIMARY KEY,
  "saldo" bigint NOT NULL
);

CREATE TABLE "entry_types" (
  "id" smallserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "entries" (
  "no_rekening" bigint NOT NULL,
  "nominal" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "id" bigserial PRIMARY KEY,
  "type_id" smallint NOT NULL
);

CREATE INDEX ON "accounts" ("customer_id");

CREATE INDEX ON "entries" ("no_rekening");

CREATE INDEX ON "entries" ("type_id");

CREATE INDEX ON "entries" ("no_rekening", "type_id");

COMMENT ON TABLE "entry_types" IS 'init it with types of DEPOSIT, WITHDRAWAL';

COMMENT ON COLUMN "entries"."nominal" IS 'can be negative or positive';

ALTER TABLE "accounts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("no_rekening") REFERENCES "accounts" ("no_rekening");

ALTER TABLE "entries" ADD FOREIGN KEY ("type_id") REFERENCES "entry_types" ("id");
