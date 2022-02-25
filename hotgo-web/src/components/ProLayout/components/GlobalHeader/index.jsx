import './index.less'

import debounce from 'lodash/debounce'
import PropTypes from 'ant-design-vue/es/_util/vue-types'
import { triggerEvent, inBrowser, isFun } from '../../utils/util'
import 'ant-design-vue/es/icon/style'
import Icon from 'ant-design-vue/es/icon'
import { defaultRenderLogo } from '../SiderMenu/SiderMenu'

export const GlobalHeaderProps = {
  collapsed: PropTypes.bool,
  handleCollapse: PropTypes.func,
  isMobile: PropTypes.bool.def(false),
  fixedHeader: PropTypes.bool.def(false),
  logo: PropTypes.any,
  menuRender: PropTypes.any,
  collapsedButtonRender: PropTypes.any,
  headerContentRender: PropTypes.any,
  rightContentRender: PropTypes.any
}

const defaultRenderCollapsedButton = (h, collapsed) => (
  <Icon type={collapsed ? 'menu-unfold' : 'menu-fold'}/>
)

const GlobalHeader = {
  name: 'GlobalHeader',
  props: GlobalHeaderProps,
  render (h) {
    const { isMobile, logo, rightContentRender, headerContentRender } = this.$props
    const toggle = () => {
      const { collapsed, handleCollapse } = this.$props
      if (handleCollapse) handleCollapse(!collapsed)
      this.triggerResizeEvent()
    }
    const renderCollapsedButton = () => {
      const {
        collapsed,
        collapsedButtonRender = defaultRenderCollapsedButton,
        menuRender
      } = this.$props
      if (collapsedButtonRender !== false && menuRender !== false) {
        return (
          <span class="ant-pro-global-header-trigger" onClick={toggle}>
            {isFun(collapsedButtonRender) && collapsedButtonRender(h, collapsed) || collapsedButtonRender}
          </span>
        )
      }
      return null
    }

    const headerCls = 'ant-pro-global-header'

    return (
      <div class={headerCls}>
        {isMobile && (
          <a class={`${headerCls}-logo`} key="logo" href={'/'}>
            {defaultRenderLogo(h, logo)}
          </a>
        )}
        {renderCollapsedButton()}
        {headerContentRender && (
          <div class={`${headerCls}-content`}>
            {isFun(headerContentRender) && headerContentRender(h, this.$props) || headerContentRender}
          </div>
        )}
        {isFun(rightContentRender) && rightContentRender(h, this.$props) || rightContentRender}
      </div>
    )
  },
  methods: {
    triggerResizeEvent: debounce(() => {
      inBrowser && triggerEvent(window, 'resize')
    })
  },
  beforeDestroy () {
    this.triggerResizeEvent.cancel && this.triggerResizeEvent.cancel()
  }
}

export default GlobalHeader
