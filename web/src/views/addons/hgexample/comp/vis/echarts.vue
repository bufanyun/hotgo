<template>
  <Iframe ref="iframeRef" style="zoom: 0.8" />
</template>
<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import Iframe from '../../../../iframe/index.vue';

  const iframeRef = ref();

  onMounted(() => {
    iframeRef.value?.open('https://www.makeapie.cn/echarts');
  });
</script>

<style lang="less" scoped></style>
