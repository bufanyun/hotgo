import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: 'IP地址',
    key: 'ip',
  },
  {
    title: '备注',
    key: 'remark',
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.status == 1 ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => (row.status == 1 ? '正常' : '隐藏'),
        }
      );
    },
  },

  {
    title: '创建时间',
    key: 'createdAt',
  },
];
