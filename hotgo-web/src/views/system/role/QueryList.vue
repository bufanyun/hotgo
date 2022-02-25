<template>
  <div>
    <a-row type="flex" :gutter="10">
      <a-col :span="12">
        <a-card :bordered="false" style="min-height:calc(100vh - 125px);">
          <advance-table
            :columns="columns"
            :data-source="list"
            title="角色管理-页面"
            :loading="loading"
            rowKey="id"
            size="middle"
            tableKey="system-role-SysRoleIndex-table"
            @refresh="getList"
            :customRow="onClickRow"
            :format-conditions="true"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
            :pagination="{
              current: queryParam.page,
              pageSize: queryParam.limit,
              total: total,
              showSizeChanger: true,
              showLessItems: true,
              showQuickJumper: true,
              showTotal: (total, range) => `第 ${range[0]}-${range[1]} 条，总计 ${total} 条`,
              onChange: changeSize,
              onShowSizeChange: onShowSizeChange,
            }">
            <div class="table-operations" slot="button">
              <a-input-search
                placeholder="请输入编码或名称"
                v-model="queryParam.codeOrName"
                style="width: 150px"
                @search="handleQuery" />
              <a-button type="primary" size="small" @click="addRow()" v-hasPermi="['system:role:add']">
                <a-icon type="plus" />新增
              </a-button>
              <a-button type="" size="small" @click="batchSaveRole()" v-hasPermi="['system:role:add']">
                <a-icon type="save" />保存
              </a-button>
              <a-button type="danger" v-if="!multiple" @click="batchDeleteRow" v-hasPermi="['system:role:remove']">
                <a-icon type="delete" />删除
              </a-button>
              <a-button type="" @click="handleExport" v-hasPermi="['system:role:export']">
                <a-icon type="download" />导出
              </a-button>
            </div>
            <span slot="status" slot-scope="{text,record}">
              <a-tooltip placement="top">
                <template slot="title">
                  <span>点击修改状态</span>
                </template>
                <a-popconfirm
                  ok-text="是"
                  cancel-text="否"
                  @confirm="confirmHandleStatus(record)"
                >
                  <span slot="title">确认<b>{{ record.status === '1' ? '启用' : '停用' }}</b>【{{ record.name }}】的角色吗?</span>
                  <a-badge :status="record.status === '1' ? 'processing' : 'error'" :text=" statusFormat(record) " />
                </a-popconfirm>
              </a-tooltip>
            </span>

            <span slot="name" slot-scope="{text,record}">
              <a-tag color="blue" v-if="!record.editable">
                {{ text }}
              </a-tag>
              <a-input placeholder="请输入" v-model="record.name" v-if="record.editable" />
            </span>
            <span slot="key" slot-scope="{text,record}">
              <a-tag color="blue" v-if="!record.editable || record.key ==='admin'">
                {{ text }}
              </a-tag>
              <a-input placeholder="请输入" v-model="record.key" v-if="record.editable && record.key !=='admin'" />
            </span>
            <span slot="sort" slot-scope="{text,record}">
              <span v-if="!record.editable">
                {{ text }}
              </span>
              <a-input-number placeholder="请输入"v-model="record.sort" v-if="record.editable" :min="0" style="width: 100%"/>
            </span>
            <span slot="operation" slot-scope="{text, record}">
              <div v-if="record.editable">
                <a @click.stop="cancelEditCell(record)" >
                  取消
                </a>
                <a-divider type="vertical" />
                <a @click="deleteRow(record,'delete')">删除</a>
              </div>
              <div v-else>
                <a @click.stop="updateRow(record)" v-hasPermi="['system:role:edit']">
                  修改
                </a>
                <a-divider type="vertical" v-hasPermi="['system:role:remove']" />
                <a @click.stop="deleteRow(record)" v-hasPermi="['system:role:remove']">
                  删除
                </a>
              </div>
            </span>
          </advance-table>
        </a-card>
      </a-col>
      <a-col :span="12">
        <a-card :bordered="false" style="min-height:calc(100vh - 125px);">
          <select-user selectModel="multi" v-model="selectedUser" v-show="false" ref="selectUserRef" />
          <advance-table
            :columns="roleUserColumns"
            :data-source="subList"
            title="角色用户"
            :loading="subLoading"
            rowKey="id"
            size="middle"
            tableKey="system-roleuser-SysRoleIndex-table"
            @refresh="getRoleUserList"
            :format-conditions="true"
            :row-selection="{ selectedRowKeys: selectedSubRowKeys, onChange: onSelectSubChange}"
            :pagination="{
              current: querySubParam.page,
              pageSize: querySubParam.limit,
              total: subTotal,
              showSizeChanger: true,
              showLessItems: true,
              showQuickJumper: true,
              showTotal: (subTotal, range) => `第 ${range[0]}-${range[1]} 条，总计 ${subTotal} 条`,
              onChange: changeSubSize,
              onShowSizeChange: onShowSizeSubChange,
            }">
            <div class="table-operations" slot="button">
              <a-input-search
                placeholder="请输入编码或名称"
                v-model="querySubParam.userNameOrName"
                style="width: 150px"
                @search="handleQueryRoleUser" />
              <a-button
                type="primary"
                v-if="currentSelectRoleId.indexOf('newRow-')<0"
                size="small"
                @click="handleAddUser()"
                v-hasPermi="['system:role:add']">
                <a-icon type="plus" />添加用户
              </a-button>
              <a-button
                type="danger"
                v-if="!subMultiple"
                @click="handleDeleteSub"
                v-hasPermi="['system:role:remove']">
                <a-icon type="delete" />删除
              </a-button>
            </div>

            <span slot="status" slot-scope="{record}">
              <a-badge status="processing" :text=" statusFormat(record) " />
            </span>
            <span slot="key" slot-scope="{text}">
              <a-tag color="blue">
                {{ text }}
              </a-tag>
            </span>
            <span slot="operation" slot-scope="{text, record}">
              <a @click.stop="handleDeleteSub(record)" v-hasPermi="['system:role:remove']">
                删除
              </a>
            </span>
          </advance-table>
        </a-card>
      </a-col>
    </a-row>

  </div>
</template>
<script>
  import {
    listRole,
    exportRole,
    batchSaveRole,
    delRoleUser,
    changeRoleStatus
  } from '@/api/system/role'
  import {
    getRoleUserList,
    saveRoleUser
  } from '@/api/system/user'
  import AdvanceTable from '@/components/pt/table/AdvanceTable'
  import SelectUser from '@/components/pt/selectUser/SelectUser'
  import {
    randomUUID
  } from '@/utils/util'
  export default {
    name: 'Role',
    components: {
      AdvanceTable,
      SelectUser
    },
    data () {
      return {
        list: [],
        // 表格缓存的数据 - 用来点击取消时回显数据
        cacheData: [],
        deleteData: [], // 可编辑表格待删除数据，数据库已存在数据界面假删除，保存到该集合，最终请求数据库删除
        subList: [],
        selectedRowKeys: [],
        selectedSubRowKeys: [],
        selectedRows: [],
        selectedSubRows: [],
        selectedUser: '',
        // 高级搜索 展开/关闭
        advanced: false,
        // 非单个禁用
        single: true,
        // 非多个禁用
        multiple: true,
        subMultiple: true,
        currentSelectRoleId: '',
        selectItem: {},
        ids: [],
        subIds: [], // 子表选择id集合
        loading: false,
        subLoading: false,
        total: 0,
        subTotal: 0,
        // 状态数据字典
        statusOptions: [],
        // 日期范围
        dateRange: [],
        labelCol: {
          span: 6
        },
        wrapperCol: {
          span: 18
        },
        queryParam: {
          page: 1,
          limit: 10,
          name: undefined,
          key: undefined,
          status: undefined
        },
        querySubParam: {
          page: 1,
          limit: 10
        },
        addModalRefName: 'addModal', // 添加弹窗ref名称
        columns: [{
            title: '角色名称',
            dataIndex: 'name',
            ellipsis: true,
            scopedSlots: {
              customRender: 'name'
            }
          },
          {
            title: '角色编码',
            dataIndex: 'key',
            width: '150px',
            ellipsis: true,
            scopedSlots: {
              customRender: 'key'
            }
          },
          {
            title: '排序号',
            dataIndex: 'sort',
            align: 'center',
            scopedSlots: {
              customRender: 'sort'
            }
          },
          {
            title: '状态',
            dataIndex: 'status',
            scopedSlots: {
              customRender: 'status'
            },
            align: 'center'
          },
          {
            title: '操作',
            dataIndex: 'operation',
            width: '100px',
            scopedSlots: {
              customRender: 'operation'
            }
          }
        ],
        roleUserColumns: [{
            title: '用户名称',
            dataIndex: 'realname',
            ellipsis: true
          },
          {
            title: '登录名',
            dataIndex: 'username',
            ellipsis: true
          },
          {
            title: '状态',
            dataIndex: 'status',
            scopedSlots: {
              customRender: 'status'
            },
            align: 'center'
          },
          {
            title: '操作',
            dataIndex: 'operation',
            width: '100px',
            scopedSlots: {
              customRender: 'operation'
            }
          }
        ]
      }
    },
    filters: {},
    created () {
      this.getList()
      this.getDicts('sys_normal_disable').then(response => {
        this.statusOptions = response
      })
    },
    computed: {},
    watch: {
      selectedUser (val) {
        const saveRoleUserForm = {
          'userId': val.ids,
          'role': this.currentSelectRoleId
        }
        saveRoleUser(saveRoleUserForm).then(response => {
          this.$message.success('添加成功', 3)
          this.getRoleUserList()
        })
      },
      selectItem (val) {
        this.renderRowStyle(val)
        this.getRoleUserList(val.id)
      }
    },
    methods: {
      statusFormat (row) {
        return this.selectDictLabel(this.statusOptions, row.status)
      },
      /** 查询定时任务列表 */
      getList () {
        this.showAddModal = false
        this.showEditModal = false
        this.showDataScopeModal = false
        this.loading = true
        listRole(this.addDateRange(this.queryParam, this.dateRange)).then(response => {
          this.list = response.list
          this.list.map((item) => {
            item.operation = item.remark
          })
          this.total = response.total_count
          this.loading = false
          if (this.list.length > 0) {
            this.$nextTick(() => (
              this.selectItem = this.list[0]
            ))
           // this.getRoleUserList(this.list[0].id)
          }
        })
      },
      getRoleUserList (roleId) {
        roleId = roleId.toString()
        if (typeof (roleId) === 'string') {
          if (roleId === '' || roleId === undefined) {
            roleId = this.currentSelectRoleId
          } else {
            this.currentSelectRoleId = roleId
          }
        } else {
          roleId = this.currentSelectRoleId
        }
        if (roleId.indexOf('newRow-') < 0) {
          // 只有保存的数据才加载字表数据
          this.subLoading = true
          this.querySubParam.role = roleId
          getRoleUserList(this.addDateRange(this.querySubParam, this.dateRange)).then(response => {
            this.subList = response.list
            this.subList.map((item) => {
              item.operation = item.remark
            })
            this.subTotal = response.total_count
            this.subLoading = false
          })
        } else {
          // 新添加行不需要加载后台，直接清空子表数据
          this.subList = []
          this.subTotal = 0
        }
        this.subIds = []
        this.selectedSubRowKeys = []
        this.subMultiple = true
      },
      onClickRow (record, index) {
        return {
          on: {
            click: (event) => {
              if (event.target !== undefined && event.target.localName === 'td') {
                if (this.validaData()) {
                  this.selectItem = record
                  // this.updateRow(record)
                }
              }
            }
          }
        }
      },
      /** 搜索按钮操作 */
      handleQuery () {
        this.queryParam.page = 1
        this.getList()
      },
      handleQueryRoleUser () {
        this.querySubParam.page = 1
        this.getRoleUserList()
      },
      addRow () {
        const flag = this.validaData()
        if (flag) {
          // 角色添加行方法
          const id = randomUUID()
          const newRow = {
            id: 'newRow-' + id, // newRow 表示该行是新增的，提交成功后 key替换为数据库ID
            handleType: 'add',
            name: '',
            key: '',
            status: '0',
            operation: '',
            editable: true
          }
          // 关闭所有已打卡行编辑
          this.closeEditCell(this.list)
          this.list.unshift(newRow)
          this.$nextTick(() => (
            this.selectItem = newRow
          ))
        }
      },
      updateRow (record) {
        // if (record.key === 'admin') {
        //   this.$message.info('管理员不允许修改！')
        // } else {
          // 角色编辑行方法
          const newData = [...this.list]
          this.closeEditCell(newData)
          record.handleType = 'edit'
          record.editable = true
          this.list = newData
          // this.getRoleUserList(record.id)
          this.$nextTick(() => (
            this.selectItem = record
          ))
        // }
      },
      cancelEditCell (record) {
        if (this.validaData()) {
          // 取消行编辑
          const newData = [...this.list]
          record.editable = false
          this.list = newData
        }
      },
      realDeleteRow (record, type) {
        if (record.key === 'admin') {
          this.$message.info('管理员不允许删除！')
        } else {
          const newData = [...this.list]
          const id = record.id
          record.handleType = 'delete'
          if (id.indexOf('newRow-') < 0) {
            // 数据库数据需要保存到待删除集合最终提交数据库删除
            this.deleteData.push(record)
          }
          const target = newData.filter(item => record.id !== item.id)
          if (target) {
            this.list = target
          }
          if (type !== 'batch') {
            this.selectFirstRecord()
          }
        }
      },
      deleteRow (record, type) {
        if (type !== 'batch') {
          const id = record.id
          if (id.indexOf('newRow-') >= 0) {
            // 针对前台添加为保存数据库的数据直接删除，无需询问
           this.realDeleteRow(record, type)
          } else {
            const that = this
            this.$confirm({
              title: '删除角色会删除相关的菜单，用户关系等，确认删除吗?',
              onOk () {
                that.realDeleteRow(record, type)
              },
              onCancel () {}
            })
          }
        } else {
          this.realDeleteRow(record, type)
        }
      },
      batchDeleteRow () {
        const that = this
        this.$confirm({
          title: '删除角色会删除相关的菜单，用户关系等，确认删除吗?',
          // content: '当前选中编号为' + roleIds + '的数据',
          onOk () {
            that.ids.forEach(id => {
              const deleteRowData = that.list.filter(item => id === item.id)[0]
              that.deleteRow(deleteRowData, 'batch')
            })
            that.selectFirstRecord()
          },
          onCancel () {}
        })
      },
      selectFirstRecord () {
        // 定位选中行到第一条数据
        if (this.list.length > 0) {
           this.selectItem = this.list[0]
        } else {
          // 移除子表数据
           this.subList = []
           this.subTotal = 0
        }
      },
      batchSaveRole () {
        // 批量保存角色数据
         let saveList = this.list.filter(item => (item.handleType === 'add' || item.handleType === 'edit'))
         saveList = saveList.concat(this.deleteData)
         if (saveList.length === 0) {
             this.$message.info('没有可保存的数据！')
         } else {
           const flag = this.validaData()
           if (flag) {
              batchSaveRole(saveList).then(response => {
                this.$message.success(
                  '保存成功',
                  3
                )
                this.getList()
              })
             }
         }
      },
      validaData () {
        let flag = true
        this.list.forEach(item => {
          if (item.handleType !== 'undefined' && (item.handleType === 'add' || item.handleType === 'edit')) {
            if (item.name === '' || item.name === 'undefined' || item.name === null) {
              this.$message.info('请先维护角色名称为空的数据！')
              flag = false
            }
            if (flag && (item.key === '' || item.key === 'undefined' || item.key === null)) {
              this.$message.info('请先维护角色编码为空的数据！')
              flag = false
            }
          }
        })
        return flag
      },
      closeEditCell (data) {
        // 关闭所有打开的可编辑行
        const list = data.filter(item => item.editable === true)
        list.forEach(item => {
          item.editable = false
        })
      },
      /* 角色状态修改 */
          confirmHandleStatus (row) {
            const text = row.status === '1' ? '启用' : '停用'
            row.status = row.status === '0' ? '1' : '0'
            changeRoleStatus(row.id, row.status)
            .then(() => {
              this.$message.success(
                text + '成功',
                3
              )
            }).catch(function () {
              this.$message.error(
                text + '异常',
                3
              )
            })
          },
      handleAddUser () {
        this.$nextTick(() => (
          this.$refs.selectUserRef.showSelectUser()
        ))
      },
      onShowSizeChange (current, limit) {
        this.queryParam.limit = limit
        this.getList()
      },
      onShowSizeSubChange (current, limit) {
        this.querySubParam.limit = limit
        this.getRoleUserList()
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
      changeSubSize (current, limit) {
        this.querySubParam.page = current
        this.querySubParam.limit = limit
        this.getRoleUserList()
      },
      onSelectChange (selectedRowKeys, selectedRows) {
        this.selectedRowKeys = selectedRowKeys
        this.selectedRows = selectedRows
        this.ids = this.selectedRows.map(item => item.id)
        this.single = selectedRowKeys.length !== 1
        this.multiple = !selectedRowKeys.length
      },
      onSelectSubChange (selectedSubRowKeys, selectedSubRows) {
        this.selectedSubRowKeys = selectedSubRowKeys
        this.selectedSubRows = selectedSubRows
        this.subIds = this.selectedSubRows.map(item => item.id)
        this.subMultiple = !selectedSubRowKeys.length
      },
      handleDeleteSub (row) {
         var that = this
        const userIds = row.id || this.subIds
        this.$confirm({
          title: '确认删除所选中数据?',
          onOk () {
             return delRoleUser(that.currentSelectRoleId, userIds)
              .then(() => {
                that.onSelectSubChange([], [])
                that.getRoleUserList()
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
            return exportRole(that.queryParam)
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
      renderRowStyle (currentRow) { // 增加表格选中行样式
        // 类数组
        const rowEles = document.getElementsByClassName('ant-table-row')
        const rowSelectEles = document.getElementsByClassName('row-selection')
        let rowList
        if (rowSelectEles.length) {
          rowSelectEles[0].classList.remove('row-selection')
        }
        if (rowEles.length) {
          rowList = [...rowEles]
          // 这里不用 === 是因为获取的 rowKey 是 String 类型，而给与的原数据 key 为 Number 类型
          // 若要用 === ，事先进行类型转换再用吧
          if (rowList.find(row => row.dataset.key === currentRow.id) !== undefined) {
            rowList.find(row => row.dataset.key === currentRow.id).classList.add('row-selection')
          }
        }
      }
    }
  }
</script>
<style lang="less">
  .ant-table-row.row-selection {
    background-color: #F0F2F5;
  }
</style>
