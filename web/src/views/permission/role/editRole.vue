<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑角色 #' + formValue.id : '添加角色'"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            :model="formValue"
            :rules="rules"
            ref="formRef"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-form-item label="上级角色" path="pid">
              <n-tree-select
                :options="editRoleOption"
                v-model:value="formValue.pid"
                key-field="id"
                label-field="name"
                clearable
                filterable
                default-expand-all
              />
            </n-form-item>
            <n-form-item label="角色名称" path="name">
              <n-input placeholder="请输入名称" v-model:value="formValue.name" />
            </n-form-item>
            <n-form-item label="角色编码" path="key">
              <n-input placeholder="请输入" v-model:value="formValue.key" />
            </n-form-item>
            <n-form-item label="排序" path="sort">
              <n-input-number v-model:value="formValue.sort" clearable style="width: 100%" />
            </n-form-item>

            <n-form-item label="状态" path="status">
              <n-radio-group v-model:value="formValue.status" name="status">
                <n-radio-button
                  v-for="status in statusOptions"
                  :key="status.value"
                  :value="status.value"
                  :label="status.label"
                />
              </n-radio-group>
            </n-form-item>
            <n-form-item label="备注" path="remark">
              <n-input type="textarea" placeholder="请输入" v-model:value="formValue.remark" />
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
  import { computed, ref } from 'vue';
  import { Edit, getRoleList } from '@/api/system/role';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { statusOptions } from '@/enums/optionsiEnum';
  import { newState, State } from '@/views/permission/role/model';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref();
  const formBtnLoading = ref(false);
  const rawRoleOption = ref([]);

  const editRoleOption = computed(() => {
    return rawRoleOption.value;
  });

  const rules = {
    name: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入名称',
    },
    key: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入角色编码',
    },
  };

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formValue.value).then((_res) => {
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

  function loadDataList() {
    loading.value = true;
    getRoleList({ pageSize: 100, page: 1 }).then((res) => {
      rawRoleOption.value = res.list;
      loading.value = false;
    });
  }

  function openModal(record: Recordable) {
    loadDataList();
    showModal.value = true;
    formValue.value = newState(record);
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
