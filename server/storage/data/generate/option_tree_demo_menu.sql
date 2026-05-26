-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：C:\Users\Administrator\Desktop\gosrc\hotgo_dev\server\storage\data\generate\option_tree_demo_menu.sql
-- Version: 2.13.1
-- Date: 2024-04-09 17:22:28
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
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2366', '选项树表', 'optionTreeDemo', '/optionTreeDemo', '', '1', '/develop/generateDemo/optionTreeDemo/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '3', 'tr_2097 tr_2366 ', '300', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, '选项树表列表', 'optionTreeDemoIndex', 'index', '', '2', '', '/optionTreeDemo/list', '', '/optionTreeDemo/index', '1', 'optionTreeDemo', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2097 tr_2366 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '选项树表详情', 'optionTreeDemoView', '', '', '3', '', '/optionTreeDemo/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增选项树表', 'optionTreeDemoEdit', '', '', '3', '', '/optionTreeDemo/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();

-- 获取最大排序
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @editId, '获取选项树表最大排序', 'optionTreeDemoMaxSort', '', '', '3', '', '/optionTreeDemo/maxSort', '', '', '1', '', '0', '0', '', '0', '0', '0', '6', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId, ' tr_', @editId,' '), '30', '', '1', @now, @now);


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除选项树表', 'optionTreeDemoDelete', '', '', '3', '', '/optionTreeDemo/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);


-- 更新状态
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '修改选项树表状态', 'optionTreeDemoStatus', '', '', '3', '', '/optionTreeDemo/status', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '50', '', '1', @now, @now);



-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出选项树表', 'optionTreeDemoExport', '', '', '3', '', '/optionTreeDemo/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


-- 关系树选项
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '获取选项树表关系树选项', 'optionTreeDemoTreeOption', '', '', '3', '', '/optionTreeDemo/treeOption', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2097 tr_2366 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);

COMMIT;