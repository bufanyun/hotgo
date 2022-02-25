# hotgo

#### HotGo 开发框架是一款 Golang 语言 web 开发框架

## 技术选型

* 前端：用基于 JeeSite Mobile Uni-App+aidex-sharp 构建基础页面。
* 后端：用 goframe 快速搭建基础API，goframe 是一个go语言编写的Web框架。
* 数据库：采用MySql(5.7)版本，使用 gorm 实现对数据库的基本操作。
* 缓存：使用Redis实现记录当前活跃用户的jwt令牌并实现多点登录限制。
* API文档：使用Swagger构建自动化文档。
* 消息队列：同时兼容 kafka、redis、rocketmq，一键配置切换到自己想用的MQ。


## 感谢(以下排名不分先后)

* goframe https://goframe.org
* JeeSite Mobile Uni-App https://gitee.com/thinkgem/jeesite4-uniapp
* aidex-sharp https://gitee.com/big-hedgehog/aidex-sharp

## 声明
* 目前项目还在开发中，部分源码尚未更新完整，但不影响搭建演示，如有疑问请联系作者！