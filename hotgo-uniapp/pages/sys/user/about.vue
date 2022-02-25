<template>
	<view class="wrap">
		<view class="u-p-t-30 u-p-b-30 u-flex u-flex-col u-text-center">
			<u-avatar size="250" src="/static/jeesite/logo200.png"></u-avatar>
			<view class="title">{{vuex_config.productName}}</view>
			<view class="version">{{vuex_config.productVersion}}</view>
		</view>
		<u-cell-group class="form" :border="false">
			<u-cell-item :arrow="true" title="检查更新" @click="upgrade()"></u-cell-item>
			<navigator url="comment" open-type="navigate">
				<u-cell-item :arrow="true" title="意见反馈"></u-cell-item>
			</navigator>
			<navigator url="/pages/common/webview?title=公司首页&url=https://jeesite.com" open-type="navigate">
				<u-cell-item :arrow="true" title="公司首页">https://jeesite.com</u-cell-item>
			</navigator>
			<navigator url="/pages/common/webview?title=服务条款&url=http://s.jeesite.com/" open-type="navigate">
				<u-cell-item :arrow="true" title="服务条款">http://s.jeesite.com/</u-cell-item>
			</navigator>
		</u-cell-group>
		<view class="copyright">
			<view>卓源软件 版权所有</view>
			<view>Copyright &copy; 2021 jeesite.com</view>
			<view>All Rights Reserved</view>
		</view>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	methods: {
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
			this.$u.toast('小程序端或H5端已是最新版，无需检查更新！');
			// #endif
		}
	}
};
</script>
<style lang="scss">
page {
	background-color: #f8f8f8;
}
.title {
	display: flex;
	margin: 30rpx 0;
	font-size: 50rpx;
}
.version {
	margin-bottom: 10rpx;
	font-size: 40rpx;
}
.copyright {
	margin-top: 50rpx;
	text-align: center;
	line-height: 60rpx;
	color: #999;
}
</style>
