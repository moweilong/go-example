CREATE DATABASE IF NOT EXISTS `hotgo` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `hotgo`;

CREATE TABLE IF NOT EXISTS `hg_admin_member` (
  `id` bigint AUTO_INCREMENT NOT NULL COMMENT '管理员ID',
  `dept_id` bigint DEFAULT '0' COMMENT '部门ID',
  `role_id` bigint DEFAULT '10' COMMENT '角色ID',
  `real_name` varchar(32) DEFAULT '' COMMENT '真实姓名',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '帐号',
  `password_hash` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` char(16) NOT NULL COMMENT '密码盐',
  `password_reset_token` varchar(150) DEFAULT '' COMMENT '密码重置令牌',
  `integral` decimal(10,2) DEFAULT '0.00' COMMENT '积分',
  `balance` decimal(10,2) DEFAULT '0.00' COMMENT '余额',
  `avatar` char(150) DEFAULT '' COMMENT '头像',
  `sex` tinyint DEFAULT '1' COMMENT '性别',
  `qq` varchar(20) DEFAULT '' COMMENT 'qq',
  `email` varchar(60) DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机号码',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `city_id` bigint DEFAULT '0' COMMENT '城市编码',
  `address` varchar(100) DEFAULT '' COMMENT '联系地址',
  `pid` bigint NOT NULL COMMENT '上级管理员ID',
  `level` int DEFAULT '1' COMMENT '关系树等级',
  `tree` varchar(512) NOT NULL COMMENT '关系树',
  `invite_code` varchar(12) DEFAULT NULL COMMENT '邀请码',
  `cash` json DEFAULT NULL COMMENT '提现配置',
  `last_active_at` datetime DEFAULT NULL COMMENT '最后活跃时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`)
) COMMENT='管理员_用户表';

INSERT INTO `hg_admin_member` (`id`, `dept_id`, `role_id`, `real_name`, `username`, `password_hash`, `salt`, `password_reset_token`, `integral`, `balance`, `avatar`, `sex`, `qq`, `email`, `mobile`, `birthday`, `city_id`, `address`, `pid`, `level`, `tree`, `invite_code`, `cash`, `last_active_at`, `remark`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'admin', 'admin', 'a7c588fffeb2c1d99b29879d7fe97c78', '6541561', '', '88.00', '99289.78', 'http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8er9nfkchdopav.png', 1, '133814250', '133814250@qq.com', '15303830571', '2016-04-16', 410172, '莲花街001号', 0, 1, '', '111', '{"name": "admin", "account": "15303830571", "payeeCode": "http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8mqal5isvcb58g.jpg"}', '2024-08-27 19:02:49', NULL, 1, '2021-02-12 17:59:45', '2024-08-27 19:02:49');