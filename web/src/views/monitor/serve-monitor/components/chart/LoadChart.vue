<template>
  <div ref="orderChartWrapper" style="height: 100%"></div>
</template>

<script lang="ts">
  import useEcharts from '@/hooks/useEcharts';
  import { defineComponent, nextTick, onBeforeUnmount, onMounted, ref, watch } from 'vue';
  import { dispose, graphic } from 'echarts';

  export default defineComponent({
    name: 'LoadChart',
    props: {
      dataModel: {
        type: Array,
        default: () => {
          // eslint-disable-next-line vue/require-valid-default-prop
          return {};
        },
      },
    },
    setup(props) {
      const data = ref<any>([]);
      const option = ref({
        tooltip: {
          trigger: 'item',
          axisPointer: {
            type: 'cross',
            label: {
              backgroundColor: '#6a7985',
            },
          },
        },
        grid: {
          x: '-5%',
          y: 0,
          x2: '-5%',
          y2: 0,
        },
        xAxis: {
          type: 'category',
          splitLine: { show: false },
        },
        yAxis: [
          {
            type: 'value',
            splitLine: { show: false },
          },
        ],
        series: [
          {
            type: 'line',
            smooth: true,
            lineStyle: {
              width: 0,
            },
            showSymbol: false,
            areaStyle: {
              opacity: 0.8,
              color: new graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(128, 255, 165)',
                },
                {
                  offset: 1,
                  color: 'rgba(1, 191, 236)',
                },
              ]),
            },
            data: [],
          },
        ],
      });

      const loading = ref(true);
      const orderChartWrapper = ref<HTMLDivElement | null>(null);
      const init = () => {
        for (let i = 0; i < props.dataModel?.length; i++) {
          const v : any = props.dataModel[i]
          data.value.push({
            name: 'CPU分钟负载比率',
            value: [v?.time, v?.ratio],
          });
        }

        // 基于准备好的dom，初始化echarts实例
        setTimeout(() => {
          loading.value = false;
          nextTick(() => {
            useEcharts(orderChartWrapper.value as HTMLDivElement).setOption(option.value);
          });
        }, 100);

        // 绘制图表
        option.value.series.forEach((item) => {
          item.data = data.value;
        });
        useEcharts(orderChartWrapper.value as HTMLDivElement).setOption(option.value);
      };
      const updateChart = () => {
        useEcharts(orderChartWrapper.value as HTMLDivElement).resize();
      };
      onMounted(init);
      onBeforeUnmount(() => {
        dispose(orderChartWrapper.value as HTMLDivElement);
      });
      watch(props, (newVal, _oldVal) => {
        let last : any = newVal.dataModel[newVal.dataModel.length - 1];
        data.value.shift();
        data.value.push({
          name: 'CPU分钟负载比率',
          value: [last?.time, last?.ratio],
        });
        useEcharts(orderChartWrapper.value as HTMLDivElement).setOption(option.value);
      });
      return {
        loading,
        orderChartWrapper,
        updateChart,
      };
    },
  });
</script>
