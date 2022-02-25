<template>
  <div style="padding: 20px;">
    <a-form-model
      ref="genInfoForm"
      :model="info"
      :rules="rules"
    >
      <a-row>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="tplCategory">
            <span slot="label">
              生成模板
            </span>
            <a-select v-model="info.tplCategory" placeholder="请选择">
              <a-select-option value="crud">单表（增删改查）</a-select-option>
              <a-select-option value="treegrid">树表（增删改查）</a-select-option>
              <a-select-option value="sub">主子表（增删改查）</a-select-option>
              <a-select-option value="tree">左树右表（增删改查）</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item prop="functionAuthor">
            <span slot="label">
              个人模板
            </span>
            <a-select
              v-model="info.templateName"
              @change="functionAuthorChange"
              :getPopupContainer="(triggerNode)=>{ return triggerNode.parentNode || document.body}"
              placeholder="请选择个人模板">
              <a-select-option
                v-for="item in templateListOptions"
                :key="item.id"
                :value="item.id">
                {{ item.templateName }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="functionAuthor">
            <span slot="label">
              作者
            </span>
            <a-input placeholder="请输入邮箱" v-model="info.functionAuthor" />
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item prop="functionAuthorEmail">
            <span slot="label">
              邮箱
            </span>
            <a-input placeholder="请输入邮箱" v-model="info.functionAuthorEmail" />
          </a-form-model-item>
        </a-col>

      </a-row>
      <a-row>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="workspacePath">
            <span slot="label">
              工作空间路径
            </span>
            <a-input placeholder="请输入工作空间路径" v-model="info.workspacePath" />
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item prop="webWorkspacePath">
            <span slot="label">
              前端工作空间路径
            </span>
            <a-input placeholder="请输入前端工作空间路径" v-model="info.webWorkspacePath" />
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="moduleName">
            <span slot="label">
              生成模块名
              <a-tooltip>
                <template slot="title">
                  可理解为子系统名，例如 system
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-input placeholder="请输入生成模块名" v-model="info.moduleName" />
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item prop="packageName">
            <span slot="label">
              包路径
            </span>
            <a-input placeholder="请输入生成包路径" v-model="info.packageName" />
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="functionName">
            <span slot="label">
              生成功能名
              <a-tooltip>
                <template slot="title">
                  用作类描述，例如 用户
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-input placeholder="请输入生成功能名" v-model="info.functionName" />
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item prop="businessName">
            <span slot="label">
              生成业务名
              <a-tooltip>
                <template slot="title">
                  可理解为功能英文名，例如 user
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-input placeholder="请输入生成业务名" v-model="info.businessName" />
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row>
        <a-col :span="9" :offset="3">
          <a-form-model-item>
            <span slot="label">
              生成代码方式
              <a-tooltip>
                <template slot="title">
                  默认为zip压缩包下载，也可以自定义生成路径
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-radio-group v-model="info.genType">
              <a-radio :value="'0'">zip压缩包</a-radio>
              <a-radio :value="'1'">自定义路径</a-radio>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="4" :offset="1">
          <a-form-model-item>
            <span slot="label">
              上级菜单
              <a-tooltip>
                <template slot="title">
                  分配到指定菜单下，例如 系统管理
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-tree-select
              v-model="info.parentMenuId"
              :tree-data="menus"
              style="width: 100%"
              :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
              placeholder="请选择系统菜单"
              :replaceFields="treeReplaceFields"
            />
          </a-form-model-item>
        </a-col>
        <a-col :span="4" :offset="1">
          <a-form-model-item label="图标" prop="info.menuIcon">
            <span slot="label">
              菜单图标
              <a-tooltip>
                <template slot="title">
                  选择当前模块对应的菜单图标
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-space size="large" class="selectIconBox">
              <a-icon :component="allIcon[info.menuIcon + 'Icon']" v-if="allIcon[info.menuIcon + 'Icon']" class="selectIcon" />
              <a-icon :type="info.menuIcon" v-if="!allIcon[info.menuIcon + 'Icon']" />
              <a @click="selectIcon" class="selectup">
                <a-icon :type="SelectIcon" />
              </a>
            </a-space>
            <a-card :bordered="false" v-if="iconVisible">
              <icon-selector v-model="info.menuIcon" @change="handleIconChange" :svgIcons="iconList" :allIcon="allIcon" />
            </a-card>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row v-show="info.tplCategory == 'treegrid' || info.tplCategory == 'tree'">
        <a-col :span="18" :offset="3">
          <a-divider orientation="left">
            树表信息
          </a-divider>
        </a-col>
        <a-col :span="9" :offset="3">
          <a-form-model-item>
            <span slot="label">
              树编码字段
              <a-tooltip>
                <template slot="title">
                  树显示的编码字段名， 如：dept_id
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-input placeholder="树编码字段" v-model="info.treeCode" value="id" disabled/>
            <!--            <a-select v-model="info.treeCode" placeholder="请选择">
              <a-select-option v-for="(item, index) in info.columns" :key="index" :value="item.columnName" >
                {{ item.columnName + (item.columnComment === null ? '': '：' + item.columnComment) }}
              </a-select-option>
            </a-select>-->
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item>
            <span slot="label">
              树父编码字段
              <a-tooltip>
                <template slot="title">
                  树显示的父编码字段名， 如：parent_Id
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-input placeholder="树父编码字段" v-model="info.treeParentCode" value="parent_id" disabled />
            <!--            <a-select v-model="info.treeParentCode" placeholder="请选择">
              <a-select-option v-for="(item, index) in info.columns" :key="index" :value="item.columnName" >
                {{ item.columnName + (item.columnComment === null ? '': '：' + item.columnComment) }}
              </a-select-option>
            </a-select>-->
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="3">
          <a-form-model-item>
            <span slot="label">
              树名称字段
              <a-tooltip>
                <template slot="title">
                  树节点的显示名称字段名， 如：dept_name
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-select v-model="info.treeName" placeholder="请选择">
              <a-select-option v-for="(item, index) in info.columns" :key="index" :value="item.columnName" >
                {{ item.columnName + (item.columnComment === null ? '': '：' + item.columnComment) }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row v-show="info.tplCategory == 'sub' || info.tplCategory == 'tree'">
        <a-col :span="18" :offset="3">
          <a-divider orientation="left">
            关联信息
          </a-divider>
        </a-col>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="subTableName">
            <span slot="label">
              关联子表的表名
              <a-tooltip>
                <template slot="title">
                  关联子表的表名， 如：sys_user
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-select v-model="info.subTableName" placeholder="请选择" @change="handleSubTableNameChange">
              <a-select-option v-for="(item, index) in subTableList" :key="index" :value="item.tableName">
                {{ item.tableName + '：' + item.tableComment }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="1">
          <a-form-model-item prop="subTableFkName">
            <span slot="label">
              子表关联的外键名
              <a-tooltip>
                <template slot="title">
                  子表关联的外键名， 如：user_id
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-select v-model="info.subTableFkName" placeholder="请选择">
              <a-select-option v-for="(item, index) in subColumn" :key="index" :value="item.columnName" >
                {{ item.columnName + (item.columnComment === null ? '': '：' + item.columnComment) }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row>
        <a-col :span="18" :offset="3">
          <a-divider orientation="left">
            其他选项
          </a-divider>
        </a-col>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="attachOption">
            <span slot="label">
              是否可上传附件
              <a-tooltip>
                <template slot="title">
                  默认新增页面，编辑页面不上传附件
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-radio-group v-model="info.attachOption">
              <a-radio :value="'1'">是</a-radio>
              <a-radio :value="'0'">否</a-radio>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="9" :offset="3">
          <a-form-model-item prop="disableEnableOption">
            <span slot="label">
              是否有停用启用
              <a-tooltip>
                <template slot="title">
                  必须要包含status字段才可以
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-radio-group v-model="info.disableEnableOption">
              <a-radio :value="'1'">是</a-radio>
              <a-radio :value="'0'">否</a-radio>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
      </a-row>
    </a-form-model>
  </div>
</template>
<script>
import allIcon from '@/core/icons'
import icons from '@/utils/requireIcons'
import IconSelector from '@/components/IconSelector'
import { listTable, getGenTable } from '@/api/tool/gen'
export default {
  name: 'GenInfoForm',
  props: {
    info: {
      type: Object,
      default: null
    },
    menus: {
      type: Array,
      default: null
    },
    templateListOptions: {
      type: Array,
      default: null
    }
  },
  components: {
    IconSelector
  },
  created () {
    this.initInfoData()
    this.getDbList()
  },
  data () {
    return {
      SelectIcon: 'down',
      allIcon,
      iconVisible: false,
      iconList: icons,
      visible: false,
      loading: false,
      subTableList: [],
      subColumn: [],
      // 模态框数据
      data: {},
      rules: {
        tplCategory: [{ required: true, message: '请选择生成模板', trigger: 'blur' }],
        packageName: [{ required: true, message: '请输入生成包路径', trigger: 'blur' }],
        moduleName: [{ required: true, message: '请输入生成模块名', trigger: 'blur' }],
        businessName: [{ required: true, message: '请输入生成业务名', trigger: 'blur' }],
        functionName: [{ required: true, message: '请输入生成功能名', trigger: 'blur' }],
        subTableName: [{ required: true, message: '请输入关联子表的表名', trigger: 'blur' }],
        subTableFkName: [{ required: true, message: '请输入子表关联的外键名', trigger: 'blur' }]
      },
      // 类型数据字典
      typeOptions: [],
      // 类型数据字典
      statusOptions: [],
      treeReplaceFields: {
        children: 'children',
        title: 'menuName',
        key: 'id',
        value: 'id'
      }
    }
  },
  methods: {
    initInfoData () {
      if (this.info.attachOption === null) {
        this.info.attachOption = '0'
      }
      if (this.info.disableEnableOption === null) {
        this.info.disableEnableOption = '0'
      }
    },
    handleSubTableNameChange (tableName, option, extra) {
      this.info.subTableFkName = ''
      const tableOption = this.subTableList.filter(item => tableName === item.tableName)
      if (tableOption.length > 0) {
        const optionValue = tableOption[0].tableId
        getGenTable(optionValue).then(res => {
          if (res.code === 200) {
            this.subColumn = res.data.rows
          }
        })
      }
    },
    getDbList () {
      listTable({}).then(res => {
        if (res.code === 200) {
          this.subTableList = res.rows
          // 子表集合去除主表项
          this.subTableList = this.subTableList.filter(item => this.info.tableName !== item.tableName)
          const tableOption = this.subTableList.filter(item => this.info.subTableName === item.tableName)
          if (tableOption.length > 0) {
            const optionValue = tableOption[0].tableId
            getGenTable(optionValue).then(res => {
              if (res.code === 200) {
                this.subColumn = res.data.rows
              }
            })
          }
        }
      })
    },
    // 关闭模态框
    close () {
      this.visible = false
    },
    functionAuthorChange () {
      const option = this.templateListOptions.filter(item => this.info.templateName === item.id)
      if (option.length > 0) {
        const optionValue = option[0]
        this.info.templateName = optionValue.templateName
        this.info.functionAuthor = optionValue.functionAuthor
        this.info.functionAuthorEmail = optionValue.functionAuthorEmail
        this.info.workspacePath = optionValue.workspacePath
        this.info.moduleName = optionValue.moduleName
        this.info.packageName = optionValue.packageName
        this.info.webWorkspacePath = optionValue.webWorkspacePath
      }
    },
    handleIconChange (icon) {
      this.SelectIcon = 'down'
      this.iconVisible = false
      this.info.menuIcon = icon
    },
    changeIcon (type) {
      this.currentSelectedIcon = type
    },
    selectIcon () {
      this.iconVisible = !this.iconVisible
      if (this.iconVisible) {
        this.SelectIcon = 'up'
      } else {
        this.SelectIcon = 'down'
      }
    },
    cancelSelectIcon () {
      this.iconVisible = false
    },
    // 打开抽屉(由外面的组件调用)
    show (data) {
      if (data) {
        this.data = data
      }
      this.visible = true
    }
  }
}
</script>
