import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts, formatBefore } from '@/utils/dateUtil';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
import { NetOption } from '@/api/monitor/monitor';

export interface State {
  id: number;
  group: string;
  name: string;
  appid: string;
  secretKey: string;
  remoteAddr: string;
  onlineLimit: number;
  loginTimes: number;
  lastLoginAt: string;
  lastActiveAt: string;
  routes: any;
  allowedIps: string;
  endAt: string;
  remark: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  group: '',
  name: '',
  appid: '',
  secretKey: '',
  remoteAddr: '',
  onlineLimit: 1,
  loginTimes: 0,
  lastLoginAt: '',
  lastActiveAt: '',
  routes: null,
  allowedIps: '',
  endAt: '',
  remark: '',
  status: 1,
  createdAt: '',
  updatedAt: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  sys_normal_disable: [],
  group: [],
  routes: [],
});

export const rules = {
  group: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入分组',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入许可名称',
  },
  appid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入应用ID',
  },
  endAt: {
    required: true,
    trigger: ['blur', 'input', 'focus'],
    type: 'string',
    message: '请输入授权结束时间',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInput',
    label: '许可ID',
    componentProps: {
      placeholder: '请输入许可ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'group',
    component: 'NSelect',
    label: '授权分组',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择授权分组',
      options: options.value.group,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'name',
    component: 'NInput',
    label: '许可名称',
    componentProps: {
      placeholder: '请输入许可名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'appid',
    component: 'NInput',
    label: 'APPID',
    componentProps: {
      placeholder: '请输入APPID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'endAt',
    component: 'NDatePicker',
    label: '过期时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '许可ID',
    key: 'id',
    width: 100,
  },
  {
    title: '授权分组',
    key: 'group',
    width: 100,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'info',
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.group, row.group),
        }
      );
    },
  },
  {
    title: '授权许可',
    key: 'name',
    render(row) {
      return h('p', { id: 'app' }, [
        h('div', {
          innerHTML: '<div><p>名称：' + row.name + '</p></div>',
        }),
        h('div', {
          innerHTML: '<div><p>APPID：' + row.appid + '</p></div>',
        }),
      ]);
    },
    width: 180,
  },
  {
    title: '在线',
    key: 'online',
    render(row) {
      return row.online + ' / ' + row.onlineLimit;
    },
    width: 100,
  },
  {
    title: '授权有效期',
    key: 'endAt',
    width: 150,
  },
  {
    title: '状态',
    key: 'status',
    width: 100,
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
  {
    title: '最后连接',
    key: 'remoteAddr',
    width: 150,
  },
  {
    title: '最近登录 / 心跳',
    key: 'name',
    render(row) {
      if (row.lastLoginAt === null) {
        return '从未登录';
      }
      return (
        formatBefore(new Date(row.lastLoginAt)) + ' / ' + formatBefore(new Date(row.lastActiveAt))
      );
    },
    width: 180,
  },
  {
    title: '累计登录',
    key: 'loginTimes',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 150,
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: ['sys_normal_disable'],
  });

  const netOption = await NetOption();
  options.value.group = netOption.licenseGroup;
  options.value.routes = netOption.routes;

  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
      case 'group':
        item.componentProps.options = options.value.group;
        break;
    }
  }
}

await loadOptions();
