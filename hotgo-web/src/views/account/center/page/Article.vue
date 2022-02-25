<template>
  <a-list
    size="large"
    rowKey="id"
    :loading="loading"
    itemLayout="vertical"
    :dataSource="data"
  >
    <a-list-item :key="item.id" slot="renderItem" slot-scope="item">
      <template slot="actions">
      </template>
      <a-list-item-meta>
        <a slot="title" href="https://vue.ant.design/">{{ item.title }}</a>
        <template slot="description">
          <span>
            <a-tag>Ant Design</a-tag>
            <a-tag>设计语言</a-tag>
            <a-tag>蚂蚁金服</a-tag>
          </span>
        </template>
      </a-list-item-meta>
      <article-list-content :description="item.description" :owner="item.owner" :avatar="item.avatar" :href="item.href" :updateAt="item.updatedAt" />
    </a-list-item>
    <div slot="footer" v-if="data.length > 0" style="text-align: center; margin-top: 16px;">
      <a-button @click="loadMore" :loading="loadingMore">加载更多</a-button>
    </div>
  </a-list>
</template>

<script>
import { ArticleListContent } from '@/components'

export default {
  name: 'Article',
  components: {
    ArticleListContent
  },
  data () {
    return {
      loading: true,
      loadingMore: false,
      data: []
    }
  },
  mounted () {
    this.getList()
  },
  methods: {
    getList () {
      for (let i = 0; i < 6; i++) {
        this.data.push({
          id: i,
          title: '色彩在界面设计中的作用',
          avatar: 'https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png',
          description: '图形化用户界面，用户在看到一个界面时，是一个由背景色、矩形按钮及图标ICON 设计、卡片等设计元素构成的。在设计界面时，借助色彩可以直观的展示背景色、导航栏、状态栏和操作按钮等构成的设计元素，并让产品设计的界面的逻辑架构和信息层级得到很好的展现。',
          star: 20,
          like: 20,
          message: 20,
          owner: 'Aidex',
          href: 'https://youdaoyun',
          updatedAt: '2020-07-22 23:11:33'
        })
      }
      this.loading = false
    },
    loadMore () {
      this.loadingMore = true
    }
  }
}
</script>

<style scoped>

</style>
