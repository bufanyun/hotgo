-- 删除表（如果存在）
DROP TABLE IF EXISTS hg_banner;

-- 创建表
CREATE TABLE hg_banner (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  cover VARCHAR(255) DEFAULT NULL,
  link VARCHAR(255) DEFAULT NULL,
  type INTEGER DEFAULT 0,
  status SMALLINT DEFAULT 1,
  sort INTEGER DEFAULT 0,
  created_at TIMESTAMP DEFAULT NULL,
  updated_at TIMESTAMP DEFAULT NULL
);

-- 添加表注释
COMMENT ON TABLE hg_banner IS '轮播图表';

-- 添加列注释
COMMENT ON COLUMN hg_banner.name IS '轮播图名称';
COMMENT ON COLUMN hg_banner.cover IS '图片URL';
COMMENT ON COLUMN hg_banner.link IS '跳转链接，小程序内用相对地址';
COMMENT ON COLUMN hg_banner.type IS '类型默认不传';
COMMENT ON COLUMN hg_banner.status IS '1可用,2不可用';
COMMENT ON COLUMN hg_banner.sort IS '排序，数字越大越靠前';

-- 添加 updated_at 字段的更新触发器（模拟 MySQL 的 ON UPDATE CURRENT_TIMESTAMP）
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
    BEFORE UPDATE ON hg_banner
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 添加菜单项
-- 先查找或创建"内容管理"父菜单（如果不存在）
INSERT INTO hg_admin_menu (pid, level, tree, title, name, path, icon, type, redirect, permissions, permission_name, component, always_show, active_menu, is_root, is_frame, frame_src, keep_alive, hidden, affix, sort, remark, status, updated_at, created_at)
SELECT 0, 1, '', '内容管理', 'content', '/content', 'BookOutlined', 1, '/content', '', '', 'LAYOUT', 2, '', 1, 2, '', 2, 2, 2, 12, '内容管理模块', 1, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM hg_admin_menu WHERE name = 'content');

-- 获取内容管理菜单ID并添加轮播图管理菜单
WITH content_parent AS (
    SELECT id FROM hg_admin_menu WHERE name = 'content' LIMIT 1
),
inserted_menu AS (
    INSERT INTO hg_admin_menu (pid, level, tree, title, name, path, icon, type, redirect, permissions, permission_name, component, always_show, active_menu, is_root, is_frame, frame_src, keep_alive, hidden, affix, sort, remark, status, updated_at, created_at)
    SELECT 
        cp.id, 
        2, 
        'tr_' || cp.id::text || ' ', 
        '轮播图管理', 
        'flashbanner', 
        '/flashbanner', 
        '', 
        2, 
        '', 
        '/flashbanner/banner/list', 
        '', 
        '/addons/flashbanner/index', 
        2, 
        '', 
        2, 
        2, 
        '', 
        1, 
        2, 
        2, 
        10, 
        '轮播图管理模块', 
        1, 
        NOW(), 
        NOW()
    FROM content_parent cp
    RETURNING id
)
-- 添加子菜单和按钮
INSERT INTO hg_admin_menu (pid, level, tree, title, name, path, icon, type, redirect, permissions, permission_name, component, always_show, active_menu, is_root, is_frame, frame_src, keep_alive, hidden, affix, sort, remark, status, updated_at, created_at)
SELECT 
    im.id,
    3,
    'tr_' || (SELECT id FROM content_parent)::text || ' tr_' || im.id::text || ' ',
    '新增轮播',
    'addbanner',
    '',
    '',
    3,
    '',
    '/flashbanner/banner/create',
    '',
    '',
    2,
    '',
    2,
    2,
    '',
    2,
    2,
    2,
    10,
    '新增轮播图权限',
    1,
    NOW(),
    NOW()
FROM inserted_menu im
UNION ALL
SELECT 
    im.id,
    3,
    'tr_' || (SELECT id FROM content_parent)::text || ' tr_' || im.id::text || ' ',
    '轮播编辑',
    'editbanner',
    '',
    '',
    3,
    '',
    '/flashbanner/banner/update',
    '',
    '',
    2,
    '',
    2,
    2,
    '',
    2,
    2,
    2,
    10,
    '编辑轮播图权限',
    1,
    NOW(),
    NOW()
FROM inserted_menu im
UNION ALL
SELECT 
    im.id,
    3,
    'tr_' || (SELECT id FROM content_parent)::text || ' tr_' || im.id::text || ' ',
    '删除轮播',
    'delbanner',
    '',
    '',
    3,
    '',
    '/flashbanner/banner/delete',
    '',
    '',
    2,
    '',
    2,
    2,
    '',
    2,
    2,
    2,
    10,
    '删除轮播图权限',
    1,
    NOW(),
    NOW()
FROM inserted_menu im;