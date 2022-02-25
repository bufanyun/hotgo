<template>
  <div class="ant-pro-pages-account-projects-cardList">
    <a-list :loading="loading" :data-source="data" :grid="{ gutter: 24, xxl: 3, xl: 2, lg: 2, md: 2, sm: 2, xs: 1 }">
      <a-list-item slot="renderItem" slot-scope="item">
        <a :href="item.url" target="_blank">
          <a-card class="ant-pro-pages-account-projects-card" hoverable>
            <img slot="cover" :src="item.cover" :alt="item.title" />
            <a-card-meta :title="item.title">
              <template slot="description">
                <ellipsis :length="50">{{ item.description }}</ellipsis>
              </template>
            </a-card-meta>
            <div class="cardItemContent">
              <span>{{ item.updatedAt | fromNow }}</span>
            </div>
          </a-card>
        </a>
      </a-list-item>
    </a-list>
  </div>
</template>

<script>
import moment from 'moment'
import { TagSelect, StandardFormRow, Ellipsis, AvatarList } from '@/components'
const TagSelectOption = TagSelect.Option
const AvatarListItem = AvatarList.AvatarItem

export default {
  name: 'Project',
  components: {
    AvatarList,
    AvatarListItem,
    Ellipsis,
    TagSelect,
    TagSelectOption,
    StandardFormRow
  },
  data () {
    return {
      data: [
        {
          title: 'Ant Design Vue',
          cover: 'https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png',
          url: 'ttps://github.com/vueComponent/ant-design-vue',
          description: 'An enterprise-class UI components based on Ant Design and Vue. '
        },
        {
          title: 'Antd Pro Vue',
          cover: 'https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png',
          url: 'https://github.com/vueComponent/ant-design-vue-pro',
          description: 'An out-of-box UI solution for enterprise applications as a Vue boilerplate. based on Ant Design of Vue.'
        },
        {
          title: 'vditor',
          cover: 'https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png',
          url: 'https://github.com/Vanessa219/vditor',
          description: '一款浏览器端的 Markdown 编辑'
        }
      ],
      form: this.$form.createForm(this),
      loading: false
    }
  },
  filters: {
    fromNow (date) {
      return moment(date).fromNow()
    }
  },
  mounted () {
    this.getList()
  },
  methods: {
    handleChange (value) {
      console.log(`selected ${value}`)
    },
    getList () {
    }
  }
}
</script>

<style lang="less" scoped>
  .ant-pro-pages-account-projects-cardList {
    margin-top: 24px;

    /deep/ .ant-card-meta-title {
      margin-bottom: 4px;
    }

    /deep/ .ant-card-meta-description {
      height: 44px;
      overflow: hidden;
      line-height: 22px;
    }

    .cardItemContent {
      display: flex;
      height: 20px;
      margin-top: 16px;
      margin-bottom: -4px;
      line-height: 20px;

      > span {
        flex: 1 1;
        color: rgba(0,0,0,.45);
        font-size: 12px;
      }

      /deep/ .ant-pro-avatar-list {
        flex: 0 1 auto;
      }
    }
  }
</style>
