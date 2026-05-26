<template>
  <div class="box">
    <template v-if="column === 'visitor'">
      <div class="flex flex-col">
        <div class="flex flex-row mb-1">
          <div class="text">
            {{
              state.memberId === 0
                ? state.memberName
                : state.memberName + '(' + state.memberId + ')'
            }}
          </div>
        </div>
        <div class="flex flex-row">
          <div>访问IP：</div>
          <div class="text">{{ state.ip }}</div>
        </div>
      </div>
      <div class="text">
        {{ state.cityLabel !== '' ? state.cityLabel : '局域网' }}
      </div>
    </template>

    <template v-if="column === 'request'">
      <div class="flex flex-col">
        <div class="flex flex-row">
          <div class="text mr-1 mb-1">
            <n-button :type="state.method === 'GET' ? 'tertiary' : 'primary'" size="tiny">
              {{ state.method }}
            </n-button>
          </div>
          <div class="text">{{ state.url }}</div>
        </div>
        <div class="flex flex-row">
          <div class="text">{{ state.tags }}</div>
        </div>
      </div>
      <div>
        {{ state.summary }}
      </div>
    </template>

    <template v-if="column === 'response'">
      <div class="flex flex-col">
        <div class="flex flex-row">
          <div class="text mr-1 mb-1">
            <n-button v-if="state.errorCode === 0" type="tertiary" size="tiny">
              {{ state.errorMsg }}
            </n-button>
            <n-button v-else type="error" size="tiny">
              {{ state.errorCode }} → {{ state.errorMsg }}
            </n-button>
          </div>
        </div>
        <div class="flex flex-row">
          <div class="text">处理耗时：{{ state.takeUpTime }}ms</div>
        </div>
        <div class="flex flex-row">
          <div class="text">响应时间：{{ timestampToTime(state.timestamp) }}</div>
        </div>
      </div>
    </template>
  </div>
</template>

<script lang="ts" setup>
  import { basicProps } from './props';
  import { timestampToTime } from '@/utils/dateUtil';

  const props = defineProps(basicProps);
</script>

<style lang="less" scoped>
  .text {
    display: inline;
  }
</style>
