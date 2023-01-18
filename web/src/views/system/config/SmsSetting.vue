<template>
  <div>
    <n-spin :show="show" description="正在获取配置...">
      <n-grid cols="2 s:2 m:2 l:2 xl:2 2xl:2" responsive="screen">
        <n-grid-item>
          <n-divider title-placement="left">通用配置</n-divider>
          <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
            <n-form-item label="默认驱动" path="smsDrive">
              <n-select
                placeholder="默认发送驱动"
                :options="driveList"
                v-model:value="formValue.smsDrive"
              />
            </n-form-item>

            <n-form-item label="最小发送间隔" path="smsMinInterval">
              <n-input-number
                :show-button="false"
                placeholder="请输入"
                v-model:value="formValue.smsMinInterval"
              >
                <template #suffix> 秒 </template>
              </n-input-number>
              <template #feedback> 同号码</template>
            </n-form-item>
            <n-form-item label="IP最大发送次数" path="smsMaxIpLimit">
              <n-input-number v-model:value="formValue.smsMaxIpLimit" placeholder="" />
              <template #feedback> 同IP每天最大允许发送次数 </template>
            </n-form-item>
            <n-form-item label="验证码有效期" path="smsCodeExpire">
              <n-input-number
                :show-button="false"
                placeholder="请输入"
                v-model:value="formValue.smsCodeExpire"
              >
                <template #suffix> 秒 </template>
              </n-input-number>
            </n-form-item>

            <n-divider title-placement="left">阿里云</n-divider>
            <n-form-item label="AccessKeyID" path="smsAliyunAccessKeyID">
              <n-input
                v-model:value="formValue.smsAliyunAccessKeyID"
                placeholder=""
                type="password"
              />
              <template #feedback
                >应用key和密钥你可以通过 https://ram.console.aliyun.com/manage/ak 获取</template
              >
            </n-form-item>

            <n-form-item label="AccessKeySecret" path="smsAliyunAccessKeySecret">
              <n-input
                type="password"
                v-model:value="formValue.smsAliyunAccessKeySecret"
                placeholder=""
              />
            </n-form-item>

            <n-form-item label="签名" path="smsAliyunSign">
              <n-input v-model:value="formValue.smsAliyunSign" placeholder="" />
              <template #feedback
                >申请地址：https://dysms.console.aliyun.com/domestic/text/sign</template
              >
            </n-form-item>

            <n-form-item label="短信模板" path="smsAliyunTemplate">
              <n-dynamic-input
                v-model:value="formValue.smsAliyunTemplate"
                preset="pair"
                key-placeholder="key"
                value-placeholder="模板CODE"
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
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';

  const group = ref('sms');
  const show = ref(false);

  const rules = {
    smsDrive: {
      required: true,
      message: '请输入默认驱动',
      trigger: 'blur',
    },
  };

  const driveList = [
    {
      label: '阿里云',
      value: 'aliyun',
    },
    {
      label: '腾讯云',
      value: 'tencent',
    },
  ];
  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    smsDrive: 'aliyun',
    smsAliyunAccessKeyID: '',
    smsAliyunAccessKeySecret: '',
    smsAliyunSign: '',
    smsAliyunTemplate: null,
    smsMinInterval: 60,
    smsMaxIpLimit: 10,
    smsCodeExpire: 600,
  });

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        console.log('formValue.value:' + JSON.stringify(formValue.value));

        updateConfig({ group: group.value, list: formValue.value })
          .then((res) => {
            console.log('res:' + JSON.stringify(res));
            message.success('更新成功');
            load();
          })
          .catch((error) => {
            message.error(error.toString());
          });
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
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          show.value = false;
          res.list.smsAliyunTemplate = JSON.parse(res.list.smsAliyunTemplate);
          formValue.value = res.list;
        })
        .catch((error) => {
          show.value = false;
          message.error(error.toString());
        });
    });
  }
</script>
