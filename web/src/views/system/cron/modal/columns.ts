import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    width: 100,
  },
  {
    title: '上级ID',
    dataIndex: 'pid',
    key: 'pid',
    width: 100,
  },
  {
    title: '分组名称',
    dataIndex: 'name',
    key: 'name',
    width: 200,
  },
  {
    title: '是否默认',
    dataIndex: 'isDefault',
    key: 'isDefault',
    render(row) {
      return row.is_default === 1 ? '是' : '否';
    },
    width: 100,
  },
  {
    title: '排序',
    dataIndex: 'sort',
    key: 'sort',
    width: 100,
  },

  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.status === 1 ? 'info' : 'error',
          bordered: false,
        },
        {
          default: () => (row.status === 1 ? '正常' : '禁用'),
        }
      );
    },
    width: 150,
  },
  {
    title: '备注',
    dataIndex: 'remark',
    key: 'remark',
    width: 150,
  },
  // {
  //   title: '更新时间',
  //   dataIndex: 'updatedAt',
  //   key: 'updatedAt',
  //   width: 160,
  // },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    width: 160,
  },
];
