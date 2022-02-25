<template>
  <pro-layout
    :menus="menus"
    :collapsed="collapsed"
    :mediaQuery="query"
    :isMobile="isMobile"
    :handleMediaQuery="handleMediaQuery"
    :handleCollapse="handleCollapse"
    :multiTab="multiTab"
    :i18nRender="i18nRender"
    v-bind="settings"
  >
    <!-- layout content -->
    <!-- 2021.01.15 默认固定页头，去掉样式paddingTop: fixedHeader ? '64' : '0'  -->
    <template v-slot:headerBottomRender>
      <a-layout-content >
        <multi-tab v-if="multiTab" @reload="reload"></multi-tab>
        <!-- <tabs-view @reload="reload" /> -->
        <transition name="page-transition">
        </transition>
      </a-layout-content>
    </template>
    <!-- 1.0.0+ 版本 pro-layout 提供 API，
          我们推荐使用这种方式进行 LOGO 和 title 自定义
    -->
    <template v-slot:menuHeaderRender>
      <div>
        <img src="~@/assets/logo.png" class="logo" alt="logo">
        <h1>{{ title }}</h1>
      </div>
    </template>
    <template v-slot:headerContentRender>
      <page-header-wrapper :title="false"></page-header-wrapper>
    </template>
    <setting-drawer :settings="settings" @change="handleSettingChange" ref="settingRef"/>
    <template v-slot:rightContentRender>
      <right-content
        :top-menu="settings.layout === 'topmenu'"
        @showSetting="showSetting()"
        :is-mobile="isMobile"
        :theme="settings.theme"
        @reloadTab="reloadTab"
        ref="rightContentRef"
        @designPortal="designPortal"/>
    </template>
    <span v-if="!multiTab">
      <router-view v-if="isRouterAlive && !multiTab" ref="routerViewRef"/> <!--没有开启多页签模式缓存逻辑按照之前逻辑-->
    </span>
    <span v-else>
      <keep-alive v-if="isRouterAlive && multiTab"><!--开启多页签模式所有页面均缓存-->
        <router-view ref="routerViewRef"/>
      </keep-alive>
    </span>
    <template v-slot:footerRender>
      <global-footer />
    </template>
  </pro-layout>
</template>

<script>
import MultiTab from '@/components/MultiTab'
import { SettingDrawer, updateTheme } from '@/components/ProLayout'
import { i18nRender } from '@/locales'
import { mapState } from 'vuex'
import { CONTENT_WIDTH_TYPE, SIDEBAR_TYPE, TOGGLE_MOBILE_TYPE } from '@/store/mutation-types'

import defaultSettings from '@/config/defaultSettings'
import RightContent from '@/components/GlobalHeader/RightContent'
import GlobalFooter from '@/components/GlobalFooter'
import Ads from '@/components/Other/CarbonAds'
import LogoSvg from '../assets/logo.png?inline'
// import tabsView from './modules/tabs-view'

export default {
  name: 'BasicLayout',
  components: {
    SettingDrawer,
    RightContent,
    GlobalFooter,
    LogoSvg,
    Ads,
    MultiTab
    // tabsView
  },
  provide () {
    return {
      reload: this.reload
    }
  },
  data () {
    return {
      isRouterAlive: true,
      showPortalDefined: false,
      // preview.pro.antdv.com only use.
      isProPreviewSite: process.env.VUE_APP_PREVIEW === 'true' && process.env.NODE_ENV !== 'development',
      // end
      multiTab: defaultSettings.multiTab,
      fixedHeader: defaultSettings.fixedHeader,
      // base
      menus: [],
      // 侧栏展开状态
      collapsed: false,
      title: defaultSettings.title,
      settings: {
        // 布局类型
        layout: defaultSettings.layout, // 'sidemenu', 'topmenu'
        // CONTENT_WIDTH_TYPE
        contentWidth: defaultSettings.layout === 'sidemenu' ? CONTENT_WIDTH_TYPE.Fluid : defaultSettings.contentWidth,
        // 主题 'dark' | 'light'
        theme: defaultSettings.navTheme,
        // 主色调
        primaryColor: defaultSettings.primaryColor,
        fixedHeader: defaultSettings.fixedHeader,
        fixSiderbar: defaultSettings.fixSiderbar,
        colorWeak: defaultSettings.colorWeak,

        hideHintAlert: false,
        hideCopyButton: false
      },
      // 媒体查询
      query: {},

      // 是否手机模式
      isMobile: false
    }
  },
  computed: {
    ...mapState({
      // 动态主路由
      mainMenu: state => state.permission.addRouters
    })
  },
  created () {
    const routes = this.mainMenu.find(item => item.path === '/')
    this.menus = (routes && routes.children) || []
    // 处理侧栏展开状态
    this.$watch('collapsed', () => {
      this.$store.commit(SIDEBAR_TYPE, this.collapsed)
    })
    this.$watch('isMobile', () => {
      this.$store.commit(TOGGLE_MOBILE_TYPE, this.isMobile)
    })
  },
  mounted () {
    const userAgent = navigator.userAgent
    if (userAgent.indexOf('Edge') > -1) {
      this.$nextTick(() => {
        this.collapsed = !this.collapsed
        setTimeout(() => {
          this.collapsed = !this.collapsed
        }, 16)
      })
    }

    // first update color
    // TIPS: THEME COLOR HANDLER!! PLEASE CHECK THAT!!
    if (process.env.NODE_ENV !== 'production' || process.env.VUE_APP_PREVIEW === 'true') {
      updateTheme(this.settings.primaryColor)
    }
  },
  methods: {
    i18nRender,
    showSetting () {
      this.$refs.settingRef.setShow(true)
    },
    reloadTab (val) {
       this.$refs.routerViewRef.reloadTab(val)
    },
    designPortal (val, portalConfigs, type) {
       this.$refs.routerViewRef.designPortal(val, portalConfigs, type)
    },
    handleMediaQuery (val) {
      this.query = val
      if (this.isMobile && !val['screen-xs']) {
        this.isMobile = false
        return
      }
      if (!this.isMobile && val['screen-xs']) {
        this.isMobile = true
        this.collapsed = false
        this.settings.contentWidth = CONTENT_WIDTH_TYPE.Fluid
        // this.settings.fixSiderbar = false
      }
    },
    handleCollapse (val) {
      this.collapsed = val
    },
    handleSettingChange ({ type, value }) {
      console.log('type', type, value)
      type && (this.settings[type] = value)
      switch (type) {
        case 'contentWidth':
          this.settings[type] = value
          break
        case 'layout':
          if (value === 'sidemenu') {
            this.settings.contentWidth = CONTENT_WIDTH_TYPE.Fluid
          } else {
            this.settings.fixSiderbar = false
            this.settings.contentWidth = CONTENT_WIDTH_TYPE.Fixed
          }
          break
      }
    },
    reload () {
      this.isRouterAlive = false
      this.$nextTick(function () {
        this.isRouterAlive = true
      })
    }
  }
}
</script>

<style lang="less">
@import "./BasicLayout.less";
</style>
