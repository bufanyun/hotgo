import { getRole, addRole, updateRole, checkRoleNameUnique, checkRoleKeyUnique, findMaxSort } from '@/api/system/role'
import { treeselect as menuTreeselect, roleMenuTreeselect } from '@/api/system/menu'
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
      spinning: false,
      delayTime: 200,
      menuExpandedKeys: [],
      autoExpandParent: false,
      menuCheckedKeys: [],
      halfCheckedKeys: [],
      menuOptions: [],
      formTitle: '',
      // 表单参数
      form: {
        roleId: undefined,
        roleName: undefined,
        roleKey: undefined,
        sort: 0,
        status: '0',
        menuIds: [],
        menuCheckStrictly: true,
        remark: undefined
      },
      open: false,
      menuExpand: false,
      menuNodeAll: false,
      rules: {
        roleName: [{ required: true, message: '角色名称不能为空', trigger: 'blur' },
        { validator: this.checkRoleNameUnique }],
        roleKey: [{ required: true, message: '角色编码不能为空', trigger: 'blur' },
        { validator: this.checkRoleKeyUnique }],
        sort: [{ required: true, message: '显示顺序不能为空', trigger: 'blur' }]
      },
      defaultProps: {
        children: 'children',
        title: 'label',
        key: 'id'
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
    onExpandMenu (expandedKeys) {
      this.menuExpandedKeys = expandedKeys
      this.autoExpandParent = false
    },
    /** 查询菜单树结构 */
    getMenuTreeselect () {
      return menuTreeselect(0, 10).then(response => {
        return response
      })
    },
    // 所有菜单节点数据
    getMenuAllCheckedKeys () {
      // 全选与半选
      return Array.from(new Set(this.menuCheckedKeys.concat(this.halfCheckedKeys)))
    },
    getAllMenuNode (nodes) {
      if (!nodes || nodes.length === 0) {
        return []
      }
      nodes.forEach(node => {
        this.menuCheckedKeys.push(node.id)
        return this.getAllMenuNode(node.children)
      })
    },
    // 回显过滤
    selectNodefilter (nodes, parentIds) {
      if (!nodes || nodes.length === 0) {
        return []
      }
      nodes.forEach(node => {
        // 父子关联模式且当前元素有父级
        const currentIndex = this.menuCheckedKeys.indexOf(node.id)
        // 当前节点存在,且父节点不存在，则说明父节点应是半选中状态
        if (currentIndex !== -1) {
          parentIds.forEach(parentId => {
            if (this.halfCheckedKeys.indexOf(parentId) === -1) {
              this.halfCheckedKeys.push(parentId)
            }
          })
          parentIds = []
        }
        // 防重
        const isExist = this.halfCheckedKeys.indexOf(node.id)
        const isExistParentIds = parentIds.indexOf(node.id)
        if (isExist === -1 && isExistParentIds === -1 && currentIndex === -1) {
          parentIds.push(node.id)
        }
        return this.selectNodefilter(node.children, parentIds)
      })
    },
    handleCheckedTreeNodeAll (value) {
      this.menuNodeAll = !this.menuNodeAll
      if (value.target.checked) {
        this.getAllMenuNode(this.menuOptions)
      } else {
        this.menuCheckedKeys = []
        this.halfCheckedKeys = []
      }
    },
    handleCheckedTreeExpand (value) {
      this.menuExpand = !this.menuExpand
      if (value.target.checked) {
        const treeList = this.menuOptions
        this.treeExpandWithLevel(treeList, -1)
      } else {
        this.menuExpandedKeys = []
        this.treeExpandWithLevel(this.menuOptions, 1)
      }
    },
    treeExpandWithLevel (treeNodeList, level) {
      level--
      if (level !== 0) {
        treeNodeList.forEach(node => {
          this.menuExpandedKeys.push(node.id)
          if (node.children) {
            this.treeExpandWithLevel(node.children, level)
          }
        })
      }
    },
    // 树权限（父子联动）
    handleCheckedTreeConnect (value) {
      this.form.menuCheckStrictly = !this.form.menuCheckStrictly
    },
    /** 根据角色ID查询菜单树结构 */
    getRoleMenuTreeselect (roleId) {
      return roleMenuTreeselect(roleId).then(response => {
        return response
      })
    },
    onCheck (checkedKeys, info) {
      if (!this.form.menuCheckStrictly) {
        let currentCheckedKeys = []
        if (this.menuCheckedKeys.checked) {
          currentCheckedKeys = Array.from(new Set(currentCheckedKeys.concat(this.menuCheckedKeys.checked)))
        }
        if (this.menuCheckedKeys.halfChecked) {
          currentCheckedKeys = Array.from(new Set(currentCheckedKeys.concat(this.menuCheckedKeys.halfChecked)))
        }
        this.menuCheckedKeys = currentCheckedKeys
      } else {
        // 半选节点
        this.halfCheckedKeys = info.halfCheckedKeys
        this.menuCheckedKeys = checkedKeys
      }
    },
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
         this.form.sort = response.data
       })
      const roleMenu = this.getMenuTreeselect()
      roleMenu.then(res => {
        this.menuOptions = res.data
        this.treeExpandWithLevel(this.menuOptions, 1)
      })
      this.open = true
      this.formTitle = '添加角色'
    },
    /** 修改按钮操作 */
    handleUpdate (row, ids) {
      this.open = true
      this.spinning = !this.spinning
      this.reset()
      this.menuExpand = false
      this.menuNodeAll = false
      const roleId = row ? row.id : ids
      const roleMenu = this.getRoleMenuTreeselect(roleId)
      getRole(roleId).then(response => {
        this.form = response.data
      })
      roleMenu.then(res => {
        this.menuOptions = res.data.menus
        this.menuCheckedKeys = res.data.checkedKeys
        // 过滤回显时的半选中node(父)
        if (this.form.menuCheckStrictly) {
          this.selectNodefilter(this.menuOptions, [])
        }
        this.treeExpandWithLevel(this.menuOptions, 1)
      })
      this.formTitle = '修改角色'
      this.spinning = !this.spinning
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
           if (this.form.id !== undefined) {
             this.form.menuIds = this.getMenuAllCheckedKeys()
             updateRole(this.form).then(response => {
               this.$message.success(
                 '修改成功',
                 3
               )
               this.open = false
               this.$emit('ok')
             })
           } else {
             this.form.menuIds = this.getMenuAllCheckedKeys()
             addRole(this.form).then(response => {
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
    checkRoleNameUnique (rule, value, callback) {
      const msg = '角色名称已存在'
      if (value === '') {
        callback()
      } else {
        const checkData = {
          roleName: value,
          id: this.form.id !== undefined ? this.form.id : ''
        }
        checkRoleNameUnique(checkData).then(response => {
          if (response.data.code === '1') {
            callback()
          } else {
           callback(msg)
          }
        })
      }
    },
    checkRoleKeyUnique (rule, value, callback) {
       const msg = '角色编码已存在'
      if (value === '') {
        callback()
      } else {
        const checkData = {
          roleKey: value,
          id: this.form.id !== undefined ? this.form.id : ''
        }
        checkRoleKeyUnique(checkData).then(response => {
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
