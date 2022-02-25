import { getConfig, saveConfig, checkConfigKeyUnique } from '@/api/system/config'
import AntModal from '@/components/pt/dialog/AntModal'
export default {
  name: 'CreateForm',
  props: {
    typeOptions: {
      type: Array,
      required: true
    }
  },
  components: {
    AntModal
  },
  data () {
    return {
      loading: false,
      formTitle: '',
      // 表单参数
      form: {
        id: undefined,
        configName: undefined,
        configKey: undefined,
        configValue: undefined,
        configType: 'Y',
        remark: undefined
      },
      open: false,
      rules: {
        configName: [{ required: true, message: '参数名称不能为空', trigger: 'blur' }],
        configKey: [{ required: true, message: '参数编码不能为空', trigger: 'blur' },
          { validator: this.checkConfigKeyUnique }],
        configValue: [{ required: true, message: '参数值不能为空', trigger: 'blur' }]
      }
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
    // 取消按钮
    cancel () {
      this.open = false
      this.$emit('close')
    },
    // 表单重置
    reset () {
      if (this.$refs.form !== undefined) {
        this.$refs.form.resetFields()
      }
    },
    /** 新增按钮操作 */
    handleAdd () {
      this.open = true
      this.formTitle = '添加参数'
    },
    /** 修改按钮操作 */
    handleUpdate (row, ids) {
      const configId = row ? row.id : ids
      getConfig(configId).then(response => {
        this.form = response.data
        this.open = true
        this.formTitle = '修改参数'
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            saveConfig(this.form).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          } else {
            saveConfig(this.form).then(response => {
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
    checkConfigKeyUnique (rule, value, callback) {
      const msg = '参数编码已存在'
      if (value === '') {
        callback()
      } else {
        const checkData = {
          configKey: value,
          id: this.form.id !== undefined ? this.form.id : ''
        }
        checkConfigKeyUnique(checkData).then(response => {
          if (response.data.code === '1') {
            callback()
          } else {
            callback(msg)
          }
        })
      }
    }
  }
}
