<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="选项树表">
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
    <n-grid cols="1 s:1 m:1 l:4 xl:4 2xl:4" responsive="screen" :x-gap="12">
      <n-gi span="1">
        <n-card :segmented="{ content: true }" :bordered="false" size="small" class="proCard">
          <template #header>
            <n-space>
              <n-button type="info" icon-placement="left" @click="addTable" v-if="hasPermission(['/optionTreeDemo/edit'])">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <PlusOutlined />
                    </n-icon>
                  </div>
                </template>
                添加
              </n-button>
              <n-button v-if="hasPermission(['/optionTreeDemo/edit'])" type="info" icon-placement="left" @click="handleEdit(selectedState)" :disabled="selectedState.id < 1">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <EditOutlined />
                    </n-icon>
                  </div>
                </template>
                编辑
              </n-button>
              <n-button v-if="hasPermission(['/optionTreeDemo/delete'])" type="error" icon-placement="left" @click="handleDelete(selectedState)" :disabled="selectedState.id < 1">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <DeleteOutlined />
                    </n-icon>
                  </div>
                </template>
                删除
              </n-button>
              <n-button type="info" icon-placement="left" @click="handleAllExpanded">
                {{ expandedKeys.length ? '收起' : '展开' }}
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <AlignLeftOutlined />
                    </n-icon>
                  </div>
                </template>
              </n-button>
            </n-space>
          </template>
          <div class="w-full menu">
            <n-input v-model:value="pattern" placeholder="输入名称搜索">
              <template #suffix>
                <n-icon size="18" class="cursor-pointer">
                  <SearchOutlined />
                </n-icon>
              </template>
            </n-input>
            <div class="py-3 menu-list">
              <template v-if="loading">
                <div class="flex items-center justify-center py-4">
                  <n-spin size="medium" />
                </div>
              </template>
              <n-tree v-else show-line block-line cascade virtual-scroll :pattern="pattern" :data="treeOption" :expandedKeys="expandedKeys" style="height: 75vh" key-field="id" label-field="title" @update:selected-keys="handleSelected" @update:expanded-keys="handleOnExpandedKeys" />
            </div>
          </div>
        </n-card>
      </n-gi>
      <n-gi span="3">
        <n-card :bordered="false" class="proCard">
          <template #header v-if="selectedState.id > 0">
            <n-space>
              <n-icon size="18">
                <FormOutlined />
              </n-icon>
              <span>
                正在编辑 {{ selectedState.title }}
              </span>
            </n-space>
          </template>
          <n-result v-show="selectedState.id < 1" status="info" title="提示" description="请先从列表选择一项后，进行编辑">
            <template #footer>
              <n-button type="info" icon-placement="left" @click="handleAdd(selectedState)"  v-if="hasPermission(['/optionTreeDemo/edit'])">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <PlusOutlined />
                    </n-icon>
                  </div>
                </template>
                添加
              </n-button>
            </template>
          </n-result>
          <BasicForm v-if="selectedState.id > 0" ref="searchFormRef" @register="register" @submit="reloadTable" @reset="reloadTable" @keyup.enter="reloadTable">
            <template #statusSlot="{ model, field }">
              <n-input v-model:value="model[field]" />
            </template>
          </BasicForm>
          <BasicTable v-if="selectedState.id > 0" ref="actionRef" openChecked :columns="columns" :request="loadDataTable" :row-key="(row) => row.id" :actionColumn="actionColumn" :scroll-x="scrollX" :resizeHeightOffset="-10000"  :checked-row-keys="checkedIds" @update:checked-row-keys="handleOnCheckedRow">
            <template #tableTitle>
              <n-button type="primary" @click="handleAdd(selectedState)"  class="min-left-space" v-if="hasPermission(['/optionTreeDemo/edit'])">
                <template #icon>
                  <n-icon>
                    <PlusOutlined />
                  </n-icon>
                </template>
                添加
              </n-button>
              <n-button type="error" @click="handleBatchDelete" class="min-left-space" v-if="hasPermission(['/optionTreeDemo/delete'])">
                <template #icon>
                  <n-icon>
                    <DeleteOutlined />
                  </n-icon>
                </template>
                批量删除
              </n-button>
            </template>
          </BasicTable>
        </n-card>
      </n-gi>
    </n-grid>
    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, computed, onMounted, unref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useDictStore } from '@/store/modules/dict';
  import { List, Delete, TreeOption } from '@/api/optionTreeDemo';
  import { PlusOutlined, EditOutlined, DeleteOutlined, AlignLeftOutlined, FormOutlined, SearchOutlined } from '@vicons/antd';
  import { columns, schemas, loadOptions, loadTreeOption, treeOption, State, newState } from './model';
  import { adaTableScrollX, getTreeKeys } from '@/utils/hotgo';
  import Edit from './edit.vue';

  const dict = useDictStore();
  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();

  const checkedIds = ref([]);

  const expandedKeys = ref([]);
  const pattern = ref('');
  const selectedState = ref<State>(newState(null));
  const loading = ref(false);

  const actionColumn = reactive({
    width: 144,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record: State) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/optionTreeDemo/edit'],
          },

          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/optionTreeDemo/delete'],
          },
        ],

      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  // 加载选项式树表数据
  const loadDataTable = async (res = {}) => {
    if (selectedState.value.id < 1) {
      return;
    }

    // 刷新树选项
    loadTreeOption();

    // 获取选中的下级列表
    const params = {
      ...(searchFormRef.value?.formModel ?? {}),
      ...res,
      pid: selectedState.value.id,
    };
    return await List(params);
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  // 添加数据
  function addTable() {
    editRef.value.openModal(null);
  }

  // 添加树节点下级数据
  function handleAdd(record: Recordable) {
    const state = newState(null);
    state.pid = record.id;
    editRef.value.openModal(state);
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 单个删除
  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 批量删除
  function handleBatchDelete() {
    if (checkedIds.value.length < 1){
      message.error('请至少选择一项要删除的数据');
      return;
    }

    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 选中树节点
  function handleSelected(keys, option) {
    if (keys.length) {
      selectedState.value = newState(option[0]);
      reloadTable();
    } else {
      selectedState.value = newState(null);
    }
  }

  // 展开指定节点
  function handleOnExpandedKeys(keys) {
    expandedKeys.value = keys;
  }

  // 展开全部节点
  function handleAllExpanded() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = getTreeKeys(unref(treeOption), 'id');
    }
  }

  // 首次加载树选项，默认展开全部
  function firstLoadTreeOption() {
    loading.value = true;
    TreeOption().then((res) => {
      treeOption.value = res;
      expandedKeys.value = getTreeKeys(unref(treeOption), 'id');
      loading.value = false;
    });
  }

  onMounted(() => {
    loadOptions();
    firstLoadTreeOption();
  });
</script>

<style lang="less" scoped></style>