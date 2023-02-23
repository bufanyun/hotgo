## 目录结构

目录

- 服务端
- web前端
- uinapp端（待开放）


#### 服务端
```
/server
├── addons           
│   ├── modules        
│   ├── xxx插件 
│   |   ├── api
│   |   ├── controller
│   |   ├── global
│   |   ├── logic
│   |   ├── model
│   |   ├── router
│   |   ├── service
│   |   ├── main.go
│   |   └── README.md
├── api
│   ├── admin        
│   ├── api           
│   ├── home          
│   ├── websocket     
├── hack
├── internal
│   ├── cmd
│   ├── consts
│   ├── controller
│   ├── crons
│   ├── dao
│   ├── global
│   ├── library
│   ├── logic
│   ├── model
│   |   ├── do
│   │   ├── entity
│   │   └── input
│   ├── packed
│   ├── queues
│   ├── router
│   ├── service
│   └── websocket
├── manifest
├── resource
├── storage
├── utility
├── go.mod
├── main.go
├── Makefile
└── README.md
```

| 目录                       | 描述                                                              |
|--------------------------|-----------------------------------------------------------------|
| 基于gf的工程目录结构做了部分调整        | 参考地址： https://goframe.org/pages/viewpage.action?pageId=30740166 |
| **addons**               | 所有的插件模块都放在这里                                                    |
| --- modules              | 为插件模块提供隐式初始化                                                    |
| --- xxx插件                | 插件模块名称                                                          |
| --- --- api              | 对外接口。提供服务的输入/输出数据结构定义                                           |
| --- --- --- admin        | 后台接口                                                            |
| --- --- --- api          | 前台通用接口，包含PC页面、uinapp接口等                                         |
| --- --- --- home         | 前台PC端页面                                                         |
| --- --- --- websocket    | 可同时为多应用提供websocket接口                                            |
| --- --- controller       | 接收/解析用户输入参数的入口/接口层，也可以理解为控制器                                    |
| --- --- global           | 项目内主要的全局变量和系统的一些初始化操作                                           |
| --- --- logic            | 业务逻辑封装管理，特定的业务逻辑实现和封装往往是项目中最复杂的部分                               |
| --- --- model            | 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义                                 |
| --- --- --- input        | 对内接口。用于controller调用service或service之间调用时的输入/输出结构定义和输入过滤和预处理      |
| --- --- router           | 注册对外接口和分组中间件                                                    |
| --- --- service          | 用于业务模块解耦的接口定义层具体的接口实现在logic中进行注入                                |
| --- main.go              | 插件始化文件和模块插拔接口                                                   |
| **api**                  | 对外接口。提供服务的输入/输出数据结构定义                                           |
| --- admin                | 后台接口                                                            |
| --- api                  | 前台通用接口，包含PC页面、uinapp接口等                                         |
| --- home                 | 前台PC端页面                                                         |
| --- websocket            | 可同时为多应用提供websocket接口                                            |
| **hack**                 | 存放项目开发工具、脚本等内容例如，CLI工具的配置，各种shell/bat脚本等文件                      |
| **internal**             | 业务逻辑存放目录通过Golang internal特性对外部隐藏可见性                             |
| --- cmd                  | 命令行管理目录可以管理维护多个命令行                                              |
| --- consts               | 项目内主要的常量定义                                                      |
| --- controller           | 接收/解析用户输入参数的入口/接口层，也可以理解为控制器                                    |
| --- crons                | 项目中由系统统一接管的定时任务处理                                               |
| --- dao                  | 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法                      |
| --- global               | 项目内主要的全局变量和系统的一些初始化操作                                           |
| --- library              | 项目内常用功能的扩展库                                                     |
| --- logic                | 业务逻辑封装管理，特定的业务逻辑实现和封装往往是项目中最复杂的部分                               |
| --- model                | 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义                                 |
| --- --- do               | 用于dao数据操作中业务模型与实例模型转换，由工具维护，用户不能修改                              |
| --- --- entity           | 与数据集合绑定的程序数据结构定义，通常和数据表一一对应                                     |
| --- --- input            | 对内接口。用于controller调用service或service之间调用时的输入/输出结构定义和输入过滤和预处理      |
| --- packed               | 将静态资源打包进可执行文件，无需单独部署                                            |
| --- queues               | 为项目内所有的消息队列的消费者提供统一的初始化和处理                                      |
| --- router               | 注册对外接口和分组中间件                                                    |
| --- service              | 用于业务模块解耦的接口定义层具体的接口实现在logic中进行注入                                |
| **manifest**             | 包含程序编译、部署、运行、配置的文件常见内容如下：                                       |
| --- config               | 配置文件存放目录                                                        |
| --- docker               | Docker镜像相关依赖文件，脚本文件等等                                           |
| --- deploy               | 部署相关的文件默认提供了Kubernetes集群化部署的Yaml模板，通过kustomize管理                |
| **resource**           	 | 静态资源文件。这些文件往往可以通过 资源打包/镜像编译 的形式注入到发布文件中                         |
| **storage**           	  | 本地数据存储目录，例如文件缓存、磁盘队列数据、sql数据文件、SSL证书等                           |
| **utility**              | 一些常用的工具方法                                                       |
| go.mod                   | 使用Go Module包管理的依赖描述文件                                           |
| main.go                  | 程序入口文件                                                          |
| Makefile                 | 程序构建发布和开发快捷指令                                                   |
| README.md                | 项目介绍文件                                                          |




#### web前端
```
/web
├── build # 打包脚本相关
│   ├── config # 配置文件
│   ├── generate # 生成器
│   ├── script # 脚本
│   └── vite # vite配置
├── mock # mock文件夹
├── public # 公共静态资源目录
├── src # 主目录
│   ├── api # 接口文件
│   ├── assets # 资源文件
│   │   ├── icons # icon sprite 图标文件夹
│   │   ├── images # 项目存放图片的文件夹
│   │   └── svg # 项目存放svg图片的文件夹
│   ├── components # 公共组件
│   ├── design # 样式文件
│   ├── directives # 指令
│   ├── enums # 枚举/常量
│   ├── hooks # hook
│   │   ├── component # 组件相关hook
│   │   ├── core # 基础hook
│   │   ├── event # 事件相关hook
│   │   ├── setting # 配置相关hook
│   │   └── web # web相关hook
│   ├── layouts # 布局文件
│   │   ├── default # 默认布局
│   │   ├── iframe # iframe布局
│   │   └── page # 页面布局
│   ├── locales # 多语言
│   ├── logics # 逻辑
│   ├── main.ts # 主入口
│   ├── router # 路由配置
│   ├── settings # 项目配置
│   │   ├── componentSetting.ts # 组件配置
│   │   ├── designSetting.ts # 样式配置
│   │   ├── encryptionSetting.ts # 加密配置
│   │   ├── localeSetting.ts # 多语言配置
│   │   ├── projectSetting.ts # 项目配置
│   │   └── siteSetting.ts # 站点配置
│   ├── store # 数据仓库
│   ├── utils # 工具类
│   └── views # 页面
├── types # 类型文件
├── vite.config.ts # vite配置文件
└── windi.config.ts # windcss配置文件
```


#### uinapp端
```
// 待开放
```