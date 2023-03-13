<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="网站名称" path="basicName">
          <n-input v-model:value="formValue.basicName" placeholder="请输入网站名称" />
        </n-form-item>

        <n-form-item label="网站logo" path="basicLogo">
          <UploadImage
            :maxNumber="1"
            v-model:value="formValue.basicLogo"
            :helpText="
              '网站logo适用于客户端使用，图片大小不超过' + componentSetting.upload.maxSize + 'MB'
            "
          />
        </n-form-item>

        <n-form-item label="网站域名" path="basicDomain">
          <n-input v-model:value="formValue.basicDomain" placeholder="请输入网站域名" />
          <template #feedback>
            如果客户端通过本地或内网访问，则认为是调试模式，走实际请求地址，否则走该配置
          </template>
        </n-form-item>

        <n-form-item label="websocket地址" path="basicWsAddr">
          <n-input v-model:value="formValue.basicWsAddr" placeholder="请输入websocket地址" />
          <template #feedback>
            如果客户端通过本地或内网访问，则认为是调试模式，走实际请求地址，否则走该配置
          </template>
        </n-form-item>

        <n-form-item label="用户是否可注册开关" path="basicRegisterSwitch">
          <n-radio-group v-model:value="formValue.basicRegisterSwitch" name="basicRegisterSwitch">
            <n-space>
              <n-radio :value="1">开启</n-radio>
              <n-radio :value="0">关闭</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="验证码开关" path="basicCaptchaSwitch">
          <n-radio-group v-model:value="formValue.basicCaptchaSwitch" name="basicCaptchaSwitch">
            <n-space>
              <n-radio :value="1">开启</n-radio>
              <n-radio :value="0">关闭</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="网站开启访问" path="basicSystemOpen">
          <n-switch
            size="large"
            v-model:value="formValue.basicSystemOpen"
            @update:value="systemOpenChange"
          />
        </n-form-item>

        <n-form-item label="网站关闭提示" path="basicCloseText">
          <n-input
            v-model:value="formValue.basicCloseText"
            type="textarea"
            placeholder="请输入网站关闭提示"
          />
        </n-form-item>

        <n-form-item label="备案编号" path="basicIcpCode">
          <n-input placeholder="请输入备案编号" v-model:value="formValue.basicIcpCode" />
        </n-form-item>

        <n-form-item label="版权所有" path="basicCopyright">
          <n-input placeholder="版权所有" v-model:value="formValue.basicCopyright" />
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
  import { ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import componentSetting from '@/settings/componentSetting';
  import { getConfig, updateConfig } from '@/api/sys/config';

  const group = ref('basic');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();

  const formValue = ref({
    basicName: 'HotGo',
    basicLogo: '',
    basicDomain: 'https://hotgo.facms.cn',
    basicWsAddr: 'wss://hotgo.facms.cn/socket',
    basicIcpCode: '',
    basicLoginCode: 0,
    basicRegisterSwitch: 1,
    basicCaptchaSwitch: 1,
    basicCopyright: '© 2021 - 2023 HotGo All Rights Reserved.',
    basicCloseText:
      '网站维护中，暂时无法访问！本网站正在进行系统维护和技术升级，网站暂时无法访问，敬请谅解！',
    basicSystemOpen: true,
  });

  const rules = {
    basicName: {
      required: true,
      message: '请输入网站名称',
      trigger: 'blur',
    },
  };

  function systemOpenChange(value) {
    if (!value) {
      dialog.warning({
        title: '提示',
        content: '您确定要关闭系统访问吗？该操作保存后立马生效，请慎重操作！',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          // message.success('操作成功');
        },
        onNegativeClick: () => {
          formValue.value.basicSystemOpen = true;
        },
      });
    }
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

  function load() {
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
