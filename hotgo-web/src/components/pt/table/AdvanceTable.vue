<template>
  <div
    ref="table"
    :id="id"
    class="advanced-table"
    :style="{'min-height': minHeight} "
  >
    <a-spin :spinning="loading">
      <div :class="['header-bar', size]">
        <div class="title">
          <template v-if="title">
            {{ title }}
          </template>
          <slot v-else-if="$slots.title" name="title"></slot>
          <template v-else>
            高级表格
          </template>
        </div>
        <div class="button"><slot name="button"></slot></div>

        <div
          class="topSet"
          v-if="setingFlag === 'top' && isShowSetBtn"
        >
          <a-tooltip title="刷新" v-if="isShowRefresh"><a-icon @click="refresh" class="action" :type="loading ? 'loading' : 'reload'" /></a-tooltip>
          <action-size v-model="sSize" @input="saveSysTableConfig" class="action" v-if="isShowSize" />
          <a-tooltip title="列配置" v-show="isShowColumns">
            <action-columns
              :columns="myColumns"
              @reset="onColumnsReset"
              @setColumnHeight="setColumnHeight()"
              @changeColumns="saveSysTableConfig"
              ref="columnHeightRef"
              class="action"
            >
              <template :slot="slot" v-for="slot in slots">
                <slot :name="slot"></slot>
              </template>
            </action-columns>
          </a-tooltip>
          <a-tooltip title="全屏" v-if="isShowFull"><a-icon @click="toggleScreen" class="action" :type="fullScreen ? 'fullscreen-exit' : 'fullscreen'" /></a-tooltip>
        </div>

      </div>

      <div class="table-box" style="position: relative;">
        <div v-if="setingFlag === 'table'">
          <transition name="slide-fade">
            <div
              v-show="isShowSetting"
              class="actions tableSet"
              :style="{ background: sSize=== 'small' ? '#fff' : '#fafafa',
                        top: settingTop ,
                        right : isShowSetting ? '40px' : '16px'} "
            >
              <action-size v-model="sSize" @input="saveSysTableConfig" class="action" v-if="isShowSize" />
              <a-tooltip title="列配置" v-show="isShowColumns">
                <action-columns
                  :columns="myColumns"
                  @reset="onColumnsReset"
                  @setColumnHeight="setColumnHeight()"
                  @changeColumns="saveSysTableConfig"
                  ref="columnHeightRef"
                  class="action"
                >
                  <template :slot="slot" v-for="slot in slots">
                    <slot :name="slot"></slot>
                  </template>
                </action-columns>
              </a-tooltip>
            </div>
          </transition>
        </div>
        <a-table
          v-bind="{ ...$attrs, ...$props, columns: visibleColumns, title: undefined, loading: false }"
          :size="sSize"
          @expandedRowsChange="onExpandedRowsChange"
          @change="onChange"
          @expand="onExpand"
        >
          <!-- v-on="$listeners" table所有方法已经重写，无需继承-->
          <template slot-scope="text, record, index" :slot="slot" v-for="slot in scopedSlots">
            <slot :name="slot" v-bind="{ text, record, index }"></slot>
          </template>
          <template :slot="slot" v-for="slot in slots">
            <slot :name="slot"></slot>
          </template>
          <template slot-scope="record, index, indent, expanded" :slot="$scopedSlots.expandedRowRender ? 'expandedRowRender' : ''">
            <slot v-bind="{ record, index, indent, expanded }" :name="$scopedSlots.expandedRowRender ? 'expandedRowRender' : ''"></slot>
          </template>
        </a-table>
      </div>
    </a-spin>
  </div>
</template>
<script>
import ActionSize from '@/components/pt/table/ActionSize'
import ActionColumns from '@/components/pt/table/ActionColumns'
export default {
  name: 'AdvanceTable',
  components: { ActionColumns, ActionSize },
  props: {
    tableKey: String, // 表格唯一键值，用于保存个性配置使用
    tableLayout: String,
    bordered: Boolean,
    childrenColumnName: { type: String, default: 'children' },
    columns: Array,
    components: Object,
    dataSource: Array,
    defaultExpandAllRows: Array[String],
    expandedRowKeys: Array[String],
    expandedRowRender: Function,
    expandRowByClick: Boolean,
    expandIconColumnIndex: Number,
    footer: Function,
    indentSize: Number,
    loading: Boolean,
    locale: Object,
    pagination: [Boolean, Object],
    rowClassName: Function,
    rowKey: [String, Function],
    rowSelection: Object,
    // scroll: { type: Object, default: () => ({ y: 'calc(100vh - 330px)' }) },//表格自适应使用
    scroll: Object,
    showHeader: { type: Boolean, default: true },
    size: String,
    title: String,
    customHeaderRow: Function,
    customRow: Function,
    getPopupContainer: Function,
    transformCellText: Function,
    formatConditions: Boolean,
    minHeight: String,
    isShowSetBtn: { type: Boolean, default: true }, // 是否显示刷新区域
    isShowRefresh: { type: Boolean, default: true }, // 是否显示刷新区域
    isShowSize: { type: Boolean, default: true }, // 是否显示密度区域
    isShowColumns: { type: Boolean, default: true }, // 是否显示列配置
    isShowFull: { type: Boolean, default: true }, // 是否显示全屏
    isTableConfig: { type: Boolean, default: true }, // 列表配置是否存储数据库
    setingFlag: { type: String, default: 'table' }
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
      myColumns: this.columns,
      fullScreen: false,
      conditions: {},
      loadType: !this.isTableConfig, // 存储数据库的话需要等数据加载完成，不存储数据库则直接使用传入列信息
      sysTableConfig: {},
      isShowSetting: false,
      isShowSettingBtn: false,
      configType: '',
      settingTop: '4px'
    }
  },
  watch: {
    sSize (val) {
      this.getSettingBtnTop(val)
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
          return true
        }
      })
      return filterColumns
    },
    visibleColumns () {
      if (this.loadType) {
        return this.myColumns.filter(col => col.visible)
      } else {
        return []
      }
    }
  },
  created () {
    this.getTableConfig()
    this.getSettingBtnTop(this.sSize)
  },
  mounted () {},
  beforeDestroy () {
    this.removeListener()
  },
  methods: {
    getSettingBtnTop (sSizeValue) { // 获取配置区域图标距离表头顶部距离
      if (sSizeValue === 'small') {
        this.settingTop = '4px'
      } else if (sSizeValue === 'middle') {
        this.settingTop = '10px'
      } else {
        this.settingTop = '14px'
      }
    },
    refresh () {
      this.$emit('refresh', this.conditions)
    },
    setColumnHeight () {
      let height = this.$refs.table.offsetHeight
      if (height > 430) {
        height = 430
      }
      this.$refs.columnHeightRef.setHeight(height)
    },
    setColumnHeight1 () {
      let height = this.$refs.table.offsetHeight
      if (height > 430) {
        height = 430
      }
      this.$refs.columnHeightRef.setHeight(height)
    },
    onSearchChange (conditions, searchOptions) {
      this.conditions = conditions
      this.$emit('search', conditions, searchOptions)
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
    setSysTableConfig (data) {
      this.sysTableConfig.id = data.id
      this.sysTableConfig.version = data.version
      this.sysTableConfig.userId = data.userId
    },
    getTableConfig () {
      this.isShowSettingBtn = !(!this.isShowRefresh && !this.isShowSize && !this.isShowColumns && !this.isShowFull)
      this.loadType = true
    },
    saveSysTableConfig (type) {

    }
  }
}
</script>
<style scoped lang="less">
.advanced-table {
  min-height: calc(100vh - 240px);
  overflow-y: auto;
  background-color: #fff;
  .header-bar {
    padding: 16px 24px;
    display: flex;
    align-items: center;
    border-radius: 4px;
    transition: all 0.3s;
    &.middle {
      padding: 12px 16px;
    }
    &.small {
      padding: 8px 12px;
      border: 1px solid #fff;
      border-bottom: 0;
      .title {
        font-size: 16px;
      }
    }
    .title {
      transition: all 0.3s;
      font-size: 18px;
      color: rgba(0, 0, 0, 0.65);
      font-weight: 700;
    }
    .button {
      flex: 1;
      text-align: right;
      margin: 0 24px;
    }
  }
}

.actions {
  text-align: right;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  background: #f7f7f7;
  position: absolute;
  top: 8px;
  right: 16px;
  z-index: 999;
  .action {
    margin: 0 8px;
    cursor: pointer;
    &:hover {
    }
  }
}
 // /deep/.ant-table-body{
 //    min-height: 300px;
 //  }
/*配置显隐动画  */
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.3, 1);
}
.slide-fade-enter,
.slide-fade-leave-to {
  transform: translateX(10px);
  opacity: 0;
}
.topSet{
  >.action{
    font-size: 16px;
    margin: 0 0.25rem;
    padding:0 4px;
  }
}
</style>
