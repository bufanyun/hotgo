import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 100,
  },
  {
    title: '岗位',
    key: 'name',
    width: 200,
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
    title: '岗位编码',
    key: 'code',
    width: 100,
  },
  {
    title: '状态',
    key: 'status',
    width: 100,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.status == 1 ? 'info' : 'error',
          bordered: false,
        },
        {
          default: () => (row.status == 1 ? '正常' : '已禁用'),
        }
      );
    },
  },
  {
    title: '备注',
    key: 'sort',
    width: 150,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
  },
];
