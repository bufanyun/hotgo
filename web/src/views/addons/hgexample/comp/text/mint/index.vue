<template>
  <n-card
    :segmented="{ content: true, footer: true }"
    header-style="padding:10px"
    footer-style="padding:10px"
  >
    <template #header> 敏感词汇验证 </template>
    <n-space vertical>
      <n-space>
        <n-tag
          :type="idx % 2 ? 'success' : 'warning'"
          :key="idx"
          v-for="(item, idx) in compData.keys"
          >{{ item }}</n-tag
        >
      </n-space>
      <n-input
        placeholder="输入需要验证的词汇文本"
        type="textarea"
        size="small"
        v-model:value="compData.text"
        :autosize="{ minRows: 3, maxRows: 5 }"
        @update:value="compData.handleUpdateText"
      />
      <n-space>
        <n-tag
          :type="idx % 2 ? 'error' : 'info'"
          :key="idx"
          v-for="(item, idx) in compData.words"
          >{{ item }}</n-tag
        >
      </n-space>
    </n-space>
  </n-card>
</template>
<script lang="ts" setup>
  import { reactive } from 'vue';
  import Mint from 'mint-filter';

  const compData = reactive({
    keys: ['敏感词', '胡萝卜', '香蕉', '苹果'],
    text: '',
    words: [],
  });
  const mint = new Mint(compData.keys);
  compData.handleUpdateText = (value: string | [string, string]) => {
    const test = mint.filter(value);
    compData.words = test.words;
  };
</script>
