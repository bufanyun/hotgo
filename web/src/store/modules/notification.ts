import { defineStore } from 'pinia';
import { store } from '@/store';
import { PullMessages } from '@/api/apply/notice';
import { MessageRow, MessageTab, parseMessage } from '@/enums/systemMessageEnum';

export interface INotificationStore {
  messages: MessageTab[];
  notifyUnread: number;
  noticeUnread: number;
  letterUnread: number;
  newMessage: MessageRow | null;
}

export const notificationStore = defineStore({
  id: 'notificationStore',
  state: (): INotificationStore => ({
    messages: [
      {
        key: 1,
        name: '通知',
        badgeProps: { type: 'warning' },
        list: [],
      },
      {
        key: 2,
        name: '公告',
        badgeProps: { type: 'error' },
        list: [],
      },
      {
        key: 3,
        name: '私信',
        badgeProps: { type: 'info' },
        list: [],
      },
    ],
    notifyUnread: 0,
    noticeUnread: 0,
    letterUnread: 0,
    newMessage: null,
  }),
  getters: {
    getMessages(): MessageTab[] {
      return this.messages;
    },
  },
  actions: {
    setMessages(messages) {
      this.messages = messages;
    },
    triggerNewMessages(message) {
      message = parseMessage(message);
      this.addMessages(message);
      this.newMessage = message;
    },
    addMessages(message: MessageRow) {
      switch (message.type) {
        case 1:
          this.messages[0].list.push(message);
          this.notifyUnread++;
          break;
        case 2:
          this.messages[1].list.push(message);
          this.noticeUnread++;
          break;
        case 3:
          this.messages[2].list.push(message);
          this.letterUnread++;
          break;
      }
    },
    pullMessages() {
      PullMessages().then((res) => {
        if (res.list === undefined) {
          return;
        }

        this.messages[0].list = [];
        this.messages[1].list = [];
        this.messages[2].list = [];

        if (res.list?.length > 0) {
          for (let i = 0; i < res.list.length; i++) {
            this.addMessages(parseMessage(res.list[i]));
          }
        }

        this.notifyUnread = res.notifyCount;
        this.noticeUnread = res.noticeCount;
        this.letterUnread = res.letterCount;
      });
    },
    getUnreadCount() {
      return this.notifyUnread + this.noticeUnread + this.letterUnread;
    },
  },
});

// Need to be used outside the setup
export function notificationStoreWidthOut() {
  return notificationStore(store);
}
