-- Modify "messages" table
ALTER TABLE `messages` ADD COLUMN `created_at` timestamp NULL, ADD COLUMN `updated_at` timestamp NULL;
-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `created_at` timestamp NULL, ADD COLUMN `updated_at` timestamp NULL;
