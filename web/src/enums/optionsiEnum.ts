import { Option } from '@/utils/hotgo';

export const switchOptions = [
  {
    value: 1,
    label: '已开启',
  },
  {
    value: 2,
    label: '已关闭',
  },
].map((s) => {
  return s;
});

export const sexOptions = [
  {
    value: 1,
    label: '男',
  },
  {
    value: 2,
    label: '女',
  },
  {
    value: 3,
    label: '未知',
  },
].map((s) => {
  return s;
});

export const statusOptions = [
  {
    value: 1,
    label: '正常',
  },
  {
    value: 2,
    label: '已禁用',
  },
].map((s) => {
  return s;
});

export const hiddenOptions = [
  {
    value: 1,
    label: '是',
  },
  {
    value: 2,
    label: '否',
  },
].map((s) => {
  return s;
});

// 操作类
export const statusActions = [
  {
    label: '设为启用',
    key: 1,
  },
  {
    label: '设为禁用',
    key: 2,
  },
];

// 标签
export const tagOptions = [
  {
    label: '灰色',
    value: 'default',
  },
  {
    label: '主色',
    value: 'primary',
  },
  {
    label: '蓝色',
    value: 'info',
  },
  {
    label: '绿色',
    value: 'success',
  },
  {
    label: '橙色',
    value: 'warning',
  },
  {
    label: '红色',
    value: 'error',
  },
];

// 登录状态
export const loginStatusOptions: Option[] = [
  {
    value: 1,
    label: '成功',
    key: 1,
    listClass: 'success',
  },
  {
    value: 2,
    label: '失败',
    key: 2,
    listClass: 'warning',
  },
];
