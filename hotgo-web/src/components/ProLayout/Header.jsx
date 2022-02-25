import './Header.less'

import 'ant-design-vue/es/layout/style'
import Layout from 'ant-design-vue/es/layout'

import PropTypes from 'ant-design-vue/es/_util/vue-types'
import BaseMenu from './components/RouteMenu/BaseMenu'
import { defaultRenderLogoAntTitle, SiderMenuProps } from './components/SiderMenu/SiderMenu'
import GlobalHeader, { GlobalHeaderProps } from './components/GlobalHeader'
import { VueFragment } from './components'
import { isFun } from './utils/util'

const { Header } = Layout

export const HeaderViewProps = {
  ...GlobalHeaderProps,
  ...SiderMenuProps,
  isMobile: PropTypes.bool.def(false),
  collapsed: PropTypes.bool,
  logo: PropTypes.any,
  hasSiderMenu: PropTypes.bool,
  autoHideHeader: PropTypes.bool,
  menuRender: PropTypes.any,
  headerRender: PropTypes.any,
  rightContentRender: PropTypes.any,
  visible: PropTypes.bool.def(true)
}

const renderContent = (h, props) => {
  const isTop = props.layout === 'topmenu'
  const maxWidth = 1200 - 280 - 120
  const contentWidth = props.contentWidth === 'Fixed'
  const baseCls = 'ant-pro-top-nav-header'
  const { logo, title, theme, isMobile, headerRender, rightContentRender, menuHeaderRender } = props
  const rightContentProps = { theme, isTop, isMobile }
  let defaultDom = <GlobalHeader {...{ props: props }} />
  if (isTop && !isMobile) {
    defaultDom = (
      <div class={[baseCls, theme]}>
        <div class={[`${baseCls}-main`, contentWidth ? 'wide' : '']}>
          {menuHeaderRender && (
            <div class={`${baseCls}-left`}>
              <div class={`${baseCls}-logo`} key="logo" id="logo">
                {defaultRenderLogoAntTitle(h, { logo, title, menuHeaderRender })}
              </div>
            </div>
          )}
          <div class={`${baseCls}-menu`} style={{ maxWidth: `${maxWidth}px`, flex: 1 }}>
            <BaseMenu {...{ props: props }} />
          </div>
          {isFun(rightContentRender) && rightContentRender(h, rightContentProps) || rightContentRender}
        </div>
      </div>
    )
  }
  if (headerRender) {
    return headerRender(h, props)
  }
  return defaultDom
}

const HeaderView = {
  name: 'HeaderView',
  props: HeaderViewProps,
  render (h) {
    const {
      visible,
      isMobile,
      layout,
      collapsed,
      siderWidth,
      fixedHeader,
      hasSiderMenu
    } = this.$props
    const props = this.$props
    const isTop = layout === 'topmenu'

    const needSettingWidth = fixedHeader && hasSiderMenu && !isTop && !isMobile

    const className = {
      'ant-pro-fixed-header': fixedHeader,
      'ant-pro-top-menu': isTop
    }

    // 没有 <></> 暂时代替写法
    return (
      visible ? (
        <VueFragment>
          { fixedHeader && <Header />}
          <Header
            style={{
              padding: 0,
              width: needSettingWidth
                ? `calc(100% - ${collapsed ? 60 : siderWidth}px)`
                : '100%',
              zIndex: 9,
              right: fixedHeader ? 0 : undefined
            }}
            class={className}
          >
            {renderContent(h, props)}
          </Header>
        </VueFragment>
      ) : null
    )
  }
}

export default HeaderView
