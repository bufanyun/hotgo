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
    width: 100,
  },
  {
    title: '操作人',
    key: 'memberName',
    render(row) {
      if (row.memberId === 0) {
        return row.memberName;
      }
      return row.memberName + '(' + row.memberId + ')';
    },
    width: 150,
  },
  {
    title: '请求方式',
    key: 'method',
    width: 80,
  },
  {
    title: '请求路径',
    key: 'url',
    width: 200,
  },
  {
    title: '访问IP',
    key: 'ip',
    width: 150,
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
    width: 150,
  },
  {
    title: '处理耗时',
    key: 'takeUpTime',
    render(row) {
      return row.takeUpTime + ' ms';
    },
    width: 120,
  },
  {
    title: '访问时间',
    key: 'createdAt',
    width: 150,
  },
];
