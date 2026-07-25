DROP TRIGGER IF EXISTS update_transactions_updated_at ON transactions;

DROP FUNCTION IF EXISTS update_updated_at_column ();

DROP TABLE IF EXISTS ledger_entries CASCADE;

DROP TABLE IF EXISTS transactions CASCADE;

DROP TYPE IF EXISTS entry_type CASCADE;

DROP TYPE IF EXISTS transaction_status CASCADE;

DROP TYPE IF EXISTS transaction_type CASCADE;