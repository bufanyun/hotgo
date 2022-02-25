<template>
  <div :class="wrpCls" style="margin-right:16px">
    <a-space size="middle">
      <a-tooltip>
        <a-dropdown>
          <a class="ant-dropdown-link" style="color:#fff;">
            切换工作台
            <a-icon type="down"/>
          </a>
          <a-menu slot="overlay" class="setUl" style="left:-10px;top:12px; width:200px">
            <a-menu-item
              :key="item.id"
              v-for="(item) in portalConfigs"
              :style="{'background-color': defaultPortal.id === item.id ? '#f0f5ff' : '' }"
              style="position: relative;">
              <a-icon
                v-if="defaultPortal.id === item.id"
                style="left: 10px;"
                type="check"
                :style="{'color': defaultPortal.id === item.id ? '#2f54eb' : '#999999' }"/>
              <a class="homeTit" target="_blank" @click="toIndex(item)">{{ item.name }}</a>
              <a-icon style="right: 8px;" type="delete" target="_blank" @click="toDesignIndex(item,'delete')"/>
              <a-icon style="right: 28px;" type="edit" target="_blank" @click="toDesignIndex(item)"/>
            </a-menu-item>
            <a-menu-divider/>
            <a-menu-item class="menu-operation" key="3" v-if="portalConfigs.length <=15">
              <a target="_blank" @click="toDesignIndex()">
                <a-icon type="plus"/>
                添加工作台</a>
            </a-menu-item>
          </a-menu>
        </a-dropdown>
      </a-tooltip>
      <a-tooltip v-if="userType === '1'">
        <template slot="title">
          控制台
        </template>
        <a-icon type="desktop" @click="toConsole" :style="{ fontSize: '18px'}"/>
      </a-tooltip>
      <a-tooltip @click="toNotice" style="cursor:pointer">
        <template slot="title">
          消息
        </template>
        <a-badge :count="msgCount">
          <a-icon type="sound" :style="{ fontSize: '18px'}"/>
        </a-badge>
      </a-tooltip>
      <a-tooltip>
        <template slot="title">
          换肤
        </template>
        <a-icon type="setting" @click="showColorSetting()" :style="{ fontSize: '18px'}"/>
      </a-tooltip>
      <a-tooltip>
        <template slot="title">
          {{ fullScreen ? '退出全屏' : '切为全屏' }}
        </template>
        <a-icon
          :type="fullScreen ? 'fullscreen-exit' : 'fullscreen'"
          @click="toggleFullScreen"
          :style="{ fontSize: '18px'}"/>
      </a-tooltip>
      <avatar-dropdown :menu="showMenu" :current-user="currentUser" :class="prefixCls"/>
      <!-- 暂只支持中文，国际化可自行扩展 -->
      <select-lang :class="prefixCls"/>
    </a-space>
    <platform-version
      v-if="modalVisible"
      ref="platformVersionModal"
      @close="modalVisible = false"
    />
  </div>
</template>

<script>
  import AvatarDropdown from './AvatarDropdown'
  import SelectLang from '@/components/SelectLang'
  import PlatformVersion from './PlatformVersion'
  import {
    mapGetters
  } from 'vuex'

  export default {
    name: 'RightContent',
    components: {
      AvatarDropdown,
      SelectLang,
      PlatformVersion
    },
    props: {
      prefixCls: {
        type: String,
        default: 'ant-pro-global-header-index-action'
      },
      isMobile: {
        type: Boolean,
        default: () => false
      },
      topMenu: {
        type: Boolean,
        required: true
      },
      theme: {
        type: String,
        required: true
      }
    },
    data() {
      return {
        modalVisible: false,
        showMenu: true,
        showPortalDefined: false,
        currentUser: {},
        fullScreen: false,
        msgCount: 0,
        docUrl: 'https://docs.geekera.cn/AiDex-Antdv/',
        githubUrl: 'https://github.com/fuzui/AiDex-Antdv'
      }
    },
    computed: {
      wrpCls() {
        return {
          'ant-pro-global-header-index-right': true,
          [`ant-pro-global-header-index-${(this.isMobile || !this.topMenu) ? 'light' : this.theme}`]: true
        }
      },
      ...mapGetters([
        'userType',
        'portalConfigs',
        'defaultPortal',
        'sysNoticeList'
      ])
    },
    mounted() {
      setTimeout(() => {
        this.currentUser = {
          name: 'RuoYi'
        }
      }, 1500)
      this.msgCount = this.sysNoticeList.length
    },
    methods: {
      showColorSetting() {
        this.$emit('showSetting')
      },
      toConsole() {
        this.$message.success(
          '尚未实现',
          3
        )
      },
      toNotice() {
        this.$router.push({
          path: '/system/notice/NoticeReadIndex'
        })
        this.msgCount = 0
      },
      toIndex(item) {
        this.$router.push({
          name: 'index',
          params: {
            key: item.id
          }
        })
        if (item.applicationRange === 'U') {
          // 当选中小页时用户自定义时，修改选中小页为默认小页
          this.defaultPortal.id = item.id
        }
        this.$emit('reloadTab', item)
      },
      toDesignIndex(item, type) {
        this.$message.success(
          '尚未实现',
          3
        )
      },
      // 全屏切换
      toggleFullScreen() {
        if (!document.fullscreenElement) {
          document.documentElement.requestFullscreen()
        } else {
          if (document.exitFullscreen) {
            document.exitFullscreen()
          }
        }
        this.fullScreen = !this.fullScreen
      },
      versionInfo() {
        this.modalVisible = true
        this.$nextTick(() => (
          this.$refs.platformVersionModal.showVersion()
        ))
      }
    }
  }
</script>
<style lang="less" scoped="scoped">
  .ant-pro-global-header {
    .anticon {
      vertical-align: middle;
    }
  }

  .ant-modal-confirm-content {
    p {
      height: 15px;
    }
  }

  .setUl {
    .ant-dropdown-menu-item {
      padding: 5px 32px;
      font-size: 12px;
    }

    .ant-dropdown-menu-item > .anticon:first-child {
      font-size: 12px;
    }

    .ant-dropdown-menu-item i {
      position: absolute;
      top: 10px;
      font-size: 12px;
      color: #969696;
    }

    .ant-dropdown-menu-item > a.homeTit {
      width: 150px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      color: #333;
    }

    .menu-operation {
      text-align: center;

      i {
        position: relative;
        top: 0px;
        margin-right: 5px;
      }
    }

    .menu-operation:hover {
      i {
        color: #1890ff
      }
    }
  }
</style>
