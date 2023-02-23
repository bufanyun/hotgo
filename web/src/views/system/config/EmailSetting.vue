<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="SMTP服务器" path="smtpHost">
          <n-input v-model:value="formValue.smtpHost" placeholder="" />
          <template #feedback> 错误的配置发送邮件会导致服务器超时</template>
        </n-form-item>

        <n-form-item label="SMTP端口" path="smtpPort">
          <n-input-number v-model:value="formValue.smtpPort" placeholder="" :show-button="false" />
          <template #feedback> (不加密默认25,SSL默认465,TLS默认587)</template>
        </n-form-item>
        <n-form-item label="SMTP用户名" path="smtpUser">
          <n-input v-model:value="formValue.smtpUser" placeholder="" />
          <template #feedback>填写完整用户名</template>
        </n-form-item>

        <n-form-item label="SMTP密码" path="smtpPass">
          <n-input
            v-model:value="formValue.smtpPass"
            placeholder=""
            type="password"
            show-password-on="click"
          >
            <template #password-visible-icon>
              <n-icon :size="16" :component="GlassesOutline" />
            </template>
            <template #password-invisible-icon>
              <n-icon :size="16" :component="Glasses" />
            </template>
          </n-input>
          <template #feedback>填写您的密码</template>
        </n-form-item>

        <n-form-item label="发件人名称" path="smtpSendName">
          <n-input v-model:value="formValue.smtpSendName" placeholder="" />
        </n-form-item>

        <n-form-item label="管理员邮箱" path="smtpAdminMailbox">
          <n-input v-model:value="formValue.smtpAdminMailbox" placeholder="" />
        </n-form-item>

        <n-divider title-placement="left">发信限制</n-divider>
        <n-form-item label="最小发送间隔" path="smtpMinInterval">
          <n-input-number
            :show-button="false"
            placeholder="请输入"
            v-model:value="formValue.smtpMinInterval"
          >
            <template #suffix> 秒 </template>
          </n-input-number>
          <template #feedback> 同地址</template>
        </n-form-item>
        <n-form-item label="IP最大发送次数" path="smtpMaxIpLimit">
          <n-input-number v-model:value="formValue.smtpMaxIpLimit" placeholder="" />
          <template #feedback> 同IP每天最大允许发送次数 </template>
        </n-form-item>
        <n-form-item label="验证码有效期" path="smtpCodeExpire">
          <n-input-number
            :show-button="false"
            placeholder="请输入"
            v-model:value="formValue.smtpCodeExpire"
          >
            <template #suffix> 秒 </template>
          </n-input-number>
        </n-form-item>

        <n-form-item label="邮件模板" path="smtpTemplate">
          <n-dynamic-input
            v-model:value="formValue.smtpTemplate"
            preset="pair"
            key-placeholder="事件KEY"
            value-placeholder="模板路径"
          />
        </n-form-item>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
            <n-button type="default" @click="sendTest">发送测试邮件</n-button>
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
          <n-input
            type="textarea"
            placeholder="多个用;隔开"
            v-model:value="formParams.to"
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
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, sendTestEmail, updateConfig } from '@/api/sys/config';
  import { GlassesOutline, Glasses } from '@vicons/ionicons5';

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
    smtpMinInterval: 60,
    smtpMaxIpLimit: 10,
    smtpCodeExpire: 600,
    smtpTemplate: null,
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
        updateConfig({ group: group.value, list: formValue.value }).then((res) => {
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

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          res.list.smtpTemplate = JSON.parse(res.list.smtpTemplate);
          formValue.value = res.list;
        })
        .finally(() => {
          show.value = false;
        });
    });
  }
</script>
