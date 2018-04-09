CREATE TABLE IF NOT EXISTS `consumer.sql`(
  `id` varchar(32) NOT NULL COMMENT "id",
  `phone` varchar(100) COMMENT "phone" DEFAULT NULL,
  `email` varchar(256) COMMENT "email" DEFAULT NULL,
  `addr` text NOT NULL COMMENT "地址",
  `notes` text NOT NULL COMMENT "备注",
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "修改时间",
  `name` varchar(256) NOT NULL DEFAULT "" COMMENT "姓名",
  PRIMARY KEY (`id`),
  UNIQUE KEY (`phone`),
  UNIQUE KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;