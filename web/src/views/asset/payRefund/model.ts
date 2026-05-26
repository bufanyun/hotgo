import { ref } from 'vue';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';
import { renderOptionTag } from '@/utils';

const dict = useDictStore();

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
      return renderOptionTag('sys_normal_disable', row.status);
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

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}

loadOptions();
