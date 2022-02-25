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
          text: '玫瑰图',
          textStyle: {
            fontSize: 16,
            fontWeight: '600',
            color: '#333' // 主标题文字颜色
          },
          left: 8,
          top: 8
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          left: 'center',
          top: 'bottom',
          data: ['类型1', '类型2', '类型3', '类型4', '类型5', '类型6', '类型7', '类型8']
        },
        series: [
          {
            name: '面积模式',
            type: 'pie',
            radius: [60, 140],
            center: ['50%', '60%'],
            roseType: 'area',
            itemStyle: {
              borderRadius: 4,
              borderColor: '#fff',
              borderWidth: 4
            },
            data: [
              { value: 30, name: '类型 1' },
              { value: 28, name: '类型 2' },
              { value: 26, name: '类型 3' },
              { value: 24, name: '类型 4' },
              { value: 22, name: '类型 5' },
              { value: 20, name: '类型 6' },
              { value: 18, name: '类型 7' },
              { value: 16, name: '类型 8' }
            ]
          }
        ]
      })
    }
  }
}
</script>

<style></style>
