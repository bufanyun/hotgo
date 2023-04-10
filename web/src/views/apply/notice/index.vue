<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="通知公告">
        在这里你可以发送通知、公告、私信到平台中的用户
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
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
        :openChecked="true"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1800"
        :resizeHeightOffset="-20000"
      >
        <template #tableTitle>
          <n-button
            type="warning"
            @click="addTable(1)"
            class="min-left-space"
            v-if="hasPermission(['/notice/editNotify'])"
          >
            <template #icon>
              <n-icon>
                <NotificationOutlined />
              </n-icon>
            </template>
            发通知
          </n-button>

          <n-button
            type="error"
            @click="addTable(2)"
            class="min-left-space"
            v-if="hasPermission(['/notice/editNotice'])"
          >
            <template #icon>
              <n-icon>
                <BellOutlined />
              </n-icon>
            </template>
            发公告
          </n-button>

          <n-button
            type="info"
            @click="addTable(3)"
            class="min-left-space"
            v-if="hasPermission(['/notice/editLetter'])"
          >
            <template #icon>
              <n-icon>
                <SendOutlined />
              </n-icon>
            </template>
            发私信
          </n-button>

          <n-button
            type="error"
            @click="batchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/notice/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        :block-scroll="false"
        :mask-closable="false"
        preset="dialog"
        :title="
          formParams.id > 0
            ? '编辑' + getOptionLabel(noticeTypeOptions, formParams.type) + ' #' + formParams.id
            : '发送' + getOptionLabel(noticeTypeOptions, formParams.type)
        "
        :style="{
          width: dialogWidth,
        }"
      >
        <n-alert :show-icon="false" type="info">
          消息发送成功后如果接收人在线会立即收到一条消息通知，编辑已发送的消息不会再次通知
        </n-alert>
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="消息标题" path="title">
            <n-input placeholder="请输入消息标题" v-model:value="formParams.title" />
          </n-form-item>

          <n-form-item label="接收人" path="receiver" v-if="formParams.type === 3">
            <n-select
              multiple
              :options="options"
              :render-label="renderLabel"
              :render-tag="renderMultipleSelectTag"
              v-model:value="formParams.receiver"
              filterable
            />
          </n-form-item>

          <n-form-item label="消息内容" path="content">
            <template v-if="formParams.type === 1">
              <n-input
                type="textarea"
                :autosize="{ minRows: 3, maxRows: 30 }"
                placeholder="请输入通知内容"
                v-model:value="formParams.content"
              />
            </template>
            <template v-else>
              <Editor style="height: 450px" v-model:value="formParams.content" />
            </template>
          </n-form-item>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="标签" path="tag">
                <n-select
                  clearable
                  placeholder="可以不填"
                  :render-tag="renderTag"
                  v-model:value="formParams.tag"
                  :options="noticeTagOptions"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="排序" path="sort">
                <n-input-number style="width: 100%" v-model:value="formParams.sort" clearable />
              </n-form-item>
            </n-gi>
          </n-grid>

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
            <n-input
              type="textarea"
              placeholder="请输入备注，没有可以不填"
              v-model:value="formParams.remark"
            />
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">立即发送</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import {
    Delete,
    EditNotify,
    EditLetter,
    EditNotice,
    List,
    MaxSort,
    Status,
  } from '@/api/apply/notice';
  import { columns } from './columns';
  import { BellOutlined, DeleteOutlined, NotificationOutlined, SendOutlined } from '@vicons/antd';
  import { statusOptions } from '@/enums/optionsiEnum';
  import {
    noticeTagOptions,
    noticeTypeOptions,
    personOption,
    renderLabel,
    renderMultipleSelectTag,
  } from '@/enums/systemMessageEnum';
  import { adaModalWidth, getOptionLabel, renderTag } from '@/utils/hotgo';
  import Editor from '@/components/Editor/editor.vue';
  import { cloneDeep } from 'lodash-es';
  import { GetMemberOption } from '@/api/org/user';
  import { usePermission } from '@/hooks/web/usePermission';
  const { hasPermission } = usePermission();

  const rules = {
    title: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入消息标题',
    },
  };

  const schemas: FormSchema[] = [
    {
      field: 'type',
      component: 'NSelect',
      label: '消息类型',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择消息类型',
        options: noticeTypeOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'title',
      component: 'NInput',
      label: '消息标题',
      componentProps: {
        placeholder: '请输入消息标题',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入消息标题', trigger: ['blur'] }],
    },
    {
      field: 'content',
      component: 'NInput',
      label: '消息内容',
      componentProps: {
        placeholder: '请输入消息内容关键词',
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
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const dialogWidth = ref('75%');
  const options = ref<personOption[]>();

  const resetFormParams = {
    id: 0,
    title: '',
    type: 1,
    tag: 0,
    content: '',
    receiver: null,
    remark: '',
    sort: 0,
    status: 1,
  };
  let formParams = ref<any>(cloneDeep(resetFormParams));

  const actionColumn = reactive({
    width: 180,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '已启用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/notice/status'],
          },
          {
            label: '已禁用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/notice/status'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/notice/edit'],
            type: 'primary',
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/notice/delete'],
          },
        ],
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function addTable(type) {
    showModal.value = true;
    formParams.value = cloneDeep(resetFormParams);
    formParams.value.type = type;
    MaxSort().then((res) => {
      formParams.value.sort = res.sort;
    });
  }

  const loadDataTable = async (res) => {
    return await List({ ...res, ...searchFormRef.value?.formModel });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
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
        switch (formParams.value.type) {
          case 1:
            EditNotify(formParams.value).then((_res) => {
              message.success('操作成功');
              setTimeout(() => {
                showModal.value = false;
                reloadTable();
              });
            });
            break;
          case 2:
            EditNotice(formParams.value).then((_res) => {
              message.success('操作成功');
              setTimeout(() => {
                showModal.value = false;
                reloadTable();
              });
            });
            break;
          case 3:
            EditLetter(formParams.value).then((_res) => {
              message.success('操作成功');
              setTimeout(() => {
                showModal.value = false;
                reloadTable();
              });
            });
            break;
          default:
            message.error('公告类型不支持');
        }
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = cloneDeep(record);
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  function handleStatus(record: Recordable, status: number) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  async function getMemberOption() {
    options.value = await GetMemberOption();
  }

  onMounted(async () => {
    adaModalWidth(dialogWidth);
    await getMemberOption();
  });
</script>

<style lang="less" scoped></style>
