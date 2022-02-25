<template>
  <div class="typical-home" ref="portaletDiv">
    <a-row :gutter="[24,24]" class="module-list">
      <!-- 待办事项 -->
      <a-col :span="24" >
        <div class="module-in module-in01 modlule-show" style="height: 260px;">
          <a-page-header :ghost="false" title="平台简介">
            <template slot="extra">
              <a-icon type="more" />
            </template>
          </a-page-header>
          <p>
            <a-alert message="免费开源【欢迎star和fork项目,请点击下方按钮访问】" type="info" show-icon />
          </p>
          <p style="text-indent: 30px;">
            基于若依-ruoyi-vue项目扩展，前端采用Ant-Design-VUE，代码易读易懂、界面简洁美观，不仅仅是一个后台开发框架，它是一个企业级快速开发解决方案，我们将把UI交互、快速开发能力追求到极致，适配国产数据库，国产中间件，将支持多租户、flowable工作流，移动APP，更多插件正在扩展中
          </p>
          <p>
            当前版本：V3.0（后台与ruoyi-vue定期同步）
          </p>

          <p>
            <a-row :gutter="[24,24]">
              <a-col :span="8">
                <a-button block style="height: 40px;" type="primary" @click="goTarget('https://gitee.com/big-hedgehog/aidex-sharp')">
                  访问AiDex
                </a-button>
              </a-col>
              <a-col :span="8">
                <a-button block style="height: 40px;" type="primary" ghost @click="goTarget('https://www.bilibili.com/video/BV13L411u7oC?spm_id_from=333.999.0.0')">
                  我的B站
                </a-button>
              </a-col>
              <a-col :span="8">
                <a-button block style="height: 40px;" type="primary" ghost @click="goTarget('https://gitee.com/y_project/RuoYi-Vue')">
                  访问若依
                </a-button>
              </a-col>
            </a-row>
          </p>
        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    data () {
      return {

      }
    },
    computed: {
    },
    created () {
    },
    methods: {
      goTarget (url) {
        window.open(url)
      },
      caseDisplay (url) {
        this.$message.success(
          '案例正在研发中，目前尚未实现，敬请期待！',
          3
         )
      }
    }
  }
</script>
