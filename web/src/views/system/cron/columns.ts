import { h } from 'vue';
import { NTag } from 'naive-ui';

const policyOptions = {
  1: '并行策略',
  2: '单例策略',
  3: '单次策略',
  4: '多次策略',
};

export const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 80,
  },
  {
    title: '任务标题',
    key: 'title',
    width: 150,
  },
  {
    title: '任务分组',
    key: 'groupName',
    width: 100,
  },
  {
    title: '执行方法',
    key: 'name',
    width: 100,
  },
  // {
  //   title: '执行参数',
  //   key: 'params',
  //   render(row) {
  //     return row.params;
  //   },
  // },
  {
    title: '执行策略',
    key: 'policy',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'info',
          bordered: false,
        },
        {
          default: () => policyOptions[row.policy] ?? '未知',
        }
      );
    },
    width: 100,
  },
  {
    title: '表达式',
    key: 'pattern',
    width: 150,
  },
  // {
  //   title: '执行次数',
  //   key: 'count',
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
          default: () => (row.status == 1 ? '运行中' : '已结束'),
        }
      );
    },
    width: 100,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
  },
];
