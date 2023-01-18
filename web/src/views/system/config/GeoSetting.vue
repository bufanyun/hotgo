<template>
  <div>
    <n-spin :show="show" description="正在获取配置...">
      <n-grid cols="2 s:2 m:2 l:2 xl:2 2xl:2" responsive="screen">
        <n-grid-item>
          <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
            <n-form-item label="高德Web服务key" path="geoAmapWebKey">
              <n-input v-model:value="formValue.geoAmapWebKey" placeholder="" type="password" />
              <template #feedback> 申请地址：https://console.amap.com/dev/key/app</template>
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

  const group = ref('geo');
  const show = ref(false);

  const rules = {
    geoAmapWebKey: {
      required: true,
      message: '请输入高德Web服务key',
      trigger: 'blur',
    },
  };

  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    geoAmapWebKey: '',
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
