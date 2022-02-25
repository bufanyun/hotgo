import axios from 'axios'
import store from '@/store'
import storage from 'store'
import notification from 'ant-design-vue/es/notification'
import { VueAxios } from './axios'
import { ACCESS_TOKEN } from '@/store/mutation-types'
import errorCode from '@/utils/errorCode'

// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: process.env.VUE_APP_BASE_API,
  // baseURL: 'https://aidex.setworld.net',
  timeout: 20000 // 请求超时时间
})

// 异常拦截处理器
const errorHandler = (error) => {
  console.log('err' + error)
  let { message } = error
  if (message === 'Network Error') {
    message = '后端接口连接异常'
  } else if (message.includes('timeout')) {
    message = '系统接口请求超时'
  } else if (message.includes('Request failed with status code')) {
    message = '系统接口' + message.substr(message.length - 3) + '异常'
  }
  notification.error({
    message: message,
    duration: 5
  })
  //  return Promise.reject(error) //注释该行，否则接口请求失败会一直请求刷新错误数据，这里保证请求一次
}

// request interceptor
request.interceptors.request.use(config => {
  const token = storage.get(ACCESS_TOKEN)
  // 如果 token 存在
  // 让每个请求携带自定义 token 请根据实际情况自行修改
  if (token) {
    config.headers['Authorization'] = 'Bearer ' + token // 让每个请求携带自定义token 请根据实际情况自行修改
    // config.headers['accessAccess-Token'] = token
  }
  // get请求映射params参数
  if (config.method === 'get' && config.params) {
    let url = config.url + '?'
    for (const propName of Object.keys(config.params)) {
      const value = config.params[propName]
      var part = encodeURIComponent(propName) + '='
      // 修改漏洞
      if (value != null && typeof (value) !== 'undefined') {
        if (typeof value === 'object') {
          for (const key of Object.keys(value)) {
            const params = propName + '[' + key + ']'
            var subPart = encodeURIComponent(params) + '='
            url += subPart + encodeURIComponent(value[key]) + '&'
          }
        } else {
          url += part + encodeURIComponent(value) + '&'
        }
      }
    }
    url = url.slice(0, -1)
    config.params = {}
    config.url = url
  }
  return config
}, errorHandler)

// response interceptor
request.interceptors.response.use((res) => {
  // 请求rul
  const requestUrl = res.config.url
  // 未设置状态码则默认成功状态
  const code = res.data.code || 0
  // 获取错误信息
  const msg = errorCode[code] || res.data.message || errorCode['default']

  if (code === 401) {
    notification.open({
      key: 'loginExpireTip',
      message: '系统提示',
      description: '登录状态已过期，请重新登录',
      btn: h => {
        return h(
          'a-button',
          {
            props: {
              type: 'primary',
              size: 'small'
            },
            on: {
              click: () => {
                store.dispatch('Logout').then(() => {
                  location.href = '/index'
                })
              }
            }
          },
          '确认'
        )
      },
      duration: 0,
      onClose: close
    })
  } else if (code === 500) {
    if (requestUrl !== '/login') {
      notification.error({
        message: msg,
        description: msg
      })
    }
  } else if (code !== 0) {
    notification.error({
      message: msg,
      duration: 5
    })
  } else {
    return res.data.data
  }
  return Promise.reject(msg)
}, errorHandler)

const installer = {
  vm: {},
  install(Vue) {
    Vue.use(VueAxios, request)
  }
}

export default request

export {
  installer as VueAxios,
  request as axios
}
