# HotGo-V2
<div align="center">
	<img width="140px" src="https://bufanyun.cn-bj.ufileos.com/hotgo/logo.sig.png">
    <p>
        <h1>HotGo V2.1</h1>
    </p>
    <p align="center">
        <a href="https://goframe.org/pages/viewpage.action?pageId=1114119" target="_blank">
	        <img src="https://img.shields.io/badge/goframe-2.3-green" alt="goframe">
	    </a>
	    <a href="https://v3.vuejs.org/" target="_blank">
	        <img src="https://img.shields.io/badge/vue.js-vue3.x-green" alt="vue">
	    </a>
	    <a href="https://www.naiveui.com" target="_blank">
	        <img src="https://img.shields.io/badge/naiveui-%3E2.0.0-blue" alt="naiveui">
	    </a>
		<a href="https://www.tslang.cn/" target="_blank">
	        <img src="https://img.shields.io/badge/typescript-%3E4.0.0-blue" alt="typescript">
	    </a>
		<a href="https://vitejs.dev/" target="_blank">
		    <img src="https://img.shields.io/badge/vite-%3E2.0.0-yellow" alt="vite">
		</a>
		<a href="https://github.com/bufanyun/hotgo/blob/v2.0/LICENSE" target="_blank">
		    <img src="https://img.shields.io/badge/license-MIT-success" alt="license">
		</a>
	</p>
</div>


## 平台简介
* 基于全新Go Frame 2+Vue3+Naive UI开发的全栈前后端分离的管理系统
* 前端采用naive-ui-admin 、Vue、Naive UI。


## 特征
* 高生产率：几分钟即可搭建一个后台管理系统
* 模块化：单应用多系统的模式，将一个完整的应用拆分为多个系统，后续扩展更加便捷，增加代码复用性。
* 插件化： 可通过插件的方式扩展系统功能
* 认证机制：采用jwt的用户状态认证及casbin的权限认证
* 路由模式：得利于goframe2.0提供了规范化的路由注册方式,无需注解自动生成api文档
* 面向接口开发


## 内置功能

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、岗位），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5. 角色管理：角色菜单权限分配、设置角色按机构或按上下级关系进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 配置管理：对系统动态配置常用参数。
8. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 服务日志：服务端运行所产生的警告、异常、崩溃日志的详细数据和堆栈信息。
11. 在线用户：当前系统中活跃用户状态监控。
12. 定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
13. 代码生成：支持自动化生成前后端代码。CURD关联表、树表、消息队列、定时任务一键生成等。
14. 服务监控：监视当前系统CPU、内存、磁盘、网络、堆栈等相关信息。
15. 附件管理：文件上传，多种上传方式适配。
16. 消息队列：同时兼容 kafka、redis、rocketmq、磁盘队列，一键配置切换到场景适用的MQ。
17. 通知公告：采用websocket实时推送在线用户最新通知、公告、私信消息。
18. 地区编码：整合国内通用省市区编码，运用于项目于一身，支持动态省市区选项。
19. 常用工具：集成常用的工具包和命令行工具，可以快速开发自定义命令行，多种启动入口。


> HotGo开源以来得到了大家的很多支持，本项目初衷只为互相学习交流，没有任何盈利性目的！欢迎为HotGo贡献代码或提供建议！

## 演示地址
-  [https://hotgo.facms.cn/admin](https://hotgo.facms.cn/admin)
>  账号：admin  密码：123456


## 环境要求
- node版本 >= v16.0.0
- golang版本 >= v1.18
- gf版本 >=v2.3.1 (会保持同步gf最新版本，gf小版本更新可能存在兼容问题，旧版本需自行处理，如非必要不建议更新！)
- IDE推荐：Goland
- mysql版本 >=5.7
- redis版本 >=3.0

## 快速开始
一、拉取代码到你已经安装好以上环境的服务器中
 ```shell script
git clone https://github.com/bufanyun/hotgo.git && cd hotgo
 ```

二、配置你的站点信息

服务端：
- 项目数据库文件 `resource/data/db.sql` 创建数据库并导入
- 修改配置 `manifest/config/config.yaml.bak` 复制改为`manifest/config/config.yaml`


后台前端：
- 配置服务端地址，包含在以下文件中：
* /hotgo/web/.env.development
* /hotgo/web/.env.production
* /hotgo/web/.env

其中必改配置
```
VITE_PROXY=[["/admin","http://你的IP:8000/admin"]]
```

三、 启动服务
服务端：
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
      
      # 如果顺利，至此到浏览器打开：http://你的IP:8000/admin，即可看到后台登录地址
      # 登录账号：admin, 密码：123456
      # 好奇为什么没有运行前端代码就能将后台运行起来？这要得益于gf强大的pack功能！
      # 当然这只是为了降低运行门槛，如果想对前端进行开发请继续往下看

```

web端：
   ```shell script
    cd web
    # 首先确定你以安装node16.0以上版本并安装了包[npm、yarn]，否则可能会出现一些未知报错
    
    # 安装依赖
    yarn install 
    
    # 启动web项目
    yarn dev 
    
    # 如果顺利，至此到浏览器打开：http://你的IP:8001/admin
```


## 文档地址
> 文档正在书写中，请耐心等一等。


## 演示图

<table>
    <tr>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/1.png"/></td>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/2.png"/></td>
    </tr>
    <tr>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/3.png"/></td>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/4.png"/></td>
    </tr>
    <tr>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/5.png"/></td>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/6.png"/></td>
    </tr>
    <tr>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/7.png"/></td>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/8.png"/></td>
    </tr>
    <tr>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/9.png"/></td>
        <td><img src="https://bufanyun.cn-bj.ufileos.com/hotgo/example/10.png"/></td>
    </tr>
</table>

## 感谢(排名不分先后)
> gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf)
>
> naive-ui [https://www.naiveui.com](https://www.naiveui.com)
>
> naive-ui-admin [https://github.com/jekip/naive-ui-admin](https://github.com/jekip/naive-ui-admin)
>
> websocket [https://github.com/gorilla/websocket](github.com/gorilla/websocket)
> 
> casbin [https://github.com/casbin/casbin](https://github.com/casbin/casbin)


## 交流QQ群
交流群①：190966648  <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=mJafkvme3VNyiQlCFIFNRtY8Xlr7pj9U&jump_from=webapi&authKey=jL10vIESr+vO8wpxwyd6DlChzkrbHpzN9uhAsIHgAinL/Vvd+nvuRyilf2UqUlCy"><img border="0" src="https://bufanyun.cn-bj.ufileos.com/hotgo/group.png" alt="HotGo框架交流1群" title="HotGo框架交流1群"></a>
> <img src="https://bufanyun.cn-bj.ufileos.com/hotgo/hotgo1qun.png" width="400px"/>  


> 感谢你使用HotGo，公司团队精力时间有限，因此我们不再提供免费的技术服务！
>
> 同时您也可以联系我们，雇佣我们团队为您干活，谢谢合作！


## 商用说明

> HotGo 是开源免费的，遵循 MIT 开源协议，意味着您无需支付任何费用，也无需授权，即可将它应用到您的产品中。

* 使用本项目必须保留所有版权信息。

* 本项目包含的第三方源码和二进制文件之版权信息另行标注。

* 版权所有Copyright © 2020-2023 by Ms (https://github.com/bufanyun/hotgo)

* All rights reserved。


## 免责声明：
* HotGo为开源学习项目，一切商业行为与HotGo无关。

* 用户不得利用HotGo从事非法行为，用户应当合法合规的使用，发现用户在使用产品时有任何的非法行为，HotGo有权配合有关机关进行调查或向政府部门举报，HotGo不承担用户因非法行为造成的任何法律责任，一切法律责任由用户自行承担，如因用户使用造成第三方损害的，用户应当依法予以赔偿。

* 所有与使用HotGo相关的资源直接风险均由用户承担。


#### 如果对您有帮助，您可以点右上角 💘Star💘支持



## [感谢JetBrains提供的免费GoLand](https://jb.gg/OpenSource)
[![avatar](https://camo.githubusercontent.com/323657c6e81419b8e151e9da4c71f409e3fcc65d630535170c59fe4807dbc905/68747470733a2f2f676f6672616d652e6f72672f646f776e6c6f61642f7468756d626e61696c732f313131343131392f6a6574627261696e732e706e67)](https://jb.gg/OpenSource)


## License
[MIT © HotGo-2023](./LICENSE)
  


  

