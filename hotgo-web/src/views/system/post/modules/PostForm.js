import { getPost, savePost, checkPostCodeUnique, findMaxSort, checkPostNameUnique } from '@/api/system/post'
import AntModal from '@/components/pt/dialog/AntModal'
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
  data () {
    return {
      loading: false,
      formTitle: '',
      // 表单参数
      form: {
        id: undefined,
        postCode: undefined,
        postName: undefined,
        sort: 0,
        status: '1',
        remark: undefined
      },
      open: false,
      rules: {
        name: [{ required: true, message: '岗位名称不能为空', trigger: 'blur' },
          { validator: this.checkPostNameUnique }],
        code: [{ required: true, message: '岗位编码不能为空', trigger: 'blur' },
          { validator: this.checkPostCodeUnique }],
        sort: [{ required: true, message: '显示顺序不能为空', trigger: 'blur' }]
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
      this.reset()
      this.$emit('close')
    },
    // 表单重置
    reset () {
    },
    /** 新增按钮操作 */
    handleAdd () {
      this.reset()
      /** 获取最大编号 */
      findMaxSort().then(response => {
        this.form.sort = response.sort
        this.open = true
        this.formTitle = '添加岗位'
      })
    },
    /** 修改按钮操作 */
    handleUpdate (row, ids) {
      this.reset()
      const postId = row ? row.id : ids
      getPost(postId).then(response => {
        this.form = response
        this.open = true
        this.formTitle = '修改岗位'
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            savePost(this.form).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          } else {
            savePost(this.form).then(response => {
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
    checkPostCodeUnique (rule, value, callback) {
      const msg = '岗位编码已存在'
      if (value === '') {
        callback()
      } else {
        const checkData = {
          code: value,
          id: this.form.id !== undefined ? this.form.id : ''
        }
        checkPostCodeUnique(checkData).then(response => {
          if (response.is_unique) {
            callback()
          } else {
            callback(msg)
          }
        })
      }
    },
    checkPostNameUnique (rule, value, callback) {
      const msg = '岗位名称已存在'
      if (value === '') {
        callback()
      } else {
        const checkData = {
          name: value,
          id: this.form.id !== undefined ? this.form.id : ''
        }
        checkPostNameUnique(checkData).then(response => {
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
