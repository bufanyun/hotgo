import { defineStore } from 'pinia';
import { createStorage, storage } from '@/utils/Storage';
import { store } from '@/store';
import { ACCESS_TOKEN, CURRENT_CONFIG, CURRENT_USER, IS_LOCKSCREEN } from '@/store/mutation-types';
import { ResultEnum } from '@/enums/httpEnum';
import { getConfig, getUserInfo, login } from '@/api/system/user';
const Storage = createStorage({ storage: localStorage });

export interface UserInfoState {
  id: number;
  deptName: string;
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
}

export interface ConfigState {
  domain: string;
  version: string;
  wsAddr: string;
}

export interface IUserState {
  token: string;
  username: string;
  realName: string;
  avatar: string;
  permissions: any[];
  info: UserInfoState | null;
  config: ConfigState | null;
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
    // 登录
    async login(userInfo) {
      try {
        const response = await login(userInfo);
        const { data, code } = response;
        if (code === ResultEnum.SUCCESS) {
          const ex = 7 * 24 * 60 * 60 * 1000;
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
    // 获取用户配置
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
    // 登出
    async logout() {
      this.setPermissions([]);
      this.setUserInfo(null);
      storage.remove(ACCESS_TOKEN);
      storage.remove(CURRENT_USER);
      return Promise.resolve('');
    },
  },
});

// Need to be used outside the setup
export function useUserStoreWidthOut() {
  return useUserStore(store);
}
