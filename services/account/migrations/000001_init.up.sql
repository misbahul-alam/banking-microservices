CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE account_type AS ENUM ('savings', 'checking', 'business');

CREATE TYPE account_status AS ENUM ('active', 'frozen', 'closed');

CREATE TABLE
    customers (
        user_id UUID PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        phone VARCHAR(30),
        status VARCHAR(20) NOT NULL,
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE INDEX customers_email_idx ON customers (email);

CREATE INDEX customers_status_idx ON customers (status);

CREATE TABLE
    accounts (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        customer_id UUID NOT NULL REFERENCES customers (user_id),
        account_number VARCHAR(30) NOT NULL UNIQUE,
        account_type account_type NOT NULL,
        currency CHAR(3) NOT NULL,
        balance NUMERIC(20, 4) NOT NULL DEFAULT 0,
        version INT NOT NULL DEFAULT 0,
        status account_status NOT NULL DEFAULT 'active',
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE UNIQUE INDEX accounts_num_uq_idx ON accounts (account_number);

CREATE INDEX accounts_customer_idx ON accounts (customer_id);

CREATE INDEX accounts_status_cust_idx ON accounts (status, customer_id);

CREATE TABLE
    beneficiaries (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        account_id UUID NOT NULL REFERENCES accounts (id) ON DELETE CASCADE,
        beneficiary_account VARCHAR(30) NOT NULL,
        beneficiary_name VARCHAR(150) NOT NULL,
        bank_name VARCHAR(150),
        nickname VARCHAR(100),
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE INDEX beneficiaries_account_idx ON beneficiaries (account_id);

CREATE UNIQUE INDEX beneficiaries_account_target_uq_idx ON beneficiaries (account_id, beneficiary_account);

CREATE TABLE
    outbox_events (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        aggregate_type VARCHAR(50) NOT NULL,
        aggregate_id UUID NOT NULL,
        event_type VARCHAR(100) NOT NULL,
        payload JSONB NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        published_at TIMESTAMPTZ,
        attempts INT NOT NULL DEFAULT 0,
        last_error TEXT
    );

CREATE INDEX outbox_events_unpublished_idx ON outbox_events (created_at)
WHERE
    published_at IS NULL;

CREATE INDEX outbox_events_aggregate_idx ON outbox_events (aggregate_type, aggregate_id);

CREATE INDEX outbox_events_published_cleanup_idx ON outbox_events (published_at);

CREATE
OR REPLACE FUNCTION update_updated_at_column () RETURNS TRIGGER AS '
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
' LANGUAGE plpgsql;

CREATE TRIGGER update_customers_updated_at BEFORE
UPDATE ON customers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();

CREATE TRIGGER update_accounts_updated_at BEFORE
UPDATE ON accounts FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();

CREATE TRIGGER update_beneficiaries_updated_at BEFORE
UPDATE ON beneficiaries FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();