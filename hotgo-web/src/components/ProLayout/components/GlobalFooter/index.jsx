import './index.less'

import PropTypes from 'ant-design-vue/es/_util/vue-types'
import { getComponentFromProp, hasProp } from 'ant-design-vue/lib/_util/props-util'

const GlobalFooterProps = {
  links: PropTypes.array,
  copyright: PropTypes.any
}

const GlobalFooter = {
  name: 'GlobalFooter',
  props: GlobalFooterProps,
  render () {
    const copyright = getComponentFromProp(this, 'copyright')
    const links = getComponentFromProp(this, 'links')
    const linksType = hasProp(links)

    return (
      <footer class="ant-pro-global-footer">
        <div class="ant-pro-global-footer-links">
          {linksType && (
            links.map(link => (
              <a
                key={link.key}
                title={link.key}
                target={link.blankTarget ? '_blank' : '_self'}
                href={link.href}
              >
                {link.title}
              </a>
            ))
          ) || links}
        </div>
        {copyright && (
          <div class="ant-pro-global-footer-copyright">{copyright}</div>
        )}
      </footer>
    )
  }
}

export default GlobalFooter
