<template>
  <div ref="portaletDiv">
    <a-card :bordered="false">
      <div ref="commandstats" style="height: 380px" />
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
          text: '折线图',
          textStyle: {
            fontSize: 16,
            fontWeight: '600',
            color: '#333' // 主标题文字颜色
          },
          left: 8,
          top: 8
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            label: {
              backgroundColor: '#6a7985'
            }
          }
        },
        legend: {
          top: 10,
          right: 20,
          textStyle: {
            color: '#666'
          },
          itemGap: 20,
          itemWidth: 10,
          data: ['邮件营销', '联盟广告']
        },
        toolbox: {
          feature: {}
        },
        grid: {
          left: '1%',
          right: '1%',
          bottom: '2%',
          containLabel: true
        },
        xAxis: [
          {
            type: 'category',
            boundaryGap: false,
            data: [
              '1',
              '2',
              '3',
              '4',
              '5',
              '6',
              '7',
              '8',
              '9',
              '10',
              '11',
              '12',
              '13',
              '14',
              '15',
              '16',
              '17',
              '18',
              '19',
              '20'
            ],
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
            type: 'value',
            textStyle: {
              color: '#666',
              fontSize: '12'
            }
          }
        ],
        series: [
          {
            name: '邮件营销',
            type: 'line',
            stack: '总量',
            itemStyle: {
              color: 'rgba(160,101,251,1)'
            },
            symbolSize: 2,
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(160,101,251,0.2)'
                },
                {
                  offset: 1,
                  color: 'rgba(160,101,251,0.01)'
                }
              ])
            },
            data: [
              60,
              132,
              140,
              134,
              90,
              230,
              210,
              90,
              182,
              191,
              234,
              290,
              330,
              310,
              210,
              132,
              140,
              134,
              140,
              230,
              210,
              90,
              182,
              191
            ]
          },
          {
            name: '联盟广告',
            type: 'line',
            stack: '总量',
            itemStyle: {
              color: 'rgba(58,150,253,1)'
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: 'rgba(58,150,253,0.2)'
                },
                {
                  offset: 1,
                  color: 'rgba(58,77,233,0.01)'
                }
              ])
            },
            data: [
              80,
              132,
              160,
              180,
              160,
              230,
              210,
              90,
              182,
              191,
              234,
              290,
              330,
              280,
              240,
              132,
              140,
              134,
              120,
              230,
              210,
              90,
              182,
              191
            ]
          }
        ]
      })
    }
  }
}
</script>

<style></style>
