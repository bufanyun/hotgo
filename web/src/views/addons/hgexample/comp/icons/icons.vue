<template>
  <div>
    <n-card
      :segmented="{ content: true, footer: true }"
      header-style="padding:10px"
      footer-style="padding:10px"
      content-style=""
    >
      <template #header>
        <n-grid y-gap="20" x-gap="10" cols="24" item-responsive responsive="screen">
          <n-grid-item span="24 m:12 l:6">
            <n-input-group>
              <n-button>图标大小</n-button>
              <n-input-number style="width: 100%" v-model:value="compData.size" />
            </n-input-group>
          </n-grid-item>
          <n-grid-item span="24 m:12 l:6">
            <n-color-picker
              :modes="['hex']"
              style="width: 100%"
              v-model:value="compData.color"
              :swatches="['#FFFFFF', '#18A058', '#2080F0', '#F0A020', '#D03050','#000000']"
            >
              <template #label>
                <div style="color: white">图标颜色 {{ compData.color }}</div>
              </template>
            </n-color-picker>
          </n-grid-item>
        </n-grid>
      </template>
      <div class="icons">
        <n-grid y-gap="20" x-gap="0" cols="24" item-responsive responsive="screen">
          <n-grid-item v-for="(item, idx) in icons" :key="idx" span="12 m:4 l:3 xl:2">
            <div class="icons-item">
              <div class="icons-item_content">
                <n-icon class="icon" :color="compData.color" :size="compData.size">
                  <component :is="item" />
                </n-icon>
                <span class="copy" v-copy="item.name" @click="handleCopy(item)">复制</span>
              </div>
            </div>
          </n-grid-item>
        </n-grid>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import * as AntdIcons from '@vicons/antd';
  import * as Ionicons5 from '@vicons/ionicons5';
  import { useMessage } from 'naive-ui';
  import { useDesignSettingStore } from '@/store/modules/designSetting';

  const designStore = useDesignSettingStore();
  const message = useMessage();

  const icons = ref<any>([]);
  const compData = ref({
    color: designStore.appTheme,
    size: 32,
  });

  const color = computed(() => {
    return compData.value.color;
  });

  function handleCopy(icon) {
    message.success(`已复制，${icon.name}`);
  }

  function open(type: string) {
    if (type === 'antd') {
      icons.value = AntdIcons;
    }
    if (type === 'ionicons5') {
      icons.value = Ionicons5;
    }
  }

  defineExpose({
    open,
  });
</script>
<style lang="less" scoped>
  :deep(.n-color-picker-trigger .n-color-picker-trigger__fill) {
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    overflow: hidden;
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
  }

  .icons {
    &-item {
      width: 100%;
      text-align: center;
      padding: 0 10px;
      position: relative;
      overflow: hidden;
      cursor: pointer;
      &_content {
        padding: 20px 10px;
        border: 1px solid rgb(240 240 245);
        box-sizing: border-box;
      }
      .icon {
        transition: top 0.3s;
        position: relative;
        top: 0;
      }
      &:hover {
        .icon {
          top: -10px;
        }
        .copy {
          bottom: 0;
        }
      }
      .copy {
        position: absolute;
        bottom: -30px;
        height: 30px;
        left: 10px;
        right: 10px;
        line-height: 30px;
        background-color: v-bind(color);
        transition: bottom 0.3s;
        color: rgb(240 240 245);
      }
    }
  }
</style>
