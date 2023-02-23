<template>
  <n-form class="py-4">
    <n-form-item label="测试入口" path="index">
      <n-input-group>
        <n-input placeholder="请输入" :default-value="url" :disabled="true" />
        <n-button v-copy="url" type="primary" @click="copy"> 复制链接 </n-button>
      </n-input-group>
    </n-form-item>

    <n-form-item label="二维码" path="index">
      <div class="text-center">
        <qrcode-vue :value="url" :size="220" class="canvas" style="margin: 0 auto" />
      </div>
    </n-form-item>
  </n-form>
</template>

<script lang="ts" setup>
  import { useUserStoreWidthOut } from '@/store/modules/user';
  import QrcodeVue from 'qrcode.vue';
  import { useMessage } from 'naive-ui';

  const message = useMessage();

  interface Props {
    path: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    path: '',
  });

  const copy = () => {
    message.success('复制成功');
  };

  const useUserStore = useUserStoreWidthOut();
  const url = useUserStore.config?.domain + props.path;
</script>

<style scoped>
  ::v-deep(.card-tabs .n-tabs-nav--bar-type) {
    padding-left: 4px;
  }
</style>
