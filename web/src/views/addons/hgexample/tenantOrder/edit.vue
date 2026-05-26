<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑购买订单 #' + formValue.id : '添加购买订单'"
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
              <n-gi span="1" v-if="userStore.isCompanyDept">
                <n-form-item label="租户ID" path="tenantId">
                  <n-input placeholder="请输入租户ID" v-model:value="formValue.tenantId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="userStore.isCompanyDept || userStore.isTenantDept">
                <n-form-item label="商户ID" path="merchantId">
                  <n-input placeholder="请输入商户ID" v-model:value="formValue.merchantId" />
                </n-form-item>
              </n-gi>
              <n-gi
                span="1"
                v-if="userStore.isCompanyDept || userStore.isTenantDept || userStore.isMerchantDept"
              >
                <n-form-item label="用户ID" path="userId">
                  <n-input placeholder="请输入用户ID" v-model:value="formValue.userId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="购买产品" path="productName">
                  <n-input placeholder="请输入购买产品" v-model:value="formValue.productName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="关联订单号" path="orderSn">
                  <n-input placeholder="请输入关联订单号" v-model:value="formValue.orderSn" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="充值金额" path="money">
                  <n-input-group>
                    <n-input-number
                      :min="1"
                      :show-button="false"
                      style="width: 100%"
                      placeholder="请输入充值金额"
                      v-model:value="formValue.money"
                    />
                    <n-input-group-label>元</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="备注" path="remark">
                  <n-input placeholder="请输入备注" v-model:value="formValue.remark" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="支付状态" path="status">
                  <n-select
                    v-model:value="formValue.status"
                    :options="dict.getOptionUnRef('payStatus')"
                  />
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
  import { Edit, View } from '@/api/addons/hgexample/tenantOrder';
  import { State, newState, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useUserStore } from '@/store/modules/user';
  import { useDictStore } from '@/store/modules/dict';

  const emit = defineEmits(['reloadTable']);
  const dict = useDictStore();
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const userStore = useUserStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  function openModal(state: State) {
    showModal.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

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
