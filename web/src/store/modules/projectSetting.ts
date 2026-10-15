import { defineStore } from 'pinia';
import { store } from '@/store';
import projectSetting from '@/settings/projectSetting';
import type { IheaderSetting, ImenuSetting, ImultiTabsSetting, IcrumbsSetting } from '/#/config';

const {
  navMode,
  navTheme,
  isMobile,
  headerSetting,
  showFooter,
  menuSetting,
  multiTabsSetting,
  crumbsSetting,
  permissionMode,
  isPageAnimate,
  pageAnimateType,
} = projectSetting;

interface ProjectSettingState {
  navMode: string; //导航模式
  navTheme: string; //导航风格
  headerSetting: IheaderSetting; //顶部设置
  showFooter: boolean; //页脚
  menuSetting: ImenuSetting; //多标签
  multiTabsSetting: ImultiTabsSetting; //多标签
  crumbsSetting: IcrumbsSetting; //面包屑
  permissionMode: string; //权限模式
  isPageAnimate: boolean; //是否开启路由动画
  pageAnimateType: string; //路由动画类型
  isMobile: boolean; // 是否处于移动端模式
}

export const useProjectSettingStore = defineStore({
  id: 'app-project-setting',
  state: (): ProjectSettingState => ({
    navMode: navMode,
    navTheme,
    isMobile,
    headerSetting,
    showFooter,
    menuSetting,
    multiTabsSetting,
    crumbsSetting,
    permissionMode,
    isPageAnimate,
    pageAnimateType,
  }),
  getters: {
    getNavMode(): string {
      return this.navMode;
    },
    getNavTheme(): string {
      return this.navTheme;
    },
    getIsMobile(): boolean {
      return this.isMobile;
    },
    getHeaderSetting(): object {
      if (this.navMode === 'double-column') {
        return { ...this.headerSetting, fixed: false };
      }
      return this.headerSetting;
    },
    getShowFooter(): boolean {
      return this.showFooter;
    },
    getMenuSetting(): object {
      if (this.navMode === 'double-column') {
        return { ...this.menuSetting, mixMenu: false };
      }
      return this.menuSetting;
    },
    getMultiTabsSetting(): object {
      return this.multiTabsSetting;
    },
    getCrumbsSetting(): object {
      return this.multiTabsSetting;
    },
    getPermissionMode(): string {
      return this.permissionMode;
    },
    getIsPageAnimate(): boolean {
      return this.isPageAnimate;
    },
    getPageAnimateType(): string {
      return this.pageAnimateType;
    },
  },
  actions: {
    setNavTheme(value: string): void {
      this.navTheme = value;
    },
    setIsMobile(value: boolean): void {
      this.isMobile = value;
    },
    // 切换导航模式；双栏同时关掉分割菜单和固定顶栏
    setNavMode(mode: string): void {
      this.navMode = mode;
      this.menuSetting.mixMenu = false;
      if (mode === 'double-column') {
        this.headerSetting.fixed = false;
      }
    },
    // 回填已保存的项目配置，双栏不保留分割菜单和固定顶栏
    setAllSettings(ps: ProjectSettingState): void {
      this.navMode = ps.navMode;
      this.navTheme = ps.navTheme;
      this.isMobile = ps.isMobile;
      this.headerSetting = ps.headerSetting;
      this.showFooter = ps.showFooter;
      this.menuSetting = ps.menuSetting;
      if (this.navMode === 'double-column') {
        this.menuSetting.mixMenu = false;
        this.headerSetting.fixed = false;
      }
      this.multiTabsSetting = ps.multiTabsSetting;
      this.crumbsSetting = ps.crumbsSetting;
      this.permissionMode = ps.permissionMode;
      this.isPageAnimate = ps.isPageAnimate;
      this.pageAnimateType = ps.pageAnimateType;
    },
  },
});

// Need to be used outside the setup
export function useProjectSettingStoreWithOut() {
  return useProjectSettingStore(store);
}
