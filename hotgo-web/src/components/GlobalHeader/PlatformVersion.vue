<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    :isShowTitle="false"
    :closeAble="true"
    :footer="null"
    modalWidth="600"
    modalHeight="350"
    @cancel="cancel"
  >

    <a-row slot="content">
      <a-col :span="8">
        <div class="copyright-icon"><a-icon type="key" /></div>
      </a-col>
      <a-col :span="16">
        <div class="copyright-content">
          <div class="copyright-text">
            <h2>平台授权信息.</h2>
            <h3>非常感谢您对我们产品的认可与支持！</h3>
            {{ versionContent[0] }}<br>
            {{ versionContent[1] }}<br>
            {{ versionContent[2] }}<br>
            授权产品名称：AiDex<br>
            当前平台版本：V1.2.1
          </div>
          <a-button type="primary" icon="close-circle" @click="cancel">
            关闭页面
          </a-button>
        </div>
      </a-col>
    </a-row>
  </ant-modal>
</template>
<script>
    import AntModal from '@/components/pt/dialog/AntModal'
    import { mapGetters } from 'vuex'
    export default {
      name: 'CreateForm',
      components: {
        AntModal
      },
      data () {
        return {
          loading: false,
          formTitle: '',
          open: false,
          versionContent: []
        }
      },
      filters: {
      },
      created () {
      },
      computed: {
        ...mapGetters([
          'platformVersion'
        ])
      },
      watch: {
      },
      methods: {
        cancel () {
          this.open = false
          this.$emit('close')
        },
        showVersion () {
          this.open = true
          this.formTitle = '授权信息'
          if (this.platformVersion !== null && this.platformVersion !== '') {
            const licenseInfo = JSON.parse(this.platformVersion)
            const customName = licenseInfo.customName
            const versionDes = licenseInfo.versionDes
            const version = licenseInfo.version
            let deadLine = licenseInfo.deadLine
            if (version === '2') {
              const beforeYear = deadLine.split('-')[0]
              let myDate = new Date()
              myDate = myDate.getFullYear()
              if ((beforeYear - myDate) >= 10) {
                deadLine = '无限制'
              }
            }
             this.versionContent.push('授权对象：' + customName)
             this.versionContent.push('版本信息：' + versionDes)
             this.versionContent.push('到期时间：' + deadLine)
          } else {
            this.versionContent.push('授权对象：未知')
            this.versionContent.push('版本信息：未知')
            this.versionContent.push('到期时间：未知')
          }
        }
      }
    }
</script>
<style lang="less">
  .copyright-content{
    padding: 30px 10px 20px;
  }
  .copyright-icon{
    text-align: center;
    font-size: 80px;
    padding: 20px 30px 10px;
    color: #85c1fb;
  }
  .copyright-text{
    margin-bottom: 20px;
    h2{
      font-size:22px;
      color: #333333;
    }
  }
</style>
