<template>
  <div>
    <n-tabs type="line" animated>
      <n-tab-pane v-for="(view, index) in views" :key="index" :name="view.name" :tab="view.name">
        <n-tag :type="view.tag.type" class="tag-margin">
          {{ view.tag.label }}
          <template #icon>
            <n-icon :component="view.tag.icon" />
          </template>
          {{ view.path }}
        </n-tag>
        <n-scrollbar class="code-scrollbar" trigger="none">
          <n-code :code="view.content" :hljs="hljs" language="goLang" show-line-numbers />
        </n-scrollbar>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script lang="ts" setup>
  import { computed } from 'vue';
  import hljs from 'highlight.js/lib/core';
  import goLang from 'highlight.js/lib/languages/go';
  import { cloneDeep } from 'lodash-es';
  import {
    CheckmarkCircle,
    CheckmarkDoneCircle,
    CloseCircleOutline,
    HelpCircleOutline,
    RemoveCircleOutline,
  } from '@vicons/ionicons5';

  hljs.registerLanguage('goLang', goLang);

  interface Props {
    previewModel: any;
    showModal: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    previewModel: cloneDeep({ views: {} }),
    showModal: false,
  });

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
</style>
