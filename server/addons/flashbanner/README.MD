## 轮播图管理

### 简介




### 使用说明




### 迁移或安装

1、安装 HotGo (2.17.8及以上)

项目介绍：https://github.com/bufanyun/hotgo

2、将当前插件项目拷贝进 HotGo 根目录的 server/addons 目录下

3、在 HotGo 根目录的 server/addons/modules 目录下创建go文件:flashbanner.go，内容如下：
```go
package modules

import _ "hotgo/addons/flashbanner"
```

4、HotGo 后台进入 开发工具->插件管理->找到 轮播图管理 (flashbanner) 进行安装

5、重启服务即可生效


### 常用命令行

```shell
# 接口维护-gen service
gf gen service -s=addons/flashbanner/logic -d=addons/flashbanner/service

```

### `gf gen dao` 之后迁移model文件
```
cp .\internal\dao\internal\banner* .\addons\flashbanner\dao\internal\
cp .\internal\model\entity\banner* .\addons\flashbanner\model\entity\
```
