import { getDept, addDept, updateDept, findMaxSort, validateDeptNameUnique, listDeptTree } from '@/api/system/dept'
import AntModal from '@/components/pt/dialog/AntModal'
import SelectDept from '@/components/pt/selectDept/SelectDept'
export default {
  name: 'CreateForm',
  props: {
    deptTypeOptions: {
      type: Array,
      required: true
    },
    statusOptions: {
      type: Array,
      required: true
    },
    deptOptions: {
      type: Array,
      required: true
    },
    expandedKeys: {
      type: Array,
      required: true
    }
  },
  components: {
    AntModal,
    SelectDept
  },
  data () {
    const validateDeptName = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('部门名不允许为空'))
        } else {
          validateDeptNameUnique(this.form.name, this.form.parentId.ids, this.form.id)
          callback()
        }
    }
    const validateParentId = (rule, value, callback) => {
      if (value.ids === '') {
        callback(new Error('上级部门不允许为空'))
      } else {
        callback()
      }
    }
    return {
      loading: false,
      parentIdShow: false,
      hasChild: false,
      selectScope: 'all',
      deptTypeEnableValue: [],
      formTitle: '',
      currentRow: undefined,
      oldParentId: '',
      spinning: false,
      delayTime: 200,
      customStyle: 'background: #fff;ssborder-radius: 4px;margin-bottom: 24px;border: 0;overflow: hidden',
      // 表单参数
      form: {
        id: undefined,
        pid: undefined,
        name: undefined,
        deptFullName: undefined,
        sort: 0,
        deptType: 'dept',
        leader: undefined,
        phone: undefined,
        address: undefined,
        zipCode: undefined,
        email: undefined,
        status: '0'
      },
      open: false,
      rules: {
        pid: [{ required: true, message: '上级部门不能为空', trigger: 'blur', validator: validateParentId }],
        deptType: [{ required: true, message: '部门类型不能为空', trigger: 'blur' }],
        name: [{ required: true, message: '部门名称不能为空', validator: validateDeptName, trigger: 'change' }],
        sort: [{ required: true, message: '排序不能为空', trigger: 'blur' }],
        email: [
          {
            type: 'email',
            message: '请输入正确的邮箱地址',
            trigger: ['blur', 'change']
          }
        ],
        phone: [
          {
            pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
            message: '请输入正确的手机号码',
            trigger: 'blur'
          }
        ]
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
    onDeptTypeChange (item) {
      if (item.target.value === 'company') {
        this.selectScope = 'nonDept'
      } else {
        this.selectScope = 'all'
      }
    },
    onSelectDept (result) {
      this.getDeptTypeEnableValue(result.types)
    },
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
    getDeptTypeEnableValue (parentDeptType) {
      const id = this.form.id
      /* if (parentDeptType === null) {
        this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue === 'org' })
          .map(function (item) { return item.dictValue })
      } else if (this.hasChild) {
        this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue === 'company' })
          .map(function (item) { return item.dictValue })
      } else if (parentDeptType === 'org' || parentDeptType === 'company' || parentDeptType === '') {
        this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue !== 'org' })
          .map(function (item) { return item.dictValue })
      } else {
        this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue === 'dept' })
          .map(function (item) { return item.dictValue })
      } */
      if (parentDeptType === 'org' || parentDeptType === 'company') {
        this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue !== 'org' })
          .map(function (item) { return item.dictValue })
      } else {
        this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue === 'dept' })
          .map(function (item) { return item.dictValue })
      }
      if (id !== null && id !== '' && id !== 'undefined' && id !== undefined) {
        // 编辑页面
        if (parentDeptType === null) {
          this.deptTypeEnableValue = this.deptTypeOptions.filter(function (item) { return item.dictValue === 'org' })
            .map(function (item) { return item.dictValue })
        }
       // 编辑页面当切换上级部门后需要判断当前部门类型是否在可选集合，如果在类型保持不变，如果不在需要重新赋值
       const deptType = this.form.deptType
       const selectDeptType = this.deptTypeEnableValue.filter(function (item) { return item === deptType })
       this.form.deptType = selectDeptType.length === 0 ? this.deptTypeEnableValue[0] : deptType
      } else {
       // 添加页面
       this.form.deptType = this.deptTypeEnableValue[0]
      }
    },
     /** 新增按钮操作  */
    handleAdd (row) {
      this.parentIdShow = true
      this.oldParentId = ''
      this.deptTypeEnableValue = this.deptTypeOptions.map(function (item) { return item.dictValue })
      if (row !== undefined) {
        this.currentRow = row
        this.oldParentId = row.id
        this.form.pid = { ids: row.id, names: row.name }
        this.getDeptTypeEnableValue(row.deptType)
      }
      /** 获取最大编号 */
      findMaxSort(row !== undefined ? row.id : '0').then(response => {
        this.form.treeSort = response
      })
      this.$emit('getTreeselect')
      this.formTitle = '添加部门'
      this.open = true
    },
    setNodeData (data) {
      this.currentRow.name = data.name
      this.currentRow.deptCode = data.deptCode
      this.currentRow.leader = data.leader
      this.currentRow.phone = data.phone
      this.currentRow.email = data.email
      this.currentRow.status = data.status
      this.currentRow.treeSort = data.treeSort
      this.currentRow.createTime = data.createTime
      this.currentRow.deptType = data.deptType
    },
    /** 修改按钮操作 */
    handleUpdate (row) {
      this.currentRow = row
      this.open = true
      this.formTitle = '修改部门'
      this.hasChild = row.children.length > 0
      this.spinning = !this.spinning
      const id = row.id
      getDept(id).then(response => {
        this.oldParentId = response.pid
        this.form = response
        if (response.pid !== '0') {
          this.parentIdShow = true
          this.getDeptTypeEnableValue(response.parentDeptType)
        } else {
          this.parentIdShow = false
          this.getDeptTypeEnableValue(null)
        }
        if (response.deptType === 'company') {
          this.selectScope = 'nonDept'
        }
        this.form.pid = { ids: response.pid, names: response.parentName }
        this.spinning = !this.spinning
      })
       this.$emit('getTreeselect', row)
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          const saveForm = JSON.parse(JSON.stringify(this.form))
          if (this.form.pid !== undefined) {
            saveForm.pid = this.form.pid.ids
          }
          if (this.form.id !== undefined) {
            updateDept(saveForm).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              if (this.oldParentId !== this.form.pid.ids) {
                // 如果修改父节点则刷新树
                this.$emit('ok')
              } else {
                this.setNodeData(response)
              }
              this.cancel()
            })
          } else {
            addDept(saveForm).then(response => {
              this.$message.success(
                '新增成功',
                3
              )
              // 修改父节点后刷新整个树，如果直接添加子节点不更换父节点则追加节点
              if (this.oldParentId !== this.form.pid.ids) {
                // 如果修改父节点则刷新树
                this.$emit('ok')
              } else {
                  this.appendTreeNode(this.currentRow, response)
              }
               this.cancel()
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
        }
  }
}
