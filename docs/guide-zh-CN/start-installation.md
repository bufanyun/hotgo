## 系统安装

目录

- 环境要求
- 安装

### 环境要求

- node版本 >= v20.0.0
- golang版本 >= v1.23
- goframe版本 >=v2.9.4
- mysql版本 >=5.7 或 postgresql版本 >=14

> 必须先看[环境搭建文档](start-environment.md)，如果安装遇到问题务必先查看[常见问题文档](start-issue.md)

### 安装


一、克隆项目

```
git clone https://github.com/bufanyun/hotgo.git && cd hotgo
```

二、配置你的站点信息

1、服务端：
- 项目数据库文件 `storage/data/hotgo.sql` 创建数据库并导入
- 将配置文件 `manifest/config/config.example.yaml` 复制后改为`manifest/config/config.yaml`
- 将`manifest/config/config.yaml`中的`database.default.link`数据库配置改为你自己的：
```yaml
# Database. 配置参考：https://goframe.org/pages/viewpage.action?pageId=1114245
database:
  logger:
    path: "logs/database"                       # 日志文件路径。默认为空，表示关闭，仅输出到终端
    <<: *defaultLogger
    stdout: true
  default:
#   link: "mysql:账号:密码@tcp(127.0.0.1:3306)/数据库名?loc=Local&parseTime=true&charset=utf8mb4"
    link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo?loc=Local&parseTime=true&charset=utf8mb4"
    debug: true
    Prefix: "hg_"
```

- 将`hack/config.yaml`中的`gfcli.gen.dao[0].link`数据库配置改为你自己的：
```yaml
gfcli:
  gen:
    dao:
      - link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo?loc=Local&parseTime=true&charset=utf8mb4"
        group: "default"                                                # 分组 使用hotgo代码生成功能时必须填
        #        tables:          ""                                    # 指定当前数据库中需要执行代码生成的数据表。如果为空，表示数据库的所有表都会生成。
        tablesEx:        "hg_sys_addons_install"                        # 指定当前数据库中需要排除代码生成的数据表。
        removePrefix: "hg_"
        descriptionTag: true
        noModelComment: true
        jsonCase: "CamelLower"
        gJsonSupport: true
        clear: false
```

2、web前端：
- 配置服务端地址，包含在以下文件中：
* /hotgo/web/.env.development
* /hotgo/web/.env.production
* /hotgo/web/.env


三、 启动服务

1、服务端：
```shell
      cd server
      
      # 设置国内代理，如果已经设置好了代理可以跳过
      go env -w GOPROXY=https://goproxy.io,direct
      
      # 更新包
      go mod tidy  
      
      # 查看命令行方法
      go run main.go help
      
      # 启动所有服务
      go run main.go  # 热编译启动： gf run main.go
```

2、web前端：
```shell
    cd web
    # 首先确定你以安装node16.0以上版本并安装了包[npm、pnpm]，否则可能会出现一些未知报错
    
    # 安装依赖
    pnpm install 
    
    # 启动web项目
    pnpm run dev 
    
    # 如果顺利，至此到浏览器打开：http://你的IP:8001/admin
    # 登录账号：admin, 密码：123456
```





