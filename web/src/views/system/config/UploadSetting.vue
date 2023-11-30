<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="默认驱动" path="uploadDrive">
          <n-select
            placeholder="默认驱动"
            :options="options.config_upload_drive"
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

        <template v-if="formValue.uploadDrive == 'local'">
          <n-divider title-placement="left">本地存储</n-divider>
          <n-form-item label="本地存储路径" path="uploadLocalPath">
            <n-input v-model:value="formValue.uploadLocalPath" placeholder="" />
            <template #feedback>填对外访问的相对路径</template>
          </n-form-item>
        </template>

        <template v-if="formValue.uploadDrive == 'ucloud'">
          <n-divider title-placement="left">ucloud对象存储</n-divider>
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
        </template>

        <template v-if="formValue.uploadDrive == 'cos'">
          <n-divider title-placement="left">腾讯云COS存储</n-divider>
          <n-form-item label="secretId" path="uploadCosSecretId">
            <n-input
              type="password"
              v-model:value="formValue.uploadCosSecretId"
              show-password-on="click"
            >
              <template #password-visible-icon>
                <n-icon :size="16" :component="GlassesOutline" />
              </template>
              <template #password-invisible-icon>
                <n-icon :size="16" :component="Glasses" />
              </template>
            </n-input>
            <template #feedback
              >子账号密钥获取地址：https://cloud.tencent.com/document/product/598/37140
            </template>
          </n-form-item>

          <n-form-item label="secretKey" path="uploadCosSecretKey">
            <n-input
              type="password"
              v-model:value="formValue.uploadCosSecretKey"
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
          <n-form-item label="存储路径" path="uploadCosBucketURL">
            <n-input v-model:value="formValue.uploadCosBucketURL" placeholder="" />
            <template #feedback>填对对象存储中的相对路径</template>
          </n-form-item>
          <n-form-item label="地域API" path="uploadCosPath">
            <n-input v-model:value="formValue.uploadCosPath" placeholder="" />
            <template #feedback
              >控制台查看地址：https://console.cloud.tencent.com/cos5/bucket
            </template>
          </n-form-item>
        </template>

        <template v-if="formValue.uploadDrive == 'oss'">
          <n-divider title-placement="left">阿里云OSS存储</n-divider>
          <n-form-item label="AccessKey ID" path="uploadOssSecretId">
            <n-input
              type="password"
              v-model:value="formValue.uploadOssSecretId"
              show-password-on="click"
            >
              <template #password-visible-icon>
                <n-icon :size="16" :component="GlassesOutline" />
              </template>
              <template #password-invisible-icon>
                <n-icon :size="16" :component="Glasses" />
              </template>
            </n-input>
            <template #feedback>创建地址：https://ram.console.aliyun.com/manage/ak</template>
          </n-form-item>

          <n-form-item label="AccessKey Secret" path="uploadOssSecretKey">
            <n-input
              type="password"
              v-model:value="formValue.uploadOssSecretKey"
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
          <n-form-item label="Endpoint" path="uploadOssEndpoint">
            <n-input v-model:value="formValue.uploadOssEndpoint" placeholder="" />
            <template #feedback> Endpoint（地域节点）</template>
          </n-form-item>
          <n-form-item label="存储路径" path="uploadOssPath">
            <n-input v-model:value="formValue.uploadOssPath" placeholder="" />
            <template #feedback>填对对象存储中的相对路径</template>
          </n-form-item>
          <n-form-item label="存储空间名称" path="uploadOssBucket">
            <n-input v-model:value="formValue.uploadOssBucket" placeholder="" />
          </n-form-item>
          <n-form-item label="Bucket 域名" path="uploadOssBucketURL">
            <n-input v-model:value="formValue.uploadOssBucketURL" placeholder="" />
          </n-form-item>
        </template>

        <template v-if="formValue.uploadDrive == 'qiniu'">
          <n-divider title-placement="left">七牛云对象存储</n-divider>
          <n-form-item label="AccessKey" path="uploadQiNiuAccessKey">
            <n-input
              type="password"
              v-model:value="formValue.uploadQiNiuAccessKey"
              show-password-on="click"
            >
              <template #password-visible-icon>
                <n-icon :size="16" :component="GlassesOutline" />
              </template>
              <template #password-invisible-icon>
                <n-icon :size="16" :component="Glasses" />
              </template>
            </n-input>
            <template #feedback>创建地址：https://portal.qiniu.com/user/key</template>
          </n-form-item>

          <n-form-item label="SecretKey" path="uploadQiNiuSecretKey">
            <n-input
              type="password"
              v-model:value="formValue.uploadQiNiuSecretKey"
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
          <n-form-item label="储存路径" path="uploadQiNiuPath">
            <n-input v-model:value="formValue.uploadQiNiuPath" placeholder="" />
            <template #feedback>填对对象存储中的相对路径</template>
          </n-form-item>
          <n-form-item label="存储空间名称" path="uploadQiNiuBucket">
            <n-input v-model:value="formValue.uploadQiNiuBucket" placeholder="" />
          </n-form-item>
          <n-form-item label="访问域名" path="uploadQiNiuDomain">
            <n-input v-model:value="formValue.uploadQiNiuDomain" placeholder="" />
          </n-form-item>
        </template>

        <template v-if="formValue.uploadDrive == 'minio'">
          <n-divider title-placement="left">minio配置</n-divider>
          <n-form-item label="AccessKey ID" path="uploadMinioAccessKey">
            <n-input
              type="password"
              v-model:value="formValue.uploadMinioAccessKey"
              show-password-on="click"
            >
              <template #password-visible-icon>
                <n-icon :size="16" :component="GlassesOutline" />
              </template>
              <template #password-invisible-icon>
                <n-icon :size="16" :component="Glasses" />
              </template>
            </n-input>
            <template #feedback>相关文档：https://min.io/</template>
          </n-form-item>

          <n-form-item label="AccessKey Secret" path="uploadMinioSecretKey">
            <n-input
              type="password"
              v-model:value="formValue.uploadMinioSecretKey"
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
          <n-form-item label="Endpoint" path="uploadMinioEndpoint">
            <n-input v-model:value="formValue.uploadMinioEndpoint" placeholder="" />
            <template #feedback> Endpoint（不带http://和路径）</template>
          </n-form-item>
          <n-form-item path="uploadMinioUseSSL">
            <n-switch
              v-model:value="formValue.uploadMinioUseSSL"
              :checked-value="1"
              :unchecked-value="2"
            >
              <template #checked> 启用SSL</template>
              <template #unchecked> 禁用SSL</template>
            </n-switch>
          </n-form-item>
          <n-form-item label="储存路径" path="uploadMinioPath">
            <n-input v-model:value="formValue.uploadMinioPath" placeholder="" />
          </n-form-item>
          <n-form-item label="存储桶名称" path="uploadMinioBucket">
            <n-input v-model:value="formValue.uploadMinioBucket" placeholder="" />
          </n-form-item>
          <n-form-item label="对外访问域名" path="uploadMinioDomain">
            <n-input v-model:value="formValue.uploadMinioDomain" placeholder="" />
          </n-form-item>
        </template>
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
  import { Dicts } from '@/api/dict/dict';
  import { Options } from '@/utils/hotgo';

  const group = ref('upload');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();

  const rules = {
    uploadDrive: {
      required: true,
      message: '请输入默认驱动',
      trigger: 'blur',
    },
  };

  const options = ref<Options>({
    config_upload_drive: [],
  });

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
    uploadCosSecretId: '',
    uploadCosSecretKey: '',
    uploadCosBucketURL: '',
    uploadCosPath: '',
    uploadOssSecretId: '',
    uploadOssSecretKey: '',
    uploadOssEndpoint: '',
    uploadOssBucketURL: '',
    uploadOssPath: '',
    uploadOssBucket: '',
    uploadQiNiuAccessKey: '',
    uploadQiNiuSecretKey: '',
    uploadQiNiuDomain: '',
    uploadQiNiuPath: '',
    uploadQiNiuBucket: '',
    uploadMinioAccessKey: '',
    uploadMinioSecretKey: '',
    uploadMinioEndpoint: '',
    uploadMinioUseSSL: 2,
    uploadMinioPath: '',
    uploadMinioBucket: '',
    uploadMinioDomain: '',
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

  onMounted(async () => {
    load();
    await loadOptions();
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

  async function loadOptions() {
    options.value = await Dicts({
      types: ['config_upload_drive'],
    });
  }
</script>
