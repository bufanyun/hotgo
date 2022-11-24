import { h } from 'vue';
import { NAvatar, NTag } from 'naive-ui';
import { timestampToTime, formatBefore, formatAfter } from '@/utils/dateUtil';

export const columns = [
  {
    title: '会话编号',
    key: 'id',
    width: 240,
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
          default: () => row.id,
        }
      );
    },
  },
  {
    title: '登录应用',
    key: 'app',
    width: 80,
    render(row) {
      return row.app;
    },
  },
  // {
  //   title: '用户ID',
  //   key: 'userId',
  //   width: 100,
  // },
  {
    title: '用户名',
    key: 'username',
    width: 100,
  },
  {
    title: '头像',
    key: 'avatar',
    width: 80,
    render(row) {
      return h(NAvatar, {
        size: 32,
        src: row.avatar,
      });
    },
  },
  {
    title: '登录地址',
    key: 'addr',
    width: 150,
  },

  // {
  //   title: 'IP地区',
  //   key: 'region',
  // },
  {
    title: '浏览器',
    key: 'browser',
    width: 200,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'success',
          bordered: false,
        },
        {
          default: () => row.browser,
        }
      );
    },
  },
  {
    title: '操作系统',
    key: 'os',
    width: 150,
    render(row) {
      return row.os;
    },
  },
  {
    title: '授权过期',
    key: 'expTime',
    width: 80,
    render: (rows, _) => {
      return formatAfter(new Date(rows.expTime * 1000));
    },
  },
  {
    title: '最后活跃',
    key: 'heartbeatTime',
    width: 80,
    render: (rows, _) => {
      return formatBefore(new Date(rows.heartbeatTime * 1000));
    },
  },
  {
    title: '登录时间',
    key: 'firstTime',
    width: 220,
    render: (rows, _) => {
      return timestampToTime(rows.firstTime);
    },
  },
];
