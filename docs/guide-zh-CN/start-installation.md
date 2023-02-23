## 系统安装

目录

- 环境要求
- 安装

### 环境要求

- node版本 >= v16.0.0
- golang版本 >= v1.18
- goframe版本 >=v2.3.2
- mysql版本 >=5.7

> 必须先看[环境搭建文档](start-environment.md)，如果安装遇到问题务必先查看[常见问题文档](start-issue.md)

### 安装


一、克隆项目

```
git clone https://github.com/bufanyun/hotgo.git && cd hotgo
```

二、配置你的站点信息

1、服务端：
- 项目数据库文件 `storage/data/db.sql` 创建数据库并导入
- 修改配置 `manifest/config/config.yaml.bak` 复制改为`manifest/config/config.yaml`
- 将`manifest/config/config.yaml`中的数据库配置改为你自己的：
```yaml
database:
  logger:
    level: "all"
    stdout: true
  default:
    link: "mysql:hotgo:hg123456.@tcp(127.0.0.1:3306)/hotgo?loc=Local&parseTime=true"
    debug: true
    Prefix: "hg_"
```

2、web前端：
- 配置服务端地址，包含在以下文件中：
* /hotgo/web/.env.development
* /hotgo/web/.env.production
* /hotgo/web/.env

其中必改配置
```
VITE_PROXY=[["/admin","http://你的IP:8000/admin"]]
```


三、 启动服务

1、服务端：
```shell script
      cd server
      
      # 设置国内代理，如果已经设置好了代理可以跳过
      go env -w GOPROXY=https://goproxy.io,direct
      
      # 更新包
      go mod tidy  
      
      # 查看命令行方法
      go run main.go hlep
      
      # 启动所有服务
      go run main.go  # 热编译启动： gf run main.go
```

2、web前端：
```shell script
    cd web
    # 首先确定你以安装node16.0以上版本并安装了包[npm、yarn]，否则可能会出现一些未知报错
    
    # 安装依赖
    yarn install 
    
    # 启动web项目
    yarn dev 
    
    # 如果顺利，至此到浏览器打开：http://你的IP:8001/admin
    # 登录账号：admin, 密码：123456
```





