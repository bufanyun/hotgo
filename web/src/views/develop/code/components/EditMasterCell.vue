<template>
  <n-spin :show="show" description="加载中...">
    <n-card :bordered="false" class="proCard">
      <BasicTable
        :single-line="false"
        size="small"
        :striped="true"
        :resizable="true"
        :columns="columns"
        :dataSource="dataSource"
        :openChecked="false"
        :showTopRight="false"
        :row-key="(row) => row.id"
        ref="actionRef"
        :canResize="true"
        :resizeHeightOffset="-20000"
        :pagination="false"
        :scroll-x="1090"
        :scrollbar-props="{ trigger: 'none' }"
      >
        <template #tableTitle>
          <n-tooltip placement="top-start" trigger="hover">
            <template #trigger>
              <n-button type="primary" @click="reloadFields(true)" class="min-left-space">
                <template #icon>
                  <n-icon>
                    <Reload />
                  </n-icon>
                </template>
                重置字段
              </n-button>
            </template>
            主要用于重置字段设置或数据库表字段发生变化时重新载入
          </n-tooltip>
        </template>
      </BasicTable>
    </n-card>
  </n-spin>
</template>

<script lang="ts" setup>
  import { computed, h, onMounted, ref } from 'vue';
  import { BasicTable } from '@/components/Table';
  import { genInfoObj, selectListObj } from '@/views/develop/code/components/model';
  import { ColumnList } from '@/api/develop/code';
  import { NButton, NCheckbox, NInput, NSelect, NTooltip, NTreeSelect } from 'naive-ui';
  import { HelpCircleOutline, Reload } from '@vicons/ionicons5';
  import { renderIcon } from '@/utils';
  import { cloneDeep } from 'lodash-es';

  const renderTooltip = (trigger, content) => {
    return h(NTooltip, null, {
      trigger: () => trigger,
      default: () => content,
    });
  };

  const emit = defineEmits(['update:value']);

  interface Props {
    value?: any;
    selectList: any;
  }

  const props = withDefaults(defineProps<Props>(), {
    value: genInfoObj,
    selectList: selectListObj,
  });

  const formValue = computed({
    get() {
      return props.value;
    },
    set(value) {
      emit('update:value', value);
    },
  });

  const actionRef = ref();
  const columns = ref<any>([]);
  const show = ref(false);
  const dataSource = ref(formValue.value.masterColumns);

  async function reloadFields(loading = false) {
    dataSource.value = [];
    if (loading) {
      show.value = true;
    }

    formValue.value.masterColumns = await ColumnList({
      name: formValue.value.dbName,
      table: formValue.value.tableName,
    });
    dataSource.value = formValue.value.masterColumns;
    if (loading) {
      show.value = false;
    }
  }

  onMounted(async () => {
    show.value = true;
    if (formValue.value.masterColumns.length === 0) {
      await reloadFields();
    }

    columns.value = [
      {
        title: '位置',
        key: 'id',
        width: 50,
      },
      {
        title(_column) {
          return renderTooltip(
            h(
              NButton,
              {
                ghost: true,
                strong: true,
                size: 'small',
                text: true,
                iconPlacement: 'right',
              },
              { default: () => '字段', icon: renderIcon(HelpCircleOutline) }
            ),
            'Go类型和属性定义取决于你在/hack/config.yaml中的配置参数'
          );
        },
        key: 'field',
        align: 'center',
        width: 800,
        children: [
          {
            title: '字段列名',
            key: 'name',
            width: 150,
          },
          {
            title: '物理类型',
            key: 'sqlType',
            width: 150,
          },
          {
            title: 'Go属性',
            key: 'goName',
            width: 130,
          },
          {
            title: 'Go类型',
            key: 'goType',
            width: 100,
          },
          {
            title: 'Ts属性',
            key: 'tsName',
            width: 130,
          },
          {
            title: 'Ts类型',
            key: 'tsType',
            width: 100,
          },
          {
            title: '字段描述',
            key: 'dc',
            width: 150,
            render(row) {
              return h(NInput, {
                value: row.dc,
                onUpdateValue: function (e) {
                  row.dc = e;
                },
              });
            },
          },
        ],
      },
      {
        width: 800,
        title(_column) {
          return renderTooltip(
            h(
              NButton,
              {
                ghost: true,
                strong: true,
                size: 'small',
                text: true,
                iconPlacement: 'right',
              },
              { default: () => '新增/编辑表单', icon: renderIcon(HelpCircleOutline) }
            ),
            '勾选编辑以后会在新增、编辑表单中显示该字段;当同时勾选列表查询时，会优先使用配置的表单组件'
          );
        },
        key: 'edit',
        align: 'center',
        children: [
          {
            align: 'center',
            title: '编辑',
            key: 'isEdit',
            width: 50,
            render(row) {
              return h(NCheckbox, {
                defaultChecked: row.isEdit,
                disabled: row.name === 'id',
                onUpdateChecked: function (e) {
                  row.isEdit = e;
                },
              });
            },
          },
          {
            title: '必填',
            key: 'required',
            width: 50,
            align: 'center',
            render(row) {
              return h(NCheckbox, {
                defaultChecked: row.required,
                disabled: row.name === 'id',
                onUpdateChecked: function (e) {
                  row.required = e;
                },
              });
            },
          },
          {
            title: '唯一',
            key: 'unique',
            width: 50,
            align: 'center',
            render(row) {
              return h(NCheckbox, {
                defaultChecked: row.unique,
                disabled: row.name === 'id',
                onUpdateChecked: function (e) {
                  row.unique = e;
                },
              });
            },
          },
          {
            title: '表单组件',
            key: 'formMode',
            width: 200,
            render(row) {
              return h(NSelect, {
                value: row.formMode,
                options: getFormModeOptions(row.tsType),
                // render: function (row) {
                //   return props.selectList?.formMode ?? [];
                // },
                // onFocus: function (e) {
                //   console.log('表单组件  onFocus row:', e);
                // },
                onUpdateValue: function (e) {
                  row.formMode = e;
                },
              });
            },
          },
          {
            title: '表单验证',
            key: 'formRole',
            width: 200,
            render(row) {
              return h(NSelect, {
                value: row.formRole,
                disabled: row.name === 'id',
                options: props.selectList?.formRole ?? [],
                onUpdateValue: function (e) {
                  row.formRole = e;
                },
              });
            },
          },
          {
            title: '字典类型',
            key: 'dictType',
            width: 300,
            render(row) {
              return h(NTreeSelect, {
                value: row.dictType,
                disabled: row.name === 'id',
                options: props.selectList?.dictMode ?? [],
                onUpdateValue: function (e) {
                  row.dictType = e;
                },
              });
            },
          },
        ],
      },
      {
        width: 800,
        title: '列表',
        key: 'list',
        align: 'center',
        children: [
          {
            title: '列表',
            key: 'isList',
            width: 50,
            align: 'center',
            render(row) {
              return h(NCheckbox, {
                defaultChecked: row.isList,
                onUpdateChecked: function (e) {
                  row.isList = e;
                },
              });
            },
          },
          {
            title: '导出',
            key: 'isExport',
            width: 50,
            align: 'center',
            render(row) {
              return h(NCheckbox, {
                defaultChecked: row.isExport,
                onUpdateChecked: function (e) {
                  row.isExport = e;
                },
              });
            },
          },
          {
            title: '查询',
            key: 'isQuery',
            width: 50,
            align: 'center',
            render(row) {
              return h(NCheckbox, {
                defaultChecked: row.isQuery,
                onUpdateChecked: function (e) {
                  row.isQuery = e;
                },
              });
            },
          },
          {
            title: '查询条件',
            key: 'queryWhere',
            width: 300,
            render(row) {
              return h(NSelect, {
                value: row.queryWhere,
                disabled: row.name === 'id',
                options: props.selectList?.whereMode ?? [],
                onUpdateValue: function (e) {
                  row.queryWhere = e;
                },
              });
            },
          },
        ],
      },
    ];

    show.value = false;
  });

  function getFormModeOptions(type: string) {
    const options = cloneDeep(props.selectList?.formMode ?? []);
    if (options.length === 0) {
      return [];
    }
    switch (type) {
      case 'number':
        for (let i = 0; i < options.length; i++) {
          const allows = ['InputNumber', 'Radio', 'Select', 'Switch', 'Rate'];
          if (!allows.includes(options[i].value)) {
            options[i].disabled = true;
          }
        }
        break;
      default:
    }
    return options;
  }
</script>

<style lang="less" scoped></style>
