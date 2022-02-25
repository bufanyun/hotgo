import AntModal from '@/components/pt/dialog/AntModal'
import { getSysPortlet, addSysPortlet, updateSysPortlet, findMaxSort, checkCodeUnique } from '@/api/system/sysPortlet'

export default {
  name: 'CreateForm',
  props: {
    showTitleOptions: {
      type: Array,
      required: true
    },
    isAllowDragOptions: {
      type: Array,
      required: true
    },
    statusOptions: {
      type: Array,
      required: true
    } },
  components: {
    AntModal },
  data () {
    const validateCode = (rule, value, callback) => {
      if (value === '' || value === undefined || value === null) {
        callback(new Error('小页编码不允许为空'))
      } else {
        checkCodeUnique(this.form.id, this.form.code)
        callback()
      }
    }
    return {
      open: false,
      spinning: false,
      delayTime: 100,
      labelCol: { span: 4 },
      wrapperCol: { span: 14 },
      loading: false,
      total: 0,
      id: undefined,
      formTitle: '添加工作台小页管理',
      // 表单参数
      form: {
      },
      rules: {
        name: [{ required: true, message: '小页名称不能为空', trigger: 'blur' }],
        code: [{ required: true, message: '小页编码不能为空', validator: validateCode, trigger: 'blur' }],
        status: [{ required: true, message: '状态不能为空', trigger: 'blur' }]
      }
    }
  },
  filters: {},
  created () {},
  computed: {},
  watch: {},
  mounted () {},
  methods: {
    onClose () {
      this.open = false
      this.reset()
      this.$emit('close')
    },
    // 取消按钮
    cancel () {
      this.open = false
      this.reset()
      this.$emit('close')
    },
    // 表单重置
    reset () {
      this.form = {
        id: undefined,
        name: undefined,
        code: undefined,
        url: undefined,
        refreshRate: undefined,
        showTitle: 'N',
        isAllowDrag: 'N',
        xGridNumber: '2',
        yGridNumber: 100,
        sort: undefined,
        status: '0'
      }
    },
    /** 新增按钮操作 */
    handleAdd () {
      this.reset()
      /** 获取最大编号 */
      findMaxSort().then(response => {
        this.form.sort = response.data
        this.open = true
        this.formTitle = '添加工作台小页管理'
      })
    },
    /** 修改按钮操作 */
    handleUpdate (row) {
      this.reset()
      this.open = true
      this.spinning = !this.spinning
      const sysPortletId = row.id
      getSysPortlet(sysPortletId).then(response => {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  this.form = response.data
        this.formTitle = '修改工作台小页管理'
        this.spinning = !this.spinning
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
            const saveForm = JSON.parse(JSON.stringify(this.form))
                                                                                        if (this.form.id !== undefined) {
            updateSysPortlet(saveForm).then(response => {
                this.$message.success('新增成功', 3)
                this.open = false
                this.$emit('ok')
                this.$emit('close')
            })
          } else {
            addSysPortlet(saveForm).then(response => {
                this.$message.success('新增成功', 3)
                this.open = false
                this.$emit('ok')
                this.$emit('close')
            })
          }
        } else {
          return false
        }
      })
    },
    back () {
      const index = '/system/sysportlet/index'
      this.$router.push(index)
    }
  }
}
