DROP TRIGGER IF EXISTS update_users_updated_at ON users;

DROP FUNCTION IF EXISTS update_updated_at_column ();

DROP TABLE IF EXISTS outbox_events CASCADE;

DROP TABLE IF EXISTS password_reset_tokens CASCADE;

DROP TABLE IF EXISTS email_verifications CASCADE;

DROP TABLE IF EXISTS user_sessions CASCADE;

DROP TABLE IF EXISTS refresh_tokens CASCADE;

DROP TABLE IF EXISTS users CASCADE;

DROP TYPE IF EXISTS user_status CASCADE;