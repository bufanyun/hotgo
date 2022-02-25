<template>
  <div>
    <a-card :bordered="false"style="margin-bottom:10px" >
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="字典名称">
                <a-input v-model="queryParam.name" placeholder="请输入字典名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="字典类型">
                <a-input v-model="queryParam.type" placeholder="请选择字典类型" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="状态">
                <a-select placeholder="字典状态" v-model="queryParam.status" style="width: 100%">
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24" v-if="advanced">
              <a-form-item label="创建时间">
                <a-range-picker style="width: 100%" v-model="dateRange" valueFormat="YYYY-MM-DD" format="YYYY-MM-DD" />
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
      <!-- 修改 -->
      <dict-type-Edit-form
        v-if="showEditModal"
        ref="dictTypeEditForm"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showEditModal = false"
      />
      <!-- 添加 -->
      <dict-type-add-form
        v-if="showAddModal"
        ref="dictTypeAddForm"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showAddModal = false"
      />
      <!-- 数据展示 -->
      <advance-table
        :loading="loading"
        rowKey="id"
        :columns="columns"
        :expandedRowKeys="expandedKeys"
        @expand="onExpandCurrent"
        :data-source="list"
        title="字典管理"
        size="middle"
        tableKey="system-dict-DictIndex-table"
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
        <!-- 操作 -->
        <div class="table-operations" slot="button">
          <a-button type="primary" @click="handleAdd()" v-hasPermi="['system:dict:add']">
            <a-icon type="plus" />新增
          </a-button>
          <a-button type="danger" v-if="!multiple" @click="handleDelete" v-hasPermi="['system:dict:remove']">
            <a-icon type="delete" />删除
          </a-button>
          <a-button type="" @click="handleExport" v-hasPermi="['system:dict:export']">
            <a-icon type="download" />导出
          </a-button>
          <a-button type="" @click="handleRefreshCache" v-hasPermi="['system:dict:remove']">
            <a-icon type="redo" />刷新缓存
          </a-button>
        </div>
        <span
          slot="expandedRowRender"
          slot-scope="{ record}"
        >
          <dict-data-index
            ref="dictDataIndex"
            :title="record.name"
            :dictId="record.id"
            :type="record.type"
          />
        </span>
        <span slot="status" slot-scope="{text, record}" >
          <a-badge :status="record.status == '1' ? 'processing' : 'error'" :text="statusFormat(record) "/>
        </span>
        <span slot="created_at" slot-scope="{text, record}">
          {{ parseTime(record.created_at) }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleUpdate(record, undefined)" v-hasPermi="['system:dict:edit']">
            修改
          </a>
          <a-divider type="vertical" />
          <a @click="handleDelete(record)" v-hasPermi="['system:dict:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>

import { listType, delType, refreshCache } from '@/api/system/dict/type' // exportType,
import DictTypeEditForm from './modules/DictTypeEditForm'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import DictTypeAddForm from './modules/DictTypeAddForm'
import DictDataIndex from './DictDataIndex'
import { exportDownload } from '@/utils/aidex'
export default {
  name: 'Dict',
  components: {
    DictTypeEditForm,
    DictTypeAddForm,
    DictDataIndex,
    AdvanceTable
  },
  data () {
    return {
      showAddModal: false,
      showEditModal: false,
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
      names: [],
      loading: false,
      total: 0,
      // 状态数据字典
      statusOptions: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      // 日期范围
      dateRange: [],
      queryParam: {
        page: 1,
        limit: 10,
        name: undefined,
        type: undefined,
        status: undefined
      },
      expandedKeys: [],
      columns: [
        {
          title: '字典名称',
          dataIndex: 'name',
          ellipsis: true,
          align: 'left'
        },
        {
          title: '字典类型',
          dataIndex: 'type',
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
          align: 'left'
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
          width: '10%',
          scopedSlots: { customRender: 'operation' },
          align: 'center'
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
    /** 查询定时任务列表 */
    getList () {
      this.showAddModal = false
      this.showEditModal = false
      this.loading = true
      listType(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.list
          this.total = response.total_count
          this.loading = false
        }
      )
    },
    // 参数系统内置字典翻译
    statusFormat (row) {
      return this.selectDictLabel(this.statusOptions, row.status)
    },
    /** 搜索按钮操作 */
    handleQuery () {
      this.queryParam.page = 1
      this.getList()
    },
    handleAdd () {
      this.showAddModal = true
      this.$nextTick(() => (
        this.$refs.dictTypeAddForm.handleAdd()
      ))
    },
    handleUpdate (record, ids) {
      this.showEditModal = true
      this.$nextTick(() => (
        this.$refs.dictTypeEditForm.handleUpdate(record, ids)
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
      this.names = this.selectedRows.map(item => item.name)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const dictIds = row.id || this.ids
      const names = row.name || this.names
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中字典编号为"' + names + '"的数据',
        onOk () {
          return delType(dictIds)
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
          exportDownload('/dict_type/export', that.queryParam)
          // return exportType(that.queryParam)
          //   .then(response => {
          //     that.download(response.msg)
          //     // window.location.href = process.env.VUE_APP_BASE_API + '/common/download?fileName=' + encodeURI('1.xlsx') + '&delete=true'
          //     that.$message.success(
          //       '导出成功',
          //       3
          //     )
          //   })
        },
        onCancel () {}
      })
    },
    /** 清理缓存按钮操作 */
    handleRefreshCache () {
      refreshCache().then(response => {
        this.$message.success(
          '刷新成功',
          3
        )
      })
    },
    onExpandCurrent (expandedKeys, row) {
      if (expandedKeys) {
        this.expandedKeys = [row.id]
      } else {
        this.expandedKeys = []
      }
    }
  }
}
</script>
