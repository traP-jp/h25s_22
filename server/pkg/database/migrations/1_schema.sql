-- +goose Up
CREATE TABLE IF NOT EXISTS `rooms` (
	`id` uuid NOT NULL,
	`placeMax` INT,
	`centerPoint` VARCHAR(255) NOT NULL,
	`radius` FLOAT NOT NULL,
	`create_at` TIMESTAMP NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `people` (
	`id` uuid NOT NULL,
	`roomID` uuid NOT NULL,
	`userName` VARCHAR(255),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`roomID`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `places` (
	`id` uuid NOT NULL,
	`roomID` uuid NOT NULL,
	`placeID` VARCHAR(255),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`roomID`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `votes` (
	`id` uuid NOT NULL,
	`peopleID` uuid NOT NULL,
	`placeID` uuid NOT NULL,
	`rank` INT NOT NULL,
	`time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`peopleID`) REFERENCES `people`(`id`) ON DELETE CASCADE,
	FOREIGN KEY (`placeID`) REFERENCES `places`(`id`) ON DELETE CASCADE
);