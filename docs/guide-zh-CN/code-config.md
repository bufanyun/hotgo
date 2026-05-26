## 生成配置

目录

- 模板配置
- CLI配置
- 多数据库使用

### 模板配置

- 配置路径：server/manifest/config/config.yaml

```yaml
# 生成代码
hggen:
  allowedIPs: [ "127.0.0.1", "*" ]                                      # 白名单，*代表所有，只有允许的IP后台才能使用生成代码功能
  selectDbs: [ "default" ]                                              # 可选生成表的数据库配置名称，支持多库
  disableTables: [ "hg_sys_gen_codes","hg_admin_role_casbin" ]          # 禁用的表，禁用以后将不会在选择表中看到
  delimiters: [ "@{", "}" ]                                             # 模板引擎变量分隔符号
  # 生成应用模型，所有生成模板允许自定义，可以参考default模板进行改造
  application:
    # CRUD和关系树列表模板
    crud:
      templates:
        # 默认的主包模板
        - group: "default"                                              # 分组名称
          isAddon: false                                                # 是否为插件模板 false｜true
          masterPackage: "sys"                                          # 主包名称，需和controllerPath、logicPath、inputPath保持关联
          templatePath: "./resource/generate/default/curd"              # 模板路径
          apiPath: "./api/admin"                                        # goApi生成路径
          controllerPath: "./internal/controller/admin/sys"             # 控制器生成路径
          logicPath: "./internal/logic/sys"                             # 主要业务生成路径
          inputPath: "./internal/model/input/sysin"                     # 表单过滤器生成路径
          routerPath: "./internal/router/genrouter"                     # 生成路由表路径
          sqlPath: "./storage/data/generate"                            # 生成sql语句路径
          webApiPath: "../web/src/api"                                  # webApi生成路径
          webViewsPath: "../web/src/views"                              # web页面生成路径

        # 默认的插件包模板，{$name}会自动替换成实际的插件名称
        - group: "addon"                                                # 分组名称
          isAddon: true                                                 # 是否为插件模板 false｜true
          masterPackage: "sys"                                          # 主包名称，需和controllerPath、logicPath、inputPath保持关联
          templatePath: "./resource/generate/default/curd"              # 模板路径
          apiPath: "./addons/{$name}/api/admin"                         # goApi生成路径
          controllerPath: "./addons/{$name}/controller/admin/sys"       # 控制器生成路径
          logicPath: "./addons/{$name}/logic/sys"                       # 主要业务生成路径
          inputPath: "./addons/{$name}/model/input/sysin"               # 表单过滤器生成路径
          routerPath: "./addons/{$name}/router/genrouter"               # 生成路由表路径
          sqlPath: "./storage/data/generate/addons"                     # 生成sql语句路径
          webApiPath: "../web/src/api/addons/{$name}"                   # webApi生成路径
          webViewsPath: "../web/src/views/addons/{$name}"               # web页面生成路径

    # 消息队列模板
    queue:
      templates:
        - group: "default"
          templatePath: "./resource/generate/default/queue"

    # 定时任务模板
    cron:
      templates:
        - group: "default"
          templatePath: "./resource/generate/default/cron"

  # 生成插件模块，通过后台创建新插件时使用的模板，允许自定义，可以参考default模板进行改造
  addon:
    srcPath: "./resource/generate/default/addon"                    # 生成模板路径
    webApiPath: "../web/src/api/addons/{$name}"                     # webApi生成路径
    webViewsPath: "../web/src/views/addons/{$name}"                 # web页面生成路径
```

### CLI配置

- hotgo在生成dao、service配置时，默认了和gf官方一致的配置方式和代码生成规则。所以无论你是通过hotgo亦或gf命令生成，最终代码格式完全一致，遵循一致的代码规范。

- 配置路径：[server/hack/config.yaml](../../server/hack/config.yaml)

```yaml
gfcli:
  build:
    name: "hotgo"                          # 编译后的可执行文件名称
    #    arch: "all"                           #不填默认当前系统架构，可选：386,amd64,arm,all
    #    system: "all"                         #不填默认当前系统平台，可选：linux,darwin,windows,all
    mod: "none"
    cgo: 0
    packSrc: "resource"                    # 将resource目录打包进可执行文件，静态资源无需单独部署
    packDst: "internal/packed/packed.go"   # 打包后生成的Go文件路径，一般使用相对路径指定到本项目目录中
    version: ""
    output: "./temp/hotgo"                 # 可执行文件生成路径
    extra: ""

  gen:
    dao:
      - link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo?loc=Local&parseTime=true"
#      - link: "pgsql:postgres:hg123456@tcp(127.0.0.1:5432)/hotgo"
        group: "default"                                                # 分组 使用hotgo代码生成功能时必须填
        #        tables:          ""                                    # 指定当前数据库中需要执行代码生成的数据表。如果为空，表示数据库的所有表都会生成。
        tablesEx:        "hg_sys_addons_install"                        # 指定当前数据库中需要排除代码生成的数据表。
        removePrefix: "hg_"
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        clear: true

    service: # 生成业务配置
      srcFolder: "internal/logic"
      dstFolder: "internal/service"
      dstFileNameCase: "CamelLower"
      clear: true
```

### 多数据库使用

- 假设我们要增加一个库名为`hotgo2`、分组为`default2`的数据库，并要为其生成代码

1. 配置[server/hack/config.yaml](../../server/hack/config.yaml) 如下：
```yaml
  gen:
    dao:
      - link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo?loc=Local&parseTime=true"
        group: "default"                                                # 分组 使用hotgo代码生成功能时必须填
        tablesEx:        "hg_sys_addons_install"                        # 指定当前数据库中需要排除代码生成的数据表。
        removePrefix: "hg_"
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        clear: false
      - link: "mysql:hotgo2:hg123456.@tcp(127.0.0.1:3306)/hotgo2?loc=Local&parseTime=true"
        group: "default2"                                                # 分组 使用hotgo代码生成功能时必须填
        tablesEx:        "hg_sys_addons_install"                         # 指定当前数据库中需要排除代码生成的数据表。
        removePrefix: ""
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        clear: false
```

2. 配置`server/manifest/config/config.yaml`,

`database`配置如下：
```yaml
database:
  logger:
    level: "all"
    stdout: true
  default:
    link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo?loc=Local&parseTime=true"
#    link: "pgsql:postgres:hg123456@tcp(127.0.0.1:5432)/hotgo"
    debug: true
    Prefix: "hg_"
  default2:
    link: "mysql:hotgo2:hg123456.@tcp(127.0.0.1:3306)/hotgo2?loc=Local&parseTime=true"
    debug: true
    Prefix: ""
```

`hggen`配置如下：
```yaml
hggen:
  allowedIPs: ["127.0.0.1", "*"]                                      # 白名单，*代表所有，只有允许的IP后台才能使用生成代码功能
  selectDbs: [ "default", "default2" ]                                # 可选生成表的数据库配置名称，支持多库
  disableTables : ["hg_sys_gen_codes","hg_admin_role_casbin"]         # 禁用的表，禁用以后将不会在选择表中看到
  delimiters: ["@{", "}"]                                             # 模板引擎变量分隔符号
```

3. 登录HotGo后台 -> 开发工具 -> 代码生成 -> 找到立即生成按钮并打开，就会发现`数据库`选项增加了一个`default2`，后续生成步骤和生成例子完全一样

> 注意：上述的配置中所有的`default2`名称必须保持一致

