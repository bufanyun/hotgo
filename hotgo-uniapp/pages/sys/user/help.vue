<template>
	<view class="wrap">
		<view class="search">
			<u-search v-model="keywords" @custom="search" @search="search"></u-search>
		</view>
		<view class="u-p-t-5 u-p-b-5">
			<u-collapse class="box" :accordion="false" :arrow="false">
				<view class="item" v-for="(item, index) in list" :key="item.code">
					<u-collapse-item :open="true">
						<view class="title" slot="title">
							<u-icon :name="item.icon != '' ? item.icon : 'home'" :size="35"></u-icon>
							<view class="text">{{item.name}}</view>
						</view>
						<u-cell-group class="list" :border="false">
							<u-cell-item :arrow="true" v-for="(child, index2) in item.childList" :key="child.code" @click="navTo('/pages/sys/msg/form')">
								<text slot="title">{{child.name}}</text>
								<text slot="label">发送者：{{child.createByName}} &nbsp;|&nbsp; 时间：{{child.createDate}}</text>
							</u-cell-item>
						</u-cell-group>
					</u-collapse-item>
				</view>
			</u-collapse>
		</view>
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
					name: '常见问题',
					icon: 'error-circle',
					childList: [
						{
							code: 'a1',
							name: '为什么没有消息提醒？',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a2',
							name: '怎么关闭消息提醒？',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a3',
							name: '怎么办理工单？',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						}
					]
				},
				{
					code: 'a-1',
					name: '任务相关',
					icon: 'clock',
					childList: [
						{
							code: 'a1',
							name: '怎么办理任务？',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a2',
							name: '怎么完成任务？',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						},
						{
							code: 'a3',
							name: '怎么查询任务？',
							createByName: '管理员',
							createDate: '2021-4-6 12:10'
						}
					]
				}
			],
				
		};
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
