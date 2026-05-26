-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：C:\Users\Administrator\Desktop\gosrc\hotgo_dev\server\storage\data\generate\normal_tree_demo_menu.sql
-- Version: 2.13.1
-- Date: 2024-04-09 17:05:12
-- Link https://github.com/bufanyun/hotgo

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `hotgo`
--

-- --------------------------------------------------------

--
-- 插入表中的数据 `hg_admin_menu`
--


SET @now := now();


-- 菜单目录
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2366', '普通树表', 'normalTreeDemo', '/normalTreeDemo', '', '1', '/develop/generateDemo/normalTreeDemo/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '3', 'tr_2097 tr_2366 ', '200', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, '普通树表列表', 'normalTreeDemoIndex', 'index', '', '2', '', '/normalTreeDemo/list', '', '/normalTreeDemo/index', '1', 'normalTreeDemo', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2097 tr_2366 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '普通树表详情', 'normalTreeDemoView', '', '', '3', '', '/normalTreeDemo/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增普通树表', 'normalTreeDemoEdit', '', '', '3', '', '/normalTreeDemo/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();

-- 获取最大排序
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @editId, '获取普通树表最大排序', 'normalTreeDemoMaxSort', '', '', '3', '', '/normalTreeDemo/maxSort', '', '', '1', '', '0', '0', '', '0', '0', '0', '6', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId, ' tr_', @editId,' '), '30', '', '1', @now, @now);


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除普通树表', 'normalTreeDemoDelete', '', '', '3', '', '/normalTreeDemo/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);


-- 更新状态
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '修改普通树表状态', 'normalTreeDemoStatus', '', '', '3', '', '/normalTreeDemo/status', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '50', '', '1', @now, @now);



-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出普通树表', 'normalTreeDemoExport', '', '', '3', '', '/normalTreeDemo/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


-- 关系树选项
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '获取普通树表关系树选项', 'normalTreeDemoTreeOption', '', '', '3', '', '/normalTreeDemo/treeOption', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);

COMMIT;