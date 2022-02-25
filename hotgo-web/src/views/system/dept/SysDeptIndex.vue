<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper" ref="search">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="部门名称">
                <a-input v-model="queryParam.name" placeholder="请输入部门名称" allow-clear @keyup.enter.native="handleQuery"/>
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
    </a-card>

    <a-card :bordered="false" class="table-card">

      <sys-dept-add-form
        v-if="showAddModal"
        ref="sysDeptAddForm"
        :deptTypeOptions="deptTypeOptions"
        :statusOptions="statusOptions"
        :deptOptions="deptOptions"
        :expandedKeys="expandedKeys"
        @getTreeselect="getTreeselect"
        @ok="getList"
        @close="showAddModal = false"
      />

      <sys-dept-edit-form
        v-if="showEditModal"
        ref="sysDeptEditForm"
        :deptTypeOptions="deptTypeOptions"
        :statusOptions="statusOptions"
        :deptOptions="deptOptions"
        :expandedKeys="expandedKeys"
        @getTreeselect="getTreeselect"
        @ok="getList"
        @close="showEditModal = false"
      />
      <!-- 上传文件 -->
      <import-excel
        ref="importExcel"
        title="导入部门"
        upload-msg="是否更新已经存在的部门数据"
        upload-action="/system/user/importData"
        @ok="getList"
        @importTemplate="importTemplate"
      />
      <!-- 数据展示 -->
      <advance-table
        :loading="loading"
        title="部门管理"
        rowKey="id"
        @refresh="getList"
        :expandIconColumnIndex="1"
        :columns="columns"
        :data-source="list"
        :indentSize="16"
        size="middle"
        tableKey="system-dept-SysDeptIndex-table"
        :defaultExpandedRowKeys="expandedRowKeys"
        :expandedRowKeys="expandedRowKeys"
        :expandIcon="expandIcon"
        @expand="expandNode"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }">
        <div class="table-operations" slot="button">
          <a-space align="center">
            <a-button type="primary" @click="handleAdd()" v-hasPermi="['system:dept:add']">
              <a-icon type="plus" />新增
            </a-button>
            <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:dept:remove']">
              <a-icon type="delete" />删除
            </a-button>
            <a-button type="" @click="$refs.importExcel.importExcelHandleOpen()" v-hasPermi="['system:user:import']">
              <a-icon type="import" />导入
            </a-button>
            <!-- <a-button @click="handleImport" v-hasPermi="['system:dept:import']">
              <a-icon type="download" />导入
            </a-button> -->
            <a-button @click="handleExport" v-hasPermi="['system:dept:export']">
              <a-icon type="download" />导出
            </a-button>
          </a-space>
        </div>
        <span slot="name" slot-scope="{text, record}">
          <a-icon v-show="record.deptType ==='org' " class="depIcon" :component="allIcon.companyFillIcon" />
          <a-icon v-show="record.deptType === 'company'" :component="allIcon.companyIcon" class="depIcon" />
          <a-icon v-show="record.deptType === 'dept'" :component="allIcon.connectionsIcon" class="depIcon" />
          <span v-if="text.indexOf(queryParam.name) > -1">
            {{ text.substr(0, text.indexOf(queryParam.name)) }}
            <span style="color: #f50">{{ queryParam.name }}</span>
            {{ text.substr(text.indexOf(queryParam.name) + queryParam.name.length) }}
          </span>
          <span v-else>{{ text }}</span>
        </span>
        <span slot="status" slot-scope="{record}">
          <a-badge :status="record.status === '1' ? 'processing' : 'error'" :text=" statusFormat(record) " />
        </span>
        <span slot="created_at" slot-scope="{text, record}">
          {{ parseTime(record.created_at) }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleUpdate(record)" v-hasPermi="['system:dept:edit']">
            修改
          </a>
          <a-divider type="vertical" v-hasPermi="['system:dept:add']" />
          <a @click="handleAdd(record)" v-hasPermi="['system:dept:add']">
            添加子部门
          </a>
          <a-divider type="vertical" v-if="record.pid != 0" v-hasPermi="['system:dept:remove']" />
          <a @click="handleDelete(record)" v-if="record.pid != 0" v-hasPermi="['system:dept:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>

<script>

import { listDept, delDept, searchDeptList, listDeptTree, listDeptExcludeChild } from '@/api/system/dept'
import SysDeptEditForm from './modules/SysDeptEditForm'
import SysDeptAddForm from './modules/SysDeptAddForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import { exportConfig } from '@/api/system/config'
import ImportExcel from '@/components/pt/import/ImportExcel'
import allIcon from '@/core/icons'
import { importTemplate } from '@/api/system/user'

export default {
  name: 'Dept',
  components: {
    SysDeptAddForm,
    SysDeptEditForm,
    AdvanceTable,
    ImportExcel,
    allIcon
  },
  data () {
    return {
      showAddModal: false,
      showEditModal: false,
      allIcon,
      list: [],
      // 部门树选项
      deptOptions: [],
      // 状态数据字典
      statusOptions: [],
      expandedKeys: [], // 表单页面默认展开节点
      expandedRowKeys: [], // 树表格展开节点
      selectedRowKeys: [],
      selectedRows: [],
      ids: [],
      multiple: true,
      loading: false,
      // 状态数据字典
      deptTypeOptions: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      queryParam: {
        name: ''
      },
      columns: [
        {
          title: '部门名称',
          dataIndex: 'name',
          width: '250px',
          scopedSlots: { customRender: 'name' }
        },
        {
          title: '负责人',
          dataIndex: 'leader',
          scopedSlots: { customRender: 'leader' }
        },
        {
          title: '联系电话',
          dataIndex: 'phone',
          scopedSlots: { customRender: 'phone' }
        },
        {
          title: '邮箱',
          dataIndex: 'email',
          ellipsis: true,

          scopedSlots: { customRender: 'email' }
        },
        {
          title: '状态',
          dataIndex: 'status',
           width: '120px',
          scopedSlots: { customRender: 'status' },
           align: 'center'
        },
        {
          title: '排序',
          dataIndex: 'sort',
          align: 'center'
        },
        {
          title: '创建时间',
          dataIndex: 'created_at',
          ellipsis: true,
          scopedSlots: { customRender: 'created_at' },
          align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: '240px',
          scopedSlots: { customRender: 'operation' }
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
    this.getDicts('sys_dept_type').then(response => {
      this.deptTypeOptions = response
    })
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response
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
    expandNode (expanded, record) {
      //  展开收缩时需要动态修改展开行集合
      if (expanded) {
         this.expandedRowKeys.push(record.id)
      } else {
        this.expandedRowKeys = this.expandedRowKeys.filter(
        function (item) { return item !== record.id }
        )
      }
      if (expanded && (record.children == null || record.children.length === 0)) {
       this.loading = true
        listDept(this.queryParam, record.id, 1).then(response => {
            // record.children = this.handleTree(response, 'id')
            record.children = response
            this.loading = false
          }
        )
        }
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
      this.ids = this.selectedRows.map(item => item.id)
      this.multiple = !selectedRowKeys.length
    },
    /** 查询菜单下拉树结构 */
    getTreeselect (row) {
      if (!row) {
        listDeptTree('', 3).then(response => {
           this.expandTree(response, 3, this.expandedRowKeys)
           this.deptOptions = response
        })
      } else {
        listDeptExcludeChild(row.id).then(response => {
          // this.deptOptions = this.handleTree(response, 'id')
          this.deptOptions = response
        })
      }
    },
    /** 查询定时任务列表 */
    getList () {
      this.loading = true
      listDept(this.queryParam, '', 3).then(response => {
          this.expandedRowKeys = []
           // 加载完数据后需要迭代计算需要展开的行数据
          this.expandTree(response, 3, this.expandedRowKeys)
          // this.list = this.handleTree(response, 'id')
           this.list = response
          this.loading = false
        }
      )
    },

    /** 搜索按钮操作 */
    handleQuery () {
      if (this.queryParam.name === '') {
         this.expandedRowKeys = []
         this.getList()
      } else {
          this.loading = true
          searchDeptList(this.queryParam).then(response => {
          this.expandedRowKeys = []
          if (response != null && response.length !== 0) {
            this.getAllDeptNode(response)
            // this.list = this.handleTree(response, 'id')
            this.list = response
          } else {
             this.list = []
          }

          this.loading = false
        }
      )
      }
    },
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.queryParam = {
        name: ''
      }
      this.handleQuery()
    },
    handleAdd (record) {
      this.showAddModal = true
      this.$nextTick(() => (
          this.$refs.sysDeptAddForm.handleAdd(record)
       ))
    },
     handleUpdate (record) {
      this.showEditModal = true
      this.$nextTick(() => (
         this.$refs.sysDeptEditForm.handleUpdate(record)
       ))
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const id = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        onOk () {
          return delDept(id)
            .then(() => {
              if (row !== null) {
                that.removeTreeNode(that.list, row)
              } else {
                 that.onSelectChange([], [])
                 that.getList()
              }
              that.$message.success(
                '删除成功',
                3
              )
          })
        },
        onCancel () {}
      })
    },
    /** 导入按钮操作 */
    handleImport () {
      alert('暂不支持')
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
    getAllDeptNode (nodes) {
      if (!nodes || nodes.length === 0) {
        return []
      }
      nodes.forEach(node => {
        this.expandedRowKeys.push(node.id)
        return this.getAllDeptNode(node.children)
      })
    },
    expandIcon (props) {
      if (props.record.treeLeaf === 'y') {
        return <span style="margin-right:22px"></span>
      } else {
        if (props.expanded) {
          return (
            <a style="color: 'black',margin-right:0px"
               onClick={(e) => {
                 props.onExpand(props.record, e)
               }}
            >
              <a-icon type="caret-down" />
            </a>
          )
        } else {
          return (
            <a style="color: 'black' ,margin-right:0px"
               onClick={(e) => {
                 props.onExpand(props.record, e)
               }}
            >
              <a-icon type="caret-right" />
            </a>
          )
        }
      }
    }
  }
}
</script>
