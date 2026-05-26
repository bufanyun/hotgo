<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        :title="data.id ? '日志详情 ID：' + data.id : '日志详情'"
      >
        <template #header-extra>
          <n-button icon-placement="right" @click="goBackOrToPage({ name: 'log' })">
            <template #icon>
              <n-icon>
                <ArrowRightOutlined />
              </n-icon>
            </template>
            返回
          </n-button>
        </template>
        <n-descriptions label-placement="left" class="py-2" :column="settingStore.isMobile ? 1 : 3">
          <n-descriptions-item label="请求方式">{{ data.method }}</n-descriptions-item>
          <n-descriptions-item>
            <template #label>请求地址</template>
            {{ data.url }}
          </n-descriptions-item>
          <n-descriptions-item label="接口名称"
            >{{ data.tags }} / {{ data.summary }}</n-descriptions-item
          >
          <n-descriptions-item label="访问IP">{{ data.ip }}</n-descriptions-item>
          <n-descriptions-item label="IP归属地">{{ data.cityLabel }}</n-descriptions-item>
          <n-descriptions-item label="链路ID">{{ data.reqId }}</n-descriptions-item>
          <n-descriptions-item label="请求耗时">{{ data.takeUpTime }} ms</n-descriptions-item>
          <n-descriptions-item label="响应时间">{{
            data.timestamp > 0 ? timestampToTime(data.timestamp) : '--'
          }}</n-descriptions-item>

          <n-descriptions-item label="访问时间">{{ data.createdAt }}</n-descriptions-item>
        </n-descriptions>
      </n-card>
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="访问代理"
      >
        {{ data.userAgent }}
      </n-card>
      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="报错信息"
      >
        <n-descriptions label-placement="left" class="py-2">
          <n-descriptions-item label="错误状态码"> {{ data.errorCode }} </n-descriptions-item>
          <n-descriptions-item label="错误提示">
            <n-tag type="error"> {{ data.errorMsg }} </n-tag>
          </n-descriptions-item>
        </n-descriptions>
      </n-card>

      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="堆栈打印"
      >
        <JsonViewer
          :value="data.errorData"
          :expand-depth="5"
          copyable
          boxed
          sort
          class="json-width"
        />
      </n-card>

      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="Header请求头"
      >
        <JsonViewer
          :value="data.headerData"
          :expand-depth="5"
          copyable
          boxed
          sort
          class="json-width"
        />
      </n-card>

      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="GET参数"
      >
        <JsonViewer
          :value="data.getData"
          :expand-depth="5"
          copyable
          boxed
          sort
          class="json-width"
        />
      </n-card>

      <n-card
        :bordered="false"
        class="proCard mt-4"
        size="small"
        :segmented="{ content: true }"
        title="POST参数"
      >
        <JsonViewer
          :value="data.postData"
          :expand-depth="5"
          copyable
          boxed
          sort
          class="json-width"
        />
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { JsonViewer } from 'vue3-json-viewer';
  import 'vue3-json-viewer/dist/vue3-json-viewer.css';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/log/log';
  import { timestampToTime } from '@/utils/dateUtil';
  import { ArrowRightOutlined } from '@vicons/antd';
  import { goBackOrToPage } from '@/utils/urlUtils';
  import { State } from '@/views/log/log/model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';

  const message = useMessage();
  const router = useRouter();
  const settingStore = useProjectSettingStore();
  const params = router.currentRoute.value.params;
  const loading = ref(false);
  const data = ref<State>(new State());

  const getInfo = () => {
    loading.value = true;
    View(params)
      .then((res) => {
        data.value = res as unknown as State;
      })
      .finally(() => {
        loading.value = false;
      });
  };

  onMounted(() => {
    if (!params.id) {
      message.error('日志ID不正确，请检查！');
      return;
    }
    getInfo();
  });
</script>

<style lang="less" scoped>
  ::v-deep(.json-width) {
    width: 100%;
    min-width: 3.125rem;
  }
</style>
