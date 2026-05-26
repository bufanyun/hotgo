<template>
  <div>
    <n-modal
      v-model:show="isShowModal"
      :show-icon="false"
      preset="dialog"
      :title="isUpdate ? '编辑 #' + params?.id : '添加'"
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
        <n-form-item label="上级地区" path="pid">
          <n-tree-select :options="optionTreeData" :default-value="params.pid" />
        </n-form-item>

        <n-form-item label="地区ID" path="id">
          <n-input-number
            style="width: 100%"
            placeholder="请输入地区ID"
            v-model:value="params.id"
            :disabled="isUpdate"
            path="handleChangeId"
          />
        </n-form-item>

        <n-form-item label="地区名称" path="title">
          <n-input placeholder="请输入地区名称" v-model:value="params.title" />
        </n-form-item>

        <n-form-item label="拼音" path="pinyin">
          <n-input placeholder="请输入拼音" v-model:value="params.pinyin" />
        </n-form-item>

        <n-grid x-gap="24" :cols="2">
          <n-gi>
            <n-form-item label="经度" path="lng">
              <n-input placeholder="经度" v-model:value="params.lng" />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="纬度" path="lat">
              <n-input placeholder="纬度" v-model:value="params.lat" />
            </n-form-item>
          </n-gi>
        </n-grid>

        <n-form-item label="排序" path="sort">
          <n-input-number v-model:value="params.sort" clearable />
        </n-form-item>

        <n-form-item label="状态" path="status">
          <n-radio-group v-model:value="params.status" name="status">
            <n-radio-button
              v-for="status in dict.getOptionUnRef('sys_normal_disable')"
              :key="Number(status.value)"
              :value="Number(status.value)"
              :label="status.label"
            />
          </n-radio-group>
        </n-form-item>
      </n-form>

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
  import { ref, computed, watch } from 'vue';
  import { State, newState } from './model';
  import { Edit, MaxSort, CheckProvincesUniqueId } from '@/api/apply/provinces';
  import { FormItemRule, useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useDictStore } from '@/store/modules/dict';
  const emit = defineEmits(['reloadTable', 'updateShowModal']);

  interface Props {
    showModal: boolean;
    formParams?: State;
    optionTreeData: any;
    isUpdate: boolean;
  }

  const props = withDefaults(defineProps<Props>(), {
    showModal: false,
    formParams: () => {
      return newState(null);
    },
    optionTreeData: [],
    isUpdate: false,
  });

  const isShowModal = computed({
    get: () => {
      return props.showModal;
    },
    set: (value) => {
      emit('updateShowModal', value);
    },
  });

  const params = computed(() => {
    return props.formParams;
  });

  const rules = {
    id: {
      required: true,
      async validator(rule: FormItemRule, value: string, callback: Function) {
        if (!value) {
          callback(new Error('请填写地区ID'));
        } else if (!/^\d*$/.test(value)) {
          callback(new Error('地区ID应该为整数'));
        } else if (!(await isUniqueId(value))) {
          callback(new Error('地区ID已存在'));
        } else {
          callback();
        }
      },
      trigger: ['input', 'blur'],
    },
    title: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入地区名称',
    },
  };

  const dict = useDictStore();
  const message = useMessage();
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth();
  });

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

  function closeForm() {
    isShowModal.value = false;
  }

  watch(
    () => params.value,
    (value) => {
      params.value.oldId = Number(value.id);
      if (value.id === 0 || value.id === null) {
        MaxSort().then((res) => {
          params.value.sort = res.sort;
        });
      }
    }
  );

  async function isUniqueId(newId: any) {
    const res = await CheckProvincesUniqueId({ oldId: params.value.oldId, newId: newId });
    return res.unique;
  }
</script>

<style lang="less"></style>
