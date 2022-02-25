<template>
  <div class="padding-card">
    <a-spin tip="玩命加载中..." :spinning="loading">
      <a-row :gutter="16">
        <a-col :span="24">
          <a-card title="基本信息" :bordered="false">
            <a-descriptions :column="5" >
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
            <div ref="commandstats" style="height: 420px ;z-index: 2;" />
            <div class="commandstats_bg"></div>
          </a-card>
        </a-col>
        <a-col :span="12">
          <a-card title="内存信息" :bordered="false">
            <div ref="usedmemory" style="height: 420px" />
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
      this.getList()
    },
    computed: {},
    watch: {},
    methods: {
      /** 查询服务信息 */
      getList () {
        this.loading = true
        getCache().then(response => {
          const cache = response.data
          this.cache = cache
          this.loading = false
          this.$nextTick(() => {
            this.commandstats = echarts.init(this.$refs.commandstats, 'macarons')
            this.commandstats.setOption({
              tooltip: {
                trigger: 'item',
                formatter: '{a} <br/> {b} : {c} ({d}%)'
              },
              color: ['#36cfc9', '#4CAF50', '#F44336', '#FF9800', '#c8d2e0', '#4782DA', '#9254de'],
              legend: {
                bottom: '1%',
                left: 'center',
                itemGap: 8 // 图例之间的间距
              },
              series: [{
                name: '命令',
                type: 'pie',
                radius: ['60%', '70%'],
                center: ['50%', '45%'],
                data: response.data.commandStats,
                animationEasing: 'cubicInOut',
                animationDuration: 1000,
                label: {
                  fontSize: '18',
                  fontWeight: 'bold',
                  position: 'center',
                  formatter: '{b}\r\n\r\n {c}({d}%) ',
                  show: true,
                  textBorderWidth: '16'
                },
                itemStyle: {
                  borderRadius: 0,
                  borderColor: '#fff',
                  borderWidth: 2
                },
                emphasis: {
                  label: {
                    fontSize: '18',
                    show: true,
                    textBorderWidth: '16'
                  }
                }
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
                color: ['#4782DA'],
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
                    borderColor: ['#4782DA']
                  }
                },
                axisLabel: {
                  distance: 25,
                  color: '#999',
                  fontSize: 14
                },
                radius: '90%',
                center: ['50%', '50%'],
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
<style >

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
     .ant-descriptions-item-label{
       color: #888;
     }
     .ant-descriptions-item-content{
       color: #222;
       font-weight: 600;
     }
      .ant-descriptions-row > th,  .ant-descriptions-row > td{
           padding-left: 24px;
     }

</style>
