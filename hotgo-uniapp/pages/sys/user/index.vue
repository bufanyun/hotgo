<template>
	<view class="wrap">
		<js-lang title="user.title"></js-lang>
		<view class="header">
			<view class="userinfo">
				<view class="image" @click="navTo('info')"><image :src="avatarUrl"></image></view>
				<view class="info">
					<view class="username">{{ vuex_user.username || $t('login.noLogin') }}</view>
					<view class="realname">{{ vuex_user.realname || $t('login.noLogin') }}</view>
				</view>
			</view>
			<view class="logout"><u-button type="success" shape="circle" size="mini" @click="logout">{{$t('login.logoutButton')}}</u-button></view>
		</view>
		<!-- <view class="toolbar">
			<view class="box">
				<navigator class="item" hover-class="hover" url="info">
					<view class="icon"><u-icon class="u-icon" :style="{ color: '#ea9a44' }" name="account"></u-icon></view>
					<text class="label">个人信息</text>
				</navigator>
				<navigator class="item" hover-class="hover" url="help">
					<view class="icon"><u-icon class="u-icon" :style="{ color: '#a571fd' }" name="question-circle"></u-icon></view>
					<text class="label">帮助中心</text>
				</navigator>
				<navigator class="item" hover-class="hover" url="info">
					<view class="icon"><u-icon class="u-icon" :style="{ color: '#ea9a44' }" name="account"></u-icon></view>
					<text class="label">个人信息</text>
				</navigator>
				<navigator class="item" hover-class="hover" url="help">
					<view class="icon"><u-icon class="u-icon" :style="{ color: '#a571fd' }" name="question-circle"></u-icon></view>
					<text class="label">帮助中心</text>
				</navigator>
			</view>
		</view> -->
		<view class="u-p-t-10 u-p-b-20">
			<view class="u-m-t-20">
				<u-cell-group>
					<u-cell-item icon="account" :iconSize="iconSize" :iconStyle="{color:'#266bff'}"
						title="个人信息" @click="navTo('info')"></u-cell-item>
					<u-cell-item icon="lock" :iconSize="iconSize" :iconStyle="{ color: '#1bca6a' }"
						title="修改密码" @click="navTo('pwd')"></u-cell-item>
					<u-cell-item icon="question-circle" :iconSize="iconSize" :iconStyle="{ color: '#d99e59' }"
						title="帮助中心" @click="navTo('help')"></u-cell-item>
				</u-cell-group>
			</view>
			<view class="u-m-t-20">
				<u-cell-group>
					<u-cell-item icon="heart" :iconSize="iconSize" :iconStyle="{ color: '#0a1aff' }"
						title="关于我们" @click="navTo('about')"></u-cell-item>
					<u-cell-item icon="kefu-ermai" :iconSize="iconSize" :iconStyle="{ color: '#a571fd' }"
						title="意见反馈" @click="navTo('comment')"></u-cell-item>
					<u-cell-item icon="clock" :iconSize="iconSize" :iconStyle="{ color: '#ff6f27' }"
						title="检查更新" @click="upgrade()"></u-cell-item>
				</u-cell-group>
			</view>
			<view class="u-m-t-20">
				<u-cell-group>
					<u-cell-item icon="setting" :iconSize="iconSize" :iconStyle="{ color: '#1a94ff' }"
						title="系统设置" @click="navTo('setting')"></u-cell-item>
				</u-cell-group>
			</view>
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
			iconSize: 38
		};
	},
	computed: {
		avatarUrl() {
			let url = this.vuex_user.avatar || '/ctxPath/static/images/user1.jpg';
			url = url.replace('/ctxPath/', this.vuex_config.baseUrl + '/');
			return url + '?t=' + new Date().getTime();
		}
	},
	methods: {
		navTo(url) {
			uni.navigateTo({
				url: url
			});
		},
		logout() {
			this.$u.api.logout().then(res => {
				this.$u.toast(res.message);
				setTimeout(() => {
					uni.reLaunch({
						url: '/pages/sys/login/index'
					});
				}, 500);
			});
		},
		upgrade(){
			// #ifdef APP-PLUS
			this.$u.api.upgradeCheck().then(res => {
				if (res.result == 'true'){
					uni.showModal({
						title: '提示',
						content: res.message + '是否下载更新？',
						showCancel: true,
						success: function (res2) {
							if (res2.confirm) {
								plus.runtime.openURL(res.data.apkUrl);
							}
						}
					});
				}else{
					this.$u.toast(res.message);
				}
			});
			// #endif
			// #ifndef APP-PLUS
			this.$u.toast('小程序端或H5端无需检查更新')
			// #endif
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
