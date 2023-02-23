<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="默认驱动" path="uploadDrive">
          <n-select
            placeholder="默认驱动"
            :options="uploadDriveList"
            v-model:value="formValue.uploadDrive"
          />
        </n-form-item>

        <n-divider title-placement="left">上传限制</n-divider>
        <n-form-item label="图片大小限制" path="uploadImageSize">
          <n-input-number
            :show-button="false"
            placeholder="请输入"
            v-model:value="formValue.uploadImageSize"
          >
            <template #suffix> MB</template>
          </n-input-number>
        </n-form-item>
        <n-form-item label="图片类型限制" path="uploadImageType">
          <n-input v-model:value="formValue.uploadImageType" placeholder="" />
        </n-form-item>

        <n-form-item label="文件大小限制" path="uploadFileSize">
          <n-input-number
            :show-button="false"
            placeholder="请输入"
            v-model:value="formValue.uploadFileSize"
          >
            <template #suffix> MB</template>
          </n-input-number>
        </n-form-item>
        <n-form-item label="文件类型限制" path="uploadFileType">
          <n-input v-model:value="formValue.uploadFileType" placeholder="" />
        </n-form-item>

        <n-divider title-placement="left">本地存储</n-divider>
        <n-form-item label="本地存储路径" path="uploadLocalPath">
          <n-input v-model:value="formValue.uploadLocalPath" placeholder="" />
          <template #feedback>填对外访问的相对路径</template>
        </n-form-item>

        <n-divider title-placement="left">UCloud存储</n-divider>
        <n-form-item label="公钥" path="uploadUCloudPublicKey">
          <n-input
            type="password"
            v-model:value="formValue.uploadUCloudPublicKey"
            show-password-on="click"
          >
            <template #password-visible-icon>
              <n-icon :size="16" :component="GlassesOutline" />
            </template>
            <template #password-invisible-icon>
              <n-icon :size="16" :component="Glasses" />
            </template>
          </n-input>
          <template #feedback>获取地址：https://console.ucloud.cn/ufile/token</template>
        </n-form-item>

        <n-form-item label="私钥" path="uploadUCloudPrivateKey">
          <n-input
            type="password"
            v-model:value="formValue.uploadUCloudPrivateKey"
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
        <n-form-item label="存储路径" path="uploadUCloudPath">
          <n-input v-model:value="formValue.uploadUCloudPath" placeholder="" />
          <template #feedback>填对对象存储中的相对路径</template>
        </n-form-item>
        <n-form-item label="地域API" path="uploadUCloudBucketHost">
          <n-input v-model:value="formValue.uploadUCloudBucketHost" placeholder="" />
        </n-form-item>
        <n-form-item label="存储桶名称" path="uploadUCloudBucketName">
          <n-input v-model:value="formValue.uploadUCloudBucketName" placeholder="" />
          <template #feedback>存储空间名称</template>
        </n-form-item>
        <n-form-item label="存储桶地域host" path="uploadUCloudFileHost">
          <n-input v-model:value="formValue.uploadUCloudFileHost" placeholder="" />
          <template #feedback>不需要包含桶名称</template>
        </n-form-item>
        <n-form-item label="访问域名" path="uploadUCloudEndpoint">
          <n-input v-model:value="formValue.uploadUCloudEndpoint" placeholder="" />
          <template #feedback>格式，http://abc.com 或 https://abc.com，不可为空</template>
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
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import { Glasses, GlassesOutline } from '@vicons/ionicons5';

  const group = ref('upload');
  const show = ref(false);

  const rules = {
    uploadDrive: {
      required: true,
      message: '请输入默认驱动',
      trigger: 'blur',
    },
  };

  const uploadDriveList = [
    {
      label: '本地存储',
      value: 'local',
    },
    {
      label: 'UC云存储',
      value: 'ucloud',
    },
  ];

  const formRef: any = ref(null);
  const message = useMessage();

  const formValue = ref({
    uploadDrive: 'local',
    uploadImageSize: 2,
    uploadImageType: '',
    uploadFileSize: 10,
    uploadFileType: '',
    uploadLocalPath: '',
    uploadUCloudPath: '',
    uploadUCloudPublicKey: '',
    uploadUCloudPrivateKey: '',
    uploadUCloudBucketHost: 'api.ucloud.cn',
    uploadUCloudBucketName: '',
    uploadUCloudFileHost: 'cn-bj.ufileos.com',
    uploadUCloudEndpoint: '',
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
