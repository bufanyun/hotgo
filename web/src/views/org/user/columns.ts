import { h } from 'vue';
import { NAvatar, NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 50,
  },
  {
    title: '用户名',
    key: 'username',
    width: 100,
  },
  {
    title: '姓名',
    key: 'realname',
    width: 100,
  },
  {
    title: '头像',
    key: 'avatar',
    width: 50,
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
              row.realname !== '' ? row.realname.substring(0, 1) : row.username.substring(0, 2),
          }
        );
      }
    },
  },
  {
    title: '绑定角色',
    key: 'role_name',
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
          default: () => row.role_name,
        }
      );
    },
  },
  {
    title: '所属部门',
    key: 'dept_name',
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
          default: () => row.dept_name,
        }
      );
    },
  },
  {
    title: '状态',
    key: 'status',
    width: 50,
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
    title: '访问次数',
    key: 'visitCount',
    width: 80,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 100,
    render: (rows, _) => {
      return rows.createdAt;
    },
  },
];
