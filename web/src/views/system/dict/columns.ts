import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'id',
    key: 'id',
  },
  {
    title: '字典类型',
    key: 'type',
  },
  {
    title: '字典标签',
    key: 'label',
  },
  {
    title: '字典键值',
    key: 'value',
  },

  // {
  //   title: '备注',
  //   key: 'remark',
  // },
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
  // {
  //   title: '创建时间',
  //   key: 'createdAt',
  //   width: 100,
  // },
];
