<template>
  <n-form
    ref="formRef"
    label-placement="left"
    size="large"
    :model="mode === 'account' ? formInline : formMobile"
    :rules="mode === 'account' ? rules : mobileRules"
  >
    <template v-if="mode === 'account'">
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
          show-password-on="click"
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

          <n-loading-bar-provider :to="loadingBarTargetRef" container-style="position: absolute;">
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
    </template>

    <template v-if="mode === 'mobile'">
      <n-form-item path="mobile">
        <n-input
          @keyup.enter="handleMobileSubmit"
          v-model:value="formMobile.mobile"
          placeholder="请输入手机号码"
        >
          <template #prefix>
            <n-icon size="18" color="#808695">
              <MobileOutlined />
            </n-icon>
          </template>
        </n-input>
      </n-form-item>

      <n-form-item path="code">
        <n-input-group>
          <n-input
            @keyup.enter="handleMobileSubmit"
            v-model:value="formMobile.code"
            placeholder="请输入验证码"
          >
            <template #prefix>
              <n-icon size="18" color="#808695" :component="SafetyCertificateOutlined" />
            </template>
          </n-input>
          <n-button
            type="primary"
            ghost
            @click="sendMobileCode"
            :disabled="isCounting"
            :loading="sendLoading"
          >
            {{ sendLabel }}
          </n-button>
        </n-input-group>
      </n-form-item>
    </template>

    <n-space :vertical="true" :size="24">
      <div class="flex-y-center justify-between">
        <n-checkbox v-model:checked="autoLogin">自动登录</n-checkbox>
        <n-button :text="true" @click="handleResetPassword">忘记密码？</n-button>
      </div>
      <n-button type="primary" size="large" :block="true" :loading="loading" @click="handleLogin">
        登录
      </n-button>

      <FormOther moduleKey="register" tag="注册账号" @updateActiveModule="updateActiveModule" />
    </n-space>

    <DemoAccount @login="handleDemoAccountLogin" />
  </n-form>
</template>

<script lang="ts" setup>
  import '../components/style.less';
  import { ref, onMounted } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useUserStore } from '@/store/modules/user';
  import { useMessage, useLoadingBar } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { PersonOutline, LockClosedOutline } from '@vicons/ionicons5';
  import { PageEnum } from '@/enums/pageEnum';
  import { SafetyCertificateOutlined, MobileOutlined } from '@vicons/antd';
  import { GetCaptcha } from '@/api/base';
  import { aesEcb } from '@/utils/encrypt';
  import DemoAccount from './demo-account.vue';
  import FormOther from '../components/form-other.vue';
  import { useSendCode } from '@/hooks/common';
  import { SendSms } from '@/api/system/user';
  import { validate } from '@/utils/validateUtil';

  interface Props {
    mode: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    mode: 'account',
  });

  interface FormState {
    username: string;
    pass: string;
    cid: string;
    code: string;
    password: string;
  }

  interface FormMobileState {
    mobile: string;
    code: string;
  }

  const formRef = ref();
  const message = useMessage();
  const loading = ref(false);
  const autoLogin = ref(true);
  const codeBase64 = ref('');
  const loadingBar = useLoadingBar();
  const loadingBarTargetRef = ref<undefined | HTMLElement>(undefined);
  const userStore = useUserStore();
  const router = useRouter();
  const route = useRoute();
  const { sendLabel, isCounting, loading: sendLoading, activateSend } = useSendCode();
  const emit = defineEmits(['updateActiveModule']);
  const LOGIN_NAME = PageEnum.BASE_LOGIN_NAME;

  const formInline = ref<FormState>({
    username: '',
    pass: '',
    cid: '',
    code: '',
    password: '',
  });

  const formMobile = ref<FormMobileState>({
    mobile: '',
    code: '',
  });

  const rules = {
    username: { required: true, message: '请输入用户名', trigger: 'blur' },
    pass: { required: true, message: '请输入密码', trigger: 'blur' },
  };

  const mobileRules = {
    mobile: { required: true, message: '请输入手机号码', trigger: 'blur' },
    code: { required: true, message: '请输入验证码', trigger: 'blur' },
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    formRef.value.validate(async (errors) => {
      if (!errors) {
        if (userStore.loginConfig?.loginCaptchaSwitch === 1 && formInline.value.code === '') {
          message.error('请输入验证码');
          return;
        }

        const params = {
          username: formInline.value.username,
          password: aesEcb.encrypt(formInline.value.pass),
          cid: formInline.value.cid,
          code: formInline.value.code,
        };
        await handleLoginResp(userStore.login(params));
      } else {
        message.error('请填写完整信息，并且进行验证码校验');
      }
    });
  };

  async function refreshCode() {
    if (userStore.loginConfig?.loginCaptchaSwitch !== 1) {
      return;
    }
    loadingBar.start();
    const data = await GetCaptcha();
    codeBase64.value = data.base64;
    formInline.value.cid = data.cid;
    formInline.value.code = '';
    loadingBar.finish();
  }

  async function handleDemoAccountLogin(user: { username: string; password: string }) {
    const params = {
      username: user.username,
      password: aesEcb.encrypt(user.password),
      isLock: true,
    };
    await handleLoginResp(userStore.login(params));
  }

  const handleMobileSubmit = (e) => {
    e.preventDefault();
    formRef.value.validate(async (errors) => {
      if (!errors) {
        const params = {
          mobile: formMobile.value.mobile,
          code: formMobile.value.code,
        };
        await handleLoginResp(userStore.mobileLogin(params));
      } else {
        message.error('请填写完整信息，并且进行验证码校验');
      }
    });
  };

  function updateActiveModule(key: string) {
    emit('updateActiveModule', key);
  }

  function sendMobileCode() {
    validate.phone(mobileRules.mobile, formMobile.value.mobile, function (error?: Error) {
      if (error === undefined) {
        activateSend(SendSms({ mobile: formMobile.value.mobile, event: 'login' }));
        return;
      }
      message.error(error.message);
    });
  }

  function handleResetPassword() {
    message.info('如果你忘记了密码，请联系管理员找回');
  }

  function handleLogin(e) {
    if (props.mode === 'account') {
      handleSubmit(e);
      return;
    }

    handleMobileSubmit(e);
  }

  async function handleLoginResp(request: Promise<any>) {
    message.loading('登录中...');
    loading.value = true;
    try {
      const { code, message: msg } = await request;
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        const toPath = decodeURIComponent((route.query?.redirect || '/') as string);
        message.success('登录成功，即将进入系统');
        if (route.name === LOGIN_NAME) {
          await router.replace('/');
        } else await router.replace(toPath);
      } else {
        message.destroyAll();
        message.info(msg || '登录失败');
        await refreshCode();
      }
    } finally {
      loading.value = false;
    }
  }

  onMounted(() => {
    setTimeout(function () {
      refreshCode();
    });
  });
</script>
