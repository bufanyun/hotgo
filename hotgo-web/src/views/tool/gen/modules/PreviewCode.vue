<template>
  <div>
    <a-drawer
      title="代码预览"
      :width="800"
      :visible="visible"
      :confirmLoading="loading"
      @close="close"
    >
      <a-tabs v-if="isShowCodeTabs">
        <a-tab-pane
          v-for="(value, key) in previewData"
          :tab="key.substring(key.lastIndexOf('/')+1,key.indexOf('.vm'))"
          :key="key"
        >
          <div id="codeView" v-highlight>
            <pre><code v-text="value"></code></pre>
          </div>
        </a-tab-pane>
      </a-tabs>
      <div class="bottom-control">
        <a-space>
          <a-button type="dashed" @click="close">
            关闭
          </a-button>
        </a-space>
      </div>
    </a-drawer>
  </div>
</template>
<script>
import { previewTable } from '@/api/tool/gen'
export default {
  data () {
    return {
      isShowCodeTabs: false,
      previewData: {},
      visible: false,
      loading: false,
      // 模态框数据
      data: {},
      labelCol: {
        xs: { span: 12, push: 1 },
        sm: { span: 6 }
      },
      wrapperCol: {
        xs: { span: 24, push: 1 },
        sm: { span: 18 }
      }
    }
  },
  created () {},
  methods: {
    // 关闭模态框
    close () {
      this.visible = false
    },
    // 打开抽屉(由外面的组件调用)
    handlePreview (data) {
      if (data) {
        const tableId = data.tableId
        previewTable(tableId).then(response => {
            this.previewData = response.data
            this.isShowCodeTabs = true
        })
      }
      this.visible = true
    }
  }
}
</script>
