DROP TABLE IF EXISTS `registered_book`;

CREATE TABLE IF NOT EXISTS `registered_book`
(
  `id`               INT             AUTO_INCREMENT,
  `name`             VARCHAR         NOT NULL,
  `isbn`             INT(13)         NOT NULL,
  `image_url`        TEXT            NOT NULL,
  `created_at`       TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at`       TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


DROP TABLE IF EXISTS `tag`;

CREATE TABLE IF NOT EXISTS `tag`
(
  `id`               INT             AUTO_INCREMENT,
  `name`             VARCHAR         NOT NULL,
  `created_at`       TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at`       TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


DROP TABLE IF EXISTS `tag_book`;

CREATE TABLE IF NOT EXISTS `tag_book`
(
  `id`               INT             AUTO_INCREMENT,
  `tag_id`           INT             NOT NULL,
  `book_id`          INT             NOT NULL,
  `created_at`       TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at`       TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY(`tag_id`)  REFERENCES tag(id),
  FOREIGN KEY(`book_id`) REFERENCES book(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
