import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '公告标题',
    key: 'title',
    render(row) {
      return row.title;
    },
  },
  {
    title: '公告类型',
    key: 'type',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.type == 1 ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => (row.type == 1 ? '通知' : '公告'),
        }
      );
    },
  },
  {
    title: '公告内容',
    key: 'content',
  },
  {
    title: '备注',
    key: 'remark',
  },
  {
    title: '排序',
    key: 'sort',
  },
  {
    title: '公告状态',
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
  {
    title: '已读人数',
    key: 'receiveNum',
  },
  {
    title: '发布时间',
    key: 'createdAt',
  },
];
