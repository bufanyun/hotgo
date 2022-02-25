<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="公告标题">
                <a-input v-model="queryParam.noticeTitle" placeholder="请输入" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="操作人员">
                <a-input v-model="queryParam.createBy" placeholder="请输入" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="公告类型">
                <a-select placeholder="请选择" v-model="queryParam.noticeType" style="width: 100%" allow-clear>
                  <a-select-option v-for="(d, index) in typeOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-select-option>
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
      <notice-view-form
        v-if="showViewModal"
        ref="noticeViewForm"
        :typeOptions="typeOptions"
        :statusOptions="statusOptions"
        @ok="getList"
        @close="showViewModal = false"
      />
      <advance-table
        :columns="columns"
        :data-source="list"
        title="通知公告"
        :loading="loading"
        rowKey="id"
        size="middle"
        tableKey="system-notic-NoticReadIndex-table"
        @refresh="getList"
        :isTableConfig="false"
        :isShowSetBtn="false"
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
          <a-button type="primary" :disabled="multiple" @click="updateNoticeToRead">
            <a-icon type="pushpin" />标记已读
          </a-button>
        </div>
        <span slot="noticeType" slot-scope="{text, record}">
          <a-tag :color="text | noticeTypeFilter">
            {{ typeFormat(record) }}
          </a-tag>
        </span>
        <span slot="isRead" slot-scope="{text, record}">
          <a-badge :status="text === '1' ? 'processing' : 'error'" :text="isReadFormat(record) " />
        </span>
        <span slot="createTime" slot-scope="{text, record}">
          {{ parseTime(record.createTime, '{y}-{m}-{d}') }}
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleView(record)" >
            阅读
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
import { listNoticeByUser, updateNoticeToRead } from '@/api/system/notice'
import AdvanceTable from '@/components/pt/table/AdvanceTable'
import NoticeViewForm from '@/views/system/notice/modules/NoticeViewForm'
export default {
  name: 'Notice',
  components: {
    AdvanceTable,
    NoticeViewForm
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
        pageNum: 1,
        pageSize: 10,
        noticeTitle: undefined,
        createBy: undefined,
        status: undefined
      },
      addModalRefName: 'addModal', // 添加弹窗ref名称
      columns: [
        {
          title: '公告标题',
          dataIndex: 'noticeTitle',
          ellipsis: true,
           width: '20%'
        },
        {
          title: '公告类型',
          dataIndex: 'noticeType',
          scopedSlots: { customRender: 'noticeType' },
          align: 'center'
        },
        {
          title: '是否已读',
          dataIndex: 'isRead',
          scopedSlots: { customRender: 'isRead' },
          align: 'center'
        },
        {
          title: '发布人',
          dataIndex: 'createByName',
          ellipsis: true,
          width: '8%'
        },
        {
          title: '发布',
          dataIndex: 'createTime',
          scopedSlots: { customRender: 'createTime' }
        },
        {
          title: '操作',
          dataIndex: 'operation',
          width: '200',
          scopedSlots: { customRender: 'operation' }
        }
      ],
      showViewModal: false
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
      this.statusOptions = response.data
    })
    this.getDicts('sys_notice_type').then(response => {
      this.typeOptions = response.data
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
      listNoticeByUser(this.queryParam).then(response => {
          this.list = response.data.list
          this.total = response.data.total
          this.loading = false
        }
      )
    },
    isReadFormat (row) {
      if (row.isRead !== null && row.isRead !== '') {
        if (row.isRead === '1') {
          return '已读'
        } else {
          return '未读'
        }
      } else {
        return '未读'
      }
    },
    // 公告类型字典翻译
    typeFormat (row) {
      return this.selectDictLabel(this.typeOptions, row.noticeType)
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
        noticeTitle: undefined,
        createBy: undefined,
        status: undefined
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
    handleView (record) {
      this.showViewModal = true
      this.$nextTick(() => (
         this.$refs.noticeViewForm.handleView(record)
       ))
    },
    updateNoticeToRead (row) {
      var that = this
      const noticeIds = row.id || this.ids
      this.$confirm({
        title: '确认标记所选中数据?',
        onOk () {
          return updateNoticeToRead(noticeIds)
            .then(() => {
              that.onSelectChange([], [])
              that.getList()
              that.$message.success(
                '标记成功',
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
