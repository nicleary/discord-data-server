-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` varchar(64) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "messages" table
CREATE TABLE `messages` (`id` bigint NOT NULL AUTO_INCREMENT, `contents` varchar(8192) NOT NULL, `sent_at` timestamp NULL, `sender_id` bigint NOT NULL, PRIMARY KEY (`id`), INDEX `messages_users_messages` (`sender_id`), CONSTRAINT `messages_users_messages` FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
