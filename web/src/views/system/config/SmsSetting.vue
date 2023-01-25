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
                :options="options.config_sms_drive"
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
              <n-input v-model:value="formValue.smsAliyunAccessKeyID" placeholder="" />
              <template #feedback
                >应用key和密钥你可以通过 https://ram.console.aliyun.com/manage/ak 获取</template
              >
            </n-form-item>

            <n-form-item label="AccessKeySecret" path="smsAliyunAccessKeySecret">
              <n-input
                type="password"
                v-model:value="formValue.smsAliyunAccessKeySecret"
                show-password-on="click"
              >
                <template #password-visible-icon>
                  <n-icon :size="16" :component="GlassesOutline" />
                </template>
                <template #password-invisible-icon>
                  <n-icon :size="16" :component="Glasses" />
                </template>
              </n-input>
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
                key-placeholder="事件KEY"
                value-placeholder="模板CODE"
              />
            </n-form-item>

            <div>
              <n-space>
                <n-button type="primary" @click="formSubmit">保存更新</n-button>
                <n-button type="default" @click="sendTest">发送测试短信</n-button>
              </n-space>
            </div>
          </n-form>
        </n-grid-item>
      </n-grid>
    </n-spin>

    <n-modal
      :block-scroll="false"
      :mask-closable="false"
      v-model:show="showModal"
      :show-icon="false"
      preset="dialog"
      title="发送测试短信"
    >
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formTestRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-form-item label="事件模板" path="event">
          <n-select :options="options.config_sms_template" v-model:value="formParams.event" />
        </n-form-item>

        <n-form-item label="手机号" path="mobile">
          <n-input
            placeholder="请输入接收手机号"
            v-model:value="formParams.mobile"
            :required="true"
          />
        </n-form-item>

        <n-form-item label="验证码" path="code">
          <n-input
            placeholder="请输入要接收的验证码"
            v-model:value="formParams.code"
            :required="true"
          />
        </n-form-item>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showModal = false)">关闭</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">发送</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, sendTestSms, updateConfig } from '@/api/sys/config';
  import { Dicts } from '@/api/dict/dict';
  import { Options } from '@/utils/hotgo';
  import { GlassesOutline, Glasses } from '@vicons/ionicons5';

  const group = ref('sms');
  const show = ref(false);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formParams = ref({ mobile: '', event: '', code: '1234' });

  const rules = {
    smsDrive: {
      required: true,
      message: '请输入默认驱动',
      trigger: 'blur',
    },
  };

  const formTestRef = ref<any>();
  const formRef: any = ref(null);
  const message = useMessage();

  const options = ref<Options>({
    config_sms_template: [],
    config_sms_drive: [],
  });

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

  function sendTest() {
    showModal.value = true;
    formBtnLoading.value = false;
  }

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

  async function load() {
    show.value = true;
    await loadOptions();
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

  async function loadOptions() {
    options.value = await Dicts({
      types: ['config_sms_template', 'config_sms_drive'],
    });
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formTestRef.value.validate((errors) => {
      if (!errors) {
        sendTestSms(formParams.value).then((_res) => {
          message.success('发送成功');
          showModal.value = false;
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }
</script>
