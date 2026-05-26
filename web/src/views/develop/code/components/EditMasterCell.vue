<template>
  <n-card :bordered="false" class="proCard">
    <n-spin :show="show" description="加载中...">
      <BasicTable
        size="small"
        striped
        resizable
        canResize
        virtual-scroll
        :single-line="false"
        :showTopRight="false"
        :pagination="false"
        :columns="columns"
        :dataSource="dataSource"
        :row-key="(row) => row.id"
        :resizeHeightOffset="-20000"
        :scroll-x="columnCollapse ? 1400 : 2400"
        :scroll-y="720"
        :scrollbar-props="{ trigger: 'none' }"
      >
        <template #tableTitle>
          <n-space>
            <n-popconfirm @positive-click="reloadColumns">
              <template #trigger>
                <n-button type="primary" class="min-left-space">
                  <template #icon>
                    <n-icon>
                      <Reload />
                    </n-icon>
                  </template>
                  重置字段
                </n-button>
              </template>
              重置后将从数据库重新加载表，不保留当前字段配置，确定要重置吗？
            </n-popconfirm>

            <n-popconfirm @positive-click="syncColumns">
              <template #trigger>
                <n-button type="primary" class="min-left-space">
                  <template #icon>
                    <n-icon>
                      <Sync />
                    </n-icon>
                  </template>
                  同步字段
                </n-button>
              </template>
              同步是从数据库重新加载表，保留当前有效的字段配置，确定要同步吗？
            </n-popconfirm>

            <n-button type="default" class="min-left-space" @click="handleMove">
              <template #icon>
                <n-icon>
                  <MoveOutline />
                </n-icon>
              </template>
              移动字段
            </n-button>
          </n-space>
        </template>
      </BasicTable>
    </n-spin>
  </n-card>
  <Move ref="moveRef" v-model:columns="dataSource" />
</template>

<script lang="ts" setup>
  import { computed, h, onMounted, ref } from 'vue';
  import { BasicTable } from '@/components/Table';
  import {
    formatColumns,
    formGridColsOptions,
    formGridSpanOptions,
    genInfoObj,
    selectListObj,
  } from '@/views/develop/code/components/model';
  import { ColumnList } from '@/api/develop/code';
  import { NInputNumber, NSpace, NButton, NCheckbox, NInput, NSelect, NCascader } from 'naive-ui';
  import { HelpCircleOutline, Reload, Sync, MoveOutline } from '@vicons/ionicons5';
  import { cloneDeep } from 'lodash-es';
  import { renderIcon, renderTooltip } from '@/utils';
  import Move from './Move.vue';

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

  const show = ref(false);
  const dataSource = ref([]);
  const moveRef = ref();
  const columnCollapse = ref(true);
  const columnsCollapseData = computed(() => {
    return columnCollapse.value
      ? [
          {
            title: '字段列名',
            key: 'name',
            width: 100,
          },
          {
            title: '字段描述',
            key: 'dc',
            width: 100,
            render(row) {
              return h(NInput, {
                value: row.dc,
                onUpdateValue: function (e) {
                  row.dc = e;
                },
              });
            },
          },
        ]
      : [
          {
            title: '字段列名',
            key: 'name',
            width: 100,
          },
          {
            title: '物理类型',
            key: 'sqlType',
            width: 80,
          },
          {
            title: 'Go属性',
            key: 'goName',
            width: 100,
          },
          {
            title: 'Go类型',
            key: 'goType',
            width: 80,
          },
          {
            title: 'Ts属性',
            key: 'tsName',
            width: 100,
          },
          {
            title: 'Ts类型',
            key: 'tsType',
            width: 80,
          },
          {
            title: '字段描述',
            key: 'dc',
            width: 100,
            render(row) {
              return h(NInput, {
                value: row.dc,
                onUpdateValue: function (e) {
                  row.dc = e;
                },
              });
            },
          },
        ];
  });

  const columns = computed(() => {
    return [
      {
        title: '',
        key: 'id',
        width: 30,
        render(row, index) {
          return index + 1;
        },
      },
      {
        title(_column) {
          return h('div', null, [
            renderTooltip(
              h(
                NButton,
                {
                  strong: true,
                  size: 'small',
                  text: true,
                  iconPlacement: 'right',
                },
                { default: () => '字段', icon: renderIcon(HelpCircleOutline) }
              ),
              'Go类型和属性定义取决于你在/hack/config.yaml中的配置参数'
            ),
            h(
              NButton,
              {
                strong: true,
                size: 'small',
                text: true,
                type: 'primary',
                style: { 'margin-left': '20px' },
                onClick: () => (columnCollapse.value = !columnCollapse.value),
              },
              { default: () => (columnCollapse.value ? '展开 >>' : '折叠 <<') }
            ),
          ]);
        },
        key: 'field',
        align: 'center',
        width: 800,
        children: columnsCollapseData.value,
      },
      {
        width: 800,
        title(_column) {
          return renderTooltip(
            h(
              NButton,
              {
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
            width: 30,
            render(row) {
              const disabled = isEditDisabled(row);
              const checkbox = h(NCheckbox, {
                defaultChecked: row.isEdit,
                disabled: disabled,
                onUpdateChecked: function (e) {
                  row.isEdit = e;
                },
              });
              if (!disabled) {
                return checkbox;
              }
              return renderTooltip(checkbox, '该字段属性由系统维护，无需单独配置！');
            },
          },
          {
            title: '必填',
            key: 'required',
            width: 30,
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
            width: 30,
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
            width: 100,
            render(row) {
              return h(NSelect, {
                consistentMenuWidth: false,
                value: row.formMode,
                options: getFormModeOptions(row.tsType),
                onUpdateValue: function (e) {
                  row.formMode = e;
                },
              });
            },
          },
          {
            title: '绑定字典',
            key: 'dictType',
            width: 100,
            render(row) {
              if (row.dictType == 0) {
                row.dictType = null;
              }
              return h(NCascader, {
                placeholder: ' ',
                filterable: true,
                clearable: true,
                showPath: false,
                checkStrategy: 'child',
                disabled: row.name === 'id',
                value: row.dictType,
                options: props.selectList?.dictMode ?? [],
                onUpdateValue: function (e) {
                  row.dictType = e;
                },
              });
            },
          },
          {
            title: '验证规则',
            key: 'formRole',
            width: 100,
            render(row) {
              return h(NSelect, {
                consistentMenuWidth: false,
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
            title(_column) {
              return h(NSpace, { inline: true }, [
                renderTooltip(
                  h(
                    NButton,
                    {
                      strong: true,
                      size: 'small',
                      text: true,
                      iconPlacement: 'right',
                    },
                    { default: () => '栅格', icon: renderIcon(HelpCircleOutline) }
                  ),
                  '表单每行摆放组件的个数。响应式栅格，小屏幕自动转为每行摆放一个组件。参考文档：https://www.naiveui.com/zh-CN/os-theme/components/grid#responsive-item.vue'
                ),
                h(NSelect, {
                  style: { width: '100px' },
                  size: 'small',
                  consistentMenuWidth: false,
                  value: formValue.value.options.presetStep.formGridCols,
                  options: formGridColsOptions,
                  onUpdateValue: function (e) {
                    formValue.value.options.presetStep.formGridCols = e;
                  },
                }),
              ]);
            },
            key: 'formGridSpan',
            width: 120,
            render(row) {
              return h(NSelect, {
                consistentMenuWidth: false,
                disabled: row.name === 'id',
                value: row.formGridSpan,
                options: getFormGridSpanOptions(formValue.value.options.presetStep.formGridCols),
                onUpdateValue: function (e) {
                  row.formGridSpan = e;
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
            width: 30,
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
            width: 30,
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
            width: 30,
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
            width: 90,
            render(row) {
              return h(NSelect, {
                consistentMenuWidth: false,
                value: row.queryWhere,
                disabled: row.name === 'id',
                options: props.selectList?.whereMode ?? [],
                onUpdateValue: function (e) {
                  row.queryWhere = e;
                },
              });
            },
          },
          {
            title: '排列方式',
            key: 'align',
            width: 80,
            render(row) {
              return h(NSelect, {
                consistentMenuWidth: false,
                value: row.align,
                options: props.selectList?.tableAlign ?? [],
                onUpdateValue: function (e) {
                  row.align = e;
                },
              });
            },
          },
          {
            title(_column) {
              return renderTooltip(
                h(
                  NButton,
                  {
                    strong: true,
                    size: 'small',
                    text: true,
                    iconPlacement: 'right',
                  },
                  { default: () => '列宽', icon: renderIcon(HelpCircleOutline) }
                ),
                '选填。设定固定值时表格生成自动计算scroll-x，未设定默认每列按100计算'
              );
            },
            key: 'width',
            width: 50,
            render(row) {
              return h(NInputNumber, {
                value: row.width,
                placeholder: ' ',
                min: -1,
                max: 2000,
                showButton: false,
                onUpdateValue: function (e) {
                  row.width = e;
                },
              });
            },
          },
        ],
      },
    ];
  });

  // 同步字段
  function syncColumns() {
    show.value = true;
    dataSource.value = [];

    const params = {
      name: formValue.value.dbName,
      table: formValue.value.tableName,
    };

    ColumnList(params)
      .then((res) => {
        const columns = formatColumns(res);
        for (let i = 0; i < columns.length; i++) {
          // 相同字段名称和类型，保留原字段属性
          const index = formValue.value.masterColumns.findIndex(
            (item) => item.name == columns[i].name && item.dataType == columns[i].dataType
          );
          if (index !== -1) {
            columns[i] = formValue.value.masterColumns[index];
          }
        }

        formValue.value.masterColumns = columns;
        dataSource.value = formValue.value.masterColumns;
      })
      .finally(() => {
        show.value = false;
      });
  }

  // 重载字段属性
  function reloadColumns() {
    show.value = true;
    dataSource.value = [];

    const params = {
      name: formValue.value.dbName,
      table: formValue.value.tableName,
    };

    ColumnList(params)
      .then((res) => {
        formValue.value.masterColumns = formatColumns(res);
        dataSource.value = formValue.value.masterColumns;
      })
      .finally(() => {
        show.value = false;
      });
  }

  function getFormModeOptions(type: string) {
    const options = cloneDeep(props.selectList?.formMode ?? []);
    if (options.length === 0) {
      return [];
    }
    switch (type) {
      case 'number':
        for (let i = 0; i < options.length; i++) {
          const allows = [
            'Input',
            'InputNumber',
            'Radio',
            'Select',
            'Switch',
            'Rate',
            'TreeSelect',
            'Cascader',
          ];
          if (!allows.includes(options[i].value)) {
            options[i].disabled = true;
          }
        }
        break;
      default:
    }
    options.sort((a, b) => (a.disabled === b.disabled ? 0 : a.disabled ? 1 : -1));
    return options;
  }

  function getFormGridSpanOptions(cols: number) {
    if (cols < 1) {
      cols = 1;
    }
    if (cols > 4) {
      cols = 4;
    }
    for (let i = 0; i < formValue.value.masterColumns.length; i++) {
      if (!formValue.value.masterColumns[i].formGridSpan) {
        formValue.value.masterColumns[i].formGridSpan = 1;
      }
      if (formValue.value.masterColumns[i].formGridSpan > cols) {
        formValue.value.masterColumns[i].formGridSpan = cols;
      }
    }
    return formGridSpanOptions.slice(0, Math.min(cols, formGridSpanOptions.length));
  }

  // 禁止编辑的字段，由系统维护
  function isEditDisabled(row) {
    const disabledNames = [
      'id',
      'created_by',
      'updated_by',
      'deleted_by',
      'created_at',
      'updated_at',
      'deleted_at',
    ];
    if (disabledNames.includes(row.name)) {
      return true;
    }

    if (formValue.value.genType == 11) {
      const disabledTreeNames = ['pid', 'level', 'tree'];
      if (disabledTreeNames.includes(row.name)) {
        return true;
      }
    }
    return false;
  }

  function handleMove() {
    moveRef.value.openModal();
  }

  onMounted(() => {
    if (formValue.value.masterColumns.length === 0) {
      reloadColumns();
    } else {
      show.value = true;
      setTimeout(function () {
        dataSource.value = formValue.value.masterColumns;
        show.value = false;
      }, 100);
    }
  });
</script>

<style lang="less" scoped>
  .tree-tips {
    margin-left: 12px;
    color: #18a058;
    font-weight: 600;
  }
</style>
