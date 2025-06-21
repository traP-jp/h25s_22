-- +goose Up
CREATE TABLE IF NOT EXISTS `rooms` (
	`id` CHAR(36) NOT NULL,
	`place_max` INT,
	`center_point` VARCHAR(255) NOT NULL,
	`radius` FLOAT NOT NULL,
	`create_at` TIMESTAMP NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `users` (
	`id` CHAR(36) NOT NULL,
	`room_id` CHAR(36) NOT NULL,
	`name` VARCHAR(255),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `places` (
	`id` CHAR(36) NOT NULL,
	`room_id` CHAR(36) NOT NULL,
	`google_place_id` VARCHAR(255),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `placeVotes` (
	`id` CHAR(36) NOT NULL,
	`user_id` CHAR(36) NOT NULL,
	`place_id` CHAR(36) NOT NULL,
	`rank` INT NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
	FOREIGN KEY (`place_id`) REFERENCES `places`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `timeOptions` (
	`id` CHAR(36) NOT NULL,
	`start_time` TIMESTAMP NULL,
	`end_time` TIMESTAMP NULL,
	PRIMARY KEY (`id`)
);	

CREATE TABLE IF NOT EXISTS `timeVotes` (
	`id` CHAR(36) NOT NULL,
	`users_id` CHAR(36) NOT NULL,
	`time_id` CHAR(36) NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`users_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
	FOREIGN KEY (`time_id`) REFERENCES `timeOptions`(`id`) ON DELETE CASCADE
);