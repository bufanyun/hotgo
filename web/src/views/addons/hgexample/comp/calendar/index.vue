<template>
  <div class="n-layout-page-header">
    <n-card :bordered="false" title="日历"> 一个普通的日历 </n-card>
  </div>
  <n-card :bordered="false" class="mt-4 proCard">
    <n-calendar
      v-model:value="value"
      #="{ year, month, date }"
      :is-date-disabled="isDateDisabled"
      @update:value="handleUpdateValue"
    >
      {{ year }}-{{ month }}-{{ date }}
    </n-calendar>
  </n-card>
</template>

<script lang="ts">
  import { defineComponent, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { isYesterday, addDays } from 'date-fns';

  export default defineComponent({
    setup() {
      const message = useMessage();
      return {
        value: ref(addDays(Date.now(), 1).valueOf()),
        handleUpdateValue(
          _: number,
          { year, month, date }: { year: number; month: number; date: number }
        ) {
          message.success(`${year}-${month}-${date}`);
        },
        isDateDisabled(timestamp: number) {
          if (isYesterday(timestamp)) {
            return true;
          }
          return false;
        },
      };
    },
  });
</script>
