<template>
  <n-cascader
    v-bind="$props"
    :value="valueLabel"
    :options="dataOptions"
    :placeholder="placeholder"
    :check-strategy="checkStrategy"
    clearable
    cascade
    :on-update:value="onValueChange"
    :on-load="handleLoad"
    :on-focus="focusLoad"
    remote
  />
</template>

<script lang="ts" setup>
  import { GetCityLabel, ProvincesSelect } from '@/api/apply/provinces';
  import { computed, ref, watch } from 'vue';
  import type { CascaderOption } from 'naive-ui';
  const emits = defineEmits(['update:value', 'update:label']);
  import { basicProps } from './props';
  const props = defineProps({
    ...basicProps,
  });

  const valueLabel = ref<string | null>(null);
  const dataOptions = ref([]);
  const placeholder = computed(() => {
    if (props.dataType === 'p') {
      return '请选择省份';
    } else if (props.dataType === 'pc') {
      return '请选择省市';
    } else {
      return '请选择省市区';
    }
  });

  function onValueChange(
    value: string | number | Array<string | number> | null,
    option: CascaderOption | Array<CascaderOption | null> | null,
    pathValues: Array<CascaderOption | null>
  ) {
    const tempPathValues = pathValues
      ? pathValues.map((it: CascaderOption | null) => ({
          label: it?.label,
          value: it?.value,
          level: it?.level,
        }))
      : null;

    emits('update:value', value);
    valueLabel.value = getLabel(tempPathValues);
  }

  function getLabel(values): string | null {
    if (values === null || values === undefined) {
      return null;
    }
    let label = '';
    const length = values.length;
    for (let i = 0; i < length; i++) {
      const item = values[i];
      label += item.label;
      if (i + 1 < length) {
        label += props.separator;
      }
    }
    return label;
  }

  watch(
    () => props.value,
    async () => {
      if (!props.value || props.value === 0) {
        valueLabel.value = null;
        return;
      }

      if (valueLabel.value === null) {
        valueLabel.value = await GetCityLabel({ id: props.value, spilt: props.separator });
      }
    },
    {
      immediate: true,
      deep: true,
    }
  );

  async function load(option) {
    const data = await ProvincesSelect({ dataType: props.dataType, ...option });
    return data.list;
  }

  async function handleLoad(option: CascaderOption) {
    option.children = await load({ dataType: props.dataType, ...option });
    return;
  }

  async function focusLoad() {
    if (dataOptions.value.length === 0) {
      dataOptions.value = await load({ dataType: props.dataType });
    }
  }
</script>
