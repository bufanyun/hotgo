<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        :title="params?.id > 0 ? '编辑许可证 #' + params?.id : '添加许可证'"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-alert :show-icon="false" type="info" v-if="params?.id > 0">
          注意：如果服务在线，为了确保更新后的许可证信息生效，服务需要重新登录
        </n-alert>
        <n-form
          :model="params"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="100"
          class="py-4"
        >
          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="分组" path="group">
                <n-select v-model:value="params.group" :options="options.group" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="许可名称" path="name">
                <n-input placeholder="请输入许可名称" v-model:value="params.name" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="应用ID" path="appid">
                <n-input placeholder="请输入应用ID" v-model:value="params.appid" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="应用秘钥" path="secretKey">
                <n-input
                  placeholder="请输入应用秘钥"
                  v-model:value="params.secretKey"
                  type="password"
                  show-password-on="click"
                />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="在线限制" path="onlineLimit">
                <n-input-number
                  placeholder="请输入在线数量限制"
                  v-model:value="params.onlineLimit"
                  style="width: 100%"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="授权有效期" path="endAt">
                <DatePicker v-model:formValue="params.endAt" type="datetime" style="width: 100%" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-form-item label="IP白名单" path="allowedIps">
            <n-input
              type="textarea"
              placeholder="*代表所有，支持IP段，多个IP用,隔开。只有允许的IP才能连接到tcp服务"
              v-model:value="params.allowedIps"
            />
          </n-form-item>

          <n-form-item label="授权状态" path="status">
            <n-select v-model:value="params.status" :options="options.sys_normal_disable" />
          </n-form-item>

          <n-form-item label="备注" path="remark">
            <n-input type="textarea" placeholder="备注" v-model:value="params.remark" />
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
  import { Edit, View } from '@/api/serveLicense';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
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
    console.log('value:' + JSON.stringify(value));
    // 新增
    if (value.id < 1) {
      params.value = newState(value);
      loading.value = false;
      return;
    }

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
