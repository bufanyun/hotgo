import { defineStore } from 'pinia';
import { createStorage, storage } from '@/utils/Storage';
import { store } from '@/store';
import {
  ACCESS_TOKEN,
  CURRENT_CONFIG,
  CURRENT_DICT,
  CURRENT_LOGIN_CONFIG,
  CURRENT_USER,
  IS_LOCKSCREEN,
} from '@/store/mutation-types';
import { ResultEnum } from '@/enums/httpEnum';
import {
  getConfig,
  getLoginConfig,
  getUserInfo,
  login,
  logout,
  mobileLogin,
} from '@/api/system/user';
import { isWechatBrowser } from '@/utils/is';
import { DeptTypeEnum } from '@/enums/deptEnum';
const Storage = createStorage({ storage: localStorage });

export interface UserInfoState {
  id: number;
  deptName: string;
  deptType: string;
  roleName: string;
  cityLabel: string;
  permissions: string[];
  username: string;
  realName: string;
  avatar: string;
  balance: number;
  integral: number;
  sex: number;
  qq: string;
  email: string;
  mobile: string;
  birthday: string;
  cityId: number;
  address: string;
  cash: {
    name: string;
    account: string;
    payeeCode: string;
  };
  createdAt: string;
  loginCount: number;
  lastLoginAt: string;
  lastLoginIp: string;
  openId: string;
  inviteCode: string;
}

export interface ConfigState {
  domain: string;
  version: string;
  wsAddr: string;
  mode: string;
}

export interface LoginConfigState {
  loginRegisterSwitch: number;
  loginCaptchaSwitch: number;
  loginAutoOpenId: number;
  loginProtocol: string;
  loginPolicy: string;
  i18nSwitch: boolean;
  defaultLanguage: string;
  projectName: string;
}

export interface IUserState {
  token: string;
  username: string;
  realName: string;
  avatar: string;
  permissions: any[];
  info: UserInfoState | null;
  config: ConfigState | null;
  loginConfig: LoginConfigState | null;
}

export const useUserStore = defineStore({
  id: 'app-member',
  state: (): IUserState => ({
    token: Storage.get(ACCESS_TOKEN, ''),
    username: '',
    realName: '',
    avatar: '',
    permissions: [],
    info: Storage.get(CURRENT_USER, null),
    config: Storage.get(CURRENT_CONFIG, null),
    loginConfig: Storage.get(CURRENT_LOGIN_CONFIG, null),
  }),
  getters: {
    getToken(): string {
      return this.token;
    },
    getAvatar(): string {
      return this.avatar;
    },
    getUsername(): string {
      return this.username;
    },
    getRealName(): string {
      return this.realName;
    },
    getPermissions(): [any][] {
      return this.permissions;
    },
    getUserInfo(): UserInfoState | null {
      return this.info;
    },
    getConfig(): ConfigState | null {
      return this.config;
    },
    getLoginConfig(): LoginConfigState | null {
      return this.loginConfig;
    },
    isCompanyDept(): boolean {
      return this.info?.deptType == DeptTypeEnum.Company;
    },
    isTenantDept(): boolean {
      return this.info?.deptType == DeptTypeEnum.Tenant;
    },
    isMerchantDept(): boolean {
      return this.info?.deptType == DeptTypeEnum.Merchant;
    },
    isUserDept(): boolean {
      return this.info?.deptType == DeptTypeEnum.User;
    },
  },
  actions: {
    setToken(token: string) {
      this.token = token;
    },
    setAvatar(avatar: string) {
      this.avatar = avatar;
    },
    setUsername(username: string) {
      this.username = username;
    },
    setRealName(realName: string) {
      this.realName = realName;
    },
    setPermissions(permissions: string[]) {
      this.permissions = permissions;
    },
    setUserInfo(info: UserInfoState | null) {
      this.info = info;
    },
    setConfig(config: ConfigState | null) {
      this.config = config;
    },
    setLoginConfig(config: LoginConfigState | null) {
      this.loginConfig = config;
    },
    // 账号登录
    async login(userInfo) {
      return await this.handleLogin(login(userInfo));
    },
    // 手机号登录
    async mobileLogin(userInfo) {
      return await this.handleLogin(mobileLogin(userInfo));
    },
    async handleLogin(request: Promise<any>) {
      try {
        const response = await request;
        const { data, code } = response;
        if (code === ResultEnum.SUCCESS) {
          const ex = 30 * 24 * 60 * 60 * 1000;
          storage.set(ACCESS_TOKEN, data.token, ex);
          storage.set(CURRENT_USER, data, ex);
          storage.set(IS_LOCKSCREEN, false);
          this.setToken(data.token);
          this.setUserInfo(data);
        }
        return Promise.resolve(response);
      } catch (e) {
        return Promise.reject(e);
      }
    },
    // 获取用户信息
    GetInfo() {
      const that: any = this;
      return new Promise((resolve, reject) => {
        getUserInfo()
          .then((res) => {
            const result = res as UserInfoState;
            if (result.permissions && result.permissions.length) {
              const permissionsList = result.permissions;
              that.setPermissions(permissionsList);
              that.setUserInfo(result);
              that.setAvatar(result.avatar);
              that.setUsername(result.username);
              that.setRealName(result.realName);
            } else {
              reject(new Error('getInfo: permissionsList must be a non-null array !'));
            }
            resolve(result);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
    // 获取基础配置
    GetConfig() {
      const that = this;
      return new Promise((resolve, reject) => {
        getConfig()
          .then((res) => {
            const result = res;
            that.setConfig(result);
            storage.set(CURRENT_CONFIG, result);
            resolve(res);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
    // 获取登录配置
    LoadLoginConfig: function () {
      const that = this;
      return new Promise((resolve, reject) => {
        getLoginConfig()
          .then((res) => {
            const result = res as unknown as LoginConfigState;
            that.setLoginConfig(result);
            storage.set(CURRENT_LOGIN_CONFIG, result);
            resolve(res);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
    // 是否允许获取微信openid
    allowWxOpenId(): boolean {
      if (!isWechatBrowser()) {
        return false;
      }

      if (this.loginConfig !== null && this.loginConfig.loginAutoOpenId !== 1) {
        return false;
      }
      return this.info === null || this.info.openId === '';
    },
    // 登出
    async logout() {
      try {
        const response = await logout();
        const { code } = response;
        if (code === ResultEnum.SUCCESS) {
          this.setPermissions([]);
          this.setUserInfo(null);
          storage.remove(ACCESS_TOKEN);
          storage.remove(CURRENT_USER);
          storage.remove(CURRENT_DICT);
        }
        return Promise.resolve(response);
      } catch (e) {
        return Promise.reject(e);
      }
    },
  },
});

// Need to be used outside the setup
export function useUserStoreWidthOut() {
  return useUserStore(store);
}
