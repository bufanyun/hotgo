<template>
  <a-modal
    :title="title"
    :visible="open"
    @ok="submitForm"
    @cancel="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules">
      <a-form-model-item has-feedback label="旧密码" prop="oldPassword">
        <a-input-password v-model="form.oldPassword" placeholder="请输入旧密码" />
      </a-form-model-item>
      <a-form-model-item has-feedback label="新密码" prop="newPassword">
        <a-input-password v-model="form.newPassword" placeholder="请输入新密码" />
      </a-form-model-item>
      <a-form-model-item has-feedback label="确认密码" prop="confirmPassword">
        <a-input-password v-model="form.confirmPassword" placeholder="请确认密码" />
      </a-form-model-item>
    <!-- <a-form-model ref="ruleForm" :model="ruleForm" :rules="rules" v-bind="layout">
      <a-form-model-item has-feedback label="Password" prop="pass">
        <a-input v-model="ruleForm.pass" type="password" autocomplete="off" />
      </a-form-model-item>
      <a-form-model-item has-feedback label="Confirm" prop="checkPass">
        <a-input v-model="ruleForm.checkPass" type="password" autocomplete="off" />
      </a-form-model-item>
      <a-form-model-item has-feedback label="Age" prop="age">
        <a-input v-model.number="ruleForm.age" />
      </a-form-model-item> -->
    </a-form-model>
  </a-modal>
</template>
<script>
import { updateUserPwd } from '@/api/system/user'

export default {
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
        oldPassword: undefined,
        newPassword: undefined,
        confirmPassword: undefined
      },
      rules: {
        oldPassword: [
          { required: true, message: '密码不能为空', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, validator: validateNewPass, trigger: 'change' }
        ],
        confirmPassword: [
          { required: true, validator: validateConfirmPass, trigger: 'change' }
        ]
      },
      layout: {
        labelCol: { span: 4 },
        wrapperCol: { span: 14 }
      }
    }
  },
  methods: {
    // 取消按钮
    cancel () {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset () {
      this.form = {
        oldPassword: undefined,
        newPassword: undefined,
        confirmPassword: undefined
      }
    },
    handleUpdatePwd () {
      this.open = true
    },
    /** 重置密码按钮操作 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          updateUserPwd(this.form.oldPassword, this.form.newPassword).then(response => {
            this.$message.success(
              '修改成功',
              3
            )
            this.open = false
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>
