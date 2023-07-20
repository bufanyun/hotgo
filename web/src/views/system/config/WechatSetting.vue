<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-divider title-placement="left">公众号</n-divider>
        <n-form-item label="AppID" path="officialAccountAppId">
          <n-input v-model:value="formValue.officialAccountAppId" placeholder="" />
          <template #feedback>请填写微信公众平台后台的AppId</template>
        </n-form-item>

        <n-form-item label="AppSecret" path="officialAccountAppSecret">
          <n-input v-model:value="formValue.officialAccountAppSecret" placeholder="" clearable />
          <template #feedback>请填写微信公众平台后台的AppSecret</template>
        </n-form-item>

        <n-form-item label="Token" path="officialAccountToken">
          <n-input v-model:value="formValue.officialAccountToken" placeholder="" clearable />
          <template #feedback
            >与公众平台接入设置值一致，必须为英文或者数字，长度为3到32个字符</template
          >
        </n-form-item>

        <n-form-item label="EncodingAESKey" path="officialAccountEncodingAESKey">
          <n-input
            v-model:value="formValue.officialAccountEncodingAESKey"
            placeholder=""
            clearable
          />
          <template #feedback
            >与公众平台接入设置值一致，必须为英文或者数字，长度为43个字符
          </template>
        </n-form-item>

        <n-divider title-placement="left">开放平台</n-divider>
        <n-form-item label="AppID" path="openPlatformAppId">
          <n-input v-model:value="formValue.openPlatformAppId" placeholder="" />
          <template #feedback>请填写微信开放平台后台的AppId</template>
        </n-form-item>

        <n-form-item label="AppSecret" path="openPlatformAppSecret">
          <n-input v-model:value="formValue.openPlatformAppSecret" placeholder="" clearable />
          <template #feedback>请填写微信开放平台后台的AppSecret</template>
        </n-form-item>

        <n-form-item label="Token" path="openPlatformToken">
          <n-input v-model:value="formValue.openPlatformToken" placeholder="" clearable />
          <template #feedback
            >与开放平台接入设置值一致，必须为英文或者数字，长度为3到32个字符</template
          >
        </n-form-item>

        <n-form-item label="EncodingAESKey" path="openPlatformEncodingAESKey">
          <n-input v-model:value="formValue.openPlatformEncodingAESKey" placeholder="" clearable />
          <template #feedback
            >与开放平台接入设置值一致，必须为英文或者数字，长度为43个字符</template
          >
        </n-form-item>

        <!--        <n-divider title-placement="left">小程序</n-divider>-->

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
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';

  const group = ref('wechat');
  const show = ref(false);
  const rules = {};
  const formRef: any = ref(null);
  const message = useMessage();
  const formValue = ref({
    officialAccountAppId: '',
    officialAccountAppSecret: '',
    officialAccountToken: '',
    officialAccountEncodingAESKey: '',
    openPlatformAppId: '',
    openPlatformAppSecret: '',
    openPlatformToken: '',
    openPlatformEncodingAESKey: '',
  });

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value }).then((_res) => {
          message.success('更新成功');
          load();
        });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  onMounted(() => {
    load();
  });

  async function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
