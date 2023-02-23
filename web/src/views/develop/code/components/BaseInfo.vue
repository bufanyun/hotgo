<template>
  <div>
    <n-spin :show="bodyShow" description="请稍候...">
      <n-card
        :bordered="true"
        title="基本设置"
        class="proCard mt-2"
        size="small"
        :segmented="{ content: true }"
      >
        <n-form ref="formRef" :model="formValue">
          <n-row :gutter="24">
            <n-col :span="6" style="min-width: 200px">
              <n-form-item label="生成类型" path="title">
                <n-select
                  placeholder="请选择"
                  :options="selectList.genType"
                  v-model:value="formValue.genType"
                />
              </n-form-item>
            </n-col>

            <n-col :span="6" style="min-width: 200px">
              <n-form-item label="实体命名" path="varName">
                <n-input placeholder="请输入" v-model:value="formValue.varName" />
              </n-form-item>
            </n-col>

            <n-col :span="6" style="min-width: 200px">
              <n-form-item
                label="数据库"
                path="dbName"
                v-show="formValue.genType >= 10 && formValue.genType < 20"
              >
                <n-select
                  placeholder="请选择"
                  :options="selectList.db"
                  v-model:value="formValue.dbName"
                  @update:value="handleDbUpdateValue"
                />
              </n-form-item>
            </n-col>

            <n-col :span="6" style="min-width: 200px">
              <n-form-item
                label="数据库表"
                path="tableName"
                v-show="formValue.genType >= 10 && formValue.genType < 20"
              >
                <n-select
                  filterable
                  tag
                  :loading="tablesLoading"
                  placeholder="请选择"
                  :options="tablesOption"
                  v-model:value="formValue.tableName"
                  @update:value="handleTableUpdateValue"
                  :disabled="formValue.dbName === ''"
                />
              </n-form-item>
            </n-col>

            <n-col :span="18">
              <n-form-item
                label="表格头部按钮组"
                path="tableName"
                v-show="formValue.genType >= 10 && formValue.genType < 20"
              >
                <n-checkbox-group v-model:value="formValue.options.headOps">
                  <n-space item-style="display: flex;">
                    <n-checkbox value="add" label="新增表单按钮" />
                    <n-checkbox value="batchDel" label="批量删除按钮" />
                    <n-checkbox value="export" label="导出按钮" />
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </n-col>

            <n-col :span="24">
              <n-form-item
                label="表格列操作"
                path="columnOps"
                v-show="formValue.genType >= 10 && formValue.genType < 20"
              >
                <n-checkbox-group v-model:value="formValue.options.columnOps">
                  <n-space item-style="display: flex;">
                    <n-checkbox value="edit" label="编辑" />
                    <n-checkbox value="status" label="状态修改" />
                    <n-popover trigger="hover">
                      <template #trigger>
                        <n-icon size="15" class="tips-help-icon" color="#2d8cf0">
                          <QuestionCircleOutlined />
                        </n-icon>
                      </template>
                      <span>主表中存在`status`字段时才会生效</span>
                    </n-popover>
                    <n-checkbox value="del" label="删除" />
                    <n-checkbox value="view" label="详情页" />
                    <n-checkbox value="check" label="开启勾选列" />
                    <n-checkbox value="switch" label="操作开关" />
                    <n-popover trigger="hover">
                      <template #trigger>
                        <n-icon size="15" class="tips-help-icon" color="#2d8cf0">
                          <QuestionCircleOutlined />
                        </n-icon>
                      </template>
                      <span>主表中存在`switch`字段时才会生效</span>
                    </n-popover>
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </n-col>

            <n-col :span="24">
              <n-form-item
                label="自动化操作"
                path="autoOps"
                v-show="formValue.genType >= 10 && formValue.genType < 20"
              >
                <n-checkbox-group v-model:value="formValue.options.autoOps">
                  <n-space item-style="display: flex;">
                    <n-checkbox value="genMenuPermissions" label="生成菜单权限" />
                    <n-checkbox value="runDao" label="生成前运行 [gf gen dao]" />
                    <n-popover trigger="hover">
                      <template #trigger>
                        <n-icon size="15" class="tips-help-icon" color="#2d8cf0">
                          <QuestionCircleOutlined />
                        </n-icon>
                      </template>
                      <span>如果你选择的表已经生成过dao相关代码，可以忽略</span>
                    </n-popover>
                    <n-checkbox value="runService" label="生成后运行 [gf gen service]" />
                    <n-popover trigger="hover">
                      <template #trigger>
                        <n-icon size="15" class="tips-help-icon" color="#2d8cf0">
                          <QuestionCircleOutlined />
                        </n-icon>
                      </template>
                      <span>如果是插件模块，勾选后也会自动在对应插件下运行service相关代码生成</span>
                    </n-popover>
                    <n-checkbox value="forcedCover" label="强制覆盖" />
                    <n-popover trigger="hover">
                      <template #trigger>
                        <n-icon size="15" class="tips-help-icon" color="#2d8cf0">
                          <QuestionCircleOutlined />
                        </n-icon>
                      </template>
                      <span>只会强制覆盖需要生成的文件，但不包含SQL文件</span>
                    </n-popover>
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </n-col>

            <n-col
              :span="6"
              style="min-width: 200px"
              v-show="formValue.options?.autoOps?.includes('genMenuPermissions')"
            >
              <n-form-item label="上级菜单" path="pid">
                <n-tree-select
                  :options="optionMenuTree"
                  :value="formValue.options.menu.pid"
                  @update:value="handleUpdateMenuPid"
                />
              </n-form-item>
            </n-col>

            <n-col
              :span="6"
              style="min-width: 200px"
              v-show="formValue.options?.autoOps?.includes('genMenuPermissions')"
            >
              <n-form-item label="菜单名称" path="tableComment">
                <n-input placeholder="请输入" v-model:value="formValue.tableComment" />
              </n-form-item>
            </n-col>

            <n-col
              :span="6"
              style="min-width: 200px"
              v-show="formValue.options?.autoOps?.includes('genMenuPermissions')"
            >
              <n-form-item label="菜单图标" path="menuIcon">
                <IconSelector style="width: 100%" v-model:value="formValue.options.menu.icon" />
              </n-form-item>
            </n-col>

            <n-col
              :span="6"
              style="min-width: 200px"
              v-show="formValue.options?.autoOps?.includes('genMenuPermissions')"
            >
              <n-form-item label="菜单排序" path="menuIcon">
                <n-input-number
                  style="width: 100%"
                  placeholder="请输入"
                  v-model:value="formValue.options.menu.sort"
                  clearable
                />
              </n-form-item>
            </n-col>
          </n-row>
        </n-form>
      </n-card>

      <n-card
        :bordered="true"
        title="关联表设置"
        class="proCard mt-2"
        size="small"
        :segmented="{ content: true }"
        v-show="formValue.genType >= 10 && formValue.genType < 20"
      >
        <template #header-extra>
          <n-space>
            <n-button
              type="warning"
              @click="addJoin"
              :disabled="formValue.options?.join?.length >= 3"
              >新增关联表</n-button
            >
          </n-space>
        </template>

        <n-form ref="formRef" :model="formValue">
          <n-alert :show-icon="false">关联表数量建议在三个以下</n-alert>

          <n-row :gutter="6" v-for="(join, index) in formValue.options.join" :key="index">
            <n-col :span="6" style="min-width: 200px">
              <n-form-item label="关联表" path="join.linkTable">
                <n-select
                  filterable
                  tag
                  :loading="tablesLoading"
                  placeholder="请选择"
                  :options="linkTablesOption"
                  v-model:value="join.linkTable"
                  @update:value="handleLinkTableUpdateValue(join)"
                  :disabled="formValue.dbName === ''"
                />
              </n-form-item>
            </n-col>

            <n-col :span="3" style="min-width: 100px">
              <n-form-item
                label="别名"
                path="join.alias"
                v-show="formValue.genType >= 10 && formValue.genType < 20"
              >
                <n-input
                  placeholder="请输入"
                  v-model:value="join.alias"
                  @update:value="updateJoinAlias"
                />

                <template #feedback> {{ joinAliasFeedback }}</template>
              </n-form-item>
            </n-col>

            <n-col :span="3" style="min-width: 100px">
              <n-form-item label="关联方式" path="join.linkMode">
                <n-select
                  placeholder="请选择"
                  :options="selectList.linkMode"
                  v-model:value="join.linkMode"
                />
              </n-form-item>
            </n-col>
            <n-col :span="5" style="min-width: 180px">
              <n-form-item label="关联字段" path="join.field">
                <n-select
                  filterable
                  tag
                  :loading="linkColumnsLoading"
                  placeholder="请选择"
                  :options="linkColumnsOption[join.uuid]"
                  v-model:value="join.field"
                />
              </n-form-item>
            </n-col>

            <n-col :span="5" style="min-width: 180px">
              <n-form-item label="主表关联字段" path="join.masterField">
                <n-select
                  filterable
                  tag
                  :loading="columnsLoading"
                  placeholder="请选择"
                  :options="columnsOption"
                  v-model:value="join.masterField"
                />
              </n-form-item>
            </n-col>

            <n-col :span="2" style="min-width: 50px">
              <n-space>
                <n-form-item label="操作" path="title">
                  <n-button @click="delJoin(join, index)" size="small" strong secondary type="error"
                    >移除</n-button
                  >
                </n-form-item>
              </n-space>
            </n-col>
          </n-row>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref, computed, watch } from 'vue';
  import { FormInst } from 'naive-ui';
  import { newState, selectListObj } from './model';
  import { TableSelect, ColumnSelect } from '@/api/develop/code';
  import { getRandomString } from '@/utils/charset';
  import IconSelector from '@/components/IconSelector/index.vue';
  import { QuestionCircleOutlined } from '@vicons/antd';
  import { getMenuList } from '@/api/system/menu';
  import { cloneDeep } from 'lodash-es';
  import { isLetterBegin } from '@/utils/is';

  const timer = ref();
  const formRef = ref<FormInst | null>(null);
  const bodyShow = ref(true);
  const tablesLoading = ref(false);
  const columnsLoading = ref(false);
  const linkColumnsLoading = ref(false);
  const tablesOption = ref<any>([]); // 数据库表选项
  const columnsOption = ref<any>([]); // 主表字段选项
  const linkTablesOption = ref<any>([]); // 关联表选项
  const linkColumnsOption = ref<any>([]); // 关联表字段选项

  const optionMenuTree = ref([
    {
      id: 0,
      key: 0,
      label: '根目录',
      pid: 0,
      title: '根目录',
      type: 1,
    },
  ]);

  const emit = defineEmits(['update:value']);

  interface Props {
    value?: any;
    selectList: any;
  }

  const props = withDefaults(defineProps<Props>(), {
    value: newState(null),
    selectList: selectListObj,
  });

  watch(props, async (newVal, oldVal) => {
    if (newVal.value.dbName != oldVal.value.dbName) {
      await instLoad();
    }
  });

  const formValue = computed({
    get() {
      return props.value;
    },
    set(value) {
      emit('update:value', value);
    },
  });

  onMounted(() => {
    timer.value = setInterval(async () => {
      if (props.value.id > 0) {
        clearTimeout(timer.value);
        await instLoad();
        // 切换tab时会导致选项被清空，这里重新进行加载
        await loadLinkColumnsOption();
        await loadMenuTreeOption();
        bodyShow.value = false;
      }
    }, 30);
  });

  const loadMenuTreeOption = async () => {
    const options = await getMenuList();
    optionMenuTree.value = [
      {
        id: 0,
        key: 0,
        label: '根目录',
        pid: 0,
        title: '根目录',
        type: 1,
      },
    ];
    optionMenuTree.value = optionMenuTree.value.concat(options.list);
  };

  const loadSelect = async () => {
    columnsOption.value = await loadColumnSelect(formValue.value.tableName);
  };

  async function instLoad() {
    columnsLoading.value = true;
    tablesLoading.value = true;
    await loadSelect();
    await loadTableSelect(formValue.value.dbName);
    tablesLoading.value = false;
    columnsLoading.value = false;
  }

  async function loadLinkColumnsOption() {
    if (formValue.value.options.join === undefined) {
      return;
    }
    for (let i = 0; i < formValue.value.options.join.length; i++) {
      linkColumnsLoading.value = true;
      linkColumnsOption.value[formValue.value.options.join[i].uuid] = await loadColumnSelect(
        formValue.value.options.join[i].linkTable
      );
      linkColumnsLoading.value = false;
    }
  }

  // 处理选项更新
  async function handleDbUpdateValue(value, _option) {
    tablesLoading.value = true;
    await loadTableSelect(value);
    tablesLoading.value = false;
  }

  async function loadTableSelect(value) {
    const options = await TableSelect({ name: value });
    tablesOption.value = cloneDeep(options);
    linkTablesOption.value = cloneDeep(options);
  }

  async function loadColumnSelect(value) {
    return await ColumnSelect({ name: formValue.value.dbName, table: value });
  }

  function handleTableUpdateValue(value, option) {
    formValue.value.varName = option?.defVarName as string;
    formValue.value.daoName = option?.daoName as string;
    formValue.value.tableComment = option?.defTableComment as string;
  }

  function addJoin() {
    if (formValue.value.options.join === undefined) {
      formValue.value.options.join = [];
    }
    let uuid = getRandomString(16, true);
    formValue.value.options.join.push({
      uuid: uuid,
      linkTable: '',
      alias: '',
      linkMode: 1,
      field: '',
      masterField: '',
      daoName: '',
      columns: [],
    });
    linkColumnsOption.value[uuid] = [];
  }

  function delJoin(join, index) {
    formValue.value.options.join.splice(index, 1);
    delete linkColumnsOption.value[join.uuid];
    let i = linkTablesOption.value.findIndex((res) => res.value === join.linkTable);
    if (i > -1) {
      linkTablesOption.value[i].disabled = false;
    }
  }

  async function handleLinkTableUpdateValue(join) {
    let i = linkTablesOption.value.findIndex((res) => res.value === join.linkTable);
    if (i > -1) {
      join.alias = linkTablesOption.value[i].defAlias;
      join.daoName = linkTablesOption.value[i].daoName;
      linkTablesOption.value[i].disabled = true;
    }

    linkColumnsLoading.value = true;
    linkColumnsOption.value[join.uuid] = await loadColumnSelect(join.linkTable);
    // 清空更新前的字段
    join.field = '';
    linkColumnsLoading.value = false;
  }

  const joinAliasFeedback = ref('');
  function updateJoinAlias(value: string) {
    if (value.length < 3) {
      joinAliasFeedback.value = '别名不能小于3位';
      return;
    }

    if (!isLetterBegin(value)) {
      joinAliasFeedback.value = '别名必须以字母开头';
      return;
    }
    joinAliasFeedback.value = '';
  }

  function handleUpdateMenuPid(value: string | number | Array<string | number> | null) {
    formValue.value.options.menu.pid = value;
  }
</script>

<style lang="less" scoped>
  ::v-deep(.default_text_value) {
    color: var(--n-tab-text-color-active);
  }
  ::v-deep(.tips-help-icon) {
    margin-left: -16px;
    margin-top: 5px;
    display: block;
  }
</style>
