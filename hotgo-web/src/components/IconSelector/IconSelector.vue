<template>
  <div :class="prefixCls">
    <a-tabs v-model="currentTab" @change="handleTabChange">
      <a-tab-pane v-for="v in icons" :key="v.key">
        <span slot="tab" :style="{ fontSize: '10px' }">
          {{ v.title }}
        </span>
        <ul v-if="v.key != 'custom'">
          <li v-for="(icon, key) in v.icons" :key="`${v.key}-${key}`" :class="{ 'active': selectedIcon==icon }" @click="handleSelectedIcon(icon)" >
            <a-icon :type="icon" :style="{ fontSize: '24px' }" />
          </li>
        </ul>
        <ul v-if="v.key == 'custom'">
          <li v-for="(icon, key) in v.icons" :key="`${v.key}-${key}`" :class="{ 'active': selectedIcon==icon }" @click="handleSelectedIcon(icon)" >
            <a-icon :component="allIcon[icon + 'Icon']" :style="{ fontSize: '24px' }"/>
          </li>
        </ul>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import icons from './icons'

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
      icons: icons
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
  },
  methods: {
    handleSelectedIcon (icon) {
      this.selectedIcon = icon
      this.$emit('change', icon)
    },
    handleTabChange (activeKey) {
      this.currentTab = activeKey
    },
    autoSwitchTab () {
      const icons = this.icons
      icons.some(item => item.icons.some(icon => icon === this.value) && (this.currentTab = item.key))
    }
  }
}
</script>

<style lang="less" scoped>
  @import "../index.less";

  ul{
    list-style: none;
    padding: 0;
    overflow-y: scroll;
    height: 250px;

    li{
      display: inline-block;
      padding: @padding-sm;
      margin: 3px 0;
      border-radius: @border-radius-base;

      &:hover, &.active{
        // box-shadow: 0px 0px 5px 2px @primary-color;
        cursor: pointer;
        color: @white;
        background-color: @primary-color;
      }
    }
  }
</style>
