DROP TABLE IF EXISTS `hg_banner`;
CREATE TABLE `hg_banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '轮播图名称',
  `cover` varchar(255) DEFAULT NULL COMMENT '图片URL',
  `link` varchar(255) DEFAULT NULL COMMENT '跳转链接，小程序内用相对地址',
  `type` int(11) DEFAULT 0 COMMENT '类型默认不传',
  `status` tinyint(1) DEFAULT 1 COMMENT '1可用,2不可用',
  `sort` int(11) DEFAULT 0 COMMENT '排序，数字越大越靠前',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='轮播图表';


-- 添加菜单项
-- 先查找或创建"内容管理"父菜单（如果不存在）
INSERT IGNORE INTO `hg_admin_menu`(`pid`, `level`, `tree`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `sort`, `remark`, `status`, `updated_at`, `created_at`) VALUES 
(0, 1, '', '内容管理', 'content', '/content', 'BookOutlined', 1, '/content', '', '', 'LAYOUT', 2, '', 1, 2, '', 2, 2, 2, 12, '内容管理模块', 1, NOW(), NOW());

-- 获取内容管理菜单ID并添加轮播图管理菜单
SET @content_parent_id = (SELECT id FROM `hg_admin_menu` WHERE `name` = 'content' LIMIT 1);

INSERT INTO `hg_admin_menu`(`pid`, `level`, `tree`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `sort`, `remark`, `status`, `updated_at`, `created_at`) VALUES 
(@content_parent_id, 2, CONCAT('tr_', @content_parent_id, ' '), '轮播图管理', 'flashbanner', '/flashbanner', '', 2, '', '/flashbanner/banner/list', '', '/addons/flashbanner/index', 2, '', 2, 2, '', 1, 2, 2, 10, '轮播图管理模块', 1, NOW(), NOW());

-- 添加子菜单和按钮
SET @flashbanner_menu_id = LAST_INSERT_ID();

INSERT INTO `hg_admin_menu`(`pid`, `level`, `tree`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `sort`, `remark`, `status`, `updated_at`, `created_at`) VALUES 
(@flashbanner_menu_id, 3, CONCAT('tr_', @content_parent_id, ' tr_', @flashbanner_menu_id, ' '), '新增轮播', 'addbanner', '', '', 3, '', '/flashbanner/banner/create', '', '', 2, '', 2, 2, '', 2, 2, 2, 10, '新增轮播图权限', 1, NOW(), NOW()),
(@flashbanner_menu_id, 3, CONCAT('tr_', @content_parent_id, ' tr_', @flashbanner_menu_id, ' '), '轮播编辑', 'editbanner', '', '', 3, '', '/flashbanner/banner/update', '', '', 2, '', 2, 2, '', 2, 2, 2, 10, '编辑轮播图权限', 1, NOW(), NOW()),
(@flashbanner_menu_id, 3, CONCAT('tr_', @content_parent_id, ' tr_', @flashbanner_menu_id, ' '), '删除轮播', 'delbanner', '', '', 3, '', '/flashbanner/banner/delete', '', '', 2, '', 2, 2, '', 2, 2, 2, 10, '删除轮播图权限', 1, NOW(), NOW());