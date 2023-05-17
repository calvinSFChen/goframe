CREATE TABLE `test_system_admin` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '后台管理员账号',
  `head_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '管理员头像',
  `password` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '后台管理员密码',
  `real_name` varchar(16) NOT NULL DEFAULT '' COMMENT '后台管理员姓名',
  `roles` varchar(128) NOT NULL DEFAULT '' COMMENT '后台管理员权限',
  `last_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '后台管理员最后一次登录ip',
  `last_time` int unsigned NOT NULL DEFAULT '0' COMMENT '后台管理员最后一次登录时间',
  `login_count` int unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `level` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '后台管理员级别',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '后台管理员状态 1有效0无效',
  `division_id` int NOT NULL DEFAULT '0' COMMENT '事业部id',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `operator` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人',
  `phone` decimal(11,0) DEFAULT NULL COMMENT '手机号码',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `account` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 COMMENT='后台管理员表';

CREATE TABLE `test_system_menus` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` smallint unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `icon` varchar(16) NOT NULL DEFAULT '' COMMENT '图标',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '名称',
  `sort` tinyint NOT NULL DEFAULT '1' COMMENT '排序',
  `is_show` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否为隐藏菜单0=隐藏菜单,1=显示菜单',
  `access` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '子管理员是否可用',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '路径',
  `unique_auth` varchar(150) NOT NULL DEFAULT '' COMMENT '前台唯一标识',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `operator` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1  COMMENT='菜单表';

CREATE TABLE `test_system_route` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `pid` smallint unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '名称',
  `methods` varchar(10) NOT NULL DEFAULT '' COMMENT '提交方式POST GET PUT DELETE',
  `access` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '子管理员是否可用',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '路径',
  `unique_auth` varchar(150) NOT NULL DEFAULT '' COMMENT '前台唯一标识',
  `is_show` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '0关闭1开启',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `operator` varchar(32) NOT NULL DEFAULT '' COMMENT '操作人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB COMMENT='菜单表';


