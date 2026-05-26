<template>
  <n-grid cols="1" responsive="screen" class="-mt-5">
    <n-grid-item>
      <n-list>
        <n-list-item>
          <template #suffix>
            <n-button type="primary" text @click="openUpdatePassForm">修改</n-button>
          </template>
          <n-thing title="绑定微信">
            <template #description><span class="text-gray-400">已绑定微信号：xxx</span></template>
          </n-thing>
        </n-list-item>
        <n-list-item>
          <template #suffix>
            <n-button type="primary" text @click="openUpdateMobileForm">修改</n-button>
          </template>
          <n-thing title="绑定抖音">
            <template #description><span class="text-gray-400">已绑定抖音号：xxx</span></template>
          </n-thing>
        </n-list-item>
      </n-list>
    </n-grid-item>
  </n-grid>

  <n-modal
    v-model:show="showModal"
    :show-icon="false"
    preset="dialog"
    title="修改登录密码"
    :style="{
      width: dialogWidth,
    }"
  >
    <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
      <n-form-item label="当前密码" path="oldPassword">
        <n-input
          type="password"
          v-model:value="formValue.oldPassword"
          placeholder="请输入当前密码"
        />
      </n-form-item>

      <n-form-item label="新密码" path="newPassword">
        <n-input type="password" v-model:value="formValue.newPassword" placeholder="请输入新密码" />
      </n-form-item>

      <div>
        <n-space justify="end">
          <n-button @click="showModal = false">取消</n-button>
          <n-button type="primary" @click="formSubmit">修改并重新登录</n-button>
        </n-space>
      </div>
    </n-form>
  </n-modal>

  <n-modal
    :block-scroll="false"
    :mask-closable="false"
    v-model:show="showMobileModal"
    :show-icon="false"
    preset="dialog"
    title="修改手机号"
    :style="{
      width: dialogWidth,
    }"
  >
    <n-form :label-width="80" :model="formMobileValue" ref="formMobileRef">
      <n-form-item label="短信验证码" path="code" v-if="userStore.info?.mobile !== ''">
        <n-input-group>
          <n-input v-model:value="formMobileValue.code" placeholder="请输入验证码" />
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

        <template #feedback> 接收号码：+86{{ userStore.info?.mobile }} </template>
      </n-form-item>

      <n-form-item label="换绑手机号" path="mobile">
        <n-input v-model:value="formMobileValue.mobile" placeholder="请输入换绑手机号" />
      </n-form-item>
      <div>
        <n-space justify="end">
          <n-button @click="showMobileModal = false">取消</n-button>
          <n-button type="primary" :loading="formMobileBtnLoading" @click="formMobileSubmit"
            >保存更新</n-button
          >
        </n-space>
      </div>
    </n-form>
  </n-modal>

  <n-modal
    :block-scroll="false"
    :mask-closable="false"
    v-model:show="showEmailModal"
    :show-icon="false"
    preset="dialog"
    title="修改邮箱"
    :style="{
      width: dialogWidth,
    }"
  >
    <n-form :label-width="80" :model="formEmailValue" ref="formEmailRef">
      <n-form-item label="邮箱验证码" path="code" v-if="userStore.info?.email !== ''">
        <n-input-group>
          <n-input v-model:value="formEmailValue.code" placeholder="请输入验证码" />
          <n-button
            type="primary"
            ghost
            @click="sendEmailCode"
            :disabled="isCounting"
            :loading="sendLoading"
          >
            {{ sendLabel }}
          </n-button>
        </n-input-group>
        <template #feedback> 接收邮箱：{{ userStore.info?.email }} </template>
      </n-form-item>

      <n-form-item label="换绑邮箱" path="email">
        <n-input v-model:value="formEmailValue.email" placeholder="请输入换绑邮箱" />
      </n-form-item>
      <div>
        <n-space justify="end">
          <n-button @click="showEmailModal = false">取消</n-button>
          <n-button type="primary" :loading="formEmailBtnLoading" @click="formEmailSubmit"
            >保存更新</n-button
          >
        </n-space>
      </div>
    </n-form>
  </n-modal>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { useRouter, useRoute } from 'vue-router';
  import { useSendCode } from '@/hooks/common';
  import { adaModalWidth } from '@/utils/hotgo';
  import {
    updateMemberPwd,
    updateMemberMobile,
    updateMemberEmail,
    SendBindEmail,
    SendBindSms,
  } from '@/api/system/user';
  import { TABS_ROUTES } from '@/store/mutation-types';
  import { useUserStore } from '@/store/modules/user';

  const { sendLabel, isCounting, loading: sendLoading, activateSend } = useSendCode();
  const userStore = useUserStore();
  const rules = {
    basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },
  };

  const formRef: any = ref(null);
  const message = useMessage();
  const router = useRouter();
  const route = useRoute();
  const showModal = ref(false);
  const formValue = ref({
    oldPassword: '',
    newPassword: '',
  });
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateMemberPwd({
          oldPassword: formValue.value.oldPassword,
          newPassword: formValue.value.newPassword,
        })
          .then((_res) => {
            message.success('更新成功');

            userStore.logout().then(() => {
              message.success('成功注销登录');
              // 移除标签页
              localStorage.removeItem(TABS_ROUTES);
              router
                .replace({
                  name: 'Login',
                  query: {
                    redirect: route.fullPath,
                  },
                })
                .finally(() => location.reload());
            });
          })
          .finally(() => {
            showModal.value = false;
          });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  function openUpdatePassForm() {
    message.error('未开放');
    return;
    // showModal.value = true;
    // formValue.value.newPassword = '';
    // formValue.value.oldPassword = '';
  }

  const formMobileBtnLoading = ref(false);
  const formMobileRef: any = ref(null);
  const showMobileModal = ref(false);
  const formMobileValue = ref({
    mobile: '',
    code: '',
  });

  function formMobileSubmit() {
    formMobileRef.value.validate((errors) => {
      if (!errors) {
        formMobileBtnLoading.value = true;
        updateMemberMobile({
          mobile: formMobileValue.value.mobile,
          code: formMobileValue.value.code,
        })
          .then((_res) => {
            message.success('更新成功');
            showMobileModal.value = false;
            userStore.GetInfo();
          })
          .finally(() => {
            formMobileBtnLoading.value = false;
          });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  function openUpdateMobileForm() {
    message.error('未开放');
    return;
    // showMobileModal.value = true;
    // formMobileValue.value.mobile = '';
    // formMobileValue.value.code = '';
  }

  const formEmailBtnLoading = ref(false);
  const formEmailRef: any = ref(null);
  const showEmailModal = ref(false);
  const formEmailValue = ref({
    email: '',
    code: '',
  });

  function formEmailSubmit() {
    formEmailRef.value.validate((errors) => {
      if (!errors) {
        formEmailBtnLoading.value = true;
        updateMemberEmail({
          email: formEmailValue.value.email,
          code: formEmailValue.value.code,
        })
          .then((_res) => {
            message.success('更新成功');
            showEmailModal.value = false;
            userStore.GetInfo();
          })
          .finally(() => {
            formEmailBtnLoading.value = false;
          });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  function openUpdateEmailForm() {
    showEmailModal.value = true;
    formEmailValue.value.email = '';
    formEmailValue.value.code = '';
  }

  function sendMobileCode() {
    activateSend(SendBindSms());
  }

  function sendEmailCode() {
    activateSend(SendBindEmail());
  }
</script>
