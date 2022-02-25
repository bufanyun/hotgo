<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol" ref="queryForm">
          <a-row :gutter="32">
            <a-col :span="6" >
              <a-form-item label="小页名称">
                <a-input v-model="queryParam.name" placeholder="请输入小页名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :span="6" >
              <a-form-item label="小页编码">
                <a-input v-model="queryParam.code" placeholder="请输入小页编码" allow-clear @keyup.enter.native="handleQuery"/>
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
      <!-- 增加 -->
      <sys-portlet-add-form
        v-if="showAddModal"
        ref="sysPortletAddForm"
        :showTitleOptions="showTitleOptions"
        :isAllowDragOptions="isAllowDragOptions"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showAddModal = false"
      />
      <!-- 编辑 -->
      <sys-portlet-edit-form
        v-if="showEditModal"
        ref="sysPortletEditForm"
        :showTitleOptions="showTitleOptions"
        :isAllowDragOptions="isAllowDragOptions"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showEditModal = false"
      />
      <advance-table
        title="工作台小页管理"
        :pagination="{
          current: queryParam.pageNum,
          pageSize: queryParam.pageSize,
          total: total,
          showSizeChanger: true,
          showLessItems: true,
          showQuickJumper: true,
          showTotal: (total, range) => `第 ${range[0]}-${range[1]} 条，总计 ${total} 条`,
          onChange: changeSize,
          onShowSizeChange: onShowSizeChange
        }"
        rowKey="id"
        size="middle"
        @refresh="getList"
        :columns="columns"
        :data-source="sysPortletList"
        :loading="loading"
        :format-conditions="true"
        tableKey="system-sysportlet-index-table"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
      >
        <div class="table-operations" slot="button">
          <a-button type="primary" @click="handleAdd" v-hasPermi="['system:sysPortlet:add']">
            <a-icon type="plus" />新增
          </a-button>
          <a-button type="" @click="handleExport" v-hasPermi="['system:sysPortlet:export']">
            <a-icon type="download" />导出
          </a-button>
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:sysPortlet:remove']">
            <a-icon type="delete" />删除
          </a-button>
        </div>
        <span slot="showTitle" slot-scope="{record}">
          {{ showTitleFormat(record) }}
        </span>
        <span slot="isAllowDrag" slot-scope="{record}">
          {{ isAllowDragFormat(record) }}
        </span>
        <span slot="status" slot-scope="{record}">
          {{ statusFormat(record) }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleUpdate(record)" v-hasPermi="['system:sysPortlet:edit']">
            修改
          </a>
          <a-divider type="vertical" v-hasPermi="['system:sysPortlet:remove']"/>
          <a @click="handleDelete(record)" v-hasPermi="['system:sysPortlet:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { listSysPortlet, delSysPortlet, exportSysPortlet } from '@/api/system/sysPortlet'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import SysPortletAddForm from '@/views/system/sysportlet/modules/SysPortletAddForm'
import SysPortletEditForm from '@/views/system/sysportlet/modules/SysPortletEditForm'
export default {
  name: 'SysPortlet',
  components: {
      AdvanceTable,
      SysPortletAddForm,
      SysPortletEditForm },
  data () {
    return {
      showAddModal: false,
      showEditModal: false,
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 选中的主键集合
      selectedRowKeys: [],
      // 选中的数据集合
      selectedRows: [],
      // 高级搜索 展开/关闭
      advanced: false,
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // label的百分比
      labelCol: { span: 6 },
      // 内容区域的百分比
      wrapperCol: { span: 18 },
      // 工作台小页管理表格数据
      sysPortletList: [],
      // 是否显示标题字典
      showTitleOptions: [],
      // 是否允许拖拽字典
      isAllowDragOptions: [],
      // 状态字典
      statusOptions: [],
      // 查询参数
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
        code: undefined
      },
      columns: [
        {
          title: '小页名称',
          dataIndex: 'name',
          ellipsis: true,
          width: '10%'
        },
        {
          title: '小页编码',
          dataIndex: 'code',
          ellipsis: true,
          width: '10%'
        },
        {
          title: '小页URL',
          dataIndex: 'url',
          ellipsis: true,
          width: '10%'
        },
        {
          title: '刷新频率',
          dataIndex: 'refreshRate',
          ellipsis: true,
          width: '10%'
        },
        {
          title: '是否显示标题',
          dataIndex: 'showTitle',
          scopedSlots: { customRender: 'showTitle' },
          width: '10%'
        },
        {
          title: '是否允许拖拽',
          dataIndex: 'isAllowDrag',
          scopedSlots: { customRender: 'isAllowDrag' },
          width: '10%'
        },
        {
          title: '排序号',
          dataIndex: 'sort',
          width: '10%'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          width: '10%'
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
  created () {
    this.getList()
    this.getDicts('sys_yes_no').then(response => {
      this.showTitleOptions = response.data
    })
    this.getDicts('sys_yes_no').then(response => {
      this.isAllowDragOptions = response.data
    })
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response.data
    })
  },
  methods: {
    /** 查询工作台小页管理列表 */
    getList () {
      this.loading = true
      listSysPortlet(this.queryParam).then(response => {
        this.sysPortletList = response.data.list
        this.total = response.data.total
        this.loading = false
      })
    },
    // 是否显示标题字典翻译
    showTitleFormat (row) {
        if (row.showTitle) {
            return this.selectDictLabel(this.showTitleOptions, row.showTitle)
        } else {
            return ''
        }
    },
    // 是否允许拖拽字典翻译
    isAllowDragFormat (row) {
        if (row.isAllowDrag) {
            return this.selectDictLabel(this.isAllowDragOptions, row.isAllowDrag)
        } else {
            return ''
        }
    },
    // 状态字典翻译
    statusFormat (row) {
        if (row.status) {
            return this.selectDictLabel(this.statusOptions, row.status)
        } else {
            return ''
        }
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
        name: undefined,
        code: undefined
      }
      this.handleQuery()
    },
    /** 翻页操作 */
    onShowSizeChange (current, pageSize) {
      this.queryParam.pageSize = pageSize
      this.getList()
    },
    /** 翻页操作 */
    onSizeChange (current, size) {
        this.queryParam.pageNum = 1
        this.queryParam.pageSize = size
        this.getList()
    },
    /** 翻页操作 */
    changeSize (current, pageSize) {
      this.queryParam.pageNum = current
      this.queryParam.pageSize = pageSize
      this.getList()
    },
    /** 翻页操作 */
    onSelectChange (selectedRowKeys, selectedRows) {
        this.selectedRowKeys = selectedRowKeys
        this.selectedRows = selectedRows
        this.ids = this.selectedRows.map(item => item.id)
        this.single = selectedRowKeys.length !== 1
        this.multiple = !selectedRowKeys.length
    },
    /** 查询折叠和展开操作 */
    toggleAdvanced () {
        this.advanced = !this.advanced
    },
    handleAdd () {
          this.showAddModal = true
          this.$nextTick(() => (
                  this.$refs.sysPortletAddForm.handleAdd()
          ))
    },
    handleUpdate (record, ids) {
          this.showEditModal = true
          this.$nextTick(() => (
                  this.$refs.sysPortletEditForm.handleUpdate(record, ids)
          ))
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const sysPortletIds = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        onOk () {
          return delSysPortlet(sysPortletIds)
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
          return exportSysPortlet(that.queryParam)
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
