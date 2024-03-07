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
import { getOptionLabel, getOptionTag, Option, Options, errorImg } from '@/utils/hotgo';

import { usePermission } from '@/hooks/web/usePermission';
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {
  public id = 0; // ID
  public categoryId = 0; // 分类ID
  public title = ''; // 标题
  public description = ''; // 描述
  public content = ''; // 内容
  public image = ''; // 单图
  public attachfile = ''; // 附件
  public cityId = 0; // 所在城市
  public switch = 2; // 显示开关
  public sort = 0; // 排序
  public status = 1; // 状态
  public createdBy = 0; // 创建者
  public updatedBy = 0; // 更新者
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间
  public deletedAt = ''; // 删除时间

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

export interface IOptions extends Options {
  sys_normal_disable: Option[]; 
};

export const options = ref<IOptions>({
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
        onError: errorImg,
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