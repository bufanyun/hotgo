<template>
  <div ref="content">
    <a-card :bordered="false">
      <!-- 编辑用户,单独封装了组件 -->
      <sys-user-edit-form
        v-if="showEditModal"
        ref="sysUserEditForm"
        :statusOptions="statusOptions"
        :sexOptions="sexOptions"
        :userTypeOptions="userTypeOptions"
        :defalutExpandedKeys="expandedKeys"
        @ok="getList"
        @close="showEditModal = false"
      />
      <a-page-header
        style="border-bottom: 1px solid rgb(235, 237, 240);padding:8px 16px;"
        title=""/>
      <split splitHeight="100%" leftWidth="220" >
        <template slot="paneL">
          <!-- 部门树 -->
          <dept-tree
            ref="deptTree"
            :deptOptions="deptOptions"
            :defalutExpandedKeys="expandedKeys"
            @setDataOptionInfo="setDataOptionInfo"
            @select="clickDeptNode"
          />
        </template>
        <template slot="paneR">
          <!-- 条件搜索 -->
          <div class="table-page-search-wrapper">
            <a-form
              :labelCol="labelCol"
              :wrapperCol="wrapperCol"
            >
              <a-row :gutter="48">
                <a-col :md="8" :sm="24">
                  <a-form-item label="姓名">
                    <a-input v-model="queryParam.realname" placeholder="请输入姓名" allow-clear @keyup.enter.native="handleQuery"/>
                  </a-form-item>
                </a-col>
                <a-col :md="8" :sm="24">
                  <a-form-item label="登录名">
                    <a-input v-model="queryParam.username" placeholder="请输入登录名" allow-clear @keyup.enter.native="handleQuery"/>
                  </a-form-item>
                </a-col>
                <a-col :md="8" :sm="24" v-if="advanced">
                  <a-form-item label="手机号">
                    <a-input v-model="queryParam.mobile" placeholder="请输入手机号" allow-clear @keyup.enter.native="handleQuery"/>
                  </a-form-item>
                </a-col>
                <a-col :md="8" :sm="24" v-if="advanced">
                  <a-form-item label="状态">
                    <a-select placeholder="请选择状态" v-model="queryParam.status" style="width: 100%" allow-clear>
                      <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                    </a-select>
                  </a-form-item>
                </a-col>
                <a-col :md="8" :sm="24" v-if="advanced">
                  <a-form-item label="创建时间">
                    <a-range-picker style="width: 100%" v-model="dateRange" valueFormat="YYYY-MM-DD" format="YYYY-MM-DD" allow-clear />
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
          <a-divider />
          <!-- 增加修改 -->
          <sys-user-add-form
            v-if="showAddModal"
            ref="sysUserAddForm"
            :statusOptions="statusOptions"
            :sexOptions="sexOptions"
            :userTypeOptions="userTypeOptions"
            :defalutExpandedKeys="expandedKeys"
            :deptCheckedValue="deptCheckedValue"
            @ok="getList"
            @close="showAddModal = false"
          />
          <!-- 修改密码抽屉 -->
          <reset-password
            v-if="showResetPassword"
            ref="resetPassword"
            @close="showResetPassword = false"
          />
          <!-- 上传文件 -->
          <import-excel
            title="导入用户"
            upload-msg="是否更新已经存在的用户数据"
            upload-action="/system/user/importData"
            ref="importExcel"
            @ok="getList"
            @importTemplate="importTemplate"
          />
          <advance-table
            :columns="columns"
            :data-source="list"
            :title="tableTitle"
            :loading="loading"
            rowKey="id"
            size="middle"
            tableKey="system-user-SysUserIndex-table"
            @change="handleTableChange"
            @refresh="getList"
            :format-conditions="true"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
            :pagination="{
              current: queryParam.page,
              limit: queryParam.limit,
              total_count: total_count,
              showSizeChanger: true,
              showLessItems: true,
              showQuickJumper: true,
              showTotal: (total_count, range) => `第 ${range[0]}-${range[1]} 条，总计 ${total_count} 条`,
              onChange: changeSize,
              onShowSizeChange: onShowSizeChange,
            }"
          >
            <div class="table-operations" slot="button">
              <a-button type="primary" @click="handleAdd()" v-hasPermi="['system:user:add']">
                <a-icon type="plus" />新增
              </a-button>
              <a-button type="danger" v-if="!multiple" @click="handleDelete" v-hasPermi="['system:user:remove']">
                <a-icon type="delete" />删除
              </a-button>
            </div>
            <span slot="status" slot-scope="{record}">
              <a-badge :status="record.status === '1' ? 'processing' : 'error'" :text=" statusFormat(record) " />
            </span>
            <span slot="userType" slot-scope="{record}">
              {{ userTypeFormat(record) }}
            </span>

            <span slot="realname" slot-scope="{text,record,index}">
              <a @click="onClickRow(record,index)">
                <a-avatar :style="{backgroundColor:getRandomColor(index)}" style="color: #fff; ">
                  {{ getShowName(text) }}
                </a-avatar>
                {{ text }}
              </a>
            </span>
            <span slot="created_at" slot-scope="text, record">
              {{ parseTime(record.created_at) }}
            </span>
            <span slot="operation" slot-scope="{text, record}">
              <a @click="handleUpdate(record,undefined)" v-hasPermi="['system:user:edit']">
                修改
              </a>
              <a-divider type="vertical" v-if="record.id !== 1" v-hasPermi="['system:user:remove']" />
              <a @click="handleDelete(record)" v-if="record.id !== 1" v-hasPermi="['system:user:remove']">
                删除
              </a>
              <a-divider type="vertical" v-hasPermi="['system:user:resetPwd']" />
              <a @click="handleResetPwd(record)" v-hasPermi="['system:user:resetPwd']">
                重置密码
              </a>
            </span>
          </advance-table>
        </template>
      </split>
    </a-card>
  </div>
</template>

<script>
import { listUser, delUser, exportUser, importTemplate } from '@/api/system/user'
import { listDeptTree } from '@/api/system/dept'
import ResetPassword from './modules/ResetPassword'
import SysUserAddForm from './modules/SysUserAddForm'
import SysUserEditForm from './modules/SysUserEditForm'
import ImportExcel from '@/components/pt/import/ImportExcel'
import DeptTree from './modules/DeptTree'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import Split from '@/components/pt/split/Index'
export default {
  name: 'User',
  components: {
    ResetPassword,
    SysUserEditForm,
    ImportExcel,
    AdvanceTable,
    SysUserAddForm,
    Split,
    DeptTree
  },
  data () {
    return {
      showAddModal: false,
      showEditModal: false,
      showResetPassword: false,
      list: [],
      colorList: ['#F38709', '#813AFD', '#00C4AA', '#4B7AEE'],
      tableTitle: '用户管理',
      selectedRowKeys: [],
      selectedRows: [],
      // 高级搜索 展开/关闭
      advanced: false,
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      ids: [],
      userNames: [],
      expandedKeys: [],
      loading: false,
      total_count: 0,
      // 状态数据字典
      statusOptions: [],
      sexOptions: [],
      userTypeOptions: [],
      deptCheckedValue: {},
      // 部门树选项
      deptOptions: [],
      // 日期范围
      dateRange: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      queryParam: {
        page: 1,
        limit: 10,
        username: undefined,
        mobile: undefined,
        status: undefined,
        dept_id: undefined
      },
      columns: [
        {
          title: '姓名',
          dataIndex: 'realname',
          width: 120,
          scopedSlots: { customRender: 'realname' }
        },
        {
          title: '登录名',
          dataIndex: 'username'
        },
        {
          title: '用户状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
           align: 'center'
        },
        {
          title: '用户角色',
          dataIndex: 'role_name',
          scopedSlots: { customRender: 'role_name' },
           align: 'center'
        },
        {
          title: '部门',
          dataIndex: 'dept_name'
        },
        {
          title: '创建时间',
          ellipsis: true,
          sorter: true,
          width: 200,
          dataIndex: 'created_at',
          align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: 200,
          scopedSlots: { customRender: 'operation' }
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
    this.getTreeselect()
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response
    })
    this.getDicts('sys_user_sex').then(response => {
      this.sexOptions = response
    })
    this.getDicts('sys_user_type').then(response => {
      this.userTypeOptions = response
    })
  },
  computed: {
  },
  watch: {
  },
  methods: {
    importTemplate (expanded, record) {
      importTemplate().then(response => {
        this.download(response.msg)
      })
    },
    /** 查询部门下拉树结构 */
    getTreeselect () {
       listDeptTree('0', 3).then(response => {
        this.deptOptions = response
        this.getExpandedKeys(this.deptOptions, 3)
      })
    },
    getExpandedKeys (nodes, expandLevel) {
      if (expandLevel > 1) {
        // 最后一层不展开
        nodes.forEach(node => {
        this.expandedKeys.push(node.id)
        expandLevel = expandLevel - 1
        return this.getExpandedKeys(node.children, expandLevel)
        })
      }
    },
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    userTypeFormat (row) {
      return this.selectDictLabel(this.userTypeOptions, row.userType)
    },
    /** 查询定时任务列表 */
    getList () {
      this.showAddModal = false
      this.showEditModal = false
      this.loading = true
      listUser(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.list
          this.total_count = response.total_count
          this.loading = false
        }
      )
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
        username: undefined,
        mobile: undefined,
        status: undefined,
        dept_id: undefined
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
      this.userNames = this.selectedRows.map(item => item.realname)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    clickDeptNode (node) {
      // console.log('node:' + JSON.stringify(node.$options.propsData))
      // this.queryParam.deptId = node.$options.propsData.eventKey
      // this.deptCheckedValue.ids = node.$options.propsData.eventKey
      // this.tableTitle = node.$options.propsData.dataRef.name
      // this.deptCheckedValue = { ids: node.$options.propsData.eventKey, names: node.$options.propsData.dataRef.name }
      this.queryParam.dept_id = node.$options.propsData.eventKey
      this.deptCheckedValue.ids = node.$options.propsData.eventKey
      this.tableTitle = node.$options.propsData.label
      if (node.$options.propsData.type === 'dept') {
      // if (node.$options.propsData.dataRef.attributes.deptType === 'dept') {
        this.deptCheckedValue = { ids: node.$options.propsData.eventKey, names: node.$options.propsData.label }
      } else {
        this.deptCheckedValue = {}
      }
      this.getList()
    },
    handleAdd () {
      this.showAddModal = true
      this.$nextTick(() => (
         this.$refs.sysUserAddForm.handleAdd()
       ))
    },
     handleUpdate (record, ids) {
      this.showEditModal = true
      this.$nextTick(() => (
         this.$refs.sysUserEditForm.handleUpdate(record, ids)
       ))
    },
    handleResetPwd (record) {
      this.showResetPassword = true
      this.$nextTick(() => (
         this.$refs.resetPassword.handleResetPwd(record)
       ))
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const userIds = row.id || this.ids
      const userNames = row.realname || this.userNames
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中编号为' + userNames + '的数据',
        onOk () {
          return delUser(userIds)
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
          return exportUser(that.queryParam)
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
    getRandomColor (index) {
         // const randomValue = Math.round(Math.random() * 3)
         return this.colorList[index % 4]
    },
    getShowName (name) {
       if (name.length > 2) {
         name = name.substring(name.length - 2)
       }
       return name
    },
    onClickRow (record, index) {
         this.handleUpdate(record, '')
    },
    setDataOptionInfo (treeDataOption) {
       this.deptOptions = treeDataOption
    },
    handleTableChange (pagination, filters, sorter) {
      if (sorter.field !== undefined && sorter.field !== null && sorter.field !== '') {
        this.queryParam.orderByColumn = 't.' + sorter.field
        this.queryParam.isAsc = sorter.order
      }
      this.getList()
    }
  }
}
</script>
<style lang="less" scoped>
  .ant-divider-horizontal {
      margin:  0;
      background: rgb(235, 237, 240);
  }
  .demo-split{
      height: 200px;
  }
  .demo-split-pane{
      padding: 10px;
  }
  .ant-avatar{
        font-size: 12px;
        border-radius: 4px;
        vertical-align: middle;
        margin-right: 8px;
  }
</style>
