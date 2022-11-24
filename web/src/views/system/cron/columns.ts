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
  },
  {
    title: '任务分组',
    key: 'groupName',
  },
  {
    title: '任务名称',
    key: 'name',
  },
  {
    title: '执行参数',
    key: 'params',
    render(row) {
      return row.params;
    },
  },
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
  },
  {
    title: '表达式',
    key: 'pattern',
  },
  {
    title: '执行次数',
    key: 'count',
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
          default: () => (row.status == 1 ? '运行中' : '已结束'),
        }
      );
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
];
