import storage from 'store'
import { login, loginBySms, getInfo, logout } from '@/api/login'
import { ACCESS_TOKEN } from '@/store/mutation-types'

const user = {
  state: {
    token: '',
    name: '',
    userType: '',
    welcome: '',
    avatar: '',
    roles: [],
    portalConfigs: [],
    defaultPortal: {},
    info: {},
    platformVersion: '',
    sysNoticeList: []
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_INFO: (state, info) => {
      state.info = info
    },
    SET_PERMISSIONS: (state, permissions) => {
      state.permissions = permissions
    },
    SET_USER_TYPE: (state, userType) => {
      state.userType = userType
    },
    SET_PORTAL_CONFIG: (state, portalConfigs) => {
      state.portalConfigs = portalConfigs
    },
    SET_DEFAULT_PORTAL: (state, defaultPortal) => {
      state.defaultPortal = defaultPortal
    },
    SET_PLATFORM_VERSION: (state, platformVersion) => {
      state.platformVersion = platformVersion
    },
    SET_NOTICE_LIST: (state, sysNoticeList) => {
      state.sysNoticeList = sysNoticeList
    }
  },

  actions: {
    // 登录
    Login({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        login(userInfo).then(res => {
          storage.set(ACCESS_TOKEN, res.token, 7 * 24 * 60 * 60 * 1000)
          commit('SET_TOKEN', res.token)
          resolve()
        })
          .catch(error => {
            reject(error)
          })
      })
    },
    // 根据验证码登录
    LoginBySms({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        loginBySms(userInfo).then(res => {
          storage.set(ACCESS_TOKEN, res.token, 7 * 24 * 60 * 60 * 1000)
          commit('SET_TOKEN', res.token)
          resolve()
        })
          .catch(error => {
            reject(error)
          })
      })
    },
    // 获取用户信息
    GetInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        getInfo(state.token).then(res => {
          // eslint-disable-next-line no-unused-vars
          const res1 = {
            'msg': '操作成功',
            'code': 200,
            'lincenseInfo': null,
            'userPortalConfig': [{
              'createByName': null,
              'createDeptName': null,
              'importErrInfo': null,
              'id': '73c217ba0fb24945a8faef74eb10d302',
              'searchValue': null,
              'createBy': null,
              'createDept': null,
              'createTime': null,
              'updateBy': null,
              'updateTime': null,
              'updateIp': null,
              'remark': null,
              'version': null,
              'delFlag': '0',
              'handleType': null,
              'params': {},
              'name': '首页',
              'code': '6c297eb4651940edbb45c87c75be00d7',
              'applicationRange': 'U',
              'isDefault': 'Y',
              'resourceId': '1',
              'resourceName': null,
              'systemDefinedId': 'app1',
              'content': '[{"id":"4ae60dd1debe462096698e1da993317a","x":0,"y":0,"w":3,"h":262,"i":"4ae60dd1debe462096698e1da993317a","key":"kuaijierukou","isShowTitle":"N","isAllowDrag":false,"name":"快捷入口","type":"smallPage","url":"dashboard/portal/CommonUse","options":"{\\"titleRequired\\":true,\\"moreUrl\\":\\"\\",\\"refresh\\":1}","moved":false},{"id":"fd7290c27f644234b38d18faf5d75783","x":0,"y":262,"w":3,"h":1370,"i":"fd7290c27f644234b38d18faf5d75783","key":"todo","isShowTitle":"N","isAllowDrag":true,"name":"待办小页","type":"smallPage","url":"dashboard/portal/ToDo","options":"{\\"titleRequired\\":true,\\"moreUrl\\":\\"\\",\\"refresh\\":1}","moved":false}]',
              'sort': null,
              'saveType': null,
              'status': '0',
              'recordLog': true
            }],
            'permissions': ['*:*:*'],
            'sysNoticeList': [],
            'roles': ['admin'],
            'defaultPortalConfig': {
              'createByName': null,
              'createDeptName': null,
              'importErrInfo': null,
              'id': '73c217ba0fb24945a8faef74eb10d302',
              'searchValue': null,
              'createBy': null,
              'createDept': null,
              'createTime': null,
              'updateBy': null,
              'updateTime': null,
              'updateIp': null,
              'remark': null,
              'version': null,
              'delFlag': '0',
              'handleType': null,
              'params': {},
              'name': '首页',
              'code': '6c297eb4651940edbb45c87c75be00d7',
              'applicationRange': 'U',
              'isDefault': 'Y',
              'resourceId': '1',
              'resourceName': null,
              'systemDefinedId': 'app1',
              'content': '[{"id":"4ae60dd1debe462096698e1da993317a","x":0,"y":0,"w":3,"h":262,"i":"4ae60dd1debe462096698e1da993317a","key":"kuaijierukou","isShowTitle":"N","isAllowDrag":false,"name":"快捷入口","type":"smallPage","url":"dashboard/portal/CommonUse","options":"{\\"titleRequired\\":true,\\"moreUrl\\":\\"\\",\\"refresh\\":1}","moved":false},{"id":"fd7290c27f644234b38d18faf5d75783","x":0,"y":262,"w":3,"h":1370,"i":"fd7290c27f644234b38d18faf5d75783","key":"todo","isShowTitle":"N","isAllowDrag":true,"name":"待办小页","type":"smallPage","url":"dashboard/portal/ToDo","options":"{\\"titleRequired\\":true,\\"moreUrl\\":\\"\\",\\"refresh\\":1}","moved":false}]',
              'sort': null,
              'saveType': null,
              'status': '0',
              'recordLog': true
            },
            'user': {
              'createByName': null,
              'createDeptName': null,
              'importErrInfo': null,
              'id': '1',
              'searchValue': null,
              'createBy': 'admin',
              'createDept': null,
              'createTime': '2021-01-30 13:27:43',
              'updateBy': null,
              'updateTime': null,
              'updateIp': null,
              'remark': '管理员',
              'version': null,
              'delFlag': '0',
              'handleType': null,
              'params': {},
              'deptId': '100',
              'name': '管理员',
              'nameEn': null,
              'no': null,
              'userName': 'admin',
              'nickName': '111',
              'userType': '1',
              'email': '1125373330@qq.com',
              'phonenumber': '1125373330a',
              'sex': '1',
              'avatar': '/profile/avatar/2021/11/11/2022/01/14/74359886-5bd9-4ad2-99f6-aab2ad85a8bb.jpeg',
              'birthday': null,
              'nation': null,
              'birthAddress': null,
              'polity': null,
              'title': null,
              'officeTel': '029-03456751111',
              'fax': null,
              'workSpace': null,
              'sort': null,
              'userPinyin': null,
              'salt': null,
              'status': '0',
              'loginIp': '112.24.62.102',
              'loginDate': '2022-01-18 09:56:46',
              'sysDept': {
                'createByName': null,
                'createDeptName': null,
                'importErrInfo': null,
                'id': '100',
                'searchValue': null,
                'createBy': null,
                'createDept': null,
                'createTime': null,
                'updateBy': null,
                'updateTime': null,
                'updateIp': null,
                'remark': null,
                'version': null,
                'delFlag': '0',
                'handleType': null,
                'params': {},
                'parentId': '0',
                'parentIds': null,
                'treeSort': 10,
                'treeSorts': null,
                'treeLevel': null,
                'treeLeaf': null,
                'children': [],
                'deptCode': null,
                'deptName': '集团',
                'leader': '管理员',
                'phone': null,
                'email': null,
                'status': '0',
                'deptFullName': null,
                'deptType': null,
                'address': null,
                'zipCode': null,
                'deptPinyin': null,
                'subtitle': null,
                'searchText': null,
                'parentName': null,
                'parentDeptType': null,
                'recordLog': true
              },
              'sysRoles': [{
                'createByName': null,
                'createDeptName': null,
                'importErrInfo': null,
                'id': '1',
                'searchValue': null,
                'createBy': null,
                'createDept': null,
                'createTime': null,
                'updateBy': null,
                'updateTime': null,
                'updateIp': null,
                'remark': null,
                'version': null,
                'delFlag': null,
                'handleType': null,
                'params': {},
                'roleName': '技术部',
                'roleKey': '007',
                'sort': '4',
                'dataScope': '1',
                'menuCheckStrictly': false,
                'deptCheckStrictly': false,
                'status': '0',
                'flag': false,
                'menuIds': null,
                'deptIds': null,
                'codeOrName': null,
                'option': null,
                'admin': true,
                'recordLog': true
              }],
              'roleIds': null,
              'postIds': null,
              'roleId': null,
              'userNameOrName': null,
              'admin': true,
              'recordLog': true
            }
          }
          const user = res.user
          const avatar = (user.avatar === '' || user.avatar === null) ? require('@/assets/images/profile.jpg') : user.avatar
          if (res.roles && res.roles.length > 0) { // 验证返回的roles是否是一个非空数组
            commit('SET_ROLES', res.roles)
            commit('SET_PERMISSIONS', res.permissions)
          } else {
            commit('SET_ROLES', ['ROLE_DEFAULT'])
          }
          commit('SET_PORTAL_CONFIG', res.userPortalConfig)
          commit('SET_DEFAULT_PORTAL', res.defaultPortalConfig)
          commit('SET_NAME', user.username)
          commit('SET_AVATAR', avatar)
          commit('SET_USER_TYPE', user.type)
          commit('SET_PLATFORM_VERSION', res.lincenseInfo)
          commit('SET_NOTICE_LIST', res.sysNoticeList ? res.sysNoticeList : [])
          resolve(res)
        }).catch(error => {
          console.log('error:' + error)
          reject(error)
        })
      })
    },

    // 登出
    Logout({ commit, state }) {
      return new Promise((resolve) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          commit('SET_PERMISSIONS', [])
          storage.remove(ACCESS_TOKEN)
          resolve()
        }).catch(() => {
          resolve()
        }).finally(() => {
        })
      })
    }

  }
}

export default user
