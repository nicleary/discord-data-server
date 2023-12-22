-- Create "messages" table
CREATE TABLE `messages` (`id` bigint NOT NULL AUTO_INCREMENT, `contents` varchar(8192) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
