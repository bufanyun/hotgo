<template>
  <n-input-group>
    <n-select
      v-bind="$props"
      :style="{ width: '68%' }"
      v-model:value="memberId"
      :options="memberOption"
      :render-label="renderLabel"
      :filter="handleReceiverFilter"
      @update:value="handleChangeMemberId"
      clearable
      filterable
    />
    <n-select
      :consistent-menu-width="false"
      :style="{ width: '32%', 'min-width': '90px' }"
      :options="selectOptions"
      v-model:value="optVal"
      @update:value="handleUpdateOptions"
    />
  </n-input-group>
</template>

<script setup lang="ts">
  import { h, onMounted, ref, watch } from 'vue';
  import { GetMemberOption } from '@/api/org/user';
  import { useUserStore } from '@/store/modules/user';
  import { NText, SelectRenderLabel } from 'naive-ui';
  import { basicProps } from './props';

  const props = defineProps({
    ...basicProps,
  });
  const emit = defineEmits(['update:value']);
  const userStore = useUserStore();
  const memberId = ref(null);
  const optVal = ref('-1');
  const memberOption = ref([]);
  const selectOptions = [
    {
      label: '查全部',
      value: '-1',
    },
    {
      label: '查本人',
      value: '1',
    },
    {
      label: '查下级',
      value: '2',
    },
  ];

  const renderLabel: SelectRenderLabel = (option: { username: ''; label: '' }) => {
    return h(
      'div',
      {
        style: {
          display: 'flex',
          alignItems: 'center',
        },
      },
      [
        h(
          'div',
          {
            style: {
              padding: '4px 0',
              display: 'flex',
            },
          },
          [
            h(
              'div',
              {
                style: {
                  marginRight: '4px',
                },
              },
              [option.username as string]
            ),
            h(
              NText,
              { depth: 3, tag: 'div' },
              {
                default: () => option.label as string,
              }
            ),
          ]
        ),
      ]
    );
  };

  function handleReceiverFilter(pattern: string, option: object): boolean {
    const isPatternInLabel = option.label.includes(pattern);
    const isPatternInUsername = option.username.includes(pattern);
    const isValueEqual = option.value.toString() === pattern;
    return isPatternInLabel || isPatternInUsername || isValueEqual;
  }

  function handleChangeMemberId(v: string) {
    memberId.value = v;
    handleUpdateValue();
  }

  function handleUpdateOptions(value: string) {
    optVal.value = value;
    handleUpdateValue();
  }

  function handleUpdateValue() {
    if (memberId.value) {
      emit('update:value', [memberId.value, optVal.value]);
    } else {
      emit('update:value', null);
    }
  }

  function resetFields() {
    memberId.value = null;
    optVal.value = '-1';
  }

  watch(
    () => props.value,
    (v) => {
      if (!v) {
        resetFields();
      }
    },
    { deep: true }
  );

  onMounted(() => {
    GetMemberOption().then((res) => {
      if (res) {
        memberOption.value = res;
      }
    });
  });
</script>

<style scoped lang="less"></style>
