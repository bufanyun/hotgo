<template>
  <div>
    <textarea id="copy-code" :value="content"></textarea>
    <n-tabs type="line" animated>
      <n-tab-pane v-for="(view, index) in views" :key="index" :name="view.name" :tab="view.name">
        <n-space justify="space-between">
          <n-tag :type="view.tag.type" class="tag-margin">
            {{ view.tag.label }}
            <template #icon>
              <n-icon :component="view.tag.icon" />
            </template>
            {{ view.path }}
          </n-tag>
          <n-button type="primary" size="small" class="tag-margin" @click="handleCopy(view.content)"
            >复制本页代码</n-button
          >
        </n-space>
        <n-scrollbar class="code-scrollbar" trigger="none">
          <n-code
            :class="'code-' + getFileExtension(view.path)"
            :code="view.content"
            :hljs="hljs"
            :language="getFileExtension(view.path)"
            show-line-numbers
          />
        </n-scrollbar>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import hljs from 'highlight.js/lib/core';
  import go from 'highlight.js/lib/languages/go';
  import typescript from 'highlight.js/lib/languages/typescript';
  import xml from 'highlight.js/lib/languages/xml';
  import sql from 'highlight.js/lib/languages/sql';
  import { cloneDeep } from 'lodash-es';
  import {
    CheckmarkCircle,
    CheckmarkDoneCircle,
    CloseCircleOutline,
    HelpCircleOutline,
    RemoveCircleOutline,
  } from '@vicons/ionicons5';
  import { useMessage } from 'naive-ui';

  hljs.registerLanguage('go', go);
  hljs.registerLanguage('ts', typescript);
  hljs.registerLanguage('sql', sql);
  hljs.registerLanguage('vue', xml);

  interface Props {
    previewModel: any;
    showModal: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    previewModel: cloneDeep({ views: {} }),
    showModal: false,
  });
  const message = useMessage();
  const content = ref('');
  const views = computed(() => {
    let tmpViews: any = [];
    let i = 0;
    for (const [k, v] of Object.entries(props.previewModel.views)) {
      let item = v as any;
      item.name = k;
      switch (item.meth) {
        case 1:
          item.tag = { type: 'success', label: '创建文件', icon: CheckmarkCircle };
          break;
        case 2:
          item.tag = { type: 'warning', label: '覆盖文件', icon: CheckmarkDoneCircle };
          break;
        case 3:
          item.tag = { type: 'info', label: '已存在跳过', icon: CloseCircleOutline };
          break;
        case 4:
          item.tag = { type: 'error', label: '不生成', icon: RemoveCircleOutline };
          break;
        default:
          item.tag = { type: 'error', label: '未知状态', icon: HelpCircleOutline };
      }
      tmpViews[i] = item;
      i++;
    }
    return tmpViews;
  });

  function getFileExtension(path: string): string {
    const parts = path.split('.');
    if (parts.length > 1) {
      return parts[parts.length - 1];
    }
    return '';
  }

  function handleCopy(code: string) {
    content.value = code;
    setTimeout(function () {
      const copyVal = document.getElementById('copy-code');
      copyVal.select();
      document.execCommand('copy');
      message.success('已复制');
    }, 20);
  }
</script>

<style lang="less" scoped>
  ::v-deep(.alert-margin) {
    margin-bottom: 20px;
  }

  ::v-deep(.tag-margin) {
    margin-bottom: 10px;
  }

  ::v-deep(.code-scrollbar) {
    height: calc(100vh - 300px);
    background: #282b2e;
    color: #e0e2e4;
    padding: 10px;
  }

  ::v-deep(.code-vue .hljs-tag) {
    color: rgb(242, 197, 92);
  }

  ::v-deep(.code-vue .hljs-name) {
    color: rgb(242, 197, 92);
  }

  ::v-deep(.code-vue .hljs-attr) {
    color: rgb(49, 104, 213);
  }

  ::v-deep(.code-go .hljs-params) {
    color: rgb(49, 104, 213);
  }

  ::v-deep(.code-ts .hljs-params) {
    color: rgb(49, 104, 213);
  }

  ::v-deep(.code-ts .hljs-property) {
    color: rgb(49, 104, 213);
  }

  ::v-deep(.code-ts .hljs-function) {
    color: rgb(49, 104, 213);
  }

  #copy-code {
    position: fixed;
    top: -100px;
    left: -100px;
  }
</style>
