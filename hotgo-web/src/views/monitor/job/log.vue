<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="任务名称">
                <a-input v-model="queryParam.jobName" placeholder="请输入任务名称" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="任务组名">
                <a-select placeholder="请选择任务组名" v-model="queryParam.jobGroup" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in jobGroupOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <template>
              <a-col :md="6" :sm="24">
                <a-form-item label="执行状态">
                  <a-select placeholder="请选择执行状态" v-model="queryParam.status" style="width: 100%" allow-clear>
                    <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-select-option>
                  </a-select>
                </a-form-item>
              </a-col>
              <a-col :md="6" :sm="24" v-if="advanced">
                <a-form-item
                  label="
                执行时间">
                  <a-range-picker style="width: 100%" v-model="dateRange" valueFormat="YYYY-MM-DD" format="YYYY-MM-DD" allow-clear/>
                </a-form-item>
              </a-col>
            </template>
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
      <!-- 详细信息 -->
      <log-view-form ref="logViewForm" :jobGroupOptions="jobGroupOptions" />

      <advance-table
        :columns="columns"
        :data-source="list"
        title="定时任务日志"
        :loading="loading"
        rowKey="jobLogId"
        tableKey="monitor-job_log-index-table"
        :isTableConfig="false"
        :isShowSetBtn="false"
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
          onShowSizeChange: onShowSizeChange,
        }"
      >
        <span slot="jobGroup" slot-scope="{text, record}">
          {{ jobGroupFormat(record) }}
        </span>
        <span slot="status" slot-scope="{text, record}">
          {{ statusFormat(record) }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="$refs.logViewForm.handleView(record)" v-hasPermi="['monitor:job:query']">
            <a-icon type="eye" theme="twoTone"/>详细
          </a>
        </span>
        <div class="table-operations" slot="button">
          <a-button type="danger" v-if="!multiple" @click="handleDelete" v-hasPermi="['monitor:job:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button type="danger" @click="handleClean" v-hasPermi="['monitor:job:remove']">
            <a-icon type="delete" />清空
          </a-button>
          <a-button type="primary" @click="handleExport" v-hasPermi="['monitor:job:export']">
            <a-icon type="download" />导出
          </a-button>
        </div>
      </advance-table>
    </a-card>
  </div>
</template>

<script>

import { listJobLog, delJobLog, exportJobLog, cleanJobLog } from '@/api/monitor/jobLog'
import LogViewForm from './modules/LogViewForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
export default {
  name: 'JobLog',
  components: {
    LogViewForm,
    AdvanceTable
  },
  data () {
    return {
      list: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      total: 0,
      // 状态数据字典
      statusOptions: [],
      jobGroupOptions: [],
      // 日期范围
      dateRange: [],
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        jobName: undefined,
        jobGroup: undefined,
        status: undefined
      },
      // 高级搜索 展开/关闭
      advanced: false,
      loading: false,
      // 选中数组
      ids: [],
      selectedRowKeys: [],
      selectedRows: [],
      // 非多个禁用
      multiple: true,
      columns: [
        {
          title: '日志编号',
          dataIndex: 'jobLogId',
          align: 'center'
        },
        {
          title: '系统模块',
          dataIndex: 'jobName',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '操作类型',
          dataIndex: 'jobGroup',
          scopedSlots: { customRender: 'jobGroup' },
          align: 'center'
        },
        {
          title: '调用目标字符串',
          dataIndex: 'invokeTarget',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '日志信息',
          dataIndex: 'jobMessage',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '执行状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        },
        {
          title: '执行时间',
          dataIndex: 'createTime',
          align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'operation',
          scopedSlots: { customRender: 'operation' },
          align: 'center'
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
    this.getDicts('sys_job_status').then(response => {
      this.statusOptions = response.data
    })
    this.getDicts('sys_job_group').then(response => {
      this.jobGroupOptions = response.data
    })
  },
  computed: {
  },
  watch: {
  },
  methods: {
    /** 查询登录日志列表 */
    getList () {
      this.loading = true
      listJobLog(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.rows
          this.total = response.total
          this.loading = false
        }
      )
    },
    // 执行状态字典翻译
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    // 任务组名字典翻译
    jobGroupFormat (row) {
      return this.selectDictLabel(this.jobGroupOptions, row.jobGroup)
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
        jobName: undefined,
        jobGroup: undefined,
        status: undefined
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
      this.ids = this.selectedRows.map(item => item.jobLogId)
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const jobLogIds = row.jobLogId || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中日志编号为' + jobLogIds + '的数据',
        onOk () {
          return delJobLog(jobLogIds)
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
    /** 清空按钮操作 */
    handleClean () {
      var that = this
      this.$confirm({
        title: '是否确认清空?',
        content: '此操作将会清空所有调度日志数据项',
        onOk () {
          return cleanJobLog()
            .then(() => {
              that.onSelectChange([], [])
              that.getList()
              that.$message.success(
                '清空成功',
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
          return exportJobLog(that.queryParam)
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
    }
  }
}
</script>
