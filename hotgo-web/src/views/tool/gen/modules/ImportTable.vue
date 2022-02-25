<template>
  <a-modal
    ref="createModal"
    :title="'导入代码'"
    :width="900"
    :visible="visible"
    @cancel="close"
    @ok="confirm"
    :confirmLoading="confirmLoading"
  >
    <div class="table-page-search-wrapper">
      <a-form layout="inline">
        <a-row :gutter="48">
          <a-col :md="8" :sm="24">
            <a-form-item label="表名称">
              <a-input v-model="queryParam.tableName" placeholder="请输入表名称" allow-clear @keyup.enter.native="handleQuery"/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <a-form-item label="表描述">
              <a-input v-model="queryParam.tableComment" placeholder="请输入表描述" allow-clear @keyup.enter.native="handleQuery"/>
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24">
            <span class="table-page-search-submitButtons">
              <a-button @click="handleQuery" type="primary">查询</a-button>
              <a-button @click="resetQuery" style="margin-left: 8px">重置</a-button>
            </span>
          </a-col>
        </a-row>
      </a-form>
    </div>
    <div class="page-header-content">
      <a-card :bordered="false" class="content">
        <a-table
          :loading="loading"
          rowKey="tableName"
          :columns="columns"
          :data-source="list"
          :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
          :scroll="{ y: tableHeight }"
          :pagination="false">
          <span slot="createTime" slot-scope="text, record">
            {{ parseTime(record.createTime) }}
          </span>
          <span slot="updateTime" slot-scope="text, record">
            {{ parseTime(record.updateTime) }}
          </span>
        </a-table>
        <!-- 分页 -->
        <a-pagination
          class="ant-table-pagination"
          show-size-changer
          show-quick-jumper
          :current="queryParam.pageNum"
          :total="total"
          :page-size="queryParam.pageSize"
          :showTotal="total => `共 ${total} 条`"
          @showSizeChange="onShowSizeChange"
          @change="changeSize"
        />
      </a-card>
    </div>
  </a-modal>
</template>
<script>
import { listDbTable, importTable } from '@/api/tool/gen'
export default {
  name: 'ImportTable',
  props: {
  },
  data () {
    return {
      // 表格数据
      list: [],
      selectedRowKeys: [],
      selectedRows: [],
        // 表格的高度
      tableHeight: document.documentElement.scrollHeight - 500 + 'px',
      // 选中表数组
      tableNames: [],
      loading: false,
      total: 0,
      // 当前控件配置:
      confirmLoading: false,
      visible: false,
      // 查询参数
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        tableName: undefined,
        tableComment: undefined
      },
      // 表格属性
      columns: [
        {
          title: '表名称',
          dataIndex: 'tableName',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '表描述',
          dataIndex: 'tableComment',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '创建时间',
          dataIndex: 'createTime',
          scopedSlots: { customRender: 'createTime' },
          align: 'center'
        },
        {
          title: '更新时间',
          dataIndex: 'updateTime',
          scopedSlots: { customRender: 'updateTime' },
          align: 'center'
        }
      ]
    }
  },
  created () {
    this.getList()
  },
  methods: {
    // 查询表数据
    getList () {
      this.loading = true
      listDbTable(this.queryParam).then(res => {
        if (res.code === 200) {
          this.list = res.rows
          this.total = res.total
          this.loading = false
        }
      })
    },
    // 关闭模态框
    close () {
      this.visible = false
      this.selectedRowKeys = []
      this.selectedRows = []
      this.queryParam = {
        pageNum: 1,
        pageSize: 10,
        tableName: undefined,
        tableComment: undefined
      }
    },
    // 打开抽屉(由外面的组件调用)
    show () {
      this.visible = true
      this.getList()
    },
    // 确认导入
    confirm () {
      this.confirmLoading = true
      importTable({ tables: this.tableNames.join(',') }).then(res => {
        this.$message.success(res.msg)
        this.confirmLoading = false
        this.visible = false
        this.$emit('ok')
      })
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.pageNum = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.queryParam = {
        pageNum: 1,
        pageSize: 10,
        tableName: undefined,
        tableComment: undefined
      }
      this.handleQuery()
    },
    onShowSizeChange (current, pageSize) {
      this.queryParam.pageSize = pageSize
      this.getList()
    },
    changeSize (current, pageSize) {
      this.queryParam.pageNum = current
      this.queryParam.pageSize = pageSize
      this.getList()
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
      this.tableNames = this.selectedRows.map(item => item.tableName)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    }
  }
}
</script>
