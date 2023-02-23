<template>
  <div>
    <n-card
      :content-style="{ padding: '10px' }"
      :header-style="{ padding: '10px' }"
      :segmented="true"
    >
      <template #header>
        <n-skeleton text style="width: 50%" v-if="loading" />
        <template v-else>
          <div class="text-sm"> 实时网卡流量（全部网卡）</div>
        </template>
      </template>

      <div class="chart-item-container">
        <n-skeleton text v-if="loading" :repeat="10" />
        <template v-else>
          <n-grid responsive="screen" cols="1 s:2 m:4 l:4 xl:4 2xl:4" x-gap="5" y-gap="5">
            <n-grid-item class="item-wrapper">
              <n-card
                :bordered="false"
                :content-style="{ padding: '10px' }"
                :header-style="{ padding: '5px' }"
                :segmented="true"
              >
                <div class="text-number">{{ last?.up }} KB</div>
                <div class="title-text">
                  <n-badge :value="1" dot color="rgb(58, 104, 255)" />
                  上行
                </div>
              </n-card>
            </n-grid-item>
            <n-grid-item class="item-wrapper">
              <n-card
                :bordered="false"
                :content-style="{ padding: '10px' }"
                :header-style="{ padding: '5px' }"
                :segmented="true"
              >
                <div class="text-number"> {{ last?.down ?? 0 }} KB</div>
                <div class="title-text">
                  <n-badge :value="1" dot color="rgb(241, 136, 136)" />
                  下行
                </div>
              </n-card>
            </n-grid-item>
            <n-grid-item class="item-wrapper">
              <n-card
                :bordered="false"
                :content-style="{ padding: '10px' }"
                :header-style="{ padding: '5px' }"
                :segmented="true"
              >
                <div class="text-number">{{ last?.bytesSent ?? 0 }}</div>
                <div class="title-text"> 总发送</div>
              </n-card>
            </n-grid-item>
            <n-grid-item class="item-wrapper">
              <n-card
                :bordered="false"
                :content-style="{ padding: '10px' }"
                :header-style="{ padding: '5px' }"
                :segmented="true"
              >
                <div class="text-number">{{ last?.bytesRecv ?? 0 }}</div>
                <div class="title-text"> 总接收</div>
              </n-card>
            </n-grid-item>
          </n-grid>
          <div ref="fullYearSalesChart" class="chart-item"></div>
        </template>
      </div>
    </n-card>
  </div>
</template>
<script lang="ts">
  import useEcharts from '@/hooks/useEcharts';
  import { defineComponent, nextTick, onBeforeUnmount, ref, watch } from 'vue';
  import { dispose, graphic } from 'echarts';
  import { object } from 'vue-types';

  export default defineComponent({
    name: 'FullYearSalesChart',
    props: {
      dataModel: {
        type: Array || object,
        default: () => {
          return {};
        },
      },
      loading: {
        type: Boolean,
        default: () => {
          return false;
        },
      },
    },
    setup(props) {
      const last = ref<any>({
        bytesSent: '0B',
        bytesRecv: '0B',
        down: '0',
        up: '0',
      });
      const s = ref([]);
      const x = ref([]);
      const sName = ref('上行宽带');
      const xName = ref('下行宽带');
      const months = ref([]);
      const option = ref({
        title: {
          subtext: '单位：KB',
        },
        color: ['rgb(58,104,255)', 'rgb(241,136,136)'],
        grid: {
          top: '10%',
          left: '2%',
          right: '2%',
          bottom: '5%',
          containLabel: true,
        },
        // legend: {
        //   data: [sName.value, xName.value],
        // },
        tooltip: {
          trigger: 'axis',
        },
        xAxis: {
          type: 'category',
          data: months.value,
          boundaryGap: false,
        },
        yAxis: {
          type: 'value',
        },
        series: [
          {
            type: 'line',
            showSymbol: false,
            name: sName.value,
            stack: '总量',
            data: s.value,
            smooth: true,
            lineStyle: {
              color: 'rgba(24, 160, 88, 0.5)',
            },
            label: {
              show: true,
              formatter(val: any) {
                return val.data + 'KB';
              },
            },
            areaStyle: {
              opacity: 0.8,
              color: new graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(85, 193, 250, 0.1)',
                },
                {
                  offset: 1,
                  color: 'rgba(156, 21, 214, 0.2)',
                },
              ]),
            },
          },
          {
            type: 'line',
            showSymbol: false,
            name: xName.value,
            stack: '总量',
            data: x.value,
            smooth: true,
            lineStyle: {
              color: 'rgba(24, 160, 88, 0.5)',
            },
            label: {
              show: true,
              formatter(val: any) {
                return val.data + 'KB';
              },
            },
            areaStyle: {
              opacity: 0.8,
              color: new graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(132, 248, 187, 0.1)',
                },
                {
                  offset: 1,
                  color: 'rgba(51, 209, 125, 0.2)',
                },
              ]),
            },
          },
        ],
      });

      const fullYearSalesChart = ref<HTMLDivElement | null>(null);
      watch(props, (_newVal, _oldVal) => {
        last.value = _newVal.dataModel[_newVal.dataModel.length - 1];

        if (months.value.length < 10) {
          for (let i = 0; i < _newVal.dataModel?.length; i++) {
            s.value.push(_newVal.dataModel[i].up);
            x.value.push(_newVal.dataModel[i].down);
            months.value.push(_newVal.dataModel[i].time);
          }
        } else {
          s.value.shift();
          s.value.push(last.value.up);

          x.value.shift();
          x.value.push(last.value.down);

          months.value.shift();
          months.value.push(last.value.time);
        }

        setTimeout(() => {
          nextTick(() =>
            useEcharts(fullYearSalesChart.value as HTMLDivElement).setOption(option.value)
          );
        }, 10);
      });
      const updateChart = () => {
        useEcharts(fullYearSalesChart.value as HTMLDivElement).resize();
      };
      onBeforeUnmount(() => {
        if (fullYearSalesChart.value !== null) {
          dispose(fullYearSalesChart.value as HTMLDivElement);
        }
      });
      return {
        fullYearSalesChart,
        updateChart,
        last,
      };
    },
  });
</script>

<style lang="less" scoped>
  .chart-item-container {
    width: 100%;

    .chart-item {
      height: 345px;
    }
  }

  .light-green {
    height: 108px;
    background-color: rgba(0, 128, 0, 0.12);
  }

  .green {
    height: 108px;
    background-color: rgba(0, 128, 0, 0.24);
  }
</style>
