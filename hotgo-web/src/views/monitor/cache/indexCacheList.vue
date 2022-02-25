<template>
  <div>
    <a-row type="flex" :gutter="10">
      <a-col :span="8">
        <a-card :bordered="false" style="height:calc(100vh - 125px);">
          <advance-table
            :columns="columns"
            :data-source="list"
            title="缓存列表"
            :loading="loading"
            :scroll="{ y: 'calc(100vh - 125px)'}"
            rowKey="cacheId"
            size="middle"
            :isTableConfig="false"
            :isShowSetBtn="false"
            tableKey="monitor-cache-cache-name-list-table"
            @refresh="getList"
            :customRow="onClickRow"
            :format-conditions="true"
          >
            <span slot="indexRender" slot-scope="{text, record, index}">
              {{ index + 1 }}
            </span>
            <span slot="operation" slot-scope="{text, record}">
              <a @click="handleClear(record)" v-hasPermi="['monitor:cache:list']">
                删除
              </a>
            </span>
          </advance-table>
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card :bordered="false" style="height:calc(100vh - 125px);">
          <advance-table
            :columns="cacheKeysColumns"
            :data-source="subList"
            title="键名列表"
            :scroll="{ y: 'calc(100vh - 125px)'}"
            :loading="subLoading"
            rowKey="cacheKey"
            size="middle"
            :customRow="onClickSubRow"
            :isTableConfig="false"
            :isShowSetBtn="false"
            tableKey="monitor-cache-cache-key-list-table"
            @refresh="getCacheKeyList"
            :format-conditions="true"
          >
            <span slot="indexSubRender" slot-scope="{text, record, index}">
              {{ index + 1 }}
            </span>
            <span slot="operation" slot-scope="{text, record}">
              <a @click.stop="handleClearByKey(record)" v-hasPermi="['monitor:cache:list']">
                删除
              </a>
            </span>
          </advance-table>
        </a-card>
      </a-col>
      <a-col :span="8">
        <a-card :bordered="false" style="height:calc(100vh - 125px);overflow-y: auto;overflow-x: hidden;">
          <a-form-model ref="form" :model="form" layout="vertical">
            <a-row :gutter="32">
              <a-col :offset="1" :span="22">
                <a-form-model-item label="缓存名称:" prop="cacheName">
                  <a-input v-model="form.cacheName" :readOnly="true"/>
                </a-form-model-item>
              </a-col>
              <a-col :offset="1" :span="22">
                <a-form-model-item label="缓存键名:" prop="cacheKey">
                  <a-input v-model="form.cacheKey" :readOnly="true"/>
                </a-form-model-item>
              </a-col>
              <a-col :offset="1" :span="22">
                <a-form-model-item label="缓存内容:" prop="cacheValue">
                  <a-textarea v-model="form.cacheValue" :rows="16" :readOnly="true"/>
                </a-form-model-item>
              </a-col>
            </a-row>
          </a-form-model>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>
<script>
  import { listCacheName, clearCache, listCacheKey, clearCacheByKey, getCacheValue } from '@/api/monitor/cache'
  import AdvanceTable from '@/components/pt/table/AdvanceTable'
  export default {
    name: 'CacheName',
    components: {
      AdvanceTable
    },
    data () {
      return {
        list: [],
        // 表格缓存的数据 - 用来点击取消时回显数据
        cacheData: [],
        subList: [],
        selectedRowKeys: [],
        selectedSubRowKeys: [],
        selectedRows: [],
        selectedSubRows: [],
        // 高级搜索 展开/关闭
        advanced: false,
        // 非单个禁用
        single: true,
        // 非多个禁用
        multiple: true,
        subMultiple: true,
        currentSelectCacheId: '',
        selectItem: {},
        selectSubItem: {},
        subLoading: false,
        // 表单参数
        form: {},
        total: 0,
        subTotal: 0,
        labelCol: {
          span: 6
        },
        wrapperCol: {
          span: 18
        },
        columns: [
        {
            title: '序号',
            width: '20%',
            align: 'center',
            scopedSlots: {
              customRender: 'indexRender'
            }
          },
          {
            title: '缓存名称',
            dataIndex: 'cacheName',
            ellipsis: true
          },
          {
            title: '备注',
            dataIndex: 'remarks',
            ellipsis: true
          },
          {
            title: '操作',
            dataIndex: 'operation',
            width: '80px',
            scopedSlots: {
              customRender: 'operation'
            }
          }
        ],
        cacheKeysColumns: [
          {
              title: '序号',
              width: '20%',
              align: 'center',
              scopedSlots: {
                customRender: 'indexSubRender'
              }
            },
          {
            title: '缓存键名',
            dataIndex: 'cacheKey',
            ellipsis: true
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
      this.form = {
        cacheName: undefined,
        cacheId: undefined,
        cacheKey: undefined,
        cacheValue: undefined,
        fileType: undefined,
        fileUrl: undefined,
        fileCreateTime: undefined,
        methodName: undefined,
        dtoValue: undefined,
        dto: undefined,
        serviceBean: undefined,
        status: undefined
      }
      this.getList()
    },
    computed: {},
    watch: {
      selectItem (val) {
        this.renderStyle(val, 'main')
        this.getCacheKeyList(val.cacheId)
      },
      selectSubItem (val) {
        this.renderStyle(val, 'sub')
      }
    },
    methods: {
      /** 查询定时任务列表 */
      getList () {
        this.loading = true
        listCacheName().then(response => {
          this.list = response.data
          this.loading = false
          if (this.list.length > 0) {
            this.$nextTick(() => (
              this.selectItem = this.list[0]
            ))
          }
        })
      },

      /** 删除缓存操作 */
      handleClear (row) {
        var that = this
        this.$confirm({
          title: '确认清除所选中数据的缓存吗?',
          onOk () {
            return clearCache(row.cacheId)
              .then(() => {
                that.onSelectChange([], [])
                that.getList()
                that.$message.success(
                  '缓存清除成功',
                  3
                )
            })
          },
          onCancel () {}
        })
      },
      /** 删除缓存key操作 */
      handleClearByKey (row) {
        var that = this
        this.$confirm({
          title: '确认清除所选中数据的缓存吗?',
          onOk () {
            return clearCacheByKey(row.cacheId, row.cacheKey)
              .then(() => {
                that.$message.success(
                  '缓存清除成功',
                  3
                )
                listCacheKey(row.cacheId).then(response => {
                  that.subList = response.data
                  that.subLoading = false
                })
            })
          },
          onCancel () {}
        })
      },
      getCacheKeyList (cacheId) {
        if (typeof (cacheId) === 'string') {
          if (cacheId === null || cacheId === '' || cacheId === undefined) {
            cacheId = this.currentSelectCacheId
          } else {
            this.currentSelectCacheId = cacheId
          }
        } else {
          cacheId = this.currentSelectCacheId
        }
        // 只有保存的数据才加载字表数据
        this.subLoading = true
        listCacheKey(cacheId).then(response => {
          this.subList = response.data
          this.subLoading = false
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
      onClickSubRow (record, index) {
        return {
          on: {
            click: (event) => {
              this.selectSubItem = record
              getCacheValue(record.cacheId, record.cacheKey).then(response => {
                this.form = response.data
              })
            }
          }
        }
      },
      /** 搜索按钮操作 */
      handleQuery () {
        this.getList()
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
      onSelectChange (selectedRowKeys, selectedRows) {
        this.selectedRowKeys = selectedRowKeys
        this.selectedRows = selectedRows
        this.single = selectedRowKeys.length !== 1
        this.multiple = !selectedRowKeys.length
      },
      onSelectSubChange (selectedSubRowKeys, selectedSubRows) {
        this.selectedSubRowKeys = selectedSubRowKeys
        this.selectedSubRows = selectedSubRows
        this.subMultiple = !selectedSubRowKeys.length
      },
      renderStyle (currentRow, type) { // 增加表格选中行样式
        // 类数组
        const rowEles = document.getElementsByClassName('ant-table-row')
        const rowSelectEles = document.getElementsByClassName(type + '-row-selection')
        let rowList
        if (rowSelectEles.length) {
          rowSelectEles[0].classList.remove(type + '-row-selection')
        }
        if (rowEles.length) {
          rowList = [...rowEles]
          // 这里不用 === 是因为获取的 rowKey 是 String 类型，而给与的原数据 key 为 Number 类型
          // 若要用 === ，事先进行类型转换再用吧
          if (type === 'main') {
          rowList.find(row => row.dataset.rowKey === currentRow.cacheId).classList.add(type + '-row-selection')
          } else if (type === 'sub') {
          rowList.find(row => row.dataset.rowKey === currentRow.cacheKey).classList.add(type + '-row-selection')
          }
        }
      }
    }
  }
</script>
<style lang="less" scoped>
  /deep/.main-row-selection {
    background-color: #F0F2F5;
  }
  /deep/.sub-row-selection {
    background-color: #F0F2F5;
  }
 /deep/.ant-table-body{
    height: calc(100vh - 225px);
  }
  /deep/.ant-table-placeholder{
    // 将暂无数据提示隐藏，否则没有数据时将表格高度制定导致暂无数据挤到最底下
    display: none;
  }
</style>
