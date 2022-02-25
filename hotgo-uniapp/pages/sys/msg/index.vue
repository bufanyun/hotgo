<template>
	<view class="wrap">
		<js-lang title="msg.title"></js-lang>
		<view class="search">
			<u-search v-model="keywords" @custom="search" @search="search"></u-search>
		</view>
		<scroll-view class="scroll-list" scroll-y="true">
			<u-collapse class="box" :accordion="false" :arrow="false">
				<view class="item" v-for="(item, index) in list" :key="item.code">
					<u-collapse-item :open="true">
						<view class="title" slot="title">
							<u-icon :name="item.icon != '' ? item.icon : 'home'" :size="35"></u-icon>
							<view class="text">{{item.name}}</view>
							<u-badge v-if="item.count && item.count > 0" :count="item.count"></u-badge>
						</view>
						<u-cell-group class="list" :border="false">
							<u-cell-item :arrow="true" v-for="(child, index2) in item.childList" :key="child.code" @click="navTo('/pages/sys/msg/form')">
								<text slot="title">{{child.name}}</text>
								<text slot="label">发送者：{{child.createByName}} | 时间：{{child.createDate}}</text>
							</u-cell-item>
						</u-cell-group>
					</u-collapse-item>
				</view>
			</u-collapse>
		</scroll-view>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	data() {
		return {
			keywords: '',
			
			list: [
				{
					code: 'a',
					name: '重要消息',
					icon: 'error-circle',
					count: 3,
					childList: [
						{
							code: 'a1',
							name: '铁马冰河入梦来铁马冰河入梦来铁马冰河入梦来河入梦来铁马冰河入梦来河入梦来铁马冰河入梦来河入梦来',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a2',
							name: '发文',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a3',
							name: '查询',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						}
					]
				},
				{
					code: 'a-1',
					name: '一般消息',
					icon: 'chat',
					childList: [
						{
							code: 'a1',
							name: '铁马冰河入梦来铁马冰河入梦来铁马冰河入梦来河入梦来',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a2',
							name: '发文',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a3',
							name: '查询',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						}
					]
				},
				{
					code: 'a-2',
					name: '系统通知',
					icon: 'bell',
					childList: [
						{
							code: 'a1',
							name: '铁马冰河入梦来铁马冰河入梦来铁马冰河入梦来河入梦来',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a2',
							name: '发文',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a3',
							name: '查询',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						}
					]
				},
				{
					code: 'a-3',
					name: '已读消息',
					icon: 'volume',
					childList: [
						{
							code: 'a1',
							name: '铁马冰河入梦来铁马冰河入梦来铁马冰河入梦来河入梦来',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a2',
							name: '发文',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a3',
							name: '查询',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						}
					]
				},
			],
				
		};
	},
	onLoad() {
		
	},
	methods: {
		navTo(url) {
			uni.navigateTo({
				url: url
			});
		},
		search(value) {
			this.$u.toast('搜索内容为：' + value)
		}
		
	}
};
</script>
<style lang="scss">
page {
	background-color: #f8f8f8;
}
</style>
