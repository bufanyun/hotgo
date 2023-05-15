<template>
  <QuillEditor
    ref="quillEditorRef"
    toolbar="full"
    v-model:content="content"
    @ready="readyQuill"
    class="quillEditor"
    :id="id"
    :modules="modules"
    @focus="onEditorFocus"
    @blur="onEditorBlur"
    @update:content="onUpdateContent"
  />
</template>

<script lang="ts" setup>
  import { ref, watch, onMounted } from 'vue';
  import { QuillEditor } from '@vueup/vue-quill';
  import '@vueup/vue-quill/dist/vue-quill.snow.css';
  import ImageUploader from 'quill-image-uploader';
  import MagicUrl from 'quill-magic-url';
  import { getRandomString } from '@/utils/charset';
  import { UploadImage } from '@/api/base';
  import componentSetting from '@/settings/componentSetting';
  import { isNullOrUnDef } from '@/utils/is';
  import { useMessage } from 'naive-ui';

  export interface Props {
    value: string;
    id?: string;
  }

  const emit = defineEmits(['update:value']);
  const message = useMessage();
  const initFinish = ref(false);
  const quillEditorRef = ref();
  const content = ref();
  const props = withDefaults(defineProps<Props>(), {
    value: '',
    id: 'quillEditorId-' + getRandomString(16, true),
  });

  function readyQuill() {
    quillEditorRef.value.setHTML(props.value);
  }

  watch(
    () => props.value,
    (newValue) => {
      if (!initFinish.value) {
        quillEditorRef.value?.setHTML(newValue);
      }
    },
    {
      immediate: true,
      deep: true,
    }
  );

  function onEditorFocus(val) {
    initFinish.value = true;
    console.log(val);
  }

  function onEditorBlur(val) {
    console.log(val);
  }

  function onUpdateContent() {
    emit('update:value', quillEditorRef.value.getHTML());
  }

  function checkFileType(map: string[], fileType: string) {
    if (isNullOrUnDef(map)) {
      return true;
    }
    return map.includes(fileType);
  }

  onMounted(async () => {
    // 兼容表单分组 n-form-item-blank
    let dom = document.getElementById(props.id);
    if (dom && dom.parentNode) {
      const parent = dom.parentNode as Element;
      if ('n-form-item-blank' === parent.className) {
        parent.setAttribute('style', 'display: block;');
      }
    }
  });

  const modules = [
    {
      name: 'imageUploader',
      module: ImageUploader,
      options: {
        upload: (file) => {
          return new Promise((resolve, reject) => {
            if (!checkFileType(componentSetting.upload.imageType, file.type)) {
              message.error(`只能上传图片类型为${componentSetting.upload.imageType.join(',')}`);
              reject('Upload failed');
              return;
            }

            const formData = new FormData();
            formData.append('file', file);
            UploadImage(formData)
              .then((res) => {
                console.log(res);
                resolve(res.fileUrl);
              })
              .catch((err) => {
                reject('Upload failed');
                console.error('Error:', err);
              });
          });
        },
      },
    },
    {
      name: 'magicUrl',
      module: MagicUrl,
    },
  ];
</script>

<style lang="less" scoped>
  :deep(.ql-container) {
    height: auto;
  }
  :deep(.ql-container.ql-snow) {
    border: none;
  }
  :deep(.ql-toolbar.ql-snow) {
    border: none;
    border-bottom: 1px solid #ccc;
  }
  :deep(.ql-editor.ql-blank::before) {
    color: #afb4bd;
    font-size: 14px;
    font-style: normal;
  }
  .dark .priview-content {
    background: #5a5a5a;
    color: #fff;
  }
  .light .priview-content {
    background: #fff;
    color: #333;
  }
</style>
