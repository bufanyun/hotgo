## 功能案例

### 简介

系统的一些功能案例


### 使用说明

系统自带的功能使用示例及其说明，包含一些简单的交互


### 安装

1、安装 HotGo (2.11.4及以上)

项目介绍：https://github.com/bufanyun/hotgo

2、将当前插件项目拷贝进 HotGo 根目录的 server/addons 目录下

3、在 HotGo 根目录的 server/addons/modules 目录下创建go文件:hgexample.go，内容如下：
```go
package modules

import _ "hotgo/addons/hgexample"
```

4、HotGo 后台进入 开发工具->插件管理->找到 功能案例 (hgexample) 进行安装

5、重启服务即可生效


### 常用命令行

```shell
# 接口维护-gen service
gf gen service -s=addons/hgexample/logic -d=addons/hgexample/service

```
