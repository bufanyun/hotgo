<template>
  <n-card
    :segmented="{ content: true, footer: true }"
    header-style="padding:10px"
    footer-style="padding:10px"
  >
    <template #header> 时间线 </template>
    <n-space>
      <n-card title="竖向">
        <n-timeline>
          <n-timeline-item content="啊" />
          <n-timeline-item type="success" title="成功" content="哪里成功" time="2018-04-03 20:46" />
          <n-timeline-item type="error" content="哪里错误" time="2018-04-03 20:46" />
          <n-timeline-item type="warning" title="警告" content="哪里警告" time="2018-04-03 20:46" />
          <n-timeline-item
            type="info"
            title="信息"
            content="是的"
            time="2018-04-03 20:46"
            line-type="dashed"
          />
          <n-timeline-item content="啊" />
        </n-timeline>
      </n-card>
      <n-card title="横向">
        <div style="overflow: auto">
          <n-timeline horizontal>
            <n-timeline-item content="啊" />
            <n-timeline-item
              type="success"
              title="成功"
              content="哪里成功"
              time="2018-04-03 20:46"
            />
            <n-timeline-item type="error" content="哪里失败" time="2018-04-03 20:46" />
            <n-timeline-item
              type="warning"
              title="警告"
              content="哪里警告"
              time="2018-04-03 20:46"
            />
            <n-timeline-item type="info" title="信息" content="是的" time="2018-04-03 20:46" />
          </n-timeline>
        </div>
      </n-card>
      <n-card title="自定义">
        <n-timeline :icon-size="20">
          <n-timeline-item color="grey" content="啊">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
          </n-timeline-item>
          <n-timeline-item type="success" title="成功" content="哪里成功" time="2018-04-03 20:46">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
          </n-timeline-item>
          <n-timeline-item type="error" content="哪里错误" time="2018-04-03 20:46">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
          </n-timeline-item>
          <n-timeline-item type="warning" title="警告" content="哪里警告" time="2018-04-03 20:46">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
          </n-timeline-item>
          <n-timeline-item type="info" title="信息" content="是的" time="2018-04-03 20:46">
            <template #icon>
              <n-icon>
                <cash-icon />
              </n-icon>
            </template>
          </n-timeline-item>
        </n-timeline>
      </n-card>
    </n-space>
  </n-card>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';
  import { CashOutline as CashIcon } from '@vicons/ionicons5';

  export default defineComponent({
    components: {
      CashIcon,
    },
  });
</script>
