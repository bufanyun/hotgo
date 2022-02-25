import { constantRouterMap } from '@/config/router.config'
import { getRouters } from '@/api/menu'
import { BasicLayout } from '@/layouts'

const RouteView = {
  name: 'RouteView',
  render: (h) => h('router-view')
}
// /**
//  * 过滤账户是否拥有某一个权限，并将菜单从加载列表移除
//  *
//  * @param permission
//  * @param route
//  * @returns {boolean}
//  */
// function hasPermission (permission, route) {
//   if (route.meta && route.meta.permission) {
//     let flag = false
//     for (let i = 0, len = permission.length; i < len; i++) {
//       flag = route.meta.permission.includes(permission[i])
//       if (flag) {
//         return true
//       }
//     }
//     return false
//   }
//   return true
// }

/**
 * 单账户多角色时，使用该方法可过滤角色不存在的菜单
 *
 * @param roles
 * @param route
 * @returns {*}
 */
// eslint-disable-next-line
function hasRole(roles, route) {
  if (route.meta && route.meta.roles) {
    return route.meta.roles.includes(roles.id)
  } else {
    return true
  }
}

function filterAsyncRouter (asyncRouterMap) {
  const accessedRouters = asyncRouterMap.filter(route => {
    // if (hasPermission(roles.permissionList, route)) {
    //   if (route.children && route.children.length) {
    //     route.children = filterAsyncRouter(route.children)
    //   }
    //   return true
    // }
      // if (route.children && route.children.length) {
      //   route.children = filterAsyncRouter(route.children)
      // }
    if (route.component) {
      // Layout ParentView 组件特殊处理
      if (route.component === 'Layout') {
        route.component = BasicLayout
      } else {
        route.component = RouteView
      }
    }
    if (route.children != null && route.children && route.children.length) {
      route.children = filterAsyncRouter(route.children)
    }
    return true
  })
  return accessedRouters
}

const permission = {
  state: {
    routers: constantRouterMap,
    addRouters: []
  },
  mutations: {
    SET_ROUTERS: (state, routers) => {
      state.addRouters = routers
      state.routers = constantRouterMap.concat(routers)
    }
  },
  actions: {
    GenerateRoutes ({ commit }) {
      return new Promise(resolve => {
        getRouters().then(res => {
          const accessedRoutes = filterAsyncRouter(res.data)
          // const accessedRoutes = filterAsyncRouter(constantRouterMap)
          // const accessedRoutes = constantRouterMap
          // accessedRoutes.push({ path: '*', redirect: '/404', hidden: true })
          commit('SET_ROUTERS', accessedRoutes)
          resolve(accessedRoutes)
        })
        // const { roles } = data
        // const accessedRouters = filterAsyncRouter(asyncRouterMap, roles)
        // commit('SET_ROUTERS', accessedRouters)
        // resolve()
      })
    }
  }
}

export default permission
