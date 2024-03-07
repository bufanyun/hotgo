<template>
  <div>
    <div>
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="图标选择器">
          图标选择器，用于表单选择图标，支持直接选定图标和输入图标名称来选择
        </n-card>
      </div>
      <n-space>
        <n-card :bordered="false" class="mt-4 proCard" title="Antd">
          <IconSelector style="width: 100%" v-model:value="antdValue" option="antd" />
        </n-card>

        <n-card :bordered="false" class="mt-4 proCard" title="Ionicons5">
          <IconSelector style="width: 100%" v-model:value="ionicons5Value" option="ionicons5" />
        </n-card>
      </n-space>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import IconSelector from '@/components/IconSelector/index.vue';

  const antdValue = ref('');
  const ionicons5Value = ref('');
</script>
<style lang="less" scoped></style>
