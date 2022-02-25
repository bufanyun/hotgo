<template style="background:#f5f6fa;">
  <div class="typical-home" ref="portaletDiv">
    <a-row :gutter="[24,24]" class="module-list">
      <a-col :span="8">
        <div class="module-in module-in01">
          <a-page-header :ghost="false" title="技术选型">
            <template slot="extra">
              <a-icon type="more" />
            </template>
          </a-page-header>
          <a-row :gutter="[24,24]">
            <a-col :span="12">
              <div class="technology">
                <a-list size="small" :split="false" :data-source="consoleTech">
                  <a-list-item slot="renderItem" slot-scope="item">
                    {{ item }}
                  </a-list-item>
                  <div slot="header">
                    后端
                  </div>
                </a-list>
              </div>
            </a-col>
            <a-col :span="12">
              <div class="technology">
                <a-list size="small" :split="false" :data-source="frontTech" >
                  <a-list-item slot="renderItem" slot-scope="item" >
                    {{ item }}
                  </a-list-item>
                  <div slot="header">
                    前端
                  </div>
                </a-list>
              </div>
            </a-col>
          </a-row>
        </div>
      </a-col>
      <a-col :span="8">
        <div class="module-in module-in01 ">
          <a-page-header :ghost="false" title="联系我们">
            <template slot="extra">
              <a-icon type="more" />
            </template>
          </a-page-header>
          <p>
            <a-alert message="专业UI设计（5年+）+VUE前端功能开发（3年+）" type="info" show-icon />
          </p>
          <a-list item-layout="horizontal" :data-source="detaildata" class="list-detail">
            <a-list-item class="module-text">
              <a-list-item-meta>
                <template slot="description">
                  <span>【1】</span>
                  <span>接AiDex Sharp系列架构的定制服务</span>
                </template>
              </a-list-item-meta>
            </a-list-item>
            <a-list-item class="module-text">
              <a-list-item-meta>
                <template slot="description">
                  <span>【2】</span>
                  <span >接3个月以内工期的vue、springboot、springcloud、app、小程序等软件定制服务</span>
                </template>
              </a-list-item-meta>
            </a-list-item>
            <a-list-item class="module-text">
              <a-list-item-meta>
                <template slot="description">
                  <span>【3】</span>
                  <span>UI原型页面设计及切图服务</span>
                </template>
              </a-list-item-meta>
            </a-list-item>
            <a-list-item class="module-text">
              <a-list-item-meta>
                <template slot="description">
                  <span>【4】</span>
                  <span>有意向请联系唯一指定联系人：皮皮大刺猬（右侧扫码）</span>
                </template>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </div>
      </a-col>
      <!-- 个人信息 -->
      <a-col :span="8">
        <div class="module-in module-in07" style="height: 370px;">
          <a-page-header :ghost="false" title="关于我">
            <template slot="extra">
              <a-icon type="more" />
            </template>
          </a-page-header>
          <a-row>
            <a-col>
              <a-col class="name">
                <span>皮皮大刺猬</span>
                <a-icon type="form" @click="compileshowModal" />
                <a-modal v-model="compile" title="Basic Modal" @ok="compilehandleOk" centered>
                  <p>Some contents...</p>
                  <p>Some contents...</p>
                  <p>Some contents...</p>
                </a-modal>
              </a-col>
              <a-col>
                <span>QQ号：</span>
                <span>1125373330</span>
              </a-col>
              <a-col>
                <span>QQ群：</span>
                <span>208511180</span>
              </a-col>
              <a-col>
                <span>微信号：</span>
                <span>big-hebgehog</span>
              </a-col>
              <a-col>
                <span>我的特长：</span>
                <span>专业UI设计（5年+）+VUE前端功能开发（3年+）</span>
              </a-col>
              <a-col>
                <span>项目：</span>
                <span><a style="color: red;" @click="goTarget('https://gitee.com/big-hedgehog/aidex-sharp')">若依-Adiex Sharp快速开发平台</a></span>
              </a-col>
              <a-col>
                <a-row>
                  <a-col :span="12">
                    <div class="wxcode" style="text-align: center;background: #f7f7f7;border: 1px solid #e8e8e8;width: 124px; margin: 0 auto;border-radius: 8px;padding-top: 8px;">
                      <img src="../images/wxCode.png" style="height:100px">
                      <p style="margin-bottom: 8px;">微信</p>
                    </div>
                  </a-col>
                  <a-col :span="12">
                    <div class="QQcode" style="text-align: center;background: #f7f7f7;border: 1px solid #e8e8e8;width: 124px; margin: 0 auto;border-radius: 8px;padding-top: 8px;">
                      <img src="../images/QQCode.png"style="height:100px">
                      <p style="margin-bottom: 8px;">QQ</p>
                    </div>
                  </a-col>
                </a-row>
              </a-col>
            </a-col>
          </a-row>

        </div>
      </a-col>
    </a-row>
    <a-row :gutter="[16,16]">
      <!-- 顶部列表 -->
      <a-col :span="16" class="module-list">
        <a-row :gutter="[16,16]">
          <!-- 待办事项 -->
          <a-col :span="12">
            <div class="module-in module-in01">
              <a-page-header :ghost="false" title="待办事项">
                <template slot="extra">
                  <a-icon type="more" />
                </template>
              </a-page-header>
              <a-list :grid="{ gutter: 8, column: 3 }" :data-source="data" class="list-totality">
                <a-list-item slot="renderItem" slot-scope="item">
                  <a-card>
                    <p>{{ item.title }}</p>
                    <span>{{ item.data }}</span>
                  </a-card>
                </a-list-item>
              </a-list>
              <a-list item-layout="horizontal" :data-source="detaildata" class="list-detail">
                <a-list-item slot="renderItem" slot-scope="item">
                  <a-list-item-meta>
                    <template slot="description">
                      <span>{{ item.state }}</span>
                      <span>{{ item.txt }}</span>
                      <a-tooltip placement="top">
                        <template slot="title">
                          <span>{{ item.hint }}</span>
                        </template>
                        <a-icon type="question-circle" />
                      </a-tooltip>
                      <a>办理</a>
                    </template>
                  </a-list-item-meta>
                </a-list-item>
              </a-list>
            </div>
          </a-col>
          <!-- 关注项目总览 -->
          <a-col :span="12">
            <div class="module-in module-in02">
              <a-page-header :ghost="false" title="关注项目总览">
                <template slot="extra">
                  <a-icon type="more" />
                </template>
              </a-page-header>
              <a-table :columns="projectcolumns" :data-source="projectdata" size="small">
                <a-badge slot="status" slot-scope="text,record" :status="record.icon" :text="record.text" />
              </a-table>
            </div>
          </a-col>
          <!-- 组件API搜索 -->
          <a-col :span="24">
            <div class="module-in module-in03">
              <a-page-header :ghost="false" title="组件API搜索">
                <template slot="extra">
                  <a-icon type="more" />
                </template>
              </a-page-header>
              <a-input placeholder="输入组件或 API 名称，支持中英文模糊搜索" />
              <div class="subtitle">最近使用</div>
              <!-- 默认最多只显示4个最近使用项 -->
              <a-list :grid="{ gutter: 16, column: 4 }" :data-source="apidata">
                <a-list-item slot="renderItem" slot-scope="item,">
                  <a-card>
                    <a-icon :type="item.icon" />
                    <span>{{ item.con }}</span>
                  </a-card>
                </a-list-item>
              </a-list>
            </div>
          </a-col>
          <!-- 趋势图 -->
          <a-col :span="24">
            <div class="module-in module-in04">
              <!--  <a-page-header
                :ghost="false"
                title="趋势图"
              >
                <template slot="extra">
                  <a-radio-group default-value="a" button-style="solid" size="small">
                    <a-radio-button value="a">
                      一个月
                    </a-radio-button>
                    <a-radio-button value="b">
                      三个月
                    </a-radio-button>
                    <a-radio-button value="c">
                      六个月
                    </a-radio-button>
                  </a-radio-group>
                </template>
              </a-page-header> -->
              <div style="height:410px;padding-top:8px ;">
                <!-- chart盒子 -->
                <line-radar></line-radar>
              </div>
            </div>
          </a-col>
          <!-- 帮助文档 -->
          <a-col :span="24">
            <div class="module-in module-in05">
              <a-page-header :ghost="false" title="帮助文档">
              </a-page-header>
              <a-row :gutter="[16,16]">
                <a-col :span="4">
                  <a-list size="small" :data-source="aiddata">
                    <a-list-item slot="renderItem" slot-scope="item">
                      {{ item }}
                    </a-list-item>
                    <div slot="header">
                      组件使用方法
                    </div>
                  </a-list>
                </a-col>
                <a-col :span="4">
                  <a-list size="small" :data-source="aiddata">
                    <a-list-item slot="renderItem" slot-scope="item">
                      {{ item }}
                    </a-list-item>
                    <div slot="header">
                      组件使用方法
                    </div>
                  </a-list>
                </a-col>
                <a-col :span="4">
                  <a-list size="small" :data-source="aiddata">
                    <a-list-item slot="renderItem" slot-scope="item">
                      {{ item }}
                    </a-list-item>
                    <div slot="header">
                      组件使用方法
                    </div>
                  </a-list>
                </a-col>
                <a-col :span="4">
                  <a-list size="small" :data-source="aiddata">
                    <a-list-item slot="renderItem" slot-scope="item">
                      {{ item }}
                    </a-list-item>
                    <div slot="header">
                      组件使用方法
                    </div>
                  </a-list>
                </a-col>
              </a-row>
            </div>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="8" class="module-list">
        <a-row :gutter="[0,16]">

          <!-- 公告 -->
          <a-col>
            <div class="module-in module-in06">
              <a-page-header :ghost="false" title="公告">
                <template slot="extra">
                  <a-icon type="more" />
                </template>
              </a-page-header>
              <a-tabs default-active-key="1" @change="callback" size="small">
                <a-tab-pane key="1" tab="首推">
                  <a-list item-layout="horizontal" :data-source="noticedata">
                    <a-list-item slot="renderItem" slot-scope="item">
                      <a-list-item-meta>
                        <a slot="description" :href="item.url">{{ item.state }}{{ item.title }}</a>
                      </a-list-item-meta>
                    </a-list-item>
                    <a-button type="link">
                      更多
                    </a-button>
                  </a-list>
                </a-tab-pane>
                <a-tab-pane key="2" tab="升级" force-render>
                  Content of Tab Pane 2
                </a-tab-pane>
                <a-tab-pane key="3" tab="安全">
                  Content of Tab Pane 3
                </a-tab-pane>
                <a-tab-pane key="4" tab="备案">
                  Content of Tab Pane 3
                </a-tab-pane>
                <a-tab-pane key="5" tab="其他">
                  Content of Tab Pane 3
                </a-tab-pane>
              </a-tabs>
            </div>
          </a-col>

          <!-- 产品热度 -->
          <a-col>
            <div class="module-in module-in08">
              <a-page-header :ghost="false" title="产品热度">
                <template slot="extra">
                  <a-icon type="more" />
                </template>
              </a-page-header>
              <a-list item-layout="horizontal" :data-source="heatdata">
                <a-list-item slot="renderItem" slot-scope="item">
                  <a-list-item-meta>
                    <a slot="description" :href="item.url">
                      <div class="order">{{ item.index }}</div>
                      <a-icon class="icon" :type="item.icon ? item.icon : 'pie-chart'" />
                      <div class="txt">{{ item.txt }}</div>
                      <a-statistic :value="item.value" :precision="1">
                        <template #suffix>
                          <a-icon :type="item.arrow ? item.arrow : 'pie-chart'" :style="{ color:item.color }" />
                        </template>
                      </a-statistic>
                    </a>
                  </a-list-item-meta>
                </a-list-item>
              </a-list>
            </div>
          </a-col>
          <!-- 统计信息 -->
          <a-col>
            <div class="module-in module-in09">
              <a-page-header :ghost="false" title="统计信息">
                <template slot="extra">
                  <a-date-picker :format="dateFormat" />
                  <a-icon type="more" />
                </template>
              </a-page-header>
              <div style="height:260px; text-align: center;">
                <!-- chart -->
                <img src="@/views/dashboard/images/img03.png" alt="">
              </div>
              <a-table
                :columns="statisticscolumns"
                :data-source="statisticsdata"
                size="small"
                :pagination="false"
                style="height:238px;">
                <span slot="tags" slot-scope="tags">
                  <a-tag :color="tags >= 0 ?'green': 'volcano' ">
                    {{ tags >= 0 ? '+' + tags+'%': tags+'%' }}
                  </a-tag>
                </span>
                <a-badge slot="status" slot-scope="text,record" :status="record.icon" :text="record.text" />
              </a-table>
            </div>
          </a-col>
        </a-row>
      </a-col>
    </a-row>
    </a-row></div>
</template>
<style type="less" >
  body{
    .modlule-show{
          height: 240px;
        .ant-btn-lg {
          height: 40px;
          padding: 0 15px;
          font-size: 16px;
          border-radius: 4px;
       }
    }
  }

</style>
<script>
  import lineRadar from '@/views/demo/chart/line-chart'
  // 顶部列表
  const dataList = [{
      icon: 'area-chart',
      text: 'OKR',
      color: '#5584fd'
    },
    {
      icon: 'pie-chart',
      text: '汇报',
      color: '#3470ff'
    },
    {
      icon: 'bar-chart',
      text: '日报',
      color: '#ff8801'
    },
    {
      icon: 'dot-chart',
      text: '周报',
      color: '#00d6b9'
    },
    {
      icon: 'line-chart',
      text: '审批',
      color: '#ff8801'
    },
    {
      icon: 'radar-chart',
      text: '审批',
      color: '#ff8801'
    },
    {
      icon: 'sliders',
      text: '公告',
      color: '#ff8801'
    },
    {
      icon: 'android',
      text: '月报',
      color: '#7e3bf3'
    }
  ]
  // 待办事项
  const data = [{
      title: '待审批',
      data: '10'
    },
    {
      title: '待提交',
      data: '10'
    },
    {
      title: '待阅示',
      data: '4'
    }
  ]
  const detaildata = [{
      state: '【待审批】',
      txt: '公告发布审批公告发布审批公告发布审批公告发布审批公告发布审批公告发布审批公告发布审批公告发布审批',
      hint: '公告发布审批说明1'
    },
    {
      state: '【待提交】',
      txt: '差旅费报销差旅费报销',
      hint: '公告发布审批说明2'
    },
    {
      state: '【待提交】',
      txt: '录用管理',
      hint: '公告发布审批说明3'
    },
    {
      state: '【待审批】',
      txt: '通知公告',
      hint: '公告发布审批说明4'
    },
    {
      state: '【待审批】',
      txt: '公告发布审批公告发布审批',
      hint: '公告发布审批说明5'
    }
  ]
  // 关注项目总览
  const projectcolumns = [{
      title: '项目名称',
      dataIndex: 'name',
      key: 'name',
      ellipsis: true
    },
    {
      title: '项目状态',
      dataIndex: 'status',
      scopedSlots: {
        customRender: 'status'
      },
      key: 'status',
      ellipsis: true,
      width: 96
    },
    {
      title: '责任人',
      dataIndex: 'take',
      key: 'take',
      ellipsis: true,
      width: 86
    }
  ]
  const projectdata = [{
      key: '1',
      name: '项目名称项目名称项目名称项目名称项目名称',
      icon: 'warning',
      text: '未开始',
      take: '张三'
    },
    {
      key: '2',
      name: '项目名称项目名称项目名称项目名称项目名称',
      icon: 'processing',
      text: '进行中',
      take: '李四'
    },
    {
      key: '3',
      name: '项目名称项目名名称项目名称',
      icon: 'processing',
      text: '进行中',
      take: '王五'
    },
    {
      key: '4',
      name: '项目名称项目名称项目名称项名称项目名称',
      icon: 'success',
      text: '已完成',
      take: '张三'
    },
    {
      key: '5',
      name: '项目名称项目名名称项目名称',
      icon: 'processing',
      text: '进行中',
      take: '王五'
    }
  ]
  // 组件API搜索
  const apidata = [{
      icon: 'pie-chart',
      con: '选择组件'
    },
    {
      icon: 'bar-chart',
      con: '上传文件组件'
    },
    {
      icon: 'dot-chart',
      con: '卡片式暂无数据组件'
    },
    {
      icon: 'line-chart',
      con: '计数文本框组件'
    }
  ]
  // 帮助文档
  const aiddata = [
    '组件使用方法',
    '组件使用方法',
    '组件使用方法',
    '组件使用方法',
    '组件使用方法'
  ]
  // 后端技术
  const consoleTech = [
   'SpringBoot',
   'Spring Security',
   'JWT',
   'MyBatis',
   'Druid',
   'Fastjson'
  ]
  // 前端技术
  const frontTech = [
   'Vue',
   'Vuex',
   'Ant-Design-VUE',
   'Axios',
   'Sass',
   'Quill'
  ]
  // 公告
  const noticedata = [{
      state: '【升级】',
      title: '1月6日DOS高防（国际）升级通知',
      url: 'https://www.baidu.com'
    },
    {
      state: '【升级】',
      title: '商标局2021年元旦期间服务器异常停止商标递交申请通知',
      url: 'https://www.antdv.com/'
    },
    {
      state: '【升级】',
      title: '1月消息列队',
      url: 'https://www.baidu.com'
    },
    {
      state: '【升级】',
      title: '1月6日DOS高防（国际）升级通知',
      url: 'https://www.baidu.com'
    },
    {
      state: '【升级】',
      title: '商标局2021年元旦期间服务器异常停止商标递交申请通知',
      url: 'https://www.baidu.com'
    }
  ]
  // 统计信息
  const statisticscolumns = [{
      title: '状态',
      dataIndex: 'status',
      scopedSlots: {
        customRender: 'status'
      },
      key: 'status',
      ellipsis: true
    },
    {
      title: '数量',
      dataIndex: 'quantity',
      key: 'quantity',
      ellipsis: true
    },
    {
      title: '增长率',
      dataIndex: 'tags',
      key: 'tags',
      scopedSlots: {
        customRender: 'tags'
      },
      ellipsis: true
    }
  ]
  const statisticsdata = [{
      key: '1',
      icon: 'error',
      text: '超期未完成',
      quantity: '10',
      tags: 35
    },
    {
      key: '2',
      icon: 'processing',
      text: '超期完成',
      quantity: '10',
      tags: -12
    },
    {
      key: '3',
      icon: 'warning',
      text: '进行中',
      quantity: '10',
      tags: 45
    },
    {
      key: '4',
      icon: 'default',
      text: '未开始',
      quantity: '10',
      tags: 24
    }
  ]
  // 热度产品
  const heatdata = [{
      index: '1',
      icon: 'copy',
      txt: '快速开发工具',
      value: '10.2',
      arrow: 'arrow-down',
      color: '#5fb38f',
      url: 'https://www.baidu.com'
    },
    {
      index: '2',
      icon: 'snippets',
      txt: '代码生成器',
      value: '11.2',
      arrow: 'arrow-up',
      color: '#ed2e2e',
      url: 'https://www.antdv.com'
    },
    {
      index: '3',
      icon: 'edit',
      txt: '流程设计器',
      value: '17.2',
      url: 'https://www.baidu.com'
    },
    {
      index: '4',
      icon: 'fund',
      txt: '菜单管理',
      value: '15.2',
      url: 'https://www.antdv.com'
    },
    {
      index: '5',
      icon: 'highlight',
      txt: '快速开发工具',
      value: '11.2',
      arrow: 'arrow-down',
      color: '#5fb38f',
      url: 'https://www.baidu.com'
    }
  ]
  // 添加弹出页
  const adddata = [{
      title: '审批',
      txt: '简单、高效、开放的审批工具',
      url: 'https://www.baidu.com',
      icon: 'ie',
      color: '#ff7741'
    },
    {
      title: '打卡',
      txt: '专业智能工具，实现高效考勤管理',
      url: 'https://www.baidu.com',
      icon: 'code-sandbox',
      color: '#3389ff'
    },
    {
      title: '公告',
      txt: '重要信息全员播报，特殊安排定向通知',
      url: 'https://www.baidu.com',
      icon: 'medium',
      color: '#ff7741'
    },
    {
      title: 'OKR',
      txt: '简单实用的团队目标管理工具',
      url: 'https://www.baidu.com',
      icon: 'slack',
      color: '#00d6b9'
    },
    {
      title: '打卡',
      txt: '专业智能工具，实现高效考勤管理',
      url: 'https://www.baidu.com',
      icon: 'ant-cloud',
      color: '#3389ff'
    }
  ]
  export default {
    data () {
      return {
        visible: false,
        data,
        detaildata,
        projectcolumns,
        projectdata,
        apidata,
        aiddata,
        consoleTech,
        frontTech,
        noticedata,
        dateFormat: 'YYYY-MM-DD',
        statisticscolumns,
        statisticsdata,
        compile: false,
        heatdata,
        adddata,
        dataList
      }
    },
    created () {
      this.reloadTab('1')
    },
    watch: {
    },
    mounted () {
      this.$emit('setHeight', this.$refs.portaletDiv.offsetHeight)
    },
    components: {
      lineRadar
    },
    methods: {
      getHeight () {
        return this.$refs.portaletDiv.offsetHeight
      },
      reloadTab (type) {
      },
      showModal () {
        this.visible = true
      },
      handleOk (e) {
        console.log(e)
        this.visible = false
      },
      compileshowModal () {
        this.compile = true
      },
      compilehandleOk (e) {
        console.log(e)
        this.compile = false
      },
      callback (key) {
        console.log(key)
      },
      onSearch (value) {
        console.log(value)
      },
      modalcallback (key) {
        console.log(key)
      },
      goTarget (href) {
        window.open(href, '_blank')
      },
      caseDisplay (href) {
        this.$message.success(
          '案例系统正在研发中，敬请期待！',
          3
         )
      }
    }
  }
</script>
<style lang="less">
  @import '../typical-home.less';
</style>
