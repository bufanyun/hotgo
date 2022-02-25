<template>
  <div>
    <a-input
      @click="showSelectDept"
      v-model="showValue"
      select-model="single"
      readOnly>
      <a-icon slot="suffix" type="apartment" />
    </a-input>
    <ant-modal
      :visible="open"
      :modal-title="title"
      :adjust-size="false"
      @cancel="cancel"
      modalHeight="480"
      modalWidth="800"
      dialogClass="deptSelect">
      >
      <a-row slot="content">
        <a-spin :spinning="spinning" :delay="delayTime" tip="Loading...">
          <a-col class="treeBox treeborder" :span="12">
            <a-input-search placeholder="请输入部门信息" @search="filterNode"/>
            <div class="deptSelectTree">
              <a-tree
                v-if="selectModel == 'multi'"
                v-model="checkedKeys"
                checkable
                :replaceFields="replaceFields"
                :default-expanded-keys="expandedKeys"
                :selected-keys="selectedKeys"
                :expanded-keys="expandedKeys"
                :auto-expand-parent="autoExpandParent"
                :tree-data="deptOptions"
                :load-data="onLoadData"
                :checkStrictly="checkStrictly"
                showIcon
                @expand="onExpand"
                @check="checkNode"
              >
                <a-icon slot="org" :component="allIcon.companyFillIcon" class="depIcon" />
                <a-icon slot="company" :component="allIcon.companyIcon" class="depIcon" />
                <a-icon slot="dept" :component="allIcon.connectionsIcon" class="depIcon" />
                <template slot="title" slot-scope="{deptTitle, attributes }">
                  <span v-if="deptTitle.indexOf(searchValue) > -1">
                    {{ deptTitle.substr(0, deptTitle.indexOf(searchValue)) }}
                    <span style="color: #f50">{{ searchValue }}</span>
                    {{ deptTitle.substr(deptTitle.indexOf(searchValue) + searchValue.length) }}
                  </span>
                  <span v-else-if="attributes.deptPinyin.indexOf(searchValue) > -1">
                    <span style="color: #f50">{{ deptTitle }}</span>
                  </span>
                  <span v-else>{{ deptTitle }}</span>
                </template>
              </a-tree>
              <a-tree
                v-if="selectModel == 'single'"
                :replaceFields="replaceFields"
                :default-expanded-keys="expandedKeys"
                :selected-keys="selectedKeys"
                :expanded-keys="expandedKeys"
                :auto-expand-parent="autoExpandParent"
                :tree-data="deptOptions"
                :load-data="onLoadData"
                showIcon
                @expand="onExpand"
                @select="selectNode"
              >
                <a-icon slot="org" :component="allIcon.companyFillIcon" class="depIcon" />
                <a-icon slot="company" :component="allIcon.companyIcon" class="depIcon" />
                <a-icon slot="dept" :component="allIcon.connectionsIcon" class="depIcon" />
                <template slot="title" slot-scope="{ deptTitle, attributes }">
                  <span v-if="deptTitle.indexOf(searchValue) > -1">
                    {{ deptTitle.substr(0, deptTitle.indexOf(searchValue)) }}
                    <span style="color: #f50">{{ searchValue }}</span>
                    {{ deptTitle.substr(deptTitle.indexOf(searchValue) + searchValue.length) }}
                  </span>
                  <span v-else-if="attributes.deptPinyin.indexOf(searchValue) > -1">
                    <span style="color: #f50">{{ deptTitle }}</span>
                  </span>
                  <span v-else>{{ deptTitle }}</span>
                </template>
              </a-tree>
            </div>
          </a-col>
          <a-col :span="12">
            <div class="contentBox">
              <div :style="{ padding: '10px 20px' }">
                <a-checkbox :indeterminate="indeterminate" @change="onCheckAllChange" :checked="checkAll">已选（{{ selectCount }}）</a-checkbox>
                <a-icon :style="{ float: 'right' }" type="delete" @click="deleteSelectDept"/>
              </div>
              <a-checkbox-group v-model="deptCheckedList" @change="onChange">
                <a-list item-layout="horizontal" :data-source="deptdata" ref="editTable">
                  <a-list-item slot="renderItem" slot-scope="item">
                    <a-checkbox :value="item.id" @change="checkBoxOnChange"/>
                    <a-list-item-meta>
                      <template slot="title">
                        <span>{{ item.name }}</span>
                      </template>
                      <a-avatar slot="avatar" :icon="item.icon" />
                    </a-list-item-meta>
                    <template slot="actions">
                      <a-icon type="drag" class="dragIconClass" v-if="selectModel == 'multi'"/>
                      <a-icon type="close-circle" @click="deletSelectDeptBypid(item.id,item.pid)"/>
                    </template>
                  </a-list-item>
                </a-list>
              </a-checkbox-group>
            </div>
          </a-col>
        </a-spin>
      </a-row>
      <template slot="footer">
        <a-button @click="cancel">
          取消
        </a-button>
        <a-button type="primary" @click="saveSelectDept">
          保存
        </a-button>
      </template>
    </ant-modal>
  </div>
</template>
<script>
  import allIcon from '@/core/icons'
  import AntModal from '@/components/pt/dialog/AntModal'
  import { listDeptTree, searchDept, getDeptInfoByIds } from '@/api/system/dept'
  import Sortable from 'sortablejs'// 列交换第三方插件
  export default {
    props: {
       title: { type: String, default: '部门选择' },
       // 单选 single ，多选 multi
       selectModel: {
         type: String,
         required: false,
         default: 'single'
       },
        // 选择范围 all dept
        selectScope: {
          type: String,
          required: false,
          default: 'dept'
        },
       // 返回数据
       value: {
         required: false
       }
    },
    data () {
      return {
        showValue: '',
        oldValue: '',
        indeterminate: false,
        delayTime: 200,
        spinning: false,
        open: false,
        replaceFields: { children: 'children', key: 'id', title: 'title' },
        expandedKeys: [],
        autoExpandParent: true,
        selectedKeys: [],
        deptOptions: [],
        checkedKeys: [],
        checkAll: false,
        deptCheckedList: [],
        selectCount: 0,
        deptdata: [],
        checkStrictly: true,
        oldDeptOptions: [], // 记录查询前数据结构
        oldExpandedKeys: [],
        searchValue: '',
        deptNodes: [],
        checkedDept: [],
        allIcon,
        sortable: undefined,
        select: { ids: '', names: '' } // 最终选择用户对象
      }
    },
    created () {
        this.getTreeselect()
        console.log(this.showValue)
    },
    watch: {
      select: {
        immediate: true,
        handler (val) {
          this.oldValue = this.select && this.select.names ? this.select.names : ''
          this.showValue = this.oldValue
        }
      },
      value: {
        immediate: true,
        handler (newV) {
          if (newV) {
            this.select = newV
          } else {
            this.select = { ids: '', names: '' }
          }
        }
      },
      checkedDept: {
        immediate: true,
        handler (val) {
          this.deptCheckedList = val
        }
      },
      deptdata: {
        immediate: true,
        deep: true,
        handler (val) {
          this.selectCount = val.length
          this.onChange()
          if (this.selectModel === 'multi' && val.length > 0 && this.sortable === undefined) {
            this.rowDrop()
          }
        }
      }
    },
    components: {
      AntModal,
      allIcon
    },
    methods: {
      checkBoxOnChange (e) {
        if (e.target.checked) {
          this.checkedDept.push(e.target.value)
        } else {
          this.checkedDept = this.checkedDept.filter(function (item) {
            return item !== e.target.value
          })
        }
      },
      filterNode (value, e) {
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
          const searchInfo = { deptName: value }
          searchDept(searchInfo).then(response => {
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
      callback (key) {
        console.log(key)
      },
      getExpandedAllKeys (nodes) {
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
      showSelectDept () {
        this.getSelectDeptInfo()
        this.open = true
        this.resetSelectDeptInfo()
      },
      resetSelectDeptInfo () {
        this.checkedKeys = []
        this.deptdata = []
      },
      getSelectDeptInfo () {
        if (this.select.ids !== undefined && this.select.ids !== 'undefined' && this.select.ids !== '') {
          const deptids = { deptIds: this.select.ids }
         getDeptInfoByIds(deptids).then(response => {
             this.deptdata = response
             this.selectedKeys = []
             this.checkedKeys = []
             this.deptdata.forEach(node => {
                this.selectedKeys.push(node.id)
                this.checkedKeys.push(node.id)
             })
          })
       }
      },
      cancel (e) {
        this.open = false
      },
      onChange (checkedList) {
        // 右侧已选用户复选框勾选触发
        this.indeterminate = !!this.checkedDept.length && this.checkedDept.length < this.deptdata.length
        this.checkAll = !!this.deptdata.length && this.checkedDept.length === this.deptdata.length
      },
      /** 查询部门下拉树结构 */
      getTreeselect () {
         listDeptTree('0', 3).then(response => {
          this.deptOptions = response
          this.getExpandedKeys(this.deptOptions, 3)
            Object.assign(this, {
              expandedKeys: this.expandedKeys,
              searchValue: '',
              autoExpandParent: true
            })
        })
      },
      getExpandedKeys (nodes, expandLevel) {
        // 递归展开指定层级
        if (expandLevel > 1) {
          if (nodes !== undefined) {
            // 最后一层不展开
            nodes.forEach(node => {
              this.expandedKeys.push(node.id)
              expandLevel = expandLevel - 1
              return this.getExpandedKeys(node.children, expandLevel)
            })
          }
        }
      },
      onLoadData (treeNode) {
        const that = this
        // 展开节点时动态加载数据
        return new Promise(resolve => {
          if (treeNode.dataRef.children) {
            resolve()
            return
          }
          this.spinning = !this.spinning
          listDeptTree(treeNode.dataRef.id, 1).then(response => {
             treeNode.dataRef.children = response
             that.expandSonData = response
             that.spinning = !that.spinning
              resolve()
          })
        })
      },
      onExpand (expandedKeys) {
        this.expandedKeys = expandedKeys
        this.autoExpandParent = false
      },
      checkNode (selectedKeys, e) {
        if (e.checked && !e.node.isLeaf) {
          const children = e.node.dataRef.children
          this.setSelectDeptInfoByNodes([e.node.$options.propsData.dataRef])
          if (children === null) {
            // 选中非叶子节点时需要加载子节点并选中
            Promise.all([this.onLoadData(e.node)]).then(res => {
              const id = e.node.$options.propsData.dataRef.id
              this.expandedKeys.push(id)
              // 选中子节点
              if (this.expandSonData.length > 0) {
                this.setSelectDeptInfoByNodes(this.expandSonData)
                this.expandSonData = []
              }
            })
          } else {
            // children.push(e.node.$options.propsData.dataRef)
            this.setSelectDeptInfoByNodes(children)
          }
        } else if (e.checked && e.node.isLeaf) {
          this.setSelectDeptInfoByNodes([e.node.$options.propsData.dataRef])
        } else {
          // 移除当前选中节点及其子节点数据
          this.removeSelectDeptByDeptTree(e.node, 'node')
        }
      },
      removeSelectDeptByDeptTree (node, dataSource) {
        let id = ''
        let childrens = null
        if (dataSource === 'node') {
          id = node.dataRef.id
          childrens = node.dataRef.children
        } else {
          id = node.id
          childrens = node.children
        }
        this.selectedKeys = this.selectedKeys.filter(function (item) {
          return item !== id
        })
        this.checkedKeys.checked = this.checkedKeys.checked.filter(function (item) {
          return item !== id
        })
        this.deptdata = this.deptdata.filter(function (item) {
          return item.id !== id
        })
        this.checkedDept = this.checkedDept.filter(function (item) {
          return item !== id
        })
        if (childrens !== null) {
          childrens.forEach(childrenNode => {
            this.removeSelectDeptByDeptTree(childrenNode, 'children')
          })
        }
        // this.selectCount = this.userdata.length
      },
      setSelectDeptInfo (checkedNodes) {
        // 过滤掉部门数据
        checkedNodes.forEach(node => {
          const name = node.componentOptions.propsData.dataRef.title
          const id = node.componentOptions.propsData.dataRef.id
          const pid = node.componentOptions.propsData.dataRef.pid
          const deptType = node.componentOptions.propsData.dataRef.attributes.deptType
          this.setSelectedDeptInfo(id, name, pid, deptType)
        })
      },
      setSelectDeptInfoByNodes (checkedNodes) {
        // 专门处理直接通过勾选父节点展开数据后的选人操作
        checkedNodes.forEach(node => {
          console.log('node:' + JSON.stringify(node))
          const name = node.title
          const id = node.id
          const pid = node.pid
          const deptType = node.attributes.deptType
          this.setSelectedDeptInfo(id, name, pid, deptType)
        })
      },
      getAllSubChildren (children, childIds) {
        children.forEach(node => {
                  childIds.push(node.id)
                  if (node.children != null) {
                     this.getAllSubChildren(node.children, childIds)
                  }
        })
        return false
      },
      selectNode (selectedKeys, e) {
        // 单选树节点触发
        var nodeData = e.node.dataRef
        console.log('nodeData:' + JSON.stringify(nodeData))
        const deptType = nodeData.type
        if (deptType !== 'dept' && this.selectScope === 'dept') {
          this.$message.warning('请选择部门添加')
        } else if (deptType === 'dept' && this.selectScope === 'nonDept') {
          this.$message.warning('请选择机构或公司添加')
        } else {
          this.selectedKeys = []
          this.deptdata = []
          this.checkAll = false
          this.deptCheckedList = []
          const id = nodeData.id
          const name = nodeData.title
          const pid = nodeData.pid
          this.setSelectedDeptInfo(id, name, pid, deptType)
        }
      },
      setSelectedDeptInfo (id, name, pid, deptType) {
        this.selectedKeys.push(id)
        if (this.checkedKeys.checked) {
          this.checkedKeys.checked.push(id)
        } else {
          this.checkedKeys.push(id)
        }
        const selectDept = {
            id: id,
            name: name,
            type: deptType,
            pid: pid,
            icon: 'apartment'
            }
          this.deptdata.push(selectDept)
          this.deptdata = this.unique(this.deptdata)
      },
      unique (arr) {
        const res = new Map()
        return arr.filter((arr) => !res.has(arr.id) && res.set(arr.id, 1))
      },
      onCheckAllChange (e) {
        Object.assign(this, {
          deptCheckedList: e.target.checked ? this.selectedKeys : [],
          checkedDept: e.target.checked ? this.selectedKeys : [],
          indeterminate: false,
          checkAll: e.target.checked
        })
      },
      deleteSelectDept () {
        const that = this
        // 右侧已选用户上方删除按钮，全选和勾选删除
        if (this.checkAll) {
          // 全选状态下直接清空数据
          this.checkedKeys = []
          this.deptdata = []
          this.checkedDept = []
          this.selectedKeys = []
          this.checkAll = false
        } else {
          if (this.checkedKeys.checked !== undefined) {
            this.checkedKeys.checked = this.checkedKeys.checked.filter(function (item) {
              return that.deptCheckedList.indexOf(item) < 0
            })
          } else {
            this.checkedKeys = this.checkedKeys.filter(function (item) {
              return that.deptCheckedList.indexOf(item) < 0
            })
          }

          this.selectedKeys = this.selectedKeys.filter(function (item) {
            return that.deptCheckedList.indexOf(item) < 0
          })
          // 移除已勾选数据
          this.deptCheckedList.forEach(checkItem => {
            this.deptdata.some((record, i) => {
              if (record.id === checkItem) {
                this.deptdata.splice(i, 1)
              }
            })
          })
        }
        this.deptCheckedList = []
        this.checkedDept = []
        this.checkAll = false
      },
      deletSelectDeptBypid (id, pid) {
        // 右侧已选用户悬浮删除按钮删除方法
        if (this.checkedKeys.checked !== undefined) {
          this.checkedKeys.checked = this.checkedKeys.checked.filter(function (item) {
            return id !== item
          })
        } else {
          this.checkedKeys = this.checkedKeys.filter(function (item) {
            return id !== item
          })
        }

        this.deptCheckedList = this.deptCheckedList.filter(function (item) {
          return id !== item
        })
        this.checkedDept = this.checkedDept.filter(function (item) {
          return id !== item
        })
        this.selectedKeys = this.selectedKeys.filter(function (item) {
          return id !== item
        })
        this.deptdata.some((record, i) => {
          if (record.id === id) {
            this.deptdata.splice(i, 1)
          }
        })
        /* this.onChange(this.deptCheckedList)
        if (this.deptCheckedList.length === 0) {
          this.checkAll = false
        } */
      },
      saveSelectDept () {
        // 保存选中数据
       let ids = ''
       let names = ''
       let types = ''
        this.deptdata.forEach(function (node, index) {
             ids = ids + ';' + node.id
             names = names + ';' + node.name
            types = types + ';' + node.type
        })
        if (ids.length > 0) {
          ids = ids.substr(1, ids.length)
          names = names.substr(1, names.length)
          types = types.substr(1, types.length)
        }
        this.showValue = names
        const result = { ids, names, types }
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
      rowDrop () {
        this.$nextTick(() => {
          const xGrid = this.$refs.editTable
          this.sortable = Sortable.create(
            xGrid.$el.querySelector('.ant-list-items'),
            {
              handle: '.ant-list-item',
              animation: 300,
              delay: 100,
              chosenClass: 'drag-list-color', // 被选中项的css 类名
              dragClass: 'drag-list-color', // 正在被拖拽中的css类名
              onEnd: ({ newIndex, oldIndex }) => {
                const currRow = this.deptdata.splice(oldIndex, 1)[0]
                this.deptdata.splice(newIndex, 0, currRow)
                this.$emit('rowDrop', this.deptdata)
              },
              onUpdate (event) {
                const newIndex = event.newIndex
                const oldIndex = event.oldIndex
                const $body = xGrid.$el.querySelector('.ant-list-items')
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
            }
          )
        })
      }
    }
  }
</script>

<style lang="less" >
body .ant-tree li .ant-tree-node-content-wrapper .depIcon {
  color: #666666;
  font-size: 20px;
}
.deptSelect .ant-modal-body {
  padding: 0;
}
.deptSelectTree {
  height: 325px;
  overflow: auto;
  padding-left: 15px;
}
.deptSelect {
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
  }
  .contentBox {
    padding: 0;
    .ant-list-items{
      margin:0 10px;
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
    .ant-checkbox-group{
      display: block;
      height:330px;
      overflow: auto;
    }
    .ant-list-item-meta-avatar {
      margin-right: 0;
    }
    .ant-list-item-meta-title {
      line-height: 30px;
      color:#323232;
      font-size: 12px;
      margin-bottom: 0px;
    }
    .ant-avatar.ant-avatar-icon {
      margin: 0 10px;
    }
    .ant-list-item-action > li {
      padding: 0 5px;
    }
    .ant-list-split .ant-list-item {
      border-bottom: 0;
      padding: 5px 5px 5px 10px;
      .ant-list-item-action {
        display: none;
      }
    }
    .ant-list-item:hover {
      background: #f0f6ff;
      .ant-list-item-action {
        display: block;
      }
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
