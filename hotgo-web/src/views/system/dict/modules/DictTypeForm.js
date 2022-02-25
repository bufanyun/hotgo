import AntModal from '@/components/pt/dialog/AntModal'
import { checkDictTypeUnique, getType, saveType } from '@/api/system/dict/type'

export default {
  name: 'CreateForm',
  props: {
    statusOptions: {
      type: Array,
      required: true
    }
  },
  components: {
    AntModal
  },
  data() {
    return {
      loading: false,
      formTitle: '',
      // 表单参数
      form: {
        id: undefined,
        name: undefined,
        type: undefined,
        status: '1',
        remark: undefined
      },
      open: false,
      rules: {
        name: [{ required: true, message: '字典名称不能为空', trigger: 'blur' }],
        type: [{ required: true, message: '字典类型不能为空', trigger: 'blur' },
          { validator: this.checkDictTypeUnique }]
      }
    }
  },
  filters: {},
  created() {
  },
  computed: {},
  watch: {},
  methods: {
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
      this.$emit('close')
    },
    // 表单重置
    reset() {
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.formTitle = '添加字典类型'
    },
    /** 修改按钮操作 */
    handleUpdate(row, ids) {
      this.reset()
      const dictId = row ? row.id : ids
      getType(dictId).then(response => {
        this.form = response
        this.form.status = '' + response.status
        this.open = true
        this.formTitle = '修改【' + this.form.name + '】类型'
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs.form.validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            saveType(this.form).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          } else {
            saveType(this.form).then(response => {
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
    checkDictTypeUnique(rule, value, callback) {
      const msg = '数据字典类型已存在'
      if (value === '') {
        callback()
      } else {
        const checkData = {
          type: value,
          id: this.form.id !== undefined ? this.form.id : ''
        }
        checkDictTypeUnique(checkData).then(response => {
          if (response.is_unique) {
            callback()
          } else {
            callback(msg)
          }
        })
      }
    }
  }
}
