<template>
  <a-modal
    :title="title"
    :visible="open"
    @ok="submitForm"
    @cancel="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules">
      <a-form-model-item label="用户名" prop="username">
        <a-input v-model="form.username" :disabled="true"/>
      </a-form-model-item>
      <a-form-model-item has-feedback label="新密码" prop="newPassword">
        <a-input-password v-model="form.newPassword" placeholder="请输入新密码" />
      </a-form-model-item>
      <a-form-model-item has-feedback label="确认密码" prop="confirmPassword">
        <a-input-password v-model="form.confirmPassword" placeholder="请确认密码" />
      </a-form-model-item>
    </a-form-model>
  </a-modal>
</template>
<script>
import { resetUserPwd } from '@/api/system/user'

export default {
  name: 'ResetPassword',
  props: {
  },
  data () {
    const validateNewPass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入新密码'))
      } else if (!/^(?![\d]+$)(?![a-zA-Z]+$)(?![^\da-zA-Z]+$)([^\u4e00-\u9fa5\s]){6,20}$/.test(value)) {
        callback(new Error('请输入6-20位英文字母、数字或者符号（除空格），且字母、数字和标点符号至少包含两种'))
      } else {
        if (this.form.confirmPassword !== '') {
          this.$refs.form.validateField('confirmPassword')
        }
        callback()
      }
    }
    const validateConfirmPass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入新密码确认'))
      } else if (value !== this.form.newPassword) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    }
    return {
      title: '重置密码',
      open: false,
      childrenDrawer: false,
      formLayout: 'horizontal',
      form: {
        username: undefined,
        newPassword: undefined,
        confirmPassword: undefined
      },
      rules: {
        newPassword: [
          { required: true, validator: validateNewPass, trigger: 'change' }
        ],
        confirmPassword: [
          { required: true, validator: validateConfirmPass, trigger: 'change' }
        ]
      }
    }
  },
  methods: {
    // 取消按钮
    cancel () {
      this.open = false
      this.$emit('close')
    },
    // 表单重置
    reset () {
      this.form = {
        id: undefined,
        username: undefined,
        newPassword: undefined,
        confirmPassword: undefined
      }
    },
    handleResetPwd (row) {
      this.form = {
        id: row.id,
        username: row.username
      }
      this.open = true
    },
    /** 重置密码按钮操作 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          resetUserPwd(this.form.id, this.form.newPassword).then(response => {
            this.$message.success(
              '重置成功',
              3
            )
            this.open = false
            this.$emit('close')
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>
