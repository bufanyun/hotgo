<template>
  <div>
    <dict-data-edit-form
      v-if="showEditSubModal"
      ref="dictDataEditForm"
      :title="title"
      :statusOptions="statusOptions"
      :type="type"
      @ok="getList"
      @close="showEditSubModal = false"
    />
    <dict-data-add-form
      v-if="showAddSubModal"
      ref="dictDataAddForm"
      :statusOptions="statusOptions"
      :type="type"
      :title="title"
      @ok="getList"
      @close="showAddSubModal = false"
    />
    <a-card :bordered="false">
      <div class="table-operations">
        <div class="table_title">{{ title }}</div>
        <a-button type="primary" @click="handleSubAdd()" v-hasPermi="['system:dict:add']" >
          <a-icon type="plus" />新增
        </a-button>
        <a-button type="danger" v-if="!multiple" @click="handleSubDelete" v-hasPermi="['system:dict:remove']">
          <a-icon type="delete" />删除
        </a-button>
        <a-tooltip title="刷新">
          <a-icon @click="getList" class="action" :type="loading ? 'loading' : 'reload'" />
        </a-tooltip>
      </div>
      <a-table
        :loading="loading"
        rowKey="id"
        size="middle"
        :columns="columns"
        :data-source="list"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
        :pagination="false">
        <span slot="status" slot-scope="text, record">
          <a-badge :status="record.status == '1' ? 'processing' : 'error'" :text=" statusFormat(record) " />
        </span>
        <span slot="createTime" slot-scope="text, record">
          {{ parseTime(record.created_at) }}
        </span>
        <span slot="operation" slot-scope="text, record">
          <a @click="handleSubUpdate(record)" v-hasPermi="['system:dict:edit']">
            修改
          </a>
          <a-divider type="vertical" />
          <a @click="handleSubDelete(record)" v-hasPermi="['system:dict:remove']">
            删除
          </a>
        </span>
      </a-table>
    </a-card>
  </div>
</template>

<script>

import { listData, delData, exportData } from '@/api/system/dict/data'
import DictDataEditForm from './modules/DictDataEditForm'
import DictDataAddForm from './modules/DictDataAddForm'
export default {
  name: 'DictData',
  props: {
    type: {
      type: String,
      required: true
    },
    dictId: {
      type: Number,
      required: true
    },
    title: {
      type: String,
      default: '子表'
    }
  },
  components: {
    DictDataEditForm,
    DictDataAddForm
  },
  data () {
    return {
      showAddSubModal: false,
      showEditSubModal: false,
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
      dictLabels: [],
      loading: false,
      total: 0,
      // 状态数据字典
      statusOptions: [],
      queryParam: {
        page: 1,
        limit: 100,
        dictName: undefined,
        type: undefined,
        status: undefined
      },
      columns: [
        {
          title: '字典标签',
          dataIndex: 'label',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '字典键值',
          dataIndex: 'value',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '字典排序',
          dataIndex: 'sort',
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
          ellipsis: true,
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
          width: '15%',
          scopedSlots: { customRender: 'operation' },
          align: 'center'
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.queryParam.type = this.type
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
    getList () {
      this.showAddSubModal = false
      this.showEditSubModal = false
      this.loading = true
      listData(this.queryParam).then(response => {
        this.list = response.list
        this.total = response.total_count
          this.loading = false
        }
      )
    },
    // 状态字典翻译
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.page = 1
      this.getList()
    },
    handleSubAdd () {
      this.showAddSubModal = true
      this.$nextTick(() => (
        this.$refs.dictDataAddForm.handleAdd()
      ))
    },
    handleSubUpdate (record) {
      this.showEditSubModal = true
      this.$nextTick(() => (
        this.$refs.dictDataEditForm.handleUpdate(record)
      ))
    },
    /** 重置按钮操作 */
    resetQuery () {
      this.dateRange = []
      this.queryParam = {
        page: 1,
        limit: 10,
        name: undefined,
        type: undefined,
        status: undefined
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
      this.dictLabels = this.selectedRows.map(item => item.dictLabel)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    /** 删除按钮操作 */
    handleSubDelete (row) {
      var that = this
      const dictCodes = row.id || this.ids
      const dictLabels = row.dictLabel || this.dictLabels
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中字典编码为"' + dictLabels + '"的数据',
        onOk () {
          return delData(dictCodes)
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
          return exportData(that.queryParam)
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
