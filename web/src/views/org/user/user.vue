<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="后台用户" />
    </div>
    <n-card :bordered="false" class="proCard">
      <n-tabs
        type="card"
        class="card-tabs"
        :value="defaultTab"
        animated
        @before-leave="handleBeforeLeave"
      >
        <n-tab-pane
          :name="item.id.toString()"
          :tab="item.name"
          v-for="item in options.roleTabs"
          :key="item.key"
        >
          <List :type="defaultTab" />
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import List from './list.vue';
  import { useRouter } from 'vue-router';
  import { options } from './model';

  const router = useRouter();
  const defaultTab = ref('-1');

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
  });

  function handleBeforeLeave(tabName: string): boolean | Promise<boolean> {
    defaultTab.value = tabName;
    return true;
  }
</script>
