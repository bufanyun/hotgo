## 常见问题

目录

- 一、后台相关
- 二、数据库相关
- 三、环境相关


### 一、后台相关

#### 1、连接超时，请刷新重试。如仍未解决请检查websocket连接是否正确！

线上或非本地运行时，请到 系统设置 -> 配置管理 -> 基本设置 -> 找到网站域名和websocket地址，改成你自己实际的地址，保存刷新页面即可

#### 2、web页面菜单切换后页面出现白屏

请参考：https://github.com/jekip/naive-ui-admin/issues/183


### 二、数据库相关

#### 1、安装数据库出现 json 报错不支持

请安装 mysql5.7 及以上版本的数据库



### 三、环境相关

#### 1、not found in resource manager or following system searching paths

> 报错信息：panic: possible config files "config" or "config.toml/yaml/yml/json/ini/xml/properties" not found in resource manager or following system searching paths:

系统没有找到配置文件，将配置文件 `manifest/config/config.yaml.bak` 复制后改为`manifest/config/config.yaml`


#### 2、net.DialTimeout failed with network

> 报错信息：connect to 127.0.0.1:8099 error: net.DialTimeout failed with network "tcp", address "127.0.0.1:8099", timeout "10s": dial tcp

- http服务没有启动或正在启动
- 通过一键启动所有服务运行时属正常情况，多服务启动时存在先后顺序问题，`tcpClient`比`tcpServer`先启动完成导致的，等`tcpServer`启动完成后会自动重连

详细请参考 - [系统安装](start-installation.md)


### 四、前端相关

#### 1、Error: connect ECONNREFUSED ::1:8000

```text
11:44:52 [vite] http proxy error at /member/info:
Error: connect ECONNREFUSED ::1:8000
    at TCPConnectWrap.afterConnect [as oncomplete] (node:net:1246:16)
```

- 服务端没有启动
- `.\wen\.env.development`中的`VITE_PROXY`配置的服务器地址或端口与实际不一致



