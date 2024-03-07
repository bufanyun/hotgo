## WebSocket客户端

目录

- 全局消息监听
- 单页面消息监听
- 发送消息

> 基于WebSocket服务器，hotgo还对客户端的上做了一些封装，使其使用起来更加方便
- [WebSocket服务器](sys-websocket-server.md)

###  全局消息监听
- 所有全局的消息监听都在这里
- 文件路径：web/src/utils/websocket/registerMessage.ts
```ts
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

```

#### 单页面消息监听
- 当你只需要某个页面使用WebSocket，这将是一个不错的选择，下面是一个简单的演示例子
- 文件路径：web/src/views/addons/hgexample/portal/websocketTest.vue
```vue
<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="测试websocket">
        尝试在下方输入框中输入任意文字消息内容，发送后websocket服务器收到会原样返回
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <n-space vertical>
        <n-input-group style="width: 520px">
          <n-input
            @keyup.enter="sendMessage"
            :style="{ width: '78%' }"
            placeholder="请输入消息内容"
            :on-focus="onFocus"
            :on-blur="onBlur"
            v-model:value="inputMessage"
          />
          <n-button type="primary" @click="sendMessage"> 发送消息</n-button>
        </n-input-group>

        <div class="mt-5"></div>

        <n-timeline :icon-size="20">
          <n-timeline-item color="grey" content="输入中.." v-if="isInput">
            <template #icon>
              <n-icon>
                <MessageOutlined />
              </n-icon>
            </template>
          </n-timeline-item>

          <n-timeline-item
            v-for="item in messages"
            :key="item"
            :type="item.type == Enum.SendType ? 'success' : 'info'"
            :title="item.type == Enum.SendType ? '发送消息' : '收到消息'"
            :content="item.content"
            :time="item.time"
          >
            <template #icon>
              <n-icon>
                <SendOutlined v-if="item.type == Enum.SendType" />
                <SoundOutlined v-if="item.type == Enum.ReceiveType" />
              </n-icon>
            </template>
          </n-timeline-item>
        </n-timeline>
      </n-space>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onBeforeUnmount, onMounted, ref } from 'vue';
  import { MessageOutlined, SendOutlined, SoundOutlined } from '@vicons/antd';
  import { format } from 'date-fns';
  import { addOnMessage, removeOnMessage, sendMsg, WebSocketMessage } from '@/utils/websocket';
  import { useMessage } from 'naive-ui';

  interface Message {
    type: Enum;
    content: string;
    time: string;
  }

  const message = useMessage();
  const messages = ref<Message[]>([]);
  const inputMessage = ref('你好，HotGo');
  const isInput = ref(false);
  const testMessageEvent = 'admin/addons/hgexample/testMessage';

  enum Enum {
    SendType = 1, // 发送类型
    ReceiveType = 2, // 接受类型
  }

  function onFocus() {
    isInput.value = true;
  }

  function onBlur() {
    isInput.value = false;
  }

  function sendMessage() {
    if (inputMessage.value == '') {
      message.error('消息内容不能为空');
      return;
    }

    // 发送消息
    sendMsg(testMessageEvent, {
      message: inputMessage.value,
    });

    const msg: Message = {
      type: Enum.SendType,
      content: inputMessage.value,
      time: format(new Date(), 'yyyy-MM-dd HH:mm:ss'),
    };
    insertMessage(msg);
    inputMessage.value = '';
  }

  // 存入本地记录
  function insertMessage(msg: Message): void {
    messages.value.unshift(msg); // 在头部插入消息
    if (messages.value.length > 10) {
      messages.value = messages.value.slice(0, 10); // 如果超过10个，则只保留最前面10个
    }
  }

  const onMessage = (res: WebSocketMessage) => {
    const msg: Message = {
      type: Enum.ReceiveType,
      content: res.data.message,
      time: format(new Date(), 'yyyy-MM-dd HH:mm:ss'),
    };
    insertMessage(msg);
  };

  onMounted(() => {
    // 在当前页面注册消息监听
    addOnMessage(testMessageEvent, onMessage);
  });

  onBeforeUnmount(() => {
    // 移除消息监听
    removeOnMessage(testMessageEvent);
  });
</script>

<style scoped></style>
```

#### 发送消息
- 向服务器发送一条消息
```ts
  import { sendMsg } from '@/utils/websocket';

    const event = 'admin/addons/hgexample/testMessage'; // 消息路由
    const data: object | null = {  // 消息内容
        message: 'message content...',
    };
    const isRetry = false; // 发送失败是否重试，不传默认为true

    // 基本使用 
    sendMsg(event, data);

    // 无消息内容 
    sendMsg(event);

    // 发送失败不重试 
    sendMsg(event, data, isRetry);
```
