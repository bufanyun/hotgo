<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="变更余额"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-alert :show-icon="false" type="info">
          通过扣除或增加你的余额来为
          <b> {{ params.realName }}</b> 加款或扣款。当扣款方余额不足时，则会操作失败
        </n-alert>
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="管理员ID" path="id">
            <n-input v-model:value="params.id" :disabled="true" />
          </n-form-item>

          <n-form-item label="TA的余额" path="balance">
            <n-input placeholder="请输入" v-model:value="params.balance" :disabled="true" />
          </n-form-item>

          <n-form-item label="操作方式" path="operateMode">
            <n-radio-group v-model:value="params.operateMode" name="operateMode">
              <n-radio-button
                v-for="status in operateModes"
                :key="status.value"
                :value="status.value"
                :label="status.label"
              />
            </n-radio-group>
          </n-form-item>

          <n-form-item label="操作数量" path="num">
            <n-input placeholder="请输入操作数量" v-model:value="params.num" />
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
  import { ref, computed, watch } from 'vue';
  import {
    addRules as rules,
    addState as State,
    addNewState as newState,
    operateModes,
  } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import { GetMemberView, AddMemberBalance } from '@/api/org/user';
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
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth();
  });

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        AddMemberBalance(params.value).then((_res) => {
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

  function loadForm(value) {
    loading.value = true;
    GetMemberView({ id: value.id })
      .then((res) => {
        params.value = res;
        params.value.operateMode = 1;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  watch(
    () => props.formParams,
    (value) => {
      if (isShowModal.value) {
        loadForm(value);
      }
    }
  );
</script>

<style lang="less"></style>
