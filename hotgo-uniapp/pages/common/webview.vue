<template>
	<view>
		<web-view :webview-styles="webviewStyles" :src="webviewUrl"></web-view>
	</view>
</template>
<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	data() {
		return {
			webviewUrl: '',
			webviewStyles: {
				progress: {
					color: '#FF7200'
				}
			}
		};
	},
	onLoad(params) {
		this.webviewUrl = decodeURIComponent(params.url);
		if (params.title != ''){
			uni.setNavigationBarTitle({
				title: params.title
			})
		}
	}
};
</script>
<style lang="scss">
	
</style>
