<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="提现管理" />
    </div>

    <n-card :bordered="false" class="proCard">
      <n-tabs
        type="card"
        class="card-tabs"
        :value="defaultTab"
        animated
        @before-leave="handleBeforeLeave"
      >
        <n-tab-pane name="" tab="全部"> <List :type="defaultTab" /></n-tab-pane>
        <n-tab-pane name="1" tab="处理中"> <List :type="defaultTab" /> </n-tab-pane>
        <n-tab-pane name="2" tab="提现成功"> <List :type="defaultTab" /> </n-tab-pane>
        <n-tab-pane name="3" tab="提现异常"> <List :type="defaultTab" /> </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import List from './list.vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const defaultTab = ref('');

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
  });

  function handleBeforeLeave(tabName: string) {
    defaultTab.value = tabName;
  }
</script>
