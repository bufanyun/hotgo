<template>
  <div class="padding-card monitor-box" style="margin-top: -16px;padding-bottom: 16px;">
    <a-spin tip="正在加载服务监控数据，请稍后" :spinning="loading">
      <a-space direction="vertical" size="middle">
        <a-row class="columns-list-ul" :gutter="16">
          <a-col :span="6">
            <a-card :bordered="false">
              <div class="text-number">{{ server.cpu ? server.cpu.cpuNum : 0 }}</div>
              <div class="title-text">CPU核心数</div>
            </a-card>
          </a-col>
          <a-col :span="6">
            <a-card :bordered="false">
              <div class="text-number">{{ server.cpu ? server.cpu.sys : 0 }}%</div>
              <div class="title-text">
                系统CPU使用率
              </div>
            </a-card>
          </a-col>
          <a-col :span="6">
            <a-card :bordered="false">
              <div class="text-number">{{ server.cpu ? server.cpu.used : 0 }}%</div>
              <div class="title-text">
                用户CPU使用率
                <!-- 比昨日 <em class="drop">- 0.05%</em> -->
              </div>
            </a-card>
          </a-col>
          <a-col :span="6">
            <a-card :bordered="false">
              <div class="text-number">{{ server.cpu ? server.cpu.free : 0 }}%</div>
              <div class="title-text">
                当前CPU空闲率
                <!-- 比昨日 <em class="rise">+ 1.25%</em> -->
              </div>
            </a-card>
          </a-col>
        </a-row>
        <a-row class="monitor-cardbox" :gutter="24">
          <a-col :span="24">
            <a-card size="small" :bordered="false" style="height: 180px;">
              <template slot="title">
                <a-icon type="fire" />
                内存
              </template>
              <a-row :gutter="16">
                <a-col :span="6" :key="item.id" v-for="item in memData">
                  <a-card-grid>
                    <div class="card-grid-name">{{ item.name }}</div>
                    <a-col class="text01" :span="12">
                      <a-statistic v-if="item.name != '使用率'" title="内存" :value="item.mem">
                        <template #suffix>
                          G
                        </template>
                      </a-statistic>
                      <a-statistic v-if="item.name == '使用率'" title="内存" :value="item.mem" :value-style="{ color: item.mem > 80 ? 'red' : '' }">
                        <template #suffix>
                          %
                          <a-icon v-if="item.jvm > 80" type="warning" style="color:#FFCC00" />
                        </template>
                      </a-statistic>
                    </a-col>
                    <a-col class="text02" :span="12">
                      <a-statistic v-if="item.name != '使用率'" title="JVM" :value="item.jvm">
                        <template #suffix>
                          M
                        </template>
                      </a-statistic>
                      <a-statistic v-if="item.name == '使用率'" title="JVM" :value="item.jvm" :value-style="{ color: item.jvm > 80 ? 'red' : '' }">
                        <template #suffix>
                          %
                          <a-icon v-if="item.jvm > 80" type="warning" style="color:#FFCC00" />
                        </template>
                      </a-statistic>
                    </a-col>
                  </a-card-grid>
                </a-col>
              </a-row>
            </a-card>
          </a-col>
        </a-row>
        <a-row :gutter="24">
          <a-col :span="24">
            <a-card size="small" :bordered="false">
              <template slot="title">
                <a-icon type="laptop" />
                磁盘状态
              </template>
              <a-row :gutter="12">
                <a-col class="gutter-row" :span="6" v-for="(item, i) in sysData" :key="item.id">
                  <div class="disc-list-ul" :class="[i % 3 === 0 && i !== 0 ? '' : 'list-border']">
                    <div class="disc-list-title">
                      <div class="disc-list-title-text">
                        <div class="disc-caption-text">
                          {{ item.dirName }}
                          <em>磁盘类型：{{ item.sysTypeName }}</em>
                        </div>
                        <div class="stat-pic">
                          <a-progress
                            type="circle"
                            :stroke-color="{
                              '0%': '#007dff',
                              '100%': '#00c9ff'
                            }"
                            :percent="item.usage"
                            :status="item.usage > 80 ? 'exception' : 'normal'"
                          >
                            <template #format="percent">
                              <span>{{ item.usage }}%</span>
                            </template>
                          </a-progress>
                        </div>
                        <div class="disc-other">
                          <a-row :gutter="24">
                            <a-col :span="12">
                              总大小
                              <br />
                              <em class="disc-other-color01">{{ item.total }}</em>
                            </a-col>
                            <a-col :span="12">
                              可用大小
                              <br />
                              <em class="disc-other-color02">{{ item.free }}</em>
                            </a-col>
                          </a-row>
                        </div>
                      </div>
                    </div>
                  </div>
                </a-col>
              </a-row>
            </a-card>
          </a-col>
        </a-row>
        <a-row :gutter="24">
          <a-col :span="24">
            <a-card class="monitor-card" size="small" :bordered="false">
              <template slot="title">
                <a-icon type="windows" />
                服务器信息
              </template>
              <a-descriptions :column="2" size="middle">
                <a-descriptions-item label="服务器名称">{{ server.sys ? server.sys.computerName : '' }}</a-descriptions-item>
                <a-descriptions-item label="操作系统">{{ server.sys ? server.sys.osName : '' }}</a-descriptions-item>
                <a-descriptions-item label="服务器IP">{{ server.sys ? server.sys.computerIp : '' }}</a-descriptions-item>
                <a-descriptions-item label="系统架构">{{ server.sys ? server.sys.osArch : '' }}</a-descriptions-item>
              </a-descriptions>
            </a-card>
          </a-col>
        </a-row>
        <a-row :gutter="24">
          <a-col :span="24">
            <a-card class="monitor-card" size="small" :bordered="false">
              <template slot="title">
                <a-icon type="hdd" />
                Java虚拟机信息
              </template>
              <a-descriptions :column="5" size="middle">
                <a-descriptions-item label="Java名称" span="3">{{ server.jvm ? server.jvm.name : '' }}</span></a-descriptions-item>
                <a-descriptions-item label="Java版本" span="2">{{ server.jvm ? server.jvm.version : '' }}</a-descriptions-item>
                <a-descriptions-item label="启动时间" span="3">{{ server.jvm ? server.jvm.startTime : '' }}</a-descriptions-item>
                <a-descriptions-item label="运行时长" span="2">{{ server.jvm ? server.jvm.runTime : '' }}</a-descriptions-item>
                <a-descriptions-item label="安装路径" span="3">{{ server.jvm ? server.jvm.home : '' }}</a-descriptions-item>
                <a-descriptions-item label="项目路径" span="2">{{ server.jvm ? server.sys.userDir : '' }}</a-descriptions-item>
              </a-descriptions>
            </a-card>
          </a-col>
        </a-row>
      </a-space>
    </a-spin>
  </div>
</template>

<script>
import { getServer } from '@/api/monitor/server'

export default {
  name: 'Server',
  data () {
    return {
      server: [],
      loading: true,
      memColumns: [
        {
          title: '属性',
          dataIndex: 'name'
        },
        {
          title: '内存',
          dataIndex: 'mem',
          scopedSlots: { customRender: 'mem' }
        },
        {
          title: 'JVM',
          dataIndex: 'jvm',
          scopedSlots: { customRender: 'jvm' }
        }
      ],
      memData: [
        {
          name: '总内存',
          mem: '0',
          jvm: '0'
        },
        {
          name: '已用内存',
          mem: '0',
          jvm: '0'
        },
        {
          name: '剩余内存',
          mem: '0',
          jvm: '0'
        },
        {
          name: '使用率',
          mem: '0',
          jvm: '0'
        }
      ],
      newMemData: {
        name: '使用率',
        mem: '0',
        usedMem: '0',
        freeMem: '0',
        jvm: '0',
        usedJvm: '0',
        freeJvm: '0'
      },
      sysColumns: [
        {
          title: '盘符路径',
          dataIndex: 'dirName',
          ellipsis: true
        },
        {
          title: '文件系统',
          dataIndex: 'sysTypeName'
        },
        {
          title: '盘符类型',
          dataIndex: 'typeName',
          ellipsis: true
        },
        {
          title: '总大小',
          dataIndex: 'total',
          scopedSlots: { customRender: 'total' }
        },
        {
          title: '可用大小',
          dataIndex: 'free',
          scopedSlots: { customRender: 'free' }
        },
        {
          title: '已用大小',
          dataIndex: 'used',
          scopedSlots: { customRender: 'used' }
        },
        {
          title: '已用百分比',
          dataIndex: 'usage',
          scopedSlots: { customRender: 'usage' }
        }
      ],
      sysData: []
    }
  },
  filters: {},
  created () {
    this.getList('0')
  },
  mounted () {
    this.timer = setInterval(() => {
      this.getList('1')
    }, 10000)
  },
  beforeDestroy () {
    clearInterval(this.timer)
  },
  computed: {},
  watch: {},
  methods: {
    /** 查询服务信息 */
    getList (type) {
      if (type === '0') {
        this.loading = true
      }
      getServer().then(response => {
        const serverData = response.data
        this.server = serverData
        this.memData = [
          {
            name: '服务器总内存',
            mem: serverData.mem.total,
            jvm: serverData.jvm.total
          },
          {
            name: '已用内存',
            mem: serverData.mem.used,
            jvm: serverData.jvm.used
          },
          {
            name: '剩余内存',
            mem: serverData.mem.free,
            jvm: serverData.jvm.free
          },
          {
            name: '使用率',
            mem: serverData.mem.usage,
            jvm: serverData.jvm.usage
          }
        ]
        this.newMemData = {
          name: '使用率',
          mem: Number(serverData.mem.usage).toFixed(),
          usedMem: serverData.mem.used,
          freeMem: serverData.mem.free,
          jvm: Number(serverData.jvm.usage).toFixed(),
          usedJvm: serverData.jvm.used,
          freeJvm: serverData.jvm.free
        }
        this.sysData = serverData.sysFiles
        setTimeout(() => {
          if (type === '0') {
            this.loading = false
          }
        }, 500)
      })
    }
  }
}
</script>
<style lang="less">
/* .ant-table-thead > tr > th {
      background: #fff !important;
  } */
.monitor-box {
  .ant-card-head {
    border-bottom: 0;
  }
  .ant-card-head-title {
    font-size: 14px;
    font-weight: bold;
    color: #262626;
    i.anticon {
      color: #2496ff;
    }
  }
  .ant-card-grid {
    width: 100%;
    padding: 0;
    box-shadow: none;
    background: #f7f9fa;
  }
  .card-grid-name {
    background: #f0f3f5;
    text-align: center;
    color: #333333;
    font-size: 14px;
    line-height: 40px;
  }
  .columns-list-ul .ant-card-body {
    text-align: center;
    padding: 10px 24px 20px;
  }
  .title-text {
    font-size: 12px;
    margin: 0;
    padding: 0;
    color: #666666;
  }
  .title-text em {
    font-style: normal;
  }
  .title-text em.rise {
    color: #cf1221;
  }
  .title-text em.drop {
    color: #347f00;
  }
  .text-number {
    color: #000000;
    font-size: 42px;
  }
  .disc-list-ul {
    padding: 10px 20px 10px 15px;
    margin-bottom: 15px;
  }
  .disc-caption-text {
    font-size: 16px;
    font-weight: bold;
    color: #333;
    margin-bottom: 5px;
    em {
      font-size: 12px;
      color: #858585;
      font-style: normal;
      font-weight: normal;
    }
  }
  .ant-progress-text {
    color: #575757;
    font-size: 18px;
    font-weight: bold;
  }
  .disc-other {
    font-size: 12px;
    color: #666666;
    line-height: 20px;
    text-align: center;
  }
  .disc-other em {
    font-style: normal;
    font-size: 22px;
    margin-top: 10px;
    display: inline-block;
  }
  .disc-other em.disc-other-color01 {
    color: #2496ff;
  }
  .disc-other em.disc-other-color02 {
    color: #3f8652;
  }
  .stat-pic {
    display: inline-block;
    width: 100%;
    text-align: center;
    margin-bottom: 12px;
  }
  .list-border {
    border-right: 1px solid #ececec;
  }
  .monitor-card .ant-card-body {
    padding: 0 0 10px;
  }
  .ant-descriptions {
    font-size: 12px;
  }
  .ant-descriptions-item-label {
    font-size: 14px;
    color: #666666;
  }
  .ant-descriptions-item-content {
    font-size: 14px;
    color: #000000;
    font-weight: normal;
  }
  .ant-card-grid {
    font-size: 12px;
  }
  .ant-statistic-title {
    font-size: 12px;
    color: #666666;
  }
  .ant-statistic-content {
    font-size: 14px;
  }
  .ant-statistic-content-value-int,
  .ant-statistic-content-value-decimal {
    font-size: 22px;
  }
  .ant-statistic-content-suffix {
    font-size: 12px;
    color: #666666;
  }
  .card-grid-name {
    color: #666666;
    font-size: 14px;
    font-weight: bold;
  }
  .ant-card-grid-hoverable:hover {
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.05);
  }
  .monitor-card .ant-descriptions-item {
    padding: 10px 30px;
  }
  .monitor-card .ant-descriptions-row:nth-child(odd) {
    background: #fbfbfb;
    border-top: 1px solid #f8f8f8;
  }
  .monitor-card .ant-descriptions-row:nth-child(even) {
    border-top: 1px solid #f8f8f8;
  }
  .ant-statistic {
    text-align: center;
    padding: 10px;
  }
  .monitor-cardbox .ant-card-body {
    padding: 10px;
  }
  .monitor-cardbox .text01 .ant-statistic-content {
    color: #3f8652;
  }
  .monitor-cardbox .text02 .ant-statistic-content {
    color: #ff6600;
  }
}
</style>
