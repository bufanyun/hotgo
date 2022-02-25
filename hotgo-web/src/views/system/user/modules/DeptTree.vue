<template>
  <div>
    <a-input-search style="margin-bottom: 8px" placeholder="输入名称回车查询" @search="filterNode" />
    <a-tree
      v-if="deptOptions.length > 0"
      :tree-data="deptOptions"
      :replaceFields="replaceFields"
      :default-expanded-keys="expandedKeys"
      :expanded-keys="expandedKeys"
      :auto-expand-parent="autoExpandParent"
      :load-data="onLoadData"
      showIcon
      @select="handleNodeClick"
      @expand="onExpand"
    >
      <a-icon slot="org" type="" :component="allIcon.companyFillIcon" class="depIcon" />
      <a-icon slot="company" type="" :component="allIcon.companyIcon" class="depIcon" />
      <a-icon slot="dept" type="" :component="allIcon.connectionsIcon" class="depIcon" />
      <template slot="title" slot-scope="{ title,attributes }">
        <span v-if="title.indexOf(searchValue) > -1">
          {{ title.substr(0, title.indexOf(searchValue)) }}
          <span style="color: #f50">{{ searchValue }}</span>
          {{ title.substr(title.indexOf(searchValue) + searchValue.length) }}
        </span>
        <span v-else-if="attributes.deptPinyin.indexOf(searchValue) > -1">
          <span style="color: #f50">{{ title }}</span>
        </span>
        <span v-else>{{ title }}</span>
      </template>
    </a-tree></div>
</template>
<script>
  import allIcon from '@/core/icons'
  import { listDeptTree, searchDept } from '@/api/system/dept'
const getParentKey = (id, tree) => {
  let parentKey
  for (let i = 0; i < tree.length; i++) {
    const node = tree[i]
    if (node.children) {
      if (node.children.some(item => item.id === id)) {
        parentKey = node.id
      } else if (getParentKey(id, node.children)) {
        parentKey = getParentKey(id, node.children)
      }
    }
  }
  return parentKey
}
export default {
  name: 'DeptTree',
  props: {
    deptOptions: {
          type: Array,
          required: true
        },
    defalutExpandedKeys: {
          type: Array
        }
  },
  components: {
    allIcon
  },
  data () {
    return {
      expandedKeys: this.defalutExpandedKeys,
      oldDeptOptions: [], // 记录查询前数据结构
      oldExpandedKeys: [],
      allIcon,
      replaceFields: { children: 'children', key: 'id', value: 'id' },
      deptNodes: [],
      searchValue: '',
      autoExpandParent: true
    }
  },
  filters: {
  },
  created () {
  },
  computed: {
  },
  watch: {
  },
  methods: {
    getAllDeptNode (nodes) {
      if (!nodes || nodes.length === 0) {
        return []
      }
      nodes.forEach(node => {
        this.deptNodes.push({ id: node.id, label: node.label })
        return this.getAllDeptNode(node.children)
      })
    },
    getExpandedKeys (nodes) {
     if (!nodes || nodes.length === 0) {
       return []
     }
        // 最后一层不展开
        nodes.forEach(node => {
        this.deptNodes.push(node.id)
        return this.getExpandedKeys(node.children)
        })
    },
    // 筛选节点
    filterNode (value, e) {
      if (this.oldDeptOptions.length === 0) {
       this.oldDeptOptions = this.deptOptions
       this.oldExpandedKeys = this.expandedKeys
      }
      if (value.trim() === '') {
        // 触发父页面设置树数据
          this.$emit('setDataOptionInfo', this.oldDeptOptions)
          Object.assign(this, {
            expandedKeys: this.oldExpandedKeys,
            searchValue: value,
            autoExpandParent: true
          })
     } else {
       const searchInfo = { deptName: value }
       searchDept(searchInfo).then(response => {
         // 触发父页面设置树数据
          this.$emit('setDataOptionInfo', response.data)
          // this.searchTree(value, response.data)
          this.getExpandedKeys(response.data)
         Object.assign(this, {
           expandedKeys: this.deptNodes,
           searchValue: value,
           autoExpandParent: true
         })
         this.deptNodes = []
       })
     }
    },
    // 节点单击事件,
    handleNodeClick (keys, event) {
      this.$emit('select', event.node)
    },
    onExpand (expandedKeys) {
      this.expandedKeys = expandedKeys
      this.autoExpandParent = false
    },
    onLoadData (treeNode) {
          return new Promise(resolve => {
            if (treeNode.dataRef.children) {
              resolve()
              return
            }
            listDeptTree(treeNode.dataRef.id, 1).then(response => {
               treeNode.dataRef.children = response.data
                resolve()
            })
          })
        },
    searchTree (value, options) {
      this.getAllDeptNode(options)
      console.log('deptNodes', this.deptNodes)
      const gData = options
      const expandedKeys = this.deptNodes
        .map(item => {
          if (item.label.indexOf(value) > -1) {
            return getParentKey(item.id, gData)
          }
          return null
        })
        .filter((item, i, self) => item && self.indexOf(item) === i)
      Object.assign(this, {
        expandedKeys: expandedKeys,
        searchValue: value,
        autoExpandParent: true
      })
      this.deptNodes = []
    }
  }
}
</script>
<style lang="less">
  .depIcon {
    color: #2f54eb;
     font-size: 20px;
  }
</style>
