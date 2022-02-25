<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="操作模块">
                <a-select placeholder="操作模块" v-model="queryParam.module" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in moduleOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="操作人员">
                <a-input v-model="queryParam.member_id" placeholder="请输入操作人员ID" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="访问路径">
                <a-input v-model="queryParam.url" placeholder="请输入访问路径" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="访问IP">
                <a-input v-model="queryParam.ip" placeholder="请输入IP地址" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="请求方式">
                <a-select placeholder="请求方式" v-model="queryParam.method" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in methodOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="耗时">
                <a-select placeholder="耗时" v-model="queryParam.take_up_time" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in takeTimeOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="状态码">
                <a-select placeholder="响应状态码" v-model="queryParam.error_code" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in errorCodeOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="操作时间">
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
      <view-form ref="viewForm" />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="系统参数"
        :loading="loading"
        rowKey="id"
        tableKey="monitor-operlog-Operlog-table"
        :isTableConfig="false"
        :isShowSetBtn="false"
        @refresh="getList"
        size="middle"
        :format-conditions="true"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        :pagination="{
          current: queryParam.page,
          limit: queryParam.limit,
          total: total,
          showSizeChanger: true,
          showLessItems: true,
          showQuickJumper: true,
          showTotal: (total, range) => `第 ${range[0]}-${range[1]} 条，总计 ${total} 条`,
          onChange: changeSize,
          onShowSizeChange: onSizeChange
        }"
      >
        <div class="table-operations" slot="button">
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['monitor:operlog:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button type="danger" @click="handleClean" v-hasPermi="['monitor:operlog:remove']">
            <a-icon type="delete" />清空
          </a-button>
          <a-button @click="handleExport" v-hasPermi="['system:config:export']">
            <a-icon type="download" />导出
          </a-button>
        </div>
        <span slot="method" slot-scope="{text, record}">
          <a-tag :color="text | operTypeFilter">
            {{ typeFormat(record) }}
          </a-tag>
        </span>
        <span slot="module" slot-scope="{text, record}">
          <a-tag :color="text | operModuleFilter">
            {{ moduleFormat(record) }}
          </a-tag>
        </span>
        <span slot="method" slot-scope="{text, record}">
          <a-tag :color="text | operMethodFilter">
            {{ methodFormat(record) }}
          </a-tag>
        </span>
        <span slot="error_code" slot-scope="{text, record}">
          <a-tag :color="text | operErrorCodeFilter">
            {{ errorCodeFormat(record) }}
          </a-tag>
        </span>
        <span slot="take_up_time" slot-scope="{text, record}">
          <a-tag :color="text | operTakeTimeFilter">
            {{ takeTimeFormat(record) }}
          </a-tag>
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="$refs.viewForm.handleView(record)" v-hasPermi="['monitor:operlog:query']">
            <a-icon type="eye" theme="twoTone"/>详细
          </a>
        </span>
        <span slot="member_name" slot-scope="{text,record}">
          <span>{{ memberNameFormat(record) }}</span>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { list, delOperlog, cleanOperlog } from '@/api/monitor/operlog'
import ViewForm from './modules/ViewForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import { exportDownload } from '@/utils/aidex'
export default {
  name: 'Operlog',
  components: {
    ViewForm,
    AdvanceTable
  },
  data () {
    return {
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
      loading: false,
      sunloading: false,
      total: 0,
      // 状态数据字典
      errorCodeOptions: [],
      takeTimeOptions: [],
      moduleOptions: [],
      methodOptions: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      // 日期范围
      dateRange: [],
      queryParam: {
        page: 1,
        limit: 10,
        module: undefined,
        member_id: undefined,
        method: undefined,
        url: undefined,
        error_code: undefined,
        take_up_time: undefined,
        ip: undefined
      },
      addModalRefName: 'addModal', // 添加弹窗ref名称
      columns: [
        {
          title: '应用端口',
          dataIndex: 'app_id',
          align: 'center'
        },
        {
          title: '操作人员',
          dataIndex: 'member_name',
          scopedSlots: { customRender: 'member_name' },
          align: 'center'
        },
        {
          title: '操作模块',
          dataIndex: 'module',
          scopedSlots: { customRender: 'module' },
          align: 'center'
        },
        {
          title: '请求方式',
          dataIndex: 'method',
          align: 'center'
        },
        {
          title: '请求路径',
          dataIndex: 'url',
          align: 'center'
        },
        {
          title: '访问IP',
          dataIndex: 'ip',
          align: 'center'
        },
        {
          title: 'IP地区',
          dataIndex: 'region',
          align: 'center'
        },
        {
          title: '状态码',
          dataIndex: 'error_code',
          scopedSlots: { customRender: 'error_code' },
          align: 'center'
        },
        {
          title: '耗时',
          dataIndex: 'take_up_time',
          scopedSlots: { customRender: 'take_up_time' },
          align: 'center'
        },
        {
          title: '操作日期',
          dataIndex: 'created_at',
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
    operTypeFilter (type) {
      let value = 'blue'
      if (type === 3 || type === 7 || type === 9) {
        value = 'orange'
      }
      return value
    },
    operModuleFilter (type) {
      let value = 'blue'
      if (type === 'api') {
        value = 'orange'
      }
      return value
    },
    operMethodFilter (method) {
      let value = 'blue'
      if (method === 'POST') {
        value = 'orange'
      }
      return value
    },
    operErrorCodeFilter (code) {
      let value = 'orange'
      if (code !== 0) {
        value = 'red'
      }
      return value
    },
    operTakeTimeFilter () {
      return 'magenta'
    }
  },
  created () {
    this.getList()
    this.getDicts('req_code').then(response => {
      this.errorCodeOptions = response
    })
    this.getDicts('req_take_up_time').then(response => {
      this.takeTimeOptions = response
    })
    this.getDicts('sys_oper_module').then(response => {
      this.moduleOptions = response
    })
    this.getDicts('sys_oper_method').then(response => {
      this.methodOptions = response
    })
  },
  computed: {
  },
  watch: {
  },
  methods: {
    onSizeChange (current, size) {
      this.queryParam.page = 1
      this.queryParam.limit = size
      this.getList()
    },
    /** 查询定时任务列表 */
    getList () {
      this.loading = true
      list(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.list
          this.total = response.total_count
          this.loading = false
        }
      )
    },
    // 操作日志状态字典翻译
    errorCodeFormat (row, column) {
      // return row.error_code
      return this.selectDictLabel(this.errorCodeOptions, row.error_code)
    },
    // 操作日志类型字典翻译
    takeTimeFormat (row, column) {
      return row.take_up_time + 'ms'
    },
    moduleFormat (row, column) {
      return this.selectDictLabel(this.moduleOptions, row.module)
    },
    methodFormat (row, column) {
      return this.selectDictLabel(this.methodOptions, row.method)
    },
    memberNameFormat(row) {
      if (row.member_id <= 0) {
        return ''
      }
      return '(#' + row.member_id + ')' + row.member_name
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.page = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.dateRange = []
      this.queryParam = {
        page: 1,
        limit: 10,
        module: undefined,
        member_id: undefined,
        method: undefined,
        url: undefined,
        error_code: undefined,
        take_up_time: undefined,
        ip: undefined
      }
      this.handleQuery()
    },
    onShowSizeChange (current, limit) {
      this.queryParam.limit = limit
      this.getList()
    },
    changeSize (current, limit) {
      this.queryParam.page = current
      this.queryParam.limit = limit
      this.getList()
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
      this.ids = this.selectedRows.map(item => item.id)
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const operIds = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中编号为' + operIds + '的数据',
        onOk () {
          return delOperlog(operIds)
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
        content: '此操作将会清空所有操作日志数据项',
        onOk () {
          return cleanOperlog()
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
          return exportDownload('/log/export', that.queryParam)
          // return exportOperlog(that.queryParam)
          //   .then(response => {
          //     that.download(response.msg)
          //     that.$message.success(
          //       '导出成功',
          //       3
          //     )
          //   })
        },
        onCancel () {}
      })
    }
  }
}
</script>
