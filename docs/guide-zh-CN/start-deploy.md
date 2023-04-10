## 生产部署

目录

- 编译配置
- 编译
- 修改生产配置文件
- 启动服务
- Nginx配置

### 编译配置

- 配置文件：server/hack/config.yaml，以下是默认配置
```yaml
gfcli:
  build:
    name: "hotgo"                              # 编译后的可执行文件名称
    #    arch: "all"                           #不填默认当前系统架构，可选：386,amd64,arm,all
    #    system: "all"                         #不填默认当前系统平台，可选：linux,darwin,windows,all
    mod: "none"
    cgo: 0
    packSrc: "resource"                        # 将resource目录打包进可执行文件，静态资源无需单独部署
    packDst: "internal/packed/packed.go"       # 打包后生成的Go文件路径，一般使用相对路径指定到本项目目录中
    version: ""
    output: "./temp/hotgo"                     # 可执行文件生成路径
    extra: ""
```


### 编译

- 以下方式任选其一即可

1、 make一键编译 （linux或mac环境推荐）
```shell
cd server &&  make build
``` 

2、 按步骤手动编译
```shell
cd server                                             # 切换到服务端目录下
rm -rf ./resource/public/admin/                       # 删除之前的web资源
mkdir ./resource/public/admin/                        # 重新创建web资源存放目录，除首次编译后续可以跳过执行此步骤
cd ../web && yarn build                               # 切换到web项目下，编译web项目
\cp -rf ./dist/*  ../server/resource/public/admin/    # 将编译好的web资源复制到server对应的资源存放路径下
echo "y" | gf build                                   # 编译hotgo服务端

# 不出意外你已经编译好了hotgo可执行文件！
``` 

3、分服务编译

待写。


### 修改生产配置文件
- 配置文件：server/manifest/config/config.yaml
> 如关闭代码生成功能、修改数据库地址、缓存驱动、队列驱动、日志路径等



### 启动服务
> 这里可以接使用gf官方推荐的启动方式，请参考：https://goframe.org/pages/viewpage.action?pageId=1114403


### Nginx配置
```
      # websocket
      location ^~ /socket  {
  			proxy_pass http://127.0.0.1:8000/socket;
  			proxy_set_header X-Real-IP $remote_addr;
  			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  			proxy_set_header Host $host;
  			proxy_set_header X-NginX-Proxy true;
  			proxy_http_version 1.1;
  			proxy_set_header Upgrade $http_upgrade;
  			proxy_set_header Connection "Upgrade";
  			proxy_connect_timeout 600s;
  			proxy_read_timeout 600;
  			proxy_send_timeout 600s;
      }

      # http
      location ^~ / {
          proxy_set_header Host $http_host;
          proxy_set_header  X-Real-IP $remote_addr;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_set_header X-Forwarded-Proto $scheme;
          proxy_pass http://127.0.0.1:8000/; # 设置代理服务器的协议和地址
          proxy_http_version 1.1;
          proxy_set_header Upgrade $http_upgrade;
          proxy_set_header Connection upgrade;
    }
```