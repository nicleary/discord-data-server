-- Create "messages" table
CREATE TABLE `messages` (`id` bigint NOT NULL AUTO_INCREMENT, `contents` varchar(10) NULL, `something` varchar(255) NOT NULL, `something else` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
