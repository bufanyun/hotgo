CREATE TABLE `hg_addon_hgexample_table` (                 -- 插件_案例_表格
`id` INTEGER NOT NULL  ,                                  -- ID
`category_id` INTEGER NOT NULL ,                          -- 分类ID
`flag` TEXT   DEFAULT NULL ,                              -- 标签
`title` TEXT NOT NULL ,                                   -- 标题
`description` TEXT NOT NULL ,                             -- 描述
`content` TEXT DEFAULT NULL ,                             -- 内容
`image` TEXT DEFAULT NULL ,                               -- 单图
`images` TEXT   DEFAULT NULL ,                            -- 多图
`attachfile` TEXT DEFAULT NULL ,                          -- 附件
`attachfiles` TEXT   DEFAULT NULL ,                       -- 多附件
`map` TEXT   DEFAULT NULL ,                               -- 动态键值对
`star` decimal(5,1) DEFAULT 0.0 ,                         -- 推荐星
`price` decimal(10,2) NOT NULL DEFAULT 0.00 ,             -- 价格
`views` INTEGER DEFAULT NULL ,                            -- 浏览次数
`activity_at` date DEFAULT NULL ,                         -- 活动时间
`start_at` datetime DEFAULT NULL ,                        -- 开启时间
`end_at` datetime DEFAULT NULL ,                          -- 结束时间
`switch` INTEGER DEFAULT NULL ,                           -- 开关
`sort` INTEGER DEFAULT NULL ,                             -- 排序
`avatar` TEXT DEFAULT '' ,                                -- 头像
`sex` INTEGER DEFAULT NULL ,                              -- 性别
`qq` TEXT DEFAULT '' ,                                    -- QQ
`email` TEXT DEFAULT '' ,                                 -- 邮箱
`mobile` TEXT DEFAULT '' ,                                -- 手机号码
`hobby` TEXT   DEFAULT NULL ,                             -- 爱好
`channel` INTEGER DEFAULT 1 ,                             -- 渠道
`city_id` INTEGER DEFAULT 0 ,                             -- 所在城市
`pid` INTEGER NOT NULL ,                                  -- 上级ID
`level` INTEGER DEFAULT 1 ,                               -- 树等级
`tree` TEXT DEFAULT NULL ,                                -- 关系树
`remark` TEXT DEFAULT NULL ,                              -- 备注
`status` INTEGER DEFAULT 1 ,                              -- 状态
`created_by` INTEGER DEFAULT 0 ,                          -- 创建者
`updated_by` INTEGER DEFAULT 0 ,                          -- 更新者
`created_at` datetime DEFAULT NULL ,                      -- 创建时间
`updated_at` datetime DEFAULT NULL ,                      -- 修改时间
`deleted_at` datetime DEFAULT NULL ,                      -- 删除时间
PRIMARY KEY (`id`)
);
CREATE TABLE `hg_addon_hgexample_tenant_order` (          -- 多租户_充值订单
  `id` INTEGER NOT NULL  ,                                -- ID
  `tenant_id` INTEGER DEFAULT NULL,                       -- 租户ID
  `merchant_id` INTEGER NOT NULL,                         -- 商户ID
  `user_id` INTEGER NOT NULL,                             -- 用户ID
  `product_name` TEXT DEFAULT NULL,                       -- 购买产品
  `order_sn` TEXT DEFAULT NULL,                           -- 订单号
  `money` decimal(10,2) NOT NULL,                         -- 充值金额
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 订单状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL,                     -- 修改时间
  PRIMARY KEY (`id`)
);
CREATE TABLE `hg_admin_cash` (                            -- 管理员_提现记录表
`id` INTEGER NOT NULL  ,                                  -- ID
`member_id` INTEGER NOT NULL ,                            -- 管理员ID
`money` decimal(10,2) NOT NULL ,                          -- 提现金额
`fee` decimal(10,2) NOT NULL ,                            -- 手续费
`last_money` decimal(10,2) NOT NULL ,                     -- 最终到账金额
`ip` TEXT NOT NULL ,                                      -- 申请人IP
`status` INTEGER NOT NULL ,                               -- 状态
`msg` TEXT NOT NULL ,                                     -- 处理结果
`handle_at` datetime DEFAULT NULL ,                       -- 处理时间
`created_at` datetime NOT NULL ,                          -- 申请时间
PRIMARY KEY (`id`)
);
CREATE TABLE `hg_admin_credits_log` (                     -- 管理员_资产变动表
  `id` INTEGER NOT NULL PRIMARY KEY,                      -- 变动ID
  `member_id` INTEGER DEFAULT 0,                          -- 管理员ID
  `app_id` TEXT DEFAULT NULL,                             -- 应用id
  `addons_name` TEXT NOT NULL DEFAULT '',                 -- 插件名称
  `credit_type` TEXT NOT NULL DEFAULT '',                 -- 变动类型
  `credit_group` TEXT DEFAULT NULL,                       -- 变动组别
  `before_num` decimal(10,2) DEFAULT 0.00,                -- 变动前
  `num` decimal(10,2) DEFAULT 0.00,                       -- 变动数据
  `after_num` decimal(10,2) DEFAULT 0.00,                 -- 变动后
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `ip` TEXT DEFAULT NULL,                                 -- 操作人IP
  `map_id` INTEGER DEFAULT 0,                             -- 关联ID
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_admin_dept` (                            -- 管理员_部门
  `id` INTEGER NOT NULL PRIMARY KEY,                      -- 部门ID
  `pid` INTEGER DEFAULT 0,                                -- 父部门ID
  `name` TEXT DEFAULT NULL,                               -- 部门名称
  `code` TEXT DEFAULT NULL,                               -- 部门编码
  `type` TEXT DEFAULT NULL,                               -- 部门类型
  `leader` TEXT DEFAULT NULL,                             -- 负责人
  `phone` TEXT DEFAULT NULL,                              -- 联系电话
  `email` TEXT DEFAULT NULL,                              -- 邮箱
  `level` INTEGER NOT NULL,                               -- 关系树等级
  `tree` TEXT DEFAULT NULL,                               -- 关系树
  `sort` INTEGER DEFAULT 0,                               -- 排序
  `status` INTEGER DEFAULT 1,                             -- 部门状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_admin_member` (                          -- 管理员_用户表
  `id` INTEGER NOT NULL PRIMARY KEY,                      -- 管理员ID
  `dept_id` INTEGER DEFAULT 0,                            -- 部门ID
  `role_id` INTEGER DEFAULT 10,                           -- 角色ID
  `real_name` TEXT DEFAULT '',                            -- 真实姓名
  `username` TEXT NOT NULL DEFAULT '',                    -- 帐号
  `password_hash` char(32) NOT NULL DEFAULT '',           -- 密码
  `salt` char(16) NOT NULL,                               -- 密码盐
  `password_reset_token` TEXT DEFAULT '',                 -- 密码重置令牌
  `integral` decimal(10,2)  DEFAULT 0.00,                 -- 积分
  `balance` decimal(10,2)  DEFAULT 0.00,                  -- 余额
  `avatar` char(150) DEFAULT '',                          -- 头像
  `sex` INTEGER DEFAULT 1,                                -- 性别
  `qq` TEXT DEFAULT '',                                   -- qq
  `email` TEXT DEFAULT '',                                -- 邮箱
  `mobile` TEXT DEFAULT '',                               -- 手机号码
  `birthday` date DEFAULT NULL,                           -- 生日
  `city_id` INTEGER DEFAULT 0,                            -- 城市编码
  `address` TEXT DEFAULT '',                              -- 联系地址
  `pid` INTEGER NOT NULL,                                 -- 上级管理员ID
  `level` INTEGER DEFAULT 1,                              -- 关系树等级
  `tree` TEXT NOT NULL,                                   -- 关系树
  `invite_code` TEXT DEFAULT NULL,                        -- 邀请码
  `cash` TEXT DEFAULT NULL,                               -- 提现配置
  `last_active_at` datetime DEFAULT NULL,                 -- 最后活跃时间
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_admin_member_post` (                     -- 管理员_用户岗位关联
  `member_id` INTEGER NOT NULL,                           -- 管理员ID
  `post_id` INTEGER NOT NULL,                             -- 岗位ID
  PRIMARY KEY (`member_id`, `post_id`)
);
CREATE TABLE `hg_admin_member_role` (                     -- 管理员_用户角色关联
  `member_id` INTEGER NOT NULL,                           -- 管理员ID
  `role_id` INTEGER NOT NULL,                             -- 角色ID
  PRIMARY KEY (`member_id`, `role_id`)
);
CREATE TABLE `hg_admin_menu` (                            -- 管理员_菜单权限
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 菜单ID
  `pid` INTEGER DEFAULT 0,                                -- 父菜单ID
  `level` INTEGER NOT NULL DEFAULT 1,                     -- 关系树等级
  `tree` TEXT NOT NULL,                                   -- 关系树
  `title` TEXT NOT NULL,                                  -- 菜单名称
  `name` TEXT NOT NULL,                                   -- 名称编码
  `path` TEXT DEFAULT NULL,                               -- 路由地址
  `icon` TEXT DEFAULT NULL,                               -- 菜单图标
  `type` INTEGER NOT NULL DEFAULT 1,                      -- 菜单类型（1目录 2菜单 3按钮）
  `redirect` TEXT DEFAULT NULL,                           -- 重定向地址
  `permissions` TEXT DEFAULT NULL,                        -- 菜单包含权限集合
  `permission_name` TEXT DEFAULT NULL,                    -- 权限名称
  `component` TEXT NOT NULL,                              -- 组件路径
  `always_show` INTEGER DEFAULT 0,                        -- 取消自动计算根路由模式
  `active_menu` TEXT DEFAULT NULL,                        -- 高亮菜单编码
  `is_root` INTEGER DEFAULT 0,                            -- 是否跟路由
  `is_frame` INTEGER DEFAULT 1,                           -- 是否内嵌
  `frame_src` TEXT DEFAULT NULL,                          -- 内联外部地址
  `keep_alive` INTEGER DEFAULT 0,                         -- 缓存该路由
  `hidden` INTEGER DEFAULT 0,                             -- 是否隐藏
  `affix` INTEGER DEFAULT 0,                              -- 是否固定
  `sort` INTEGER DEFAULT 0,                               -- 排序
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 菜单状态
  `updated_at` datetime DEFAULT NULL,                     -- 更新时间
  `created_at` datetime DEFAULT NULL                      -- 创建时间
);
CREATE TABLE `hg_admin_notice` (                          -- 管理员_通知公告
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 公告ID
  `title` TEXT NOT NULL,                                  -- 公告标题
  `type` INTEGER NOT NULL,                                -- 公告类型
  `tag` INTEGER DEFAULT NULL,                             -- 标签
  `content` TEXT NOT NULL,                                -- 公告内容
  `receiver` TEXT DEFAULT NULL,                           -- 接收者
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `sort` INTEGER NOT NULL DEFAULT 0,                      -- 排序
  `status` INTEGER DEFAULT 1,                             -- 公告状态
  `created_by` INTEGER DEFAULT NULL,                      -- 发送人
  `updated_by` INTEGER DEFAULT 0,                         -- 修改人
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL,                     -- 更新时间
  `deleted_at` datetime DEFAULT NULL                      -- 删除时间
);
CREATE TABLE `hg_admin_notice_read` (                     -- 管理员_公告已读记录
  `id` INTEGER NOT NULL ,                                 -- 记录ID
  `notice_id` INTEGER NOT NULL,                           -- 公告ID
  `member_id` INTEGER NOT NULL,                           -- 会员ID
  `clicks` INTEGER DEFAULT 1,                             -- 已读次数
  `updated_at` datetime DEFAULT NULL,                     -- 更新时间
  `created_at` datetime DEFAULT NULL,                     -- 阅读时间
PRIMARY KEY (`id`)
);
CREATE TABLE `hg_admin_oauth` (                           -- 管理员_第三方登录
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键
  `member_id` INTEGER DEFAULT 0,                          -- 用户ID
  `unionid` TEXT DEFAULT '',                              -- 唯一ID
  `oauth_client` TEXT DEFAULT NULL,                       -- 授权组别
  `oauth_openid` TEXT DEFAULT NULL,                       -- 授权开放ID
  `sex` INTEGER DEFAULT 1,                                -- 性别
  `nickname` TEXT DEFAULT NULL,                           -- 昵称
  `head_portrait` TEXT DEFAULT NULL,                      -- 头像
  `birthday` date DEFAULT NULL,                           -- 生日
  `country` TEXT DEFAULT '',                              -- 国家
  `province` TEXT DEFAULT '',                             -- 省
  `city` TEXT DEFAULT '',                                 -- 市
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_admin_order` (                           -- 管理员_充值订单
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键
  `member_id` INTEGER DEFAULT 0,                          -- 管理员id
  `order_type` TEXT NOT NULL,                             -- 订单类型
  `product_id` INTEGER DEFAULT NULL,                      -- 产品id
  `order_sn` TEXT DEFAULT '',                             -- 关联订单号
  `money` decimal(10,2) NOT NULL,                         -- 充值金额
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `refund_reason` TEXT DEFAULT NULL,                      -- 退款原因
  `reject_refund_reason` TEXT DEFAULT NULL,               -- 拒绝退款原因
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_admin_post` (                            -- 管理员_岗位
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 岗位ID
  `code` TEXT NOT NULL,                                   -- 岗位编码
  `name` TEXT NOT NULL,                                   -- 岗位名称
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `sort` INTEGER NOT NULL,                                -- 排序
  `status` INTEGER NOT NULL,                              -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_admin_role` (                            -- 管理员_角色信息
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 角色ID
  `name` TEXT NOT NULL,                                   -- 角色名称
  `key` TEXT NOT NULL,                                    -- 角色权限字符串
  `data_scope` INTEGER DEFAULT 1,                         -- 数据范围
  `custom_dept` TEXT DEFAULT NULL,                        -- 自定义部门权限
  `pid` INTEGER DEFAULT 0,                                -- 上级角色ID
  `level` INTEGER NOT NULL DEFAULT 1,                     -- 关系树等级
  `tree` TEXT DEFAULT NULL,                               -- 关系树
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `sort` INTEGER NOT NULL DEFAULT 0,                      -- 排序
  `status` INTEGER NOT NULL DEFAULT 1,                    -- 角色状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_admin_role_casbin` (
`id` INTEGER NOT NULL ,
`p_type` TEXT DEFAULT NULL,
`v0` TEXT DEFAULT NULL,
`v1` TEXT DEFAULT NULL,
`v2` TEXT DEFAULT NULL,
`v3` TEXT DEFAULT NULL,
`v4` TEXT DEFAULT NULL,
`v5` TEXT DEFAULT NULL,
PRIMARY KEY (`id`)
);
CREATE TABLE `hg_admin_role_menu` (                       -- 管理员_角色菜单关联
  `role_id` INTEGER NOT NULL,                             -- 角色ID
  `menu_id` INTEGER NOT NULL,                             -- 菜单ID
  PRIMARY KEY (`role_id`, `menu_id`)
);
CREATE TABLE `hg_pay_log` (                               -- 支付_支付日志
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键
  `member_id` INTEGER DEFAULT 0,                          -- 会员ID
  `app_id` TEXT DEFAULT NULL,                             -- 应用ID
  `addons_name` TEXT DEFAULT '',                          -- 插件名称
  `order_sn` TEXT DEFAULT '',                             -- 关联订单号
  `order_group` TEXT DEFAULT '',                          -- 组别[默认统一支付类型]
  `openid` TEXT DEFAULT '',                               -- openid
  `mch_id` TEXT DEFAULT '',                               -- 商户支付账户
  `subject` TEXT DEFAULT NULL,                            -- 订单标题
  `detail` TEXT DEFAULT NULL,                             -- 支付商品详情
  `auth_code` TEXT DEFAULT '',                            -- 刷卡码
  `out_trade_no` TEXT DEFAULT '',                         -- 商户订单号
  `transaction_id` TEXT DEFAULT NULL,                     -- 交易号
  `pay_type` TEXT NOT NULL,                               -- 支付类型
  `pay_amount` decimal(10,2) NOT NULL DEFAULT 0.00,       -- 支付金额
  `actual_amount` decimal(10,2) DEFAULT NULL,             -- 实付金额
  `pay_status` INTEGER DEFAULT 0,                         -- 支付状态
  `pay_at` datetime DEFAULT NULL,                         -- 支付时间
  `trade_type` TEXT DEFAULT '',                           -- 交易类型
  `refund_sn` TEXT DEFAULT NULL,                          -- 退款单号
  `is_refund` INTEGER DEFAULT 0,                          -- 是否退款
  `custom` text DEFAULT NULL,                             -- 自定义参数
  `create_ip` TEXT DEFAULT NULL,                          -- 创建者IP
  `pay_ip` TEXT DEFAULT NULL,                             -- 支付者IP
  `notify_url` TEXT DEFAULT NULL,                         -- 支付通知回调地址
  `return_url` TEXT DEFAULT NULL,                         -- 买家付款成功跳转地址
  `trace_ids` TEXT DEFAULT NULL,                          -- 链路ID集合
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_pay_refund` (                            -- 支付_退款记录
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键ID
  `member_id` INTEGER DEFAULT 0,                          -- 会员ID
  `app_id` TEXT DEFAULT NULL,                             -- 应用ID
  `order_sn` TEXT DEFAULT NULL,                           -- 业务订单号
  `refund_trade_no` TEXT DEFAULT NULL,                    -- 退款交易号
  `refund_money` decimal(10,2) DEFAULT NULL,              -- 退款金额
  `refund_way` INTEGER DEFAULT 1,                         -- 退款方式
  `ip` TEXT DEFAULT NULL,                                 -- 申请者IP
  `reason` TEXT DEFAULT NULL,                             -- 申请退款原因
  `remark` TEXT DEFAULT NULL,                             -- 退款备注
  `status` INTEGER DEFAULT 1,                             -- 退款状态
  `created_at` datetime DEFAULT NULL,                     -- 申请时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_addons_config` (                     -- 系统_插件配置
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 配置ID
  `addon_name` TEXT NOT NULL,                             -- 插件名称
  `group` TEXT NOT NULL,                                  -- 分组
  `name` TEXT DEFAULT '',                                 -- 参数名称
  `type` TEXT NOT NULL,                                   -- 键值类型:string,int,uint,bool,datetime,date
  `key` TEXT DEFAULT '',                                  -- 参数键名
  `value` TEXT DEFAULT '',                                -- 参数键值
  `default_value` TEXT NOT NULL,                          -- 默认值
  `sort` INTEGER NOT NULL DEFAULT 0,                      -- 排序
  `tip` TEXT DEFAULT NULL,                                -- 变量描述
  `is_default` INTEGER DEFAULT 0,                         -- 是否为系统默认
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_addons_install` (                    -- 系统_插件安装记录
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键
  `name` TEXT NOT NULL,                                   -- 插件名称
  `version` TEXT NOT NULL DEFAULT '',                     -- 版本号
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_attachment` (                        -- 系统_附件管理
  `id` INTEGER NOT NULL PRIMARY KEY ,                     -- 文件ID
  `app_id` TEXT NOT NULL,                                 -- 应用ID
  `member_id` INTEGER DEFAULT 0,                          -- 管理员ID
  `cate_id` INTEGER DEFAULT 0,                            -- 上传分类
  `drive` TEXT DEFAULT NULL,                              -- 上传驱动
  `name` TEXT DEFAULT NULL,                               -- 文件原始名
  `kind` TEXT DEFAULT NULL,                               -- 上传类型
  `mime_type` TEXT NOT NULL DEFAULT '',                   -- 扩展类型
  `naive_type` TEXT NOT NULL,                             -- NaiveUI类型
  `path` TEXT DEFAULT NULL,                               -- 本地路径
  `file_url` TEXT DEFAULT NULL,                           -- url
  `size` INTEGER DEFAULT 0,                               -- 文件大小
  `ext` TEXT DEFAULT NULL,                                -- 扩展名
  `md5` TEXT DEFAULT NULL,                                -- md5校验码
  `status` INTEGER NOT NULL DEFAULT 1,                    -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_sys_blacklist` (                         -- 系统_访问黑名单
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 黑名单ID
  `ip` TEXT DEFAULT '',                                   -- IP地址
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_config` (                            -- 系统_配置
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 配置ID
  `group` TEXT NOT NULL,                                  -- 配置分组
  `name` TEXT DEFAULT '',                                 -- 参数名称
  `type` TEXT NOT NULL,                                   -- 键值类型:string,int,uint,bool,datetime,date
  `key` TEXT DEFAULT '',                                  -- 参数键名
  `value` TEXT DEFAULT NULL,                              -- 参数键值
  `default_value` TEXT NOT NULL,                          -- 默认值
  `sort` INTEGER NOT NULL DEFAULT 0,                      -- 排序
  `tip` TEXT DEFAULT NULL,                                -- 变量描述
  `is_default` INTEGER DEFAULT 0,                         -- 是否为系统默认
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_cron` (                              -- 系统_定时任务
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 任务ID
  `group_id` INTEGER NOT NULL,                            -- 分组ID
  `title` TEXT NOT NULL,                                  -- 任务标题
  `name` TEXT DEFAULT NULL,                               -- 任务方法
  `params` TEXT DEFAULT NULL,                             -- 函数参数
  `pattern` TEXT NOT NULL,                                -- 表达式
  `policy` INTEGER NOT NULL DEFAULT 1,                    -- 策略
  `count` INTEGER NOT NULL DEFAULT 0,                     -- 执行次数
  `sort` INTEGER DEFAULT 0,                               -- 排序
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 任务状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_cron_group` (                        -- 系统_定时任务分组
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 任务分组ID
  `pid` INTEGER NOT NULL,                                 -- 父类任务分组ID
  `name` TEXT DEFAULT '',                                 -- 分组名称
  `is_default` INTEGER DEFAULT 0,                         -- 是否默认
  `sort` INTEGER DEFAULT 0,                               -- 排序
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 分组状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_dict_data` (                         -- 系统_字典数据
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 字典数据ID
  `label` TEXT DEFAULT NULL,                              -- 字典标签
  `value` TEXT DEFAULT NULL,                              -- 字典键值
  `value_type` TEXT NOT NULL DEFAULT 'string',            -- 键值数据类型：string,int,uint,bool,datetime,date
  `type` TEXT DEFAULT NULL,                               -- 字典类型
  `list_class` TEXT DEFAULT NULL,                         -- 表格回显样式
  `is_default` INTEGER DEFAULT 2,                         -- 是否为系统默认
  `sort` INTEGER DEFAULT 0,                               -- 字典排序
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_dict_type` (                         -- 系统_字典类型
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 字典类型ID
  `pid` INTEGER NOT NULL,                                 -- 父类字典类型ID
  `name` TEXT DEFAULT '',                                 -- 字典类型名称
  `type` TEXT DEFAULT '',                                 -- 字典类型
  `sort` INTEGER DEFAULT 0,                               -- 排序
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 字典类型状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_ems_log` (                           -- 系统_邮件发送记录
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键
  `event` TEXT NOT NULL,                                  -- 事件
  `email` TEXT NOT NULL,                                  -- 邮箱地址，多个用;隔开
  `code` TEXT DEFAULT '',                                 -- 验证码
  `times` INTEGER NOT NULL,                               -- 验证次数
  `content` TEXT DEFAULT NULL,                            -- 邮件内容
  `ip` TEXT DEFAULT NULL,                                 -- ip地址
  `status` INTEGER DEFAULT 1,                             -- 状态(1未验证,2已验证)
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_gen_codes` (                         -- 系统_代码生成记录
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 生成ID
  `gen_type` INTEGER UNSIGNED NOT NULL,                   -- 生成类型
  `gen_template` INTEGER DEFAULT 0,                       -- 生成模板
  `var_name` TEXT NOT NULL,                               -- 实体命名
  `options` TEXT DEFAULT NULL,                            -- 配置选项
  `db_name` TEXT DEFAULT NULL,                            -- 数据库名称
  `table_name` TEXT NOT NULL,                             -- 主表名称
  `table_comment` TEXT DEFAULT NULL,                      -- 主表注释
  `dao_name` TEXT DEFAULT NULL,                           -- 主表dao模型
  `master_columns` TEXT DEFAULT NULL,                     -- 主表字段
  `addon_name` TEXT DEFAULT NULL,                         -- 插件名称
  `status` INTEGER DEFAULT 1,                             -- 生成状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_sys_gen_curd_demo` (                     -- 系统_生成curd演示
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- ID
  `category_id` INTEGER DEFAULT 0,                        -- 分类ID
  `title` TEXT NOT NULL,                                  -- 标题
  `description` TEXT DEFAULT '',                          -- 描述
  `content` text DEFAULT NULL,                            -- 内容
  `image` TEXT DEFAULT NULL,                              -- 单图
  `attachfile` TEXT DEFAULT NULL,                         -- 附件
  `city_id` INTEGER DEFAULT 0,                            -- 所在城市
  `switch` INTEGER DEFAULT 1,                             -- 显示开关
  `sort` INTEGER DEFAULT NULL,                            -- 排序
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_by` INTEGER DEFAULT 0,                         -- 创建者
  `updated_by` INTEGER DEFAULT 0,                         -- 更新者
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL,                     -- 修改时间
  `deleted_at` datetime DEFAULT NULL                      -- 删除时间
);
CREATE TABLE `hg_sys_gen_tree_demo` (                     -- 系统_生成树演示
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- ID
  `pid` INTEGER DEFAULT NULL,                             -- 上级ID
  `level` INTEGER DEFAULT 1,                              -- 关系树级别
  `tree` TEXT DEFAULT NULL,                               -- 关系树
  `category_id` INTEGER DEFAULT 0,                        -- 分类ID
  `title` TEXT NOT NULL,                                  -- 标题
  `description` TEXT DEFAULT NULL,                        -- 描述
  `sort` INTEGER DEFAULT NULL,                            -- 排序
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_by` INTEGER DEFAULT 0,                         -- 创建者
  `updated_by` INTEGER DEFAULT 0,                         -- 更新者
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL,                     -- 修改时间
  `deleted_at` datetime DEFAULT NULL                      -- 删除时间
);
CREATE TABLE `hg_sys_log` (                               -- 系统_全局日志
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 日志ID
  `req_id` TEXT DEFAULT NULL,                             -- 对外ID
  `app_id` TEXT DEFAULT '',                               -- 应用ID
  `merchant_id` INTEGER DEFAULT 0,                        -- 商户ID
  `member_id` INTEGER DEFAULT 0,                          -- 用户ID
  `method` TEXT DEFAULT NULL,                             -- 提交类型
  `module` TEXT DEFAULT NULL,                             -- 访问模块
  `url` TEXT DEFAULT NULL,                                -- 提交url
  `get_data` TEXT,                                        -- get数据
  `post_data` TEXT,                                       -- post数据
  `header_data` TEXT,                                     -- header数据
  `ip` TEXT DEFAULT NULL,                                 -- IP地址
  `province_id` INTEGER NOT NULL DEFAULT 0,               -- 省编码
  `city_id` INTEGER NOT NULL DEFAULT 0,                   -- 市编码
  `error_code` INTEGER DEFAULT 0,                         -- 报错code
  `error_msg` TEXT DEFAULT NULL,                          -- 对外错误提示
  `error_data` TEXT,                                      -- 报错日志
  `user_agent` TEXT DEFAULT NULL,                         -- UA信息
  `take_up_time` INTEGER DEFAULT 0,                       -- 请求耗时
  `timestamp` INTEGER DEFAULT 0,                          -- 响应时间
  `status` INTEGER NOT NULL DEFAULT 1,                    -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE IF NOT EXISTS "hg_sys_login_log" (           -- 系统_登录日志
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 日志ID
  `req_id` TEXT DEFAULT NULL,                             -- 请求ID
  `member_id` INTEGER DEFAULT 0,                          -- 用户ID
  `username` TEXT DEFAULT NULL,                           -- 用户名
  `response` TEXT,                                        -- 响应数据
  `login_at` datetime DEFAULT NULL,                       -- 登录时间
  `login_ip` TEXT DEFAULT NULL,                           -- 登录IP
  `province_id` INTEGER DEFAULT NULL,                     -- 省编码
  `city_id` INTEGER DEFAULT NULL,                         -- 市编码
  `user_agent` TEXT DEFAULT NULL,                         -- UA信息
  `err_msg` TEXT DEFAULT NULL,                            -- 错误提示
  `status` INTEGER NOT NULL DEFAULT 1,                    -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_sys_serve_license` (                     -- 系统_服务许可证
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 许可ID
  `group` TEXT NOT NULL,                                  -- 分组
  `name` TEXT NOT NULL,                                   -- 许可名称
  `appid` TEXT NOT NULL,                                  -- 应用ID
  `secret_key` TEXT DEFAULT NULL,                         -- 应用秘钥
  `remote_addr` TEXT DEFAULT NULL,                        -- 最后连接地址
  `online_limit` INTEGER DEFAULT 1,                       -- 在线限制
  `login_times` INTEGER DEFAULT NULL,                     -- 登录次数
  `last_login_at` datetime DEFAULT NULL,                  -- 最后登录时间
  `last_active_at` datetime DEFAULT NULL,                 -- 最后心跳
  `routes` TEXT,                                          -- 路由表，空使用默认分组路由
  `allowed_ips` TEXT DEFAULT NULL,                        -- IP白名单
  `end_at` datetime NOT NULL,                             -- 授权有效期
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_sys_serve_log` (                         -- 系统_服务日志
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 日志ID
  `trace_id` TEXT DEFAULT NULL,                           -- 链路ID
  `level_format` TEXT DEFAULT NULL,                       -- 日志级别
  `content` text DEFAULT NULL,                            -- 日志内容
  `stack` TEXT,                                           -- 打印堆栈
  `line` TEXT NOT NULL,                                   -- 调用行
  `trigger_ns` INTEGER DEFAULT NULL,                      -- 触发时间(ns)
  `status` INTEGER NOT NULL DEFAULT 1,                    -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 修改时间
);
CREATE TABLE `hg_sys_sms_log` (                           -- 系统_短信发送记录
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 主键
  `event` TEXT NOT NULL,                                  -- 事件
  `mobile` TEXT NOT NULL DEFAULT '',                      -- 手机号
  `code` TEXT DEFAULT '',                                 -- 验证码或短信内容
  `times` INTEGER NOT NULL,                               -- 验证次数
  `ip` TEXT DEFAULT NULL,                                 -- ip地址
  `status` INTEGER DEFAULT 1,                             -- 状态(1未验证,2已验证)
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_test_category` (                         -- 测试分类
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- 分类ID
  `name` TEXT NOT NULL,                                   -- 分类名称
  `short_name` TEXT DEFAULT NULL,                         -- 简称
  `description` TEXT DEFAULT NULL,                        -- 描述
  `sort` INTEGER NOT NULL,                                -- 排序
  `remark` TEXT DEFAULT NULL,                             -- 备注
  `status` INTEGER DEFAULT 1,                             -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL,                     -- 修改时间
  `deleted_at` datetime DEFAULT NULL                      -- 删除时间
);
CREATE TABLE `hg_sys_provinces` (                         -- 系统_省市区编码
  `id` INTEGER NOT NULL PRIMARY KEY,                      -- 省市区ID
  `title` TEXT NOT NULL DEFAULT '',                       -- 栏目名称
  `pinyin` TEXT DEFAULT '',                               -- 拼音
  `lng` TEXT DEFAULT '',                                  -- 经度
  `lat` TEXT DEFAULT '',                                  -- 纬度
  `pid` INTEGER NOT NULL DEFAULT 0,                       -- 父栏目
  `level` INTEGER NOT NULL DEFAULT 1,                     -- 关系树等级
  `tree` TEXT NOT NULL,                                   -- 关系
  `sort` INTEGER DEFAULT 0,                               -- 排序
  `status` INTEGER NOT NULL DEFAULT 1,                    -- 状态
  `created_at` datetime DEFAULT NULL,                     -- 创建时间
  `updated_at` datetime DEFAULT NULL                      -- 更新时间
);
CREATE TABLE `hg_gen_curd_test` (                         -- 测试_代码生成
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        -- ID
  "name" VARCHAR(50),                                     -- 姓名
  "email" VARCHAR(255),                                   -- 邮箱
  "address" TEXT,                                         -- 地址
  "salt" VARCHAR(50),                                     -- 密码盐
  "password" VARCHAR(50),                                 -- 密码
  "mark" VARCHAR(255),                                    -- 备注
  "permission" TEXT,                                      -- 权限
  "created_user_id" INTEGER,                              -- 创建人
  "created_at" DATETIME,                                  -- 创建时间
  "updated_at" DATETIME,                                  -- 更新时间
  "deleted_at" DATETIME                                   -- 删除时间
);

CREATE INDEX `hg_addon_hgexample_tenant_order_order_sn` ON `hg_addon_hgexample_tenant_order` (`order_sn`);
CREATE INDEX `hg_addon_hgexample_tenant_order_member_id` ON `hg_addon_hgexample_tenant_order` (`user_id`);
CREATE INDEX `hg_addon_hgexample_tenant_order_merchant_id` ON `hg_addon_hgexample_tenant_order` (`merchant_id`);
CREATE INDEX `hg_addon_hgexample_tenant_order_agent_id` ON `hg_addon_hgexample_tenant_order` (`tenant_id`);
CREATE INDEX `hg_admin_cash_admin_id` ON `hg_admin_cash` (`member_id`);
CREATE INDEX `hg_admin_credits_log_member_id` ON `hg_admin_credits_log` (`member_id`);
CREATE INDEX `hg_admin_dept_pid` ON `hg_admin_dept` (`pid`);
CREATE UNIQUE INDEX `hg_admin_member_invite_code` ON `hg_admin_member` (`invite_code`);
CREATE INDEX `hg_admin_member_dept_id` ON `hg_admin_member` (`dept_id`);
CREATE INDEX `hg_admin_member_pid` ON `hg_admin_member` (`pid`);
CREATE UNIQUE INDEX `hg_admin_menu_name` ON `hg_admin_menu` (`name`);
CREATE INDEX `hg_admin_menu_pid` ON `hg_admin_menu` (`pid`);
CREATE INDEX `hg_admin_menu_status` ON `hg_admin_menu` (`status`);
CREATE INDEX `hg_admin_menu_type` ON `hg_admin_menu` (`type`);
CREATE INDEX `hg_admin_oauth_oauth_client` ON `hg_admin_oauth` (`oauth_client`);
CREATE INDEX `hg_admin_oauth_member_id` ON `hg_admin_oauth` (`member_id`);
CREATE INDEX `hg_admin_order_order_sn` ON `hg_admin_order` (`order_sn`);
CREATE INDEX `hg_admin_order_member_id` ON `hg_admin_order` (`member_id`);
CREATE UNIQUE INDEX `hg_pay_log_order_sn` ON `hg_pay_log` (`order_sn`);
CREATE INDEX `hg_pay_log_member_id` ON `hg_pay_log` (`member_id`);
CREATE INDEX `hg_pay_refund_order_sn` ON `hg_pay_refund` (`order_sn`);
CREATE UNIQUE INDEX `hg_sys_addons_config_addon_name_2` ON `hg_sys_addons_config` (`addon_name`);
CREATE INDEX `hg_addons_config_addon_name` ON `hg_sys_addons_config` (`addon_name`);
CREATE UNIQUE INDEX `hg_sys_addons_install_name` ON `hg_sys_addons_install` (`name`);
CREATE INDEX `hg_sys_attachment_md5` ON `hg_sys_attachment` (`md5`);
CREATE UNIQUE INDEX `hg_sys_blacklist_name` ON `hg_sys_blacklist` (`ip`);
CREATE INDEX `hg_sys_config_group` ON `hg_sys_config` (`group`);
CREATE INDEX `hg_sys_config_key` ON `hg_sys_config` (`key`);
CREATE INDEX `hg_sys_dict_data_dict_data_idx` ON `hg_sys_dict_data` (`type`);
CREATE UNIQUE INDEX `hg_sys_dict_type_dict_type` ON `hg_sys_dict_type` (`type`);
CREATE INDEX `hg_sys_ems_log_email` ON `hg_sys_ems_log` (`email`);
CREATE INDEX `hg_sys_log_error_code` ON `hg_sys_log` (`error_code`);
CREATE INDEX `hg_sys_log_req_id` ON `hg_sys_log` (`req_id`);
CREATE INDEX `hg_sys_log_member_id` ON `hg_sys_log` (`member_id`);
CREATE INDEX `hg_sys_login_log_member_id` ON `hg_sys_login_log` (`member_id`);
CREATE INDEX `hg_sys_login_log_req_id` ON `hg_sys_login_log` (`req_id`);
CREATE INDEX `hg_sys_provinces_pid` ON `hg_sys_provinces` (`pid`);
CREATE UNIQUE INDEX `hg_sys_serve_license_appid` ON `hg_sys_serve_license` (`appid`);
CREATE INDEX `hg_sys_serve_log_member_id` ON `hg_sys_serve_log` (`level_format`);
CREATE INDEX `hg_sys_serve_log_traceid` ON `hg_sys_serve_log` (`trace_id`);
CREATE INDEX `hg_sys_sms_log_mobile` ON `hg_sys_sms_log` (`mobile`);