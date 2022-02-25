import './index.less'

import 'ant-design-vue/es/drawer/style'
import Drawer from 'ant-design-vue/es/drawer'
import SiderMenu, { SiderMenuProps } from './SiderMenu'

const SiderMenuWrapper = {
  name: 'SiderMenuWrapper',
  model: {
    prop: 'collapsed',
    event: 'collapse'
  },
  props: SiderMenuProps,
  render (h) {
    const {
      layout,
      isMobile,
      collapsed
    } = this
    const isTopMenu = layout === 'topmenu'
    const handleCollapse = (e) => {
      this.$emit('collapse', true)
    }
    return isMobile ? (
      <Drawer
        class="ant-pro-sider-menu"
        visible={!collapsed}
        placement="left"
        maskClosable
        getContainer={null}
        onClose={handleCollapse}
        bodyStyle={{
          padding: 0,
          height: '100vh'
        }}
      >
        <SiderMenu {...{ props: { ...this.$props, collapsed: isMobile ? false : collapsed } } } />
      </Drawer>
    ) : !isTopMenu && (
      <SiderMenu class="ant-pro-sider-menu" {...{ props: this.$props }} />
    )
  }
}

SiderMenuWrapper.install = function (Vue) {
  Vue.component(SiderMenuWrapper.name, SiderMenuWrapper)
}

export {
  SiderMenu,
  SiderMenuProps
}

export default SiderMenuWrapper
