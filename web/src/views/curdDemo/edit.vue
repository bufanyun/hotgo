<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
       :title="params?.id > 0 ? '编辑 #' + params?.id : '添加'"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="分类ID" path="categoryId">
            <n-input-number placeholder="请输入分类ID" v-model:value="params.categoryId" />
          </n-form-item>

          <n-form-item label="标题" path="title">
          <n-input placeholder="请输入标题" v-model:value="params.title" />
          </n-form-item>

          <n-form-item label="描述" path="description">
            <n-input type="textarea" placeholder="描述" v-model:value="params.description" />
          </n-form-item>

          <n-form-item label="内容" path="content">
            <Editor style="height: 450px" v-model:value="params.content" />
          </n-form-item>

          <n-form-item label="单图" path="image">
            <UploadImage :maxNumber="1" v-model:value="params.image" />
          </n-form-item>

          <n-form-item label="附件" path="attachfile">
            <UploadFile :maxNumber="1" v-model:value="params.attachfile" />
          </n-form-item>

          <n-form-item label="所在城市" path="cityId">
            <CitySelector v-model:value="params.cityId" />
          </n-form-item>

          <n-form-item label="显示开关" path="switch">
            <n-switch :unchecked-value="2" :checked-value="1" v-model:value="params.switch"
        />
          </n-form-item>

          <n-form-item label="排序" path="sort">
            <n-input-number placeholder="请输入排序" v-model:value="params.sort" />
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-select v-model:value="params.status" :options="options.sys_normal_disable" />
          </n-form-item>


        </n-form>
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
  import { onMounted, ref, computed, watch } from 'vue';
  import { Edit, MaxSort, View } from '@/api/curdDemo';
  import Editor from '@/components/Editor/editor.vue';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import UploadFile from '@/components/Upload/uploadFile.vue';
  import CitySelector from '@/components/CitySelector/citySelector.vue';
  import { rules, options, State, newState } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable', 'updateShowModal']);

  interface Props {
    showModal: boolean;
    formParams?: State;
  }

  const props = withDefaults(defineProps<Props>(), {
    showModal: false,
    formParams: () => {
      return newState(null);
    },
  });

  const isShowModal = computed({
    get: () => {
      return props.showModal;
    },
    set: (value) => {
      emit('updateShowModal', value);
    },
  });

  const loading = ref(false);
  const params = ref<State>(props.formParams);
  const message = useMessage();
  const formRef = ref<any>({});
  const dialogWidth = ref('75%');
  const formBtnLoading = ref(false);

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(params.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            isShowModal.value = false;
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  onMounted(async () => {
    adaModalWidth(dialogWidth);
  });

  function closeForm() {
    isShowModal.value = false;
  }

  function loadForm(value) {
    loading.value = true;

    // 新增
    if (value.id < 1) {
      params.value = newState(value);
      MaxSort()
        .then((res) => {
          params.value.sort = res.sort;
        })
        .finally(() => {
          loading.value = false;
        });
      return;
    }

    // 编辑
    View({ id: value.id })
      .then((res) => {
        params.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
  );
</script>

<style lang="less"></style>