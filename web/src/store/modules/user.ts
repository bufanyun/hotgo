import { defineStore } from 'pinia';
import { createStorage, storage } from '@/utils/Storage';
import { store } from '@/store';
import { ACCESS_TOKEN, CURRENT_CONFIG, CURRENT_USER, IS_LOCKSCREEN } from '@/store/mutation-types';
import { ResultEnum } from '@/enums/httpEnum';
import { getConfig, getUserInfo, login } from '@/api/system/user';

const Storage = createStorage({ storage: localStorage });

export interface IUserState {
  token: string;
  username: string;
  welcome: string;
  avatar: string;
  permissions: any[];
  info: any;
  config: any;
}

export const useUserStore = defineStore({
  id: 'app-member',
  state: (): IUserState => ({
    token: Storage.get(ACCESS_TOKEN, ''),
    username: '',
    welcome: '',
    avatar: '',
    permissions: [],
    info: Storage.get(CURRENT_USER, {}),
    config: Storage.get(CURRENT_CONFIG, {}),
  }),
  getters: {
    getToken(): string {
      return this.token;
    },
    getAvatar(): string {
      return this.avatar;
    },
    getNickname(): string {
      return this.username;
    },
    getPermissions(): [any][] {
      return this.permissions;
    },
    getUserInfo(): object {
      return this.info;
    },
    getConfig(): object {
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
    setPermissions(permissions) {
      this.permissions = permissions;
    },
    setUserInfo(info) {
      this.info = info;
    },
    setConfig(config) {
      this.config = config;
    },
    // 登录
    async login(userInfo) {
      try {
        const response = await login(userInfo);
        const { data, code } = response;
        console.log('data:' + JSON.stringify(data));
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
      // eslint-disable-next-line @typescript-eslint/no-this-alias
      const that = this;
      return new Promise((resolve, reject) => {
        getUserInfo()
          .then((res) => {
            const result = res;
            if (result.permissions && result.permissions.length) {
              const permissionsList = result.permissions;
              that.setPermissions(permissionsList);
              that.setUserInfo(result);
            } else {
              reject(new Error('getInfo: permissionsList must be a non-null array !'));
            }
            that.setAvatar(result.avatar);
            resolve(res);
          })
          .catch((error) => {
            reject(error);
          });
      });
    },
    // 获取用户信息
    GetConfig() {
      // eslint-disable-next-line @typescript-eslint/no-this-alias
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
      this.setUserInfo('');
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
