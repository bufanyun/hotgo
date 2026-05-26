import { TABS_ROUTES } from '@/store/mutation-types';
import { SocketEnum } from '@/enums/socketEnum';
import { useUserStoreWidthOut } from '@/store/modules/user';
import { notificationStoreWidthOut } from '@/store/modules/notification';
import { addOnMessage, WebSocketMessage } from '@/utils/websocket/index';

// 注册全局消息监听
export function registerGlobalMessage() {
  // 心跳
  addOnMessage(SocketEnum.EventPing, function (_message: WebSocketMessage) {
    // console.log('ping..');
  });

  // 强制退出
  addOnMessage(SocketEnum.EventKick, function (_message: WebSocketMessage) {
    const useUserStore = useUserStoreWidthOut();
    useUserStore.logout().then(() => {
      // 移除标签页
      localStorage.removeItem(TABS_ROUTES);
      location.reload();
    });
  });

  // 消息通知
  addOnMessage(SocketEnum.EventNotice, function (message: WebSocketMessage) {
    const notificationStore = notificationStoreWidthOut();
    notificationStore.triggerNewMessages(message.data);
  });

  // 更多全局消息处理都可以在这里注册
  // ...
}
