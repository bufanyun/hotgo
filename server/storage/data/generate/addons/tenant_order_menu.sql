-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：/Users/mengshuai/Desktop/gosrc/hotgo_dev/server/storage/data/generate/addons/tenant_order_menu.sql
-- Version: 2.13.1
-- Date: 2024-04-13 23:37:27
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
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2228', '多租户功能演示', 'tenantOrder', '/tenantOrder', '', '1', '/addons/hgexample/tenantOrder/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '3', 'tr_2227 tr_2228 ', '200', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, '多租户功能演示列表', 'tenantOrderIndex', 'index', '', '2', '', '/hgexample/tenantOrder/list', '', '/addons/hgexample/tenantOrder/index', '1', 'tenantOrder', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2227 tr_2228 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '多租户功能演示详情', 'tenantOrderView', '', '', '3', '', '/hgexample/tenantOrder/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2227 tr_2228 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增多租户功能演示', 'tenantOrderEdit', '', '', '3', '', '/hgexample/tenantOrder/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2227 tr_2228 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除多租户功能演示', 'tenantOrderDelete', '', '', '3', '', '/hgexample/tenantOrder/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2227 tr_2228 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);




-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出多租户功能演示', 'tenantOrderExport', '', '', '3', '', '/hgexample/tenantOrder/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2227 tr_2228 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


COMMIT;