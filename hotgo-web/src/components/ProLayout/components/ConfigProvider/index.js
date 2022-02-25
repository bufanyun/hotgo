import PropTypes from 'ant-design-vue/es/_util/vue-types'

const ProConfigProviderProps = {
  i18nRender: PropTypes.oneOfType([PropTypes.func, PropTypes.bool]).def(false),
  contentWidth: PropTypes.oneOf(['Fluid', 'Fixed']).def('Fluid'),
  breadcrumbRender: PropTypes.func
}

const ConfigProvider = {
  name: 'ProConfigProvider',
  props: ProConfigProviderProps,
  provide () {
    const _self = this
    return {
      locale: _self.$props.i18nRender,
      contentWidth: _self.$props.contentWidth,
      breadcrumbRender: _self.$props.breadcrumbRender
    }
  },
  render () {
    const { $scopedSlots } = this
    const children = this.children || $scopedSlots.default
    return children()
  }
}

export default ConfigProvider
