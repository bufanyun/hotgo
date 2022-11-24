<template>
  <div>
    <n-card :bordered="false" class="proCard" title="省市区">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        @keyup.enter="handleSubmit"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>

      <BasicTable
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
        :resizeHeightOffset="-10000"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            新建
          </n-button>
          &nbsp;
          <n-button type="error" @click="batchDelete" :disabled="batchDeleteDisabled">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>

      <n-modal v-model:show="showModal" :show-icon="false" preset="dialog" title="新建">
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="公告标题" path="title">
            <n-input placeholder="请输入公告标题" v-model:value="formParams.title" />
          </n-form-item>

          <n-form-item label="公告类型" path="type">
            <n-radio-group v-model:value="formParams.type" name="type">
              <n-radio-button
                v-for="type in typeOptions"
                :key="type.value"
                :value="type.value"
                :label="type.label"
              />
            </n-radio-group>
          </n-form-item>

          <n-form-item label="公告内容" path="content">
            <n-input type="textarea" placeholder="请输入内容" v-model:value="formParams.content" />
          </n-form-item>

          <n-form-item label="接收人" path="receiver">
            <n-input
              type="textarea"
              placeholder="多个用户ID用,隔开 、不填则全部接收"
              v-model:value="formParams.receiver"
            />
          </n-form-item>

          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="formParams.sort" clearable />
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
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Delete, Edit, List, Status } from '@/api/apply/provinces';
  import { columns } from './columns';
  import { DeleteOutlined, PlusOutlined } from '@vicons/antd';
  import { statusActions, statusOptions } from '@/enums/optionsiEnum';

  const typeOptions = [
    {
      value: 1,
      label: '通知',
    },
    {
      value: 2,
      label: '公告',
    },
  ].map((s) => {
    return s;
  });
  const params = ref({
    pageSize: 10,
    title: '',
    content: '',
    status: null,
  });

  const rules = {
    title: {
      // required: true,
      trigger: ['blur', 'input'],
      message: '请输入标题',
    },
  };

  const schemas: FormSchema[] = [
    {
      field: 'title',
      component: 'NInput',
      label: '公告标题',
      componentProps: {
        placeholder: '请输入公告标题',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入公告标题', trigger: ['blur'] }],
    },
    {
      field: 'content',
      component: 'NInput',
      label: '内容',
      componentProps: {
        placeholder: '请输入内容关键词',
        showButton: false,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '状态',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择类型',
        options: statusOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref({});
  const formRef = ref({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

  const resetFormParams = {
    id: 0,
    title: '',
    name: '',
    type: 1,
    receiver: '',
    remark: '',
    sort: 0,
    status: 1,
    created_at: '',
    updated_at: '',
  };
  let formParams = ref(resetFormParams);

  const actionColumn = reactive({
    width: 220,
    title: '操作',
    key: 'action',
    fixed: 'right',
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
        dropDownActions: statusActions,
        select: (key) => {
          updateStatus(record.id, key);
        },
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function addTable() {
    showModal.value = true;
    formParams.value = resetFormParams;
  }

  const loadDataTable = async (res) => {
    return await List({ ...params.value, ...res, ...searchFormRef.value.formModel });
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
    if (rowKeys.length > 0) {
      batchDeleteDisabled.value = false;
    } else {
      batchDeleteDisabled.value = true;
    }

    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        console.log('formParams:' + JSON.stringify(formParams.value));
        Edit(formParams.value)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            setTimeout(() => {
              showModal.value = false;
              reloadTable();
              formParams.value = ref(resetFormParams);
            });
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    showModal.value = true;
    formParams.value = record;
  }

  function handleDelete(record: Recordable) {
    console.log('点击了删除', record);
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete(record)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value })
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    params.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function updateStatus(id, status) {
    Status({ id: id, status: status })
      .then((_res) => {
        console.log('_res:' + JSON.stringify(_res));
        message.success('操作成功');
        setTimeout(() => {
          reloadTable({});
        });
      })
      .catch((e: Error) => {
        message.error(e.message ?? '操作失败');
      });
  }
</script>

<style lang="less" scoped></style>
