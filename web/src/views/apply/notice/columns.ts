import { h } from 'vue';
import { NAvatar, NAvatarGroup, NTag, NTooltip } from 'naive-ui';
import { noticeTagOptions, noticeTypeOptions } from '@/enums/systemMessageEnum';
import { getOptionLabel, getOptionTag } from '@/utils/hotgo';

export const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 80,
  },
  {
    title: '消息标题',
    key: 'title',
    render(row) {
      return h('p', { id: 'app' }, [
        h('div', {
          innerHTML: '<div style="white-space: pre-wrap">' + row.title + '</div>',
        }),
      ]);
    },
    width: 280,
  },
  {
    title: '消息类型',
    key: 'type',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(noticeTypeOptions, row.type),
          bordered: false,
        },
        {
          default: () => getOptionLabel(noticeTypeOptions, row.type),
        }
      );
    },
    width: 100,
  },
  {
    title: '标签',
    key: 'tag',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(noticeTagOptions, row.tag),
          bordered: false,
        },
        {
          default: () => getOptionLabel(noticeTagOptions, row.tag),
        }
      );
    },
    width: 100,
  },
  {
    title: '接收人',
    key: 'receiver',
    render(row) {
      if (row.type === 1 || row.type === 2) {
        return '所有人';
      }
      return h(
        NAvatarGroup,
        {
          max: 4,
          size: 40,
          options: row.receiverGroup,
        },
        {
          avatar: (column) =>
            h(NTooltip, null, {
              trigger: () =>
                column.option.src !== ''
                  ? h(NAvatar, {
                      src: column.option.src,
                      round: true,
                      size: 32,
                      style: {
                        marginRight: '4px',
                      },
                    })
                  : h(
                      NAvatar,
                      {
                        round: true,
                        size: 32,
                        style: {
                          marginRight: '4px',
                        },
                      },
                      {
                        default: () => column.option.name?.substring(0, 1) as string,
                      }
                    ),
              default: () => column.option.name,
            }),
        }
      );
    },
    width: 180,
  },
  {
    title: '阅读量',
    key: 'readCount',
    width: 80,
  },
  {
    title: '排序',
    key: 'sort',
    width: 80,
  },
  {
    title: '备注',
    key: 'remark',
    width: 150,
  },
  {
    title: '发送时间',
    key: 'createdAt',
    width: 180,
  },
];
