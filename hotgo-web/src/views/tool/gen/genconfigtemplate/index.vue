<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="模板名称">
                <a-input v-model="queryParam.templateName" placeholder="请输入模板名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="作者">
                <a-input v-model="queryParam.authorName" placeholder="请输入作者" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col>
              <span class="table-page-search-submitButtons" style="float:right;">
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
      <gen-config-template-add-form
        ref="genConfigTemplateAddForm"
        :statusOptions="statusOptions"
        :templateDefaultOptions="templateDefaultOptions"
        @ok="getList"
      />
      <!-- 编辑 -->
      <gen-config-template-edit-form
        ref="genConfigTemplateEditForm"
        :statusOptions="statusOptions"
        :templateDefaultOptions="templateDefaultOptions"
        @ok="getList"
      />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="模板配置"
        :loading="loading"
        rowKey="id"
        size="middle"
        @refresh="getList"
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
          onChange: onSelectChange,
          onShowSizeChange: onSizeChange,
        }"
      >
        <div class="table-operations" slot="button">
          <a-button type="primary" @click="$refs.genConfigTemplateAddForm.handleAdd()" >
            <a-icon type="plus" />新增
          </a-button>
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete">
            <a-icon type="delete" />删除
          </a-button>
        </div>
        <span slot="createTime" slot-scope="{text, record}">
          {{ parseTime(record.createTime, '{y}-{m}-{d}') }}
        </span>
        <span slot="status" slot-scope="{text, record}">
          <a-popconfirm
            ok-text="是"
            cancel-text="否"
            @confirm="confirmHandleStatus(record)"
            @cancel="cancelHandleStatus(record)"
          >
            <span slot="title">确认<b>{{ record.status === '1' ? '启用' : '停用' }}</b>{{ record.templateName }}的模板吗?</span>
            <a-switch checked-children="正常" un-checked-children="停用" :checked="record.status == 0" />
          </a-popconfirm>
        </span>
        <span slot="templateDefault" slot-scope="{text, record}">
          <a-popconfirm
            ok-text="是"
            cancel-text="否"
            @confirm="confirmHandleTemplateDefault(record)"
            @cancel="cancelHandleTemplateDefault(record)"
          >
            <span slot="title">确认<b>{{ record.templateDefault === 'N' ? '设置' : '取消' }}</b>{{ record.templateName }}的模板为默认模板吗?</span>
            <a-switch checked-children="是" un-checked-children="否" :checked="record.templateDefault == 'Y'" />
          </a-popconfirm>
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="$refs.genConfigTemplateEditForm.handleUpdate(record)" v-hasPermi="['gen:template:edit']">
            修改
          </a>
          <a-divider type="vertical" v-hasPermi="['gen:template:remove']" />
          <a @click="handleDelete(record)" v-hasPermi="['gen:template:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { listTemplate, changeStatus, changeTemplateDefault, delTemplate } from '@/api/tool/genConfigTemplate'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import GenConfigTemplateAddForm from '@/views/tool/gen/genconfigtemplate/module/GenConfigTemplateAddForm'
import GenConfigTemplateEditForm from '@/views/tool/gen/genconfigtemplate/module/GenConfigTemplateEditForm'
export default {
  name: 'Template',
  components: {
    AdvanceTable,
    GenConfigTemplateAddForm,
    GenConfigTemplateEditForm
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
      total: 0,
      // 类型数据字典
      statusOptions: [],
      templateDefaultOptions: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        templateName: undefined,
        authorName: undefined
      },
      // 添加弹窗ref名称
      addModalRefName: 'addModal',
      columns: [
        {
          title: '模板名称',
          dataIndex: 'templateName',
          ellipsis: true,
          align: 'left',
          width: '10%'
        },
        {
          title: '作者',
          dataIndex: 'functionAuthor',
          align: 'left'
        },
        {
          title: '邮箱',
          dataIndex: 'functionAuthorEmail',
          ellipsis: true,
          align: 'left'
        },
        {
          title: '创建时间',
          dataIndex: 'createTime',
          scopedSlots: { customRender: 'createTime' },
          align: 'center'
        },
        {
          title: '工作空间',
          dataIndex: 'workspacePath',
          ellipsis: true,
          align: 'left'
        },
        {
          title: '模块名',
          dataIndex: 'moduleName',
          align: 'left'
        },
        {
          title: '包路径',
          dataIndex: 'packageName',
          align: 'left'
        },
        {
          title: '前端工作空间',
          dataIndex: 'webWorkspacePath',
          ellipsis: true,
          align: 'left'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' },
          align: 'center'
        },
        {
          title: '是否默认',
          dataIndex: 'templateDefault',
          scopedSlots: { customRender: 'templateDefault' },
          align: 'center'
        },
        {
          title: '排序',
          dataIndex: 'sort',
          align: 'center'
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
    this.getDicts('sys_normal_disable').then(response => {
      this.statusOptions = response.data
    })
    this.getDicts('sys_yes_no').then(response => {
      this.templateDefaultOptions = response.data
    })
    this.getList()
  },
  computed: {
  },
  watch: {
  },
  methods: {
    /** 查询列表 */
    getList () {
      this.loading = true
      listTemplate(this.queryParam).then(response => {
          this.list = response.data.list
          this.total = response.data.total
          this.loading = false
        }
      )
    },
    /* 状态修改 */
    confirmHandleStatus (row) {
      const text = row.status === '1' ? '启用' : '停用'
      row.status = row.status === '0' ? '1' : '0'
      changeStatus(row.id, row.status)
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
    /* 是否默认修改 */
    confirmHandleTemplateDefault (row) {
      const text = row.templateDefault === 'Y' ? '取消默认' : '设置默认'
      row.templateDefault = row.templateDefault === 'Y' ? 'N' : 'Y'
      changeTemplateDefault(row.id, row.templateDefault)
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
    cancelHandleTemplateDefault (row) {
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
        templateName: undefined,
        authorName: undefined
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
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const templateIds = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        onOk () {
          return delTemplate(templateIds)
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
    /** 修改按钮操作 */
    handleUpdate (row, ids) {
      const templateId = row ? row.id : ids
      this.$router.push({
        name: 'NoticeForm',
        params:
          {
            id: templateId,
            formTitle: '修改模板'
          }
      })
    }
  }
}
</script>
