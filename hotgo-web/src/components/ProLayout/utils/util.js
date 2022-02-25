import triggerEvent from 'ant-design-vue/es/_util/triggerEvent'
import { inBrowser } from 'ant-design-vue/es/_util/env'

const getComponentFromProp = (instance, prop) => {
  const slots = instance.slots && instance.slots()
  return slots[prop] || instance.props[prop]
}

const isFun = (func) => {
  return typeof func === 'function'
}

// 兼容 0.3.4~0.3.8
export const contentWidthCheck = (contentWidth) => {
  return Object.prototype.toString.call(contentWidth) === '[object Boolean]'
    ? contentWidth === true && 'Fixed' || 'Fluid'
    : contentWidth
}

export const layoutContentWidth = (contentType) => {
  return contentType !== 'Fluid'
}

const themeConfig = {
  daybreak: 'geekblue',
  '#2F54EB': 'geekblue',
  '#1890ff': 'daybreak',
  '#F5222D': 'dust',
  '#FA541C': 'volcano',
  '#FAAD14': 'sunset',
  '#13C2C2': 'cyan',
  '#52C41A': 'green',
  '#722ED1': 'purple'
}

const invertKeyValues = (obj) =>
  Object.keys(obj).reduce((acc, key) => {
    acc[obj[key]] = key
    return acc
  }, {})

/**
 * #1890ff -> daybreak
 * @param val
 */
export function genThemeToString (val) {
  return val && themeConfig[val] ? themeConfig[val] : val
}

/**
 * daybreak-> #1890ff
 * @param val
 */
export function genStringToTheme (val) {
  const stringConfig = invertKeyValues(themeConfig)
  return val && stringConfig[val] ? stringConfig[val] : val
}

export {
  triggerEvent,
  inBrowser,
  getComponentFromProp,
  isFun
}
