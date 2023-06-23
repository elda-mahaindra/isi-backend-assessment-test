ALTER TABLE "entries" DROP CONSTRAINT IF EXISTS "entries_type_id_fkey";
ALTER TABLE "entries" DROP CONSTRAINT IF EXISTS "entries_no_rekening_fkey";

ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_customer_id_fkey";

DROP TABLE IF EXISTS "entries";
DROP TABLE IF EXISTS "entry_types";
DROP TABLE IF EXISTS "accounts";
DROP TABLE IF EXISTS "customers";