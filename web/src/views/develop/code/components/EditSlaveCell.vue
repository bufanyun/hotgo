<template>
  <n-spin :show="show" description="加载中...">
    <n-card :bordered="false" class="proCard">
      <BasicTable
        size="small"
        striped
        resizable
        canResize
        :single-line="false"
        :showTopRight="false"
        :pagination="false"
        :columns="columns"
        :dataSource="dataSource"
        :row-key="(row) => row.id"
        :resizeHeightOffset="-20000"
        :scroll-x="columnCollapse ? 880 : 1880"
        :scroll-y="720"
        :scrollbar-props="{ trigger: 'none' }"
      >
        <template #tableTitle>
          <n-space>
            <n-popconfirm @positive-click="reloadColumns(getIndex())">
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
              重置后将从数据库重新加载该表的默认配置，确定要重置吗？
            </n-popconfirm>

            <n-popconfirm @positive-click="syncColumns(getIndex())">
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
    </n-card>
    <Move ref="moveRef" v-model:columns="dataSource" />
  </n-spin>
</template>

<script lang="ts" setup>
  import { computed, h, onMounted, ref } from 'vue';
  import { BasicTable } from '@/components/Table';
  import { formatColumns, genInfoObj, selectListObj } from '@/views/develop/code/components/model';
  import { ColumnList } from '@/api/develop/code';
  import {
    NButton,
    NCheckbox,
    NIcon,
    NInput,
    NInputNumber,
    NSelect,
    NSpace,
    NTooltip,
  } from 'naive-ui';
  import { HelpCircleOutline, MoveOutline, Reload, Sync, WarningOutline } from '@vicons/ionicons5';
  import { renderIcon, renderTooltip } from '@/utils';
  import Move from './Move.vue';

  const emit = defineEmits(['update:value']);

  interface Props {
    value?: any;
    selectList: any;
    uuid: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    value: genInfoObj,
    selectList: selectListObj,
    uuid: '',
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
            width: 150,
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
                  // await saveProductCustom(row.id, 'frontShow', e);
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
            width: 150,
          },
          {
            title: 'Go类型',
            key: 'goType',
            width: 100,
          },
          {
            title: 'Ts属性',
            key: 'tsName',
            width: 150,
          },
          {
            title: 'Ts类型',
            key: 'tsType',
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
        ];
  });

  const columns = computed(() => {
    return [
      {
        title: '',
        key: 'id',
        width: 50,
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
            width: 100,
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
            width: 100,
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
                min: 0,
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

  function handleMove() {
    moveRef.value.openModal();
  }
  function getIndex() {
    if (formValue.value.options.join.length === 0) {
      return -1;
    }
    for (let i = 0; i < formValue.value.options.join.length; i++) {
      if (formValue.value.options.join[i].uuid === props.uuid) {
        return i;
      }
    }
    return -1;
  }

  // 同步字段属性
  function syncColumns(index: number) {
    show.value = true;
    dataSource.value = [];

    const join = formValue.value.options.join[index];
    const params = {
      name: formValue.value.dbName,
      table: join.linkTable,
      isLink: 1,
      alias: join.alias,
    };

    ColumnList(params)
      .then((res) => {
        const columns = formatColumns(res);
        for (let i = 0; i < columns.length; i++) {
          // 相同字段名称和类型，保留原字段属性
          const index2 = join.columns.findIndex(
            (item) => item.name == columns[i].name && item.dataType == columns[i].dataType
          );
          if (index2 !== -1) {
            columns[i] = join.columns[index2];
          }
        }
        join.columns = columns;
        dataSource.value = join.columns;
      })
      .finally(() => {
        show.value = false;
      });
  }

  // 重载字段属性
  function reloadColumns(index: number) {
    show.value = true;
    dataSource.value = [];

    const join = formValue.value.options.join[index];
    const params = {
      name: formValue.value.dbName,
      table: join.linkTable,
      isLink: 1,
      alias: join.alias,
    };

    ColumnList(params)
      .then((res) => {
        join.columns = formatColumns(res);
        dataSource.value = join.columns;
      })
      .finally(() => {
        show.value = false;
      });
  }

  onMounted(() => {
    show.value = true;
    setTimeout(() => {
      const index = getIndex();

      // 已存在直接加载
      if (formValue.value.options.join[index].columns.length > 0) {
        dataSource.value = formValue.value.options.join[index].columns;
        show.value = false;
        return;
      }

      reloadColumns(index);
    }, 100);
  });
</script>

<style lang="less" scoped></style>
