<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="tagsView"> 标签页操作项 </n-card>
    </div>
    <n-card>
      <n-space>
        <n-button type="info" @click="handleSignal('1')"> 刷新页面路由 </n-button>
        <n-button type="error" @click="handleSignal('2')"> 关闭当前页面 </n-button>
        <n-button type="warning" @click="handleSignal('3')"> 关闭其他页面 </n-button>
        <n-button type="success" @click="handleSignal('4')"> 关闭全部页面 </n-button>
      </n-space>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { useTabsViewStore } from '@/store/modules/tabsView';

  const tabsViewStore = useTabsViewStore();

  function handleSignal(signal: string) {
    tabsViewStore.closeSignal(signal);
  }
</script>
<style lang="less" scoped></style>
