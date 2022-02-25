import './ThemeColor.less'

import PropTypes from 'ant-design-vue/es/_util/vue-types'
import { genThemeToString } from '../../utils/util'
import 'ant-design-vue/es/tooltip/style'
import Tooltip from 'ant-design-vue/es/tooltip'
import 'ant-design-vue/es/icon/style'
import Icon from 'ant-design-vue/es/icon'

const baseClassName = 'theme-color'

export const TagProps = {
  color: PropTypes.string,
  check: PropTypes.bool
}

const Tag = {
  props: TagProps,
  functional: true,
  render (h, content) {
    const { props: { color, check }, data } = content
    return (
      <div {...data} style={{ backgroundColor: color }}>
        { check ? <Icon type="check" /> : null }
      </div>
    )
  }
}

export const ThemeColorProps = {
  colors: PropTypes.array,
  title: PropTypes.string,
  value: PropTypes.string,

  i18nRender: PropTypes.oneOfType([PropTypes.func, PropTypes.bool]).def(false)
}

const ThemeColor = {
  props: ThemeColorProps,
  inject: ['locale'],
  render (h) {
    const { title, value, colors = [] } = this
    const i18n = this.$props.i18nRender || this.locale
    const handleChange = (key) => {
      this.$emit('change', key)
    }

    return (
      <div class={baseClassName} ref={'ref'}>
        <h3 class={`${baseClassName}-title`}>{title}</h3>
        <div class={`${baseClassName}-content`}>
          {colors.map(item => {
            const themeKey = genThemeToString(item.key)
            const check = value === item.key || genThemeToString(value) === item.key
            return (
              <Tooltip
                key={item.color}
                title={themeKey ? i18n(`app.setting.themecolor.${themeKey}`) : item.key}
              >
                <Tag
                  class={`${baseClassName}-block`}
                  color={item.color}
                  check={check}
                  onClick={() => handleChange(item.key)}
                />
              </Tooltip>
            )
          })}
        </div>
      </div>
    )
  }
}

export default ThemeColor
