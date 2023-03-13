## 常见问题

目录

- 一、后台相关
- 二、数据库相关
- 三、环境相关


### 一、后台相关

#### 1、连接超时，请刷新重试。如仍未解决请检查websocket连接是否正确！

线上或非本地运行时，请到 系统设置 -> 配置管理 -> 基本设置 -> 找到网站域名和websocket地址，改成你自己实际的地址，保存刷新页面即可



### 二、数据库相关

#### 1、安装数据库出现 json 报错不支持

请安装 mysql5.7 及以上版本的数据库



### 三、环境相关

#### 1、not found in resource manager or following system searching paths

> 报错信息：panic: possible config files "config" or "config.toml/yaml/yml/json/ini/xml/properties" not found in resource manager or following system searching paths:

这是因为系统没有找到配置文件，将配置文件 `manifest/config/config.yaml.bak` 复制后改为`manifest/config/config.yaml`


详细请参考 - [系统安装](start-installation.md)
