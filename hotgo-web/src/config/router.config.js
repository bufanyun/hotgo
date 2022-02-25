// eslint-disable-next-line
import { UserLayout } from '@/layouts'

/**
 * 工作台
 */
export const indexRouterMap = [
  {
    path: '/index',
    name: 'index',
    component: 'DashBoard',
    meta: { title: '工作台', keepAlive: true, icon: 'home', noCache: false }
  },
  // {
  //   name: 'index',
  //   path: '/',
  //   component: 'Layout',
  //   meta: { title: '首页', icon: 'home', hideHeader: true },
  //   redirect: '/index',
  //   children: [
  //     {
  //       path: '/index',
  //       name: 'index',
  //       component: 'DashBoard',
  //       meta: { title: '首页', keepAlive: true, icon: 'home', noCache: false }
  //     }
  //   ]
  // },
  {
    path: '/account/center',
    name: 'center',
    component: 'AccountCenter',
    meta: { title: '个人中心', keepAlive: true, noCache: false },
    hidden: true
  },
  // {
  //   path: '/dashboard/console',
  //   name: 'center',
  //   component: 'Console',
  //   meta: { title: '控制台', keepAlive: true, noCache: false },
  //   hidden: true
  // },
  {
    path: '/account/settings',
    name: 'settings',
    component: 'AccountSettings',
    meta: { title: '个人设置', hideHeader: true },
    redirect: '/account/settings/base',
    hidden: true,
    children: [
      {
        path: '/account/settings/base',
        name: 'BaseSettings',
        component: 'BaseSettings',
        hidden: true,
        meta: { title: '基本设置', hidden: true, keepAlive: true, noCache: false }
      },
      {
        path: '/account/settings/security',
        name: 'SecuritySettings',
        component: 'SecuritySettings',
        meta: { title: '安全设置', hidden: true, keepAlive: true, noCache: false }
      }
    ]
  },
  {
    path: '/monitor/job/log',
    name: 'JobLog',
    component: 'JobLog',
    meta: { title: '调度日志', keepAlive: true, noCache: false },
    hidden: true
  },
  {
    path: '/system/notice/NoticeReadIndex',
    name: 'NoticeReadIndex',
    component: 'NoticeReadIndex',
    meta: { title: '通知公告阅读', keepAlive: true, noCache: false },
    hidden: true
  },
  {
    path: '/system/notice/form',
    name: 'NoticeForm',
    component: 'NoticeForm',
    meta: { title: '公告编辑', keepAlive: true, noCache: false },
    hidden: true
  },
  {
    path: '/gen/edit',
    name: 'GenEdit',
    component: 'GenEdit',
    meta: { title: '修改生成配置', keepAlive: true, noCache: false },
    hidden: true
  }
]
/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      }
    ]
  },

  {
    path: '/404',
    name: '404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  },
  {
    path: '/applyLicense',
    name: 'applyLicense',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/applyLicense/ApplyLicense')
  }

]
