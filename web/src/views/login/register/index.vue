<template>
  <n-form ref="formRef" label-placement="left" size="large" :model="formInline" :rules="rules">
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
        placeholder="请输入密码"
        show-password-on="click"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <LockClosedOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>

    <n-form-item path="confirmPwd">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.confirmPwd"
        type="password"
        placeholder="再次输入密码"
        show-password-on="click"
      >
        <template #prefix>
          <n-icon size="18" color="#808695">
            <LockClosedOutline />
          </n-icon>
        </template>
      </n-input>
    </n-form-item>

    <n-form-item path="mobile">
      <n-input
        @keyup.enter="handleSubmit"
        v-model:value="formInline.mobile"
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
          @keyup.enter="handleSubmit"
          v-model:value="formInline.code"
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

    <n-form-item path="inviteCode">
      <n-input
        :style="{ width: '100%' }"
        placeholder="邀请码(选填)"
        @keyup.enter="handleSubmit"
        v-model:value="formInline.inviteCode"
        :disabled="inviteCodeDisabled"
      >
        <template #prefix>
          <n-icon size="18" color="#808695" :component="TagOutlined" />
        </template>
      </n-input>
    </n-form-item>

    <n-form-item class="default-color">
      <Agreement
        v-model:value="agreement"
        @clickProtocol="handleClickProtocol"
        @clickPolicy="handleClickPolicy"
      />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="handleSubmit" size="large" :loading="loading" block>
        注册
      </n-button>
    </n-form-item>

    <FormOther moduleKey="login" tag="登录账号" @updateActiveModule="updateActiveModule" />
  </n-form>

  <n-modal
    v-model:show="showModal"
    :show-icon="false"
    :mask-closable="false"
    preset="dialog"
    :closable="false"
    :style="{
      width: dialogWidth,
      position: 'top',
      bottom: '15vw',
    }"
  >
    <n-space justify="center">
      <div class="agree-title">《{{ agreeTitle }}》</div>
    </n-space>

    <div v-html="modalContent"></div>

    <n-divider />
    <n-space justify="center">
      <n-button type="info" ghost strong @click="handleAgreement(true)">我已知晓并接受</n-button>
      <n-button type="error" ghost strong @click="handleAgreement(false)">我拒绝</n-button>
    </n-space>
  </n-modal>
</template>

<script lang="ts" setup>
  import '../components/style.less';
  import { ref, onMounted, computed } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { PersonOutline, LockClosedOutline } from '@vicons/ionicons5';
  import { SafetyCertificateOutlined, MobileOutlined, TagOutlined } from '@vicons/antd';
  import { aesEcb } from '@/utils/encrypt';
  import Agreement from './agreement.vue';
  import FormOther from '../components/form-other.vue';
  import { useSendCode } from '@/hooks/common';
  import { validate } from '@/utils/validateUtil';
  import { register, SendSms } from '@/api/system/user';
  import { useUserStore } from '@/store/modules/user';
  import { adaModalWidth } from '@/utils/hotgo';

  interface FormState {
    username: string;
    pass: string;
    confirmPwd: string;
    mobile: string;
    code: string;
    inviteCode: string;
    password: string;
  }

  const emit = defineEmits(['updateActiveModule']);
  const formRef = ref();
  const router = useRouter();
  const message = useMessage();
  const userStore = useUserStore();
  const loading = ref(false);
  const showModal = ref(false);
  const agreeTitle = ref('');
  const modalContent = ref('');
  const { sendLabel, isCounting, loading: sendLoading, activateSend } = useSendCode();
  const agreement = ref(false);
  const inviteCodeDisabled = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth();
  });

  const formInline = ref<FormState>({
    username: '',
    pass: '',
    confirmPwd: '',
    mobile: '',
    code: '',
    inviteCode: '',
    password: '',
  });

  const rules = {
    username: { required: true, message: '请输入用户名', trigger: 'blur' },
    pass: { required: true, message: '请输入密码', trigger: 'blur' },
    mobile: { required: true, message: '请输入手机号码', trigger: 'blur' },
    code: { required: true, message: '请输入验证码', trigger: 'blur' },
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    formRef.value.validate(async (errors) => {
      if (!errors) {
        if (formInline.value.pass !== formInline.value.confirmPwd) {
          message.info('两次输入的密码不一致，请检查');
          return;
        }

        if (!agreement.value) {
          message.info('请确认你已经仔细阅读并接受《用户协议》和《隐私权政策》并已勾选接受选项');
          return;
        }

        message.loading('注册中...');
        loading.value = true;

        try {
          const { code, message: msg } = await register({
            username: formInline.value.username,
            password: aesEcb.encrypt(formInline.value.pass),
            mobile: formInline.value.mobile,
            code: formInline.value.code,
            inviteCode: formInline.value.inviteCode,
          });
          message.destroyAll();
          if (code == ResultEnum.SUCCESS) {
            message.success('注册成功，请登录！');
            updateActiveModule('login');
          } else {
            message.info(msg || '注册失败');
          }
        } finally {
          loading.value = false;
        }
      } else {
        message.error('请填写完整信息，并且进行验证码校验');
      }
    });
  };

  onMounted(() => {
    const inviteCode = router.currentRoute.value.query?.inviteCode as string;
    if (inviteCode) {
      inviteCodeDisabled.value = true;
      formInline.value.inviteCode = inviteCode;
    }
  });

  function updateActiveModule(key: string) {
    emit('updateActiveModule', key);
  }

  function sendMobileCode() {
    validate.phone(rules.mobile, formInline.value.mobile, function (error?: Error) {
      if (error === undefined) {
        activateSend(SendSms({ mobile: formInline.value.mobile, event: 'register' }));
        return;
      }
      message.error(error.message);
    });
  }

  function handleClickProtocol() {
    showModal.value = true;
    agreeTitle.value = '用户协议';
    modalContent.value = userStore.loginConfig?.loginProtocol as string;
  }

  function handleClickPolicy() {
    showModal.value = true;
    agreeTitle.value = '隐私权政策';
    modalContent.value = userStore.loginConfig?.loginPolicy as string;
  }

  function handleAgreement(agree: boolean) {
    showModal.value = false;
    agreement.value = agree;
  }
</script>

<style lang="less" scoped>
  .agree-title {
    font-size: 18px;
    margin-bottom: 22px;
  }
</style>
