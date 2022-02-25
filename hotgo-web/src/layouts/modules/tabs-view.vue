<template>
  <div class="tabs-view">
    <a-tabs
      v-if="panes.length !== 0"
      @change="changePage"
      v-model="activeKey"
      hide-add
      type="editable-card"
      @edit="onEdit"
    >
      <template v-for="(pane, index) in panes" :tab="pane.title" :closable="pane.closable">
        <a-tab-pane :key="pane.fullPath" >
          <template #tab>
            <a-dropdown :trigger="['contextmenu']">
              <div style="display: inline-block">
                {{ pane.title }}
              </div>
              <template v-slot:overlay>
                <a-menu style="user-select: none">
                  <a-menu-item @click="reloadPage(pane)" :disabled="activeKey !== pane.fullPath" key="1">
                    <a-icon type="reload" />
                    刷新
                  </a-menu-item>
                  <a-menu-item @click="removeTab(pane, index)" key="2">
                    <a-icon type="close" />
                    关闭
                  </a-menu-item>
                  <a-menu-divider/>
                  <a-menu-item @click="closeLeft(pane, index)" key="3">
                    <a-icon type="vertical-right" />
                    关闭左侧
                  </a-menu-item>
                  <a-menu-item @click="closeRight(pane, index)" key="4">
                    <a-icon type="vertical-left" />
                    关闭右侧
                  </a-menu-item>
                  <a-menu-divider/>
                  <a-menu-item @click="closeOther(pane, index)" key="5">
                    <a-icon type="column-width" />
                    关闭其他
                  </a-menu-item>
                  <a-menu-item @click="closeAll" key="6">
                    <a-icon type="minus" />
                    关闭全部
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </template>
        </a-tab-pane>
      </template>
    </a-tabs>
  </div>
</template>

<script>
export default {
  name: 'TabsView',
  data () {
    const panes = [
      { title: '工作台', key: '/dashboard/workplace', fullPath: '/dashboard/workplace' }
    ]
    return {
      activeKey: panes[0].key,
      panes,
      newTabIndex: 0
    }
  },
  created () {},
  // 监听,当路由发生变化的时候执行
  watch: {
    $route: {
      handler: function (val, oldVal) {
        const title = val.meta.title
        const fullPath = val.fullPath
        const routerObj = Object.assign({
          'title': title,
          'fullPath': fullPath,
          'key': fullPath
        })
        // 先判断标签路由中是否已经存在相同的,如果存在相同的跳转后不push
        let flag = true
        this.panes.forEach(item => {
          if (item.fullPath === fullPath) {
            flag = false
          }
        })
        if (flag) {
          this.panes.push(routerObj)
        }
        this.activeKey = fullPath
      },
      // 深度观察监听
      deep: true
    }
  },
  methods: {
    remove (targetKey) {
      let index = -1
      this.panes.forEach((pane, i) => {
        if (pane.key === targetKey) {
          index = i
        }
      })
      if (index !== -1) {
        this.panes.splice(index, 1)
        if (this.panes.length === 0) {
          this.$router.push({
            path: '/dashboard/workplace'
          })
        }
        const prevItem = this.panes[index - 1]
        if (prevItem) {
          this.activeKey = prevItem.fullPath
          this.$router.push({
            path: prevItem.fullPath
          })
        }
        const nextItem = this.panes[index + 1]
        if (nextItem) {
          this.activeKey = nextItem.fullPath
          this.$router.push({
            path: nextItem.fullPath
          })
        }
      }
    },
    onEdit (targetKey, action) {
      this[action](targetKey)
    },
    changePage (key) {
      this.activeKey = key
      this.$router.push({ path: key })
    },
    // 刷新页面
    reloadPage (pane) {
      this.$emit('reload')
    },
    // 关闭当前
    removeTab (pane, index) {
      // 判断如果当前是最后一个元素了,就跳转到工作台
      this.panes.splice(index, 1)
      if (this.panes.length === 0) {
        this.$router.push({
          path: '/dashboard/workplace'
        })
      }
      const prevItem = this.panes[index - 1]
      if (prevItem) {
        this.activeKey = prevItem.fullPath
        this.$router.push({
          path: prevItem.fullPath
        })
      }
      const nextItem = this.panes[index + 1]
      if (nextItem) {
        this.activeKey = nextItem.fullPath
        this.$router.push({
          path: nextItem.fullPath
        })
      }
    },
    // 关闭其他
    closeOther (pane, index) {
      const delIndex = index + 1
      // 删除其所有右侧元素
      this.panes.splice(delIndex, this.panes.length - delIndex)
      // 删除其所有左侧元素
      this.panes.splice(0, index)
    },
    // 关闭左侧
    closeLeft (pane, index) {
      this.panes.splice(0, index)
    },
    // 关闭右侧
    closeRight (pane, index) {
      const delIndex = index + 1
      this.panes.splice(delIndex, this.panes.length - delIndex)
    },
    // 关闭全部
    closeAll () {
      // 删除数组全部元素->数组置空
      this.panes = []
      this.$router.push({
        path: '/index'
      })
    }
  }
}
</script>
<style lang="less">
  .ant-tabs-nav-container{
    background: #fff;
    border-bottom: 1px solid #e8e8e8
  }
  .ant-tabs-nav-wrap{
    padding: 0 8px;
  }
</style>
