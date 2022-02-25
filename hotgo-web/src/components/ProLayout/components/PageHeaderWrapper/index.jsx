import './index.less'

import PropTypes from 'ant-design-vue/es/_util/vue-types'
import { isArray } from 'ant-design-vue/lib/_util/vue-types/utils'
import GridContent from '../GridContent'
import 'ant-design-vue/es/page-header/style'
import PageHeader, { PageHeaderProps } from 'ant-design-vue/es/page-header'
import 'ant-design-vue/es/tabs/style'
import Tabs from 'ant-design-vue/es/tabs'
import { getComponentFromProp } from 'ant-design-vue/lib/_util/props-util'

const prefixedClassName = 'ant-pro-page-header-wrap'

const PageHeaderTabConfig = {
  tabList: PropTypes.array,
  tabActiveKey: PropTypes.string,
  tabProps: PropTypes.object,
  tabChange: PropTypes.func
}

const PageHeaderWrapperProps = {
  ...PageHeaderTabConfig,
  ...PageHeaderProps,
  title: PropTypes.oneOfType([PropTypes.string, PropTypes.bool]),
  content: PropTypes.any,
  extraContent: PropTypes.any,
  pageHeaderRender: PropTypes.func,
  breadcrumb: PropTypes.oneOfType([PropTypes.object, PropTypes.bool]).def(true),
  back: PropTypes.func,

  // only use `pro-layout` in children
  i18nRender: PropTypes.oneOfType([PropTypes.func, PropTypes.bool]).def(false)
}

const defaultI18nRender = (t) => t

const useContext = (route) => {
  return route && {
    ...route.meta
  } || null
}

const noop = () => {
}

// TODO :: tabList tab 支持图标 优化
const renderFooter = (h, tabConfigProps, i18nRender) => {
  const {
    tabList,
    tabActiveKey,
    tabChange: onTabChange,
    tabBarExtraContent,
    tabProps
  } = tabConfigProps
  return tabList && tabList.length > 0 && (
    <Tabs
      class={`${prefixedClassName}-tabs`}
      activeKey={tabActiveKey}
      onChange={key => {
        if (onTabChange) {
          onTabChange(key)
        }
      }}
      tabBarExtraContent={tabBarExtraContent}
      {...tabProps}
    >
      {tabList.map(item => (
        <Tabs.TabPane {...item} tab={i18nRender(item.tab)} key={item.key}/>
      ))}
    </Tabs>
  )
}

const renderPageHeader = (h, content, extraContent) => {
  if (!content && !extraContent) {
    return null
  }
  return (
    <div class={`${prefixedClassName}-detail`}>
      <div class={`${prefixedClassName}-main`}>
        <div class={`${prefixedClassName}-row`}>
          {content && (
            <div class={`${prefixedClassName}-content`}>{content}</div>
          )}
          {extraContent && (
            <div class={`${prefixedClassName}-extraContent`}>
              {extraContent}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

const defaultPageHeaderRender = (h, props, pageMeta, i18nRender) => {
  const {
    title: propTitle,
    tags,
    content,
    pageHeaderRender,
    extra,
    extraContent,
    breadcrumb,
    back: handleBack,
    ...restProps
  } = props
  if (pageHeaderRender) {
    return pageHeaderRender({ ...props })
  }
  let pageHeaderTitle = propTitle
  if (!propTitle && propTitle !== false) {
    pageHeaderTitle = pageMeta.title
  }
  // title props 不是 false 且不是 array 则直接渲染 title
  // 反之认为是 VNode, 作为 render 参数直接传入到 PageHeader
  const title = isArray(pageHeaderTitle)
    ? pageHeaderTitle
    : pageHeaderTitle && i18nRender(pageHeaderTitle)
  const tabProps = {
    breadcrumb,
    extra,
    tags,
    title,
    footer: renderFooter(h, restProps, i18nRender)
  }
  if (!handleBack) {
    tabProps.backIcon = false
  }

  return (
    <PageHeader {...{ props: tabProps }} onBack={handleBack || noop}>
      {renderPageHeader(h, content, extraContent)}
    </PageHeader>
  )
  // return
}

const PageHeaderWrapper = {
  name: 'PageHeaderWrapper',
  props: PageHeaderWrapperProps,
  inject: ['locale', 'contentWidth', 'breadcrumbRender'],
  render (h) {
    const { $route, $listeners } = this
    const children = this.$slots.default
    const title = getComponentFromProp(this, 'title')
    const tags = getComponentFromProp(this, 'tags')
    const content = getComponentFromProp(this, 'content')
    const extra = getComponentFromProp(this, 'extra')
    const extraContent = getComponentFromProp(this, 'extraContent')

    const pageMeta = useContext(this.$props.route || $route)
    const i18n = this.$props.i18nRender || this.locale || defaultI18nRender
    const contentWidth = this.$props.contentWidth || this.contentWidth || false
    // 当未设置 back props 或未监听 @back，不显示 back
    // props 的 back 事件优先级高于 @back，需要注意
    const onBack = this.$props.back || $listeners.back
    const back = onBack && (() => {
      // this.$emit('back')
      // call props back func
      onBack && onBack()
    }) || undefined

    const onTabChange = this.$props.tabChange
    const tabChange = (key) => {
      this.$emit('tabChange', key)
      onTabChange && onTabChange(key)
    }

    let breadcrumb = {}
    const propsBreadcrumb = this.$props.breadcrumb
    if (propsBreadcrumb === true) {
      const routes = $route.matched.concat().map(route => {
        return {
          path: route.path,
          breadcrumbName: i18n(route.meta.title),
          redirect: route.redirect
        }
      })

      const defaultItemRender = ({ route, params, routes, paths, h }) => {
        return (route.redirect === 'noRedirect' || routes.indexOf(route) === routes.length - 1) && (
          <span>{route.breadcrumbName}</span>
        ) || (
          <router-link to={{ path: route.path || '/', params }}>{route.breadcrumbName}</router-link>
        )
      }

      // If custom breadcrumb render undefined
      // use default breadcrumb..
      const itemRender = this.breadcrumbRender || defaultItemRender
      routes.splice(0, 1)
      breadcrumb = { props: { routes, itemRender } }
    } else {
      breadcrumb = propsBreadcrumb || null
    }

    const props = {
      ...this.$props,
      title,
      tags,
      content,
      extra,
      extraContent,
      breadcrumb,
      tabChange,
      back
    }

    return (
      <div class="ant-pro-page-header-wrap">
        <div class={`${prefixedClassName}-page-header-warp`}>
          <GridContent>{defaultPageHeaderRender(h, props, pageMeta, i18n)}</GridContent>
        </div>
        {children ? (
          <GridContent contentWidth={contentWidth}>
            <div class={`${prefixedClassName}-children-content`}>
              {children}
            </div>
          </GridContent>
        ) : null}
      </div>
    )
  }
}

PageHeaderWrapper.install = function (Vue) {
  Vue.component(PageHeaderWrapper.name, PageHeaderWrapper)
  Vue.component('page-container', PageHeaderWrapper)
}

export default PageHeaderWrapper
