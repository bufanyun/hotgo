<template>
  <div>
    <n-card
      :bordered="false"
      class="proCard mt-4"
      size="small"
      :segmented="{ content: true }"
      :title="data.id ? '日志详情 ID：' + data.id : '日志详情'"
    >
      <n-descriptions label-placement="left" class="py-2">
        <n-descriptions-item label="请求方式">{{ data.method }}</n-descriptions-item>
        <n-descriptions-item>
          <template #label>请求地址</template>
          {{ data.url }}
        </n-descriptions-item>
        <n-descriptions-item label="请求耗时">{{ data.takeUpTime }} ms</n-descriptions-item>
        <n-descriptions-item label="访问IP">{{ data.ip }}</n-descriptions-item>
        <n-descriptions-item label="IP归属地">河南 郑州</n-descriptions-item>
        <n-descriptions-item label="链路ID">{{ data.reqId }}</n-descriptions-item>
        <n-descriptions-item label="响应时间">{{
          timestampToTime(data.timestamp)
        }}</n-descriptions-item>

        <n-descriptions-item label="创建时间">{{ data.createdAt }}</n-descriptions-item>
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
        <n-descriptions-item label="报错状态码"> {{ data.errorCode }} </n-descriptions-item>
        <n-descriptions-item label="报错消息">
          <n-tag type="success"> {{ data.errorMsg }} </n-tag>
        </n-descriptions-item>
        <n-descriptions-item label="报错日志">
          <n-tag type="success"> {{ data.errorData }} </n-tag>
        </n-descriptions-item>
      </n-descriptions>
    </n-card>

    <n-card
      :bordered="false"
      class="proCard mt-4"
      size="small"
      :segmented="{ content: true }"
      title="Header请求头"
    >
      <JsonViewer
        :value="JSON.parse(data.headerData ?? '{}')"
        :expand-depth="5"
        copyable
        boxed
        sort
        style="width: 100%; min-width: 3.125rem"
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
        :value="JSON.parse(data.getData ?? '{}')"
        :expand-depth="5"
        copyable
        boxed
        sort
        style="width: 100%; min-width: 3.125rem"
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
        :value="JSON.parse(data.postData ?? '{}')"
        :expand-depth="5"
        copyable
        boxed
        sort
        style="width: 100%; min-width: 3.125rem"
      />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { JsonViewer } from 'vue3-json-viewer';
  import 'vue3-json-viewer/dist/index.css';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/log/log';
  import { timestampToTime } from '@/utils/dateUtil';

  const message = useMessage();
  const router = useRouter();
  const logId = Number(router.currentRoute.value.params.id);

  onMounted(async () => {
    if (logId === undefined || logId < 1) {
      message.error('ID不正确，请检查！');
      return;
    }

    await getInfo();
  });

  const data = ref({});

  const getInfo = async () => {
    data.value = await View({ id: logId });
  };
</script>

<style lang="less" scoped></style>
