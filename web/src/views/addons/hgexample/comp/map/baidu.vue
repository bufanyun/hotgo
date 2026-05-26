<template>
  <div class="container" ref="container"></div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import useCreateScript from '@/hooks/useCreateScript';

  const SCRIPT_URL =
    'http://api.map.baidu.com/getscript?v=3.0&ak=WxbQmaOc3bvSGSaKWcyeFSf8fnYCWpKd&services=&t=' +
    new Date().getTime();
  const container = ref<HTMLDivElement | null>(null);
  const { createScriptPromise } = useCreateScript(SCRIPT_URL);
  const initMap = () => {
    createScriptPromise.then(() => {
      const bMap = (window as any).BMap;
      const map: any = new bMap.Map(container.value);
      const point = new bMap.Point(116.404, 39.915);
      map.centerAndZoom(point, 7);
      map.enableScrollWheelZoom();
      map.setMapStyleV2({ styleId: 'ea4652613f3629247d47666706ce7e89' });
    });
  };
  onMounted(initMap);
</script>

<style lang="less">
  .container {
    width: 100%;
    height: 720px;
    margin-top: -32px;
  }
</style>
