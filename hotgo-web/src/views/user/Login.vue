<template>
  <div class="main">
    <a-icon class="QRcodeIcon" :component="allIcon.QRcodeIcon"/>
    <a-form-model id="formLogin" ref="form" class="user-layout-login" :model="form" :rules="rules">
      <a-page-header title="用户登录"/>
      <a-form-model-item prop="username">
        <a-input v-model="form.username" size="large" placeholder="用户名">
          <a-icon slot="prefix" type="user" :style="{ color: 'rgba(0,0,0,.25)'}"/>
        </a-input>
      </a-form-model-item>
      <a-form-model-item prop="password">
        <a-input-password v-model="form.password" size="large" placeholder="密码">
          <a-icon slot="prefix" type="lock" :style="{ color: 'rgba(0,0,0,.25)' }"/>
        </a-input-password>
      </a-form-model-item>
      <a-row :gutter="16" v-if="captchaOnOff">
        <a-col class="gutter-row" :span="16">
          <a-form-model-item prop="code">
            <a-input v-model="form.code" size="large" type="text" autocomplete="off" placeholder="验证码">
              <a-icon slot="prefix" type="security-scan" :style="{ color: 'rgba(0,0,0,.25)' }"/>
            </a-input>
          </a-form-model-item>
        </a-col>
        <a-col class="gutter-row" :span="8">
          <img class="getCaptcha" :src="codeUrl" @click="getCode">
        </a-col>
      </a-row>
      <a-form-item>
        <a-checkbox v-decorator="['rememberMe', { valuePropName: 'checked' }]">记住密码</a-checkbox>
        <div style="float: right;line-height: 30px;">
          还没有账号？
          <a-button type="link">立即注册</a-button>
          <a-button type="link">忘记密码</a-button>
        </div>
      </a-form-item>
      <a-form-item style="margin-top:24px">
        <a-button
          size="large"
          type="primary"
          htmlType="submit"
          class="login-button"
          :loading="logining"
          :disabled="logining"
          @click="handleSubmit"
        >确定
        </a-button>
      </a-form-item>
      <a-space>
        <div style="margin-top: -10px;">其它登录方式</div>
        <div class="icons-list">
          <a-icon class="dingtalk" :component="allIcon.dingtalkIcon"/>
          <a-icon class="WeChat" :component="allIcon.WeChatIcon"/>
          <a-icon class="Alipay" :component="allIcon.AlipayIcon"/>
          <a-icon class="Sina" :component="allIcon.SinaIcon"/>
          <a-button type="link" @click="applyLicense">授权申请</a-button>
        </div>
      </a-space>
    </a-form-model>
  </div>
</template>

<script>

  // import md5 from 'md5'
  import { mapActions } from 'vuex'
  import { timeFix } from '@/utils/util'
  import { getCodeImg } from '@/api/login'
  import allIcon from '@/core/icons'

  export default {
    components: {
      allIcon
    },
    data() {
      return {
        allIcon,
        codeUrl: '',
        form: {
          username: 'admin',
          password: '123456',
          code: 1, // undefined
          cid: ''
        },
        // 验证码开关
        captchaOnOff: true,
        rules: {
          username: [{ required: true, message: '请输入帐户名', trigger: 'blur' }],
          password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
          code: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
          cid: [{ required: true, message: '验证码ID不能为空', trigger: 'blur' }]
        },
        logining: false
      }
    },
    created() {
      this.getCode()
    },
    methods: {
      getCode() {
        getCodeImg().then(res => {
          this.captchaOnOff = res.captchaOnOff === undefined ? true : res.captchaOnOff
          if (this.captchaOnOff) {
            this.codeUrl = res.base64
            this.form.cid = res.cid
          }
        })
      },
      applyLicense() {
        window.open('/applyLicense', '_blank')
      },
      ...mapActions(['Login', 'Logout']),
      handleSubmit() {
        // 关闭登录过期的提示框
        this.$notification.close('loginExpireTip')
        this.logining = true
        this.$refs.form.validate(valid => {
          if (valid) {
            this.Login(this.form)
              .then((res) => this.loginSuccess(res))
              .catch(err => this.requestFailed(err))
              .finally(() => {
                this.logining = false
              })
          } else {
            setTimeout(() => {
              this.logining = false
            }, 600)
            if (this.captchaOnOff) {
              this.getCode()
            }
          }
        })
      },
      loginSuccess(res) {
        this.$router.push({ path: '/' })
        // 延迟 1 秒显示欢迎信息
        setTimeout(() => {
          this.$notification.success({
            message: '欢迎',
            description: `${timeFix()}，欢迎回来`
          })
        }, 1000)
      },
      requestFailed(err) {
        // this.isLoginError = true
        // this.loginErrorInfo = err
        this.form.code = undefined
        this.getCode()
        this.$message.error(err)
      },
      handleCloseLoginError() {
        // this.isLoginError = false
        // this.loginErrorInfo = ''
      }
    }
  }
</script>

<style lang="less" scoped>
  .QRcodeIcon {
    position: absolute;
    right: 0;
    font-size: 32px;
    color: #1890ff;
    margin-top: 5px;
    margin-right: 5px;
  }

  .ant-space-align-center {
    color: #8f959e;
    line-height: 30px;
    height: 30px;
  }

  .user-layout-login label {
    font-size: 12px !important;
  }

  .user-layout-login {
    label {
      font-size: 14px;
    }

    .ant-page-header {
      padding: 60px 0 45px 0;
    }

    .getCaptcha {
      display: block;
      width: 100%;
      height: 40px;
    }

    button.login-button {
      padding: 0 15px;
      font-size: 16px;
      height: 40px;
      width: 100%;
      border-radius: 4px;
    }

    .icons-list {
      .anticon {
        font-size: 30px;
        width: 40px;
      }

      .dingtalk {
        color: #0089FF;
      }

      .WeChat {
        color: #51C332;
      }

      .Alipay {
        color: #06B4FD;
      }

      .Sina {
        color: #D81E06;
      }
    }
  }

</style>
