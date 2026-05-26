import { RouteLocationRaw } from 'vue-router';
import router from '@/router';

/**
 * 将对象添加当作参数拼接到URL上面
 * @param baseUrl 需要拼接的url
 * @param obj 参数对象
 * @returns {string} 拼接后的对象
 * 例子:
 *  let obj = {a: '3', b: '4'}
 *  setObjToUrlParams('www.baidu.com', obj)
 *  ==>www.baidu.com?a=3&b=4
 */
export function setObjToUrlParams(baseUrl: string, obj: object): string {
  let parameters = '';
  let url = '';
  for (const key in obj) {
    parameters += key + '=' + encodeURIComponent(obj[key]) + '&';
  }
  parameters = parameters.replace(/&$/, '');
  if (/\?$/.test(baseUrl)) {
    url = baseUrl + parameters;
  } else {
    url = baseUrl.replace(/\/?$/, '?') + parameters;
  }
  return url;
}

export function encodeParams(obj) {
  const arr = [];
  for (const p in obj) {
    // @ts-ignore
    arr.push(encodeURIComponent(p) + '=' + encodeURIComponent(obj[p]));
  }
  return arr.join('&');
}

/**
 * 获取文件后缀
 */
export function getFileExt(fileName: string) {
  if (fileName === undefined || fileName === '') {
    return ``;
  }
  return fileName.substring(fileName.lastIndexOf('.') + 1);
}

/**
 * 获取当访问的url，不含参数
 */
export function getNowUrl(): string {
  const w = window.location;
  return w.protocol + '//' + w.host + w.pathname;
}

// 返回上一页
export function goBackOrToPage(to: RouteLocationRaw): void {
  if (router.currentRoute.value.matched.length > 0) {
    router.go(-1);
  } else {
    router.push(to).catch((e) => {
      console.log('返回上一页失败:', e);
    });
  }
}

//  更新hash路由get参数，不刷新页面
export function pushHashRouterParameter(url, key, value) {
  const urlParts = url.split('?');
  let newUrl = '';
  if (urlParts.length >= 2) {
    const baseUrl = urlParts[0];
    const queryString = urlParts[1];
    const searchParams = new URLSearchParams(queryString);
    searchParams.set(key, value);
    newUrl = `${baseUrl}?${searchParams.toString()}`;
  } else {
    newUrl = `${url}?${key}=${value}`;
  }
  window.history.pushState({ path: newUrl }, '', newUrl);
}
