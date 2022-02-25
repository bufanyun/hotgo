<template>
  <div ref="portaletDiv">
    <a-card :bordered="false">
      <div ref="commandstats" style="height: 144px" />
    </a-card>
  </div>
</template>

<script>
import * as echarts from 'echarts'
export default {
  name: 'Cache',
  data () {
    return {
      commandstats: null,
      chartData: [
      ]
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
        backgroundColor: '#148be4',
        title: {
          text: ''
        },
        tooltip: {
          trigger: 'axis'
        },
        label: {
          show: false
        },
        legend: {
          show: false,
          left: 'center',
          data: ['工时分类']
        },
        radar: [
          {
            indicator: [
              { text: '设计', max: 100 },
              { text: '交互', max: 100 },
              { text: '平面', max: 100 },
              { text: '产品', max: 100 },
              { text: '规划', max: 100 }
            ],
            center: ['50%', '55%'],
            radius: 60,
            name: {
              formatter: '{value}',
              textStyle: {
                fontSize: 10,
                color: '#148be4'
              }
            }
          }
        ],
        series: [
          {
            type: 'radar',
            tooltip: {
              trigger: 'item'
            },
            data: [
              {
                value: [60, 73, 85, 40, 50],
                name: '工时分类',
                itemStyle: {
                  normal: {
                    color: '#00b1ff',
                    borderColor: '#00b1ff',
                    borderWidth: 4
                  }
                },
                lineStyle: {
                  normal: {
                    color: '#00b1ff',
                    width: 2,
                    shadowColor: 'rgba(0,177,255,0.6)',
                    shadowBlur: 50,
                    shadowOffsetY: 15
                  }
                },
                areaStyle: {
                  normal: {
                  color: 'rgba(0,177,255,0.6)'
                  }
                }
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
