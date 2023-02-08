<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-grid cols="2 s:2 m:2 l:2 xl:2 2xl:2" responsive="screen">
        <n-grid-item>
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
        </n-grid-item>
      </n-grid>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, reactive, ref, unref } from 'vue';
  import { useMessage } from 'naive-ui';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import { BasicUpload } from '@/components/Upload';
  import { useGlobSetting } from '@/hooks/setting';
  import { useUserStoreWidthOut } from '@/store/modules/user';
  import { getUserInfo, updateMemberCash } from '@/api/system/user';

  const show = ref(false);
  const useUserStore = useUserStoreWidthOut();
  const globSetting = useGlobSetting();
  const { uploadUrl } = globSetting;
  const uploadHeaders = reactive({
    Authorization: useUserStore.token,
  });

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

  function uploadChange(list: string[]) {
    // 单图模式，只需要第一个索引
    if (list.length > 0) {
      formValue.value.payeeCode = unref(list[0]);
    } else {
      formValue.value.payeeCode = unref('');
    }
  }

  onMounted(() => {
    load();
  });

  function load() {
    show.value = true;
    getUserInfo()
      .then((res) => {
        formValue.value = res.cash;
        formValue.value.password = '';
      })
      .finally(() => {
        show.value = false;
      });
  }
</script>
