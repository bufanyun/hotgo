<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑部门 #' + formValue.id : '添加部门'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="上级部门" path="pid">
                  <n-tree-select
                    :options="treeOption"
                    v-model:value="formValue.pid"
                    key-field="id"
                    label-field="name"
                    clearable
                    filterable
                    default-expand-all
                    show-path
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="部门名称" path="name">
                  <n-input placeholder="请输入部门名称" v-model:value="formValue.name" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="部门编码" path="code">
                  <n-input placeholder="请输入部门编码" v-model:value="formValue.code" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="部门类型" path="type">
                  <n-radio-group v-model:value="formValue.type" name="type">
                    <n-space>
                      <n-radio v-for="item in dict.getOptionUnRef('deptType')" :value="item.value">
                        {{ item.label }}
                      </n-radio>
                    </n-space>
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="负责人" path="leader">
                  <n-input placeholder="请输入负责人" v-model:value="formValue.leader" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系电话" path="phone">
                  <n-input placeholder="请输入联系电话" v-model:value="formValue.phone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="邮箱" path="email">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.email" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="排序" path="sort">
                  <n-input-number
                    placeholder="请输入排序"
                    v-model:value="formValue.sort"
                    clearable
                    style="width: 100%"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="status in dict.getOptionUnRef('sys_normal_disable')"
                      :key="status.value"
                      :value="status.value"
                      :label="status.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm"> 取消 </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 确定 </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { Edit, View, MaxSort } from '@/api/org/dept';
  import { State, newState, treeOption, loadTreeOption, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useDictStore } from '@/store/modules/dict';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(520);
  });

  function openModal(state: State) {
    showModal.value = true;

    // 加载关系树选项
    loadTreeOption();

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

      loading.value = true;
      MaxSort()
        .then((res) => {
          formValue.value.sort = res.sort;
        })
        .finally(() => {
          loading.value = false;
        });
      return;
    }

    // 编辑
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formValue.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            closeForm();
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
