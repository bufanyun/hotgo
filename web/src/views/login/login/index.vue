<template>
  <n-tabs type="segment" justify-content="space-evenly">
    <n-tab-pane name="account" tab="账号登录">
      <Form @updateActiveModule="updateActiveModule" mode="account" />
    </n-tab-pane>
    <n-tab-pane name="mobile" tab="手机号登录">
      <Form @updateActiveModule="updateActiveModule" mode="mobile" />
    </n-tab-pane>
  </n-tabs>
</template>

<script lang="ts" setup>
  import Form from './form.vue';

  const emit = defineEmits(['updateActiveModule']);

  function updateActiveModule(key: string) {
    emit('updateActiveModule', key);
  }
</script>

<style lang="less" scoped></style>
