import AntModal from '@/components/pt/dialog/AntModal'
import { getSysPortalConfig, addSysPortalConfig, updateSysPortalConfig, checkCodeUnique } from '@/api/system/sysPortalConfig'
import SelectUser from '@/components/pt/selectUser/SelectUser'
export default {
  name: 'CreateForm',
  props: {
    applicationRangeOptions: {
      type: Array,
      required: true
    },
    isDefaultOptions: {
      type: Array,
      required: true
    }
    },
  components: {
    AntModal,
    SelectUser
  },
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
      roleOptions: [],
      isShowResourceId: true,
      total: 0,
      id: undefined,
      formTitle: '添加多栏目门户配置',
      // 表单参数
      form: {},
      rules: {
        name: [{ required: true, message: '名称不能为空', trigger: 'blur' }],
        code: [{ required: true, message: '编码不能为空', validator: validateCode, trigger: 'blur' }]
      }
    }
  },
  filters: {},
  created () {},
  computed: {},
  watch: {
  },
  mounted () {},
  methods: {
    rangeChange (e) {
        const applicationRange = this.form.applicationRange
        if (applicationRange === 'R' || applicationRange === 'U') {
          this.isShowResourceId = true
        } else {
          this.isShowResourceId = false
        }
        this.form.systemDefinedId = ''
    },
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
        applicationRange: 'R',
        isDefault: 'N',
        resourceId: undefined,
        systemDefinedId: undefined,
        content: undefined,
        sort: undefined,
        status: undefined
      }
    },
    /** 新增按钮操作 */
    handleAdd () {
      this.reset()
      this.open = true
      this.formTitle = '添加多栏目门户配置'
      getSysPortalConfig().then(response => {
        this.roleOptions = response.data.roles
      })
    },
    /** 修改按钮操作 */
    handleUpdate (row) {
      this.reset()
      this.open = true
      this.spinning = !this.spinning
      const sysPortalConfigId = row.id
      getSysPortalConfig(sysPortalConfigId).then(response => {
        this.form = response.data.data
        this.rangeChange()
        const applicationRange = this.form.applicationRange
        if (applicationRange === 'U') {
          let userIds = ''
          let userNames = ''
          response.data.listMap.forEach(node => {
            userIds += node.id + ','
            userNames += node.name + ','
          })
          userIds = userIds.substr(0, userIds.length - 1)
          userNames = userNames.substr(0, userNames.length - 1)
          this.form.resourceId = { ids: userIds, names: userNames }
        }
        this.roleOptions = response.data.roles
        this.formTitle = '修改多栏目门户配置'
        this.spinning = !this.spinning
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
            const saveForm = JSON.parse(JSON.stringify(this.form))
            if (this.form.resourceId !== undefined && this.form.applicationRange === 'U') {
              saveForm.resourceId = this.form.resourceId.ids
            }
           if (this.form.id !== undefined) {
            updateSysPortalConfig(saveForm).then(response => {
                this.$message.success('新增成功', 3)
                this.open = false
                this.$emit('ok')
                this.$emit('close')
            })
          } else {
            addSysPortalConfig(saveForm).then(response => {
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
      const index = '/system/sysportalconfig/index'
      this.$router.push(index)
    }
  }
}
