<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="任务名称">
                <a-input v-model="queryParam.jobName" placeholder="请输入任务名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="任务组名">
                <a-input v-model="queryParam.jobGroup" placeholder="请选择任务组名" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item
                label="
              任务状态">
                <a-select placeholder="请选择任务状态" v-model="queryParam.status" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col>
              <span class="table-page-search-submitButtons" style="float: right;">
                <a-button type="primary" @click="handleQuery"><a-icon type="search" />查询</a-button>
                <a-button style="margin-left: 8px" @click="resetQuery"><a-icon type="redo" />重置</a-button>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>
    </a-card>
    <a-card :bordered="false" class="table-card">
      <!-- 详细信息 -->
      <view-form ref="viewForm" :jobGroupOptions="jobGroupOptions" />
      <!-- 增加修改 -->
      <create-form
        ref="createForm"
        :statusOptions="statusOptions"
        :jobGroupOptions="jobGroupOptions"
        @ok="getList"
      />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="定时任务"
        :loading="loading"
        rowKey="jobId"
        tableKey="monitor-job-index-table"
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
        <div class="table-operations" slot="button">
          <a-button type="primary" @click="$refs.createForm.handleAdd()" v-hasPermi="['monitor:job:add']">
            <a-icon type="plus" />新增
          </a-button>
          <!-- <a-button type="primary" v-if="!single" :disabled="single" @click="$refs.createForm.handleUpdate(undefined, ids)" v-hasPermi="['monitor:job:edit']">
            <a-icon type="edit" />修改
          </a-button> -->
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['monitor:job:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button type="primary" @click="handleExport" v-hasPermi="['monitor:job:export']">
            <a-icon type="download" />导出
          </a-button>
          <a-button type="dashed" @click="handleJobLog" v-hasPermi="['monitor:job:query']">
            <a-icon type="snippets" />日志
          </a-button>
        </div>
        <span slot="jobGroup" slot-scope="{text, record}">
          {{ jobGroupFormat(record) }}
        </span>
        <span slot="status" slot-scope="{text, record}">
          <a-popconfirm
            ok-text="是"
            cancel-text="否"
            @confirm="confirmHandleStatus(record)"
            @cancel="cancelHandleStatus(record)"
          >
            <span slot="title">确认<b>{{ record.status === '1' ? '开启' : '关闭' }}</b>{{ record.jobName }}的任务吗?</span>
            <a-switch checked-children="开" un-checked-children="关" :checked="record.status == 0" />
          </a-popconfirm>
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a-popconfirm
            ok-text="是"
            cancel-text="否"
            @confirm="confirmHandleRun(record)"
            @cancel="cancelHandleRun(record)"
          >
            <span slot="title">确认执行一次{{ record.jobName }}的任务吗?</span>
            <a v-hasPermi="['monitor:job:changeStatus']">
              <a-icon type="caret-right" />
              执行一次
            </a>
          </a-popconfirm>

          <a-divider type="vertical" />
          <a @click="$refs.createForm.handleUpdate(record)" v-hasPermi="['monitor:job:edit']">
            <a-icon type="eye" />修改
          </a>
          <a-divider type="vertical" />
          <a @click="$refs.viewForm.handleView(record)" v-hasPermi="['monitor:job:query']">
            <a-icon type="eye" />详细
          </a>
        </span>
      </advance-table>

    </a-card>
  </div>
</template>

<script>

import { listJob, delJob, exportJob, runJob, changeJobStatus } from '@/api/monitor/job'
import CreateForm from './modules/CreateForm'
import ViewForm from './modules/ViewForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
export default {
  name: 'Job',
  components: {
    CreateForm,
    ViewForm,
    AdvanceTable
  },
  data () {
    return {
      list: [],
      selectedRowKeys: [],
      selectedRows: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      // 高级搜索 展开/关闭
      advanced: false,
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      ids: [],
      loading: false,
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
      columns: [
        {
          title: '任务编号',
          dataIndex: 'jobId',
          width: '70px',
          align: 'center'
        },
        {
          title: '任务名称',
          dataIndex: 'jobName',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '任务组名',
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
          title: 'cron执行表达式',
          dataIndex: 'cronExpression',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        },
        {
          title: '备注',
          dataIndex: 'remark',
           ellipsis: true
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: '250px',
          scopedSlots: { customRender: 'operation' }
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
    /** 查询定时任务列表 */
    getList () {
      this.loading = true
      listJob(this.queryParam).then(response => {
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
      this.ids = this.selectedRows.map(item => item.jobId)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 任务日志列表查询 */
    handleJobLog () {
      this.$router.push({ path: '/log/joblog' })
    },
    /* 任务状态修改 */
    confirmHandleStatus (row) {
      const text = row.status === '1' ? '启用' : '停用'
      row.status = row.status === '0' ? '1' : '0'
      changeJobStatus(row.jobId, row.status)
      .then(() => {
        this.$message.success(
          text + '成功',
          3
        )
      }).catch(function () {
        this.$message.error(
          text + '发生异常',
          3
        )
      })
    },
    cancelHandleStatus (row) {
    },
    /* 立即执行一次 */
    confirmHandleRun (row) {
      runJob(row.jobId, row.jobGroup)
      .then(() => {
        this.$message.success(
          '执行成功',
          3
        )
      }).catch(function () {
        this.$message.error(
          '发生异常',
          3
        )
      })
    },
    cancelHandleRun (row) {
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      this.ids = this.selectedRows.map(item => item.jobId)
      const jobIds = row.jobId || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中定时任务编号为' + jobIds + '的数据',
        onOk () {
          return delJob(jobIds)
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
          return exportJob(that.queryParam)
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
