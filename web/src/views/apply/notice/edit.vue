<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable>
        <template #header>
          {{
            formValue.id > 0
              ? '编辑' + dict.getLabel('noticeTypeOptions', formValue.type) + ' #' + formValue.id
              : '发送' + dict.getLabel('noticeTypeOptions', formValue.type)
          }}
        </template>
        <n-scrollbar style="max-height: 87vh" class="pr-5">
          <n-spin :show="loading" description="请稍候...">
            <n-alert :show-icon="false" type="info">
              消息发送成功后如果接收人在线会立即收到一条消息通知，编辑已发送的消息不会再次通知
            </n-alert>
            <n-form
              :model="formValue"
              :rules="rules"
              ref="formRef"
              :label-placement="settingStore.isMobile ? 'top' : 'left'"
              :label-width="80"
              class="py-4"
            >
              <n-form-item label="消息标题" path="title">
                <n-input placeholder="请输入消息标题" v-model:value="formValue.title" />
              </n-form-item>

              <n-form-item label="接收人" path="receiver" v-if="formValue.type === 3">
                <n-select
                  multiple
                  :options="options"
                  :render-label="renderLabel"
                  :render-tag="renderMultipleSelectTag"
                  v-model:value="formValue.receiver"
                  filterable
                />
              </n-form-item>

              <n-form-item label="消息内容" path="content">
                <template v-if="formValue.type === 1">
                  <n-input
                    type="textarea"
                    :autosize="{ minRows: 3, maxRows: 10 }"
                    placeholder="请输入通知内容"
                    v-model:value="formValue.content"
                  />
                </template>
                <template v-else>
                  <Editor style="height: 320px" v-model:value="formValue.content" />
                </template>
              </n-form-item>

              <n-grid x-gap="24" :cols="2">
                <n-gi>
                  <n-form-item label="标签" path="tag">
                    <n-select
                      clearable
                      placeholder="可以不填"
                      :render-tag="renderTag"
                      v-model:value="formValue.tag"
                      :options="dict.getOptionUnRef('noticeTagOptions')"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="排序" path="sort">
                    <n-input-number style="width: 100%" v-model:value="formValue.sort" clearable />
                  </n-form-item>
                </n-gi>
              </n-grid>

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
                <n-input
                  type="textarea"
                  placeholder="请输入备注，没有可以不填"
                  v-model:value="formValue.remark"
                />
              </n-form-item>
            </n-form>
          </n-spin>
        </n-scrollbar>
        <template #footer>
          <n-space>
            <n-button @click="closeForm"> 取消 </n-button>
            <n-button type="primary" :loading="formBtnLoading" @click="confirmForm">
              立即发送
            </n-button>
          </n-space>
        </template>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useDictStore } from '@/store/modules/dict';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { personOption, renderLabel, renderMultipleSelectTag } from '@/enums/systemMessageEnum';
  import Editor from '@/components/Editor/editor.vue';
  import { statusOptions } from '@/enums/optionsiEnum';
  import { GetMemberOption } from '@/api/org/user';
  import { MaxSort, EditLetter, EditNotice, EditNotify } from '@/api/apply/notice';
  import { renderTag } from '@/utils';
  import { adaModalWidth } from '@/utils/hotgo';
  import { State, newState, rules } from './model';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const dict = useDictStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const options = ref<personOption[]>();
  const dialogWidth = ref(adaModalWidth(840));

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        switch (formValue.value.type) {
          case 1:
            EditNotify(formValue.value).then((_res) => {
              confirmComplete();
            });
            break;
          case 2:
            EditNotice(formValue.value).then((_res) => {
              confirmComplete();
            });
            break;
          case 3:
            EditLetter(formValue.value).then((_res) => {
              confirmComplete();
            });
            break;
          default:
            message.error('公告类型不支持');
        }
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function confirmComplete() {
    message.success('操作成功');
    setTimeout(() => {
      closeForm();
      emit('reloadTable');
    });
  }

  function getMemberOption() {
    GetMemberOption().then((res) => {
      options.value = res;
    });
  }

  // 关闭抽屉
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开抽屉
  function openModal(state: State, type: number) {
    showModal.value = true;
    dialogWidth.value = adaModalWidth(840);
    getMemberOption();

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      formValue.value.type = type;

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
    formValue.value = newState(state);
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
