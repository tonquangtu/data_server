DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `username`          varchar(64)         NOT NULL COMMENT 'Username',
    `password_hash`     varchar(255)        NOT NULL COMMENT 'Hash of password',
    `created_at`        bigint(20) unsigned NOT NULL COMMENT 'Created at',
    `updated_at`        bigint(20) unsigned NOT NULL COMMENT 'Updated at',
    `deleted_at`        bigint(20) unsigned NOT NULL COMMENT 'Deleted at',
    PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `device`;
CREATE TABLE `device`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id`           bigint(20) unsigned       NOT NULL COMMENT 'User id',
    `device_name`         varchar(255)        NOT NULL COMMENT 'Device name',
    `device_detail`     varchar(512) NOT NULL COMMENT 'Device detail',
    `created_at`        bigint(20) unsigned NOT NULL COMMENT 'Created at',
    `updated_at`        bigint(20) unsigned NOT NULL COMMENT 'Updated at',
    `deleted_at`        bigint(20) unsigned NOT NULL COMMENT 'Deleted at',
    PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `note`;
CREATE TABLE `note`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id`           bigint(20) unsigned         NOT NULL COMMENT 'User id',
    `device_id`         bigint(20) unsigned        NOT NULL COMMENT 'Device id',
    `data`              text NOT NULL COMMENT 'Device detail',
    `created_at`        bigint(20) unsigned NOT NULL COMMENT 'Created at',
    `updated_at`        bigint(20) unsigned NOT NULL COMMENT 'Updated at',
    `deleted_at`        bigint(20) unsigned NOT NULL COMMENT 'Deleted at',
    PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8;