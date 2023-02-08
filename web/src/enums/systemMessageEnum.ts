import { NAvatar, NTag, NText, SelectRenderLabel, SelectRenderTag } from 'naive-ui';
import { Component, h } from 'vue';
import { getOptionLabel, getOptionTag, Option } from '@/utils/hotgo';
import { BellOutlined, NotificationOutlined, SendOutlined } from '@vicons/antd';

export const noticeTypeOptions: Option[] = [
  {
    key: 1,
    value: 1,
    label: '通知',
    listClass: 'warning',
  },
  {
    key: 2,
    value: 2,
    label: '公告',
    listClass: 'error',
  },
  {
    key: 3,
    value: 3,
    label: '私信',
    listClass: 'info',
  },
];

export const noticeTagOptions: Option[] = [
  {
    value: 0,
    label: '无标签',
    key: 0,
    listClass: 'default',
  },
  {
    value: 1,
    label: '一般',
    key: 1,
    listClass: 'info',
  },
  {
    value: 2,
    label: '紧急',
    key: 2,
    listClass: 'error',
  },
  {
    value: 3,
    label: '重要',
    key: 3,
    listClass: 'warning',
  },
  {
    value: 4,
    label: '提醒',
    key: 4,
    listClass: 'success',
  },
  {
    value: 5,
    label: '次要',
    key: 5,
    listClass: 'default',
  },
];

export interface personOption {
  value: number;
  label: string;
  username: string;
  avatar: string;
}

export const renderMultipleSelectTag: SelectRenderTag = ({ option, handleClose }) => {
  // @ts-ignore
  return h(
    NTag,
    {
      style: {
        padding: '0 6px 0 4px',
      },
      round: true,
      closable: true,
      onClose: (e) => {
        e.stopPropagation();
        handleClose();
      },
    },
    {
      default: () =>
        h(
          'div',
          {
            style: {
              display: 'flex',
              alignItems: 'center',
            },
          },
          [
            option.avatar !== ''
              ? h(NAvatar, {
                  src: option.avatar as string,
                  round: true,
                  size: 22,
                  style: {
                    marginRight: '4px',
                  },
                })
              : h(
                  NAvatar,
                  {
                    round: true,
                    size: 22,
                    style: {
                      marginRight: '4px',
                    },
                  },
                  {
                    default: () =>
                      option.label !== ''
                        ? ((option?.label as string).substring(0, 1) as string)
                        : ((option?.username as string).substring(0, 1) as string),
                  }
                ),
            option.label as string,
          ]
        ),
    }
  );
};

export const renderLabel: SelectRenderLabel = (option) => {
  return h(
    'div',
    {
      style: {
        display: 'flex',
        alignItems: 'center',
      },
    },
    [
      option.avatar !== ''
        ? h(NAvatar, {
            src: option.avatar as string,
            round: true,
            size: 'small',
          })
        : h(
            NAvatar,
            {
              round: true,
              size: 'small',
            },
            {
              default: () =>
                option.label !== ''
                  ? ((option?.label as string).substring(0, 1) as string)
                  : ((option?.username as string).substring(0, 2) as string),
            }
          ),
      h(
        'div',
        {
          style: {
            marginLeft: '12px',
            padding: '4px 0',
          },
        },
        [
          h('div', null, [option.label as string]),
          h(
            NText,
            { depth: 3, tag: 'div' },
            {
              default: () => option.username,
            }
          ),
        ]
      ),
    ]
  );
};

export interface MessageTab {
  /** tab的key */
  key: number;
  /** tab名称 */
  name: string;
  /** badge类型 */
  badgeProps?: import('naive-ui').BadgeProps;
  /** 消息数据 */
  list: MessageRow[];
}

export interface MessageRow {
  /** 消息ID */
  id: number;
  /** 消息类型 */
  type: number;
  /** 消息标题 */
  title: string;
  /** 消息内容 */
  content: string;
  /** 发送时间 */
  createdAt: string;
  /** 是否已读 */
  isRead: boolean;
  /** 发送者头像 */
  senderAvatar: string;
  /** 标签ID */
  tag: number;
  /** 标签名称 */
  tagTitle?: string;
  /** 标签props */
  tagProps?: import('naive-ui').TagProps;
}

// 获取消息的展示图标
export function getIcon(row: MessageRow): Component {
  if (row.type === 1) {
    return NotificationOutlined;
  }
  if (row.type === 2) {
    return BellOutlined;
  }
  return SendOutlined;
}

// 解析消息
export function parseMessage(row): MessageRow {
  row = row as MessageRow;
  if (row.tag <= 0) {
    return row;
  }

  row.tagTitle = getOptionLabel(noticeTagOptions, row.tag);
  row.tagProps = { type: getOptionTag(noticeTagOptions, row.tag) };
  return row;
}
