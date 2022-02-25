<template>
	<view v-if="showBtn" class="js-lang" @tap="switchLang">
		<u-icon size="46" color="warning" :name="lang"></u-icon>
	</view>
</template>
<script>
/**
 * 语言切换组件
 * @property {String} title 顶部导航的标题 i18n 编码
 * @property {Boolean} showBtn 是否显示语言切换按钮
 * @example <js-lang title="login.title" :showBtn="true"></js-lang>
 * @description Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 * @author ThinkGem
 * @version 2021-3-11
 */
export default {
	props: {
		title: {
			type: String,
			default: ''
		},
		showBtn: {
			type: Boolean,
			default: false
		}		
	},
	computed: {
		lang() {
			return this.$i18n.locale == 'zh_CN' ? 'zh' : 'en';
		}
	},
	created(){
		this.setBarTitle();
	},
	methods: {
		switchLang() {
			this.$i18n.locale = this.$i18n.locale == 'zh_CN' ? 'en' : 'zh_CN';
			this.$u.vuex('vuex_locale', this.$i18n.locale);
			this.$u.api.lang({lang: this.vuex_locale});
			this.setBarTitle();
		},
		setBarTitle (){
			uni.setNavigationBarTitle({
				title: this.$t(this.title)
			});
			uni.setTabBarItem({
				index: 0,
				text: this.$t('nav.msg')
			});
			uni.setTabBarItem({
				index: 1,
				text: this.$t('nav.home')
			});
			uni.setTabBarItem({
				index: 2,
				text: this.$t('nav.user')
			});
		}
	}
}
</script>
<style lang="scss" scoped>
.js-lang {
	position: absolute;
	z-index: 10000;
	top: 15px;
	right: 15px;
}
</style>
