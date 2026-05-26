import { h } from 'vue';
import { NAvatar, NTag, NText } from 'naive-ui';
import { formatBefore } from '@/utils/dateUtil';

export const columns = [
  {
    title: '管理员ID',
    key: 'id',
    width: 100,
  },
  {
    title: '用户名',
    key: 'username',
    width: 100,
  },
  {
    title: '姓名',
    key: 'realName',
    width: 100,
    render(row) {
      if (row.realName == '') {
        return h(NText, { depth: 3 }, { default: () => '未设置' });
      }
      return row.realName;
    },
  },
  {
    title: '头像',
    key: 'avatar',
    width: 70,
    render(row) {
      if (row.avatar !== '') {
        return h(NAvatar, {
          circle: true,
          size: 'small',
          src: row.avatar,
        });
      } else {
        return h(
          NAvatar,
          {
            circle: true,
            size: 'small',
          },
          {
            default: () =>
              row.realName !== '' ? row.realName.substring(0, 1) : row.username.substring(0, 2),
          }
        );
      }
    },
  },
  {
    title: '绑定角色',
    key: 'roleName',
    width: 100,
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
          default: () => row.roleName,
        }
      );
    },
  },
  {
    title: '所属部门',
    key: 'deptName',
    width: 100,
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
          default: () => row.deptName,
        }
      );
    },
  },
  {
    title: '余额',
    key: 'balance',
    width: 120,
    render(row) {
      return '￥' + Number(row.balance).toFixed(2);
    },
  },
  {
    title: '积分',
    key: 'integral',
    width: 120,
    render(row) {
      return Number(row.integral).toFixed(2);
    },
  },
  {
    title: '状态',
    key: 'status',
    width: 80,
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
    title: '最近活跃',
    key: 'lastActiveAt',
    width: 100,
    render(row) {
      if (row.lastActiveAt === null) {
        return '从未登录';
      }
      return formatBefore(new Date(row.lastActiveAt));
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
    render: (rows, _) => {
      return rows.createdAt;
    },
  },
];
