<template>
  <n-date-picker v-bind="$props" v-model:value="modelValue" :shortcuts="shortcuts" />
</template>

<script lang="ts">
  import { computed, defineComponent, onMounted, ref } from 'vue';
  import {
    dateToTimestamp,
    formatToDate,
    formatToDateTime,
    timestampToTime,
    defShortcuts,
    defRangeShortcuts,
  } from '@/utils/dateUtil';
  import { basicProps } from './props';

  export default defineComponent({
    name: 'BasicUpload',
    props: {
      ...basicProps,
    },
    emits: ['update:formValue', 'update:startValue', 'update:endValue'],
    setup(props, { emit }) {
      const shortcuts = ref<any>({});

      function getTimestamp(value) {
        let t = dateToTimestamp(value);
        if (t === 0) {
          return new Date().getTime();
        }
        return t;
      }

      function setTimestamp(value) {
        if (!isTimeType()) {
          return formatToDate(new Date(Number(value)).toDateString());
        } else {
          return formatToDateTime(timestampToTime(Number(value / 1000)));
        }
      }

      function isRangeType() {
        return props.type.indexOf('range') != -1;
      }

      function isTimeType() {
        return props.type.indexOf('time') != -1;
      }

      const modelValue = computed({
        get() {
          if (!isRangeType()) {
            return getTimestamp(props.formValue);
          } else {
            return [getTimestamp(props.startValue), getTimestamp(props.endValue)];
          }
        },
        set(value) {
          if (!isRangeType()) {
            emit('update:formValue', setTimestamp(value));
          } else {
            emit('update:startValue', setTimestamp(value[0]));
            emit('update:endValue', setTimestamp(value[1]));
          }
        },
      });

      onMounted(async () => {
        if (!isRangeType()) {
          shortcuts.value = defShortcuts();
        } else {
          shortcuts.value = defRangeShortcuts();
        }
      });

      return {
        modelValue,
        shortcuts,
      };
    },
  });
</script>

<style lang="less"></style>
