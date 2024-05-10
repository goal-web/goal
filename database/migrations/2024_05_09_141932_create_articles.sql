CREATE TABLE IF NOT EXISTS articles
(
    `id`       INT UNSIGNED AUTO_INCREMENT,
    title      varchar(20) not null,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;