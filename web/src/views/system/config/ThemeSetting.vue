<template>
  <div>
    <n-spin :show="show" description="正在获取配置...">
      <n-grid cols="2 s:2 m:2 l:2 xl:2 2xl:2" responsive="screen">
        <n-grid-item>
          <n-form :label-width="80" :model="formValue" :rules="rules" ref="formRef">
            <n-form-item label="默认主题" path="themeDarkTheme">
              <n-input v-model:value="formValue.themeDarkTheme" placeholder="" />
              <template #feedback> 可选：'dark' 、 'light' </template>
            </n-form-item>

            <n-form-item label="默认系统主题" path="themeAppTheme">
              <n-input v-model:value="formValue.themeAppTheme" placeholder="" />
              <template #feedback> 默认：#2d8cf0 </template>
            </n-form-item>
            <n-form-item label="默认侧边栏风格" path="themeNavTheme">
              <n-input v-model:value="formValue.themeNavTheme" placeholder="" />
              <template #feedback>可选：'light'、 'header-dark'</template>
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

  const group = ref('theme');
  const show = ref(false);

  const rules = {
    themeDarkTheme: {
      required: true,
      message: '请输入默认主题',
      trigger: 'blur',
    },
  };

  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    themeDarkTheme: 'dark',
    themeAppTheme: '#2d8cf0',
    themeNavTheme: 'dark',
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
