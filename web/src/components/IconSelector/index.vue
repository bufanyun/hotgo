<template>
  <div>
    <n-input-group>
      <template v-if="option === 'ionicons5'">
        <Ionicons5Selector v-model:value="formValue" />
      </template>
      <template v-else>
        <AntdSelector v-model:value="formValue" />
      </template>
      <n-input
        v-bind="$props"
        :value="formValue"
        :style="{ width: '70%' }"
        placeholder="请选择图标"
      />
    </n-input-group>
  </div>
</template>

<script lang="ts">
  import { defineComponent, computed } from 'vue';
  import { basicProps } from '@/components/IconSelector/props';
  import Ionicons5Selector from '@/components/IconSelector/Ionicons5Selector.vue';
  import AntdSelector from '@/components/IconSelector/AntdSelector.vue';
  export default defineComponent({
    name: 'BasicUpload',
    components: { Ionicons5Selector, AntdSelector },
    props: {
      ...basicProps,
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

      return {
        formValue,
      };
    },
  });
</script>

<style lang="less"></style>
