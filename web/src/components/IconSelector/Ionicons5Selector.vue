<template>
  <n-popover trigger="click" placement="bottom" width="400">
    <template #trigger>
      <n-button>
        <template #icon>
          <n-icon size="20">
            <component :is="formValue !== '' ? formValue : 'LogoIonic'" />
          </n-icon>
        </template>
      </n-button>
    </template>
    <n-scrollbar class="grid-wrapper">
      <n-grid :cols="8" :collapsed="false" responsive="screen" style="height: 300px">
        <n-grid-item v-for="(item, index) of icons" :key="index">
          <div
            class="flex flex-col items-center justify-center p-3 icon-wrapper"
            @click="onIconClick(item)"
          >
            <n-icon size="20">
              <component :is="item" />
            </n-icon>
          </div>
        </n-grid-item>
      </n-grid>
    </n-scrollbar>
    <div class="flex justify-end mt-2 mb-2">
      <n-pagination
        :page="currentPage"
        :page-size="pageSize"
        :page-slot="8"
        :item-count="itemCount"
        @update-page="onUpdatePage"
      />
    </div>
  </n-popover>
</template>

<script lang="ts">
  import { computed, defineComponent, ref, shallowReactive } from 'vue';
  import * as Ionicons5Icons from '@vicons/ionicons5';
  export default defineComponent({
    name: 'Ionicons5Selector',
    components: Ionicons5Icons,
    props: {
      value: String,
      option: String,
    },
    emits: ['update:value'],
    setup(props, { emit }) {
      const formValue = computed({
        get() {
          return props.value;
        },
        set(value) {
          emit('update:value', value);
        },
      });

      const iconArray = Object.keys(Ionicons5Icons);
      const pageSize = 40;
      const icons = shallowReactive(iconArray.slice(0, 40));
      const currentPage = ref(1);
      const itemCount = computed(() => iconArray.length);

      function onUpdatePage(page: number) {
        currentPage.value = page;
        icons.length = 0;
        const start = (currentPage.value - 1) * pageSize;
        icons.push(...iconArray.slice(start, start + pageSize));
      }

      function onIconClick(item: any) {
        formValue.value = item;
      }
      return {
        icons,
        currentPage,
        pageSize,
        itemCount,
        onUpdatePage,
        onIconClick,
        formValue,
      };
    },
  });
</script>
<style lang="less" scoped>
  .grid-wrapper {
    .icon-wrapper {
      cursor: pointer;
      border: 1px solid #f5f5f5;
    }
  }
</style>
