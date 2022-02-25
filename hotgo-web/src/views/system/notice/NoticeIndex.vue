<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="公告标题">
                <a-input v-model="queryParam.title" placeholder="请输入" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="公告类型">
                <a-select placeholder="请选择" v-model="queryParam.type" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in typeOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
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
      <!-- 增加组件按需加载 -->
      <notice-add-form
        v-if="showAddModal"
        ref="noticeAddForm"
        :typeOptions="typeOptions"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showAddModal = false"
      />
      <!-- 组件按需加载-->
      <notice-edit-form
        v-if="showEditModal"
        ref="noticeEditForm"
        :typeOptions="typeOptions"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showEditModal = false"
      />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="通知公告"
        :loading="loading"
        rowKey="id"
        size="middle"
        tableKey="system-notic-NoticIndex-table"
        @refresh="getList"
        :format-conditions="true"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        :pagination="{
          current: queryParam.page,
          pageSize: queryParam.limit,
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
          <a-button type="primary" @click="handleAdd" v-hasPermi="['system:notice:add']">
            <a-icon type="plus" />新增
          </a-button>
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:notice:remove']">
            <a-icon type="delete" />删除
          </a-button>
        </div>
        <span slot="type" slot-scope="{text, record}">
          <a-tag :color="text | noticeTypeFilter">
            {{ typeFormat(record) }}
          </a-tag>
        </span>
        <span slot="status" slot-scope="{text, record}">
          <a-badge status="processing" :text=" statusFormat(record) " />
        </span>
        <span slot="created_at" slot-scope="{text, record}">
          {{ parseTime(record.created_at, '{y}-{m}-{d}') }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleUpdate(record)" v-hasPermi="['system:notice:edit']">
            修改
          </a>
          <a-divider type="vertical" v-hasPermi="['system:notice:remove']" />
          <a @click="handleDelete(record)" v-hasPermi="['system:notice:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { listNotice, delNotice } from '@/api/system/notice'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import NoticeAddForm from '@/views/system/notice/modules/NoticeAddForm'
import NoticeEditForm from '@/views/system/notice/modules/NoticeEditForm'
export default {
  name: 'Notice',
  components: {
    NoticeAddForm,
    AdvanceTable,
    NoticeEditForm
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
      // 类型数据字典
      typeOptions: [],
      statusOptions: [],
      // 日期范围
      dateRange: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      queryParam: {
        page: 1,
        limit: 10,
        title: undefined,
        status: undefined
      },
      addModalRefName: 'addModal', // 添加弹窗ref名称
      columns: [
        {
          title: '公告标题',
          dataIndex: 'title',
          ellipsis: true,
           width: '30%'
        },
        {
          title: '公告类型',
          dataIndex: 'type',
          scopedSlots: { customRender: 'type' },
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
          scopedSlots: { customRender: 'created_at' }
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: '200',
          scopedSlots: { customRender: 'operation' }
        }
      ],
      showAddModal: false,
      showEditModal: false
    }
  },
  filters: {
    noticeTypeFilter (type) {
      let value = 'blue'
      if (type === '1') {
        value = 'orange'
      }
      return value
    }
  },
  created () {
    this.getDicts('sys_notice_status').then(response => {
      this.statusOptions = response
    })
    this.getDicts('sys_notice_type').then(response => {
      this.typeOptions = response
    })
    this.getList()
  },
  computed: {
  },
  watch: {
  },
  methods: {
    /** 查询定时任务列表 */
    getList () {
      this.loading = true
      listNotice(this.queryParam).then(response => {
          this.list = response.list
          this.total = response.total_count
          this.loading = false
        }
      )
    },
    // 公告状态字典翻译
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    // 公告类型字典翻译
    typeFormat (row) {
      return this.selectDictLabel(this.typeOptions, row.type)
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.page = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.queryParam = {
        page: 1,
        limit: 10,
        title: undefined,
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
    handleAdd () {
      this.showAddModal = true
      this.$nextTick(() => (
         this.$refs.noticeAddForm.handleAdd()
       ))
    },
    handleUpdate (record) {
      this.showEditModal = true
      this.$nextTick(() => (
         this.$refs.noticeEditForm.handleUpdate(record)
       ))
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const noticeIds = row.id || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        onOk () {
          return delNotice(noticeIds)
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
    }
  }

}
</script>
