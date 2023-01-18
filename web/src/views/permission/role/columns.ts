import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: '角色ID',
    key: 'id',
  },
  {
    title: '角色名称',
    key: 'name',
    render(row) {
      return h(
        NTag,
        {
          type: 'info',
        },
        {
          default: () => row.name,
        }
      );
    },
  },
  {
    title: '上级角色',
    key: 'pid',
  },
  {
    title: '是否默认角色',
    key: 'isDefault',
    render(row) {
      return h(
        NTag,
        {
          type: row.id == 1 ? 'success' : 'error',
        },
        {
          default: () => (row.id == 1 ? '是' : '否'),
        }
      );
    },
  },
  {
    title: '排序',
    key: 'sort',
  },
  {
    title: '备注',
    key: 'remark',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
];
