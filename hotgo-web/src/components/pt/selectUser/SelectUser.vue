<template>
  <div>
    <a-input @click="showSelectUser" v-model="showValue" select-model="single" readOnly>
      <a-icon slot="suffix" type="user-add" />
    </a-input>
    <ant-modal
      :visible="open"
      :modal-title="title"
      :adjust-size="false"
      @cancel="cancel"
      modalHeight="480"
      modalWidth="800"
      dialogClass="personSelect">
      <a-row slot="content">
        <a-spin :spinning="spinning" :delay="delayTime" tip="Loading...">
          <a-col class="treeBox treeborder" :span="12">
            <a-input-search placeholder="请输入用户信息" @search="filterNode" />
            <div class="personSelectTree">
              <a-tree
                v-if="selectModel == 'multi'"
                v-model="checkedKeys"
                checkable
                :replaceFields="replaceFields"
                :default-expanded-keys="expandedKeys"
                :expanded-keys="expandedKeys"
                :auto-expand-parent="autoExpandParent"
                :selected-keys="selectedKeys"
                :tree-data="deptOptions"
                :load-data="onLoadData"
                showIcon
                @expand="onExpand"
                @check="checkNode">
                <a-icon slot="org" type="" :component="allIcon.companyFillIcon" class="depIcon" />
                <a-icon slot="company" type="" :component="allIcon.companyIcon" class="depIcon" />
                <a-icon slot="dept" type="" :component="allIcon.connectionsIcon" class="depIcon" />
                <a-icon slot="user" type="user" class="depIcon" />
                <template slot="title" slot-scope="{ userTitle, attributes }">
                  <span v-if="userTitle.indexOf(searchValue) > -1">
                    {{ userTitle.substr(0, userTitle.indexOf(searchValue)) }}
                    <span style="color: #f50">{{ searchValue }}</span>
                    {{ userTitle.substr(userTitle.indexOf(searchValue) + searchValue.length) }}
                  </span>
                  <span v-else-if="attributes.deptPinyin.indexOf(searchValue) > -1">
                    <span style="color: #f50">{{ userTitle }}</span>
                  </span>
                  <span v-else>{{ userTitle }}</span>
                </template>
              </a-tree>
              <a-tree
                v-if="selectModel == 'single'"
                :replaceFields="replaceFields"
                :default-expanded-keys="expandedKeys"
                :expanded-keys="expandedKeys"
                :auto-expand-parent="autoExpandParent"
                :selected-keys="selectedKeys"
                :tree-data="deptOptions"
                :load-data="onLoadData"
                showIcon
                @expand="onExpand"
                @select="selectNode">
                <a-icon slot="org" type="" :component="allIcon.companyFillIcon" class="depIcon" />
                <a-icon slot="company" type="" :component="allIcon.companyIcon" class="depIcon" />
                <a-icon slot="dept" type="" :component="allIcon.connectionsIcon" class="depIcon" />
                <a-icon slot="user" type="user" class="depIcon" />
                <template slot="title" slot-scope="{ userTitle, attributes }">
                  <span v-if="userTitle.indexOf(searchValue) > -1">
                    {{ userTitle.substr(0, userTitle.indexOf(searchValue)) }}
                    <span style="color: #f50">{{ searchValue }}</span>
                    {{ userTitle.substr(userTitle.indexOf(searchValue) + searchValue.length) }}
                  </span>
                  <span v-else-if="attributes.deptPinyin.indexOf(searchValue) > -1">
                    <span style="color: #f50">{{ userTitle }}</span>
                  </span>
                  <span v-else>{{ userTitle }}</span>
                </template>
              </a-tree>
            </div>
          </a-col>
          <a-col :span="12">
            <div class="contentBox">
              <div :style="{ padding: '10px 20px' }">
                <a-checkbox :indeterminate="indeterminate" @change="onCheckAllChange" :checked="checkAll">
                  已选（{{ selectCount }}）</a-checkbox>
                <a-icon :style="{ float: 'right' }" type="delete" @click="deleteSelectUser" />
              </div>
              <a-checkbox-group v-model="userCheckedList" @change="onChange">
                <a-list item-layout="horizontal" :data-source="userdata" ref="editTable">
                  <a-list-item slot="renderItem" slot-scope="item">
                    <a-checkbox :value="item.id" @change="checkBoxOnChange" />
                    <a-list-item-meta>
                      <template slot="title">
                        <span class="title-name">{{ item.name }}</span>
                        <span class="title-dept">({{ item.subtitle }})</span>
                      </template>
                      <a-avatar slot="avatar" :icon="item.icon" />
                    </a-list-item-meta>
                    <template slot="actions">
                      <a-icon type="drag" class="dragIconClass" v-if="selectModel === 'multi'" />
                      <a-icon type="close-circle" @click="deletSelectUserByParentIds(item.id, item.parentIds)" />
                    </template>
                  </a-list-item>
                </a-list>
              </a-checkbox-group>
            </div>
          </a-col>
        </a-spin>
      </a-row>

      <template slot="footer">
        <a-button @click="cancel">取消</a-button>
        <a-button type="primary" @click="saveSelectUser">保存</a-button>
      </template>
    </ant-modal>
  </div>
</template>
<script>
  import allIcon from '@/core/icons'
  import {
    userSelectTree,
    searchDeptUserList
  } from '@/api/system/dept'
  import {
    getUserInfoByIds
  } from '@/api/system/user'
  import AntModal from '@/components/pt/dialog/AntModal'
  import Sortable from 'sortablejs' // 列交换第三方插件
  export default {
    props: {
      title: {
        type: String,
        default: '人员选择'
      },
      // 默认值
      defaultValue: {
        required: false,
        default: null
      },
      // 返回数据
      value: {
        required: false
      },
      // 单选 single ，多选 multi
      selectModel: {
        type: String,
        required: false,
        default: 'single'
      },
      maxSelect: {
        type: Number,
        required: false,
        default: 0
      }
    },
    data() {
      return {
        showValue: '',
        oldValue: '',
        indeterminate: false,
        checkAll: false,
        userCheckedList: [],
        spinning: false,
        delayTime: 200,
        allIcon,
        sortable: undefined,
        replaceFields: {
          children: 'children',
          key: 'id',
          value: 'id'
        },
        open: false,
        expandedKeys: [],
        autoExpandParent: true,
        checkedKeys: [],
        selectedKeys: [], // 左侧树所有选中节点
        deptOptions: [],
        deptNodes: [],
        oldDeptOptions: [], // 记录查询前数据结构
        oldExpandedKeys: [],
        expandSonData: [], // 异步展开节点时记录子节点
        select: {
          ids: '',
          names: ''
        }, // 最终选择用户对象
        searchValue: '',
        userdata: [],
        checkedUser: [],
        selectCount: 0
      }
    },
    components: {
      AntModal,
      allIcon
    },
    created() {
      this.getTreeselect()
    },
    mounted() {},
    watch: {
      checkedKeys(val) {
        console.log('onCheck', val)
      },
      checkedUser: {
        immediate: true,
        handler(val) {
          this.userCheckedList = val
        }
      },
      userdata: {
        immediate: true,
        handler(val) {
          this.selectCount = val.length
          if (this.selectModel === 'multi' && val.length > 0 && this.sortable === undefined) {
            this.rowDrop()
          }
          /* if (this.userCheckedList.length === val.length) {
            this.checkAll = true
            this.indeterminate = true
          } else {
            this.checkAll = false
            this.indeterminate = false
          } */
        }
      },
      select: {
        immediate: true,
        handler(val) {
          this.oldValue = this.select && this.select.names ? this.select.names : ''
          this.showValue = this.oldValue
        }
      },
      value: {
        immediate: true,
        handler(newV) {
          if (newV) {
            this.select = newV
          } else {
            this.select = {
              ids: '',
              names: ''
            }
          }
        }
      }
    },
    methods: {
      getSelectUserInfo() {
        if (this.select.ids !== undefined && this.select.ids !== 'undefined' && this.select.ids !== '') {
          const userids = {
            userIds: this.select.ids
          }
          getUserInfoByIds(userids).then(response => {
            this.userdata = response
            this.selectedKeys = []
            this.checkedKeys = []
            this.userdata.forEach(node => {
              this.selectedKeys.push(node.id)
              this.checkedKeys.push(node.id)
            })
            // this.selectCount = this.userdata.length
          })
        }
      },
      checkBoxOnChange(e) {
        if (e.target.checked) {
          this.checkedUser.push(e.target.value)
        } else {
          this.checkedUser = this.checkedUser.filter(function(item) {
            return item !== e.target.value
          })
        }
      },
      onChange(checkedList) {
        // 右侧已选用户复选框勾选触发
        /* this.indeterminate = !!checkedList.length && checkedList.length < this.userdata.length
        this.checkAll = checkedList.length === this.userdata.length */
        this.indeterminate = !!this.checkedUser.length && this.checkedUser.length < this.userdata.length
        this.checkAll = this.checkedUser.length === this.userdata.length
      },
      onCheckAllChange(e) {
        // 右侧已选用户全选按钮触发 解决漏洞
        /* const ids = []
        this.userdata.forEach(node => {
          ids.push(node.id)
        }) */
        Object.assign(this, {
          userCheckedList: e.target.checked ? this.selectedKeys : [],
          checkedUser: e.target.checked ? this.selectedKeys : [],
          indeterminate: false,
          checkAll: e.target.checked
        })
      },
      deleteSelectUser() {
        // 右侧已选用户上方删除按钮，全选和勾选删除
        if (this.checkAll) {
          // 全选状态下直接清空数据
          this.userdata.forEach(record => {
            this.checkedKeys = this.checkedKeys.filter(function(item) {
              return record.parentIds.indexOf(item) < 0
            })
          })
          this.userdata = []
          this.userCheckedList = []
          this.checkAll = false
          // this.selectCount = this.userdata.length
        } else {
          // 移除已勾选数据
          this.userCheckedList.forEach(checkItem => {
            this.userdata.some((record, i) => {
              if (record.id === checkItem) {
                this.userdata.splice(i, 1)
                this.checkedKeys = this.checkedKeys.filter(function(item) {
                  return record.parentIds.indexOf(item) < 0
                })
              }
            })
          })
          // this.selectCount = this.userdata.length
        }
      },
      deletSelectUserByParentIds(id, parentIds) {
        // 右侧已选用户悬浮删除按钮删除方法
        this.checkedKeys = this.checkedKeys.filter(function(item) {
          return parentIds.indexOf(item) < 0
        })
        this.userdata.some((record, i) => {
          if (record.id === id) {
            this.userdata.splice(i, 1)
            this.checkedKeys = this.checkedKeys.filter(function(item) {
              return record.parentIds.indexOf(item) < 0
            })
          }
        })
        // this.selectCount = this.userdata.length
      },
      resetSelectUserInfo() {
        this.checkedKeys = []
        this.userdata = []
        // this.selectCount = this.userdata.length
      },
      /** 查询部门下拉树结构 */
      getTreeselect() {
        userSelectTree('0', 3).then(response => {
          this.deptOptions = response
          this.getExpandedKeys(this.deptOptions, 3)
          Object.assign(this, {
            expandedKeys: this.expandedKeys,
            searchValue: '',
            autoExpandParent: true
          })
        })
      },
      getExpandedKeys(nodes, expandLevel) {
        // 递归展开指定层级
        if (expandLevel > 1) {
          // 最后一层不展开
          nodes.forEach(node => {
            this.expandedKeys.push(node.id)
            expandLevel = expandLevel - 1
            return this.getExpandedKeys(node.children, expandLevel)
          })
        }
      },
      getExpandedAllKeys(nodes) {
        // 递归展开所有层
        if (!nodes || nodes.length === 0) {
          return []
        }
        // 最后一层不展开
        nodes.forEach(node => {
          this.deptNodes.push(node.id)
          return this.getExpandedAllKeys(node.children)
        })
      },
      onLoadData(treeNode) {
        // 展开节点时动态加载数据
        return new Promise(resolve => {
          if (treeNode.dataRef.children) {
            resolve()
            return
          }
          this.spinning = !this.spinning
          userSelectTree(treeNode.dataRef.id, 1).then(response => {
            treeNode.dataRef.children = response
            this.expandSonData = response
            if (treeNode.checked) {
              // 当前节点已经是选中状态时异步加载子节点的话，需要将子节点人员选择到已选人员列表
              this.setSelectUserInfoByNodes(response)
            } else {
              this.checkedKeys = this.selectedKeys
            }
            this.spinning = !this.spinning
            resolve()
          })
        })
      },
      showSelectUser() {
        this.getSelectUserInfo()
        this.open = true
        this.resetSelectUserInfo()
      },
      filterNode(value, e) {
        if (this.oldDeptOptions.length === 0) {
          this.oldDeptOptions = this.deptOptions
          this.oldExpandedKeys = this.expandedKeys
        }
        if (value.trim() === '') {
          // 触发父页面设置树数据
          this.deptOptions = this.oldDeptOptions
          Object.assign(this, {
            expandedKeys: this.oldExpandedKeys,
            searchValue: value,
            autoExpandParent: true
          })
        } else {
          const searchInfo = {
            searchText: value
          }
          searchDeptUserList(searchInfo).then(response => {
            // 触发父页面设置树数据
            this.deptOptions = response
            this.getExpandedAllKeys(response)
            Object.assign(this, {
              expandedKeys: this.deptNodes,
              searchValue: value,
              autoExpandParent: true
            })
            this.deptNodes = []
          })
        }
      },
      callback(key) {
        console.log(key)
      },
      cancel(e) {
        this.$emit('close')
        this.open = false
      },
      onExpand(expandedKeys) {
        this.expandedKeys = expandedKeys
        this.autoExpandParent = false
      },
      onCheck(checkedKeys) {
        console.log('onCheck', checkedKeys)
        this.checkedKeys = checkedKeys
      },
      selectNode(selectedKeys, e) {
        // 单选树节点触发
        var nodeData = e.node.dataRef
        const deptType = nodeData.attributes.deptType
        if (deptType === 'user') {
          this.selectedKeys = []
          this.userdata = []
          const id = nodeData.id
          const name = nodeData.title
          const parentIds = nodeData.parentIds
          const subtitle = nodeData.attributes.subtitle
          const selectUser = {
            id: id,
            name: name,
            subtitle: subtitle,
            parentIds: parentIds,
            icon: 'user'
          }
          this.selectedKeys.push(id)
          this.userdata.push(selectUser)
          // this.selectCount = this.userdata.length
        } else {
          this.$message.warning('请选择用户添加')
        }
      },
      checkNode(selectedKeys, e) {
        if (e.checked && !e.node.isLeaf) {
          const children = e.node.dataRef.children
          if (children === null) {
            // 选中非叶子节点时需要加载子节点并选中
            Promise.all([this.onLoadData(e.node)]).then(res => {
              const id = e.node.$options.propsData.dataRef.id
              this.expandedKeys.push(id)
              // 选中子节点
              if (this.expandSonData.length > 0) {
                this.setSelectUserInfoByNodes(this.expandSonData)
                this.expandSonData = []
              }
            })
          } else {
            this.setSelectUserInfo(e.checkedNodes)
          }
        } else if (e.checked && e.node.isLeaf) {
          this.setSelectUserInfo(e.checkedNodes)
        } else {
          // 移除当前选中节点及其子节点数据
          this.removeSelectUserByUserTree(e.node, 'node')
        }
      },
      unique(arr) {
        // 数据去重
        const res = new Map()
        return arr.filter(arr => !res.has(arr.id) && res.set(arr.id, 1))
      },
      removeSelectUserByUserTree(node, dataSource) {
        let id = ''
        let childrens = null
        if (dataSource === 'node') {
          id = node.dataRef.id
          childrens = node.dataRef.children
        } else {
          id = node.id
          childrens = node.children
        }
        this.selectedKeys = this.selectedKeys.filter(function(item) {
          return item !== id
        })
        this.userdata = this.userdata.filter(function(item) {
          return item.id !== id
        })

        if (childrens !== null) {
          childrens.forEach(childrenNode => {
            this.removeSelectUserByUserTree(childrenNode, 'children')
          })
        }
        // this.selectCount = this.userdata.length
      },
      setSelectUserInfo(checkedNodes) {
        // 过滤掉部门数据
        checkedNodes.forEach(node => {
          const name = node.componentOptions.propsData.dataRef.title
          const id = node.componentOptions.propsData.dataRef.id
          const parentIds = node.componentOptions.propsData.dataRef.parentIds
          const deptType = node.componentOptions.propsData.dataRef.attributes.deptType
          const subtitle = node.componentOptions.propsData.dataRef.attributes.subtitle
          this.setSelectEdUserInfo(id, name, subtitle, deptType, parentIds)
        })
      },
      setSelectUserInfoByNodes(checkedNodes) {
        // 专门处理直接通过勾选父节点展开数据后的选人操作
        checkedNodes.forEach(node => {
          const name = node.title
          const id = node.id
          const parentIds = node.parentIds
          const deptType = node.attributes.deptType
          const subtitle = node.attributes.subtitle
          this.setSelectEdUserInfo(id, name, subtitle, deptType, parentIds)
        })
      },
      setSelectEdUserInfo(id, name, subtitle, deptType, parentIds) {
        this.selectedKeys.push(id)
        this.checkedKeys.push(id)
        if (deptType === 'user') {
          const selectUser = {
            id: id,
            name: name,
            subtitle: subtitle,
            parentIds: parentIds,
            icon: 'user'
          }
          this.userdata.push(selectUser)
          this.userdata = this.unique(this.userdata)
          // this.selectCount = this.userdata.length
        }
      },
      saveSelectUser() {
        // 保存选中数据
        let ids = ''
        let names = ''
        if (this.userdata.length > this.maxSelect && this.maxSelect !== 0) {
          this.$message.warning(`已设置最多选择${this.maxSelect}人！`)
          return
        }
        this.userdata.forEach(function(node, index) {
          if (index > 0) {
            ids += ';'
            names += ';'
          }
          ids = ids + node.id
          names = names + node.name
        })
        this.showValue = names
        const result = {
          ids,
          names
        }
        this.$emit('change', result)
        this.$nextTick(() => {
          this.select = result
          // 双向绑定
          this.$emit('input', result)
          this.$emit('callBack', result)
        })
        this.open = false
      },
      /**
       * 行拖拽事件
       */
      rowDrop() {
        const that = this
        this.$nextTick(() => {
          const xGrid = this.$refs.editTable
          const el = xGrid.$el.querySelector('.ant-list-items')
          this.sortable = Sortable.create(el, {
            handle: '.ant-list-item',
            animation: 300,
            delay: 100,
            chosenClass: 'select-list-color', // 被选中项的css 类名
            dragClass: 'drag-list-color', // 正在被拖拽中的css类名
            onEnd: ({
              newIndex,
              oldIndex
            }) => {
              const currRow = that.userdata.splice(oldIndex, 1)[0]
              that.userdata.splice(newIndex, 0, currRow)
              // this.$emit('rowDrop', this.userdata)
            },
            onUpdate(event) {
              const newIndex = event.newIndex
              const oldIndex = event.oldIndex
              const $body = el
              const $tr = $body.children[newIndex]
              const $oldTr = $body.children[oldIndex]
              // 先删除移动的节点
              $body.removeChild($tr)
              // 再插入移动的节点到原有节点，还原了移动的操作
              if (newIndex > oldIndex) {
                $body.insertBefore($tr, $oldTr)
              } else {
                $body.insertBefore($tr, $oldTr.nextSibling)
              }
            }
          })
        })
      }
    }
  }
</script>
<style lang="less">
  .ant-tree-checkbox-disabled {}

  body .ant-tree li .ant-tree-node-content-wrapper .depIcon {
    color: #666666;
    font-size: 20px;
  }

  .personSelect .ant-modal-body {
    padding: 0;
  }

  .personSelectTree {
    height: 325px;
    overflow: auto;
    padding-left: 15px;
  }

  .personSelect {
    .ant-tabs-bar {
      margin: 0;
    }

    .treeborder {
      border-right: 1px solid #e8e8e8;
    }

    .treeBox {
      padding: 0;

      .ant-input-search {
        width: 100%;
        padding: 10px 15px 5px;
      }

      .ant-input-affix-wrapper .ant-input-suffix {
        right: 22px;
      }

      .ant-tree-checkbox {
        padding: 8px 0 0;
      }

      .ant-tree-checkbox-checked::after {
        border: none;
      }

      .ant-tree li .ant-tree-node-content-wrapper {
        width: calc(100% - 40px);
      }
    }

    .contentBox {
      padding: 0;

      .ant-list-items {
        margin: 0 10px;
      }

      .ant-avatar {
        width: 30px;
        height: 30px;
        line-height: 30px;
        background: #47b5e6;
      }

      .ant-checkbox-wrapper {
        font-size: 12px;
      }

      .ant-checkbox-group {
        display: block;
        height: 330px;
        overflow: auto;
      }

      .ant-list-item-meta-avatar {
        margin-right: 0;
      }

      .ant-list-item-meta-title {
        line-height: 30px;
        font-size: 12px;
        margin-bottom: 0px;
      }

      .ant-avatar.ant-avatar-icon {
        margin: 0 10px;
      }

      .ant-list-item-action>li {
        padding: 0 5px;
      }

      .ant-list-split .ant-list-item {
        border-bottom: 0;
        padding: 5px 10px;

        .ant-list-item-action {
          display: none;
          margin-left: 0;
        }

        .title-name {
          color: #323232;
          margin-right: 5px;
        }

        .title-dept {
          color: #a5a5a5;
        }
      }

      .ant-list-item:hover {
        background: #f0f6ff;
        cursor: move;

        .ant-list-item-action {
          display: block;
        }
      }

      .select-list-color {
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
        z-index: 9999;
      }

      .ant-list-item-action .ant-list-item-action-split {
        width: 0;
      }
    }

    .ant-tabs.ant-tabs-card .ant-tabs-card-bar .ant-tabs-tab {
      margin-right: 0px;
      margin-top: 0px;
      height: 40px;
      line-height: 40px;
      border: 0;
      border-right: 1px solid #fff;
      background: #f3f3f3;
      border-radius: 0;
    }

    .ant-tabs.ant-tabs-card .ant-tabs-card-bar .ant-tabs-tab-active {
      background: #fff;
    }

    .ant-tabs.ant-tabs-card .ant-tabs-card-bar .ant-tabs-nav-wrap {
      padding: 0 10px;
      background: #f3f3f3;
    }
  }
</style>
