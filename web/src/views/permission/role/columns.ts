import { h } from 'vue';
import { NTag, NButton } from 'naive-ui';
import { HelpCircleOutline } from '@vicons/ionicons5';
import { renderTooltip, renderIcon } from '@/utils';
export const columns = [
  // {
  //   title: '角色ID',
  //   key: 'id',
  // },
  {
    title(_column) {
      return renderTooltip(
        h(
          NButton,
          {
            ghost: true,
            strong: true,
            size: 'small',
            text: true,
            iconPlacement: 'right',
          },
          { default: () => '角色', icon: renderIcon(HelpCircleOutline) }
        ),
        '支持上下级角色，点击列表中左侧 > 按钮可展开下级角色列表'
      );
    },
    key: 'name',
    render(row) {
      return h(
        NTag,
        {
          type: 'info',
        },
        {
          default: () => row.name,
        }
      );
    },
    width: 200,
  },
  {
    title: '角色编码',
    key: 'key',
    // width: 150,
  },
  // {
  //   title: '上级角色',
  //   key: 'pid',
  // },
  {
    title: '默认角色',
    key: 'isDefault',
    render(row) {
      return h(
        NTag,
        {
          type: row.id == 1 ? 'success' : 'error',
        },
        {
          default: () => (row.id == 1 ? '是' : '否'),
        }
      );
    },
    // width: 80,
  },
  {
    title: '排序',
    key: 'sort',
    // width: 100,
  },
  {
    title: '备注',
    key: 'remark',
    // width: 300,
  },
  {
    title: '状态',
    key: 'status',
    // width: 80,
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
    title: '创建时间',
    key: 'createdAt',
    width: 180,
  },
];
