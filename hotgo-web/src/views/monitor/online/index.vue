<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <!-- 条件搜索 -->
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="登录名称">
                <a-input v-model="queryParam.userName" placeholder="请输入用户名称" allow-clear @keyup.enter.native="handleQuery"/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="登录地址">
                <a-input v-model="queryParam.ipaddr" placeholder="请输入登录地址" allow-clear @keyup.enter.native="handleQuery"/>
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
      <advance-table
        :columns="columns"
        :data-source="list"
        title="在线用户"
        :loading="loading"
        rowKey="tokenId"
        tableKey="monitor-online-index-table"
        :isTableConfig="false"
        :isShowSetBtn="false"
        @refresh="getList"
        size="middle"
        :format-conditions="true"
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
        <span slot="operation" slot-scope="{text, record}">
          <a-popconfirm
            ok-text="是"
            cancel-text="否"
            @confirm="confirmHandleForceLogout(record)"
            @cancel="cancelHandleForceLogout(record)"
          >
            <span slot="title">确认强退<b>{{ record.userName }}</b>的用户吗?</span>
            <a style="color:#2f54eb" v-hasPermi="['monitor:online:forceLogout']"> 强退 </a>
          </a-popconfirm>
        </span>
      </advance-table>

    </a-card>
  </div>
</template>

<script>

import { list, forceLogout } from '@/api/monitor/online'
import AdvanceTable from '@/components/pt/table/AdvanceTable'

export default {
  name: 'Online',
  components: {
    AdvanceTable
  },
  data () {
    return {
      list: [],
      loading: false,
      labelCol: { span: 6 },
      wrapperCol: { span: 18 },
      total: 0,
      // 非多个禁用
      multiple: true,
      queryParam: {
        pageNum: 1,
        pageSize: 10,
        ipaddr: undefined,
        userName: undefined
      },
      columns: [
        {
          title: '会话编号',
          dataIndex: 'tokenId',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '登录名称',
          dataIndex: 'userName',
          align: 'center'
        },
        {
          title: '部门名称',
          dataIndex: 'deptName',
          align: 'center'
        },
        {
          title: '登录地址',
          dataIndex: 'ipaddr',
          ellipsis: true,
          align: 'center'
        },
        {
          title: '登录地点',
          dataIndex: 'loginLocation',
          align: 'center'
        },
        {
          title: '浏览器',
          dataIndex: 'browser',
          align: 'center'
        },
        {
          title: '操作系统',
          dataIndex: 'os',
          align: 'center'
        },
        {
          title: '登录时间',
          dataIndex: 'loginTime',
          width: 180,
          scopedSlots: { customRender: 'loginTime' },
          align: 'center'
        },
        {
          title: '操作',
          dataIndex: 'operation',
          scopedSlots: { customRender: 'operation' }
        }
      ]
    }
  },
  filters: {
  },
  created () {
    this.getList()
  },
  computed: {
  },
  watch: {
  },
  methods: {
    /** 查询登录日志列表 */
    getList () {
      this.loading = true
      list(this.queryParam).then(response => {
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
        ipaddr: undefined,
        userName: undefined
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
    /** 强退按钮操作 */
    confirmHandleForceLogout (row) {
      forceLogout(row.tokenId)
      .then(() => {
        this.getList()
        this.$message.success(
          '已强退',
          3
        )
      }).catch(function () {
        this.$message.error(
          '发生异常',
          3
        )
      })
    },
    cancelHandleForceLogout (row) {
    }
  }
}
</script>
