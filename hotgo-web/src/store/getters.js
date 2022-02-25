const getters = {
  isMobile: state => state.app.isMobile,
  lang: state => state.app.lang,
  theme: state => state.app.theme,
  color: state => state.app.color,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  nickname: state => state.user.name,
  name: state => state.user.name,
  userType: state => state.user.userType,
  portalConfigs: state => state.user.portalConfigs,
  defaultPortal: state => state.user.defaultPortal,
  platformVersion: state => state.user.platformVersion,
  sysNoticeList: state => state.user.sysNoticeList,
  welcome: state => state.user.welcome,
  roles: state => state.user.roles,
  permissions: state => state.user.permissions,
  userInfo: state => state.user.info,
  addRouters: state => state.permission.addRouters,
  multiTab: state => state.app.multiTab
}

export default getters
