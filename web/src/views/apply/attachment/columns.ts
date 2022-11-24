import { h } from 'vue';
import { NAvatar, NTag } from 'naive-ui';

export const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: '应用',
    key: 'appId',
  },
  {
    title: '会员ID',
    key: 'memberId',
  },
  {
    title: '驱动',
    key: 'drive',
    render(row) {
      return row.drive;
    },
  },
  {
    title: '上传类型',
    key: 'kind',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.kind == 'images' ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => row.kind,
        }
      );
    },
  },
  {
    title: '文件',
    key: 'fileUrl',
    width: 80,
    render(row) {
      return h(NAvatar, {
        size: 40,
        src: row.fileUrl,
      });
    },
  },
  {
    title: '本地路径',
    key: 'path',
  },
  {
    title: '扩展名',
    key: 'ext',
  },
  {
    title: '文件大小',
    key: 'sizeFormat',
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
    title: '上传时间',
    key: 'createdAt',
  },
];
