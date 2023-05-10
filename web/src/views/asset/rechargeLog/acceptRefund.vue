<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        title="受理退款申请"
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
          <n-form-item label="业务单号" path="orderSn">
            <n-input v-model:value="params.orderSn" :disabled="true" />
          </n-form-item>

          <n-form-item label="订单金额" path="money">
            <n-input placeholder="请输入标题" v-model:value="params.money" :disabled="true" />
          </n-form-item>

          <n-form-item label="退款原因" path="refundReason">
            <n-input
              type="textarea"
              placeholder="请填写退款原因"
              v-model:value="params.refundReason"
              :disabled="true"
            />
          </n-form-item>

          <n-form-item label="更新状态" path="status">
            <n-select v-model:value="params.status" :options="options.acceptRefundStatus" />
          </n-form-item>

          <n-form-item label="拒绝原因" path="rejectRefundReason" v-if="params.status === 9">
            <n-input
              type="textarea"
              placeholder="请填拒绝退款原因"
              v-model:value="params.rejectRefundReason"
            />
          </n-form-item>

          <n-form-item label="退款备注" path="remark" v-if="params.status === 8">
            <n-input type="textarea" placeholder="请填退款备注" v-model:value="params.remark" />
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
  import { rules, State, newState, options } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import { AcceptRefund, View } from '@/api/order';
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
        AcceptRefund(params.value).then((_res) => {
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
