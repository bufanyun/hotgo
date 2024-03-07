<template>
  <Iframe ref="iframeRef" style="zoom: 0.9" />
</template>
<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import Iframe from '../../../../iframe/index.vue';

  const iframeRef = ref();

  onMounted(() => {
    iframeRef.value?.open('https://ppchart.com/#/');
  });
</script>

<style lang="less" scoped></style>
