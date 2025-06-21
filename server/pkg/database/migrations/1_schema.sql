-- +goose Up
CREATE TABLE IF NOT EXISTS `rooms` (
	`id` uuid NOT NULL,
	`place_max` INT,
	`center_point` VARCHAR(255) NOT NULL,
	`radius` FLOAT NOT NULL,
	`create_at` TIMESTAMP NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `people` (
	`id` uuid NOT NULL,
	`room_id` uuid NOT NULL,
	`user_name` VARCHAR(255),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `places` (
	`id` uuid NOT NULL,
	`room_id` uuid NOT NULL,
	`place_id` VARCHAR(255),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `votes` (
	`id` uuid NOT NULL,
	`people_id` uuid NOT NULL,
	`place_id` uuid NOT NULL,
	`rank` INT NOT NULL,
	`time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`people_id`) REFERENCES `people`(`id`) ON DELETE CASCADE,
	FOREIGN KEY (`place_id`) REFERENCES `places`(`id`) ON DELETE CASCADE
);