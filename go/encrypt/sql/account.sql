CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` char(36) NOT NULL,
  `name` varchar(50) NOT NULL,
  `phone` varchar(16) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `bio` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  UNIQUE KEY `phone_email` (`phone`,`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;