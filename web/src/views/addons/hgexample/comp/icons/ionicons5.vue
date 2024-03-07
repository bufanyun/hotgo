<template>
  <div>
    <n-spin :show="loading" description="请稍候..">
      <Icons ref="iconsRef" />
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import Icons from './icons.vue';

  const iconsRef = ref();
  const loading = ref(true);

  onMounted(() => {
    setTimeout(() => {
      iconsRef.value?.open('ionicons5');
      loading.value = false;
    }, 10);
  });
</script>
<style lang="less" scoped></style>
