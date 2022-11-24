import {http} from '@/utils/http/axios';
import {ApiEnum} from "@/enums/apiEnum";

/**
 * @description: 根据用户id获取用户菜单
 */
export function adminMenus() {
  return http.request({
    url: ApiEnum.RoleDynamic,
    method: 'GET',
  });
}

/**
 * 获取tree菜单列表
 * @param params
 */
export function getMenuList(params?) {
  return http.request({
    url: '/menu/list',
    method: 'GET',
    params,
  });
}

export function EditMenu(params?) {
  return http.request({
    url: '/menu/edit',
    method: 'POST',
    params,
  });
}

export function DeleteMenu(params?) {
  return http.request({
    url: '/menu/delete',
    method: 'POST',
    params,
  });
}
