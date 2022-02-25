<template>
  <div class="prefixCls">
    <a-tabs v-model="currentTab" @change="handleTabChange">
      <a-tab-pane v-for="v in icons" :key="v.key">
        <span slot="tab" :style="{ fontSize: '10px' }">
          {{ v.title }}
        </span>
        <ul v-if="v.key != 'custom'" style="height: calc(100vh - 196px) ;">
          <li v-for="(icon, key) in iconList" :key="`${v.key}-${key}`" :class="{ 'active': selectedIcon==icon }" @click="handleSelectedIcon(icon)" >
            <a-icon :type="icon" :component="allIcon[icon + 'Icon']" :style="{ fontSize: '24px' }" />
            <span class="anticon-class">
              <span class="ant-badge">
                {{ icon }}
              </span>
            </span>
          </li>
        </ul>
        <ul class="IconList" v-if="v.key == 'custom'" style="height: calc(100vh - 196px) ;">
          <li v-for="(icon, key) in iconList" :key="`${v.key}-${key}`" :class="{ 'active': selectedIcon==icon }" @click="handleSelectedIcon(icon,'1')" >
            <a-icon :component="allIcon[icon + 'Icon']" :type="icon"/>
            <span class="anticon-class">
              <span class="ant-badge">
                {{ icon }}
              </span>
            </span>
          </li>
        </ul>
      </a-tab-pane>
      <a-input-search class="inputsearch" slot="tabBarExtraContent" placeholder="全局搜索图标" @search="onSearchAll" />
    </a-tabs>
  </div>
</template>

<script>
import icons from './icons'
import allCustomIcon from '@/core/icons'
export default {
  name: 'IconSelect',
  props: {
    prefixCls: {
      type: String,
      default: 'ant-pro-icon-selector'
    },
    // eslint-disable-next-line
    value: {
      type: String
    },
    svgIcons: {
      type: Array,
      required: true
    },
    allIcon: {
      type: Object,
      required: true
    }
  },
  data () {
    return {
      selectedIcon: this.value || '',
      currentTab: 'custom',
      icons: icons,
      allCustomIcon,
      iconList: [], // 页面真实展示图标集合，根据搜索条件改变
      currentIconList: [] // 记录当前页面图标集合，不会根据搜索条件改变
    }
  },
  watch: {
    value (val) {
      this.selectedIcon = val
      this.autoSwitchTab()
    }
  },
  created () {
    if (this.value) {
      this.autoSwitchTab()
    }
    const custom = [{
      key: 'custom',
      title: '自定义图标',
      icons: this.svgIcons
    }]
    this.icons = custom.concat(this.icons)
    this.getCurrentIconList()
  },
  methods: {
    handleSelectedIcon (icon, type) {
       this.selectedIcon = icon
       if (allCustomIcon[icon + 'Icon']) {
         // 自定义图标，这里不能根据页签区分是否为自定义图标，因为搜索为全局搜索
         type = '1'
       } else {
         type = '2'
       }
        let copayValue = '<a-icon type="' + icon + '" />'
       if (type === '1') {
         // 自定义图标
         copayValue = '<a-icon type="" :component="allIcon.' + icon + 'Icon"/>'
       }
      var domType = document.createElement('input')
      domType.value = copayValue
      domType.id = 'creatDom'
      document.body.appendChild(domType)
      domType.select() // 选择对象
      document.execCommand('Copy') // 执行浏览器复制命令
      const creatDom = document.getElementById('creatDom')
      creatDom.parentNode.removeChild(creatDom)
      this.$message.success(
              copayValue + ' 复制成功 🎉🎉🎉',
              3
            )
    },
    handleTabChange (activeKey) {
      this.currentTab = activeKey
      this.getCurrentIconList()
    },
    autoSwitchTab () {
      const icons = this.icons
      icons.some(item => item.icons.some(icon => icon === this.value) && (this.currentTab = item.key))
    },
    getCurrentIconList () {
      this.icons.forEach((icon, index) => {
             if (icon.key === this.currentTab) {
               this.iconList = icon.icons
               this.currentIconList = icon.icons
             }
      })
    },
    onSearchAll (text) {
      if (text === '') {
        this.iconList = this.currentIconList
        return
      }
      this.iconList = []
      this.icons.forEach((icon, index) => {
            icon.icons.forEach((icon, index) => {
                   if (icon.toUpperCase().indexOf(text.toUpperCase()) >= 0) {
                       this.iconList.push(icon)
                   }
             })
      })
    }
  }
}
</script>

<style lang="less" scoped>
  @import "../index.less";
  .prefixCls{
    background: #ffffff;
    .inputsearch{
      width: 200px;
      margin-right: 15px;
    }
  }
  ul{
    list-style: none;
    padding: 0;
    overflow-y: scroll;
    li{
      display: inline-block;
      width: 100px;
      height:105px;
      padding: 0;
      margin: 15px 13px;
      text-align: center;
      border-radius: @border-radius-base;

      &:hover, &.active{
        // box-shadow: 0px 0px 5px 2px @primary-color;
        cursor: pointer;
        color: @white;
        background-color: @primary-color;
      }
      i.anticon {
            display: inline-block;
            margin-top: 25px;
            font-size: 24px;
        }
      .anticon-class {
          display: block;
          text-align: center;
          transform: scale(.83);
          margin-top: 12px;
      }
    }
  }
</style>
