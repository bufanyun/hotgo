<template>
  <div>
    <a-card
      :bordered="false"
      v-if="tableId != 0"
      style="margin-bottom: 10px;height:calc(100vh - 172px)">
      <a-tabs default-active-key="2">
        <a-tab-pane key="1" tab="基本信息" force-render>
          <basic-info-form ref="basicInfo" :info="info" style="height:calc(100vh - 280px);overflow: auto;"/>
        </a-tab-pane>
        <a-tab-pane key="2" tab="字段信息" force-render >
          <!-- 表格 -->
          <a-table
            ref="table"
            size="small"
            :scroll="{ x: 1500, y: 'calc(100vh - 320px)' }"
            bordered
            :columns="columns"
            :loading="tableLoading"
            :data-source="tableList"
            row-key="columnId"
            :pagination="false">

            <template slot="dragIcon">
              <a-icon type="drag" class="dragIconClass"/>
            </template>
            <span slot="colSpanTitle">
              <a-tooltip placement="topRight" style="cursor: help">
                <template slot="title">
                  <span>该选项表示24栅格中<br>摆放控件个数</span>
                </template>
                列数<a-icon type="question-circle" />
              </a-tooltip>
            </span>
            <template slot="columnName" slot-scope="text">
              <div style="width: 100px;">
                <span>{{ text }}</span>
              </div>
            </template>
            <!-- 字段描述 -->
            <template slot="columnComment" slot-scope="text, record">
              <a-input v-model="record.columnComment" style="width: 90px;" :disabled="checkDisabledColumn(record)"></a-input>
            </template>
            <!-- Java类型 -->
            <template slot="javaType" slot-scope="text, record">
              <!--在锁定列环境下下拉框无法跟随滚动 -->
              <a-select
                v-model="record.javaType"
                @dropdownVisibleChange="javaTypeVisible(record)"
                @change="javaTypeChange(record)"
                style="width:90px"
                :disabled="checkDisabledColumn(record)">
                <a-select-option v-for="(value, index) in javaTypeList" :key="index" :value="value" :title="value">
                  {{ value }}
                </a-select-option>
              </a-select>
            </template>
            <!-- 编辑 -->
            <template slot="isEdit" slot-scope="text, record">
              <a-checkbox v-model="record.isEdit" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 列表 -->
            <template slot="isList" slot-scope="text, record">
              <a-checkbox v-model="record.isList" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 查询 -->
            <template slot="isQuery" slot-scope="text, record">
              <a-checkbox v-model="record.isQuery" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 日志 -->
            <template slot="isLog" slot-scope="text, record">
              <a-checkbox v-model="record.isLog" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 查询方式 -->
            <template slot="queryType" slot-scope="text, record">
              <a-select
                v-model="record.queryType"
                :getPopupContainer="
                  triggerNode => {
                    return triggerNode.parentNode || document.body
                  }
                "
                style="width:100%"
                :disabled="checkDisabledColumn(record)">
                <a-select-option value="EQ">=</a-select-option>
                <a-select-option value="NE">!=</a-select-option>
                <a-select-option value="GT">></a-select-option>
                <a-select-option value="GTE">>=</a-select-option>
                <a-select-option value="LT">&lt;</a-select-option>
                <a-select-option value="LTE">&lt;=</a-select-option>
                <a-select-option value="LIKE">LIKE</a-select-option>
                <a-select-option value="BETWEEN">BETWEEN</a-select-option>
              </a-select>
            </template>
            <!-- 查询方式 -->
            <template slot="alignType" slot-scope="text, record">
              <a-select
                v-model="record.alignType"
                :getPopupContainer="
                  triggerNode => {
                    return triggerNode.parentNode || document.body
                  }
                "
                style="width:100%"
                :disabled="checkDisabledColumn(record)">
                <a-select-option value="center">居中</a-select-option>
                <a-select-option value="left">居左</a-select-option>
                <a-select-option value="right">居右</a-select-option>
              </a-select>
            </template>
            <!-- 必填 -->
            <template slot="isRequired" slot-scope="text, record">
              <a-checkbox v-model="record.isRequired" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 唯一性 -->
            <template slot="isUnique" slot-scope="text, record">
              <a-checkbox v-model="record.isUnique" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 新行 -->
            <template slot="isNewRow" slot-scope="text, record">
              <a-checkbox v-model="record.isNewRow" :disabled="checkDisabledColumn(record)"></a-checkbox>
            </template>
            <!-- 列数 -->
            <template slot="colSpan" slot-scope="text, record">
              <a-select
                v-model="record.colSpan"
                style="width:100%"
                :getPopupContainer="
                  triggerNode => {
                    return triggerNode.parentNode || document.body
                  }
                "
                :disabled="checkDisabledColumn(record)">
                <a-select-option :value="1">一列/24栅格</a-select-option>
                <a-select-option :value="2">两列/24栅格</a-select-option>
                <a-select-option :value="3">三列/24栅格</a-select-option>
                <a-select-option :value="4">四列/24栅格</a-select-option>
              </a-select>
            </template>
            <!-- 显示类型 -->
            <template slot="htmlType" slot-scope="text, record">
              <a-select
                v-model="record.htmlType"
                :getPopupContainer="
                  triggerNode => {
                    return triggerNode.parentNode || document.body
                  }
                "
                @dropdownVisibleChange="htmlTypeVisible(record)"
                style="width:100%"
                :disabled="checkDisabledColumnType(record)">
                <a-select-option v-for="(item, index) in htmlTypeList" :key="index" :value="item.code" >
                  {{ item.label }}
                </a-select-option>
              </a-select>
            </template>
            <!-- 字典类型 -->
            <template slot="dictType" slot-scope="text, record">
              <a-select
                v-model="record.dictType"
                :getPopupContainer="
                  triggerNode => {
                    return triggerNode.parentNode || document.body
                  }
                "
                placeholder="请选择"
                style="width:100%"
                show-search
                :disabled="checkDisabledColumn(record,'dictType')">
                <a-select-option value="">请选择</a-select-option>
                <a-select-option v-for="item in dictOptions" :key="item.dictType" :value="item.dictType">
                  {{ item.dictName }}
                </a-select-option>
              </a-select>
            </template>
            <!-- 字段校验 -->
            <template slot="colCheck" slot-scope="text,record">
              <a-select
                v-model="record.colCheck"
                :getPopupContainer="
                  triggerNode => {
                    return triggerNode.parentNode || document.body
                  }
                "
                style="width:100%"
                mode="multiple"
                :disabled="checkDisabledColumn(record)">
                <a-select-option value="">请选择</a-select-option>
                <a-select-option value="date">日期</a-select-option>
                <a-select-option value="number">数值</a-select-option>
                <a-select-option value="integer">整数</a-select-option>
                <a-select-option value="positiveInteger">正整数</a-select-option>
                <a-select-option value="phone">手机号</a-select-option>
                <a-select-option value="tel">固定号码</a-select-option>
                <a-select-option value="zipCode">邮编</a-select-option>
                <a-select-option value="email">邮箱</a-select-option>
                <a-select-option value="qq">QQ号</a-select-option>
              </a-select>
            </template>
          </a-table>
        </a-tab-pane>
        <a-tab-pane key="3" tab="生成信息" force-render>
          <gen-info-form
            v-if="templateListOptions.length>0"
            ref="genInfo"
            :info="info"
            :menus="menus"
            :templateListOptions="templateListOptions"
            style="height:calc(100vh - 280px);overflow: auto;"/>
        </a-tab-pane>
      </a-tabs>
    </a-card>
    <a-card :bordered="false">
      <div style="text-align: right;padding: 5px;">
        <a-space>
          <a-button type="primary" @click="submitForm">
            保存
          </a-button>
          <a-button type="primary" @click="submitFormGenCode">
            保存并生成代码
          </a-button>
          <a-button @click="back">
            取消
          </a-button>
        </a-space>
      </div>
    </a-card>
  </div>
</template>

<script>
  import { getGenTable, updateGenTable, updateGenTableNoValidated, genCode } from '@/api/tool/gen'
  import { optionselect as getDictOptionselect } from '@/api/system/dict/type'
  import { listMenu as getMenuTreeselect } from '@/api/system/menu'
  import { listTemplate } from '@/api/tool/genConfigTemplate'
  import BasicInfoForm from './BasicInfoForm'
  import GenInfoForm from './GenInfoForm'
  import Sortable from 'sortablejs'// 列交换第三方插件
  import storage from 'store'
  export default {
    name: 'GenEdit',
    components: {
      BasicInfoForm,
      GenInfoForm
    },
    data () {
      return {
        tableId: 0,
        formTitle: '修改生成配置',
        // 表格加载
        tableLoading: false,
        // 字典信息
        dictOptions: [],
         templateListOptions: [],
        javaTypeList: ['Long', 'String', 'Integer', 'Double', 'BigDecimal', 'Date', 'DateTime'],
        htmlTypeList: [
        { 'code': 'input', 'label': '文本框' },
        { 'code': 'number', 'label': '数字框' },
        { 'code': 'textarea', 'label': '文本域' },
        { 'code': 'select', 'label': '下拉框' },
        { 'code': 'selectMultiple', 'label': '下拉框（多选）' },
        { 'code': 'radio', 'label': '单选框' },
        { 'code': 'checkbox', 'label': '复选框' },
        { 'code': 'user', 'label': '用户控件' },
        { 'code': 'dept', 'label': '部门控件' },
        { 'code': 'datetime', 'label': '日期控件' },
        { 'code': 'time', 'label': '时间' },
        { 'code': 'editor', 'label': '富文本控件' }
        ],
        disabledColumn: ['id', 'create_by', 'create_dept', 'create_time', 'update_by', 'update_time', 'update_ip', 'version', 'del_flag'],
        // 菜单信息
        menus: [],
        // 表详细信息
        info: {},
        // 表数据
        tableList: [],
        // 表数据
        tableListTemp: [],
        // 表头
        columns: [
            {
                dataIndex: 'sort',
                align: 'center',
                width: 50
              },
          {
            title: '字段',
            children: [
              {
                  scopedSlots: {
                    customRender: 'dragIcon'
                  },
                  align: 'center',
                  width: 30
                },
             {
               title: '字段列名',
               dataIndex: 'columnName',
               align: 'left',
               ellipsis: true,
               scopedSlots: {
                 customRender: 'columnName'
               },
               width: 110
             },
             {
               title: '物理类型',
               dataIndex: 'columnType',
               scopedSlots: {
                 customRender: 'columnType'
               },
               align: 'left',
               width: 110
             },
             {
               title: 'Java类型',
               dataIndex: 'javaType',
               scopedSlots: {
                 customRender: 'javaType'
               },
               align: 'left',
               width: 100
             },
             {
               title: 'java属性',
               dataIndex: 'javaField',
               scopedSlots: {
                 customRender: 'javaField'
               },
               align: 'left',
               ellipsis: true,
               width: 120
             },
             {
               title: '字段描述',
               dataIndex: 'columnComment',
               scopedSlots: {
                 customRender: 'columnComment'
               },
               align: 'center',
               width: 100
             }
            ]
          },
          {
            title: '表单',
             children: [
                {
                  title: '编辑',
                  dataIndex: 'isEdit',
                  scopedSlots: {
                    customRender: 'isEdit'
                  },
                  align: 'center',
                  width: '3%'
                },
                {
                  title: '必填',
                  dataIndex: 'isRequired',
                  scopedSlots: {
                    customRender: 'isRequired'
                  },
                  align: 'center',
                  width: '3%'
                },
                {
                  title: '唯一性',
                  dataIndex: 'isUnique',
                  scopedSlots: {
                    customRender: 'isUnique'
                  },
                  align: 'center',
                  width: '4%'
                },
                {
                  title: '显示类型',
                  dataIndex: 'htmlType',
                  scopedSlots: {
                    customRender: 'htmlType'
                  },
                  align: 'center',
                  width: '10%'
                },
                {
                  dataIndex: 'colSpan',
                  scopedSlots: {
                    customRender: 'colSpan'
                  },
                  slots: { title: 'colSpanTitle' },
                  align: 'center',
                  width: 120
                },
                {
                  title: '新行',
                  dataIndex: 'isNewRow',
                  scopedSlots: {
                    customRender: 'isNewRow'
                  },
                  align: 'center',
                  width: '3%'
                },
                {
                  title: '字典类型',
                  dataIndex: 'dictType',
                  scopedSlots: {
                    customRender: 'dictType'
                  },
                  align: 'center',
                  width: 110
                },
                {
                  title: '日志',
                  dataIndex: 'isLog',
                  scopedSlots: {
                    customRender: 'isLog'
                  },
                  align: 'center',
                  width: '3%'
                }, {
                  title: '字段校验',
                  dataIndex: 'colCheck',
                  scopedSlots: {
                    customRender: 'colCheck'
                  },
                  align: 'center',
                  width: 120
                }
             ]
          },
          {
            title: '列表',
             children: [
                {
                  title: '列表',
                  dataIndex: 'isList',
                  scopedSlots: {
                    customRender: 'isList'
                  },
                  align: 'center',
                  width: '3%'
                },
                {
                  title: '查询',
                  dataIndex: 'isQuery',
                  scopedSlots: {
                    customRender: 'isQuery'
                  },
                  align: 'center',
                  width: '3%'
                },
                {
                  title: '查询方式',
                  dataIndex: 'queryType',
                  scopedSlots: {
                    customRender: 'queryType'
                  },
                  align: 'center',
                  width: '8%'
                },
                {
                  title: '对齐方式',
                  dataIndex: 'alignType',
                  scopedSlots: {
                    customRender: 'alignType'
                  },
                  align: 'center',
                  width: '8%'
                }
             ]
          }
        ]
      }
    },
    created () {
      const tableId = this.$route.params && this.$route.params.tableId
      if (tableId) {
        this.tableId = tableId
      } else {
        const genTableId = 'genTableId'
        this.tableId = storage.get(genTableId)
      }
      this.tableLoading = true
      if (this.tableId) {
        // 获取表详细信息
        getGenTable(this.tableId).then(res => {
          const tableList = res.data.rows
          this.tableListTemp = JSON.parse(JSON.stringify(tableList))
          tableList.forEach(e => {
            e.isInsert = e.isInsert === '1'
            e.isEdit = e.isEdit === '1'
            e.isList = e.isList === '1'
            e.isNewRow = e.isNewRow === '1'
            e.isQuery = e.isQuery === '1'
            e.isUnique = e.isUnique === '1'
            e.isLog = e.isLog === '1'
            e.isRequired = e.isRequired === '1'
            e.colCheck = (e.colCheck !== '' && e.colCheck !== null) ? e.colCheck.split(',') : []
          })
          this.tableList = tableList
          this.info = res.data.info
          this.tableLoading = false
          this.rowDrop()
        })
        /** 查询字典下拉列表 */
        getDictOptionselect().then(response => {
          this.dictOptions = response.data
        })
        /** 查询菜单下拉列表 */
        getMenuTreeselect().then(response => {
          this.menus = response.data
        })
        listTemplate().then(response => {
          this.templateListOptions = response.data.list
        })
      }
    },
    methods: {
      /**
       * 校验行是否可以编辑
       */
      checkDisabledColumnType (record, type) {
         let superColumn = record.superColumn
         const htmlType = record.htmlType
         if (type === 'dictType') {
           if (htmlType !== 'select' && htmlType !== 'selectMultiple' && htmlType !== 'radio' && htmlType !== 'checkbox') {
             superColumn = true
             // 清除数据
             record.dictType = ''
           }
         }
         const queryType = record.queryType
         if (htmlType === 'input' || htmlType === 'textarea') {
           // 设置字段居中方式
           record.alignType = 'left'
           if (queryType !== 'EQ' && queryType !== 'LIKE') {
             record.queryType = 'LIKE'
           }
         } else if (htmlType === 'number') {
           // 设置字段居中方式
           record.alignType = 'right'
           if (queryType === 'LIKE') {
             record.queryType = 'EQ'
           }
         } else if (htmlType === 'select' || htmlType === 'radio' || htmlType === 'checkbox' || htmlType === 'selectMultiple') {
           // 设置字段居中方式
           record.alignType = 'center'
           if (queryType === 'LIKE') {
             record.queryType = 'EQ'
           }
         } else if (htmlType === 'time' || htmlType === 'datetime') {
           if (queryType === 'LIKE') {
             record.queryType = 'EQ'
           }
           // 设置字段居中方式
           record.alignType = 'center'
         }
        return superColumn
      },
      /**
       * 校验行是否可以编辑
       */
      checkDisabledColumn (record, type) {
         let superColumn = record.superColumn
         const htmlType = record.htmlType
         if (type === 'dictType') {
           if (htmlType !== 'select' && htmlType !== 'selectMultiple' && htmlType !== 'radio' && htmlType !== 'checkbox') {
             superColumn = true
             // 清除数据
             record.dictType = ''
           }
         }
        return superColumn
      },
      javaTypeVisible (record) {
        // 根据页面加载的java类型初始化下拉框，页面加载的值为后台计算好的
        const javaType = record.javaType
        if (javaType === 'Long' || javaType === 'Integer' || javaType === 'Double' || javaType === 'BigDecimal') {
           this.javaTypeList = ['Long', 'Integer', 'Double', 'BigDecimal']
        } else if (javaType === 'Date' || javaType === 'DateTime') {
           this.javaTypeList = ['Date', 'DateTime']
        } else {
           this.javaTypeList = ['String']
        }
      },
      htmlTypeVisible (record) {
        const javaType = record.javaType
        if (javaType === 'Long' || javaType === 'Integer' || javaType === 'Double' || javaType === 'BigDecimal') {
           this.htmlTypeList = [{ 'code': 'number', 'label': '数字框' }]
        } else if (javaType === 'Date' || javaType === 'DateTime') {
           this.htmlTypeList = [{ 'code': 'datetime', 'label': '日期控件' }, { 'code': 'time', 'label': '时间控件' }]
        } else {
            this.htmlTypeList = [
              { 'code': 'input', 'label': '文本框' },
              { 'code': 'textarea', 'label': '文本域' },
              { 'code': 'select', 'label': '下拉框' },
              { 'code': 'selectMultiple', 'label': '下拉框（多选）' },
              { 'code': 'radio', 'label': '单选框' },
              { 'code': 'checkbox', 'label': '复选框' },
              { 'code': 'user', 'label': '用户控件' },
              { 'code': 'dept', 'label': '部门控件' },
              { 'code': 'editor', 'label': '富文本控件' }
            ]
        }
      },
      dictTypeVisible (record) {
          const htmlType = record.htmlType
          if (htmlType !== 'select' && htmlType !== 'selectMultiple' && htmlType !== 'radio' && htmlType !== 'checkbox') {
            this.$message.error('请选择')
            return false
          }
      },
      /** 提交按钮 */
      submitForm () {
        this.tableLoading = true
        const basicForm = this.$refs.basicInfo.info
        const genForm = this.$refs.genInfo.info
        if (basicForm && genForm) {
          const genTable = Object.assign({}, basicForm, genForm)
          const tableList = JSON.parse(JSON.stringify(this.tableList))
          tableList.forEach(e => {
            e.isInsert = e.isInsert ? '1' : '0'
            e.isEdit = e.isEdit ? '1' : '0'
            e.isNewRow = e.isNewRow ? '1' : '0'
            e.isList = e.isList ? '1' : '0'
            e.isQuery = e.isQuery ? '1' : '0'
            e.isUnique = e.isUnique ? '1' : '0'
            e.isLog = e.isLog ? '1' : '0'
            e.isRequired = e.isRequired ? '1' : '0'
            e.colCheck = e.colCheck.join(',')
          })
          genTable.columns = tableList
          genTable.params = {
            treeCode: genTable.treeCode,
            treeName: genTable.treeName,
            treeParentCode: genTable.treeParentCode,
            parentMenuId: genTable.parentMenuId,
            menuIcon: genTable.menuIcon,
            attachOption: genTable.attachOption,
            disableEnableOption: genTable.disableEnableOption
          }
          updateGenTableNoValidated(genTable).then(res => {
            if (res.code === 200) {
              this.$message.success(res.msg)
              this.tableLoading = false
              this.back()
            } else {
              this.$message.error(res.msg)
              this.tableLoading = false
            }
          }).finally(() => {
            this.tableLoading = false
          })
        } else {
          this.msgError('表单校验未通过，请重新检查提交内容')
          this.tableLoading = false
        }
      },
      /** 生成代码按钮 */
      submitFormGenCode () {
        this.tableLoading = true
        const basicForm = this.$refs.basicInfo.info
        const genForm = this.$refs.genInfo.info
        if (basicForm && genForm) {
          const genTable = Object.assign({}, basicForm, genForm)
          const tableList = JSON.parse(JSON.stringify(this.tableList))
          tableList.forEach(e => {
            e.isInsert = e.isInsert ? '1' : '0'
            e.isEdit = e.isEdit ? '1' : '0'
            e.isNewRow = e.isNewRow ? '1' : '0'
            e.isList = e.isList ? '1' : '0'
            e.isQuery = e.isQuery ? '1' : '0'
            e.isUnique = e.isUnique ? '1' : '0'
            e.isLog = e.isLog ? '1' : '0'
            e.isRequired = e.isRequired ? '1' : '0'
            e.colCheck = e.colCheck.join(',')
          })
          genTable.columns = tableList
          genTable.params = {
            treeCode: genTable.treeCode,
            treeName: genTable.treeName,
            treeParentCode: genTable.treeParentCode,
            parentMenuId: genTable.parentMenuId,
            menuIcon: genTable.menuIcon,
            attachOption: genTable.attachOption,
            disableEnableOption: genTable.disableEnableOption
          }
          updateGenTable(genTable).then(res => {
            if (res.code === 200) {
              genCode(basicForm.tableName).then(response => {
                this.$notification.open({
                    message: '提示',
                    description: response.msg,
                    duration: 3
                  })
                this.tableLoading = false
              })
              this.back()
            } else {
              this.$message.error(res.msg)
              this.tableLoading = false
            }
          }).finally(() => {
            this.tableLoading = false
          })
        } else {
          this.msgError('表单校验未通过，请重新检查提交内容')
          this.tableLoading = false
        }
      },
      /** 关闭按钮 */
      back () {
        this.$router.push('/tool/gen')
      },
      /**
       * 行拖拽事件
       */
      rowDrop () {
        const that = this
        this.$nextTick(() => {
          const xGrid = this.$refs.table
          const el = xGrid.$el.querySelector('.ant-table tbody')
          this.sortable = Sortable.create(
            el,
            {
              handle: '.dragIconClass',
              animation: 100,
              delay: 100,
              chosenClass: 'drag-list-color', // 被选中项的css 类名
              dragClass: 'drag-list-color', // 正在被拖拽中的css类名
              onEnd: ({ newIndex, oldIndex }) => {
                const currRow = that.tableList.splice(oldIndex, 1)[0]
                that.tableList.splice(newIndex, 0, currRow)
                this.$emit('rowDrop', that.tableList)
                that.tableList.forEach(function (item, index) {
                        item.sort = that.tableListTemp[index].sort
                })
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
  /deep/.ant-table-body{
    height: calc(100vh - 280px);
  }
  .dragIconClass{
    font-size: 16px;
    font-weight:bold;
    padding-right:0px;
  }
  .dragIconClass:hover{
    cursor: move;
  }
</style>
