<template>
  <div class="action-columns Selected-items" ref="root">
    <a-popover
      v-model="visible"
      placement="bottomRight"
      trigger="click"
      :arrowPointAtCenter="true"
      @click="doParentSetHeight()"
      :get-popup-container="() => $refs.root">
      <div slot="title">
        <a-checkbox :indeterminate="indeterminate" :checked="checkAll" @change="onCheckAllChange" class="check-all" />列展示
        <a-button @click="resetColumns" style="float: right" type="link" size="small">重置</a-button>
      </div>
      <div :style="{maxHeight:columnHeight+'px'}" style="overflow: auto;" slot="content" ref="editTable">
        <a-list size="small" :key="i" v-for="(col, i) in columns">
          <a-list-item>
            <a-icon style="margin-right: 3px;" type="" :component="allIcon.DragColumnIcon" />
            <a-checkbox v-model="col.visible" @change="e => onCheckChange(e, col)" />
            <template v-if="col.title">
              {{ col.title }}
            </template>
            <slot v-else-if="col.slots && col.slots.title" :name="col.slots.title"></slot>
          </a-list-item>
        </a-list>
      </div>
      <a-icon class="action" type="control" />
    </a-popover>
  </div>
</template>
<script>
  import cloneDeep from 'lodash.clonedeep'
  import Sortable from 'sortablejs' // 列交换第三方插件
  import allIcon from '@/core/icons'
  export default {
    name: 'ActionColumns',
    props: ['columns', 'visibleColumns'],
    components: {
      allIcon
    },
    data () {
      return {
        visible: false,
        indeterminate: false,
        checkAll: true,
        columnHeight: 50,
        allIcon,
        sortable: undefined,
        checkedCounts: this.columns.length,
        backColumns: cloneDeep(this.columns)
      }
    },
    watch: {
      checkedCounts (val) {
        this.checkAll = val === this.columns.length
        this.indeterminate = val > 0 && val < this.columns.length
      },
      columns (newVal, oldVal) {
        if (newVal !== oldVal) {
          this.checkedCounts = newVal.length
          this.formatColumns(newVal)
        }
      }
    },
    created () {
      this.formatColumns(this.columns)
    },
    mounted () {},
    methods: {
      doParentSetHeight () {
        this.$emit('setColumnHeight')
        this.rowDrop()
      },
      setHeight (height) {
        let columnsHeight = (height - 160)
        if (columnsHeight < 0) {
          columnsHeight = 35
        }
        this.columnHeight = columnsHeight
      },
      onCheckChange (e, col) {
        if (!col.visible) {
          this.checkedCounts -= 1
        } else {
          this.checkedCounts += 1
        }
        this.$emit('changeColumns')
      },
      fixColumn (fixed, col) {
        if (fixed !== col.fixed) {
          this.$set(col, 'fixed', fixed)
        } else {
          this.$set(col, 'fixed', undefined)
        }
      },
      setSearch (col) {
        this.$set(col, 'searchAble', !col.searchAble)
        if (!col.searchAble && col.search) {
          this.resetSearch(col)
        }
      },
      resetSearch (col) {
        // col.search.value = col.dataType === 'boolean' ? false : undefined
        col.search.value = undefined
        col.search.backup = undefined
      },
      resetColumns () {
        // const { columns, backColumns } = this
        // let counts = columns.length
        // backColumns.forEach((back, index) => {
        //   const column = columns[index]
        //   column.visible = back.visible === undefined || back.visible
        //   if (!column.visible) {
        //     counts -= 1
        //   }
        //   if (back.fixed !== undefined) {
        //     column.fixed = back.fixed
        //   } else {
        //     this.$set(column, 'fixed', undefined)
        //   }
        //    this.$set(column, 'searchAble', back.searchAble)
        //    column.searchAble = back.searchAble
        //    this.resetSearch(column)
        // })
        // this.checkedCounts = counts
        // this.visible = false
        // this.$emit('reset', this.getConditions(columns))
        this.checkedCounts = this.columns.length
        // this.columns.forEach(col => col.visible = true)
        this.columns.forEach(function (col, index) {
          col.visible = true
        })
        this.$emit('changeColumns', 'reset')
      },
      onCheckAllChange (e) {
        if (e.target.checked) {
          this.checkedCounts = this.columns.length
          // this.columns.forEach(col => col.visible = true)
          this.columns.forEach(function (col, index) {
            col.visible = true
          })
        } else {
          this.checkedCounts = 0
          // this.columns.forEach(col => col.visible = false)
          this.columns.forEach(function (col, index) {
            col.visible = false
          })
        }
        this.$emit('changeColumns')
      },
      getConditions (columns) {
        const conditions = {}
        columns.filter(item => item.search.value !== undefined && item.search.value !== '' && item.search.value !==
            null)
          .forEach(col => {
            conditions[col.dataIndex] = col.search.value
          })
        return conditions
      },
      formatColumns (columns) {
        for (const col of columns) {
          if (col.visible === undefined) {
            this.$set(col, 'visible', true)
          }
          if (!col.visible) {
            this.checkedCounts -= 1
          }
        }
      },
      /**
       * 行拖拽事件
       */
      rowDrop () {
        const that = this
        this.$nextTick(() => {
          const xGrid = this.$refs.editTable
          const el = xGrid
          this.sortable = Sortable.create(
            el, {
              handle: '.ant-list-item',
              chosenClass: 'ant-list-item-drag', // 被选中项的css 类名
              dragClass: 'ant-list-item-drag', // 正在被拖拽中的css类名
              onEnd: ({
                newIndex,
                oldIndex
              }) => {
                const currRow = that.columns.splice(oldIndex, 1)[0]
                that.columns.splice(newIndex, 0, currRow)
                that.$emit('rowDrop', that.columns)
                that.$emit('changeColumns')
              },
              onUpdate (event) {
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
            }
          )
        })
      }
    }
  }
</script>

<style scoped lang="less">
  .action-columns {
    display: inline-block;

    .check-all {
      margin-right: 8px;
    }

    .left,
    .right {
      transform: rotate(-90deg);
    }

  }
</style>
<style lang="less">
  /* 表格设置调整表头顺序部分 */
   .Selected-items{
     .ant-popover-title{
      padding: 5px 15px 5px 23px;
      .ant-btn{
        text-align: right;
        margin-top: 2px;
      }
     }
     .ant-popover-inner-content {
        padding: 0px ;
        .ant-list-item{
          padding: 8px 6px;
           width: 100%;
           .ant-checkbox-wrapper{
             margin-right: 4px;
           }
        }
        .ant-list-item:hover{
          cursor: move;
        }
        .ant-list-item-drag{
          background: #ffffff;
          box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
          box-sizing:border-box;
          z-index: 9999;
          border: 1px solid #dee0e3;
        }
      }
   }
</style>
<style lang="less" scoped>
 /* /deep/.ant-popover-placement-bottomRight{
  left: -80px!important;
 } */
</style>
