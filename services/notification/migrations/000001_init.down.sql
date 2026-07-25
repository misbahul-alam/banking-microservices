DROP TRIGGER IF EXISTS update_notification_templates_updated_at ON notification_templates;

DROP TRIGGER IF EXISTS update_notification_preferences_updated_at ON notification_preferences;

DROP FUNCTION IF EXISTS update_updated_at_column ();

DROP TABLE IF EXISTS notification_logs CASCADE;

DROP TABLE IF EXISTS notification_templates CASCADE;

DROP TABLE IF EXISTS notification_preferences CASCADE;