-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` varchar(64) NOT NULL, `date_joined` timestamp NULL, PRIMARY KEY (`id`), UNIQUE INDEX `user_id` (`user_id`), INDEX `user_user_id` (`user_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "messages" table
CREATE TABLE `messages` (`id` bigint NOT NULL AUTO_INCREMENT, `contents` varchar(8192) NOT NULL, `sent_at` timestamp NULL, `sender_id` bigint NOT NULL, PRIMARY KEY (`id`), INDEX `message_sender_id` (`sender_id`), INDEX `message_sent_at` (`sent_at`), CONSTRAINT `messages_users_messages` FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
