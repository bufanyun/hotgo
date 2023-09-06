import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { Switch } from '@/api/curdDemo';
import { isArray, isNullObject } from '@/utils/is';
import { getFileExt } from '@/utils/urlUtils';
import { defRangeShortcuts, defShortcuts, formatToDate } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { getOptionLabel, getOptionTag, Options, errorImg } from '@/utils/hotgo';

import { usePermission } from '@/hooks/web/usePermission';
const { hasPermission } = usePermission();
const $message = window['$message'];


export interface State {
  id: number;
  categoryId: number;
  title: string;
  description: string;
  content: string;
  image: string;
  attachfile: string;
  cityId: number;
  switch: number;
  sort: number;
  status: number;
  createdBy: number;
  updatedBy: number;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

export const defaultState = {
  id: 0,
  categoryId: 0,
  title: '',
  description: '',
  content: '',
  image: '',
  attachfile: '',
  cityId: 0,
  switch: 1,
  sort: 0,
  status: 1,
  createdBy: 0,
  updatedBy: 0,
  createdAt: '',
  updatedAt: '',
  deletedAt: '',
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

export const rules = {
  categoryId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入分类ID',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标题',
  },
  description: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入描述',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入内容',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'ID',
    componentProps: {
      placeholder: '请输入ID',
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
  {
    field: 'testCategoryName',
    component: 'NInput',
    label: '分类名称',
    componentProps: {
      placeholder: '请输入分类名称',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '分类ID',
    key: 'categoryId',
  },
  {
    title: '标题',
    key: 'title',
  },
  {
    title: '描述',
    key: 'description',
  },
  {
    title: '单图',
    key: 'image',
    render(row) {
      return h(NImage, {
        width: 32,
        height: 32,
        src: row.image,
        fallbackSrc: errorImg,
        style: {
          width: '32px',
          height: '32px',
          'max-width': '100%',
          'max-height': '100%',
        },
      });
    },
  },
  {
    title: '附件',
    key: 'attachfile',
    render(row) {
      if (row.attachfile === '') {
        return ``;
      }
      return h(
        NAvatar,
        {
          size: 'small',
        },
        {
          default: () => getFileExt(row.attachfile),
        }
      );
    },
  },
  {
    title: '所在城市',
    key: 'cityId',
  },
  {
    title: '显示开关',
    key: 'switch',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.switch === 1,
        checked: '开启',
        unchecked: '关闭',
        disabled: !hasPermission(['/curdDemo/switch']),
        onUpdateValue: function (e) {
          console.log('onUpdateValue e:' + JSON.stringify(e));
          row.switch = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', value: row.switch }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '排序',
    key: 'sort',
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
    title: '创建者',
    key: 'createdBy',
  },
  {
    title: '更新者',
    key: 'updatedBy',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '修改时间',
    key: 'updatedAt',
  },
  {
    title: '分类名称',
    key: 'testCategoryName',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'sys_normal_disable',
   ],
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
