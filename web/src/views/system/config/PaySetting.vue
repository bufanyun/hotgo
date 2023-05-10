<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="开启debug" path="payDebug">
          <n-switch size="large" v-model:value="formValue.payDebug" />
          <template #feedback>开启后控制台会输出支付相关的日志</template>
        </n-form-item>

        <n-divider title-placement="left">支付宝</n-divider>
        <n-alert :show-icon="false" type="info">
          确保你已经申请开通过支付宝相关产品权限，建议按照以下步骤进行配置
          <br />1.
          下载支付宝平台密钥工具（下载地址：https://opendocs.alipay.com/common/02kipk），加签方式选择证书，加密算法选择RSA2
          <br />2. 生成后的私钥请在工具中转换为PKCS1格式 <br />3.
          在支付宝中配置证书，参考地址：https://opendocs.alipay.com/common/02khjo?pathHash=5403bedd
        </n-alert>
        <n-form-item label="应用ID" path="payAliPayAppId">
          <n-input v-model:value="formValue.payAliPayAppId" placeholder="" />
          <template #feedback></template>
        </n-form-item>

        <n-form-item label="应用私钥路径" path="payAliPayPrivateKey">
          <n-input v-model:value="formValue.payAliPayPrivateKey" placeholder="" clearable />
          <template #feedback
            >RSA2 加密算法默认生成格式为 PKCS8，系统默认是RSA2加密，切记转换为 PKCS1 格式</template
          >
        </n-form-item>

        <n-form-item label="应用公钥" path="payAliPayAppCertPublicKey">
          <n-input v-model:value="formValue.payAliPayAppCertPublicKey" placeholder="" clearable />
          <template #feedback>appCertPublicKey.crt证书路径</template>
        </n-form-item>

        <n-form-item label="支付宝根证书路径" path="payAliPayRootCert">
          <n-input v-model:value="formValue.payAliPayRootCert" placeholder="" clearable />
          <template #feedback>alipayRootCert.crt证书路径"</template>
        </n-form-item>

        <n-form-item label="支付宝公钥证书路径" path="payAliPayCertPublicKeyRSA2">
          <n-input v-model:value="formValue.payAliPayCertPublicKeyRSA2" placeholder="" clearable />
          <template #feedback>alipayCertPublicKey_RSA2.crt证书路径"</template>
        </n-form-item>

        <n-divider title-placement="left">微信支付</n-divider>
        <n-form-item label="应用ID" path="payWxPayAppId">
          <n-input v-model:value="formValue.payWxPayAppId" placeholder="" />
          <template #feedback>和微信配置中的微信公众号配置保持一致</template>
        </n-form-item>

        <n-form-item label="商户ID" path="payWxPayMchId">
          <n-input v-model:value="formValue.payWxPayMchId" placeholder="" />
          <template #feedback>商户ID 或者服务商模式的 sp_mchid</template>
        </n-form-item>

        <n-form-item label="证书序列号" path="payWxPaySerialNo">
          <n-input v-model:value="formValue.payWxPaySerialNo" placeholder="" />
          <template #feedback>商户证书的证书序列号</template>
        </n-form-item>
        <n-form-item label="APIv3Key" path="payWxPayAPIv3Key">
          <n-input v-model:value="formValue.payWxPayAPIv3Key" placeholder="" clearable />
          <template #feedback>商户平台获取</template>
        </n-form-item>

        <n-form-item label="私钥" path="payWxPayPrivateKey">
          <n-input
            type="textarea"
            v-model:value="formValue.payWxPayPrivateKey"
            placeholder=""
            clearable
          />
          <template #feedback>apiclient_key.pem 读取后的内容</template>
        </n-form-item>

        <n-divider title-placement="left">QQ支付</n-divider>
        <n-form-item label="应用ID" path="payQQPayAppId">
          <n-input v-model:value="formValue.payQQPayAppId" placeholder="" />
          <template #feedback></template>
        </n-form-item>

        <n-form-item label="商户ID" path="payQQPayMchId">
          <n-input v-model:value="formValue.payQQPayMchId" placeholder="" />
          <template #feedback></template>
        </n-form-item>

        <n-form-item label="ApiKey" path="payQQPayApiKey">
          <n-input
            type="textarea"
            v-model:value="formValue.payQQPayApiKey"
            placeholder=""
            clearable
          />
          <template #feedback>API秘钥值</template>
        </n-form-item>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
<!--            <n-button type="default" @click="sendTest">测试支付</n-button>-->
          </n-space>
        </div>
      </n-form>
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

  const group = ref('pay');
  const show = ref(false);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formParams = ref({ mobile: '', event: '', code: '1234' });
  const rules = {};
  const formTestRef = ref<any>();
  const formRef: any = ref(null);
  const message = useMessage();

  const options = ref<Options>({
    config_sms_template: [],
    config_sms_drive: [],
  });

  const formValue = ref({
    payDebug: true,
    payAliPayAppId: '',
    payAliPayPrivateKey: '',
    payAliPayAppCertPublicKey: '',
    payAliPayRootCert: '',
    payAliPayCertPublicKeyRSA2: '',
    payWxPayAppId: '',
    payWxPayMchId: '',
    payWxPaySerialNo: '',
    payWxPayAPIv3Key: '',
    payWxPayPrivateKey: '',
    payQQPayAppId: '',
    payQQPayMchId: '',
    payQQPayApiKey: '',
  });

  function sendTest() {
    showModal.value = true;
    formBtnLoading.value = false;
  }

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
    await loadOptions();
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
