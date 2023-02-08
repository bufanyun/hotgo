import { SocketEnum } from '@/enums/socketEnum';
import { notificationStoreWidthOut } from '@/store/modules/notification';
import { useUserStoreWidthOut } from '@/store/modules/user';
import { TABS_ROUTES } from '@/store/mutation-types';
import { MessageRow } from '@/enums/systemMessageEnum';
import { isJsonString } from '@/utils/is';

let socket: WebSocket;
let isActive: boolean;

export function getSocket(): WebSocket {
  if (socket === undefined) {
    location.reload();
  }
  return socket;
}

export function getActive(): boolean {
  return isActive;
}

export function sendMsg(event: string, data = null, isRetry = true) {
  if (socket === undefined || !isActive) {
    if (!isRetry) {
      console.log('socket连接异常，发送失败！');
      return;
    }
    console.log('socket连接异常，等待重试..');
    setTimeout(function () {
      sendMsg(event, data);
    }, 200);
    return;
  }

  try {
    socket.send(
      JSON.stringify({
        event: event,
        data: data,
      })
    );
  } catch (err) {
    // @ts-ignore
    console.log('ws发送消息失败，等待重试，err：' + err.message);
    if (!isRetry) {
      return;
    }
    setTimeout(function () {
      sendMsg(event, data);
    }, 100);
  }
}

export function addOnMessage(onMessageList: any, func: Function) {
  let exist = false;
  for (let i = 0; i < onMessageList.length; i++) {
    if (onMessageList[i].name == func.name) {
      onMessageList[i] = func;
      exist = true;
    }
  }
  if (!exist) {
    onMessageList.push(func);
  }
}

export default (onMessage: Function) => {
  const heartCheck = {
    timeout: 5000,
    timeoutObj: setTimeout(() => {}),
    serverTimeoutObj: setInterval(() => {}),
    reset: function () {
      clearTimeout(this.timeoutObj);
      clearTimeout(this.serverTimeoutObj);
      return this;
    },
    start: function () {
      // eslint-disable-next-line @typescript-eslint/no-this-alias
      const self = this;
      clearTimeout(this.timeoutObj);
      clearTimeout(this.serverTimeoutObj);
      this.timeoutObj = setTimeout(function () {
        socket.send(
          JSON.stringify({
            event: SocketEnum.EventPing,
          })
        );
        self.serverTimeoutObj = setTimeout(function () {
          console.log('关闭服务');
          socket.close();
        }, self.timeout);
      }, this.timeout);
    },
  };

  const notificationStore = notificationStoreWidthOut();
  const useUserStore = useUserStoreWidthOut();
  let lockReconnect = false;
  let timer: ReturnType<typeof setTimeout>;
  const createSocket = () => {
    console.log('createSocket...');
    try {
      if (useUserStore.token === '') {
        throw new Error('用户未登录，稍后重试...');
      }
      socket = new WebSocket(useUserStore.config.wsAddr + '?authorization=' + useUserStore.token);
      init();
    } catch (e) {
      console.log('createSocket err:' + e);
      reconnect();
    }
    if (lockReconnect) {
      lockReconnect = false;
    }
  };
  const reconnect = () => {
    console.log('lockReconnect:' + lockReconnect);
    if (lockReconnect) return;
    lockReconnect = true;
    clearTimeout(timer);
    timer = setTimeout(() => {
      createSocket();
    }, SocketEnum.HeartBeatInterval);
  };

  const init = () => {
    socket.onopen = function (_) {
      console.log('WebSocket:已连接');
      heartCheck.reset().start();
      isActive = true;
    };

    socket.onmessage = function (event) {
      isActive = true;
      // console.log('WebSocket:收到一条消息', event.data);

      let isHeart = false;
      if (!isJsonString(event.data)) {
        console.log('socket message incorrect format:' + JSON.stringify(event));
        return;
      }

      const message = JSON.parse(event.data);
      if (message.event === 'ping') {
        isHeart = true;
      }

      // 强制退出
      if (message.event === 'kick') {
        useUserStore.logout().then(() => {
          // 移除标签页
          localStorage.removeItem(TABS_ROUTES);
          location.reload();
        });
        return;
      }

      // 通知
      if (message.event === 'notice') {
        notificationStore.triggerNewMessages(message.data);
        return;
      }

      if (onMessage && !isHeart) {
        onMessage.call(null, event);
      }
      heartCheck.reset().start();
    };

    socket.onerror = function (_) {
      console.log('WebSocket:发生错误');
      reconnect();
      isActive = false;
    };

    socket.onclose = function (_) {
      console.log('WebSocket:已关闭');
      heartCheck.reset();
      reconnect();
      isActive = false;
    };

    window.onbeforeunload = function () {
      socket.close();
      isActive = false;
    };
  };

  createSocket();
};
