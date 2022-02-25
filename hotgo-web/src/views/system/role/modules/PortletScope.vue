<template>
  <div>
    <a-card :bordered="false">
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
        :scroll="{y: 'calc(100vh - 360px)' }"
        rowKey="id"
        size="middle"
        @refresh="getList"
        :columns="columns"
        :data-source="sysPortletList"
        :loading="loading"
        :format-conditions="true"
        :isShowSetBtn="false"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }">
        <span slot="menuId" slot-scope="{text, record}">
          <a-popconfirm ok-text="是" cancel-text="否" @confirm="confirmHandleAuth(record)">
            <span slot="title">确认<b>{{ record.menuId !== null ? '取消' : '给' }}</b>该小页授权吗?</span>
            <a-switch :checked="text !== '' && text !== null" >
            </a-switch>
          </a-popconfirm>
        </span>
        <div class="table-operations" slot="button">
          <a-button type="primary" @click="handleBatchScope('add')" v-hasPermi="['system:sysPortlet:add']">
            <a-icon type="plus" />授权
          </a-button>
          <a-button type="danger" @click="handleBatchScope('del')" v-hasPermi="['system:sysPortlet:remove']">
            <a-icon type="delete" />取消授权
          </a-button>
        </div>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
  import { listSysPortletByRoleId } from '@/api/system/sysPortlet'
  import { saveRolePortlet } from '@/api/system/role'
  import AdvanceTable from '@/components/pt/table/AdvanceTable'
  export default {
    name: 'SysPortlet',
    components: {
      AdvanceTable
    },
    data () {
      return {
        selectRoleId: '',
        // 遮罩层
        loading: false,
        // 选中数组
        ids: [],
        addIds: [],
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
        labelCol: {
          span: 6
        },
        // 内容区域的百分比
        wrapperCol: {
          span: 18
        },
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
        columns: [{
            title: '小页名称',
            dataIndex: 'name',
            ellipsis: true,
            width: '30%'
          },
          {
            title: '小页编码',
            dataIndex: 'code',
            ellipsis: true,
            width: '30%'
          },
          {
            title: '是否已授权',
            dataIndex: 'menuId',
            ellipsis: true,
            width: '30%',
            scopedSlots: {
              customRender: 'menuId'
            }
          }
        ]
      }
    },
    created () {},
    methods: {
      resetData () {
        this.ids = []
        this.addIds = []
        // 选中的主键集合
        this.selectedRowKeys = []
        // 选中的数据集合
        this.selectedRows = []
      },
      loadPortlet (roleId) {
        this.selectRoleId = roleId
        this.getList()
      },
      /** 查询工作台小页管理列表 */
      getList () {
        this.resetData()
        this.loading = true
        this.queryParam.roleId = this.selectRoleId
        listSysPortletByRoleId(this.queryParam).then(response => {
          this.sysPortletList = response.data.list
          this.total = response.data.total
          this.loading = false
        })
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
        this.addIds = this.selectedRows.map(function (item) {
         if (item.menuId === null) {
            return item.id
         } else {
           return ''
         }
        })
        this.single = selectedRowKeys.length !== 1
        this.multiple = !selectedRowKeys.length
      },
      handleBatchScope (option) {
        var that = this
        let menuIds = this.ids
        if (option === 'add') {
         menuIds = this.addIds.filter(id => id !== '')
        }
        const saveData = {
          'id': this.selectRoleId,
          'menuIds': menuIds,
          'option': option
        }
        const title = option === 'add' ? '授权' : '取消授权'
        this.$confirm({
          title: '确认为所选中数据' + title + '?',
          onOk () {
            saveRolePortlet(saveData)
              .then(() => {
                that.$message.success(
                  title + '成功',
                  3
                )
                that.getList()
              }).catch(function () {
                that.$message.error(
                 title + '异常',
                  3
                )
              })
          },
          onCancel () {}
        })
      },
      confirmHandleAuth (row) {
        const text = row.menuId !== null ? '取消授权' : '授权'
        const option = row.menuId !== null ? 'del' : 'add'
        const menuIds = [row.id]
        const saveData = {
          'id': this.selectRoleId,
          'menuIds': menuIds,
          'option': option
        }
        saveRolePortlet(saveData)
          .then(() => {
            this.$message.success(
              text + '成功',
              3
            )
            this.getList()
          }).catch(function () {
            this.$message.error(
              text + '异常',
              3
            )
          })
      }
    }
  }
</script>
<style scoped lang="less">
  /deep/.ant-table-body {
    height: calc(100vh - 360px);
  }
</style>
