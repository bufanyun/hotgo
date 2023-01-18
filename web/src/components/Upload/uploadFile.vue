<template>
  <BasicUpload
    :action="`${uploadUrl}${urlPrefix}/upload/file`"
    :headers="uploadHeaders"
    :data="{ type: 0 }"
    name="file"
    :width="100"
    :height="100"
    fileType="file"
    :maxNumber="maxNumber"
    @uploadChange="uploadChange"
    v-model:value="image"
    v-model:values="images"
  />
</template>

<script lang="ts" setup>
  import { ref, onMounted, unref, reactive } from 'vue';
  import { BasicUpload } from '@/components/Upload';
  import { useGlobSetting } from '@/hooks/setting';
  import { useUserStoreWidthOut } from '@/store/modules/user';

  export interface Props {
    value: string | string[] | null;
    maxNumber: number;
  }

  const globSetting = useGlobSetting();
  const urlPrefix = globSetting.urlPrefix || '';
  const { uploadUrl } = globSetting;
  const useUserStore = useUserStoreWidthOut();
  const uploadHeaders = reactive({
    Authorization: useUserStore.token,
  });
  const emit = defineEmits(['update:value']);
  const props = withDefaults(defineProps<Props>(), { value: '', maxNumber: 1 });
  const image = ref<string>('');
  const images = ref<string[] | object>([]);

  function uploadChange(list: string | string[]) {
    if (props.maxNumber === 1) {
      image.value = unref(list as string);
      emit('update:value', image.value);
    } else {
      images.value = unref(list as string[]);
      emit('update:value', images.value);
    }
  }

  onMounted(async () => {
    if (props.maxNumber === 1) {
      image.value = props.value as string;
    } else {
      images.value = props.value as string[];
    }
  });
</script>

<style lang="less"></style>
