<template>
  <div>
    <n-grid cols="24 300:1 600:24" :x-gap="24">
      <n-grid-item span="6">
        <n-card :bordered="false" size="small" class="proCard">
          <n-thing
            class="thing-cell"
            v-for="item in typeTabList"
            :key="item.key"
            :class="{ 'thing-cell-on': type === item.key }"
            @click="switchType(item)"
          >
            <template #header>{{ item.name }}</template>
            <template #description>{{ item.desc }}</template>
          </n-thing>
        </n-card>
      </n-grid-item>
      <n-grid-item span="18">
        <n-card :bordered="false" size="small" :title="typeTitle" class="proCard">
          <BasicSetting v-if="type === 1" />
          <ThemeSetting v-if="type === 2" />
          <RevealSetting v-if="type === 3" />
          <EmailSetting v-if="type === 4" />
          <SmsSetting v-if="type === 5" />
          <CashSetting v-if="type === 7" />
          <UploadSetting v-if="type === 8" />
          <GeoSetting v-if="type === 9" />
          <PaySetting v-if="type === 10" />
          <WechatSetting v-if="type === 11" />
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>
<script lang="ts">
  import { defineComponent, reactive, toRefs } from 'vue';
  import BasicSetting from './BasicSetting.vue';
  import RevealSetting from './RevealSetting.vue';
  import EmailSetting from './EmailSetting.vue';
  import ThemeSetting from './ThemeSetting.vue';
  import CashSetting from './CashSetting.vue';
  import UploadSetting from './UploadSetting.vue';
  import GeoSetting from './GeoSetting.vue';
  import SmsSetting from './SmsSetting.vue';
  import PaySetting from './PaySetting.vue';
  import WechatSetting from './WechatSetting.vue';
  const typeTabList = [
    {
      name: '基本设置',
      desc: '系统常规设置',
      key: 1,
    },
    // {
    //   name: '主题设置',
    //   desc: '系统主题设置',
    //   key: 2,
    // },
    // {
    //   name: '显示设置',
    //   desc: '系统显示设置',
    //   key: 3,
    // },
    {
      name: '邮件设置',
      desc: '系统邮件设置',
      key: 4,
    },
    {
      name: '短信配置',
      desc: '短信验证码平台',
      key: 5,
    },
    // {
    //   name: '管理员配置',
    //   desc: '默认设置和权限屏蔽',
    //   key: 6,
    // },
    {
      name: '提现配置',
      desc: '管理员提现规则配置',
      key: 7,
    },
    {
      name: '云存储',
      desc: '配置上传文件驱动',
      key: 8,
    },
    {
      name: '地理位置',
      desc: '配置地理位置工具',
      key: 9,
    },
    {
      name: '支付配置',
      desc: '支付宝/微信/QQ支付配置等',
      key: 10,
    },
    {
      name: '微信配置',
      desc: '公众号/开放平台/小程序配置等',
      key: 11,
    },
  ];
  export default defineComponent({
    components: {
      BasicSetting,
      RevealSetting,
      EmailSetting,
      ThemeSetting,
      CashSetting,
      UploadSetting,
      GeoSetting,
      SmsSetting,
      PaySetting,
      WechatSetting,
    },
    setup() {
      const state = reactive({
        type: 1,
        typeTitle: '基本设置',
      });

      function switchType(e) {
        state.type = e.key;
        state.typeTitle = e.name;
      }

      return {
        ...toRefs(state),
        switchType,
        typeTabList,
      };
    },
  });
</script>
<style lang="less" scoped>
  .thing-cell {
    margin: 0 -16px 10px;
    padding: 5px 16px;

    &:hover {
      background: #f3f3f3;
      cursor: pointer;
    }
  }

  .thing-cell-on {
    background: #f0faff;
    color: #2d8cf0;

    ::v-deep(.n-thing-main .n-thing-header .n-thing-header__title) {
      color: #2d8cf0;
    }

    &:hover {
      background: #f0faff;
    }
  }
</style>
