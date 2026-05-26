<template>
  <n-grid :x-gap="5" :y-gap="5" :cols="pageGridCols" responsive="screen">
    <n-gi :span="2">
      <n-menu
        :options="generateKindOptions()"
        style="width: 100%"
        :default-value="defaultKindValue"
        :on-update:value="handleUpdateKind"
      />
    </n-gi>
    <n-gi :span="8">
      <n-spin style="height: 100%" :show="loading">
        <n-layout style="height: 100%" content-style="display:flex;flex-direction: column;">
          <n-layout-header>
            <BasicForm
              style="padding-top: 10px; box-sizing: border-box"
              @register="register"
              @submit="reloadTable"
              @reset="reloadTable"
              @keyup.enter="reloadTable"
              ref="searchFormRef"
            >
              <template #statusSlot="{ model, field }">
                <n-input v-model:value="model[field]" />
              </template>
            </BasicForm>
          </n-layout-header>

          <n-empty v-if="isEmptyDataSource()" description="无数据" />
          <n-layout-content style="padding: 5px; box-sizing: border-box">
            <n-grid :cols="imageGridCols" x-gap="15" y-gap="15" responsive="screen">
              <n-grid-item
                v-for="item in dataSource"
                :key="item.id"
                :class="{ imageActive: isSelected(item) }"
              >
                <n-card
                  size="small"
                  hoverable
                  content-style="padding: 3px;"
                  footer-style="padding: 0"
                  style="overflow: hidden"
                  bordered
                >
                  <div @click="handleSelect(item)">
                    <n-image
                      v-if="item.kind === 'image'"
                      preview-disabled
                      class="image-size"
                      :src="item.fileUrl"
                      :on-error="errorImg"
                    />
                    <n-avatar v-else class="image-size">
                      <span style="font-size: 24px"> {{ getFileExt(item.fileUrl) }}</span>
                    </n-avatar>
                  </div>
                  <template #footer>
                    <n-ellipsis style="padding-left: 5px">
                      {{ item.name }}
                    </n-ellipsis>
                  </template>

                  <template #action>
                    <n-space justify="center" style="padding: 5px">
                      <n-button
                        strong
                        secondary
                        size="tiny"
                        type="primary"
                        @click="item.kind === 'image' ? handlePreview(item) : handleDown(item)"
                      >
                        <template #icon>
                          <n-icon>
                            <EyeOutlined v-if="item.kind === 'image'" />
                            <DownloadOutlined v-else />
                          </n-icon>
                        </template>
                        {{ item.kind === 'image' ? '预览' : '下载' }}
                      </n-button>
                      <n-button
                        strong
                        secondary
                        size="tiny"
                        type="error"
                        @click="handleDelete(item)"
                      >
                        <template #icon>
                          <n-icon>
                            <DeleteOutlined />
                          </n-icon>
                        </template>
                        删除
                      </n-button>
                    </n-space>
                  </template>
                </n-card>
              </n-grid-item>
            </n-grid>
          </n-layout-content>
          <n-space
            v-if="!isEmptyDataSource()"
            justify="end"
            align="center"
            style="box-sizing: border-box; padding: 5px; margin: 0"
          >
            <n-pagination
              v-model:page="params.page"
              v-model:page-size="params.pageSize"
              :page-count="params.pageCount"
              :page-slot="pageSlot"
              :page-sizes="[10, 20, 30, 40]"
              :on-update:page="onUpdatePage"
              :on-update:page-size="onUpdatePageSize"
              show-size-picker
              show-quick-jumper
            />
          </n-space>
        </n-layout>
      </n-spin>
    </n-gi>
  </n-grid>
  <Preview ref="previewRef" />
</template>

<script lang="ts" setup>
  import { NSpace, NInput, NButton, useMessage, NEllipsis, useDialog } from 'naive-ui';
  import { onMounted, ref, h, computed } from 'vue';
  import { BasicForm, FormSchema, useForm } from '@/components/Form';
  import { defRangeShortcuts } from '@/utils/dateUtil';
  import { EyeOutlined, DeleteOutlined, DownloadOutlined, ClearOutlined } from '@vicons/antd';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { ClearKind, Delete, List } from '@/api/apply/attachment';
  import { constantRouterIcon } from '@/router/router-icons';
  import { errorImg } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';
  import { renderIcon } from '@/utils';
  import { Attachment, FileType, KindOption } from './model';
  import Preview from './Preview.vue';
  import { VNode } from 'vue';
  import { Option, useDictStore } from '@/store/modules/dict';

  export interface Props {
    fileList: string[] | null;
    maxNumber?: number;
    fileType?: FileType;
  }

  const props = withDefaults(defineProps<Props>(), {
    fileList: null,
    maxNumber: 1,
    fileType: 'default',
  });

  const emit = defineEmits(['saveChange']);
  const dict = useDictStore();
  const settingStore = useProjectSettingStore();
  const message = useMessage();
  const dialog = useDialog();
  const selectList = ref(props.fileList);
  const loading = ref(false);
  const previewRef = ref();
  const searchFormRef = ref<any>();
  const dataSource = ref<Attachment[]>([]);

  const defaultKindValue = computed(() => {
    if (props.fileType === 'default') {
      return '';
    }
    return props.fileType;
  });

  const pageSlot = computed(() => (settingStore.isMobile ? 3 : 10));

  const imageGridCols = computed(() =>
    settingStore.isMobile ? '2 s:1 m:2 l:2 xl:2 2xl:3' : '5 s:3 m:4 l:5 xl:5 2xl:6'
  );

  const pageGridCols = computed(() =>
    settingStore.isMobile ? '1' : '10 s:1 m:1 l:10 xl:10 2xl:10'
  );

  const options = ref({
    drive: [],
    kind: [],
  });

  const schemas = ref<FormSchema[]>([
    {
      field: 'drive',
      component: 'NSelect',
      label: '',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择上传驱动',
        options: dict.getOption('config_upload_drive'),
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'name',
      component: 'NInput',
      label: '',
      defaultValue: null,
      componentProps: {
        placeholder: '请输入文件名称',
        options: options.value.drive,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'updatedAt',
      component: 'NDatePicker',
      label: '',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        shortcuts: defRangeShortcuts(),
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ]);

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:2 xl:3 2xl:3' },
    labelWidth: 80,
    schemas,
  });

  const params = ref({
    kind: defaultKindValue.value,
    page: 1,
    pageSize: 10,
    pageCount: 0,
  });

  function handleDelete(item: Attachment) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: item.id }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function handlePreview(item: Attachment) {
    previewRef.value.openPreview(item.fileUrl);
  }

  function handleDown(item: Attachment) {
    window.open(item.fileUrl);
  }

  function handleSelect(item: Attachment) {
    if (!selectList.value || props.maxNumber == 1) {
      selectList.value = [];
    }

    const index = selectList.value.findIndex((selected) => selected === item.fileUrl);
    if (index === -1) {
      if (selectList.value.length >= props.maxNumber) {
        message.error('已达最大允许选择上限' + props.maxNumber + '个');
        return;
      }
      selectList.value.push(item.fileUrl);
    } else {
      selectList.value.splice(index, 1);
    }
    emit('saveChange', selectList.value);
  }

  function handleUpdateKind(value: string) {
    params.value.page = 1;
    params.value.kind = value;
    loadList();
  }

  function isSelected(item: Attachment) {
    if (!selectList.value) {
      return false;
    }
    return selectList.value.some((selected) => selected === item.fileUrl);
  }

  function generateKindOptions(): any {
    const option: KindOption[] = [];
    dict.getOptionUnRef('AttachmentKindOption').forEach((item: Option) => {
      const data: KindOption = {
        label: () => h(NEllipsis, null, { default: () => item.label }),
        key: item.key.toString(),
        extra: () => (item.key !== '' ? createExtraContent(item) : null),
        icon: constantRouterIcon[item.extra] || null,
        disabled: isDisabledKindOption(item),
      };
      option.push(data);
    });
    return option;
  }

  function isDisabledKindOption(item: Option): boolean {
    if (props.fileType === 'default') {
      return false;
    }
    return item.key !== props.fileType;
  }

  function createExtraContent(item: Option): VNode {
    return h(
      NButton,
      {
        quaternary: true,
        type: 'default',
        size: 'tiny',
        style: 'position: absolute; right: 15px;',
        onClick: (event: MouseEvent) => {
          event.stopPropagation();
          dialog.warning({
            title: '警告',
            content: '你确定要清空 [' + item.label + '] 分类？该操作不可恢复！',
            positiveText: '确定',
            negativeText: '取消',
            onPositiveClick: () => {
              ClearKind({ kind: item.key }).then((_res) => {
                message.success('操作成功');
                reloadTable();
              });
            },
          });
        },
      },
      { icon: renderIcon(ClearOutlined) }
    );
  }

  function loadList() {
    loading.value = true;
    List({ ...params.value, ...searchFormRef.value?.formModel }).then((res) => {
      dataSource.value = res.list;
      params.value.page = res.page;
      params.value.pageSize = res.pageSize;
      params.value.pageCount = res.pageCount;
      loading.value = false;
    });
  }

  function isEmptyDataSource(): boolean {
    return !dataSource.value || dataSource.value.length === 0;
  }

  function onUpdatePage(page: number) {
    params.value.page = page;
    loadList();
  }

  function onUpdatePageSize(pageSize: number) {
    params.value.pageSize = pageSize;
    loadList();
  }

  function reloadTable() {
    params.value.page = 1;
    loadList();
  }

  function loadOptions() {
    dict.loadOptions(['AttachmentKindOption', 'config_upload_drive']);
  }

  onMounted(async () => {
    loadOptions();
    loadList();
  });

  defineExpose({
    reloadTable,
  });
</script>

<style lang="less" scoped>
  .base-list .n-spin-content {
    height: 100% !important;
  }

  :deep(.n-card__action) {
    padding: 5px;
  }

  .imageActive {
    border: 2px solid #3086ff;
  }

  :deep(img, video) {
    max-width: 100%;
    aspect-ratio: 1/1;
  }

  :deep(.image-size) {
    aspect-ratio: 1/1;
    width: 100%;
    height: auto;
  }
</style>
