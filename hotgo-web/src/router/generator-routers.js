// eslint-disable-next-line
import { getRouters } from '@/api/menu'
import { UserLayout, BasicLayout, BlankLayout, PageView, RouteView } from '@/layouts'
import { indexRouterMap } from '@/config/router.config'
import allIcon from '@/core/icons'
// import { validURL } from '@/utils/validate'
// 前端路由表
const constantRouterComponents = {
  // 基础页面 layout 必须引入
  BasicLayout: BasicLayout,
  BlankLayout: BlankLayout,
  RouteView: RouteView,
  PageView: PageView,
  UserLayout: UserLayout, // 登陆注册页面的通用布局

  // 你需要动态引入的页面组件
  'Index': () => import('@/views/index'),
  // account
  'AccountCenter': () => import('@/views/account/center'),
  'AccountSettings': () => import('@/views/account/settings/index'),
  'BaseSettings': () => import('@/views/account/settings/BaseSetting'),
  'SecuritySettings': () => import('@/views/account/settings/Security'),
  // job log
  'JobLog': () => import('@/views/monitor/job/log'),
  // 公告新增修改
  // 修改生成配置
  'GenEdit': () => import('@/views/tool/gen/modules/GenEdit'),
  'DashBoard': () => import('@/views/dashboard/index'),
  'NoticeReadIndex': () => import('@/views/system/notice/NoticeReadIndex')
}

// 前端未找到页面路由（固定不用改）
// const notFoundRouter = {
//   path: '*', redirect: '/404', hidden: true
// }
// 根级菜单
const rootRouter = {
  key: '',
  name: 'index',
  path: '',
  component: 'BasicLayout',
  redirect: '/index',
  meta: {
    title: '工作台'
  },
  children: []
}

/**
 * 动态生成菜单
 * @param token
 * @returns {Promise<Router>}
 */
export const generatorDynamicRouter = (token) => {
  return new Promise((resolve, reject) => {
    // 向后端请求路由数据
    getRouters().then(res => {
      // eslint-disable-next-line no-unused-vars
      const res1 = [
          {
            'name': 'Org',
            'path': '/org',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '组织管理', 'icon': 'calculator', 'noCache': false, 'remark': '' },
            'children': [{
              'name': '小欣欣',
              'path': '小欣欣',
              'hidden': false,
              'component': 'ParentView',
              'isFrame': '1',
              'meta': { 'title': '学校', 'icon': 'Alipay', 'noCache': false, 'remark': null }
            }, {
              'name': 'SysAuth',
              'path': 'sysAuth',
              'hidden': false,
              'component': 'system/role/SysRoleAuth',
              'isFrame': '1',
              'meta': { 'title': '菜单授权', 'icon': 'api', 'noCache': false, 'remark': '' }
            }, {
              'name': 'SysPost',
              'path': 'sysPost',
              'hidden': false,
              'component': '11/syspost/index',
              'isFrame': '1',
              'meta': { 'title': '岗位信息', 'icon': 'DragColumn', 'noCache': false, 'remark': '' }
            }]
          },
          {
            'name': 'Auth',
            'path': '/auth',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '权限管理', 'icon': 'solution', 'noCache': false, 'remark': '' },
            'children': [{
              'name': 'Role',
              'path': 'role',
              'hidden': false,
              'component': 'system/role/QueryList',
              'isFrame': '1',
              'meta': { 'title': '角色管理', 'icon': 'contacts', 'noCache': false, 'remark': '维护平台各角色数据以及权限分配.' }
            }]
          }, {
            'name': 'Tool',
            'path': '/tool',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '系统工具', 'icon': 'appstore', 'noCache': false, 'remark': '系统工具目录' },
            'children': [{
              'name': 'Build',
              'path': 'build',
              'hidden': false,
              'component': 'tool/build/index',
              'isFrame': '1',
              'meta': { 'title': '表单构建', 'icon': 'bars', 'noCache': false, 'remark': '表单构建菜单' }
            }, {
              'name': 'Gen',
              'path': 'gen',
              'hidden': false,
              'component': 'tool/gen/index',
              'isFrame': '1',
              'meta': { 'title': '代码生成', 'icon': 'code', 'noCache': false, 'remark': '代码生成菜单' }
            }, {
              'name': 'Swagger',
              'path': 'swagger',
              'hidden': false,
              'component': 'tool/swagger/index',
              'isFrame': '1',
              'meta': { 'title': '系统接口', 'icon': 'api', 'noCache': false, 'remark': '系统接口菜单' }
            }, {
              'name': 'Template',
              'path': 'template',
              'hidden': false,
              'component': 'tool/gen/genconfigtemplate/index',
              'isFrame': '1',
              'meta': { 'title': '模板配置', 'icon': 'picture', 'noCache': false, 'remark': '' }
            }, {
              'name': 'BaseTreeTable',
              'path': 'baseTreeTable',
              'hidden': false,
              'component': '11/basetreetable/index',
              'isFrame': '1',
              'meta': { 'title': '树基础', 'icon': '#', 'noCache': false, 'remark': '' }
            }, {
              'name': 'BaseTable',
              'path': 'baseTable',
              'hidden': false,
              'component': 'test/basetable/index',
              'isFrame': '1',
              'meta': { 'title': '基础', 'icon': '#', 'noCache': false, 'remark': '' }
            }]
          }, {
            'name': 'SysSetting',
            'path': '/sysSetting',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '系统设置', 'icon': 'setting', 'noCache': false, 'remark': '' },
            'children': [{
              'name': 'Menu',
              'path': 'menu',
              'hidden': false,
              'component': 'system/menu/MenuIndex',
              'isFrame': '1',
              'meta': { 'title': '菜单管理', 'icon': 'bars', 'noCache': false, 'remark': '平台所有菜单维护' }
            }, {
              'name': 'Dict',
              'path': 'dict',
              'hidden': false,
              'component': 'system/dict/DictIndex',
              'isFrame': '1',
              'meta': { 'title': '字典管理', 'icon': 'read', 'noCache': false, 'remark': '字典管理菜单' }
            }, {
              'name': 'Config',
              'path': 'config',
              'hidden': false,
              'component': 'system/config/ConfigIndex',
              'isFrame': '1',
              'meta': { 'title': '参数设置', 'icon': 'code', 'noCache': false, 'remark': '参数设置菜单' }
            }, {
              'name': 'User',
              'path': 'user',
              'hidden': false,
              'component': 'system/user/SysUserIndex',
              'isFrame': '1',
              'meta': { 'title': '用户管理', 'icon': 'peoples', 'noCache': false, 'remark': null }
            }]
          }, {
            'name': 'Log',
            'path': '/log',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '日志管理', 'icon': 'copy', 'noCache': false, 'remark': '日志管理菜单' },
            'children': [{
              'name': 'Operlog',
              'path': 'operlog',
              'hidden': false,
              'component': 'monitor/operlog/OperlogIndex',
              'isFrame': '1',
              'meta': { 'title': '操作日志', 'icon': 'form', 'noCache': false, 'remark': '操作日志菜单' }
            }, {
              'name': 'LoginLog',
              'path': 'loginLog',
              'hidden': false,
              'component': 'monitor/loginlog/LoginLogIndex',
              'isFrame': '1',
              'meta': { 'title': '登录日志', 'icon': 'loginLog', 'noCache': false, 'remark': '登录日志菜单' }
            }, {
              'name': 'Joblog',
              'path': 'joblog',
              'hidden': false,
              'component': 'monitor/job/log',
              'isFrame': '1',
              'meta': { 'title': '调度日志', 'icon': 'bug', 'noCache': false, 'remark': '' }
            }]
          }, {
            'name': 'Monitor',
            'path': '/monitor',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '系统监控', 'icon': 'fund', 'noCache': false, 'remark': '系统监控目录' },
            'children': [{
              'name': 'Online',
              'path': 'online',
              'hidden': false,
              'component': 'monitor/online/index',
              'isFrame': '1',
              'meta': { 'title': '在线用户', 'icon': 'dot-chart', 'noCache': false, 'remark': '在线用户菜单' }
            }, {
              'name': 'Job',
              'path': 'job',
              'hidden': false,
              'component': 'monitor/job/index',
              'isFrame': '1',
              'meta': { 'title': '定时任务', 'icon': 'bar-chart', 'noCache': false, 'remark': '定时任务菜单' }
            }, {
              'name': 'Druid',
              'path': 'druid',
              'hidden': false,
              'component': 'monitor/druid/index',
              'isFrame': '1',
              'meta': { 'title': '数据监控', 'icon': 'dashboard', 'noCache': false, 'remark': '数据监控菜单' }
            }, {
              'name': 'Server',
              'path': 'server',
              'hidden': false,
              'component': 'monitor/server/index',
              'isFrame': '1',
              'meta': { 'title': '服务监控', 'icon': 'pie-chart', 'noCache': false, 'remark': '服务监控菜单' }
            }, {
              'name': 'Cache',
              'path': 'cache',
              'hidden': false,
              'component': 'monitor/cache/index',
              'isFrame': '1',
              'meta': { 'title': '缓存监控', 'icon': 'box-plot', 'noCache': false, 'remark': '缓存监控菜单' }
            }, {
              'name': 'CacheList',
              'path': 'cacheList',
              'hidden': false,
              'component': 'monitor/cache/indexCacheList',
              'isFrame': '1',
              'meta': { 'title': '缓存列表', 'icon': 'dashboardNew', 'noCache': false, 'remark': null }
            }]
          }, {
            'name': 'SysApp',
            'path': '/sysApp',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '系统应用', 'icon': 'laptop', 'noCache': false, 'remark': '' },
            'children': [{
              'name': 'Notice',
              'path': 'notice',
              'hidden': false,
              'component': 'system/notice/NoticeIndex',
              'isFrame': '1',
              'meta': { 'title': '通知公告', 'icon': 'notification', 'noCache': false, 'remark': '通知公告菜单' }
            }, {
              'name': 'Baidu.com',
              'path': 'baidu.com',
              'hidden': false,
              'component': 'ParentView',
              'isFrame': '0',
              'meta': { 'title': '测试', 'icon': 'QRcode', 'noCache': false, 'remark': null }
            }]
          }, {
            'path': '/',
            'hidden': false,
            'component': 'Layout',
            'isFrame': '1',
            'meta': { 'title': '部门管理', 'icon': 'cluster', 'noCache': false, 'remark': '部门管理菜单' },
            'children': [{
              'name': 'Dept',
              'path': 'dept',
              'hidden': false,
              'component': 'system/dept/SysDeptIndex',
              'isFrame': '1',
              'meta': { 'title': '部门管理', 'icon': 'cluster', 'noCache': false, 'remark': '部门管理菜单' }
            }]
          }, {
            'name': '字典',
            'path': '/SysDictType',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '字典', 'icon': 'redis', 'noCache': false, 'remark': null },
            'children': [{
              'name': 'SysDictType',
              'path': 'sysDictType',
              'hidden': false,
              'component': 'sysDictType/sysdicttype/index',
              'isFrame': '1',
              'meta': { 'title': '字典类型', 'icon': 'Sina', 'noCache': false, 'remark': '' }
            }]
          }, {
            'name': '测试菜单-name',
            'path': '/test_path',
            'hidden': false,
            'redirect': 'noRedirect',
            'component': 'Layout',
            'alwaysShow': true,
            'isFrame': '1',
            'meta': { 'title': '测试菜单-title', 'icon': 'setting', 'noCache': false, 'remark': '' },
            'children': [{
              'name': 'Temp_children_name',
              'path': 'temp_children_path',
              'hidden': false,
              'component': 'test/temp/home',
              'isFrame': '1',
              'meta': { 'title': '测试页面', 'icon': 'code', 'noCache': false, 'remark': '参数设置菜单' }
            }]
          }]

      // console.log('res2:' + res1.length)
      // console.log('res.data:' + JSON.stringify(res))
      const menuNav = []
      rootRouter.children = indexRouterMap.concat(res)
      menuNav.push(rootRouter)
      const routers = generator(menuNav)
      // console.log('routers', routers)
      // routers.push(notFoundRouter)
      resolve(routers)
    }).catch(err => {
      reject(err)
    })
  })
}

/**
 * 格式化树形结构数据 生成 vue-router 层级路由表
 *
 * @param routerMap
 * @param parent
 * @returns {*}
 */
export const generator = (routerMap, parent) => {
  return routerMap.map(item => {
    const { title, show, hideChildren, hiddenHeaderContent, hidden, icon, noCache } = item.meta || {}
    if (item.component) {
      // Layout ParentView 组件特殊处理
      if (item.component === 'Layout') {
        item.component = 'RouteView'
      } else if (item.component === 'ParentView') {
        // 三级菜单处理
        item.component = 'RouteView'
        item.path = '/' + item.path
      }
    }
    if (item.path) {
      // item.path = '/' + item.path
    }
    if (item.isFrame === '2') {
      item.target = '_blank'
    } else {
      item.target = ''
    }
    const currentRouter = {
      // 如果路由设置了 path，则作为默认 path，否则 路由地址 动态拼接生成如 /dashboard/workplace
      path: item.path || `${parent && parent.path || ''}/${item.path}`,
      // 路由名称，建议唯一
      name: item.name || item.key || '',
      // 该路由对应页面的 组件(动态加载)
      component: (constantRouterComponents[item.component || item.key]) || (() => import(`@/views/${item.component}`)),
      hidden: item.hidden,
      // meta: 页面标题, 菜单图标, 页面权限(供指令权限用，可去掉)
      meta: {
        title: title,
        icon: allIcon[icon + 'Icon'] || icon,
        iconStr: icon === null ? 'profile' : icon,
        hiddenHeaderContent: hiddenHeaderContent,
        // 目前只能通过判断path的http链接来判断是否外链，适配若依
        // target: validURL(item.path) ? '_blank' : '',
        target: item.target,
        permission: item.name,
        keepAlive: noCache,
        hidden: hidden,
        remark: item.meta.remark
      },
      redirect: item.redirect
    }
    // 是否设置了隐藏菜单
    if (show === false) {
      currentRouter.hidden = true
    }
    // 适配若依，若依为缩写路径，而antdv-pro的pro-layout要求每个路径需为全路径
    if (!constantRouterComponents[item.component || item.key]) {
      currentRouter.path = `${parent && parent.path || ''}/${item.path}`
    }
    // 是否设置了隐藏子菜单
    if (hideChildren) {
      currentRouter.hideChildrenInMenu = true
    }
    // 是否有子菜单，并递归处理，并将父path传入
    if (item.children && item.children.length > 0) {
      // Recursion
      currentRouter.children = generator(item.children, currentRouter)
    }
    // console.log('currentRouter:' + JSON.stringify(currentRouter))
    // console.log('======================')
    return currentRouter
  })
}
