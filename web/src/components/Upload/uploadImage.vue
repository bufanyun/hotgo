<template>
  <BasicUpload
    :action="`${uploadUrl}${urlPrefix}/upload/image`"
    :headers="uploadHeaders"
    :data="{ type: 0 }"
    name="file"
    :width="100"
    :height="100"
    :maxNumber="maxNumber"
    :helpText="helpText"
    @uploadChange="uploadChange"
    v-model:value="image"
    v-model:values="images"
  />
</template>

<script lang="ts" setup>
  import { onMounted, reactive, ref, unref, watch } from 'vue';
  import { BasicUpload } from '@/components/Upload';
  import { useGlobSetting } from '@/hooks/setting';
  import { useUserStoreWidthOut } from '@/store/modules/user';

  export interface Props {
    value: string | string[] | null;
    maxNumber: number;
    helpText?: string;
  }

  const globSetting = useGlobSetting();
  const urlPrefix = globSetting.urlPrefix || '';
  const { uploadUrl } = globSetting;
  const useUserStore = useUserStoreWidthOut();
  const uploadHeaders = reactive({
    Authorization: useUserStore.token,
  });
  const emit = defineEmits(['update:value']);
  const props = withDefaults(defineProps<Props>(), { value: '', maxNumber: 1, helpText: '' });
  const image = ref<string>('');
  const images = ref<string[]>([]);

  function uploadChange(list: string | string[]) {
    if (props.maxNumber === 1) {
      image.value = unref(list as string);
      emit('update:value', image.value);
    } else {
      images.value = unref(list as string[]);
      emit('update:value', images.value);
    }
  }

  //赋值默认图片显示
  function loadImage() {
    if (props.maxNumber === 1) {
      image.value = props.value as string;
    } else {
      images.value = props.value as string[];
    }
  }

  watch(
    () => props.value,
    () => {
      loadImage();
    },
    {
      immediate: true,
      deep: true,
    }
  );

  onMounted(async () => {
    loadImage();
  });
</script>

<style lang="less"></style>
