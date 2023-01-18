import { defineStore } from 'pinia';
import { store } from '@/store';

export interface INotificationStore {
  messages: any[];
}

export const notificationStore = defineStore({
  id: 'notificationStore',
  state: (): INotificationStore => ({
    messages: [],
  }),
  getters: {
    getMessages(): [any][] {
      return this.messages;
    },
  },
  actions: {
    setMessages(messages) {
      this.messages = messages;
    },
    addMessages(message) {
      message = JSON.parse(message);
      if (
        message.event !== undefined &&
        message.event === 'notice' &&
        message.data !== undefined &&
        message.data !== ''
      ) {
        this.messages.unshift({
          title: message.data.title,
          description: message.data.type == 1 ? '通知' : '公告',
          content: message.data.content,
          meta: message.data.updatedAt,
        });
      }
      // 数据最大提醒条数，超出进行清理
      const limit = 10;
      if (this.messages.length > limit) {
        const sub = this.messages.length - limit;
        this.messages.splice(this.messages.length - sub);
      }
    },
  },
});

// Need to be used outside the setup
export function notificationStoreWidthOut() {
  return notificationStore(store);
}
