<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-drawer v-model:show="showModal" :width="dialogWidth">
        <n-drawer-content>
          <template #header> 生成演示详情 </template>
          <template #footer>
            <n-button @click="showModal = false"> 关闭 </n-button>
          </template>
          <n-descriptions label-placement="left" class="py-2" column="1">
                <n-descriptions-item>
          <template #label>分类ID</template>
          {{ formValue.categoryId }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>标题</template>
          {{ formValue.title }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>描述</template>
          <span v-html="formValue.description"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>内容</template>
          <span v-html="formValue.content"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>单图</template>
          <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.image"
        /></n-descriptions-item>

        <n-descriptions-item>
          <template #label>附件</template>
          <div
            class="upload-card"
            v-show="formValue.attachfile !== ''"
            @click="download(formValue.attachfile)"
          >
            <div class="upload-card-item" style="height: 100px; width: 100px">
              <div class="upload-card-item-info">
                <div class="img-box">
                  <n-avatar :style="fileAvatarCSS">{{ getFileExt(formValue.attachfile) }}</n-avatar>
                </div>
              </div>
            </div>
          </div>
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>所在城市</template>
          {{ formValue.cityId }}
        </n-descriptions-item>

        <n-descriptions-item label="显示开关">
          <n-switch v-model:value="formValue.switch" :unchecked-value="2" :checked-value="1" :disabled="true"
        /></n-descriptions-item>

        <n-descriptions-item>
          <template #label>排序</template>
          {{ formValue.sort }}
        </n-descriptions-item>

        <n-descriptions-item label="状态">
          <n-tag
            :type="getOptionTag(options.sys_normal_disable, formValue?.status)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.sys_normal_disable, formValue?.status) }}</n-tag
          >
        </n-descriptions-item>


          </n-descriptions>
        </n-drawer-content>
      </n-drawer>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/curdDemo';
  import { State, newState, options } from './model';
  import { adaModalWidth, getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';

  const message = useMessage();
  const dialogWidth = ref('75%');
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const fileAvatarCSS = computed(() => {
    return {
      '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
      '--n-font-size': `18px`,
    };
  });

  //下载
  function download(url: string) {
    window.open(url);
  }

  function openModal(state: State) {
    adaModalWidth(dialogWidth, 580);
    showModal.value = true;
    loading.value = true;
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

<style lang="less" scoped></style>