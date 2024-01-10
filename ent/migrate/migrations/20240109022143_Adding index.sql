-- Modify "messages" table
ALTER TABLE `messages` ADD UNIQUE INDEX `message_id` (`message_id`), ADD INDEX `message_message_id` (`message_id`);
