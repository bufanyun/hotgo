<template>
  <div ref="portaletDiv">
    <a-card :bordered="false">
      <div ref="commandstats" style="height: 340px" />
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
    window.addEventListener('resize', () => {
      this.commandstats.resize()
    })
    this.$emit('setHeight', this.$refs.portaletDiv.offsetHeight)
  },
  computed: {},
  watch: {},
  methods: {
    getFirstChart () {
      this.commandstats = echarts.init(this.$refs.commandstats, 'macarons')
      this.commandstats.setOption({
        title: {
          text: '折线面积图',
          textStyle: {
            fontSize: 16,
            fontWeight: '600',
            color: '#333' // 主标题文字颜色
          },
          left: 8,
          top: 8
        },
        grid: {
          left: '2%',
          right: '2%',
          bottom: '1%',
          containLabel: true
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
           axisLabel: {
              // 坐标轴文本标签，详见axis.axisLabel
              show: true,
              rotate: 0,
              margin: 8,
              textStyle: {
                color: '#666',
                fontSize: '12'
              }
            },
           axisLine: {
              lineStyle: {
                color: '#dfe6ff',
                width: 1
              }
            }
        },
        yAxis: {
          type: 'value',
          boundaryGap: [0, 0]
        },
        visualMap: {
          type: 'piecewise',
          show: false,
          dimension: 0,
          seriesIndex: 0,
          pieces: [
            {
              gt: 1,
              lt: 3,
              color: 'rgba(55, 148, 252, 0.2)'
            },
            {
              gt: 5,
              lt: 7,
              color: 'rgba(55, 148, 252, 0.2)'
            }
          ]
        },
        series: [
          {
            type: 'line',
            smooth: 0.6,
            symbol: 'none',
            lineStyle: {
              color: '#3794fc',
              width: 2
            },
            markLine: {
              symbol: ['none', 'none'],
              label: { show: false },
              data: [{ xAxis: 1 }, { xAxis: 3 }, { xAxis: 5 }, { xAxis: 7 }],
              lineStyle: {
                color: '#c3cbe5',
                width: 1
              }
            },
            areaStyle: {},
            data: [
              ['10.10', 200],
              ['10.11', 560],
              ['10.12', 750],
              ['10.13', 580],
              ['10.14', 250],
              ['10.15', 300],
              ['10.16', 450],
              ['10.17', 300],
              ['10.18', 100]
            ]
          }
        ]
      })
    }
  }
}
</script>

<style></style>
