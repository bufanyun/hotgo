-- 删除表
DROP TABLE IF EXISTS hg_banner;

-- 删除 flashbanner 相关菜单
-- 先删除子菜单（权限按钮）
DELETE FROM hg_admin_menu WHERE name IN ('addbanner', 'editbanner', 'delbanner');

-- 删除主菜单
DELETE FROM hg_admin_menu WHERE name = 'flashbanner';