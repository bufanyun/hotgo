<template>
  <div class="padding-card">
    <a-spin tip="正在加载缓存监控数据，请稍后！" :spinning="loading">
      <a-row :gutter="16">
        <a-col :span="24">
          <a-card title="基本信息" :bordered="false">
            <a-descriptions :column="5">
              <a-descriptions-item label="Redis版本">
                <span v-if="cache.info">{{ cache.info.redis_version }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="运行模式">
                <span v-if="cache.info">{{ cache.info.redis_mode == 'standalone' ? '单机' : '集群' }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="端口">
                <span v-if="cache.info">{{ cache.info.tcp_port }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="客户端数">
                <span v-if="cache.info">{{ cache.info.connected_clients }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="运行时间(天)">
                <span v-if="cache.info">{{ cache.info.uptime_in_days }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="使用内存">
                <span v-if="cache.info">{{ cache.info.used_memory_human }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="使用CPU">
                <span v-if="cache.info">{{ parseFloat(cache.info.used_cpu_user_children).toFixed(2) }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="内存配置">
                <span v-if="cache.info">{{ cache.info.maxmemory_human }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="AOF是否开启">
                <span v-if="cache.info">{{ cache.info.aof_enabled == '0' ? '否' : '是' }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="RDB是否成功">
                <span v-if="cache.info">{{ cache.info.rdb_last_bgsave_status }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="Key数量">
                <span v-if="cache.dbSize">{{ cache.dbSize }}</span>
              </a-descriptions-item>
              <a-descriptions-item label="系统架构">
                <span v-if="cache.info">
                  {{ cache.info.instantaneous_input_kbps }}kps/{{ cache.info.instantaneous_output_kbps }}kps</span>
              </a-descriptions-item>
            </a-descriptions>
          </a-card>
        </a-col>
      </a-row>
      <a-row :gutter="16" style="margin-top: 16px;">
        <a-col :span="12">
          <a-card title="命令统计" :bordered="false">
            <div ref="commandstats" style="height: 380px ;z-index: 2;" />
          </a-card>
        </a-col>
        <a-col :span="12">
          <a-card title="内存信息" :bordered="false">
            <div ref="usedmemory" style="height: 380px" />
          </a-card>
        </a-col>
      </a-row>
    </a-spin>
  </div>
</template>

<script>
  import {
    getCache
  } from '@/api/monitor/cache'
  import * as echarts from 'echarts'

  export default {
    name: 'Cache',
    data () {
      return {
        loading: true,
        // 统计命令信息
        commandstats: null,
        // 使用内存
        usedmemory: null,
        // cache信息
        cache: []
      }
    },
    filters: {},
    created () {
      this.getList('0')
    },
    mounted () {
      window.addEventListener('resize', () => {
        this.commandstats.resize()
        this.usedmemory.resize()
      })
    },
    computed: {},
    watch: {},
    methods: {
      /** 查询服务信息 */
      getList (type) {
        if (type === '0') {
          this.loading = true
        }
        getCache().then(response => {
          const cache = response.data
          this.cache = cache
          if (type === '0') {
            this.loading = false
          }
          this.$nextTick(() => {
           const commandStats = response.data.commandStats
            const names = []
            const values = []
            commandStats.forEach(item => names.push(item.name))
            commandStats.forEach(item => values.push(item.value))
            this.commandstats = echarts.init(this.$refs.commandstats, 'macarons')
            this.commandstats.setOption({
              title: {
                textStyle: {
                  fontSize: 16,
                  fontWeight: '600',
                  color: '#333' // 主标题文字颜色
                },
                left: 8,
                top: 8
              },

              tooltip: {
                trigger: 'axis'
              },
              grid: {
                left: '2%',
                right: '2%',
                bottom: '2%',
                top: '8%',
                containLabel: true
              },
              xAxis: [{
                type: 'category',
                 data: names,
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
                    fontSize: '14'
                  }
                },
                axisLine: {
                  lineStyle: {
                    color: '#dfe6ff',
                    width: 1

                  }
                }
              }],
              yAxis: [{
                type: 'value'
              }],
              series: [{
                name: '使用次数',
                type: 'bar',
                barWidth: '20%',
                itemStyle: {
                  barBorderRadius: [4, 4, 0, 0],
                  color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                      offset: 0,
                      color: '#2881e6'
                    },
                    {
                      offset: 0.5,
                      color: '#308bf2'
                    },
                    {
                      offset: 1,
                      color: '#3693fb'
                    }
                  ])
                },
                data: values
              }]
            })
            this.usedmemory = echarts.init(this.$refs.usedmemory, 'macarons')
            this.usedmemory.setOption({
              tooltip: {
                formatter: '{b} <br/>{a} : ' + this.cache.info.used_memory_human
              },
              series: [{
                name: '峰值',
                type: 'gauge',
                progress: {
                  show: true,
                  width: 14
                },
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                    offset: 0,
                    color: '#2881e6'
                  },
                  {
                    offset: 0.5,
                    color: '#308bf2'
                  },
                  {
                    offset: 1,
                    color: '#3693fb'
                  }
                ]),
                axisLine: {
                  lineStyle: {
                    width: 14
                  }
                },
                axisTick: {
                  show: false
                },
                splitLine: {
                  length: 4,
                  lineStyle: {
                    width: 2,
                    color: '#999'
                  }
                },
                anchor: {
                  show: true,
                  showAbove: true,
                  size: 20,
                  itemStyle: {
                    borderWidth: 4,
                    borderColor: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                        offset: 0,
                        color: '#2881e6'
                      },
                      {
                        offset: 0.5,
                        color: '#308bf2'
                      },
                      {
                        offset: 1,
                        color: '#3693fb'
                      }
                    ])
                  }
                },
                axisLabel: {
                  distance: 25,
                  color: '#999',
                  fontSize: 14
                },
                radius: '90%',
                center: ['50%', '56%'],
                min: 0,
                max: 1000,
                detail: {
                  fontSize: 24,
                  valueAnimation: true,
                  formatter: this.cache.info.used_memory_human
                },
                data: [{
                  value: parseFloat(this.cache.info.used_memory_human),
                  name: '内存消耗'
                }]
              }]
            })
          })
        })
      }
    }
  }
</script>
<style>
  .commandstats_bg {
    height: 160px;
    width: 160px;
    position: absolute;
    left: calc(50% - 80px);
    top: calc(50% - 70px);
    border-radius: 50%;
    background: #fff;
    box-shadow: 0 0 20px rgb(40 57 123 / 12%);
    z-index: 0;
  }

  .ant-descriptions-item-label {
    color: #888;
  }

  .ant-descriptions-item-content {
    color: #222;
    font-weight: 600;
  }

  .ant-descriptions-row>th,
  .ant-descriptions-row>td {
    padding-left: 24px;
  }
</style>
