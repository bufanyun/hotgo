<template>
  <a-popconfirm
    v-if="askType == 'popconfirm'"
    :title="title"
    @confirm="handleConfirm"
    @cancel="handleCancel"
  >
    <slot>
      <a-button :type="type" icon="delete" :disabled="disabled" loading="loading" title="删除">
        删除
      </a-button>
    </slot>
  </a-popconfirm>
  <span v-else @click="handleDelete">
    <slot>
      <a-button :type="type" icon="delete" :loading="loading" :disabled="disabled" title="删除">
        删除
      </a-button>
    </slot>
  </span>
</template>
<script>
import { Modal, Popconfirm, Button } from 'ant-design-vue'
export default {
  name: 'AidexDelete',
  components: {
    APopconfirm: Popconfirm,
    AButton: Button
  },
  props: {
    askType: {
      // 询问类型:modal popconfirm,默认modal.
      type: String,
      default: 'modal'
    },
    title: {
      // 提示语
      type: String,
      default: '确认要删除选中的数据吗?'
    },
    loading: {
      // 按钮loading状态
      type: Boolean,
      default: false
    },
    type: {
      // 按钮类型
      type: String,
      default: 'danger'
    },
    disabled: {
      // 按钮是否禁用
      type: Boolean,
      default: false
    }
  },
  data () {
    return {}
  },
  methods: {
    handleDelete (e) {
      e.stopPropagation()
      // modal提示是否删除
      Modal.confirm({
        title: this.title,
        onOk: () => {
          // 确认删除
          this.handleConfirm()
        },
        onCancel: () => {
          // 取消删除
          this.handleCancel()
        }
      })
      // 确认删除
      this.$emit('click')
    },
    handleConfirm () {
      // 确认删除
      this.$emit('click-confirm')
    },
    handleCancel () {
      // 取消删除
      this.$emit('click-cancel')
    }
  }
}
</script>
