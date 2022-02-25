<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="32">
            <a-col :span="6">
              <a-form-item label="表名称">
                <a-input v-model="queryParam.tableName" placeholder="请输入表名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :span="6">
              <a-form-item label="表描述">
                <a-input v-model="queryParam.tableComment" placeholder="请输入表描述" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :span="6">
              <a-form-item label="作者">
                <a-input v-model="queryParam.functionAuthor" placeholder="请输入作者" allow-clear @keyup.enter.native="handleQuery"/>
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
      <advance-table
        :columns="columns"
        :data-source="list"
        title="代码生成"
        :loading="loading"
        rowKey="tableId"
        size="middle"
        @refresh="getList"
        :format-conditions="true"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
        :pagination="{
          current: queryParam.pageNum,
          pageSize: queryParam.pageSize,
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
          <a-button type="primary" @click="$refs.importTable.show()" v-hasPermi="['tool:gen:import']">
            <a-icon type="cloud-upload" />导入
          </a-button>
          <a-button :disabled="single" @click="handleGenTable" v-hasPermi="['tool:gen:code']">
            <a-icon type="cloud-download" />生成
          </a-button>
          <a-button :disabled="single" @click="handleEditTable" v-hasPermi="['tool:gen:edit']">
            <a-icon type="edit" />修改
          </a-button>
          <a-button type="danger" v-if="!multiple" :disabled="multiple" @click="handleDelete" v-hasPermi="['tool:gen:remove']">
            <a-icon type="delete" />删除
          </a-button>
        </div>
        <span slot="createTime" slot-scope="{text, record}">
          {{ parseTime(record.createTime) }}
        </span>
        <span slot="updateTime" slot-scope="{text, record}">
          {{ parseTime(record.updateTime) }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="$refs.previewCode.handlePreview(record)" v-hasPermi="['tool:gen:preview']">
            <a-icon type="eye" />预览
          </a>
          <a-divider type="vertical" v-hasPermi="['tool:gen:edit']" />
          <a @click="handleEditTable(record)" v-hasPermi="['tool:gen:edit']">
            <a-icon type="edit" />编辑
          </a>
          <a-divider type="vertical" v-hasPermi="['tool:gen:remove']" />
          <a @click="handleDelete(record)" v-hasPermi="['tool:gen:remove']">
            <a-icon type="delete" />删除
          </a>
          <a-divider type="vertical" v-hasPermi="['tool:gen:edit']" />
          <a @click="handleSynchDb(record)" v-hasPermi="['tool:gen:edit']">
            <a-icon type="cloud-sync" />同步
          </a>
          <a-divider type="vertical" v-hasPermi="['tool:gen:code']" />
          <a @click="handleGenTable(record)" v-hasPermi="['tool:gen:code']">
            <a-icon type="cloud-download" />生成代码
          </a>
          <a-divider type="vertical" v-hasPermi="['tool:gen:code']" />
          <a @click="handleAddMenu(record)" v-hasPermi="['tool:gen:code']">
            <a-icon type="cloud-upload" />创建菜单
          </a>
        </span>
      </advance-table>
      <!-- 预览 -->
      <preview-code ref="previewCode" />
      <!-- 导入 -->
      <import-table ref="importTable" @ok="handleOk" />
    </a-card>
  </div>
</template>

<script>
import { delTable, listTable, synchDb, genCode, addMenu } from '@/api/tool/gen'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import PreviewCode from './modules/PreviewCode'
import ImportTable from './modules/ImportTable'
import { downLoadZip } from '@/utils/zipdownload'
import storage from 'store'
export default {
  name: 'Gen',
  components: {
    PreviewCode,
    ImportTable,
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
      // 选中表数组
      tableNames: [],
      loading: false,
      total: 0,
      // 日期范围
      dateRange: [],
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      // 查询参数
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        tableName: undefined,
        tableComment: undefined
      },
      // 表头
      columns: [
        {
          title: '序号',
          dataIndex: 'tableId',
          align: 'center',
          width: 45
        },
        {
          title: '表名称',
          dataIndex: 'tableName',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '表描述',
          dataIndex: 'tableComment',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '实体',
          dataIndex: 'className',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '作者',
          dataIndex: 'functionAuthor',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '创建时间',
          dataIndex: 'createTime',
          ellipsis: true,
          scopedSlots: { customRender: 'createTime' },
          align: 'center'
        },
        {
          title: '更新时间',
          dataIndex: 'updateTime',
          scopedSlots: { customRender: 'updateTime' },
          ellipsis: true,
          align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'action',
          width: '40%',
          scopedSlots: { customRender: 'operation' },
          align: 'center'
        }
      ]
    }
  },
  created () {
    this.getList()
  },
  methods: {
    /** 查询表集合 */
    getList () {
      this.loading = true
      listTable(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.rows
          this.total = response.total
          this.loading = false
        }
      )
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
        tableName: undefined,
        tableComment: undefined
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
      this.ids = this.selectedRows.map(item => item.tableId)
      this.tableNames = this.selectedRows.map(item => item.tableName)
      this.single = selectedRowKeys.length !== 1
      this.multiple = !selectedRowKeys.length
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    /** 删除按钮操作 */
    handleDelete (row) {
      var that = this
      const tableIds = row.tableId || this.ids
      this.$confirm({
        title: '确认删除所选中数据?',
        content: '当前选中编号为' + tableIds + '的数据',
        onOk () {
          return delTable(tableIds)
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
    /** 同步数据库操作 */
    handleSynchDb (row) {
      var that = this
      const tableName = row.tableName
      this.$confirm({
        title: '确认强制同步数据?',
        content: '当前同步表名为' + tableName + '的数据',
        onOk () {
          return synchDb(tableName)
            .then(() => {
              that.onSelectChange([], [])
              that.getList()
              that.$message.success(
                '同步成功',
                3
              )
          })
        },
        onCancel () {}
      })
    },
    /** 修改按钮操作 */
    handleEditTable (row) {
      const tableId = row.tableId || this.ids[0]
      const genTableId = 'genTableId'
      storage.set(genTableId, tableId)
      this.$router.push({
        name: 'GenEdit',
        params:
        {
          tableId: tableId
        }
      })
    },
    /** 生成代码操作  */
    handleGenTable (row) {
      this.loading = true
      const tableNames = row.tableName || this.tableNames
      if (tableNames === '' || tableNames.length === 0) {
        this.$message.error(
          '请选择要生成的数据',
          3
        )
        this.loading = false
        return
      }
      if (row.genType !== undefined) {
        if (row.genType === '1') {
          genCode(row.tableName).then(response => {
            this.$notification.open({
              message: '提示',
              description: response.msg,
              duration: 3
            })
          }).finally(() => {
            this.loading = false
          })
        } else {
          downLoadZip('/tool/gen/batchGenCode?tables=' + tableNames, 'aidex')
          this.loading = false
        }
      } else {
        this.selectedRows.forEach(node => {
          if (node.genType === '1') {
            genCode(node.tableName).then(response => {
              this.$notification.open({
                message: '提示',
                description: response.msg,
                duration: 3
              })
            }).finally(() => {
              this.loading = false
            })
          } else {
            downLoadZip('/tool/gen/batchGenCode?tables=' + node.tableName, 'ruoyi')
            this.loading = false
          }
        })
      }
    },
    /** 创建菜单操作 */
    handleAddMenu (row) {
      var that = this
      const tableName = row.tableName
      this.$confirm({
        title: '确认自动生成菜单吗?',
        content: '自动生成【' + tableName + '】菜单数据',
        onOk () {
          return addMenu(tableName)
            .then(() => {
              that.onSelectChange([], [])
              that.getList()
              that.$message.success(
                '菜单创建成功',
                3
              )
          })
        },
        onCancel () {}
      })
    },
    handleOk () {
      this.resetQuery()
    }
  }
}
</script>
