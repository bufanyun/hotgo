<template>
  <div>
    <a-card :bordered="false" style="margin-bottom: 10px;">
      <div class="table-page-search-wrapper">
        <a-form :labelCol="labelCol" :wrapperCol="wrapperCol">
          <a-row :gutter="48">
            <a-col :md="6" :sm="24">
              <a-form-item label="菜单名称">
                <a-input v-model="queryParam.name" placeholder="请输入菜单名称" allow-clear/>
              </a-form-item>
            </a-col>
            <a-col :md="6" :sm="24">
              <a-form-item label="状态">
                <a-select placeholder="状态" v-model="queryParam.status" style="width: 100%">
                  <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col>
              <span class="table-page-search-submitButtons" style="float: right;">
                <a-button type="primary" @click="handleQuery"><a-icon type="search"/>查询</a-button>
                <a-button style="margin-left: 8px" @click="resetQuery"><a-icon type="redo"/>重置</a-button>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>
    </a-card>
    <a-card :bordered="false" class="table-card">
      <menu-add-form
        v-if="showAddModal"
        ref="menuAddForm"
        :menuOptions="menuOptions"
        :menuTypeOptions="menuTypeOptions"
        :statusOptions="statusOptions"
        :visibleOptions="visibleOptions"
        @ok="getList"
        @select-tree="getTreeselect"
        @close="showAddModal = false"
      />
      <menu-edit-form
        v-if="showEditModal"
        ref="menuEditForm"
        :menuOptions="menuOptions"
        :menuTypeOptions="menuTypeOptions"
        :statusOptions="statusOptions"
        :visibleOptions="visibleOptions"
        @ok="getList"
        @select-tree="getTreeselect"
        @close="showEditModal = false"
      />
      <!-- 数据展示 -->
      <advance-table
        :loading="loading"
        title="菜单管理"
        rowKey="id"
        @refresh="getList"
        :expandIconColumnIndex="1"
        :columns="columns"
        :data-source="list"
        :pagination="false"
        size="middle"
        tableKey="system-menu-MenuIndex-table"
        :defaultExpandedRowKeys="expandedRowKeys"
        :expandedRowKeys="expandedRowKeys"
        :expandIcon="expandIcon"
        @expand="expandNode"
        :indentSize="16"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }">
        <div class="table-operations" slot="button">
          <a-space align="center">
            <a-button type="primary" @click="handleAdd()" v-hasPermi="['system:dept:add']">
              <a-icon type="plus"/>
              新增
            </a-button>
          </a-space>
        </div>
        <span slot="name" slot-scope="{text, record}">
          <a-icon
            :component="allIcon[record.icon + 'Icon']"
            :type="record.icon"
            class="depIcon"
            :style="{ fontSize: '16px',marginTop:'0px' }"/>
          <span v-if="text.indexOf(queryParam.name) > -1">
            {{ text.substr(0, text.indexOf(queryParam.name)) }}
            <span style="color: #f50">{{ queryParam.name }}</span>
            {{ text.substr(text.indexOf(queryParam.name) + queryParam.name.length) }}
          </span>
          <span v-else>{{ text }}</span>
        </span>
        <span slot="sort" slot-scope="{text}">
          <a-tag style="width:50px;">
            {{ text }}
          </a-tag>
        </span>
        <span slot="type" slot-scope="{text, record}">
          <a-tag :color="text | menuTypeFilter">
            {{ menuTypeFormat(record) }}
          </a-tag>
        </span>
        <span slot="status" slot-scope="{text, record}">
          <a-badge :status="record.status == '1' ? 'processing' : 'error'" :text="statusFormat(record) "/>
        </span>
        <span slot="is_visible" slot-scope="{text,record}">
          <span :style="{ color: text === '1' ? 'red' : '#000;' }">{{ visibleFormat(record) }}</span>
        </span>
        <span slot="operation" slot-scope="{text, record}">
          <a @click="handleUpdate(record)" v-hasPermi="['system:dept:edit']">
            修改
          </a>
          <a-divider
            type="vertical"
            v-if="record.status === '1' && record.type !== 'F'"
            v-hasPermi="['system:dept:add']"/>
          <a
            @click="handleAdd(record)"
            v-if="record.status === '1' && record.type !== 'F'"
            v-hasPermi="['system:dept:add']">
            添加子菜单
          </a>
          <a-divider type="vertical" v-hasPermi="['system:dept:remove']"/>
          <a @click="handleDelete(record)" v-hasPermi="['system:dept:remove']">
            删除
          </a>
        </span>
      </advance-table>
    </a-card>
  </div>
</template>
<script>
  import { listMenu, delMenu, searchMenuList } from '@/api/system/menu'
  import MenuEditForm from './modules/MenuEditForm'
  import MenuAddForm from './modules/MenuAddForm'
  import AdvanceTable from '@/components/pt/table/AdvanceTable'
  import allIcon from '@/core/icons'

  export default {
    name: 'Menu',
    components: {
      AdvanceTable,
      MenuEditForm,
      MenuAddForm,
      allIcon
    },
    data() {
      return {
        showAddModal: false,
        showEditModal: false,
        allIcon,
        iconVisible: false,
        list: [],
        // 部门树选项
        selectedRowKeys: [],
        selectedRows: [],
        ids: [],
        menuOptions: [],
        expandedRowKeys: [],
        menuTypeOptions: [],
        loading: false,
        // 状态数据字典
        statusOptions: [],
        visibleOptions: [],
        labelCol: { span: 6 },
        wrapperCol: { span: 18 },
        queryParam: {
          name: undefined,
          visible: undefined
        },
        columns: [
          {
            title: '菜单名称',
            dataIndex: 'name',
            width: '20%',
            scopedSlots: { customRender: 'name' }
          },
          {
            title: '菜单编码',
            dataIndex: 'code',
            width: '150px',
            scopedSlots: { customRender: 'code' }
          },
          {
            title: '排序',
            dataIndex: 'sort',
            align: 'center',
            width: '50px',
            scopedSlots: { customRender: 'sort' }
          },
          {
            title: '菜单类型',
            dataIndex: 'type',
            width: '80px',
            scopedSlots: { customRender: 'type' },
            align: 'center'
          },
          {
            title: '可见',
            dataIndex: 'is_visible',
            width: '50px',
            scopedSlots: { customRender: 'is_visible' },
            align: 'center'
          },
          {
            title: '权限标识',
            dataIndex: 'perms',
            ellipsis: true,
            align: 'center'
          },
          {
            title: '组件路径',
            dataIndex: 'component',
            ellipsis: true,
            scopedSlots: { customRender: 'component' },
            align: 'center'
          },
          {
            title: '状态',
            dataIndex: 'status',
            width: '80px',
            scopedSlots: { customRender: 'status' },
            align: 'center'
          },
          {
            title: '操作',
            dataIndex: 'operation',
            width: '210px',
            scopedSlots: { customRender: 'operation' }
          }
        ]
      }
    },
    filters: {
      menuTypeFilter(type) {
        let value = '#108ee9'
        if (type === 'M') {
          value = '#2db7f5'
        } else if (type === 'C') {
          value = '#87d068'
        }
        return value
      }
    },
    created() {
      this.getList()
      this.getDicts('sys_normal_disable').then(response => {
        this.statusOptions = response
      })
      this.getDicts('sys_show_hide').then(response => {
        this.visibleOptions = response
      })
      this.menuTypeOptions = [{ menuTypeValue: 'M', menuTypeLabel: '目录' }, {
        menuTypeValue: 'C',
        menuTypeLabel: '菜单'
      }, { menuTypeValue: 'F', menuTypeLabel: '按钮' }]
    },
    computed: {},
    watch: {},
    methods: {
      expandNode(expanded, record) {
        // 展开收缩时需要动态修改展开行集合
        if (expanded) {
          this.expandedRowKeys.push(record.id)
        } else {
          this.expandedRowKeys = this.expandedRowKeys.filter(
            function(item) {
              return item !== record.id
            }
          )
        }
        if (expanded && !record.children) { // && record.children.length === 0
          this.loading = true
          listMenu(this.queryParam, record.id, 1).then(response => {
              // record.children = this.handleTree(response.data, 'id')
              record.children = response.list
              this.loading = false
            }
          )
        }
      },
      onSelectChange(selectedRowKeys, selectedRows, expandedRowKeys) {
        this.selectedRowKeys = selectedRowKeys
        this.selectedRows = selectedRows
        // this.expandedRowKeys = expandedRowKeys
        this.ids = this.selectedRows.map(item => item.id)
        this.multiple = !selectedRowKeys.length
      },
      /** 查询定时任务列表 */
      getList() {
        this.expandedRowKeys = []
        this.loading = true
        listMenu(this.queryParam, '', 1).then(response => {
            // console.log("response.data:"+JSON.stringify(response.list))
            this.expandTree(response.list, 1, this.expandedRowKeys)
            // this.list = this.handleTree(response.data, 'id')
            this.list = response.list
            this.loading = false
          }
        )
      },
      // 字典状态字典翻译
      visibleFormat(row) {
        /* if (row.menuType === 'F') {
           return ''
         } */
        return this.selectDictLabel(this.visibleOptions, row.is_visible)
      },
      statusFormat(row) {
        return this.selectDictLabel(this.statusOptions, row.status)
      },
      menuTypeFormat(row) {
        if (row.type === 'M') {
          return '目录'
        } else if (row.type === 'C') {
          return '菜单'
        } else if (row.type === 'F') {
          return '按钮'
        }
      },
      /** 搜索按钮操作 */
      handleQuery() {
        if ((this.queryParam.name === undefined && this.queryParam.status === undefined) || (this.queryParam.name === '' && this.queryParam.status === '')) {
          this.expandedRowKeys = []
          this.getList()
        } else {
          this.loading = true
          searchMenuList(this.queryParam).then(response => {
              this.expandedRowKeys = []
              if (response && response.length !== 0) {
                this.getAllMenuNode(response)
                // this.list = this.handleTree(response.data, 'id')
                this.list = response
              } else {
                this.list = []
              }

              this.loading = false
            }
          )
        }
      },
      handleAdd(record) {
        this.showAddModal = true
        this.$nextTick(() => (
          this.$refs.menuAddForm.handleAdd(record)
        ))
      },
      handleUpdate(record) {
        this.showEditModal = true
        this.$nextTick(() => (
          this.$refs.menuEditForm.handleUpdate(record)
        ))
      },
      getAllMenuNode(nodes) {
        if (!nodes || nodes.length === 0) {
          return []
        }
        nodes.forEach(node => {
          if (node.children && node.children.length !== 0) {
            this.expandedRowKeys.push(node.id)
          }
          return this.getAllMenuNode(node.children)
        })
      },
      /** 重置按钮操作 */
      resetQuery() {
        this.queryParam = {
          deptName: undefined,
          status: undefined
        }
        this.handleQuery()
      },
      /** 查询菜单下拉树结构 */
      getTreeselect() {
        searchMenuList().then(response => {
          this.menuOptions = []
          const menu = { id: 0, name: '主目录', children: [] }
          // menu.children = this.handleTree(response.data, 'id')
          menu.children = response
          console.log('response.list:' + JSON.stringify(response))
          this.menuOptions.push(menu)
        })
      },
      /** 删除按钮操作 */
      handleDelete(row) {
        var that = this
        const menuId = row.id || this.ids
        this.$confirm({
          title: '确认删除所选中数据?',
          content: '当前选中的数据',
          onOk() {
            return delMenu(menuId)
              .then(() => {
                if (row !== null && row.pid === 0) {
                  that.removeTreeNode(that.list, row)
                } else {
                  that.onSelectChange([], [], [])
                  that.getList()
                }
                that.$message.success(
                  '删除成功',
                  3
                )
              })
          },
          onCancel() {
          }
        })
      },
      // expandIcon(props) {
      //   if (props.record.treeLeaf === 'y') {
      //     return '<span style = \'margin-right:22px\' > < /span>'
      //   } else {
      //     if (props.expanded) {
      //       return ('< a style = \'color: \'black\' ,margin-right:0px\' onClick = {(e) =>{props.onExpand(props.record, e)}}><a-icon type = \'caret-down\' / > < /a>')
      //     } else {
      //       return ('< a style = \'color: \'black\' ,margin-right:0px\' onClick = {(e) =>{props.onExpand(props.record, e)}}><a-icon type = \'caret-right\' / > < /a>')
      //     }
      //   }
      // }
      expandIcon (props) {
        if (props.record.treeLeaf === 'y') {
          return <span style="margin-right:22px"></span>
        } else {
          if (props.expanded) {
            return (
              <a style="color: 'black',margin-right:0px"
            onClick={(e) => {
              props.onExpand(props.record, e)
            }}
          >
          <a-icon type="caret-down" />
              </a>
          )
          } else {
            return (
              <a style="color: 'black' ,margin-right:0px"
            onClick={(e) => {
              props.onExpand(props.record, e)
            }}
          >
          <a-icon type="caret-right" />
              </a>
          )
          }
        }
      }
    }
  }
</script>
