<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="showModal"
        :mask-closable="false"
        :show-icon="false"
        preset="dialog"
        transform-origin="center"
        :title="formValue.id > 0 ? '编辑 #' + formValue.id : '添加'"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-scrollbar style="max-height: 87vh" class="pr-5">
          <n-form
            :model="formValue"
            :rules="rules"
            ref="formRef"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-form-item label="分类ID" path="categoryId">
            <n-input-number placeholder="请输入分类ID" v-model:value="formValue.categoryId" />
          </n-form-item>

          <n-form-item label="标题" path="title">
          <n-input placeholder="请输入标题" v-model:value="formValue.title" />
          </n-form-item>

          <n-form-item label="描述" path="description">
            <n-input type="textarea" placeholder="描述" v-model:value="formValue.description" />
          </n-form-item>

          <n-form-item label="内容" path="content">
            <Editor style="height: 450px" id="content" v-model:value="formValue.content" />
          </n-form-item>

          <n-form-item label="单图" path="image">
            <UploadImage :maxNumber="1" v-model:value="formValue.image" />
          </n-form-item>

          <n-form-item label="附件" path="attachfile">
            <UploadFile :maxNumber="1" v-model:value="formValue.attachfile" />
          </n-form-item>

          <n-form-item label="所在城市" path="cityId">
            <CitySelector v-model:value="formValue.cityId" />
          </n-form-item>

          <n-form-item label="显示开关" path="switch">
            <n-switch :unchecked-value="2" :checked-value="1" v-model:value="formValue.switch"
        />
          </n-form-item>

          <n-form-item label="排序" path="sort">
            <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" />
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-select v-model:value="formValue.status" :options="options.sys_normal_disable" />
          </n-form-item>


          </n-form>
        </n-scrollbar>
        <template #action>
          <n-space>
            <n-button @click="closeForm">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Edit, MaxSort, View } from '@/api/curdDemo';
  import Editor from '@/components/Editor/editor.vue';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import UploadFile from '@/components/Upload/uploadFile.vue';
  import CitySelector from '@/components/CitySelector/citySelector.vue';
  import { rules, options, State, newState } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dialogWidth = ref('75%');
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);

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

  function openModal(state: State) {
    adaModalWidth(dialogWidth);
    showModal.value = true;
    loading.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
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
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>