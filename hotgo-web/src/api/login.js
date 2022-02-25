import request from '@/utils/request'

const userApi = {
  Login: '/login/sign',
  Logout: '/login/logout',
  // get my info
  UserInfo: '/member/info',
  SendSms: '/system/sysSms/sendSms'
}

/**
 * login func
 * @param parameter
 * @returns {*}
 */
export function login(parameter) {
  return request({
    url: userApi.Login,
    method: 'post',
    data: parameter
  })
}

/**
 * login func
 * @param parameter
 * @returns {*}
 */
export function loginBySms(parameter) {
  return request({
    url: '/system/sysSms/loginBySms',
    method: 'post',
    data: parameter
  })
}

export function getInfo() {
  return request({
    url: userApi.UserInfo,
    method: 'get',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

export function logout() {
  return request({
    url: userApi.Logout,
    method: 'post',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

// 获取验证码
export function getCodeImg() {
  return request({
    url: '/login/captcha',
    method: 'get',
    timeout: 20000
  })
}

// 获取手机验证码
export function getSmsCaptcha(parameter) {
  return request({
    url: userApi.SendSms,
    method: 'post',
    data: parameter
  })
}

// 注册方法
export function register(data) {
  return request({
    url: '/register',
    headers: {
      isToken: false
    },
    method: 'post',
    data: data
  })
}
