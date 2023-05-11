<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="支付宝姓名" path="name">
          <n-input v-model:value="formValue.name" />
        </n-form-item>

        <n-form-item label="支付宝账号" path="account ">
          <n-input v-model:value="formValue.account" />
        </n-form-item>

        <n-form-item label="支付宝收款码" path="payeeCode">
          <UploadImage
            :maxNumber="1"
            :helpText="'请上传清晰有效的收款码，图片大小不超过2M'"
            v-model:value="formValue.payeeCode"
          />
        </n-form-item>

        <n-form-item label="登录密码" path="password">
          <n-input
            type="password"
            v-model:value="formValue.password"
            placeholder="请输入登录密码验证身份"
          />
        </n-form-item>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import { getUserInfo, updateMemberCash } from '@/api/system/user';

  const show = ref(false);

  const rules = {
    password: {
      required: true,
      message: '请输入登录密码',
      trigger: 'blur',
    },
  };

  const formRef: any = ref(null);
  const message = useMessage();
  const formValue = ref({
    password: '',
    payeeCode: '',
    account: '',
    name: '',
  });

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateMemberCash({
          name: formValue.value.name,
          account: formValue.value.account,
          payeeCode: formValue.value.payeeCode,
          password: formValue.value.password,
        })
          .then((_res) => {
            message.success('更新成功');
            load();
          })
          .finally(() => {});
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    getUserInfo()
      .then((res) => {
        res.cash.password = '';
        formValue.value = res.cash;
      })
      .finally(() => {
        show.value = false;
      });
  }
</script>
