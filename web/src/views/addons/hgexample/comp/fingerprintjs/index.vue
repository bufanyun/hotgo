<template>
  <n-card :segmented="{ content: true, footer: true }" footer-style="padding:10px">
    <template #header>
      通过设备浏览器信息获取浏览器指纹的插件(官方宣称其识别精度达到99.5%)
    </template>
    <div>
      指纹ID:
      <n-text type="info">
        {{ compData.murmur }}
      </n-text>
    </div>
  </n-card>
</template>
<script lang="ts" setup>
  import { reactive } from 'vue';
  import Fingerprint2 from 'fingerprintjs2';

  const compData = reactive({
    values: {},
    murmur: '',
  });

  const createFingerprint = () => {
    Fingerprint2.get((components) => {
      compData.values = components.map((component) => component.value); // 配置的值的数组
      compData.murmur = Fingerprint2.x64hash128(compData.values.join(''), 31).toUpperCase(); // 生成浏览器指纹
    });
  };
  if (window.requestIdleCallback) {
    requestIdleCallback(() => {
      createFingerprint();
    });
  } else {
    setTimeout(() => {
      createFingerprint();
    }, 600);
  }
</script>
