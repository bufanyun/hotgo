<template>
  <div class="view-account">
    <div class="view-account-header"></div>
    <div class="view-account-container">
      <div class="view-account-top">
        <div class="view-account-top-logo">
          <img src="~@/assets/images/account-logo.png" alt="" />
        </div>
        <div class="view-account-top-desc">HotGo 后台管理系统</div>
      </div>
      <div class="view-account-form">
        <n-form
          ref="formRef"
          label-placement="left"
          size="large"
          :model="formInline"
          :rules="rules"
        >
          <n-form-item path="username">
            <n-input
              @keyup.enter="handleSubmit"
              v-model:value="formInline.username"
              placeholder="请输入用户名"
            >
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <PersonOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item path="pass">
            <n-input
              @keyup.enter="handleSubmit"
              v-model:value="formInline.pass"
              type="password"
              showpassOn="click"
              placeholder="请输入密码"
            >
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <LockClosedOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item path="code" v-show="codeBase64 !== ''">
            <n-input-group>
              <n-input
                :style="{ width: '100%' }"
                placeholder="验证码"
                @keyup.enter="handleSubmit"
                v-model:value="formInline.code"
              >
                <template #prefix>
                  <n-icon size="18" color="#808695" :component="SafetyCertificateOutlined" />
                </template>
                <template #suffix> </template>
              </n-input>

              <n-loading-bar-provider
                :to="loadingBarTargetRef"
                container-style="position: absolute;"
              >
                <img
                  ref="loadingBarTargetRef"
                  style="width: 100px"
                  :src="codeBase64"
                  @click="refreshCode"
                  loading="lazy"
                  alt="点击获取"
                />
                <loading-bar-trigger />
              </n-loading-bar-provider>
            </n-input-group>
          </n-form-item>
          <n-form-item class="default-color">
            <div class="flex justify-between">
              <div class="flex-initial">
                <n-checkbox v-model:checked="autoLogin">自动登录</n-checkbox>
              </div>
              <div class="flex-initial order-last">
                <a href="javascript:">忘记密码</a>
              </div>
            </div>
          </n-form-item>
          <n-form-item>
            <n-button type="primary" @click="handleSubmit" size="large" :loading="loading" block>
              登录
            </n-button>
          </n-form-item>
          <n-form-item class="default-color">
            <div class="flex view-account-other">
              <div class="flex-initial">
                <span>其它登录方式</span>
              </div>
              <div class="flex-initial mx-2">
                <a href="javascript:">
                  <n-icon size="24" color="#2d8cf0">
                    <LogoWechat />
                  </n-icon>
                </a>
              </div>
              <div class="flex-initial mx-2">
                <a href="javascript:">
                  <n-icon size="24" color="#2d8cf0">
                    <LogoTiktok />
                  </n-icon>
                </a>
              </div>
              <div class="flex-initial" style="margin-left: auto">
                <a @click="handleRegister">注册账号</a>
              </div>
            </div>
          </n-form-item>
        </n-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useUserStore } from '@/store/modules/user';
  import { useMessage, useLoadingBar } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { PersonOutline, LockClosedOutline, LogoWechat, LogoTiktok } from '@vicons/ionicons5';
  import { PageEnum } from '@/enums/pageEnum';
  import { SafetyCertificateOutlined } from '@vicons/antd';
  import { GetCaptcha } from '@/api/base';
  import { aesEcb } from '@/utils/encrypt';

  interface FormState {
    username: string;
    pass: string;
    cid: string;
    code: string;
    password: string;
  }

  const formRef = ref();
  const message = useMessage();
  const loading = ref(false);
  const autoLogin = ref(true);
  const codeBase64 = ref('');
  const loadingBar = useLoadingBar();
  const loadingBarTargetRef = ref<undefined | HTMLElement>(undefined);
  const LOGIN_NAME = PageEnum.BASE_LOGIN_NAME;

  const formInline = ref<FormState>({
    username: '',
    pass: '',
    cid: '',
    code: '',
    password: '',
  });

  const rules = {
    username: { required: true, message: '请输入用户名', trigger: 'blur' },
    pass: { required: true, message: '请输入密码', trigger: 'blur' },
    code: { required: true, message: '请输入验证码', trigger: 'blur' },
  };

  const userStore = useUserStore();
  const router = useRouter();
  const route = useRoute();

  const handleSubmit = (e) => {
    e.preventDefault();
    formRef.value.validate(async (errors) => {
      if (!errors) {
        message.loading('登录中...');
        loading.value = true;
        try {
          const { code, message: msg } = await userStore.login({
            username: formInline.value.username,
            password: aesEcb.encrypt(formInline.value.pass),
            cid: formInline.value.cid,
            code: formInline.value.code,
          });
          message.destroyAll();
          if (code == ResultEnum.SUCCESS) {
            const toPath = decodeURIComponent((route.query?.redirect || '/') as string);
            message.success('登录成功，即将进入系统');
            if (route.name === LOGIN_NAME) {
              await router.replace('/');
            } else await router.replace(toPath);
          } else {
            message.info(msg || '登录失败');
            await refreshCode();
          }
        } finally {
          loading.value = false;
        }
      } else {
        message.error('请填写完整信息，并且进行验证码校验');
      }
    });
  };

  async function refreshCode() {
    loadingBar.start();
    const data = await GetCaptcha();
    codeBase64.value = data.base64;
    formInline.value.cid = data.cid;
    formInline.value.code = '';
    loadingBar.finish();
  }

  onMounted(() => {
    setTimeout(function () {
      refreshCode();
    });
    console.log('window.location.href',route.path);
  });

  function handleRegister() {
    message.success('即将开放，请稍后');
    return;
  }
</script>

<style lang="less" scoped>
  .view-account {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: auto;

    &-container {
      flex: 1;
      padding: 32px 12px;
      max-width: 384px;
      min-width: 320px;
      margin: 0 auto;
    }

    &-top {
      padding: 32px 0;
      text-align: center;

      &-desc {
        font-size: 14px;
        color: #808695;
      }
    }

    &-other {
      width: 100%;
    }

    .default-color {
      color: #515a6e;

      .ant-checkbox-wrapper {
        color: #515a6e;
      }
    }
  }

  @media (min-width: 768px) {
    .view-account {
      background-image: url('../../assets/images/login.svg');
      background-repeat: no-repeat;
      background-position: 50%;
      background-size: 100%;
    }

    .page-account-container {
      padding: 32px 0 24px 0;
    }
  }
</style>
