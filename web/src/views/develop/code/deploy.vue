<template>
  <div>
    <n-spin :show="show" description="正在生成配置信息...">
      <n-card>
        <n-tabs
          type="card"
          class="card-tabs"
          :default-value="value"
          animated
          tab-style="min-width: 80px;"
          style="margin: 0 -4px"
          pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
          :on-update:value="updateTabs"
          ref="tabsRef"
          @close="handleClose"
          @add="handleAdd"
        >
          <n-tab-pane v-for="panel in panels" :key="panel" :name="panel">
            <template v-if="panel === '基本信息'">
              <BaseInfo v-model:value="genInfo" :selectList="selectList" />
            </template>

            <template v-if="panel === '主表字段'">
              <EditMasterCell v-model:value="genInfo" :selectList="selectList" />
            </template>
          </n-tab-pane>

          <n-tab-pane
            v-for="panel in slavePanels"
            :key="panel"
            :name="panel"
            v-show="slavePanels.length > 0 && slavePanels !== []"
          >
            <EditSlaveCell
              v-model:value="genInfo"
              :uuid="slaveMap[panel]"
              :selectList="selectList"
            />
          </n-tab-pane>

          <template #suffix>
            <n-space>
              <n-button type="primary" @click="preview">预览代码</n-button>
              <n-button type="success" :loading="formBtnLoading" @click="submitBuild"
                >提交生成</n-button
              >
              <n-button type="info" dashed :loading="formBtnLoading" @click="submitSave"
                >仅保存配置</n-button
              >
            </n-space>
          </template>
        </n-tabs>

        <n-modal
          v-model:show="showModal"
          :block-scroll="false"
          :mask-closable="false"
          :show-icon="false"
          preset="card"
          title="预览代码"
          style="width: 95%"
        >
          <PreviewTab :previewModel="previewModel" />
          <template #action>
            <n-space justify="end">
              <n-button @click="() => (showModal = false)">关闭</n-button>
              <n-button type="info" :loading="formBtnLoading" @click="submitBuild"
                >提交生成</n-button
              >
            </n-space>
          </template>
        </n-modal>
      </n-card>
    </n-spin>
  </div>
</template>

<style scoped>
  .card-tabs .n-tabs-nav--bar-type {
    padding-left: 4px;
  }
</style>

<script lang="ts" setup>
  import { onMounted, ref, watch } from 'vue';
  import { useRouter } from 'vue-router';
  import { useDialog, useMessage } from 'naive-ui';
  import BaseInfo from './components/BaseInfo.vue';
  import EditMasterCell from './components/EditMasterCell.vue';
  import EditSlaveCell from './components/EditSlaveCell.vue';
  import { Selects, View, Preview, Build, Edit } from '@/api/develop/code';
  import { selectListObj, newState } from '@/views/develop/code/components/model';
  import PreviewTab from '@/views/develop/code/components/PreviewTab.vue';
  import { isJsonString } from '@/utils/is';

  interface Props {
    genId?: number;
  }
  const props = withDefaults(defineProps<Props>(), { genId: 0 });
  const router = useRouter();
  const genId = Number(router.currentRoute.value.params.id ?? props.genId);
  const show = ref(false);
  const message = useMessage();
  const selectList = ref<any>(selectListObj);
  const genInfo = ref(newState(null));
  const tabsRef = ref();
  const value = ref('基本信息');
  const panels = ref(['基本信息', '主表字段']);
  const slaveMap = ref<any>([]);
  const slavePanels = ref<any>([]);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const previewModel = ref<any>();
  const dialog = useDialog();

  onMounted(async () => {
    if (genId < 1 && props.genId < 1) {
      message.error('生成ID不正确，请检查！');
      return;
    }
    await getGenInfo();
    await loadSelect();
  });

  async function getGenInfo() {
    let tmp = await View({ id: genId });
    if (isJsonString(tmp.options)) {
      tmp.options = JSON.parse(tmp.options);
    }

    if (tmp.masterColumns === undefined || tmp.masterColumns.length === 0) {
      tmp.masterColumns = [];
    }
    if (isJsonString(tmp.masterColumns)) {
      tmp.masterColumns = JSON.parse(tmp.masterColumns);
    }

    genInfo.value = tmp;
  }

  watch(
    genInfo,
    (newVal, _oldVal) => {
      if (newVal.genType >= 10 && newVal.genType < 20) {
        handleAdd('主表字段');
      } else {
        handleClose('主表字段');
      }

      if (newVal.options.join !== undefined) {
        slavePanels.value = [];
        for (let i = 0; i <= newVal.options.join.length; i++) {
          if (newVal.options.join[i]?.alias !== undefined && newVal.options.join[i]?.alias !== '') {
            handleSlaveAdd(
              '关联表[ ' + newVal.options.join[i]?.alias + ' ]',
              newVal.options.join[i]
            );
          }
        }
      }
    },
    {
      deep: true, // 是否深度监听
    }
  );

  function updateTabs(value: string | number) {
    console.log('value:' + value);
  }

  function handleAdd(name: string) {
    const nameIndex = panels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) {
      panels.value.push(name);
    }
  }

  function handleSlaveAdd(name: string, join) {
    const nameIndex = slavePanels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) {
      slavePanels.value.push(name);
      slaveMap.value[name] = join.uuid;
    }
  }

  function _handleSlaveClose(name: string) {
    const nameIndex = panels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) return;
    panels.value.splice(nameIndex, 1);
    if (name === value.value) {
      value.value = panels.value[Math.min(nameIndex, panels.value.length - 1)];
    }
  }

  function handleClose(name: string) {
    const nameIndex = panels.value.findIndex((panelName) => panelName === name);
    if (!~nameIndex) return;
    panels.value.splice(nameIndex, 1);
    if (name === value.value) {
      value.value = panels.value[Math.min(nameIndex, panels.value.length - 1)];
    }
  }

  const loadSelect = async () => {
    selectList.value = await Selects({});
  };

  async function preview() {
    previewModel.value = await Preview(genInfo.value);
    showModal.value = true;
  }

  function submitBuild() {
    dialog.warning({
      title: '警告',
      content: '你确定要提交生成吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Build(genInfo.value).then((_res) => {
          setTimeout(function () {
            location.reload();
          }, 1500);
          message.success('生成提交成功，即将刷新页面..');
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function submitSave() {
    dialog.warning({
      title: '警告',
      content: '你确定要保存生成配置吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Edit(genInfo.value).then((_res) => {
          message.success('操作成功');
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }
</script>
<style lang="less" scoped>
  ::v-deep(.alert-margin) {
    margin-bottom: 20px;
  }
  ::v-deep(.tag-margin) {
    margin-bottom: 10px;
  }

  ::v-deep(.code-scrollbar) {
    height: calc(100vh - 300px);
    background: #282b2e;
    color: #e0e2e4;
    padding: 10px;
  }
</style>
