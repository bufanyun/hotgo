<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <BasicTable
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
        :flex-height="false"
        :pagination="{ pageSize: 10 }"
        :resizeHeightOffset="-50000"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加分组
          </n-button>
        </template>
      </BasicTable>

      <n-modal v-model:show="showModal" :show-icon="false" preset="dialog" :title="modalTitle">
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="上级分组" path="pid">
            <n-tree-select
              :options="optionTreeData"
              :default-value="formParams.pid"
              @update:value="handleUpdateValue"
            />
          </n-form-item>
          <n-form-item label="分组名称" path="name">
            <n-input placeholder="请输入分组名称" v-model:value="formParams.name" />
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-radio-group v-model:value="formParams.status" name="status">
              <n-radio-button
                v-for="status in statusOptions"
                :key="status.value"
                :value="status.value"
                :label="status.label"
              />
            </n-radio-group>
          </n-form-item>

<!--          <n-form-item label="排序" path="sort">-->
<!--            <n-input-number v-model:value="formParams.sort" clearable />-->
<!--          </n-form-item>-->

          <n-form-item label="备注" path="remark">
            <n-input type="textarea" placeholder="请输入备注" v-model:value="formParams.remark" />
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { GroupDelete, GroupEdit, GroupList, getSelect } from '@/api/sys/cron';
  import { statusOptions } from '@/enums/optionsiEnum';

  const emit = defineEmits(['reloadGroupOption']);
  const optionTreeData = ref<any>([]);
  const message = useMessage();
  const statusValue = ref(1);
  const defaultValueRef = () => ({
    id: 0,
    pid: 0,
    name: '',
    sort: 0,
    remark: '',
    status: statusValue.value,
  });
  const modalTitle = ref('添加分组');
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const rules = {
    name: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入名称',
    },
  };

  function addTable() {
    showModal.value = true;
    modalTitle.value = '添加分组';
    formParams.value = defaultValueRef();
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        GroupEdit(formParams.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showModal.value = false;
            reloadTable();
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  const dialog = useDialog();
  const actionRef = ref();
  const formParams = ref<any>(defaultValueRef);
  const formRef = ref<any>({});

  const actionColumn = reactive({
    width: 150,
    title: '操作',
    key: 'action',
    // fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
      });
    },
  });

  function handleEdit(record: Recordable) {
    showModal.value = true;
    modalTitle.value = '编辑分组 ID：' + record.id;
    formParams.value = {
      id: record.id,
      pid: record.pid,
      name: record.name,
      sort: record.sort,
      remark: record.remark,
      status: record.status,
    };
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        GroupDelete(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  const dataSource = ref({
    successful_order: 0,
    transaction_money: 0,
  });

  const loadDataTable = async (res) => {
    dataSource.value = await GroupList({ ...res });
    return dataSource.value;
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
  }

  function reloadTable() {
    actionRef.value.reload();
    emit('reloadGroupOption');
  }

  async function setDictSelect() {
    const tmp = await getSelect({});
    optionTreeData.value = [
      {
        id: 0,
        key: 0,
        label: '顶级分组',
        pid: 0,
        name: '顶级分组',
      },
    ];
    optionTreeData.value = optionTreeData.value.concat(tmp.list);
  }

  onMounted(async () => {
    await setDictSelect();
  });

  // 处理选项更新
  function handleUpdateValue(value: string | number | Array<string | number> | null) {
    formParams.value.pid = value;
  }
</script>

<style lang="less" scoped></style>
