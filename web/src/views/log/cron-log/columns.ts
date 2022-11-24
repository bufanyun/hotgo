import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: '模块',
    key: 'module',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.module == 'admin' ? 'info' : 'success',
          bordered: false,
        },
        {
          default: () => row.module,
        }
      );
    },
  },
  {
    title: '操作人',
    key: 'member_name',
    render(row) {
      if (row.memberId === 0) {
        return row.member_name;
      }
      return row.member_name + '(' + row.memberId + ')';
    },
  },
  {
    title: '请求方式',
    key: 'method',
  },
  {
    title: '请求路径',
    key: 'url',
  },
  {
    title: '访问IP',
    key: 'ip',
  },
  // {
  //   title: 'IP地区',
  //   key: 'region',
  // },
  {
    title: '状态码',
    key: 'errorCode',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.errorCode == 0 ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => row.errorMsg + '(' + row.errorCode + ')',
        }
      );
    },
  },
  {
    title: 'Goroutine耗时',
    key: 'takeUpTime',
    render(row) {
      return row.takeUpTime + ' ms';
    },
  },
  {
    title: '访问时间',
    key: 'createdAt',
  },
];
