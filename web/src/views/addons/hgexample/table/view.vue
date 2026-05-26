<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="基础详情"> 基础详情，有时也用于显示只读信息。 </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" :column="4">
        <n-descriptions-item>
          <template #label>分类ID</template>
          {{ formValue.categoryId }}
        </n-descriptions-item>

        <n-descriptions-item label="标签">
          <template v-for="(item, key) in formValue?.flag" :key="key">
            <n-tag
              :type="dict.getType('sys_notice_type', item)"
              size="small"
              class="min-left-space"
              >{{ dict.getLabel('sys_notice_type', item) }}</n-tag
            >
          </template>
        </n-descriptions-item>
        <n-descriptions-item label="标题">{{ formValue.title }}</n-descriptions-item>
        <n-descriptions-item label="描述">{{ formValue.description }}</n-descriptions-item>
        <n-descriptions-item label="推荐星"
          ><n-rate readonly :default-value="formValue.star"
        /></n-descriptions-item>
        <n-descriptions-item label="价格">{{ formValue.price }}</n-descriptions-item>
        <n-descriptions-item label="浏览次数">{{ formValue.views }}</n-descriptions-item>
        <n-descriptions-item label="活动时间">{{ formValue.activityAt }}</n-descriptions-item>
        <n-descriptions-item label="开关">
          <n-switch v-model:value="formValue.switch" :unchecked-value="2" :checked-value="1"
        /></n-descriptions-item>
        <n-descriptions-item label="创建人ID">{{ formValue.createdBy }} </n-descriptions-item>
        <n-descriptions-item label="创建时间">{{ formValue.createdAt }} </n-descriptions-item>
      </n-descriptions>
    </n-card>

    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="top" title="内容" class="py-2" :column="1">
        <n-descriptions-item><span v-html="formValue.content"></span></n-descriptions-item>
      </n-descriptions>
    </n-card>

    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="top" title="单图" class="py-2" :column="1">
        <n-descriptions-item>
          <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.image"
        /></n-descriptions-item>
      </n-descriptions>

      <n-descriptions label-placement="top" title="多图" class="py-2" :column="1">
        <n-descriptions-item>
          <n-image-group>
            <n-space>
              <span v-for="(item, key) in formValue?.images" :key="key">
                <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="item" />
              </span>
            </n-space>
          </n-image-group>
        </n-descriptions-item>
      </n-descriptions>

      <n-descriptions label-placement="top" title="附件" class="py-2" :column="1">
        <n-descriptions-item>
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
      </n-descriptions>

      <n-descriptions label-placement="top" title="多附件" class="py-2" :column="1">
        <n-descriptions-item>
          <div class="upload-card">
            <n-space style="gap: 0px 0px">
              <div
                class="upload-card-item"
                style="height: 100px; width: 100px"
                v-for="(item, key) in formValue.attachfiles"
                :key="key"
              >
                <div class="upload-card-item-info">
                  <div class="img-box">
                    <n-avatar :style="fileAvatarCSS" @click="download(item)">{{
                      getFileExt(item)
                    }}</n-avatar>
                  </div>
                </div>
              </div>
            </n-space>
          </div>
        </n-descriptions-item>
      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/addons/hgexample/table';
  import { newState } from './model';
  import { getFileExt } from '@/utils/urlUtils';
  import { useDictStore } from '@/store/modules/dict';

  const dict = useDictStore();
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
