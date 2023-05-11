<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <BasicForm
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

      <BasicTable
        :openChecked="false"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :scroll-x="1800"
        :resizeHeightOffset="-10000"
        size="small"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="handleExport"
            class="min-left-space"
            v-if="hasPermission(['/creditsLog/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { BasicTable } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export } from '@/api/creditsLog';
  import { columns, schemas } from './model';
  import { ExportOutlined } from '@vicons/antd';

  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '',
  });

  const { hasPermission } = usePermission();
  const actionRef = ref();
  const message = useMessage();
  const searchFormRef = ref<any>({});

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  const loadDataTable = async (res) => {
    return await List({
      ...searchFormRef.value?.formModel,
      ...res,
      ...{ creditType: props.type },
    });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }
</script>

<style lang="less" scoped></style>
