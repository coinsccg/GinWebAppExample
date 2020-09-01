DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20)                             NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20)                             NOT NULL,
    `username`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       varchar(64) COLLATE utf8mb4_general_ci,
    `gender`      tinyint(4)                             NOT NULL DEFAULT '0',
    `create_time` timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE
        CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;


CREATE TABLE `question`
(
    `id`          bigint(20)                               NOT NULL AUTO_INCREMENT,
    `question_id` bigint(20)                               NOT NULL,
    `caption`     varchar(128) COLLATE utf8mb4_general_ci  NOT NULL,
    `content`     varchar(8192) COLLATE utf8mb4_general_ci NOT NULL,
    `author_id`   bigint(20)                               NOT NULL,
    `category_id` bigint(20)                               NOT NULL,
    `status`      tinyint(4)                               NOT NULL DEFAULT '1',
    `create_time` timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_question_id` (`question_id`),
    KEY `idx_category_id` (`category_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;


CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned NOT NULL,
  `category_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_category_id` (`category_id`),
  UNIQUE KEY `idx_category_name` (`category_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



