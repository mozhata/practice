CREATE TABLE IF NOT EXISTS `operation_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operator_id` varchar(36) NOT NULL COMMENT '操作人ID',
  `customer_id` varchar(36) NOT NULL COMMENT '客户ID',
  `op_type` tinyint(4) NOT NULL COMMENT '操作类型',
  `amount` decimal(10,0) NOT NULL COMMENT '金额(元)',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `notes` text NOT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `operator_id` (`operator_id`),
  KEY `customer_id` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4