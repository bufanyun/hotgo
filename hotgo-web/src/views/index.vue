<template style="background:#f5f6fa;">
  <div class="typical-home" >
    <a-row :gutter="[4,4]">
      <!-- 顶部列表 -->
      <a-col :span="24" style="padding-bottom:0;">
        <a-row :gutter="[16,16]" class="top-list">
          <a-col :xxl="6" :xl="12" :lg="16" style="margin: 8px;height: 148px;border-radius: 4px;overflow: hidden;padding: 0;">
            <a href="##">
            </a>
          </a-col>
          <a-col v-for="(item,index) in dataList" :key="index" :xxl="3" :xl="6" :lg="8" >
            <a-card :bordered="false">
              <a-icon :type="item.icon" :style="{ background:item.color }"/>
              <span>{{ item.text }}</span>
            </a-card>
          </a-col>
          <a-col :xxl="3" :xl="6" :lg="8">
            <a-card :bordered="false" @click="showModal">
              <a-icon type="plus" class="add-plus"/>
              <span>添加常用</span>
              <!-- ref="subcard" -->
              <a-modal v-model="visible" title="添加常用应用" centered @ok="handleOk" wrapClassName="top-list-modal">
                <a-input-search placeholder="请输入需要搜索的内容" @search="onSearch" />
                <a-tabs default-active-key="1" @change="modalcallback" tabPosition="left" size="small">
                  <a-tab-pane key="1" tab="最近使用">
                    <a-list item-layout="horizontal" :data-source="adddata">
                      <a-list-item slot="renderItem" slot-scope="item">
                        <a-list-item-meta :description="item.txt">
                          <a slot="title" :href="item.url">{{ item.title }}</a>
                          <a-avatar slot="avatar" :size="34" :icon="item.icon" :style="{ background:item.color }"/>
                        </a-list-item-meta>
                        <a-button size="small">移除</a-button>
                      </a-list-item>
                    </a-list>
                  </a-tab-pane>
                  <a-tab-pane key="2" tab="分析工具">
                    分析工具
                  </a-tab-pane>
                  <a-tab-pane key="3" tab="沟通协作">
                    沟通协作
                  </a-tab-pane>
                  <a-tab-pane key="4" tab="客户服务">
                    客户服务
                  </a-tab-pane>
                  <a-tab-pane key="5" tab="开发工具">
                    开发工具
                  </a-tab-pane>
                  <a-tab-pane key="6" tab="人力资源">
                    人力资源
                  </a-tab-pane>
                  <!-- <a-tab-pane key="7" tab="办公管理">
                    办公管理
                  </a-tab-pane>
                  <a-tab-pane key="8" tab="市场营销">
                    市场营销
                  </a-tab-pane>
                  <a-tab-pane key="9" tab="财务管理">
                    财务管理
                  </a-tab-pane>
                  <a-tab-pane key="10" tab="生产效率">
                    生产效率
                  </a-tab-pane> -->
                </a-tabs>
              </a-modal>
            </a-card>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="16" class="module-list">
        <a-row :gutter="[16,16]">
          <!-- 待办事项 -->
          <a-col :span="12">
            <div class="module-in module-in01">
              <a-page-header
                :ghost="false"
                title="待办事项"
              >
                <template slot="extra">
                  <a-icon type="dash" />
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
              <a-page-header
                :ghost="false"
                title="关注项目总览"
              >
                <template slot="extra">
                  <a-icon type="dash" />
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
              <a-page-header
                :ghost="false"
                title="组件API搜索"
              >
                <template slot="extra">
                  <a-icon type="dash" />
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
              <a-page-header
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
              </a-page-header>
              <div style="height:280px;">
                <!-- chart盒子 -->
              </div>
            </div>
          </a-col>
          <!-- 帮助文档 -->
          <a-col :span="24">
            <div class="module-in module-in05">
              <a-page-header
                :ghost="false"
                title="帮助文档"
              >
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
              <a-page-header
                :ghost="false"
                title="公告"
              >
                <template slot="extra">
                  <a-icon type="dash" />
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
          <!-- 个人信息 -->
          <a-col>
            <div class="module-in module-in07">
              <a-page-header
                :ghost="false"
                title="个人信息"
              >
                <template slot="extra">
                  <a-icon type="dash" />
                </template>
              </a-page-header>
              <a-row type="flex">
                <a-col flex="110px">
                  <a-avatar size="large" icon="user" />
                </a-col>
                <a-col>
                  <a-col class="name">
                    <span>Aidex</span>
                    <a-icon type="form" @click="compileshowModal"/>
                    <a-modal v-model="compile" title="Basic Modal" @ok="compilehandleOk" centered>
                      <p>Some contents...</p>
                      <p>Some contents...</p>
                      <p>Some contents...</p>
                    </a-modal>
                  </a-col>
                  <a-col>
                    <span>员工编号：</span>
                    <span>88888888888</span>
                  </a-col>
                  <a-col>
                    <span>组织机构：</span>
                    <span>XXXXX</span>
                  </a-col>
                  <a-col>
                    <span>所属群组：</span>
                    <span>XXXX</span>
                  </a-col>
                  <a-col>
                    <span>所属部门：</span>
                    <span>市场部</span>
                  </a-col>
                </a-col>
              </a-row>
            </div>
          </a-col>
          <!-- 产品热度 -->
          <a-col>
            <div class="module-in module-in08">
              <a-page-header
                :ghost="false"
                title="产品热度"
              >
                <template slot="extra">
                  <a-icon type="dash" />
                </template>
              </a-page-header>
              <a-list item-layout="horizontal" :data-source="heatdata">
                <a-list-item slot="renderItem" slot-scope="item">
                  <a-list-item-meta>
                    <a slot="description" :href="item.url">
                      <div class="order">{{ item.index }}</div>
                      <a-icon class="icon" :type="item.icon" />
                      <div class="txt">{{ item.txt }}</div>
                      <a-statistic
                        :value="item.value"
                        :precision="1"
                      >
                        <template #suffix>
                          <a-icon :type="item.arrow" :style="{ color:item.color }"/>
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
              <a-page-header
                :ghost="false"
                title="统计信息"
              >
                <template slot="extra">
                  <a-date-picker :default-value="moment('2015/01/01', dateFormat)" :format="dateFormat" />
                  <a-icon type="dash" />
                </template>
              </a-page-header>
              <div style="height:260px; text-align: center;">
                <!-- chart -->
              </div>
              <a-table :columns="statisticscolumns" :data-source="statisticsdata" size="small" :pagination="false" style="height:238px;">
                <span slot="tags" slot-scope="tags">
                  <a-tag
                    :color="tags >= 0 ?'green': 'volcano' "
                  >
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
  </div>
</template>
<script>
  // 顶部列表
  const dataList = [
    {
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
  const data = [
    {
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
  const detaildata = [
    {
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
  const projectcolumns = [
    {
      title: '项目名称',
      dataIndex: 'name',
      key: 'name',
      ellipsis: true
    },
    {
      title: '项目状态',
      dataIndex: 'status',
      scopedSlots: { customRender: 'status' },
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
  const projectdata = [
    {
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
    },
    {
      key: '6',
      name: '项目名称项目名称项目名称项名称项目名称',
      icon: 'processing',
      text: '进行中',
      take: '张三'
    }
  ]
  // 组件API搜索
  const apidata = [
    {
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
  // 公告
  const noticedata = [
    {
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
  const statisticscolumns = [
    {
      title: '状态',
      dataIndex: 'status',
      scopedSlots: { customRender: 'status' },
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
      scopedSlots: { customRender: 'tags' },
      ellipsis: true
    }
  ]
  const statisticsdata = [
    {
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
  const heatdata = [
    {
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
  const adddata = [
    {
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
    methods: {
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
      }
    }
  }
</script>
<style lang="less">
  @import 'dashboard/typical-home.less';
</style>
