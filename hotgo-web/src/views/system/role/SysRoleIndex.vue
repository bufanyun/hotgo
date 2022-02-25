<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper" ref="search">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="角色名">
                <a-input v-model="queryParam.name" placeholder="请输入" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="角色编码">
                <a-input v-model="queryParam.key" placeholder="请输入" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="状态">
                <a-select placeholder="请选择" v-model="queryParam.status" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="状态">
                <a-select placeholder="请选择" v-model="queryParam.status" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="状态">
                <a-select placeholder="请选择" v-model="queryParam.status" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="创建时间">
                <a-range-picker style="width: 100%" v-model="dateRange" valueFormat="YYYY-MM-DD" format="YYYY-MM-DD" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="12" :sm="24" v-if="advanced">
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
      <!-- 增加修改 -->
      <sys-role-add-form
        v-if="showAddModal"
        ref="sysRoleAddForm"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showAddModal = false"
      />
      <sys-role-edit-form
        v-if="showEditModal"
        ref="sysRoleEditForm"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showEditModal = false"
      />
      <!-- 分配角色数据权限对话框 -->
      <create-data-scope-form
        v-if="showDataScopeModal"
        ref="createDataScopeForm"
        @ok="getList"
        @close="showDataScopeModal = false"
      />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="角色管理"
        :loading="loading"
        rowKey="id"
        size="middle"
        tableKey="system-role-SysRoleIndex-table"
        @refresh="getList"
        :format-conditions="true"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
        :pagination="{
          current: queryParam.page,
          limit: queryParam.limit,
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
          <a-button type="primary" size="small" @click="handleAdd()" v-hasPermi="['system:role:add']">
            <a-icon type="plus" />新增
          </a-button>
          <a-button type="danger" v-if="!multiple" @click="handleDelete" v-hasPermi="['system:role:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button type="" @click="handleExport" v-hasPermi="['system:role:export']">
            <a-icon type="download" />导出
          </a-button>
        </div>

        <span slot="status" slot-scope="{record}">
          <a-badge :status="record.status === '1' ? 'processing' : 'error'" :text=" statusFormat(record) " />
        </span>
        <span slot="key" slot-scope="{text}">
          <a-tag color="blue">
            {{ text }}
          </a-tag>
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click.stop="handleUpdate(record)" v-hasPermi="['system:role:edit']">
            修改
          </a>
          <a-divider type="vertical" v-hasPermi="['system:role:edit']" />
          <a @click.stop="handleDataScope(record)" v-hasPermi="['system:role:edit']">
            数据权限
          </a>
          <a-divider type="vertical" v-hasPermi="['system:role:remove']" />
          <a @click.stop="handleDelete(record)" v-hasPermi="['system:role:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { listRole, delRole, exportRole } from '@/api/system/role'
import SysRoleAddForm from './modules/SysRoleAddForm'
import SysRoleEditForm from './modules/SysRoleEditForm'
import CreateDataScopeForm from './modules/CreateDataScopeForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
export default {
  name: 'Role',
  components: {
    SysRoleAddForm,
    SysRoleEditForm,
    CreateDataScopeForm,
    AdvanceTable
  },
  data () {
    return {
      showAddModal: false,
      showEditModal: false,
      showDataScopeModal: false,
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
      total: 0,
      // 状态数据字典
      statusOptions: [],
      // 日期范围
      dateRange: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      queryParam: {
        page: 1,
        limit: 10,
        name: undefined,
        key: undefined,
        status: undefined
      },
       addModalRefName: 'addModal', // 添加弹窗ref名称
      columns: [
        {
          title: '角色名',
          dataIndex: 'name',
          ellipsis: true
        },
        {
          title: '角色编码',
          dataIndex: 'key',
          ellipsis: true,
          scopedSlots: { customRender: 'key' }
        },
        {
          title: '排序',
          dataIndex: 'sort',
          align: 'center'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        },
        {
          title: '创建时间',
          dataIndex: 'created_at',
          scopedSlots: { customRender: 'created_at' },
           align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: '20%',
          scopedSlots: { customRender: 'operation' }
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response
    })
  },
  computed: {
  },
  watch: {
  },
  methods: {
    statusFormat (row) {
      return this.selectlabel(this.statusOptions, row.status)
    },
    /** 查询定时任务列表 */
    getList () {
      this.showAddModal = false
      this.showEditModal = false
      this.showDataScopeModal = false
      this.loading = true
      listRole(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.list
          this.list.map((item) => { item.operation = item.remark })
          this.total = response.total_count
          this.loading = false
        }
      )
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.page = 1
      this.getList()
    },
    handleAdd (record) {
      this.showAddModal = true
      this.$nextTick(() => (
        this.$refs.sysRoleAddForm.handleAdd(record)
      ))
    },
    handleUpdate (record) {
      this.showEditModal = true
      this.$nextTick(() => (
        this.$refs.sysRoleEditForm.handleUpdate(record)
      ))
    },
    handleDataScope (record) {
      this.showDataScopeModal = true
      this.$nextTick(() => (
        this.$refs.createDataScopeForm.handleDataScope(record)
      ))
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.dateRange = []
      this.queryParam = {
        page: 1,
        limit: 10,
        name: undefined,
        key: undefined,
        status: undefined
      }
      this.handleQuery()
    },
    onShowSizeChange (current, limit) {
      this.queryParam.limit = limit
      this.getList()
    },
    onSizeChange (current, size) {
      this.queryParam.page = 1
      this.queryParam.limit = size
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
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const roleIds = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        // content: '当前选中编号为' + roleIds + '的数据',
        onOk () {
          return delRole(roleIds)
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
          return exportRole(that.queryParam)
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
