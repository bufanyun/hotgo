export enum ApiEnum {
  // api前缀
  Prefix = '/api',

  // 基础
  SiteRegister = '/site/register', // 账号注册
  SiteAccountLogin = '/site/accountLogin', // 账号登录
  SiteMobileLogin = '/site/mobileLogin', // 手机号登录
  SiteLoginConfig = '/site/loginConfig', // 登录配置
  SiteLogout = '/site/logout', // 注销
  SiteConfig = '/site/config', // 配置信息

  // 用户
  MemberInfo = '/member/info', // 登录用户信息

  // 角色
  RoleDynamic = '/role/dynamic', // 动态路由
}
