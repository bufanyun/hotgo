<template>
	<view class="wrap">
		<u-cell-group :border="false" title="消息提醒">
			<u-cell-item title="接受消息提醒" :arrow="false">
				<u-switch v-model="message" slot="right-icon" class="u-m-l-20"></u-switch>
			</u-cell-item>
			<u-cell-item title="通知栏显示消息详情" :arrow="false">
				<u-switch v-model="messageBar" slot="right-icon" class="u-m-l-20"></u-switch>
				<text slot="label">关闭后，当收到消息的时候，只显示有提示，不显示消息内容。</text>
			</u-cell-item>
		</u-cell-group>
		<u-cell-group :border="false" title="声音与振动">
			<u-cell-item title="收到消息后播放声音或振动" @click="openSettings">
				<text slot="label">前往系统设置中，修改声音与振动</text>
			</u-cell-item>
		</u-cell-group>
		<u-cell-group :border="false" title="软件更新提醒">
			<u-cell-item title="软件更新提醒" :arrow="false">
				<u-switch v-model="upgrade" slot="right-icon" class="u-m-l-20"></u-switch>
				<text slot="label">当本软件有新版本发布时，给予提醒</text>
			</u-cell-item>
		</u-cell-group>
		<view class="u-m-40">
			<u-button type="primary" @click="logout" :hair-line="false">退出登录</u-button>
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
			message: true,
			messageBar: true,
			upgrade: true
		};
	},
	methods: {
		openSettings() {
			// #ifdef APP-PLUS
			uni.getSystemInfo({  
				success(res) {  
					if(res.platform == 'ios'){
						plus.runtime.openURL("app-settings://");
					} else if (res.platform == 'android'){
						var main = plus.android.runtimeMainActivity();  
						var Intent = plus.android.importClass("android.content.Intent");
						var mIntent = new Intent('android.settings.SOUND_SETTINGS');
						main.startActivity(mIntent);
					}
				}
			});
			// #endif
			// #ifndef APP-PLUS
			this.$u.toast('小程序端或H5端已是最新版，无需检查更新！');
			// #endif
		},
		logout() {
			this.$u.api.logout().then(res => {
				this.$u.toast(res.message);
				if (res.result == 'true') {
					let self = this;
					setTimeout(() => {
						uni.reLaunch({
							url: '/pages/sys/login/index'
						});
					}, 500);
				}
			});
		}
	}
};
</script>
<style lang="scss">
@import '../home/index.scss';

page {
	background-color: #f8f8f8;
}

/deep/ .u-cell-title {
	padding: 25rpx 30rpx;
	font-size: 30rpx;
}
</style>
