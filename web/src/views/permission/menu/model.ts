import { cloneDeep } from 'lodash-es';

export interface State {
  id: number;
  pid: number;
  title: string;
  name: string;
  path: string;
  label: string;
  icon: string;
  type: number;
  redirect: string;
  permissions: string;
  permissionName: string;
  component: string;
  alwaysShow: number;
  activeMenu: string;
  isRoot: number;
  isFrame: number;
  frameSrc: string;
  keepAlive: number;
  hidden: number;
  affix: number;
  status: number;
  sort: number;
}

export const defaultState = {
  id: 0,
  pid: 0,
  title: '',
  name: '',
  path: '',
  label: '',
  icon: '',
  type: 1,
  redirect: '',
  permissions: '',
  permissionName: '',
  component: '',
  alwaysShow: 1,
  activeMenu: '',
  isRoot: 0,
  isFrame: 2,
  frameSrc: '',
  keepAlive: 0,
  hidden: 0,
  affix: 0,
  status: 1,
  sort: 10,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}
