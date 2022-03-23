# hotgo


#### HotGo 是一个基于 vue 和 goframe2.0 开发的全栈前后端分离的开发基础平台和移动应用平台，集成jwt鉴权，动态路由，动态菜单，casbin鉴权，消息队列，定时任务等功能，提供多种常用场景文件，让您把更多时间专注在业务开发上。

## 技术选型

* 后端：用 goframe2.0 快速搭建基础API，goframe2.0 是一个go语言编写的Web框架。
* 前端：用基于 JeeSite Mobile Uni-App+aidex-sharp 构建基础页面。
* 数据库：采用MySql(8.0)版本，使用 gorm 实现对数据库的基本操作。
* 缓存：使用Redis实现记录当前活跃用户的jwt令牌并实现多点登录限制。
* API文档：使用Swagger构建自动化文档。
* 消息队列：同时兼容 kafka、redis、rocketmq，一键配置切换到自己想用的MQ。

## 系统截图
#### * web端

![image](https://user-images.githubusercontent.com/26652343/155689571-e6a0a5a3-011b-44cc-b84b-a1c82301b207.png)

![image](https://user-images.githubusercontent.com/26652343/155689646-d3395261-6061-469f-8256-3cd0ff9f5d05.png)

![image](https://user-images.githubusercontent.com/26652343/155689709-5ddac1d3-1c01-4fab-9d3a-9ece72ca5ba0.png)

#### * 移动端
![image](https://user-images.githubusercontent.com/26652343/155689481-2fc019eb-18e4-4a94-b417-50524e945089.png)
![image](https://user-images.githubusercontent.com/26652343/155689738-ac97f9c0-47ae-499b-b3fe-0cb4ce97f3bc.png)

## 环境要求
- node版本 >= v14.0.0 
- golang版本 >= v1.16
- IDE推荐：Goland
- mysql版本 >=8.0
- redis版本 >=5.0

## 快速开始
 一、拉取代码到你已经安装好以上环境的服务器中
 ```shell script
git clone https://github.com/bufanyun/hotgo.git
 ```

二、配置你的站点信息

服务端：
 - 创建mysql数据库，将数据库文件导入你的mysql，目录地址：/hotgo-server/storage/hotgo.sql
 - 将/hotgo-server/config/config.example.yaml 改为：config.yaml，并根据你实际环境情况进行配置

web+uinapp端：
 - 配置服务端地址，包含在一下文件中：
 * hotgo-uniapp/common/config.js 
 * /hotgo-uniapp/manifest.json 
 * hotgo-uniapp/common/config.js 

三、 启动服务
服务端：
   ```shell script
  cd hotgo-server
  go mod tidy  #更新包
  go run main.go  #启动服务
```

web端：
   ```shell script
cd hotgo-web
npm install #安装依赖
npm run dev #启动web项目
```
uinapp端：
- 1、下载并安装：集成开发环境 HBuilderX （推荐，也可以使用 VSCode 或 WebStorm）
- 2、菜单：文件 -> 导入 -> 从本地目录导入，选择 “jeesite4-uniapp” 文件夹。
- 3、菜单：运行 -> 运行到内置浏览器（或运行到浏览器 -> Chrome 浏览器）。
- 4、等待 HBuliderX 控制台编译完成后，会自动弹出手机登录页面。


## 特别感谢(以下排名不分先后)

* goframe2.0 https://goframe.org
* JeeSite Mobile Uni-App https://gitee.com/thinkgem/jeesite4-uniapp
* aidex-sharp https://gitee.com/big-hedgehog/aidex-sharp

## 开源声明
* 目前项目还在持续更新中，仅供参考学习，如遇到问题请联系作者下方微信！

![image](https://user-images.githubusercontent.com/26652343/155691271-1ded98d8-f0f1-4467-9079-26cec1195af5.png)