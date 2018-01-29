CREATE DATABASE IF NOT EXISTS `block_chain`
USE `block_chain`

DROP TABLE IF EXISTS `block`;
CREATE TABLE `block` (
    `id` VARCHAR(36) NOT NULL COMMENT 'uuid',
    `Data` TEXT NOT NULL COMMENT '数据',
    `PrevBlockHash` VARCHAR(36) NOT NULL COMMENT '上一块的hash',
    `Hash` VARCHAR(36) NOT NULL COMMENT 'hash',
    `Timestamp` int(64) NOT NULL COMMENT '时间戳',
    `Nonce` int(64) NOT NULL COMMENT '唯一值',
    PRIMARY KEY(`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='块信息';
