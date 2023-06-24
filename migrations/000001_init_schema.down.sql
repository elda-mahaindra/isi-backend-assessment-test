ALTER TABLE "accounts" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("no_rekening") REFERENCES "accounts" ("no_rekening");

DROP TABLE IF EXISTS "entries";
DROP TABLE IF EXISTS "accounts";
DROP TABLE IF EXISTS "customers";