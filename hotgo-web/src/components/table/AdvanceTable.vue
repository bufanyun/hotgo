<template>
  <div ref="table" :id="id" class="advanced-table">
    <a-spin :spinning="loading">
      <div :class="['header-bar', size]">
        <div class="title">
          <template v-if="title">{{ title }}</template>
          <slot v-else-if="$slots.title" name="title"></slot>
          <template v-else>高级表格</template>
        </div>
        <div class="button">
          <slot name="button"></slot>
        </div>
        <div class="actions">
          <a-tooltip title="刷新">
            <a-icon @click="refresh" class="action" :type="loading ? 'loading' : 'reload'" />
          </a-tooltip>
          <action-size v-model="sSize" class="action"/>
          <a-tooltip title="列配置">
            <action-columns :columns="columns" @reset="onColumnsReset" class="action">
              <template :slot="slot" v-for="slot in slots">
                <slot :name="slot"></slot>
              </template>
            </action-columns>
          </a-tooltip>
          <a-tooltip title="全屏">
            <a-icon @click="toggleScreen" class="action" :type="fullScreen ? 'fullscreen-exit' : 'fullscreen'" />
          </a-tooltip>
        </div>
      </div>
      <a-table
        v-bind="{...$props, columns: visibleColumns, title: undefined, loading: false}"
        :size="sSize"
        @expandedRowsChange="onExpandedRowsChange"
        @change="onChange"
        @expand="onExpand"
      >
        <template slot-scope="text, record, index" :slot="slot" v-for="slot in scopedSlots ">
          <slot :name="slot" v-bind="{text, record}"></slot>
        </template>
        <!-- <template :slot="slot" v-for="slot in slots">
          <slot :name="slot"></slot>
        </template>
        <template slot-scope="record, index, indent, expanded" :slot="$scopedSlots.expandedRowRender ? 'expandedRowRender' : ''">
          <slot v-bind="{record, index, indent, expanded}" :name="$scopedSlots.expandedRowRender ? 'expandedRowRender' : ''"></slot>
        </template> -->
      </a-table>
    </a-spin>
  </div>
</template>
<script>
  import ActionSize from '@/components/table/ActionSize'
  import ActionColumns from '@/components/table/ActionColumns'
  export default {
    name: 'AdvanceTable',
    components: { ActionColumns, ActionSize },
    props: {
      tableLayout: String,
      bordered: Boolean,
      childrenColumnName: { type: String, default: 'children' },
      columns: Array,
      components: Object,
      dataSource: Array,
      defaultExpandAllRows: Array[String],
      expandedRowKeys: Array[String],
      expandedRowRender: Function,
      expandIcon: Function,
      expandRowByClick: Boolean,
      expandIconColumnIndex: Number,
      footer: Function,
      indentSize: Number,
      loading: Boolean,
      locale: Object,
      pagination: Object,
      rowClassName: Function,
      rowKey: [String, Function],
      rowSelection: Object,
      scroll: Object,
      showHeader: { type: Boolean, default: true },
      size: String,
      title: String,
      customHeaderRow: Function,
      customRow: Function,
      getPopupContainer: Function,
      transformCellText: Function,
      formatConditions: Boolean
    },
    provide () {
      return {
        table: this
      }
    },
    data () {
      return {
        id: `${new Date().getTime()}-${Math.floor(Math.random() * 10)}`,
        sSize: this.size || 'default',
        fullScreen: false,
        conditions: {}
      }
    },
    computed: {
      slots () {
        return Object.keys(this.$slots).filter(slot => slot !== 'title')
      },
      scopedSlots () {
        return Object.keys(this.$scopedSlots).filter(slot => slot !== 'expandedRowRender' && slot !== 'title')
      },
      filterColumns () {
        const filterColumns = []
        const dataIndexColumns = typeof this.dataIndex === 'string' ? [this.dataIndex] : this.dataIndex
        dataIndexColumns.map(item => {
          if (item !== 'title') {
            filterColumns.push(item)
          }
        })
        return filterColumns
      },
      visibleColumns () {
        return this.columns.filter(col => col.visible)
      }
    },
    created () {

    },
    beforeDestroy () {
      this.removeListener()
    },
    methods: {
      refresh () {
        this.$emit('refresh', this.conditions)
      },
      onSearchChange (conditions, searchOptions) {
        this.conditions = conditions
        this.$emit('search', conditions, searchOptions)
      },
      toggleScreen () {
        if (this.fullScreen) {
          this.outFullScreen()
        } else {
          this.inFullScreen()
        }
      },
      inFullScreen () {
        const el = this.$refs.table
        el.classList.add('beauty-scroll')
        if (el.requestFullscreen) {
          el.requestFullscreen()
          return true
        } else if (el.webkitRequestFullScreen) {
          el.webkitRequestFullScreen()
          return true
        } else if (el.mozRequestFullScreen) {
          el.mozRequestFullScreen()
          return true
        } else if (el.msRequestFullscreen) {
          el.msRequestFullscreen()
          return true
        }
        this.$message.warn('对不起，您的浏览器不支持全屏模式')
        el.classList.remove('beauty-scroll')
        return false
      },
      outFullScreen () {
        if (document.exitFullscreen) {
          document.exitFullscreen()
        } else if (document.webkitCancelFullScreen) {
          document.webkitCancelFullScreen()
        } else if (document.mozCancelFullScreen) {
          document.mozCancelFullScreen()
        } else if (document.msExitFullscreen) {
          document.msExitFullscreen()
        }
        this.$refs.table.classList.remove('beauty-scroll')
      },
      onColumnsReset (conditions) {
        this.$emit('reset', conditions)
      },
      onExpandedRowsChange (expandedRows) {
        this.$emit('expandedRowsChange', expandedRows)
      },
      onChange (pagination, filters, sorter, options) {
        this.$emit('change', pagination, filters, sorter, options)
      },
      onExpand (expanded, record) {
        this.$emit('expand', expanded, record)
      },
      addListener () {
        document.addEventListener('fullscreenchange', this.fullScreenListener)
        document.addEventListener('webkitfullscreenchange', this.fullScreenListener)
        document.addEventListener('mozfullscreenchange', this.fullScreenListener)
        document.addEventListener('msfullscreenchange', this.fullScreenListener)
      },
      removeListener () {
        document.removeEventListener('fullscreenchange', this.fullScreenListener)
        document.removeEventListener('webkitfullscreenchange', this.fullScreenListener)
        document.removeEventListener('mozfullscreenchange', this.fullScreenListener)
        document.removeEventListener('msfullscreenchange', this.fullScreenListener)
      },
      fullScreenListener (e) {
        if (e.target.id === this.id) {
          this.fullScreen = !this.fullScreen
        }
      }
    }
  }
</script>
<style scoped lang="less">
.advanced-table{
  overflow-y: auto;
  background-color: #fff;
  .header-bar{
    padding: 16px 24px;
    display: flex;
    align-items: center;
    border-radius: 4px;
    transition: all 0.3s;
    &.middle{
      padding: 12px 16px;
    }
    &.small{
      padding: 8px 12px;
       border: 1px solid #fff;
      border-bottom: 0;
      .title{
        font-size: 16px;
      }
    }
    .title{
      transition: all 0.3s;
      font-size: 18px;
       color: rgba(0, 0, 0, 0.65);
      font-weight: 700;
    }
    .button{
      flex: 1;
      text-align: right;
      margin: 0 24px;
    }
    .actions{
      text-align: right;
      color: rgba(0, 0, 0, 0.65);
      font-size: 17px;
      .action{
        margin: 0 8px;
        cursor: pointer;
        &:hover{
        }
      }
    }
  }
}
</style>
