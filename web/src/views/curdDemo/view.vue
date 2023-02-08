<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="生成演示详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
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
          <template v-for="(item, key) in formValue?.status" :key="key">
            <n-tag
              :type="getOptionTag(options.sys_normal_disable, item)"
              size="small"
              class="min-left-space"
              >{{ getOptionLabel(options.sys_normal_disable, item) }}</n-tag
            >
          </template>
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/curdDemo';
  import { newState, options } from './model';
  import { getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';

  const message = useMessage();
  const router = useRouter();
  const id = Number(router.currentRoute.value.params.id);
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

  onMounted(async () => {
    if (id < 1) {
      message.error('ID不正确，请检查！');
      return;
    }
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>