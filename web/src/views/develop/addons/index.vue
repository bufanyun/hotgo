<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="插件管理">
        在这里可以管理和创建插件模板，开发独立、临时性、工具类型的功能时推荐使用插件化开发。
      </n-card>
    </div>

    <n-card :bordered="false" class="proCard">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
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
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            创建新插件
          </n-button>
        </template>
      </BasicTable>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        preset="dialog"
        title="创建新插件"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-alert :show-icon="false" type="info">
          注意：插件创建成功后会在服务端对应的项目目录中生成插件模块文件，并自动注册到当前项目中。
        </n-alert>
        <n-form
          :model="formParams"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="插件标签" path="label">
            <n-input placeholder="请输入" v-model:value="formParams.label" />
            <template #feedback>显示在插件列表中的标识. 不要超过20个字符</template>
          </n-form-item>

          <n-form-item label="插件包名" path="name">
            <n-input placeholder="请输入" v-model:value="formParams.name" />
            <template #feedback>
              系统按照此名称创建和查找插件定义， 只能英文和下划线组成，建议全部小写，例如：hgexample
            </template>
          </n-form-item>

          <n-form-item label="功能分组" path="group">
            <n-select
              placeholder="请选择"
              :options="selectList.groupType"
              v-model:value="formParams.group"
              :on-update:value="onUpdateValueGroup"
            />
          </n-form-item>

          <n-form-item label="版本" path="version">
            <n-input placeholder="请输入" v-model:value="formParams.version" />
            <template #feedback>此版本号用于插件的版本更新</template>
          </n-form-item>

          <n-form-item label="简介" path="brief">
            <n-input placeholder="请输入" v-model:value="formParams.brief" />
          </n-form-item>

          <n-form-item label="使用说明" path="description">
            <n-input
              type="textarea"
              placeholder="使用说明"
              v-model:value="formParams.description"
            />
            <template #feedback>详细介绍插件的功能和使用方法</template>
          </n-form-item>

          <n-form-item label="作者" path="author">
            <n-input placeholder="请输入" v-model:value="formParams.author" />
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确认创建</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, onBeforeMount, reactive, ref } from 'vue';
  import {
    NIcon,
    NTag,
    NIconWrapper,
    useMessage,
    NImage,
    useDialog,
    useNotification,
  } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { List, Selects, Build, UnInstall, Install, Upgrade } from '@/api/develop/addons';
  import { PlusOutlined } from '@vicons/antd';
  import { newState } from '@/views/develop/addons/components/model';
  import { errorImg } from '@/utils/hotgo';
  import { isUrl } from '@/utils/is';
  import { getIconComponent } from '@/utils/icons';

  const dialog = useDialog();
  const message = useMessage();
  const notification = useNotification();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formRef: any = ref(null);
  const actionRef = ref();
  const formParams = ref<any>();
  const dialogWidth = ref('50%');
  const checkedIds = ref([]);
  const searchFormRef = ref<any>();

  const selectList = ref({
    groupType: [],
    status: [],
  });

  const columns = [
    {
      title: '图标',
      key: 'logo',
      width: 80,
      render(row) {
        if (isUrl(row.logo)) {
          return h(NImage, {
            width: 48,
            height: 48,
            src: row.logo,
            onError: errorImg,
            style: {
              width: '48px',
              height: '48px',
              'max-width': '100%',
              'max-height': '100%',
            },
          });
        } else {
          return h(
            NIconWrapper,
            {
              size: 48,
              borderRadius: 8,
            },
            {
              default: () =>
                h(
                  NIcon,
                  {
                    size: 36,
                    style: {
                      marginTop: '-8px',
                    },
                  },
                  {
                    default: () => h(getIconComponent(row.logo)),
                  }
                ),
            }
          );
        }
      },
    },
    {
      title: '模块名称',
      key: 'name',
      width: 120,
      render(row) {
        return h('div', {
          innerHTML:
            '<div >' + row.label + '<br><span style="opacity: 0.8;">' + row.name + '</span></div>',
        });
      },
    },
    {
      title: '作者',
      key: 'author',
      width: 100,
    },
    {
      title: '分组',
      key: 'groupName',
      render(row) {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: 'info',
            bordered: false,
          },
          {
            default: () => row.groupName,
          }
        );
      },
      width: 120,
    },
    {
      title: '简介',
      key: 'brief',
      render(row) {
        return row.brief;
      },
      width: 180,
    },
    {
      title: '详细描述',
      key: 'description',
      width: 300,
      render(row) {
        return h('p', { id: 'app' }, [
          h('div', {
            innerHTML: '<div style="white-space: pre-wrap">' + row.description + '</div>',
          }),
        ]);
      },
    },
    {
      title: '版本',
      key: 'version',
      width: 100,
    },
  ];

  const schemas = ref<FormSchema[]>([
    {
      field: 'name',
      component: 'NInput',
      label: '模块名称',
      componentProps: {
        placeholder: '请输入模块名称或标签',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'group',
      component: 'NSelect',
      label: '分组',
      componentProps: {
        placeholder: '请选择分组',
        options: [],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '安装状态',
      componentProps: {
        placeholder: '请选择状态',
        options: [],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ]);

  const actionColumn = reactive({
    width: 220,
    title: '操作',
    key: 'action',
    // fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '安装模块',
            onClick: handleInstall.bind(null, record),
            ifShow: () => {
              return record.installStatus !== 1;
            },
          },
          {
            type: 'warning',
            label: '更新配置',
            onClick: handleUpgrade.bind(null, record),
            ifShow: () => {
              return record.installStatus === 1;
            },
          },
          {
            type: 'error',
            label: '卸载模块',
            onClick: handleUnInstall.bind(null, record),
            ifShow: () => {
              return record.installStatus === 1;
            },
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

  function onCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  const loadDataTable = async (res) => {
    mapWidth();
    return await List({ ...res, ...searchFormRef.value?.formModel });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleInstall(record: Recordable) {
    dialog.info({
      title: '提示',
      content: '你确定要安装 【' + record.label + '】 模块？安装成功后需要重启服务才能生效！',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Install(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleUpgrade(record: Recordable) {
    dialog.warning({
      title: '提示',
      content: '你确定要更新 【' + record.label + '】 模块？更新成功后需要重启服务才能生效！',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Upgrade(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleUnInstall(record: Recordable) {
    dialog.error({
      title: '提示',
      content: '你确定要卸载 【' + record.label + '】 模块？卸载成功后需要重启服务才能生效！',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        UnInstall(record).then((_res) => {
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

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        dialog.info({
          title: '提示',
          content:
            '你确定要提交生成吗？热编译生成后如果列表自动刷新没有出现新插件请等几秒刷新即可。否则请手动重启服务后刷新页面！',
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: () => {
            Build(formParams.value).then((_res) => {
              showModal.value = false;
              buildSuccessNotify();
            });
          },
          onNegativeClick: () => {
            // message.error('取消');
          },
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function mapWidth() {
    let val = document.body.clientWidth;
    const def = 840; // 默认宽度
    if (val < def) {
      dialogWidth.value = '100%';
    } else {
      dialogWidth.value = def + 'px';
    }

    return dialogWidth.value;
  }

  const loadSelect = async () => {
    selectList.value = await Selects({});
    for (const item of schemas.value) {
      switch (item.field) {
        case 'status':
          item.componentProps.options = selectList.value.status;
          break;
        case 'group':
          item.componentProps.options = selectList.value.groupType;
          break;
      }
    }
  };

  function onUpdateValueGroup(value) {
    formParams.value.group = value;
  }

  function buildSuccessNotify() {
    let count = 10;
    const n = notification.success({
      title: '插件构建成功',
      content: `如果你使用的热编译，页面将在 ${count} 秒后自动刷新插件列表即可生效。否则请手动重启服务后刷新页面！`,
      duration: 10000,
      closable: false,
      onAfterEnter: () => {
        const minusCount = () => {
          count--;
          n.content = `如果你使用的热编译，页面将在 ${count} 秒后自动刷新插件列表即可生效。否则请手动重启服务后刷新页面！`;
          if (count > 0) {
            window.setTimeout(minusCount, 1000);
          }
        };
        window.setTimeout(minusCount, 1000);
      },
      onAfterLeave: () => {
        location.reload();
      },
    });
  }

  onBeforeMount(async () => {
    await loadSelect();
  });
</script>

<style lang="less" scoped></style>
