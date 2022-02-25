<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="参数编码">
                <a-input v-model="queryParam.configKey" placeholder="请输入参数编码" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="参数名称">
                <a-input v-model="queryParam.configName" placeholder="请输入参数名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="是否内置参数">
                <a-select placeholder="请选择" v-model="queryParam.configType" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in typeOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="创建时间">
                <a-range-picker style="width: 100%" v-model="dateRange" valueFormat="YYYY-MM-DD" format="YYYY-MM-DD" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col>
              <span class="table-page-search-submitButtons" style="float: right;">
                <a-button type="primary" @click="handleQuery"><a-icon type="search" />查询</a-button>
                <a-button style="margin-left: 8px" @click="resetQuery"><a-icon type="redo" />重置</a-button>
                <a @click="toggleAdvanced" style="margin-left: 8px">
                  {{ advanced ? '收起' : '展开' }}
                  <a-icon :type="advanced ? 'up' : 'down'"/>
                </a>
              </span>
            </a-col>

          </a-row>
        </a-form>
      </div>
    </a-card>
    <a-card :bordered="false" class="table-card">
      <!-- 增加组件按需加载 -->
      <config-add-form
        v-if="showAddModal"
        ref="configAddForm"
        :typeOptions="typeOptions"
        @ok="getList"
        @close="showAddModal = false"
      />
      <!-- 修改组件按需加载 -->
      <config-edit-form
        v-if="showEditModal"
        ref="configEditForm"
        :typeOptions="typeOptions"
        @ok="getList"
        @close="showEditModal = false"
      />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="系统参数"
        :loading="loading"
        rowKey="id"
        tableKey="system-config-ConfigIndex-table"
        @refresh="getList"
        size="middle"
        :format-conditions="true"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        :pagination="{
          current: queryParam.pageNum,
          pageSize: queryParam.pageSize,
          total: total,
          showSizeChanger: true,
          showLessItems: true,
          showQuickJumper: true,
          showTotal: (total, range) => `第 ${range[0]}-${range[1]} 条，总计 ${total} 条`,
          onChange: changeSize,
          onShowSizeChange: onSizeChange,
        }"
      >
        <div class="table-operations" slot="button">
          <a-button type="primary" size="small" @click="handleAdd" v-hasPermi="['system:config:add']">
            <a-icon type="plus" />新增
          </a-button>
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:config:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button @click="handleExport" v-hasPermi="['system:config:export']">
            <a-icon type="download" />导出
          </a-button>
          <a-button @click="handleRefreshCache" v-hasPermi="['system:config:remove']">
            <a-icon type="redo" />刷新缓存
          </a-button>
        </div>
        <span slot="configType" slot-scope="{text, record}">
          {{ typeFormat(record) }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleUpdate(record)" v-hasPermi="['system:config:edit']">
            修改
          </a>
          <a-divider type="vertical" v-hasPermi="['system:config:remove']" />
          <a @click="handleDelete(record)" v-hasPermi="['system:config:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { listConfig, delConfig, exportConfig, refreshCache } from '@/api/system/config'
import ConfigEditForm from './modules/ConfigEditForm'
import ConfigAddForm from './modules/ConfigAddForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
export default {
  name: 'Config',
  components: {
    ConfigEditForm,
    ConfigAddForm,
    AdvanceTable
  },
  data () {
    return {
      showAddModal: false,
      showEditModal: false,
      list: [],
      selectedRowKeys: [],
      selectedRows: [],
      // 高级搜索 展开/关闭
      advanced: false,
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      ids: [],
      configNames: [],
      loading: false,
      sunloading: false,
      total: 0,
      // 类型数据字典
      typeOptions: [],
      // 日期范围
      dateRange: [],
      labelCol: { span: 10 },
      wrapperCol: { span: 14 },
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        configName: undefined,
        configKey: undefined,
        configType: undefined
      },
      addModalRefName: 'addModal', // 添加弹窗ref名称
      columns: [
        {
          title: '参数编码',
          dataIndex: 'configKey',
          ellipsis: SVGComponentTransferFunctionElement,
          sswidth: '30%'
        },
        {
          title: '参数名称',
          dataIndex: 'configName',
          ellipsis: true,
          width: '15%'
        },
        {
          title: '参数值',
          dataIndex: 'configValue',
          ellipsis: true,
           width: '15%'
        },
        {
          title: '是否内置参数',
          dataIndex: 'configType',
          scopedSlots: { customRender: 'configType' },
          align: 'center',
           width: '10%'
        },
        {
          title: '备注',
          dataIndex: 'remark',
          ellipsis: true,
          width: '15%'
        },
        {
          title: '创建时间',
          dataIndex: 'createTime',
          scopedSlots: { customRender: 'createTime' },
          width: '15%'
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: '10%',
          align: 'center',
          scopedSlots: { customRender: 'operation' }
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
    this.getDicts('sys_yes_no').then(response => {
      this.typeOptions = response.data
    })
  },
  computed: {
  },
  watch: {
  },
  methods: {
    /** 查询定时任务列表 */
    getList () {
      this.showAddModal = false
      this.showEditModal = false
      this.loading = true
      listConfig(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.data.list
          this.list.map((item) => { item.operation = item.remark })
          this.total = response.data.total
          this.loading = false
        }
      )
    },
    typeFormat (row) {
      return this.selectDictLabel(this.typeOptions, row.configType)
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.pageNum = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.dateRange = []
      this.queryParam = {
        pageNum: 1,
        pageSize: 10,
        configName: undefined,
        configKey: undefined,
        configType: undefined
      }
      this.handleQuery()
    },
    onShowSizeChange (current, pageSize) {
      this.queryParam.pageSize = pageSize
      this.getList()
    },
    onSizeChange (current, size) {
      this.queryParam.pageNum = 1
      this.queryParam.pageSize = size
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
      this.ids = this.selectedRows.map(item => item.id)
      this.configNames = this.selectedRows.map(item => item.configName)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    handleAdd () {
      this.showAddModal = true
      this.$nextTick(() => (
         this.$refs.configAddForm.handleAdd()
       ))
    },
     handleUpdate (record) {
      this.showEditModal = true
      this.$nextTick(() => (
         this.$refs.configEditForm.handleUpdate(record)
       ))
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const configIds = row.id || this.ids
      const configNames = row.configName || this.configNames
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中名称为"' + configNames + '"的数据',
        onOk () {
          return delConfig(configIds)
            .then(() => {
              that.onSelectChange([], [])
              that.getList()
              that.$message.success(
                '删除成功',
                3
              )
            })
        },
        onCancel () {}
      })
    },
    /** 导出按钮操作 */
    handleExport () {
      var that = this
      this.$confirm({
        title: '是否确认导出?',
        content: '此操作将导出当前条件下所有数据而非选中数据',
        onOk () {
          return exportConfig(that.queryParam)
            .then(response => {
              that.download(response.msg)
              that.$message.success(
                '导出成功',
                3
              )
            })
        },
        onCancel () {}
      })
    },
    /** 清理缓存按钮操作 */
    handleRefreshCache () {
      refreshCache().then(response => {
        this.$message.success(
          '刷新成功',
          3
        )
      })
    }
  }
}
</script>
