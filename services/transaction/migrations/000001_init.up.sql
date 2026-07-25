CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE transaction_type AS ENUM ('deposit', 'withdrawal', 'transfer');

CREATE TYPE transaction_status AS ENUM (
    'pending',
    'processing',
    'completed',
    'failed',
    'cancelled'
);

CREATE TYPE entry_type AS ENUM ('debit', 'credit');

CREATE TABLE
    transactions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        reference VARCHAR(50) NOT NULL UNIQUE,
        idempotency_key VARCHAR(64) NOT NULL UNIQUE,
        type transaction_type NOT NULL,
        status transaction_status NOT NULL DEFAULT 'pending',
        initiated_by UUID NOT NULL,
        description TEXT,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE UNIQUE INDEX transactions_ref_uq_idx ON transactions (reference);

CREATE UNIQUE INDEX transactions_idem_uq_idx ON transactions (idempotency_key);

CREATE INDEX transactions_user_created_idx ON transactions (initiated_by, created_at DESC);

CREATE INDEX transactions_status_created_idx ON transactions (status, created_at);

CREATE TABLE
    ledger_entries (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        transaction_id UUID NOT NULL REFERENCES transactions (id),
        account_id UUID NOT NULL,
        entry_type entry_type NOT NULL,
        amount NUMERIC(20, 4) NOT NULL CHECK (amount > 0),
        currency CHAR(3) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE INDEX ledger_tx_idx ON ledger_entries (transaction_id);

CREATE INDEX ledger_account_stmt_idx ON ledger_entries (account_id, created_at DESC);

CREATE INDEX ledger_account_curr_idx ON ledger_entries (account_id, currency, created_at);

CREATE
OR REPLACE FUNCTION update_updated_at_column () RETURNS TRIGGER AS '
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
' LANGUAGE plpgsql;

CREATE TRIGGER update_transactions_updated_at BEFORE
UPDATE ON transactions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();