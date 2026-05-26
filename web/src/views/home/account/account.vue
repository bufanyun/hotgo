<template>
  <div>
    <n-grid cols="24 300:1 600:24" :x-gap="12">
      <n-grid-item span="6">
        <n-card :bordered="false" class="proCard">
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
          <SafetySetting v-if="type === 2" />
          <CashSetting v-if="type === 3" />
          <ThirdBind v-if="type === 4" />
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>
<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import BasicSetting from './BasicSetting.vue';
  import SafetySetting from './SafetySetting.vue';
  import CashSetting from './CashSetting.vue';
  import ThirdBind from './ThirdBind.vue';
  import { useRouter } from 'vue-router';
  import {pushHashRouterParameter} from "@/utils/urlUtils";

  const router = useRouter();
  const type = ref(1);
  const typeTitle = ref('基本设置');

  const typeTabList = [
    {
      name: '基本设置',
      desc: '个人账户信息设置',
      key: 1,
    },
    {
      name: '安全设置',
      desc: '密码、手机号、邮箱等设置',
      key: 2,
    },
    {
      name: '提现设置',
      desc: '提现收款账号支付宝设置',
      key: 3,
    },
    {
      name: '第三方绑定',
      desc: '第三方快捷登录、消息推送',
      key: 4,
    },
  ];

  onMounted(() => {
    if (router.currentRoute.value.query?.type) {
      setDefaultOption();
    }
  });

  function setDefaultOption() {
    const key = router.currentRoute.value.query.type as unknown as number;
    if (key !== undefined && key > 0) {
      for (const item of typeTabList) {
        if (item.key == key) {
          switchType(item);
        }
      }
    }
  }

  function switchType(e) {
    type.value = e.key;
    typeTitle.value = e.name;
    pushHashRouterParameter(window.location.href, 'type', e.key);
  }
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
