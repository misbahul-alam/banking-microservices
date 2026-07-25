CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE
    notification_preferences (
        user_id UUID PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        phone VARCHAR(30),
        email_enabled BOOLEAN NOT NULL DEFAULT TRUE,
        sms_enabled BOOLEAN NOT NULL DEFAULT TRUE,
        push_enabled BOOLEAN NOT NULL DEFAULT TRUE,
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE TABLE
    notification_templates (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        type VARCHAR(50) NOT NULL UNIQUE,
        channel VARCHAR(20) NOT NULL,
        subject VARCHAR(255),
        body TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

CREATE TABLE
    notification_logs (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL,
        channel VARCHAR(20) NOT NULL,
        status VARCHAR(30) NOT NULL,
        provider VARCHAR(100),
        template_id UUID REFERENCES notification_templates (id),
        payload JSONB,
        error_message TEXT,
        retry_count INT NOT NULL DEFAULT 0,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        delivered_at TIMESTAMPTZ
    );

CREATE INDEX notification_logs_user_recent_idx ON notification_logs (user_id, created_at DESC);

CREATE INDEX notification_logs_status_retry_idx ON notification_logs (status, created_at);

CREATE INDEX notification_logs_template_idx ON notification_logs (template_id);

CREATE
OR REPLACE FUNCTION update_updated_at_column () RETURNS TRIGGER AS '
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
' LANGUAGE plpgsql;

CREATE TRIGGER update_notification_preferences_updated_at BEFORE
UPDATE ON notification_preferences FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();

CREATE TRIGGER update_notification_templates_updated_at BEFORE
UPDATE ON notification_templates FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();