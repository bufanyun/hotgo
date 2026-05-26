<template>
  <div>
    <n-card :bordered="false" title="充值记录" class="proCard">
      <n-tabs
        type="card"
        class="card-tabs"
        :value="defaultTab"
        animated
        @before-leave="handleBeforeLeave"
      >
        <n-tab-pane
          :name="item.key.toString()"
          :tab="item.label"
          v-for="item in dict.getOptionUnRef('orderStatus')"
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
  import { loadOptions } from './model';
  import { useDictStore } from '@/store/modules/dict';

  const dict = useDictStore();
  const router = useRouter();
  const defaultTab = ref('-1');

  function handleBeforeLeave(tabName: string) {
    defaultTab.value = tabName;
  }

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      defaultTab.value = router.currentRoute.value.query.type as string;
    }
    loadOptions();
  });
</script>
