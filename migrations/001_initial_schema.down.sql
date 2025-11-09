-- Drop indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_messages_created_at;
DROP INDEX IF EXISTS idx_messages_chat_id;
DROP INDEX IF EXISTS idx_chats_updated_at;
DROP INDEX IF EXISTS idx_chats_user_id;

-- Drop tables (in reverse order due to foreign key constraints)
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS users;

-- Drop UUID extension (only if no other objects depend on it)
-- DROP EXTENSION IF EXISTS "uuid-ossp";