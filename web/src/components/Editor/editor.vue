<template>
  <QuillEditor
    ref="quillEditor"
    :options="options"
    v-model:content="content"
    @ready="readyQuill"
    class="quillEditor"
    :id="quillEditorId"
  />
</template>

<script lang="ts" setup>
  import { ref, watch, onMounted } from 'vue';
  import { QuillEditor } from '@vueup/vue-quill';
  import '@vueup/vue-quill/dist/vue-quill.snow.css';
  import { getRandomString } from '@/utils/charset';
  export interface Props {
    value: string;
  }

  const emit = defineEmits(['update:value']);
  const quillEditorId = ref('quillEditorId-' + getRandomString(16, true));
  const quillEditor = ref();
  const content = ref();
  const props = withDefaults(defineProps<Props>(), { value: '' });
  const options = ref({
    modules: {
      toolbar: [
        ['bold', 'italic', 'underline', 'strike'], // toggled buttons
        ['blockquote', 'code-block'],

        [{ header: 1 }, { header: 2 }], // custom button values
        [{ list: 'ordered' }, { list: 'bullet' }],
        [{ script: 'sub' }, { script: 'super' }], // superscript/subscript
        [{ indent: '-1' }, { indent: '+1' }], // outdent/indent
        [{ direction: 'rtl' }], // text direction

        [{ size: ['small', false, 'large', 'huge'] }], // custom dropdown
        [{ header: [1, 2, 3, 4, 5, 6, false] }],

        [{ color: [] }, { background: [] }], // dropdown with defaults from theme
        [{ font: [] }],
        [{ align: [] }],
        ['clean'],
        ['image'],
      ],
    },
    theme: 'snow',
    placeholder: '输入您要编辑的内容！',
  });

  function readyQuill() {
    quillEditor.value.setHTML(props.value);
  }

  watch(
    () => content.value,
    (_newValue, _oldValue) => {
      if (quillEditor.value !== undefined) {
        emit('update:value', quillEditor.value.getHTML());
      }
    },
    {
      immediate: true, // 深度监听
    }
  );

  onMounted(async () => {
    // 兼容表单分组 n-form-item-blank
    let dom = document.getElementById(quillEditorId.value);
    if (dom && dom.parentNode) {
      const parent = dom.parentNode as Element;
      if ('n-form-item-blank' === parent.className) {
        parent.setAttribute('style', 'display: block;');
      }
    }
  });
</script>

<style lang="less">
  .ql-container {
    height: auto;
  }
</style>
