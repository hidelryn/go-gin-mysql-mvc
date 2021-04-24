CREATE TABLE `user`.`user` (
  `id` INT UNSIGNED NOT NULL,
  `nickname` VARCHAR(45) NULL,
  `create_at` BIGINT UNSIGNED NULL DEFAULT 0,
  `update_at` BIGINT UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `nickname_UNIQUE` (`nickname` ASC) VISIBLE);

CREATE TABLE `user`.`dayJoinCnt` (
  `yyyymmdd` INT UNSIGNED NOT NULL,
  `cnt` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`yyyymmdd`));

