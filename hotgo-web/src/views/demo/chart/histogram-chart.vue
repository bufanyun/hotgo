<template>
  <div ref="portaletDiv">
    <a-card :bordered="false">
      <div ref="commandstats" style="height: 343px" />
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
          text: '柱状图',
          textStyle: {
            fontSize: 16,
            fontWeight: '600',
            color: '#333' // 主标题文字颜色
          },
          left: 8,
          top: 8
        },
        legend: {
          top: 10,
          right: 20,
          textStyle: {
            color: '#666'
          },
          itemGap: 20,
          itemWidth: 10,
          data: ['邮件营销']
        },
        tooltip: {
          trigger: 'axis'
        },
        grid: {
          left: '2%',
          right: '2%',
          bottom: '2%',
          top: '20%',
          containLabel: true
        },
        xAxis: [
          {
            type: 'category',
            data: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12'],
            axisTick: {
              alignWithLabel: true
            },
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
          }
        ],
        yAxis: [
          {
            type: 'value'
          }
        ],
        series: [
          {
            name: '完成情况',
            type: 'bar',
            barWidth: '40%',
            itemStyle: {
              barBorderRadius: [4, 4, 0, 0],
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                { offset: 0, color: '#2881e6' },
                { offset: 0.5, color: '#308bf2' },
                { offset: 1, color: '#3693fb' }
              ])
            },
            data: [60, 52, 200, 334, 390, 330, 220, 200, 334, 390, 330, 220]
          }
        ]
      })
    }
  }
}
</script>

<style></style>
