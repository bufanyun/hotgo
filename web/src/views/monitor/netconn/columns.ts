import { h } from 'vue';
import { NButton, NTag } from 'naive-ui';
import { timestampToTime, formatBefore } from '@/utils/dateUtil';
import { getOptionLabel } from '@/utils/hotgo';
import { options } from '@/views/monitor/netconn/modal/model';
import { renderIcon, renderTooltip } from '@/utils';
import { HelpCircleOutline } from '@vicons/ionicons5';

export const columns = [
  {
    title: '连接ID',
    key: 'id',
    width: 80,
    render(row) {
      return '# ' + row.id;
    },
  },
  {
    title: '连接协议',
    key: 'proto',
    width: 80,
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
          default: () => row.proto,
        }
      );
    },
  },
  {
    title: '服务端口',
    key: 'port',
    width: 80,
    render(row) {
      return row.port;
    },
  },
  {
    title: '应用名称',
    key: 'name',
    width: 100,
    render: (rows, _) => {
      return rows.name;
    },
  },
  {
    title: '登录地址',
    key: 'addr',
    width: 120,
  },
  {
    title(_column) {
      return renderTooltip(
        h(
          NButton,
          {
            strong: true,
            size: 'small',
            text: true,
            iconPlacement: 'right',
          },
          { default: () => '认证状态', icon: renderIcon(HelpCircleOutline) }
        ),
        '成功连接到服务端口并通过登录许可认证的客户端连接会显示为 `已认证` 状态'
      );
    },
    key: 'isAuth',
    width: 80,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.isAuth ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => (row.isAuth ? '已认证' : '未认证'),
        }
      );
    },
  },
  {
    title: '授权分组',
    key: 'group',
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
          default: () => getOptionLabel(options.value.group, row.group),
        }
      );
    },
  },
  {
    title: '授权许可',
    key: 'name',
    render(row) {
      return h('p', { id: 'app' }, [
        h('div', {
          innerHTML: '<div><p>名称：' + row.licenseName + '</p></div>',
        }),
        h('div', {
          innerHTML: '<div><p>APPID：' + row.appId + '</p></div>',
        }),
      ]);
    },
    width: 180,
  },
  {
    title: '心跳',
    key: 'heartbeatTime',
    width: 100,
    render: (rows, _) => {
      return formatBefore(new Date(rows.heartbeatTime * 1000));
    },
  },
  {
    title: '登录时间',
    key: 'firstTime',
    width: 150,
    render: (rows, _) => {
      return timestampToTime(rows.firstTime);
    },
  },
];
