<template>
  <n-card
    :content-style="{ padding: '0px' }"
    :footer-style="{ padding: '0px' }"
    :bordered="false"
    :segmented="true"
  >
    <div v-if="notificationStore.messages.length > 0">
      <div
        class="flex items-center max-w-sm p-1 mx-auto space-x-2 rounded-xl"
        v-for="(item, index) of notificationStore.messages"
        :key="index"
      >
        <div class="flex-shrink-0">
          <n-icon size="40" color="#f00">
            <NotificationsCircle />
          </n-icon>
        </div>
        <div>
          <div class="text-sm font-medium">{{ item.title }}</div>
          <n-ellipsis :line-clamp="1" class="text-gray-500">{{ item.content }}</n-ellipsis>
        </div>
      </div>
    </div>
    <n-empty v-else description="暂无消息哦~" class="pt-20 pb-20" />
    <template #footer>
      <div class="flex justify-evenly">
        <n-button type="text" @click="onClearMessage">清空提醒</n-button>
        <n-button type="text" @click="onAllMessage">查看更多</n-button>
      </div>
    </template>
  </n-card>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import { NotificationsCircle } from '@vicons/ionicons5';
  import { notificationStoreWidthOut } from '@/store/modules/notification';
  import { useRouter } from 'vue-router';

  export default defineComponent({
    name: 'PopoverMessage',
    components: { NotificationsCircle },
    emits: ['clear'],
    setup(_props, { emit }) {
      const notificationStore = notificationStoreWidthOut();
      const router = useRouter();

      function onClearMessage() {
        notificationStore.setMessages([]);
        emit('clear');
      }

      function onAllMessage() {
        router.push({ name: 'apply_notice' });
      }

      return {
        onClearMessage,
        notificationStore,
        onAllMessage,
      };
    },
  });
</script>
