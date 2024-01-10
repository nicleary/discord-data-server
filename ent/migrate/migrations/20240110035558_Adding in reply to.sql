-- Modify "messages" table
ALTER TABLE `messages` ADD COLUMN `in_reply_to_id` bigint NULL, ADD INDEX `messages_messages_responders` (`in_reply_to_id`), ADD CONSTRAINT `messages_messages_responders` FOREIGN KEY (`in_reply_to_id`) REFERENCES `messages` (`id`) ON UPDATE RESTRICT ON DELETE SET NULL;
