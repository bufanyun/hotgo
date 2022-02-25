<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="登录地址">
                <a-input v-model="queryParam.ipaddr" placeholder="请输入登录地址" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="操作信息">
                <a-input v-model="queryParam.msg" placeholder="请输入操作信息" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="使用状态">
                <a-select placeholder="请选择状态" v-model="queryParam.status" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="登录名称">
                <a-input v-model="queryParam.loginName" style="width: 100%" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="登陆时间">
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
      <advance-table
        :columns="columns"
        :data-source="list"
        title="登录日志"
        :loading="loading"
        rowKey="id"
        tableKey="monitor-logininfo-Logininfo-table"
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
          onShowSizeChange: onSizeChange,
        }"
      >
        <div class="table-operations" slot="button">
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['monitor:LoginLog:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button type="danger" @click="handleClean" v-hasPermi="['monitor:LoginLog:remove']">
            <a-icon type="delete" />清空
          </a-button>
          <a-button @click="handleExport" v-hasPermi="['system:LoginLog:export']">
            <a-icon type="download" />导出
          </a-button>
        </div>
        <span slot="status" slot-scope="{text, record}">
          <a-badge :status="record.status == '0' ? 'processing' : 'error'" :text=" statusFormat(record) " />
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { list, delLoginLog, cleanLoginLog, exportLoginLog } from '@/api/monitor/loginLog'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
export default {
  name: 'LoginLog',
  components: {
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
      statusOptions: [],
      // 日期范围
      dateRange: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        ipaddr: null,
        userName: undefined,
        status: undefined
      },
      addModalRefName: 'addModal', // 添加弹窗ref名称
      columns: [
        {
          title: '用户名称',
          dataIndex: 'userName',
          align: 'center'
        },
        {
          title: '登录地址',
          dataIndex: 'ipaddr',
          align: 'center'
        },
        {
          title: '登录地点',
          dataIndex: 'loginLocation',
          align: 'center'
        },
        {
          title: '浏览器',
          dataIndex: 'browser',
          align: 'center'
        },
        {
          title: '操作系统',
          dataIndex: 'os',
          align: 'center'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        },
        {
          title: '操作信息',
          dataIndex: 'msg',
          align: 'center'
        },
        {
          title: '登录时间',
          dataIndex: 'loginTime',
          align: 'center'
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
    this.getDicts('sys_common_status').then(response => {
      this.statusOptions = response.data
    })
  },
  computed: {
  },
  watch: {
  },
  methods: {
    onSizeChange (current, size) {
      this.queryParam.pageNum = 1
      this.queryParam.pageSize = size
      this.getList()
    },
    /** 查询定时任务列表 */
    getList () {
      this.loading = true
      list(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.loading = false
        }
      )
    },
    // 执行状态字典翻译
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
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
        ipaddr: null,
        userName: undefined,
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
      this.ids = this.selectedRows.map(item => item.id)
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const infoIds = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        // content: '当前选中访问编号为' + infoIds + '的数据',
        onOk () {
          return delLoginLog(infoIds)
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
        content: '此操作将会清空所有登录日志数据项',
        onOk () {
          return cleanLoginLog()
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
          return exportLoginLog(that.queryParam)
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
