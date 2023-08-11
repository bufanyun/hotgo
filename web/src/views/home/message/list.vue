<template>
  <n-spin :show="loading">
    <n-empty v-show="!dataSource.list || dataSource.list.length === 0" description="无数据" />

    <n-list hoverable clickable class="list-item">
      <n-list-item v-for="item in dataSource.list" :key="item.id" @click="UnRead(item)">
        <n-thing
          content-indented
          :title="item.title"
          :description="item.createdAt"
          :content-style="{ padding: '10px' }"
        >
          <template #avatar>
            <n-badge v-bind="getBadgePops(item)">
              <n-avatar v-if="item.senderAvatar !== ''" round :size="28" :src="item.senderAvatar" />
              <n-icon-wrapper v-else :size="28" :border-radius="10">
                <n-icon :size="20" :component="getIcon(item)" />
              </n-icon-wrapper>
            </n-badge>
          </template>

          <template #header-extra>
            <n-tag
              v-if="item.tagTitle !== '' && item.tagTitle !== undefined"
              v-bind="item.tagProps"
              size="large"
              strong
            >
              {{ item.tagTitle }}
            </n-tag>
          </template>

          <template #footer>
            <span v-html="filters(item.content)"></span>
          </template>
        </n-thing>
      </n-list-item>
    </n-list>
  </n-spin>

  <n-space justify="end" style="margin-top: 30px">
    <n-pagination
      v-model:page="dataSource.page"
      :page-count="dataSource.pageCount"
      :page-slot="5"
      :page-sizes="[5, 10, 50, 100]"
      size="medium"
      show-quick-jumper
      show-size-picker
      :on-update:page="onUpdatePage"
      :on-update:page-size="onUpdatePageSize"
    >
      <template #prefix>共 {{ dataSource.totalCount }} 条</template>
    </n-pagination>
  </n-space>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { getIcon, MessageRow, parseMessage } from '@/enums/systemMessageEnum';
  import { MessageList, UpRead } from '@/api/apply/notice';
  import { debounce } from 'lodash-es';
  import { notificationStoreWidthOut } from '@/store/modules/notification';

  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '1',
  });

  interface dataList {
    page: number;
    pageSize: number;
    pageCount: number;
    totalCount: number;
    list: null | MessageRow[];
  }

  const dataSource = ref<dataList>({
    page: 1,
    pageSize: 5,
    pageCount: 1,
    list: [],
  });

  const loading = ref(false);
  const notificationStore = notificationStoreWidthOut();

  function loadDataSource() {
    loading.value = true;
    MessageList({
      type: props.type,
      page: dataSource.value.page,
      pageSize: dataSource.value.pageSize,
    })
      .then((res) => {
        if (res.list?.length > 0) {
          for (let i = 0; i < res.list.length; i++) {
            res.list[i] = parseMessage(res.list[i]);
          }
        }
        dataSource.value = res as dataList;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  function UnRead(item: MessageRow) {
    UpRead({ id: item.id })
      .then(() => {
        item.isRead = true;
        debounceCallback();
      })
      .finally(() => {
        loading.value = false;
      });
  }

  const debounceCallback = debounce(function () {
    notificationStore.pullMessages();
  }, 1000);

  function getBadgePops(item: MessageRow) {
    if (item.isRead) {
      return {};
    }
    return { dot: true, processing: true, offset: [-2, 2] };
  }

  function onUpdatePage(page: number) {
    dataSource.value.page = page;
    loadDataSource();
  }

  function onUpdatePageSize(pageSize: number) {
    dataSource.value.pageSize = pageSize;
    loadDataSource();
  }

  onMounted(() => {
    loadDataSource();
  });

  function filters(data) {
    return data.replace(/\n/g, '<br>');
  }
</script>

<style lang="less" scoped>
  ::v-deep(.list-item) {
    margin-left: calc(1vw);
    margin-right: calc(1vw);
  }

  :deep(img, video, audio) {
    width: 100%;
  }
</style>
