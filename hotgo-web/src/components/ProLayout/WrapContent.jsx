import PropTypes from 'ant-design-vue/es/_util/vue-types'

import 'ant-design-vue/es/layout/style'
import Layout from 'ant-design-vue/es/layout'
import ConfigProvider from 'ant-design-vue/es/config-provider'
import GridContent from './components/GridContent'

const { Content } = Layout

const WrapContentProps = {
  isChildrenLayout: PropTypes.bool,
  location: PropTypes.any,
  contentHeight: PropTypes.number,
  contentWidth: PropTypes.oneOf(['Fluid', 'Fixed']).def('Fluid')
}

const WrapContent = {
  name: 'WrapContent',
  props: WrapContentProps,
  render (h) {
    const {
      isChildrenLayout,
      contentWidth
    } = this.$props
    return (
      <Content>
        <ConfigProvider
          getPopupContainer={(el, dialogContext) => {
            if (isChildrenLayout) {
              return el.parentNode()
            }
            return document.body
          }}
        >
          <div class="ant-pro-basicLayout-children-content-wrap" >
            <GridContent contentWidth={contentWidth}>{this.$slots.default}</GridContent>
          </div>
        </ConfigProvider>
      </Content>
    )
  }
}

export default WrapContent
