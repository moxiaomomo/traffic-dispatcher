create database `traffic-dispatcher` default character set utf8;
-- grant all privileges on `traffic-dispatcher`.* to 'admin'@'%' identified by 'test123456';
-- flush privileges;
use traffic-dispatcher;

-- 创建用户表
CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role` int(11) DEFAULT 0 COMMENT '用户角色,0:passenger,1:driver,2:admin,...',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户唯一id',
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT '用户密码(加密)',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `phone` varchar(128) DEFAULT '' COMMENT '手机号',
  `email_validated` tinyint(1) DEFAULT 0 COMMENT '邮箱是否已验证',
  `phone_validated` tinyint(1) DEFAULT 0 COMMENT '手机号是否已验证',
  `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
  `profile` text COMMENT '用户属性',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态(启用/禁用/锁定/标记删除等)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_userid` (`user_id`),
  UNIQUE KEY `idx_username_role` (`user_name`,`role`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建订单表
CREATE TABLE `tbl_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单唯一id',
  `src_geo` varchar(32) NOT NULL DEFAULT '' COMMENT '起始位置',
  `dest_geo` varchar(32) NOT NULL DEFAULT '' COMMENT '目标位置',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '订单创建日期',
  `accept_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '订单accept日期',
  `cancel_at` datetime COMMENT '订单取消日期',
  `finish_at` datetime COMMENT '订单完成日期',
  `cancel_role` int(11) DEFAULT -1 COMMENT '取消订单的角色,0:passenger,1:driver',
  `cost` float DEFAULT 0 COMMENT '订单价格',
  `passenger_id` varchar(64) NOT NULL DEFAULT '' COMMENT '乘客id',
  `driver_id` varchar(64) NOT NULL DEFAULT '' COMMENT '司机id',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '订单状态(0:创建/1:进行中/2:取消/3:完成等)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_orderid` (`order_id`),
  KEY `idx_passenger` (`passenger_id`),
  KEY `idx_driver` (`driver_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
