<template>
  <div class="dasboard_body" style="padding:8px" ref="portaletDiv">
    <div>
      <a-row type="flex" :gutter="[16, 16]">
        <a-col :span="6" :order="1">
          <a-card :bordered="false">
            <div class="top_card">
              <a-statistic title="我的收益" :value="888888" :precision="2" style="" valueStyle="color:#6a7890;">
                <template #suffix>
                  <a-tag color="#ea331a" class="red_tag"> <a-icon type="arrow-up" />+15%</a-tag>
                </template>
              </a-statistic>
              <a-icon class="big_icon" type="pay-circle" />
            </div>
          </a-card>
        </a-col>
        <a-col :span="6" :order="2">
          <a-card :bordered="false">
            <div class="top_card">
              <a-statistic-countdown
                title="剩余时间"
                :value="deadline"
                format="HH:mm:ss:SSS"
                style="margin-right: 50px"
                valueStyle="color:#6a7890;"
              />
              <a-icon class="big_icon" type="clock-circle" />
            </div>
          </a-card>
        </a-col>
        <a-col :span="6" :order="3">
          <a-card :bordered="false">
            <div class="top_card">
              <a-statistic title="转换率" :value="888888" :precision="2" valueStyle="color:#6a7890;">
                <template #suffix>
                  <a-tag color="#ea331a" class="green_tag"> <a-icon type="arrow-down" />-21%</a-tag>
                </template>
              </a-statistic>
              <a-icon class="big_icon" type="smile" />
            </div>
          </a-card>
        </a-col>
        <a-col :span="6" :order="4">
          <a-card :bordered="false">
            <div class="top_card">
              <a-statistic title="转换率" :value="15" :precision="2" valueStyle="color:#6a7890;" suffix="%"></a-statistic>
              <a-icon class="big_icon" type="pie-chart" />
            </div>
          </a-card>
        </a-col>
      </a-row>
      <a-row type="flex" :gutter="[16, 16]" style="margin-top:8px">
        <a-col :span="24" :order="1">
          <a-card class="chart_card" :bordered="false">
            <line-radar></line-radar>
          </a-card>
        </a-col>
      </a-row>
      <a-row type="flex" :gutter="[16, 16]" style="margin-top:8px">
        <a-col :span="6" :order="1">
          <a-card :bordered="false" class="chart_card money">
            <div>
              <h4>月收入</h4>
            </div>
            <a-row type="flex" :gutter="[16, 16]" style="margin-top:8px">
              <a-col :span="24" :order="1"><h1>￥183,827</h1></a-col>
            </a-row>
            <div class="money_card">
              <p>上月增长</p>
              <span>￥1,838</span>
            </div>
            <div class="money_card" style="border:0;">
              <p>当月增长</p>
              <span>￥1,838</span>
            </div>
            <div class="tag_box">
              <a-tag color="#ea331a" class="green_tag">
                +23.91%
              </a-tag>
            </div>
          </a-card>
        </a-col>
        <a-col :span="12" :order="1">
          <a-card class="chart_card" :bordered="false">
            <histogramChart></histogramChart>
          </a-card>
        </a-col>
        <a-col :span="6" :order="1">
          <a-card :bordered="false" class="chart_card">
            <div style="position: relative;" >
              <pieChart></pieChart>
              <div class="pie_center"></div>
            </div>
          </a-card>
        </a-col>
      </a-row>
      <a-row type="flex" :gutter="[16, 16]" style="margin-top:8px">
        <a-col :span="12" :order="1">
          <a-card :bordered="false" class="chart_card">
            <smPieChart></smPieChart>
          </a-card>
        </a-col>
        <a-col :span="12" :order="1">
          <a-card :bordered="false" class="chart_card">
            <areaChart></areaChart>
          </a-card>
        </a-col>
      </a-row>
    </div>
  </div>
</template>
<script>
import lineRadar from './line-chart'
import histogramChart from './histogram-chart'
import dashboardChart from './dashboard-chart'
import chartRadar from './chart-radar'
import scatterChart from './scatter-chart'
import pieChart from './pie-chart'
import smPieChart from './sm-pieChart'
import areaChart from './area-chart'

export default {
  name: 'Config',
  components: {
    lineRadar,
    histogramChart,
    dashboardChart,
    chartRadar,
    scatterChart,
    pieChart,
    smPieChart,
    areaChart
  },
  data () {
    return {
      deadline: Date.now() + 1000 * 60 * 60 * 24 * 2 + 1000 * 30
    }
  },
  mounted () {
    this.$emit('setHeight', this.$refs.portaletDiv.offsetHeight)
  },
  filters: {},
  computed: {},
  watch: {},
  methods: {}
}
</script>
<style scoped lang="less">
.top_card {
  padding:16px;
}
.dasboard_body{
 .ant-card-body{
   padding: 0;
 }
}
.big_icon {
  font-size: 40px;
  position: absolute;
  top: 16px;
  right: 16px;
  color: #eaf0f7;
  z-index: 0;
}

.ant-tag-has-color {
  border-radius: 16px;
  padding: 4px 16px;
  border: 0;
  margin-left: 8px;
  font-size: 16px;
  vertical-align: top;
  font-size: 14px;
  margin-top: -2px;
  .anticon {
    margin-right: 4px;
  }
}
.red_tag {
  background: linear-gradient(45deg, #f8371c, #df3018);
}
.green_tag {
  background: linear-gradient(45deg, #50c51c, #49b21b);
}
.ant-progress-bg {
  height: 12px !important;
}
.chart_card {
  padding: 16px;
  .ant-card-head {
    padding: 0;
  }
}
.smBanner {
  background: linear-gradient(135deg, #148ae4, #148be4);
  border-radius: 4px;
  color: #fff;
}
.ant-card {
  border-radius: 4px;
}
.money {
  h4 {
    margin: 8px;
    margin-bottom: 16px;
    font-size: 16px;
    font-weight: 600;
    color: #333;
  }
  h1 {
    font-size: 40px;
    font-weight: 400;
    color: #3794fc;
    font-family: Roboto;
    text-align: center;
    margin: 16px 0;
  }
  .tag_box {
    .green_tag {
      font-size: 18px;
      line-height: 32px;
      text-align: center;
      width: 140px;
      display: block;
      color: #fff;
      border-radius: 18px;
      margin: 16px auto;
    }
  }
  .money_card {
    margin: 16px 0;
    width: 50%;
    display: inline-block;
    border-right: 1px solid #dee5ef;

    p {
      font-size: 14px;
      padding: 10px;
      text-align: center;
      margin: 0;
      color: #97a4ba;
    }
    span {
      display: block;
      color: #67758d;
      font-size: 24px;
      line-height: 54px;
      text-align: center;
    }
  }
}
.smBanner {
  padding-left: 140px;
  .smBanner_left {
    width: 100px;
    position: absolute;
    left: 16px;
    top: 16px;
  }
  .smBanner_right {
    display: block;
    position: absolute;
    left: 100px;
    top: 0;
    bottom: 0;
    right: 4px;
  }
}
.pie_center{
  width: 140px;
  height: 140px;
  border-radius:50%;
  box-shadow:  0 0 10px rgba(40, 57, 123, .2);
  position: absolute;
  left: 50%;
  top: 50%;
  margin-left: -70px ;
  margin-top: -70px ;
}
</style>
