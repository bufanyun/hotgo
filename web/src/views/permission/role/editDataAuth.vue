<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="'修改 ' + formValue.name + ' 的数据权限'"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            :model="formValue"
            ref="formRef"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-form-item label="数据范围" path="dataScope">
              <n-select v-model:value="formValue.dataScope" :options="dataScopeOption" />
            </n-form-item>
            <n-form-item label="自定义权限" path="customDept" v-if="formValue.dataScope === 4">
              <n-tree-select
                multiple
                key-field="id"
                label-field="name"
                :options="deptList"
                v-model:value="formValue.customDept"
                :default-expand-all="true"
              />
            </n-form-item>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { DataScopeEdit, DataScopeSelect } from '@/api/system/role';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { newState, State } from '@/views/permission/role/model';
  import { getDeptList } from '@/api/org/dept';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref();
  const formBtnLoading = ref(false);
  const dataScopeOption = ref<any>([]);
  const deptList = ref<any>([]);

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        DataScopeEdit(formValue.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showModal.value = false;
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

  async function loadDataScopeSelect() {
    const res = await DataScopeSelect();
    if (res.list) {
      dataScopeOption.value = res.list;
    }
  }

  async function loadDeptList() {
    const res = await getDeptList({});
    if (res.list) {
      deptList.value = res.list;
    }
  }

  async function openModal(record: Recordable) {
    showModal.value = true;
    loading.value = true;
    formValue.value = newState(record);
    await loadDeptList();
    await loadDataScopeSelect();
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
