<p align="center">
    <img alt="JeeSite" src="https://jeesite.com/assets/images/logo.png" width="120" height="120" style="margin-bottom: 10px;">
</p>
<h3 align="center" style="margin:30px 0 30px;font-weight:bold;font-size:30px;">快速开发平台 - 手机端</h3>

## 引言

JeeSite Mobile Uni-App 是 JeeSite 手机端框架/移动端框架，基于 uni-app、uView UI 实现。

uni-app 是一个使用 Vue.js 开发所有前端应用的框架，开发者编写一套代码，可发布到iOS、Android、Web、
以及各种小程序（微信/支付宝/百度/头条/QQ/钉钉/淘宝）、快应用等多个平台。

uView UI，是 uni-app 生态最优秀的 UI 框架，全面的组件和便捷的工具会让您信手拈来，如鱼得水。

## 特性

* 支持 Android，iOS，H5，微信小程序，等其它小程序平台。
* 移动端是无 Cookie 环境的，该项目对移动端进行会话环境封装，
* 让你像 Cookie 一样使用 token，无需特别处理，有框架帮你完成。
* 化繁为简，封装 vuex 的繁琐，简单通过 api 即可进行 state 存取。
* 贴心的表单组件封装，下拉框、复选框、文件上传，完美与后端 JeeSite 结合。
* uView 提供 60+ 精选组件，功能丰富，多端兼容，让您快速集成，开箱即用。
* 众多贴心的 JS 利器，让您飞镖在手，召之即来，百步穿杨。
* 众多的常用页面和布局，让您专注逻辑，事半功倍。
* 合理使用 style 的 scoped 减少包体积大小。
* 详尽的文档支持，现代化的演示效果。
* 按需引入，精简打包体积。
* 移动端完整开源。

## 功能列表

* 账号登录、记住我（下次免登录）
* 自助服务：找回密码功能、账号注册功能
* 我的主页：修改个人信息、修改头像和裁剪、修改密码
* 辅助功能：关于我们、意见反馈、检查更新、帮助中心
* 工作台功能列表主页、消息列表页面
* 增删改查示例
* 工作流引擎

## 快速体验

1、H5 APP 端访问地址：<a href="https://demo.jeesite.com/app" target="blank">https://demo.jeesite.com/app</a> （最新演示）
<br>&nbsp; &nbsp; &nbsp; 获得H5最佳体验，操作方法：Chrome 为例，在浏览器上按 F12 打开“开发者工具”，点击该工具左上角第二个按钮
“Toggle device toolbar”，显示“切换设备工具栏”，然后在该工具栏上点击“Responsive”下拉选择“iPhone6/7/8”，再按“F5”刷新页面，即可。

2、微信小程序端：通过**微信**扫码（最佳体验，但不是最新演示，更新延迟）<br><br>
<img src="https://jeesite.com/assets/images/wx_app.jpg" width="220" height="220" >

3、安卓 Android 安装包，点击下载：[JeeSite4.3.1.apk](https://gitee.com/thinkgem/jeesite4/attach_files/925161/download)

## 快速运行

1、下载并安装：<a href="https://www.dcloud.io/hbuilderx.html" target="blank">集成开发环境 HBuilderX</a>
  （推荐，也可以使用 VSCode 或 WebStorm）

2、菜单：文件 -> 导入 -> 从本地目录导入，选择 “jeesite4-uniapp” 文件夹。

3、菜单：运行 -> 运行到内置浏览器（或运行到浏览器 -> Chrome 浏览器）。

4、等待 HBuliderX 控制台编译完成后，会自动弹出手机登录页面。

## 安装服务端
 
本项目后台服务默认连接的是 demo.jeesite.com 官网演示环境，你需要替换为你的 JeeSite 后台，步骤如下：

1、安装 JeeSite 最新版：<a href="https://gitee.com/thinkgem/jeesite4#%E6%9C%AC%E5%9C%B0%E8%BF%90%E8%A1%8C" target="blank">https://gitee.com/thinkgem/jeesite4#本地运行</a>
  （本项目支持 v4.2.3 或以上版本，若已安装，请执行 `bin/package.bat` 更新依赖）

2、打开 application.yml 修改如下配置（Ajax跨域设置和与后台基础交互的请求头名）：

```yml
# Shiro 相关
shiro:

  # 是否允许跨域访问 CORS，如果允许，设置允许的域名。v4.2.3 开始支持多个域名和模糊匹配，例如：http://*.jeesite.com,http://*.jeesite.net
  accessControlAllowOrigin: '*'
  
  # 允许跨域访问时 CORS，可以获取和返回的方法和请求头
  accessControlAllowMethods: GET, POST, OPTIONS
  accessControlAllowHeaders: content-type, x-requested-with, x-ajax, x-token, x-remember
  accessControlExposeHeaders: x-remember
  
# Session 相关
session:

  # 设置接收 SessionId 请求参数和请求头名称
  sessionIdHeaderName: x-token
  
  # 记住我的请求参数和请求头的名称
  rememberMeHeaderName: x-remember
  
# Web 相关
web:

  # AJAX 接受参数名和请求头名
  ajaxHeaderName: x-ajax
  
```

3、打开手机端项目的 `/common/config.js` 修改 `config.baseUrl` 后端服务地址为你安装的 JeeSite 服务地址。

## 生态系统

* 分布式微服务系统（Spring Cloud）：<https://gitee.com/thinkgem/jeesite4-cloud>
* JFlow工作流引擎：<https://gitee.com/thinkgem/jeesite4-jflow> ：<http://ccflow.org>
* Flowable业务流程模块（BPM）：<http://jeesite.com/docs/bpm/>
* 内容管理模块（CMS）：<https://gitee.com/thinkgem/jeesite4-cms>
* 手机端移动端：<https://gitee.com/thinkgem/jeesite4-uniapp>

## 学习路线

1. <a href="https://uniapp.dcloud.io/README" target="blank">什么是 uni-app、为什么选择 uni-app</a>
2. <a href="https://www.dcloud.io/hbuilderx.html" target="blank">集成开发环境 HBuilderX 下载</a>
3. <a href="https://ke.qq.com/course/3169971" target="blank">uni-app 官方视频教程</a>
4. <a href="http://ask.dcloud.net.cn/article/35657" target="blank">如果你熟悉 h5，但不熟悉 Vue 和小程序，请看这篇白话 uni-app</a>
5. <a href="https://uniapp.dcloud.io/vue-basics" target="blank">Vue.js 视频 + 文档教程</a>

## 学习文档

* <a href="https://uniapp.dcloud.io/collocation/pages" target="blank">uni-app 框架文档</a>
* <a href="https://uniapp.dcloud.io/component/README" target="blank">uni-app 组件文档</a>
* <a href="https://uviewui.com/components/intro.html" target="blank">uView 组件文档</a>
* <a href="https://uviewui.com/js/intro.html" target="blank">uView JS 文档</a>

## 打包发布

* 打开 `/common/config.js` 找到 `config.baseUrl` 修改为正式的手机端后台服务地址
* 阅读这篇文章：<a href="https://uniapp.dcloud.io/quickstart-hx?id=%e5%8f%91%e5%b8%83uni-app" target="blank">如何发布 uni-app 软件</a>
* <a href="https://ask.dcloud.net.cn/article/34972" target="blank">uni-app 整包升级、冷更新</a>
* <a href="https://ask.dcloud.net.cn/article/35667" target="blank">uni-app 资源升级、热更新</a>

## 授权许可协议条款

1. 基于 Apache License Version 2.0 协议发布，可用于商业项目，但必须遵守以下补充条款。
2. 不得将本软件应用于危害国家安全、荣誉和利益的行为，不能以任何形式用于非法为目的的行为。
3. 在延伸的代码中（修改和有源代码衍生的代码中）需要带有原来代码中的协议、版权声明和其他原作者
   规定需要包含的说明（请尊重原作者的著作权，不要删除或修改文件中的`Copyright`和`@author`信息）
   更不要，全局替换源代码中的 jeesite 或 ThinkGem 等字样，否则你将违反本协议条款承担责任。
4. 基于本软件完成的软件作品，只能使用 JeeSite 作为后台服务，除外情况不允许二次分发或开源。
5. 您若套用本软件的一些代码或功能参考，请保留源文件中的版权和作者，需要在您的软件介绍明显位置
   说明出处，举例：本软件基于 JeeSite 快速开发平台-手机端，并附带链接：http://jeesite.com
6. 任何基于本软件而产生的一切法律纠纷和责任，均于我司无关。
7. 如果你对本软件有改进，希望可以贡献给我们，共同进步。
8. 本项目已申请软件著作权，请尊重开源，感谢阅读。

## 技术服务与支持

* 本软件免费，我们也提供了相应的收费服务，因为：
* 没有资金的支撑就很难得到发展，特别是一个好的产品，如果 JeeSite 帮助了您，请为我们点赞（本软件Git仓库首页，右上角点击 star 按钮，关注我们）。支持我们，您可以得到一些回报，有了这些我们会把公益事业做的更好，回报社区和社会，请给我们一些动力吧，在此非常感谢已支持我们的朋友！
* **联系方式（官方商务）QQ：[1766571055](http://wpa.qq.com/msgrd?v=3&uin=1766571055&site=qq&menu=yes)**
* 技术服务支持网页：<http://s.jeesite.com>
