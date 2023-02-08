import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';
import { Switch } from '@/api/test';
import { isNullObject } from '@/utils/is';
import { getFileExt } from '@/utils/urlUtils';
import { defRangeShortcuts, defShortcuts, formatToDate } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { errorImg, getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
const $message = window['$message'];
export interface State {
  id: number;
  memberId: number;
  categoryId: number;
  flag: number[] | null;
  title: string;
  content: string;
  image: string;
  images: string[] | null;
  attachfile: string;
  attachfiles: string[] | null;
  map: unknown[] | null;
  star: number;
  description: string;
  price: number;
  views: number;
  activityAt: string;
  startAt: null;
  endAt: null;
  switch: number;
  sort: number;
  avatar: string;
  sex: number;
  qq: string;
  email: string;
  mobile: string;
  channel: number;
  cityId: number;
  hobby: string[] | null;
  pid: number;
  level: number;
  tree: string;
  remark: string;
  status: number;
  createdBy: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  memberId: 0,
  categoryId: 0,
  flag: [1],
  title: '',
  content: '',
  image: '',
  images: null,
  attachfile: '',
  attachfiles: null,
  map: null,
  star: 0,
  description: '',
  price: 0,
  views: 0,
  activityAt: '',
  startAt: null,
  endAt: null,
  switch: 0,
  sort: 0,
  avatar: '',
  sex: 0,
  qq: '',
  email: '',
  mobile: '',
  channel: 0,
  cityId: 0,
  hobby: null,
  pid: 0,
  level: 1,
  tree: '',
  remark: '',
  status: 1,
  createdBy: 0,
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
  sys_user_sex: [],
  sys_notice_type: [],
  sys_user_channel: [],
  sys_user_hobby: [],
  sys_switch: [],
});

export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入标题',
  },
  price: {
    required: true,
    trigger: ['blur', 'input'],
    validator: validate.amount,
  },
  qq: {
    required: false,
    trigger: ['blur', 'input'],
    validator: validate.qq,
  },
  email: {
    required: true,
    trigger: ['blur', 'input'],
    validator: validate.email,
  },
  mobile: {
    required: true,
    trigger: ['blur', 'input'],
    validator: validate.phone,
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: {
      placeholder: '请输入标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: [{ message: '请输入标题', trigger: ['blur'] }],
  },
  {
    field: 'content',
    component: 'NInput',
    label: '内容',
    componentProps: {
      placeholder: '请输入内容关键词',
      showButton: false,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'price',
    labelMessage: '我是自定义提示',
    component: 'NInput',
    label: '价格',
    componentProps: {
      pair: true,
      separator: '-',
      clearable: true,
      placeholder: ['从', '到'],
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'activityAt',
    component: 'NDatePicker',
    label: '活动时间',
    componentProps: {
      type: 'date',
      clearable: true,
      shortcuts: defShortcuts(),
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
    field: 'flag',
    component: 'NCheckbox',
    label: '标签',
    giProps: {
      span: 1,
    },
    componentProps: {
      placeholder: '请选择标签',
      options: [],
      onUpdateChecked: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'switch',
    component: 'NRadioGroup',
    label: '开关',
    giProps: {
      //span: 24,
    },
    componentProps: {
      options: [],
      onUpdateChecked: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'hobby',
    component: 'NSelect',
    label: '爱好',
    defaultValue: null,
    componentProps: {
      multiple: true,
      placeholder: '请选择爱好',
      options: [],
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
      placeholder: '请选择类型',
      options: [],
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
    title: '标题',
    key: 'title',
    render(row) {
      return row.title;
    },
  },
  {
    title: '标签',
    key: 'flag',
    render(row) {
      if (isNullObject(row.flag)) {
        return ``;
      }
      return row.flag.map((tagKey) => {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.sys_notice_type, tagKey),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.sys_notice_type, tagKey),
          }
        );
      });
    },
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
    title: '多图',
    key: 'images',
    render(row) {
      if (isNullObject(row.images)) {
        return ``;
      }
      return row.images.map((image) => {
        return h(NImage, {
          width: 32,
          height: 32,
          src: image,
          onError: errorImg,
          style: {
            width: '32px',
            height: '32px',
            'max-width': '100%',
            'max-height': '100%',
            'margin-left': '2px',
          },
        });
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
    title: '多附件',
    key: 'attachfiles',
    render(row) {
      if (isNullObject(row.attachfiles)) {
        return ``;
      }
      return row.attachfiles.map((attachfile) => {
        return h(
          NAvatar,
          {
            size: 'small',
            style: {
              'margin-left': '2px',
            },
          },
          {
            default: () => getFileExt(attachfile),
          }
        );
      });
    },
  },
  {
    title: '推荐星',
    key: 'star',
    // width: 180,
    render(row) {
      return h(NRate, {
        allowHalf: true,
        readonly: true,
        defaultValue: row.star,
      });
    },
  },
  {
    title: '描述',
    key: 'description',
  },
  {
    title: '价格',
    key: 'price',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'success',
          bordered: false,
        },
        {
          default: () => row.price.toFixed(2),
        }
      );
    },
  },
  {
    title: '开关',
    key: 'switch',
    width: 100,
    render(row) {
      return h(NSwitch, {
        value: row.switch === 1,
        checked: '开启',
        unchecked: '关闭',
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
  // {
  //   title: '排序',
  //   key: 'sort',
  // },
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
    title: '爱好',
    key: 'hobby',
    render(row) {
      if (isNullObject(row.hobby)) {
        return ``;
      }
      return row.hobby.map((tagKey) => {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.sys_user_hobby, tagKey),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.sys_user_hobby, tagKey),
          }
        );
      });
    },
  },
  {
    title: '活动时间',
    key: 'activityAt',
    render(row) {
      return formatToDate(row.activityAt);
    },
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'sys_normal_disable',
      'sys_user_sex',
      'sys_notice_type',
      'sys_switch',
      'sys_user_hobby',
      'sys_user_channel',
    ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
      case 'flag':
        item.componentProps.options = options.value.sys_notice_type;
        break;
      case 'switch':
        item.componentProps.options = options.value.sys_switch;
        break;
      case 'hobby':
        item.componentProps.options = options.value.sys_user_hobby;
        break;
    }
  }
}

await loadOptions();
