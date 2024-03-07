<template>
  <n-space vertical>
    <n-card
      :segmented="{ content: true, footer: true }"
      header-style="padding:10px"
      footer-style="padding:10px"
    >
      <template #header> 文字水印 </template>
      <template #header-extra> 核心机密 </template>
      <n-watermark
        content="核心机密"
        cross
        selectable
        :font-size="16"
        :line-height="16"
        :width="192"
        :height="128"
        :x-offset="12"
        :y-offset="28"
        :rotate="-15"
      >
        <n-table :single-line="false">
          <thead>
            <tr>
              <th>复盘</th>
              <th>赋能</th>
              <th>协同</th>
              <th>...</th>
              <th>串联</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>拉通</td>
              <td>打通</td>
              <td>树立</td>
              <td>...</td>
              <td>履约</td>
            </tr>
            <tr>
              <td>...</td>
              <td>...</td>
              <td>...</td>
              <td>...</td>
              <td>...</td>
            </tr>
          </tbody>
        </n-table>
      </n-watermark>
    </n-card>
    <n-card
      :segmented="{ content: true, footer: true }"
      header-style="padding:10px"
      footer-style="padding:10px"
    >
      <template #header> 图片水印 </template>
      <template #header-extra>
        <n-image height="30" width="30" :src="compData.img" />
      </template>
      <n-watermark
        :image="compData.img"
        cross
        :rotate="-15"
        :font-size="16"
        :line-height="16"
        :width="192"
        :height="128"
        :x-offset="12"
        :y-offset="30"
        :image-width="34"
        :image-opacity="0.24"
      >
        <n-table :single-line="false">
          <thead>
            <tr>
              <th>复盘</th>
              <th>赋能</th>
              <th>协同</th>
              <th>...</th>
              <th>串联</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>拉通</td>
              <td>打通</td>
              <td>树立</td>
              <td>...</td>
              <td>履约</td>
            </tr>
            <tr>
              <td>...</td>
              <td>...</td>
              <td>...</td>
              <td>...</td>
              <td>...</td>
            </tr>
          </tbody>
        </n-table>
      </n-watermark>
    </n-card>
    <n-card
      :segmented="{ content: true, footer: true }"
      header-style="padding:10px"
      footer-style="padding:10px"
    >
      <template #header> 自定义水印 </template>
      <template #header-extra>
        <n-input-group>
          <n-input v-model:value="compData.text" placeholder="输入自定义水印" />
          <n-button type="primary" ghost>添加</n-button>
        </n-input-group>
      </template>
      <n-watermark
        v-model:content="compData.text"
        cross
        selectable
        :font-size="16"
        :line-height="16"
        :width="192"
        :height="128"
        :x-offset="12"
        :y-offset="28"
        :rotate="-15"
      >
        <n-table :single-line="false">
          <thead>
            <tr>
              <th>复盘</th>
              <th>赋能</th>
              <th>协同</th>
              <th>...</th>
              <th>串联</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>拉通</td>
              <td>打通</td>
              <td>树立</td>
              <td>...</td>
              <td>履约</td>
            </tr>
            <tr>
              <td>...</td>
              <td>...</td>
              <td>...</td>
              <td>...</td>
              <td>...</td>
            </tr>
          </tbody>
        </n-table>
      </n-watermark>
    </n-card>
  </n-space>
</template>
<script lang="ts">
  import { defineComponent, reactive } from 'vue';
  import logo from '@/assets/images/logo.png';
  export default defineComponent({
    setup() {
      const compData = reactive({
        img: logo,
        text: '自定义水印',
      });
      return {
        compData,
      };
    },
  });
</script>
