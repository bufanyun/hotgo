-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和在SQL文件：C:\Users\Administrator\Desktop\gosrc\hotgo_dev\server\storage\data\generate\curd_demo_menu.sql
-- Version: 2.11.5
-- Date: 2023-12-29 18:18:21
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
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2097', '生成演示', 'curdDemo', '/curdDemo', 'MenuOutlined', '1', '', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '2', 'tr_2097 ', '200', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, '生成演示列表', 'curdDemoIndex', 'index', '', '2', '', '/curdDemo/list', '', '/curdDemo/index', '1', '', '0', '0', '', '0', '0', '0', '3', CONCAT('tr_2097 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '生成演示详情', 'curdDemoView', '', '', '3', '', '/curdDemo/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增生成演示', 'curdDemoEdit', '', '', '3', '', '/curdDemo/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();

-- 获取最大排序
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @editId, '获取生成演示最大排序', 'curdDemoMaxSort', '', '', '3', '', '/curdDemo/maxSort', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId, ' tr_', @editId,' '), '30', '', '1', @now, @now);


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除生成演示', 'curdDemoDelete', '', '', '3', '', '/curdDemo/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);


-- 更新状态
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '修改生成演示状态', 'curdDemoStatus', '', '', '3', '', '/curdDemo/status', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId,' '), '50', '', '1', @now, @now);


-- 操作开关
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '操作生成演示开关', 'curdDemoSwitch', '', '', '3', '', '/curdDemo/switch', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId,' '), '60', '', '1', @now, @now);


-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出生成演示', 'curdDemoExport', '', '', '3', '', '/curdDemo/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2097 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


COMMIT;