<template>
  <div ref="portaletDiv">
    <a-card :bordered="false">
      <div ref="commandstats" style="height: 344px" />
    </a-card>
  </div>
</template>

<script>
import * as echarts from 'echarts'
export default {
  name: 'Cache',
  data () {
    return {
      commandstats: null
    }
  },
  filters: {},
  created () {},
  mounted () {
    this.getFirstChart()
    this.$emit('setHeight', this.$refs.portaletDiv.offsetHeight)
  },
  computed: {},
  watch: {},
  methods: {
    getFirstChart () {
      this.commandstats = echarts.init(this.$refs.commandstats, 'macarons')
      this.commandstats.setOption({
        title: {
          text: '仪表盘',
          textStyle: {
            fontSize: 16,
            fontWeight: '600',
            color: '#333' // 主标题文字颜色
          },
          left: 8,
          top: 8
        },
        series: [
          {
            type: 'gauge',
            startAngle: 200,
            endAngle: -20,
            min: 0,
            max: 1,
            splitNumber: 8,
            itemStyle: {
              color: '#3794fc'
            },
            axisLine: {
              lineStyle: {
                width: 0,
                color: [[1, '#dee5ef']]
              }
            },
            center: ['50%', '140px'],
            pointer: {
              length: '12%',
              width: 0,
              offsetCenter: [0, '-60%'],
              itemStyle: {
                color: 'auto'
              }
            },
            axisTick: {
              distance: 0,
              length: 18,
              lineStyle: {
                color: '#dee5ef',
                width: 2
              }
            },
            progress: {
              show: true,
              width: 18,
              color: 'red'
            },
            splitLine: {
              show: false
            },
            axisLabel: {
              color: '#464646',
              fontSize: 20,
              distance: 60,
              formatter: function (value) {
                if (value === 0.875) {
                  return ''
                } else if (value === 0.625) {
                  return ''
                } else if (value === 0.375) {
                  return ''
                } else if (value === 0.125) {
                  return ''
                }
              }
            },
            title: {},
            detail: {
              fontSize: 32,
              fontWeight: '400',
              offsetCenter: [0, '0'],
              valueAnimation: true,
              formatter: function (value) {
                return Math.round(value * 100)
              },
              color: '#3a96fd'
            },
            data: [
              {
                value: 0.7
              }
            ]
          }
        ]
      })
    }
  }
}
</script>

<style></style>
