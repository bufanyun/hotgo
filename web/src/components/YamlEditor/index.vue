<template>
  <div class="yaml-editor-container">
    <div
      ref="editorRef"
      class="yaml-editor"
      :class="{ 'error-state': hasError }"
    />
    <div v-if="errorMessage" class="error-message">
      <n-alert type="error" closable @close="clearError">
        {{ errorMessage }}
      </n-alert>
    </div>
    <div v-if="showStats" class="editor-stats">
      <n-space size="small">
        <n-tag size="small" type="info">
          行数: {{ lineCount }}
        </n-tag>
        <n-tag size="small" type="success" v-if="!hasError">
          ✓ 语法正确
        </n-tag>
        <n-tag size="small" type="error" v-if="hasError">
          ✗ 语法错误
        </n-tag>
      </n-space>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, onBeforeUnmount, watch, nextTick, computed } from 'vue';
  import { NAlert, NSpace, NTag, useMessage } from 'naive-ui';
  import { EditorView, basicSetup } from 'codemirror';
  import { EditorState } from '@codemirror/state';
  import { yaml } from '@codemirror/lang-yaml';
  import { oneDark } from '@codemirror/theme-one-dark';
  import * as yamlParser from 'js-yaml';

  export interface Props {
    value?: string;
    placeholder?: string;
    height?: string | number;
    disabled?: boolean;
    readonly?: boolean;
    darkTheme?: boolean;
    showStats?: boolean;
    validateOnChange?: boolean;
    validateOnBlur?: boolean;
    preventInvalidSubmit?: boolean;
  }

  const emit = defineEmits(['update:value', 'blur', 'focus', 'error', 'valid', 'change', 'validation-change']);
  const message = useMessage();

  const props = withDefaults(defineProps<Props>(), {
    value: '',
    placeholder: '# YAML 配置示例\n# 支持注释\nkey: value\nlist:\n  - item1\n  - item2\nconfig:\n  enabled: true\n  timeout: 30',
    height: 300,
    disabled: false,
    readonly: false,
    darkTheme: false,
    showStats: true,
    validateOnChange: false,
    validateOnBlur: true,
    preventInvalidSubmit: true,
  });

  const editorRef = ref<HTMLElement>();
  const editorView = ref<EditorView>();
  const errorMessage = ref('');
  const hasError = ref(false);
  const currentValue = ref(props.value || '');

  // 计算属性
  const lineCount = computed(() => {
    return currentValue.value.split('\n').length;
  });

  const isValid = computed(() => {
    return !hasError.value || !currentValue.value.trim();
  });

  // 监听外部value变化
  watch(
    () => props.value,
    (newValue) => {
      if (newValue !== currentValue.value && editorView.value) {
        currentValue.value = newValue || '';
        editorView.value.dispatch({
          changes: {
            from: 0,
            to: editorView.value.state.doc.length,
            insert: newValue || ''
          }
        });
        clearError();
      }
    },
    { immediate: false }
  );

  // 验证YAML格式
  function validateYaml(content: string): boolean {
    if (!content.trim()) {
      clearError();
      return true;
    }

    try {
      yamlParser.load(content);
      clearError();
      return true;
    } catch (error: any) {
      const errorMsg = `YAML语法错误: ${error.message}`;
      errorMessage.value = errorMsg;
      hasError.value = true;
      emit('error', { message: errorMsg, error });
      return false;
    }
  }

  function clearError() {
    errorMessage.value = '';
    hasError.value = false;
  }

  // 内容变化处理
  function onContentChange(content: string) {
    currentValue.value = content;
    
    // 验证 YAML 格式
    const valid = validateYaml(content);
    
    // 总是更新外部值
    emit('update:value', content);
    emit('change', content);
    
    // 发出验证状态变化事件
    emit('validation-change', valid);
    
    if (props.validateOnChange && valid) {
      emit('valid', content);
    }
  }

  // 强制添加选中样式
  function addSelectionStyles() {
    const styleId = 'yaml-editor-selection-styles';
    let existingStyle = document.getElementById(styleId);
    if (existingStyle) {
      existingStyle.remove();
    }
    
    const style = document.createElement('style');
    style.id = styleId;
    style.textContent = `
      .yaml-editor .cm-selectionBackground {
        background-color: rgba(64, 152, 252, 0.4) !important;
      }
      .yaml-editor .cm-focused .cm-selectionBackground {
        background-color: rgba(64, 152, 252, 0.5) !important;
      }
      .yaml-editor .cm-selectionLayer .cm-selectionBackground {
        background-color: rgba(64, 152, 252, 0.4) !important;
      }
      .yaml-editor .cm-content ::selection {
        background-color: rgba(64, 152, 252, 0.4) !important;
      }
      .yaml-editor .cm-line::selection {
        background-color: rgba(64, 152, 252, 0.4) !important;
      }
    `;
    document.head.appendChild(style);
  }

  // 创建编辑器
  function createEditor() {
    if (!editorRef.value) return;
    
    // 强制添加选中样式
    addSelectionStyles();

    const extensions = [
      basicSetup,
      yaml(),
      EditorView.updateListener.of((update) => {
        if (update.docChanged) {
          const content = update.state.doc.toString();
          onContentChange(content);
        }
        
        if (update.focusChanged) {
          if (update.view.hasFocus) {
            emit('focus');
          } else {
            emit('blur');
            if (props.validateOnBlur) {
              const content = update.state.doc.toString();
              const isValid = validateYaml(content);
              if (isValid) {
                emit('valid', content);
              }
            }
          }
        }
      }),
      EditorView.theme({
        '&': {
          fontSize: '14px',
          border: '1px solid #d9d9d9',
          borderRadius: '6px',
        },
        '&.cm-focused': {
          outline: 'none',
          borderColor: '#4098fc',
          boxShadow: '0 0 0 2px rgba(64, 152, 252, 0.2)',
        },
        '.cm-content': {
          padding: '12px',
          minHeight: typeof props.height === 'number' ? `${props.height}px` : props.height,
          fontFamily: '"SF Mono", Monaco, Inconsolata, "Roboto Mono", Consolas, "Courier New", monospace',
          lineHeight: '1.6',
        },
        '.cm-focused .cm-cursor': {
          borderColor: '#4098fc',
        },
        '&.error-state': {
          borderColor: '#ff4757',
        },
        '&.error-state.cm-focused': {
          borderColor: '#ff4757',
          boxShadow: '0 0 0 2px rgba(255, 71, 87, 0.2)',
        }
      }),
      EditorState.readOnly.of(props.readonly),
    ];

    // 添加暗色主题支持
    if (props.darkTheme) {
      extensions.push(oneDark);
    }

    const initialState = EditorState.create({
      doc: currentValue.value || props.placeholder,
      extensions,
    });

    editorView.value = new EditorView({
      state: initialState,
      parent: editorRef.value,
    });

    // 如果是占位符，选中所有文本
    if (!currentValue.value && props.placeholder) {
      nextTick(() => {
        if (editorView.value) {
          editorView.value.dispatch({
            selection: { anchor: 0, head: editorView.value.state.doc.length }
          });
        }
      });
    }
  }

  // 提供外部调用的验证方法
  function validate(): boolean {
    return validateYaml(currentValue.value);
  }

  // 提供外部调用的获取解析后的数据方法
  function getParsedData(): any {
    try {
      return yamlParser.load(currentValue.value);
    } catch (error) {
      return null;
    }
  }

  // 提供外部调用的设置数据方法
  function setData(data: any) {
    try {
      const yamlString = yamlParser.dump(data, {
        indent: 2,
        lineWidth: -1,
        noRefs: true,
        sortKeys: false,
        quotingType: '"',
        forceQuotes: false,
      });
      
      if (editorView.value) {
        editorView.value.dispatch({
          changes: {
            from: 0,
            to: editorView.value.state.doc.length,
            insert: yamlString
          }
        });
      }
      currentValue.value = yamlString;
      emit('update:value', yamlString);
      clearError();
    } catch (error: any) {
      const errorMsg = `数据转换YAML失败: ${error.message}`;
      errorMessage.value = errorMsg;
      hasError.value = true;
      emit('error', { message: errorMsg, error });
    }
  }

  // 获取当前值
  function getValue(): string {
    return currentValue.value;
  }

  // 设置焦点
  function focus() {
    editorView.value?.focus();
  }

  // 插入文本
  function insertText(text: string) {
    if (editorView.value) {
      const selection = editorView.value.state.selection.main;
      editorView.value.dispatch({
        changes: {
          from: selection.from,
          to: selection.to,
          insert: text
        }
      });
    }
  }

  onMounted(() => {
    nextTick(() => {
      createEditor();
    });
  });

  onBeforeUnmount(() => {
    if (editorView.value) {
      editorView.value.destroy();
    }
    
    // 清理动态样式
    const existingStyle = document.getElementById('yaml-editor-selection-styles');
    if (existingStyle) {
      existingStyle.remove();
    }
  });

  // 暴露方法给父组件
  defineExpose({
    validate,
    getParsedData,
    setData,
    getValue,
    focus,
    insertText,
    clearError,
    hasError: () => hasError.value,
    isValid: () => isValid.value,
  });
</script>

<style lang="less" scoped>
  .yaml-editor-container {
    width: 100%;
    
    .yaml-editor {
      border-radius: 6px;
      transition: all 0.2s;
      
      &.error-state {
        border-color: #ff4757;
      }
      
      :deep(.cm-editor) {
        outline: none;
      }
      
      :deep(.cm-content) {
        background: #fafafa;
      }
      
      :deep(.cm-line) {
        line-height: 1.6;
      }
      
      // YAML 语法高亮样式增强
      :deep(.cm-comment) {
        color: #8e8e93;
        font-style: italic;
      }
      
      :deep(.cm-string) {
        color: #0c7d9d;
      }
      
      :deep(.cm-number) {
        color: #1c00cf;
      }
      
      :deep(.cm-keyword) {
        color: #ad3da4;
        font-weight: 600;
      }
      
      :deep(.cm-property) {
        color: #3f6ec7;
        font-weight: 500;
      }
    }
    
    .error-message {
      margin-top: 8px;
    }
    
    .editor-stats {
      margin-top: 8px;
      display: flex;
      justify-content: flex-end;
    }
  }

  // 暗色主题支持
  .dark {
    .yaml-editor-container {
      .yaml-editor {
        :deep(.cm-content) {
          background: #1f1f1f;
          color: #d4d4d4;
        }
        
        :deep(.cm-selectionBackground) {
          background-color: rgba(64, 152, 252, 0.6) !important;
        }
        
        :deep(.cm-focused .cm-selectionBackground) {
          background-color: rgba(64, 152, 252, 0.7) !important;
        }
      }
    }
  }
</style>