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
    key: 'type',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.listClass,
          bordered: false,
        },
        {
          default: () => row.label,
        }
      );
    },
  },
  {
    title: '字典键值',
    key: 'value',
  },
  {
    title: '键值类型',
    key: 'valueType',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'default',
          bordered: false,
        },
        {
          default: () => row.valueType,
        }
      );
    },
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
];
