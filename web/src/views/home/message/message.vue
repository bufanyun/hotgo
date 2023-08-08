<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="我的消息">
        在这里，您可以方便地查看平台中的通知、公告和与您相关的私信消息
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <n-tabs
        type="card"
        class="card-tabs"
        :value="defaultTab"
        animated
        @before-leave="handleBeforeLeave"
      >
        <n-tab-pane name="1" tab="通知"> <List :type="defaultTab" /></n-tab-pane>
        <n-tab-pane name="2" tab="公告"> <List :type="defaultTab" /> </n-tab-pane>
        <n-tab-pane name="3" tab="私信"> <List :type="defaultTab" /> </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import List from './list.vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const defaultTab = ref('1');

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
  });

  function handleBeforeLeave(tabName: string) {
    defaultTab.value = tabName;
  }
</script>
