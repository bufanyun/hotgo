<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form
        :label-width="150"
        :model="formValue"
        :rules="rules"
        ref="formRef"
        label-placement="top"
      >
        <n-divider title-placement="left">基础设置</n-divider>
        <n-form-item label="默认驱动" path="smsDrive">
          <n-select
            placeholder="默认发送驱动"
            :options="dict.getOptionUnRef('config_sms_drive')"
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

        <n-tabs type="card" size="small" v-model:value="tabName">
          <n-tab-pane name="aliyun">
            <template #tab> 阿里云 </template>
            <n-divider title-placement="left"> 阿里云</n-divider>
            <n-form-item label="AccessKeyID" path="smsAliYunAccessKeyID">
              <n-input v-model:value="formValue.smsAliYunAccessKeyID" placeholder="" />
              <template #feedback>
                应用key和密钥你可以通过 https://ram.console.aliyun.com/manage/ak 获取
              </template>
            </n-form-item>

            <n-form-item label="AccessKeySecret" path="smsAliYunAccessKeySecret">
              <n-input
                type="password"
                v-model:value="formValue.smsAliYunAccessKeySecret"
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

            <n-form-item label="签名" path="smsAliYunSign">
              <n-input v-model:value="formValue.smsAliYunSign" placeholder="" />
              <template #feedback>
                申请地址：https://dysms.console.aliyun.com/domestic/text/sign
              </template>
            </n-form-item>

            <n-form-item label="短信模板" path="smsAliYunTemplate">
              <n-dynamic-input
                v-model:value="formValue.smsAliYunTemplate"
                preset="pair"
                key-placeholder="事件KEY"
                value-placeholder="模板CODE"
              />
            </n-form-item>
          </n-tab-pane>

          <n-tab-pane name="tencent">
            <template #tab> 腾讯云 </template>
            <n-divider title-placement="left"> 腾讯云</n-divider>
            <n-form-item label="APPID" path="smsTencentSecretId">
              <n-input v-model:value="formValue.smsTencentSecretId" placeholder="" />
              <template #feedback>
                子账号密钥获取地址：https://cloud.tencent.com/document/product/598/37140
              </template>
            </n-form-item>

            <n-form-item label="密钥" path="smsTencentSecretKey">
              <n-input
                type="password"
                v-model:value="formValue.smsTencentSecretKey"
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

            <n-form-item label="接入地域域名" path="smsTencentEndpoint">
              <n-input v-model:value="formValue.smsTencentEndpoint" placeholder="" />
              <template #feedback>
                默认就近地域接入域名为 <n-text code>sms.tencentcloudapi.com</n-text>
                ，也支持指定地域域名访问，例如广州地域的域名为
                <n-text code>sms.ap-guangzhou.tencentcloudapi.com</n-text>
              </template>
            </n-form-item>

            <n-form-item label="地域信息" path="smsTencentRegion">
              <n-input v-model:value="formValue.smsTencentRegion" placeholder="" />
              <template #feedback>
                支持的地域列表参考
                https://cloud.tencent.com/document/api/382/52071#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8
              </template>
            </n-form-item>

            <n-form-item label="短信应用ID" path="smsTencentAppId">
              <n-input v-model:value="formValue.smsTencentAppId" placeholder="" />
              <template #feedback>
                查看地址：https://console.cloud.tencent.com/smsv2/app-manage
              </template>
            </n-form-item>

            <n-form-item label="签名" path="smsTencentSign">
              <n-input v-model:value="formValue.smsTencentSign" placeholder="" />
              <template #feedback>
                查看地址：https://console.cloud.tencent.com/smsv2/csms-sign
              </template>
            </n-form-item>

            <n-form-item label="短信模板" path="smsTencentTemplate">
              <n-dynamic-input
                v-model:value="formValue.smsTencentTemplate"
                preset="pair"
                key-placeholder="事件KEY"
                value-placeholder="模板ID"
              />
            </n-form-item>
          </n-tab-pane>

          <!-- tencent -->
        </n-tabs>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
            <n-button type="default" @click="sendTest">发送测试短信</n-button>
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
          <n-select
            :options="dict.getOptionUnRef('config_sms_template')"
            v-model:value="formParams.event"
          />
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
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 发送 </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, watch } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, sendTestSms, updateConfig } from '@/api/sys/config';
  import { GlassesOutline, Glasses } from '@vicons/ionicons5';
  import { useDictStore } from '@/store/modules/dict';
  import type { FormRules } from 'naive-ui/es/form/src/interface';

  const dict = useDictStore();
  const message = useMessage();
  const group = ref('sms');
  const show = ref(false);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formParams = ref({ mobile: '', event: '', code: '1234' });
  const formTestRef = ref<any>();
  const formRef: any = ref(null);
  const defaultTabName = 'aliyun';
  const tabName = ref<string>(defaultTabName);

  const rules: FormRules = {
    smsDrive: {
      required: true,
      message: '请输入默认驱动',
      trigger: 'blur',
    },
  };

  const formValue = ref({
    smsDrive: defaultTabName,
    smsMinInterval: 60,
    smsMaxIpLimit: 10,
    smsCodeExpire: 600,
    smsAliYunAccessKeyID: '',
    smsAliYunAccessKeySecret: '',
    smsAliYunSign: '',
    smsAliYunTemplate: null,
    smsTencentSecretId: '',
    smsTencentSecretKey: '',
    smsTencentEndpoint: 'sms.tencentcloudapi.com',
    smsTencentRegion: 'ap-guangzhou',
    smsTencentAppId: '',
    smsTencentSign: '',
    smsTencentTemplate: null,
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

  function load() {
    show.value = true;
    getConfig({ group: group.value })
      .then((res) => {
        res.list.smsAliYunTemplate = JSON.parse(res.list.smsAliYunTemplate);
        res.list.smsTencentTemplate = JSON.parse(res.list.smsTencentTemplate);
        formValue.value = res.list;
      })
      .finally(() => {
        show.value = false;
      });
  }

  function loadOptions() {
    dict.loadOptions(['config_sms_template', 'config_sms_drive']);
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

  watch(
    () => formValue.value.smsDrive,
    (smsDrive: string) => {
      tabName.value = smsDrive;
    }
  );

  onMounted(() => {
    loadOptions();
    load();
  });
</script>
