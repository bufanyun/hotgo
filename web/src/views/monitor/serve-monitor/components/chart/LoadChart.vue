<template>
  <div ref="orderChartWrapper" style="height: 100%"></div>
</template>

<script lang="ts">
  import useEcharts from '@/hooks/useEcharts';
  import { defineComponent, onBeforeUnmount, onMounted, ref, watch } from 'vue';
  import { dispose, graphic } from 'echarts';

  type DataItem = {
    time: string;
    avg: number;
    ratio: number;
  };

  export default defineComponent({
    name: 'LoadChart',
    props: {
      dataModel: {
        type: Array,
        default: () => {
          return [];
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

      const orderChartWrapper = ref();

      function getWrapper(): HTMLElement {
        return orderChartWrapper.value as HTMLElement;
      }

      const init = () => {
        // 绘制图表
        option.value.series.forEach((item) => {
          item.data = data.value;
        });

        handleResize();

        useEcharts(getWrapper()).setOption(option.value);
      };

      // 调整图表大小
      const handleResize = () => {
        useEcharts(getWrapper()).resize();
      };

      onMounted(init);

      onBeforeUnmount(() => {
        dispose(getWrapper());
      });

      watch(props, (newVal) => {
        const newValues = newVal.dataModel.map((item: DataItem) => ({
          name: 'CPU分钟负载比率',
          value: [item.time, item.ratio],
        }));

        data.value.push(...newValues);

        // 移除超过10个的最开头元素
        if (data.value.length > 10) {
          const removeCount = data.value.length - 10;
          data.value.splice(0, removeCount);
        }

        useEcharts(getWrapper()).setOption(option.value);
      });

      return {
        orderChartWrapper,
      };
    },
  });
</script>
