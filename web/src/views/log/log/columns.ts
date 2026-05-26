import { h } from 'vue';
import Column from './components/Column.vue';

export const columns = [
  {
    title: '记录ID',
    key: 'id',
    width: 100,
  },
  {
    title: '访客',
    key: 'name',
    width: 180,
    render(row) {
      return h(Column, {
        state: row,
        column: 'visitor',
      });
    },
  },
  {
    title: '请求接口',
    key: 'name',
    width: 260,
    render(row) {
      return h(Column, {
        state: row,
        column: 'request',
      });
    },
  },
  {
    title: '接口响应',
    key: 'name',
    width: 260,
    render(row) {
      return h(Column, {
        state: row,
        column: 'response',
      });
    },
  },
  {
    title: '访问时间',
    key: 'createdAt',
    width: 180,
  },
];
