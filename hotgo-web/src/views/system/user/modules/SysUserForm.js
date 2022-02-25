import { getUser, addUser, updateUser, checkUserNameUnique, checkEmailUnique, checkPhoneUnique } from '@/api/system/user'
import { listDeptTree } from '@/api/system/dept'
import SelectDept from '@/components/pt/selectDept/SelectDept'
import AntModal from '@/components/pt/dialog/AntModal'
export default {
  name: 'CreateForm',
  props: {
    deptCheckedValue: {
      type: Object
    },
    statusOptions: {
      type: Array,
      required: true
    },
    sexOptions: {
      type: Array,
      required: true
    },
    userTypeOptions: {
      type: Array,
      required: true
    },
    defalutExpandedKeys: {
      type: Array
    }
  },
  components: {
    AntModal,
    SelectDept
  },
  data () {
    const validateDeptId = (rule, value, callback) => {
      if (value.ids === '' || value.ids === undefined || value.ids === null) {
        callback(new Error('部门不允许为空'))
      } else {
        callback()
      }
    }
    return {
      expandedKeys: this.defalutExpandedKeys,
      spinning: false,
      delayTime: 100,
      replaceFields: { children: 'children', title: 'label', key: 'id', value: 'id' },
      customStyle: 'background: #fff;ssborder-radius: 4px;margin-bottom: 24px;border: 0;overflow: hidden',
      // 岗位选项
      postOptions: [],
      // 角色选项
      roleOptions: [],
      // 默认密码
      initPassword: undefined,
      formTitle: '',
      // 表单参数
      form: {
        id: undefined,
        dept_id: 0,
        username: undefined,
        nickName: undefined,
        mobile: undefined,
        email: undefined,
        sex: '2',
        status: '1',
        userType: '2',
        remark: undefined,
        postIds: [],
        role: []
      },
      open: false,
      rules: {
        realname: [{ required: true, message: '用户名称不能为空', trigger: 'blur' }],
        id: [{ required: true, message: '用户编号不能为空', trigger: 'blur' }],
        username: [{ required: true, message: '登录名不能为空', trigger: 'blur' },
         { validator: this.checkUserNameUnique, trigger: 'change' }
         ],
        postIds: [{ required: true, message: '岗位不能为空', trigger: 'blur' }],
        secretLevel: [{ required: true, message: '密级不能为空', trigger: 'blur' }],
        dept_id: [{ required: true, message: '部门不能为空', trigger: 'blur', validator: validateDeptId }],
        email: [
          {
            type: 'email',
            message: '请正确填写邮箱地址',
            trigger: ['blur', 'change']
          },
            { validator: this.checkEmailUnique }
        ],
        mobile: [
          {
            pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
            message: '请正确填写手机号',
            trigger: 'blur'
          },
           { validator: this.checkPhoneUnique }
        ]
      }
    }
  },
  filters: {},
  created () {
    this.getConfigKey('sys.user.initPassword').then(response => {
      this.initPassword = response.value
    })
  },
  computed: {},
  watch: {},
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
      // this.$emit('select-tree')
      this.open = true
      this.formTitle = '新增用户'
      getUser().then(response => {
        this.postOptions = response.posts
        this.roleOptions = response.roles
        this.form.dept_id = this.deptCheckedValue
      })
    },
    /** 修改按钮操作 */
    handleUpdate (row, ids) {
        this.open = true
        this.formTitle = '修改【' + row.realname + '】信息'
        this.spinning = !this.spinning
        // this.$emit('select-tree')
        const id = row ? row.id : ids
        getUser(id).then(response => {
          this.form = response
          this.form.dept_id = { ids: response.dept_id, names: response.dept_name }
          this.postOptions = response.posts
          this.roleOptions = response.roles
          this.form.postIds = response.postIds !== null ? response.postIds : []
          this.form.role = response.role !== null ? response.role : []
          this.spinning = !this.spinning
        })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
            const saveForm = JSON.parse(JSON.stringify(this.form))
            if (this.form.dept_id !== undefined) {
              saveForm.dept_id = this.form.dept_id.ids
            }
            if (this.form.id !== undefined) {
              updateUser(saveForm).then(response => {
                this.$message.success(
                  '修改成功',
                  3
                )
                this.open = false
                this.$emit('ok')
              })
            } else {
              addUser(saveForm).then(response => {
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
    onLoadData (treeNode) {
          return new Promise(resolve => {
            if (treeNode.dataRef.children) {
              resolve()
              return
            }
            listDeptTree(treeNode.dataRef.id, 1).then(response => {
               treeNode.dataRef.children = response
                resolve()
            })
          })
        },
     checkUserNameUnique (rule, value, callback) {
       const msg = '登陆名称已存在'
       if (value === '') {
         callback()
       } else {
         const checkData = {
           username: value,
           id: this.form.id !== undefined ? this.form.id : ''
         }
         checkUserNameUnique(checkData).then(response => {
           if (response.is_unique) {
             callback()
           } else {
            callback(msg)
           }
         })
       }
     },
     checkEmailUnique (rule, value, callback) {
       const msg = '登陆名称已存在'
       if (value === '') {
         callback()
       } else {
         const checkData = {
           email: value,
           id: this.form.id !== undefined ? this.form.id : ''
         }
         checkEmailUnique(checkData).then(response => {
           if (response.is_unique) {
             callback()
           } else {
            callback(msg)
           }
         })
       }
     },
     checkPhoneUnique (rule, value, callback) {
       const msg = '手机号已存在'
       if (value === '') {
         callback()
       } else {
         const checkData = {
           mobile: value,
           id: this.form.id !== undefined ? this.form.id : ''
         }
         checkPhoneUnique(checkData).then(response => {
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
