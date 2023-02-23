import { h, unref } from 'vue';
import type { App, Plugin } from 'vue';
import { NIcon, NTag, NTooltip } from 'naive-ui';
import { PageEnum } from '@/enums/pageEnum';
import { isObject } from './is/index';
import { cloneDeep } from 'lodash-es';

export const renderTooltip = (trigger, content) => {
  return h(NTooltip, null, {
    trigger: () => trigger,
    default: () => content,
  });
};

/**
 * render 图标
 * */
export function renderIcon(icon, props = null) {
  return () => h(NIcon, props, { default: () => h(icon) });
}

/**
 * render new Tag
 * */
const newTagColors = { color: '#f90', textColor: '#fff', borderColor: '#f90' };

export function renderNew(type = 'warning', text = 'New', color: object = newTagColors) {
  return () =>
    h(
      NTag as any,
      {
        type,
        round: true,
        size: 'small',
        color,
      },
      { default: () => text }
    );
}

/**
 * 递归组装菜单格式
 */
export function generatorMenu(routerMap: Array<any>) {
  return filterRouter(routerMap).map((item) => {
    const isRoot = isRootRouter(item);
    const info = isRoot ? item.children[0] : item;
    const currentMenu = {
      ...info,
      ...info.meta,
      label: info.meta?.title,
      key: info.name,
      icon: isRoot ? item.meta?.icon : info.meta?.icon,
    };
    // 是否有子菜单，并递归处理
    if (info.children && info.children.length > 0) {
      // Recursion
      currentMenu.children = generatorMenu(info.children);

      // 当生成后子集为空，则删除子集空数组，否则加载时仍为目录格式！
      if (currentMenu.children.length === 0) {
        delete currentMenu.children;
      }
    }
    return currentMenu;
  });
}

/**
 * 混合菜单
 * */
export function generatorMenuMix(routerMap: Array<any>, routerName: string, location: string) {
  const cloneRouterMap = cloneDeep(routerMap);
  const newRouter = filterRouter(cloneRouterMap);
  if (location === 'header') {
    const firstRouter: any[] = [];
    newRouter.forEach((item) => {
      const isRoot = isRootRouter(item);
      const info = isRoot ? item.children[0] : item;
      info.children = undefined;
      const currentMenu = {
        ...info,
        ...info.meta,
        label: info.meta?.title,
        key: info.name,
      };
      firstRouter.push(currentMenu);
    });
    return firstRouter;
  } else {
    return getChildrenRouter(newRouter.filter((item) => item.name === routerName));
  }
}

/**
 * 递归组装子菜单
 * */
export function getChildrenRouter(routerMap: Array<any>) {
  return filterRouter(routerMap).map((item) => {
    const isRoot = isRootRouter(item);
    const info = isRoot ? item.children[0] : item;
    const currentMenu = {
      ...info,
      ...info.meta,
      label: info.meta?.title,
      key: info.name,
    };
    // 是否有子菜单，并递归处理
    if (info.children && info.children.length > 0) {
      // Recursion
      currentMenu.children = getChildrenRouter(info.children);
    }
    return currentMenu;
  });
}

/**
 * 判断根路由 Router
 * */
export function isRootRouter(item) {
  if (item.meta?.alwaysShow != true && item.children?.length === 1) {
    return true;
  }

  // if (item.meta?.alwaysShow != true) {
  //   if (item.children?.length > 0) {
  //     // 如果存在子级。且只要有一个不是隐藏状态的，则判断不是跟路由
  //     for (let i = 0; i < item.children.length; i++) {
  //       if (item.children[i]?.hidden == false) {
  //         return false;
  //       }
  //     }
  //
  //     return true;
  //   }
  // }

  return false;
}

/**
 * 强制根路由转换
 * @param item
 */
export function mandatoryRootConvert(item) {
  if (item.meta?.isRoot === true) {
  }

  // 默认
  return item.children[0];
}

/**
 * 排除Router
 * */
export function filterRouter(routerMap: Array<any>) {
  return routerMap.filter((item) => {
    return (
      (item.meta?.hidden || false) != true &&
      !['/:path(.*)*', '/', PageEnum.REDIRECT, PageEnum.BASE_LOGIN].includes(item.path)
    );
  });
}

export const withInstall = <T>(component: T, alias?: string) => {
  const comp = component as any;
  comp.install = (app: App) => {
    // @ts-ignore
    app.component(comp.name || comp.displayName, component);
    if (alias) {
      app.config.globalProperties[alias] = component;
    }
  };
  return component as T & Plugin;
};

/**
 *  找到对应的节点
 * */
let result = null;

export function getTreeItem(data: any[], key?: string | number): any {
  data.map((item) => {
    if (item.key === key) {
      result = item;
    } else {
      if (item.children && item.children.length) {
        getTreeItem(item.children, key);
      }
    }
  });
  return result;
}

/**
 *  找到所有节点
 * */
const treeAll: any[] = [];

export function getTreeAll(data: any[]): any[] {
  data.map((item) => {
    treeAll.push(item.key);
    if (item.children && item.children.length) {
      getTreeAll(item.children);
    }
  });
  return treeAll;
}

// dynamic use hook props
export function getDynamicProps<T, U>(props: T): Partial<U> {
  const ret: Recordable = {};

  // @ts-ignore
  Object.keys(props).map((key) => {
    ret[key] = unref((props as Recordable)[key]);
  });

  return ret as Partial<U>;
}

export function deepMerge<T = any>(src: any = {}, target: any = {}): T {
  let key: string;
  for (key in target) {
    src[key] = isObject(src[key]) ? deepMerge(src[key], target[key]) : (src[key] = target[key]);
  }
  return src;
}

/**
 * Sums the passed percentage to the R, G or B of a HEX color
 * @param {string} color The color to change
 * @param {number} amount The amount to change the color by
 * @returns {string} The processed part of the color
 */
function addLight(color: string, amount: number) {
  const cc = parseInt(color, 16) + amount;
  const c = cc > 255 ? 255 : cc;
  return c.toString(16).length > 1 ? c.toString(16) : `0${c.toString(16)}`;
}

/**
 * Lightens a 6 char HEX color according to the passed percentage
 * @param {string} color The color to change
 * @param {number} amount The amount to change the color by
 * @returns {string} The processed color represented as HEX
 */
export function lighten(color: string, amount: number) {
  color = color.indexOf('#') >= 0 ? color.substring(1, color.length) : color;
  amount = Math.trunc((255 * amount) / 100);
  return `#${addLight(color.substring(0, 2), amount)}${addLight(
    color.substring(2, 4),
    amount
  )}${addLight(color.substring(4, 6), amount)}`;
}

/**
 * 判断是否 url
 * */
export function isUrl(url: string) {
  return /^(http|https):\/\//g.test(url);
}
