<script>
/**
 * Copyright (c) 2013-Now http://jeesite.com All rights reserved.
 */
export default {
	onLaunch() {
		// 国际化，设置当前语言
		if (this.vuex_locale){
			this.$i18n.locale = this.vuex_locale;
			this.$u.api.lang({lang: this.vuex_locale});
		}
		// 设置底部导航栏角标
		uni.setTabBarBadge({
		  index: 0,
		  text: '3'
		});
		// uni.removeTabBarBadge({
		//   index: 0
		// });
	}
}
</script>
<style lang="scss">
@import "uview-ui/index.scss";
@import "pages/common/jeesite.scss";
</style>