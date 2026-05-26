// 本地字典
// 一些不需要放在服务端注册的字典，也可以在这里进行注册

import { DictOptions, Option } from '@/store/modules/dict';

export const noticeTypeOptions: Option[] = [
  {
    key: 1,
    value: 1,
    label: '通知',
    listClass: 'warning',
    extra: null,
  },
  {
    key: 2,
    value: 2,
    label: '公告',
    listClass: 'error',
    extra: null,
  },
  {
    key: 3,
    value: 3,
    label: '私信',
    listClass: 'info',
    extra: null,
  },
];

export const noticeTagOptions: Option[] = [
  {
    value: 0,
    label: '无标签',
    key: 0,
    listClass: 'default',
    extra: null,
  },
  {
    value: 1,
    label: '一般',
    key: 1,
    listClass: 'info',
    extra: null,
  },
  {
    value: 2,
    label: '紧急',
    key: 2,
    listClass: 'error',
    extra: null,
  },
  {
    value: 3,
    label: '重要',
    key: 3,
    listClass: 'warning',
    extra: null,
  },
  {
    value: 4,
    label: '提醒',
    key: 4,
    listClass: 'success',
    extra: null,
  },
  {
    value: 5,
    label: '次要',
    key: 5,
    listClass: 'default',
    extra: null,
  },
];

export const tagTypeOptions: Option[] = [
  {
    value: 'success',
    label: '绿色',
    key: 'success',
    listClass: 'success',
    extra: null,
  },
  {
    value: 'warning',
    label: '橙色',
    key: 'warning',
    listClass: 'warning',
    extra: null,
  },
  {
    value: 'error',
    label: '红色',
    key: 'error',
    listClass: 'error',
    extra: null,
  },
  {
    value: 'info',
    label: '蓝色',
    key: 'info',
    listClass: 'info',
    extra: null,
  },
  {
    value: 'default',
    label: '灰色',
    key: 'default',
    listClass: 'default',
    extra: null,
  },
  {
    value: 'primary',
    label: '主题色',
    key: 'primary',
    listClass: 'primary',
    extra: null,
  },
];

export const localDict: DictOptions = {
  noticeTypeOptions: noticeTypeOptions, // 消息通知类型
  noticeTagOptions: noticeTagOptions, // 消息标签类型
  tagTypeOptions: tagTypeOptions, // 标签类型选项
};
