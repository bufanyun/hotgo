import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';

export interface State {
  id: number;
  memberId: number;
  appId: string;
  addonsName: string;
  creditType: string;
  creditGroup: string;
  beforeNum: number;
  num: number;
  afterNum: number;
  remark: string;
  ip: string;
  mapId: number;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  memberId: 0,
  appId: '',
  addonsName: '',
  creditType: '',
  creditGroup: '',
  beforeNum: 0,
  num: 0,
  afterNum: 0,
  remark: '',
  ip: '',
  mapId: 0,
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
});

export const rules = {};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '变动ID',
    componentProps: {
      placeholder: '请输入变动ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'memberId',
    component: 'NInputNumber',
    label: '管理员ID',
    componentProps: {
      placeholder: '请输入管理员ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'appId',
    component: 'NInput',
    label: '应用id',
    componentProps: {
      placeholder: '请输入应用id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'creditType',
    component: 'NSelect',
    label: '变动类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择变动类型',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'creditGroup',
    component: 'NSelect',
    label: '变动的组别',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择变动的组别',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'remark',
    component: 'NInput',
    label: '备注',
    componentProps: {
      placeholder: '请输入备注',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'ip',
    component: 'NInput',
    label: '操作人IP',
    componentProps: {
      placeholder: '请输入操作人IP',
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
    title: '变动ID',
    key: 'id',
  },
  {
    title: '管理员ID',
    key: 'memberId',
  },
  {
    title: '应用id',
    key: 'appId',
  },
  {
    title: '插件名称',
    key: 'addonsName',
  },
  {
    title: '变动前',
    key: 'beforeNum',
  },
  {
    title: '变动数据',
    key: 'num',
  },
  {
    title: '变动后',
    key: 'afterNum',
  },
  {
    title: '备注',
    key: 'remark',
  },
  {
    title: '操作人IP',
    key: 'ip',
  },
  {
    title: '关联ID',
    key: 'mapId',
  },
  {
    title: '状态',
    key: 'status',
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
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '修改时间',
    key: 'updatedAt',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: ['sys_normal_disable'],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
    }
  }
}

await loadOptions();
