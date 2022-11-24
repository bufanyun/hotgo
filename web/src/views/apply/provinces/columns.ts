import { h } from 'vue';
import { NAvatar, NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '地区名称',
    key: 'title',
  },
  {
    title: '父ID',
    key: 'pid',
    render(row) {
      return row.pid;
    },
  },
  {
    title: '拼音',
    key: 'pinyin',
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
          default: () => row.pinyin,
        }
      );
    },
  },
  {
    title: '经度',
    key: 'lng',
  },
  {
    title: '维度',
    key: 'lat',
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

  {
    title: '创建时间',
    key: 'createdAt',
  },
];
