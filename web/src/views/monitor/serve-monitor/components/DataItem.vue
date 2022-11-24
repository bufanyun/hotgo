<template>
  <n-card :content-style="{ padding: '10px' }" :header-style="{ padding: '5px' }" :segmented="true">
    <template #header>
      <n-skeleton text v-if="loading" width="60%" />
      <template v-else>
        <div class="flex items-center justify-between">
          <span class="text-sm text-bold">{{ dataModel.title }}</span>
          <n-icon style="font-size: 26px">
            <div v-if="dataModel.iconClass === 'HardwareChip'">
              <HardwareChip />
            </div>
            <div v-else-if="dataModel.iconClass === 'AppsSharp'">
              <AppsSharp />
            </div>
            <div v-else-if="dataModel.iconClass === 'Analytics'">
              <Analytics />
            </div>
            <div v-else-if="dataModel.iconClass === 'PieChart'">
              <PieChart />
            </div>

            <div v-else>
              <Bookmark />
            </div>
          </n-icon>
        </div>
      </template>
    </template>
    <n-skeleton text v-if="loading" :repeat="6" />
    <template v-else>
      <div style="height: 130px" class="flex flex-col justify-between">
        <div class="flex flex-col justify-center">
          <span class="text-xxl">{{ dataModel.data }}</span>
        </div>
        <div class="flex flex-col justify-center flex-1">
          <slot name="extra" :extra="dataModel.extra"></slot>
        </div>
        <div class="divide"></div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-grey">{{ dataModel.bottomTitle }}</span>
          <span class="text-sm text-grey">{{ dataModel.totalSum }}</span>
        </div>
      </div>
    </template>
  </n-card>
</template>

<script lang="ts">
  import { defineComponent, ref } from 'vue';
  import { HardwareChip, Bookmark, AppsSharp, PieChart, Analytics } from '@vicons/ionicons5';

  export default defineComponent({
    name: 'DataItem',
    components: {
      Bookmark,
      HardwareChip,
      AppsSharp,
      PieChart,
      Analytics,
    },
    props: {
      dataModel: {
        type: Object,
        default: () => {
          return {};
        },
      },
      loading: {
        type: Boolean,
        default: () => {
          return false;
        },
      },
    },
    setup() {
      // const loading = ref(true);
      // setTimeout(() => {
      //   loading.value = false;
      // }, 1000);
      return {
        // loading,
        Bookmark,
        AppsSharp,
        PieChart,
        HardwareChip,
        Analytics,
      };
    },
  });
</script>

<style lang="less" scoped>
  .divide {
    margin: 10px 0;
    border-bottom: 1px solid #f5f5f5;
  }
</style>
