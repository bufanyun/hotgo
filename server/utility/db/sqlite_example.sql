CREATE TABLE IF NOT EXISTS `user1` (                -- 用户管理
  `id` INTEGER,                                     -- 编号
  `name` VARCHAR(50),                               -- 名称
  `email` VARCHAR(255),                             -- 邮箱
  `address` TEXT,                                   -- 地址
  `salt` VARCHAR(50),                               -- 盐
  `password` VARCHAR(50),                           -- 密码
  `mark` VARCHAR(255),                              -- 备注
  `permission` TEXT,                                -- 权限
  `created_user_id` INTEGER,                        -- 创建者编号
  `created_at` DATETIME,                            -- 创建时间
  `updated_at` DATETIME,                            -- 更新时间
  `deleted_at` DATETIME,                            -- 删除时间
  PRIMARY KEY(`id`)
);
CREATE TABLE "user2" (                              -- 用户管理
  "id" INTEGER,                                     -- 编号
  "name" VARCHAR(50),                               -- 名称
  "email" VARCHAR(255),                             -- 邮箱
  "address" TEXT,                                   -- 地址
  "salt" VARCHAR(50),                               -- 盐
  "password" VARCHAR(50),                           -- 密码
  "mark" VARCHAR(255),                              -- 备注
  "permission" TEXT,                                -- 权限
  "created_user_id" INTEGER,                        -- 创建者编号
  "created_at" DATETIME,                            -- 创建时间
  "updated_at" DATETIME,                            -- 更新时间
  "deleted_at" DATETIME,                            -- 删除时间
  PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS user3 (                  -- 用户管理
  id INTEGER,                                       -- 编号
  name VARCHAR(50),                                 -- 名称
  email VARCHAR(255),                               -- 邮箱
  address TEXT,                                     -- 地址
  salt VARCHAR(50),                                 -- 盐
  password VARCHAR(50),                             -- 密码
  mark VARCHAR(255),                                -- 备注
  permission TEXT,                                  -- 权限
  created_user_id INTEGER,                          -- 创建者编号
  created_at DATETIME,                              -- 创建时间
  updated_at DATETIME,                              -- 更新时间
  deleted_at DATETIME,                              -- 删除时间
  PRIMARY KEY(id)
);