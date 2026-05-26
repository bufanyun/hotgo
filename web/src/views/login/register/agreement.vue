<template>
  <div>
    <n-checkbox v-model:checked="checked" class="text-12px">我已阅读并接受</n-checkbox>
    <n-button text type="primary" @click="handleClickProtocol" class="text-12px">
      《用户协议》
    </n-button>
    <n-button text type="primary" @click="handleClickPolicy" class="text-12px">
      《隐私权政策》
    </n-button>
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue';

  interface Props {
    value?: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    value: true,
  });

  interface Emits {
    (e: 'update:value', value: boolean): void;
    (e: 'click-protocol'): void;
    (e: 'click-policy'): void;
  }

  const emit = defineEmits<Emits>();

  const checked = computed({
    get() {
      return props.value;
    },
    set(newValue: boolean) {
      emit('update:value', newValue);
    },
  });

  function handleClickProtocol() {
    emit('click-protocol');
  }

  function handleClickPolicy() {
    emit('click-policy');
  }
</script>

<style scoped>
  .text-12px {
    font-size: 12px;
  }
</style>
