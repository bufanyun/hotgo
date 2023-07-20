import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Option } from '@/api/creditsLog';
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
  creditType: [],
  creditGroup: [],
});

export const rules = {};

export const schemas = ref<FormSchema[]>([
  {
    field: 'memberId',
    component: 'NInput',
    label: '管理员ID',
    componentProps: {
      placeholder: '请输入管理员ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'creditGroup',
    component: 'NSelect',
    label: '组别',
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
    field: 'id',
    component: 'NInput',
    label: '变动ID',
    componentProps: {
      placeholder: '请输入变动ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '变动时间',
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
    width: 100,
  },
  {
    title: '管理员ID',
    key: 'memberId',
    width: 100,
  },
  {
    title: '变动类型',
    key: 'creditType',
    render(row) {
      if (isNullObject(row.creditType)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.creditType, row.creditType),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.creditType, row.creditType),
        }
      );
    },
    width: 150,
  },
  {
    title: '组别',
    key: 'creditGroup',
    render(row) {
      if (isNullObject(row.creditGroup)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.creditGroup, row.creditGroup),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.creditGroup, row.creditGroup),
        }
      );
    },
    width: 150,
  },
  {
    title: '变动前',
    key: 'beforeNum',
    width: 100,
    render(row) {
      return Number(row.beforeNum).toFixed(2);
    },
  },
  {
    title: '变动数量',
    key: 'num',
    width: 100,
    render(row) {
      return Number(row.num).toFixed(2);
    },
  },
  {
    title: '变动后',
    key: 'afterNum',
    width: 100,
    render(row) {
      return Number(row.afterNum).toFixed(2);
    },
  },
  {
    title: '备注',
    key: 'remark',
    width: 200,
  },
  {
    title: '操作人IP',
    key: 'ip',
    width: 150,
  },
  {
    title: '关联ID',
    key: 'mapId',
    width: 100,
    render(row) {
      if (row.mapId === 0) {
        return '-';
      }
      return row.mapId;
    },
  },
  {
    title: '变动时间',
    key: 'createdAt',
    width: 180,
  },
];

async function loadOptions() {
  options.value = await Option();
  for (const item of schemas.value) {
    switch (item.field) {
      case 'creditType':
        item.componentProps.options = options.value.creditType;
        break;
      case 'creditGroup':
        item.componentProps.options = options.value.creditGroup;
        break;
    }
  }
}

await loadOptions();
