<template>
  <div>
    <n-spin :show="show" description="正在获取配置...">
      <n-grid cols="2 s:2 m:2 l:2 xl:2 2xl:2" responsive="screen">
        <n-grid-item>
          <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
            <n-form-item label="SMTP服务器" path="smtpHost">
              <n-input v-model:value="formValue.smtpHost" placeholder="" />
              <template #feedback> 错误的配置发送邮件会导致服务器超时</template>
            </n-form-item>

            <n-form-item label="SMTP端口" path="smtpPort">
              <n-input-number
                v-model:value="formValue.smtpPort"
                placeholder=""
                :show-button="false"
              />
              <template #feedback> (不加密默认25,SSL默认465,TLS默认587)</template>
            </n-form-item>
            <n-form-item label="SMTP用户名" path="smtpUser">
              <n-input v-model:value="formValue.smtpUser" placeholder="" />
              <template #feedback>填写完整用户名</template>
            </n-form-item>

            <n-form-item label="SMTP密码" path="smtpPass">
              <n-input v-model:value="formValue.smtpPass" placeholder="" type="password" />
              <template #feedback>填写您的密码</template>
            </n-form-item>

            <n-form-item label="发件人名称" path="smtpSendName">
              <n-input v-model:value="formValue.smtpSendName" placeholder="" />
            </n-form-item>

            <n-form-item label="管理员邮箱" path="smtpAdminMailbox">
              <n-input v-model:value="formValue.smtpAdminMailbox" placeholder="" />
            </n-form-item>

            <div>
              <n-space>
                <n-button type="primary" @click="formSubmit">保存更新</n-button>
                <n-button type="default" @click="sendTest">发送测试邮件</n-button>
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
      title="发送测试邮件"
    >
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-form-item label="接收邮箱" path="to">
          <n-input placeholder="多个用;隔开" v-model:value="formParams.to" :required="true" />
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
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, sendTestEmail, updateConfig } from '@/api/sys/config';

  const group = ref('smtp');
  const show = ref(false);

  const showModal = ref(false);
  const formBtnLoading = ref(false);

  const formParams = ref({ to: '' });

  const rules = {
    smtpHost: {
      required: true,
      message: '请输入SMTP服务器',
      trigger: 'blur',
    },
  };

  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    smtpHost: 'smtpdm.aliyun.com',
    smtpPort: 25,
    smtpUser: '',
    smtpPass: '',
    smtpSendName: 'HotGo',
    smtpAdminMailbox: '',
  });

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        sendTestEmail(formParams.value).then((_res) => {
          message.success('发送成功');
          showModal.value = false;
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function sendTest() {
    showModal.value = true;
    formBtnLoading.value = false;
  }

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
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
          // state.formValue.watermarkClarity = res;
          formValue.value = res.list;
          console.log('res:' + JSON.stringify(res));
        })
        .catch((error) => {
          show.value = false;
          message.error(error.toString());
        });
    });
  }
</script>
