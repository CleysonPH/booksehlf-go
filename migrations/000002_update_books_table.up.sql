ALTER TABLE `books` ADD COLUMN `description` TEXT NULL AFTER `cover`;
ALTER TABLE `books` ADD COLUMN `published_at` DATE NULL AFTER `description`;
ALTER TABLE `books` ADD COLUMN `publisher` VARCHAR(255) NULL AFTER `published_at`;
ALTER TABLE `books` ADD COLUMN `pages` INT NULL AFTER `publisher`;
ALTER TABLE `books` ADD COLUMN `edition` INT NULL AFTER `pages`;

