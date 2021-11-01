CREATE TABLE IF NOT EXISTS `users`(
	`id` VARCHAR(36) NOT NULL,
	`username` VARCHAR(30) NOT NULL,
	`email` VARCHAR(50) NOT NULL,
	`hashed_pass` VARCHAR(255) NOT NULL,
	`hash_key` VARCHAR(255) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	PRIMARY KEY(`id`),
	UNIQUE KEY `unique_username` (`username`),
	UNIQUE KEY `unique_email` (`email`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci