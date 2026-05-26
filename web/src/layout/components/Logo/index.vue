<template>
  <div class="logo">
    <img src="~@/assets/images/logo.png" alt="" :class="{ 'mr-2': !collapsed }" />
    <h2 v-show="!collapsed" class="title">{{ projectName }}</h2>
  </div>
</template>

<script>
  import { useUserStore } from '@/store/modules/user';

  export default {
    name: 'Index',
    props: {
      collapsed: {
        type: Boolean,
      },
    },
    setup() {
      const userStore = useUserStore();
      const projectName = userStore.loginConfig?.projectName;
      return {
        projectName,
      };
    },
  };
</script>

<style lang="less" scoped>
  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 64px;
    line-height: 64px;
    overflow: hidden;
    white-space: nowrap;

    img {
      width: auto;
      height: 30px;
      border-radius: 20px;
    }

    .title {
      margin-bottom: 0px;
    }
  }
</style>
