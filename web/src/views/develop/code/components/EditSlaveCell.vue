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
      />
    </n-card>
  </n-spin>
</template>

<script lang="ts" setup>
  import { Component, computed, h, onMounted, ref } from 'vue';
  import { BasicTable } from '@/components/Table';
  import { genInfoObj, selectListObj } from '@/views/develop/code/components/model';
  import { ColumnList } from '@/api/develop/code';
  import { NButton, NCheckbox, NIcon, NInput, NSelect, NTooltip } from 'naive-ui';
  import { HelpCircleOutline } from '@vicons/ionicons5';

  const renderTooltip = (trigger, content) => {
    return h(NTooltip, null, {
      trigger: () => trigger,
      default: () => content,
    });
  };
  function renderIcon(icon: Component) {
    return () => h(NIcon, null, { default: () => h(icon) });
  }

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

  const columns = ref<any>([]);

  const formValue = computed({
    get() {
      return props.value;
    },
    set(value) {
      emit('update:value', value);
    },
  });

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

  const show = ref(false);
  const dataSource = ref([]);
  onMounted(async () => {
    show.value = true;
    setTimeout(async () => {
      const index = getIndex();
      if (formValue.value.options.join[index].columns.length === 0) {
        formValue.value.options.join[index].columns = await ColumnList({
          name: formValue.value.dbName,
          table: formValue.value.options.join[index].linkTable,
          isLink: 1,
          alias: formValue.value.options.join[index].alias,
        });
      }

      dataSource.value = formValue.value.options.join[index].columns;

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
              width: 260,
            },
            {
              title: 'Go类型',
              key: 'goType',
              width: 100,
            },
            {
              title: 'Ts属性',
              key: 'tsName',
              width: 260,
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
                    // await saveProductCustom(row.id, 'frontShow', e);
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
    }, 50);
  });

  const actionRef = ref();
</script>

<style lang="less" scoped></style>
