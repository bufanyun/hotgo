<template style="background:#f5f6fa;">
  <div class="typical-home" ref="portaletDiv">
    <a-spin :spinning="spinning" :delay="delayTime" tip="loading...">
      <a-row :gutter="[16, 16]" style="margin: ;">
        <!-- 顶部列表 -->
        <a-col :span="24" style="padding-bottom:0;">
          <a-row :gutter="[16, 16]" class="top-list">
            <a-col :xxl="6" :xl="8" :lg="16" style="height:164px;">
              <a
                style="display: block;height: 148px;overflow: hidden;border-radius: 4px;background: #516bf7;"
                v-if="bannerType === '1'"><img src="@/views/dashboard/images/banner1.png" /></a>
              <a
                style="display: block;height: 148px;overflow: hidden;border-radius: 4px;background: #516bf7;"
                v-if="bannerType === '2'"><img src="@/views/dashboard/images/banner2.png" /></a>
            </a-col>
            <a-col v-for="(item, index) in dataList" :key="index" :xxl="3" :xl="4" :lg="8">
              <a-card :bordered="false" @click="addTabMenu(item.path)" style="cursor: pointer;">
                <a-icon
                  :type="item.icon"
                  :component="allIcon[item.icon + 'Icon']"
                  :style="{ background: getMenuColor(index) }" />
                <span>{{ item.title }}</span>
              </a-card>
            </a-col>
          </a-row>
        </a-col>
      </a-row>
    </a-spin>
  </div>
</template>
<script>
  import allIcon from '@/core/icons'
  export default {
    name: 'DashboardIndex',
    props: {
      bannerType: {
        type: String,
        default: '1'
      }
    },
    data () {
      return {
        visible: false,
        spinning: false,
        delayTime: 100,
        dataList: [
          {
            path: '/org/user',
            icon: 'pound',
            title: '用户管理'
          },
          {
            path: '/org/dept',
            icon: 'lock',
            title: '部门管理'
          },
          {
            path: '/org/post',
            icon: 'unlock',
            title: '岗位管理'
          },
          {
            path: '/auth/role',
            icon: 'book',
            title: '角色管理'
          },
          {
            path: '/monitor/druid',
            icon: 'calendar',
            title: '数据监控'
          },
          {
            path: '/auth/sysAuth',
            icon: 'code',
            title: '菜单授权'
          },
          {
            path: '/tool/gen',
            icon: 'copy',
            title: '代码生成'
          },
          {
            path: '/sysSetting/menu',
            icon: 'windows',
            title: '菜单管理'
          },
          {
            path: '/sysSetting/dict',
            icon: 'aliwangwang',
            title: '字典管理'
          },
          {
            path: '/sysSetting/config',
            icon: 'code',
            title: '参数设置'
          },
          {
            path: '/log/operlog',
            icon: 'contacts',
            title: '操作日志'
          },
          {
            path: '/log/loginLog',
            icon: 'api',
            title: '登录日志'
          },
          {
            path: '/monitor/server',
            icon: 'idcard',
            title: '服务监控'
          },
          {
            path: '/monitor/cache',
            icon: 'shopping',
            title: '缓存监控'
          }
        ],
        changyongResultMenus: [], // 添加常用应用查询结果集合
        menus: [],
        sortable: undefined,
        buildMenus: [], // 重构菜单结构，用于首页显示
        activeChangyongTabKey: 'tab0',
        colorList: ['#5584fd', '#3470ff', '#ff8801', '#00d6b9', '#7e3bf3'],
        allIcon: allIcon
      }
    },
    computed: {
    },
    mounted () {
    },
    watch: {
    },
    created () {
    },
    methods: {
      getMenuColor (index) {
        return this.colorList[index % 4]
      },
      addTabMenu (path) {
        this.$router.push({
          path: path
        })
      }
    }
  }
</script>
<style lang="less">
  @import '../typical-home.less';
  .ant-list-item-drag {
    background: #ffffff;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
    box-sizing: border-box;
    z-index: 9999;
    border: 1px solid #dee0e3;
  }
  .common-iconbtn .anticon{
    margin: 0px 5px;
    color:rgba(0, 0, 0, 0.35);
    font-size: 16px;
  }
  .common-iconbtn .anticon-menu{
    cursor:move;
  }

</style>
