<template>
	<view class="wrap">
		<js-lang title="home.title" :showBtn="true"></js-lang>
		<u-swiper :height="300" :list="imgList" :title="false" @click="imgListClick"></u-swiper>
		<view class="toolbar u-m-b-20">
			<u-grid class="grid" :col="3" :border="false">
				<u-grid-item :index="0" @click="navTo('/pages/bpm/myTaskTodo')">
					<u-badge :count="todoCount"></u-badge>
					<u-icon class="grid-icon" name="clock" :size="80" :style="{ color: '#ea9a44' }"></u-icon>
					<view class="grid-text">待办任务</view>
				</u-grid-item>
				<u-grid-item :index="1" @click="navTo('/pages/bpm/myTaskHistory')">
					<u-icon class="grid-icon" name="checkmark-circle" :size="80" :style="{ color: '#47cb66' }"></u-icon>
					<view class="grid-text">已办任务</view>
				</u-grid-item>
				<u-grid-item :index="2" @click="navTo('/pages/bpm/myRuntime')">
					<u-icon class="grid-icon" name="order" :size="80" :style="{ color: '#5a98ea' }"></u-icon>
					<view class="grid-text">我相关的</view>
				</u-grid-item>
			</u-grid>
		</view>
		<u-collapse class="box u-p-b-5" :accordion="false" :arrow="true">
			<view class="item" v-for="(menu, index) in menuList" :key="menu.menuCode">
				<u-collapse-item :open="true">
					<view class="title" slot="title">
						<u-icon :name="menu.menuIcon != '' ? menu.menuIcon : 'home'" :size="40"
							:style="{ color: menu.menuColor != '' ? menu.menuColor : '#666' }"></u-icon>
						<view class="text" :style="{ color: menu.menuColor != '' ? menu.menuColor : '#666' }"
							>{{menu.menuName}}</view>
					</view>
					<u-grid class="grid u-m-t-20" :col="3" :border="false">
						<u-grid-item v-for="(child, index2) in menu.childList" :key="child.menuCode" @click="navTo(child.url)">
							<u-icon class="grid-icon" :name="child.menuIcon != '' ? child.menuIcon : 'order'" :size="80"
								:style="{ color: child.menuColor != '' ? child.menuColor : '#666' }"></u-icon>
							<view class="grid-text" :style="{ color: child.menuColor != '' ? child.menuColor : '#666' }"
								>{{child.menuName}}</view>
						</u-grid-item>
					</u-grid>
				</u-collapse-item>
			</view>
		</u-collapse>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	data() {
		return {
			
			imgList: [
				{image: '/static/jeesite/banner/1.svg'},
				{image: '/static/jeesite/banner/2.svg'},
				{image: '/static/jeesite/banner/3.svg'}
			],
			
			todoCount: 0,
			
			menuList: [
				{
					menuCode: 'a-1',
					menuName: '增删改查',
					menuIcon: 'file-text',
					menuColor: '',
					url: '',
					childList: [
						{
							menuCode: 'a13',
							menuName: '列表',
							menuIcon: 'thumb-up',
							menuColor: '',
							url: '/pages/testData/index',
						},
						{
							menuCode: 'a11',
							menuName: '新增',
							menuIcon: 'plus-circle',
							menuColor: '',
							url: '/pages/testData/form',
						},
						{
							menuCode: 'a10',
							menuName: '请假',
							menuIcon: 'calendar',
							menuColor: '',
							url: '/pages/oa/oaLeave/index',
						},
					]
				},
				{
					menuCode: 'a',
					menuName: '公文管理',
					menuIcon: 'home',
					menuColor: '#919328',
					url: '',
					childList: [
						{
							menuCode: 'a1',
							menuName: '收文',
							menuIcon: 'email',
							menuColor: '#919328',
							url: '/pages/testData/form',
						},
						{
							menuCode: 'a2',
							menuName: '发文',
							menuIcon: 'bookmark',
							menuColor: '#919328',
							url: '/pages/testData/form',
						},
						{
							menuCode: 'a3',
							menuName: '查询',
							menuIcon: 'search',
							menuColor: '#919328',
							url: '/pages/testData/index',
						}
					]
				},
				{
					menuCode: 'a-2',
					menuName: '功能列表',
					menuIcon: '',
					menuColor: '#0d9311',
					url: '',
					childList: [
						{
							menuCode: 'a21',
							menuName: '找回密码',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/login/forget',
						},
						{
							menuCode: 'a22',
							menuName: '注册用户',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/login/reg',
						},
						{
							menuCode: 'a23',
							menuName: '个人资料',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/user/info',
						},{
							menuCode: 'a24',
							menuName: '关于我们',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/user/about',
						},
						{
							menuCode: 'a25',
							menuName: '修改密码',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/user/pwd',
						},
						{
							menuCode: 'a26',
							menuName: '意见反馈',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/user/comment',
						},
						{
							menuCode: 'a27',
							menuName: '系统设置',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/sys/user/setting',
						},
						{
							menuCode: 'a28',
							menuName: '列表演示',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/testData/index',
						},
						{
							menuCode: 'a29',
							menuName: '表单演示',
							menuIcon: '',
							menuColor: '#0d9311',
							url: '/pages/testData/form',
						}
					]
				},
			],
				
		};
	},
	onLoad() {
		//this.refreshCount();
	},
	onShow() {
		this.refreshCount();
	},
	methods: {
		navTo(url) {
			uni.navigateTo({
				url: url
			});
		},
		refreshCount() {
			this.todoCount = 3;
		},
		imgListClick(index) {
			console.log(`点击了第${index + 1}页图片`)
		},
		itemClick(index) {
			console.log(index);
		}
	}
};
</script>
<style lang="scss">
@import 'index.scss';
page {
	background-color: #f8f8f8;
}
</style>
