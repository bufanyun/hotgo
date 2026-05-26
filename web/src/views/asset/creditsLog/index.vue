<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="资金变动">
        你和下级在平台中余额、积分的变动明细都可以在这里进行查看
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
        <n-tab-pane name="" tab="全部"> <List :type="defaultTab" /></n-tab-pane>
        <n-tab-pane name="balance" tab="余额明细"> <List :type="defaultTab" /> </n-tab-pane>
        <n-tab-pane name="integral" tab="积分明细"> <List :type="defaultTab" /> </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import List from './list.vue';
  import { useRouter } from 'vue-router';
  import { loadOptions } from './model';

  const router = useRouter();
  const defaultTab = ref('');

  function handleBeforeLeave(tabName: string) {
    defaultTab.value = tabName;
  }

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
    loadOptions();
  });
</script>

<style lang="less" scoped></style>
