import './index.less'

import PropTypes from 'ant-design-vue/es/_util/vue-types'
import { layoutContentWidth } from '../../utils/util'

const GridContent = {
  name: 'GridContent',
  functional: true,
  props: {
    children: PropTypes.any,
    contentWidth: PropTypes.oneOf(['Fluid', 'Fixed']).def('Fluid')
  },
  render (h, content) {
    const { contentWidth } = content.props
    const children = content.children

    const propsContentWidth = layoutContentWidth(contentWidth)
    const classNames = {
      'ant-pro-grid-content': true,
      'wide': propsContentWidth
    }

    return <div class={classNames}>{children}</div>
  }
}

export default GridContent
