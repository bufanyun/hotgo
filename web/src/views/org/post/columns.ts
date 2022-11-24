import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 100,
  },
  {
    title: '岗位名称',
    key: 'name',
    width: 100,
  },
  {
    title: '岗位名称',
    key: 'name',
    width: 100,
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
    title: '排序',
    key: 'sort',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 100,
    render: (rows, _) => {
      return rows.createdAt;
    },
  },
];
