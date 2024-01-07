-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `is_bot` bool NOT NULL DEFAULT 0;
