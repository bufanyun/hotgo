import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'id',
    key: 'id',
  },
  {
    title: '角色名称',
    key: 'name',
  },
  {
    title: '说明',
    key: 'remark',
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
    title: '创建时间',
    key: 'createdAt',
  },
];
