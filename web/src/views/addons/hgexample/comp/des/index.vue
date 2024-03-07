<template>
  <n-space vertical>
    <n-card>
      <template #header>基础用法</template>
      <n-descriptions label-placement="top" title="描述">
        <n-descriptions-item>
          <template #label> 早餐 </template>
          苹果
        </n-descriptions-item>
        <n-descriptions-item label="早午餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="午餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="晚餐" :span="2">
          两个<br />
          苹果
        </n-descriptions-item>
        <n-descriptions-item label="夜宵"> 苹果 </n-descriptions-item>
      </n-descriptions>
    </n-card>
    <n-card>
      <template #header>每个描述项都可以设定跨度</template>
      <n-descriptions label-placement="top" bordered :column="6">
        <n-descriptions-item>
          <template #label> 早餐 </template>
          苹果
        </n-descriptions-item>
        <n-descriptions-item label="早午餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="午餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="晚餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="正餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="夜宵"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="苹果"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="苹果" :span="2"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="苹果" :span="3"> 苹果 </n-descriptions-item>
      </n-descriptions>
    </n-card>
    <n-card>
      <template #header>边框</template>
      <template #header-extra> 如果有很多多行的信息，你可以把它设为 bordered </template>
      <n-descriptions bordered>
        <n-descriptions-item>
          <template #label> 早餐 </template>
          苹果
        </n-descriptions-item>
        <n-descriptions-item label="午餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="晚餐"> 苹果 </n-descriptions-item>
        <n-descriptions-item label="为什么长">
          为什么<br />长<br />长<br />长<br />长<br />长
        </n-descriptions-item>
        <n-descriptions-item label="为什么长">
          为什么<br />长<br />长<br />长<br />长<br />长
        </n-descriptions-item>
        <n-descriptions-item label="为什么长">
          为什么<br />长<br />长<br />长<br />长<br />长
        </n-descriptions-item>
      </n-descriptions>
    </n-card>
    <n-card>
      <n-list hoverable clickable>
        <n-list-item>
          <n-thing title="相见恨晚" content-style="margin-top: 10px;">
            <template #description>
              <n-space size="small" style="margin-top: 4px">
                <n-tag :bordered="false" type="info" size="small"> 暑夜 </n-tag>
                <n-tag :bordered="false" type="info" size="small"> 晚春 </n-tag>
              </n-space>
            </template>
            奋勇呀然后休息呀<br />
            完成你伟大的人生
          </n-thing>
        </n-list-item>
        <n-list-item>
          <n-thing title="他在时间门外" content-style="margin-top: 10px;">
            <template #description>
              <n-space size="small" style="margin-top: 4px">
                <n-tag :bordered="false" type="info" size="small"> 环形公路 </n-tag>
                <n-tag :bordered="false" type="info" size="small"> 潜水艇司机 </n-tag>
              </n-space>
            </template>
            最新的打印机<br />
            复制着彩色傀儡<br />
            早上好<br />
            让他带你去被工厂敲击
          </n-thing>
        </n-list-item>
      </n-list>
    </n-card>
  </n-space>
</template>
<script setup></script>
