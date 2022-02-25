<template>
	<view class="wrap">
		<view class="article-title">{{banner.title}}</view>
		<view class="article-meta">
			{{banner.author_name}} 发起于 {{banner.published_at}}</text>
		</view>
		<view class="article-content">
			<rich-text :nodes="content"></rich-text>
		</view>
		<view class="u-p-t-5"></view>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	data() {
		return {
			banner: {
				cover: '/static/jeesite/banner/2.svg',
				title: '最好用的快速开发平台',
				author_name: '系统管理员',
				published_at: '2021-4-20 12:30'
			},
			content: `<p>
JeeSite 快速开发平台，不仅仅是一个后台开发框架，它是一个企业级快速开发解决方案，基于经典技术组合 Spring Boot、Spring MVC、Apache Shiro、MyBatis、Beetl、Bootstrap、AdminLTE 采用经典开发模式，让初学者能够更快的入门并投入到团队开发中去。在线代码生成功能，包括模块如：组织机构、角色用户、菜单及按钮授权、数据权限、系统参数、内容管理、工作流等。采用松耦合设计，模块增减便捷；界面无刷新，一键换肤；众多账号安全设置，密码策略；文件在线预览；消息推送；多元化第三方登录；在线定时任务配置；支持集群，支持SAAS；支持多数据源；支持读写分离、分库分表；支持微服务应用。
</p><p>
JeeSite 快速开发平台的主要目的是能够让初级的研发人员快速的开发出复杂的业务功能（经典架构会的人多），让开发者注重专注业务，其余有平台来封装技术细节，降低技术难度，从而节省人力成本，缩短项目周期，提高软件安全质量。
</p><p>
JeeSite 自 2013 年发布以来已被广大爱好者用到了企业、政府、医疗、金融、互联网等各个领域中，JeeSite 架构精良、易于扩展、大众思维的设计模式、工匠精神打磨每一个细节，深入开发者的内心，并荣获开源中国《最受欢迎中国开源软件》奖杯，期间也帮助了不少刚毕业的大学生，教师作为入门教材，快速的去实践。
</p><p>
JeeSite4 的升级，作者结合了多年总结和经验，以及各方面的应用案例，对架构完成了一次全部重构，也纳入很多新的思想。不管是从开发者模式、底层架构、逻辑处理还是到用户界面，用户交互体验上都有很大的进步，在不忘学习成本、提高开发效率的情况下，安全方面也做和很多工作，包括：身份认证、密码策略、安全审计、日志收集等众多安全选项供你选择。努力为大中小微企业打造全方位企业级快速开发解决方案。
</p>`
		};
	},
	onLoad() {
		
	},
	methods: {
		navTo(url) {
			uni.navigateTo({
				url: url
			});
		}
	}
};
</script>
<style lang="scss">
page {
	background-color: #f8f8f8;
}
.article-title {
	font-size: 38rpx;
	font-weight: 400;
	text-align: center;
	padding-bottom: 30rpx;
	margin: 30rpx;
	border-bottom: 1px solid #c0c4cb;
	color: #51a0d5;
}
.article-meta {
	padding: 10rpx 50upx 30rpx;
	color: gray;
}
.article-content {
	padding: 0 30upx;
	overflow: hidden;
	font-size: 30rpx;
	line-height: 50rpx;
	margin-bottom: 30rpx;
	text-indent: 60rpx;
	/deep/ p {
		margin-bottom: 20rpx;
	}
}
</style>
