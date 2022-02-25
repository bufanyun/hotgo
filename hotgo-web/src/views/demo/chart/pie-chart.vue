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
            text: '饼状图',
            textStyle: {
              fontSize: 16,
              fontWeight: '600',
              color: '#333' // 主标题文字颜色
            },
            left: 8,
            top: 8
          },
          tooltip: {
            show: true,
            trigger: 'item'
          },
          color: ['#4782DA', '#FF9800', '#34C38F', '#F44336', '#EEEEEE'],
          legend: {
            bottom: '1%',
            left: 'center',
            orient: 'horizontal'
          },
          series: [{
            name: '访问来源',
            type: 'pie',
            radius: ['60%', '70%'],
            center: ['50%', '50%'],
            avoidLabelOverlap: false,
            label: {
              show: true,
              fontSize: '16',
              position: 'center',
              formatter: ' {c} \r\n \r\n {b} ',
              borderWidth: 4
            },
            itemStyle: {
              borderRadius: 4,
              borderColor: '#fff',
              borderWidth: 4
            },
            emphasis: {
              label: {
                show: true,
                fontSize: '20',
                fontWeight: 'bold'
              }
            },
            labelLine: {
              show: false
            },
            data: [{
                value: 1048,
                name: '用户数量'
              },
              {
                value: 735,
                name: '直接访问'
              },
              {
                value: 580,
                name: '邮件营销'
              },
              {
                value: 484,
                name: '联盟广告'
              }
            ]
          }]
        })
      }
    }
  }
</script>

<style></style>
