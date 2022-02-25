<template>
  <div>
    <a-row type="flex" :gutter="10">
      <a-col :span="8">
        <a-card :bordered="false" style="min-height:calc(100vh - 125px);">
          <advance-table
            :columns="columns"
            :data-source="list"
            title="角色管理"
            :loading="loading"
            rowKey="id"
            size="middle"
            :isShowSetBtn="false"
            @refresh="getList"
            :customRow="onClickRow"
            :format-conditions="true"
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
            </div>
          </advance-table>
        </a-card>
      </a-col>
      <a-col :span="16">
        <a-card :bordered="false" style="min-height:calc(100vh - 125px);">
          <a-tabs
            default-active-key="1"
            @change="tabChange"
            v-if="selectItem.id !== undefined && selectItem.id !== 1">
            <a-tab-pane key="1" tab="菜单权限" style="padding-left:20px;padding-right: 20px;">
              <a-spin :spinning="spinning" :delay="delayTime" tip="Loading...">
                <a-card>
                  <a slot="title" >
                    菜单权限：
                    <a-checkbox @change="handleCheckedTreeExpand($event)" :checked="menuExpand">
                      展开/折叠
                    </a-checkbox>
                    <a-checkbox @change="handleCheckedTreeNodeAll($event)" :checked="menuNodeAll">
                      全选/全不选
                    </a-checkbox>
                    <a-checkbox @change="handleCheckedTreeConnect($event)" :checked="form.menuCheckStrictly">
                      父子联动
                    </a-checkbox>
                  </a>
                  <a slot="extra" >
                    <a-button type="primary" @click="saveRoleMenu">
                      保存
                    </a-button>
                  </a>
                  <a-form-model ref="form" :model="form" style="height:calc(100vh - 265px);overflow-y: auto; overflow-x:hidden;padding-left:20px;">
                    <a-form-model-item >
                      <a-tree
                        v-model="menuCheckedKeys"
                        checkable
                        :checkStrictly="!form.menuCheckStrictly"
                        :expanded-keys="menuExpandedKeys"
                        :auto-expand-parent="autoExpandParent"
                        :tree-data="menuOptions"
                        @check="onCheck"
                        @expand="onExpandMenu"
                        :replaceFields="defaultProps" />
                    </a-form-model-item>
                  </a-form-model>
                </a-card>
              </a-spin>
            </a-tab-pane>
            <a-tab-pane key="3" tab="数据权限">
              <div style="height:calc(100vh - 200px);padding-left:20px;padding-right: 20px;">
                <data-scope ref="createDataScopeForm" />
              </div>
            </a-tab-pane>
          </a-tabs>

          <a-result
            style="padding: 50px;"
            v-if="selectItem.id !== undefined && selectItem.id === 1"
            status="success"
            title="超级管理员权限"
            sub-title="超级管理员不受权限控制,其余角色根据需求设置菜单,小页以及数据权限"
          >
          </a-result></a-card></a-col></a-row></div></template>
          </a-result>

        </a-card>
      </a-col>
    </a-row>

  </div>
</template>
<script>
  import {
    listRole,
    updateRole
  } from '@/api/system/role'
  import {
    roleMenuTreeselect
  } from '@/api/system/menu'
  import AdvanceTable from '@/components/pt/table/AdvanceTable'
  import DataScope from './modules/DataScope'
  import PortletScope from './modules/PortletScope'

  export default {
    name: 'Role',
    components: {
      AdvanceTable,
      DataScope,
      PortletScope
    },
    data () {
      return {
        spinning: false,
        delayTime: 200,
        showDataScopeModal: false,
        showPortletModal: false,
        tab1Flag: '', // 选中行是否变化标志
        tab2Flag: '', // 选中行是否变化标志
        tab3Flag: '', // 选中行是否变化标志
        currentSelectTabKey: '1',
        routeSelectRoleId: '', // 从角色管理跳转到集中授权页面时传入的角色ID，用于定位选中行
        list: [],
        // 表格缓存的数据 - 用来点击取消时回显数据
        cacheData: [],
        deleteData: [], // 可编辑表格待删除数据，数据库已存在数据界面假删除，保存到该集合，最终请求数据库删除
        subList: [],
        selectedRowKeys: [],
        selectedRows: [],
        // 高级搜索 展开/关闭
        advanced: false,
        // 非单个禁用
        single: true,
        // 非多个禁用
        multiple: true,
        currentSelectRoleId: '',
        selectItem: {},
        loading: false,
        total: 0,
        // 权限区域
        menuExpandedKeys: [],
        autoExpandParent: false,
        menuCheckedKeys: [],
        halfCheckedKeys: [],
        menuOptions: [],
        menuExpand: false,
        menuNodeAll: false,
        form: {},
        defaultProps: {
          children: 'children',
          title: 'label',
          key: 'id'
        },
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
            ellipsis: true,
            scopedSlots: {
              customRender: 'key'
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
      selectItem (val) {
        this.renderRowStyle(val)
        this.tabChange(this.currentSelectTabKey)
      },
      $route () {
             //  this.routeSelectRoleId = this.$route.params.roleId
             // if(this.routeSelectRoleId !== undefined){

             // }
      }
    },
    methods: {
      statusFormat (row) {
        return this.selectDictLabel(this.statusOptions, row.status)
      },
      tabChange (key) {
        if (this.selectItem.id !== 1) {
          this.currentSelectTabKey = key
          if (key === '1') {
            if ((this.selectItem.id + 1) !== this.tab1Flag) {
              this.getRoleMenuList(this.selectItem)
            }
            this.tab1Flag = this.selectItem.id + 1
          } else if (key === 2) {
            if ((this.selectItem.id + 2) !== this.tab2Flag) {
              this.$nextTick(() => (
                this.$refs.portletRef.loadPortlet(this.selectItem.id)
              ))
            }
            this.tab2Flag = this.selectItem.id + 2
          } else if (key === 3) {
            if ((this.selectItem.id + 3) !== this.tab3Flag) {
              this.$nextTick(() => (
                this.$refs.createDataScopeForm.handleDataScope(this.selectItem)
              ))
            }
            this.tab3Flag = this.selectItem.id + 3
          }
          }
      },
      /** 查询定时任务列表 */
      getList () {
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
          }
        })
      },
      getRoleMenuList (row) {
        this.spinning = true
        this.menuExpand = false
        this.menuNodeAll = false
        const roleId = row.id
        const roleMenu = this.getRoleMenuTreeselect(roleId)
        this.form = row
        roleMenu.then(res => {
          this.menuOptions = res.menus !== null ? res.menus : []
          this.menuCheckedKeys = res.checkedKeys
          // 过滤回显时的半选中node(父)
          if (this.form.menuCheckStrictly) {
            this.selectNodefilter(this.menuOptions, [])
          }
          this.treeExpandWithLevel(this.menuOptions, 1)
          this.spinning = false
        })
      },
      // 回显过滤
      selectNodefilter (nodes, parentIds) {
        if (!nodes || nodes.length === 0) {
          return []
        }
        nodes.forEach(node => {
          // 父子关联模式且当前元素有父级
          const currentIndex = this.menuCheckedKeys.indexOf(node.id)
          // 当前节点存在,且父节点不存在，则说明父节点应是半选中状态
          if (currentIndex !== -1) {
            parentIds.forEach(parentId => {
              if (this.halfCheckedKeys.indexOf(parentId) === -1) {
                this.halfCheckedKeys.push(parentId)
              }
            })
            parentIds = []
          }
          // 防重
          const isExist = this.halfCheckedKeys.indexOf(node.id)
          const isExistParentIds = parentIds.indexOf(node.id)
          if (isExist === -1 && isExistParentIds === -1 && currentIndex === -1) {
            parentIds.push(node.id)
          }
          return this.selectNodefilter(node.children, parentIds)
        })
      },
      treeExpandWithLevel (treeNodeList, level) {
        level--
        if (level !== 0) {
          treeNodeList.forEach(node => {
            this.menuExpandedKeys.push(node.id)
            if (node.children) {
              this.treeExpandWithLevel(node.children, level)
            }
          })
        }
      },
      onExpandMenu (expandedKeys) {
        this.menuExpandedKeys = expandedKeys
        this.autoExpandParent = false
      },
      onCheck (checkedKeys, info) {
        if (!this.form.menuCheckStrictly) {
          let currentCheckedKeys = []
          if (this.menuCheckedKeys.checked) {
            currentCheckedKeys = Array.from(new Set(currentCheckedKeys.concat(this.menuCheckedKeys.checked)))
          }
          if (this.menuCheckedKeys.halfChecked) {
            currentCheckedKeys = Array.from(new Set(currentCheckedKeys.concat(this.menuCheckedKeys.halfChecked)))
          }
          this.menuCheckedKeys = currentCheckedKeys
        } else {
          // 半选节点
          this.halfCheckedKeys = info.halfCheckedKeys
          this.menuCheckedKeys = checkedKeys
        }
      },
      /** 根据角色ID查询菜单树结构 */
      getRoleMenuTreeselect (roleId) {
        return roleMenuTreeselect(roleId).then(response => {
          return response
        })
      },
      onClickRow (record, index) {
        return {
          on: {
            click: (event) => {
              this.selectItem = record
            }
          }
        }
      },
      handleCheckedTreeExpand (value) {
        this.menuExpand = !this.menuExpand
        if (value.target.checked) {
          const treeList = this.menuOptions
          this.treeExpandWithLevel(treeList, -1)
        } else {
          this.menuExpandedKeys = []
          this.treeExpandWithLevel(this.menuOptions, 1)
        }
      },
      handleCheckedTreeNodeAll (value) {
        this.menuNodeAll = !this.menuNodeAll
        if (value.target.checked) {
          this.getAllMenuNode(this.menuOptions)
        } else {
          this.menuCheckedKeys = []
          this.halfCheckedKeys = []
        }
      },
      getAllMenuNode (nodes) {
        if (!nodes || nodes.length === 0) {
          return []
        }
        nodes.forEach(node => {
          this.menuCheckedKeys.push(node.id)
          return this.getAllMenuNode(node.children)
        })
      },
      // 树权限（父子联动）
      handleCheckedTreeConnect (value) {
        this.form.menuCheckStrictly = !this.form.menuCheckStrictly
      },
      /** 搜索按钮操作 */
      handleQuery () {
        this.queryParam.page = 1
        this.getList()
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
      /** 提交按钮 */
      saveRoleMenu: function () {
        this.$refs.form.validate(valid => {
          if (this.form.id !== undefined) {
            this.form.menuIds = this.getMenuAllCheckedKeys()
            updateRole(this.form).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          }
        })
      },
      // 所有菜单节点数据
      getMenuAllCheckedKeys () {
        // 全选与半选
        return Array.from(new Set(this.menuCheckedKeys.concat(this.halfCheckedKeys)))
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
          if (rowList.find(row => row.dataset.rowKey === currentRow.id) !== undefined) {
            rowList.find(row => row.dataset.rowKey === currentRow.id).classList.add('row-selection')
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
