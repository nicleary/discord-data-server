-- Modify "messages" table
ALTER TABLE `messages` ADD COLUMN `channel_id` varchar(255) NOT NULL, ADD INDEX `message_channel_id` (`channel_id`);
