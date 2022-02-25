import AntModal from '@/components/pt/dialog/AntModal'
import {
  getTemplate,
  addTemplate,
  updateTemplate,
  findMaxSort
} from '@/api/tool/genConfigTemplate'

export default {
  name: 'CreateForm',
  props: {
    statusOptions: {
      type: Array,
      required: true
    },
    templateDefaultOptions: {
      type: Array,
      required: true
    }
  },
  components: {
    AntModal
  },
  data() {
    return {
      open: false,
      spinning: false,
      labelCol: {
        span: 4
      },
      wrapperCol: {
        span: 14
      },
      loading: false,
      total: 0,
      id: undefined,
      formTitle: '',
      // 表单参数
      form: {},
      rules: {
        templateName: [{
          required: true,
          message: '模板名称不能为空',
          trigger: 'change'
        }],
        functionAuthor: [{
          required: true,
          message: '作者不能为空',
          trigger: 'change'
        }],
        functionAuthorEmail: [{
          type: 'email',
          message: '请正确填写邮箱地址',
          trigger: ['blur', 'change']
        }],
        workspacePath: [{
          required: true,
          message: '后端工作空间不能为空',
          trigger: 'change'
        }],
        wenWorkspacePath: [{
          required: true,
          message: '前端工作空间不能为空',
          trigger: 'change'
        }],
        moduleName: [{
          required: true,
          message: '模块名不能为空',
          trigger: 'change'
        }],
        packageName: [{
          required: true,
          message: '模块路径不能为空',
          trigger: 'change'
        }],
        sort: [{
          required: true,
          message: '排序不能为空',
          trigger: 'change'
        }]
      }
    }
  },
  filters: {},
  created() {},
  computed: {},
  watch: {},
  mounted() {},
  methods: {
    onClose() {
      this.open = false
      this.reset()
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        templateName: undefined,
        functionAuthor: undefined,
        functionAuthorEmail: undefined,
        workspacePath: undefined,
        moduleName: undefined,
        packageName: undefined,
        webWorkspacePath: undefined,
        templateDefault: 'N',
        status: '0',
        remark: undefined
      }
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      /** 获取最大编号 */
      findMaxSort().then(response => {
        this.form.sort = response.data
        this.open = true
        this.formTitle = '添加模板'
      })
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      this.spinning = !this.spinning
      const templateId = row.id
      getTemplate(templateId).then(response => {
        this.open = true
        this.form = response.data
        this.formTitle = '修改模板'
      })
      this.spinning = !this.spinning
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateTemplate(this.form).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          } else {
            addTemplate(this.form).then(response => {
              this.$message.success(
                '新增成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          }
        } else {
          return false
        }
      })
    },
    back () {
      this.$router.push('/tool/gen/genconfigtemplate/index')
    }
  }
}
