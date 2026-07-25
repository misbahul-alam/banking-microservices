DROP TRIGGER IF EXISTS update_beneficiaries_updated_at ON beneficiaries;

DROP TRIGGER IF EXISTS update_accounts_updated_at ON accounts;

DROP TRIGGER IF EXISTS update_customers_updated_at ON customers;

DROP FUNCTION IF EXISTS update_updated_at_column ();

DROP TABLE IF EXISTS outbox_events CASCADE;

DROP TABLE IF EXISTS beneficiaries CASCADE;

DROP TABLE IF EXISTS accounts CASCADE;

DROP TABLE IF EXISTS customers CASCADE;

DROP TYPE IF EXISTS account_status CASCADE;

DROP TYPE IF EXISTS account_type CASCADE;