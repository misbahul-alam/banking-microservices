CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE user_status AS ENUM ('active', 'suspended', 'deleted');

CREATE TABLE
    users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        email VARCHAR(255) NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        phone VARCHAR(30) UNIQUE,
        status user_status NOT NULL DEFAULT 'active',
        email_verified BOOLEAN NOT NULL DEFAULT FALSE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE UNIQUE INDEX users_email_uq_idx ON users (email);

CREATE UNIQUE INDEX users_phone_uq_idx ON users (phone)
WHERE
    phone IS NOT NULL;

CREATE INDEX users_status_created_idx ON users (status, created_at);

CREATE TABLE
    refresh_tokens (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        token_hash TEXT NOT NULL UNIQUE,
        expires_at TIMESTAMPTZ NOT NULL,
        revoked_at TIMESTAMPTZ,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE UNIQUE INDEX refresh_tokens_hash_uq_idx ON refresh_tokens (token_hash);

CREATE INDEX refresh_tokens_user_idx ON refresh_tokens (user_id);

CREATE INDEX refresh_tokens_active_idx ON refresh_tokens (user_id, expires_at)
WHERE
    revoked_at IS NULL;

CREATE TABLE
    user_sessions (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        refresh_token_id UUID REFERENCES refresh_tokens (id) ON DELETE SET NULL,
        ip_address VARCHAR(45),
        user_agent TEXT,
        device_name VARCHAR(100),
        last_activity TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        revoked_at TIMESTAMPTZ,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE INDEX user_sessions_user_idx ON user_sessions (user_id);

CREATE INDEX user_sessions_token_idx ON user_sessions (refresh_token_id);

CREATE INDEX user_sessions_active_activity_idx ON user_sessions (user_id, last_activity DESC)
WHERE
    revoked_at IS NULL;

CREATE TABLE
    email_verifications (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        token_hash TEXT NOT NULL UNIQUE,
        expires_at TIMESTAMPTZ NOT NULL,
        used_at TIMESTAMPTZ,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE UNIQUE INDEX email_verifications_hash_idx ON email_verifications (token_hash);

CREATE INDEX email_verifications_user_idx ON email_verifications (user_id);

CREATE TABLE
    password_reset_tokens (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        token_hash TEXT NOT NULL UNIQUE,
        expires_at TIMESTAMPTZ NOT NULL,
        used_at TIMESTAMPTZ,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE UNIQUE INDEX password_reset_tokens_hash_idx ON password_reset_tokens (token_hash);

CREATE INDEX password_reset_tokens_user_idx ON password_reset_tokens (user_id);

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

CREATE TRIGGER update_users_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();