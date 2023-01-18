<template>
  <div class="logo">
    <img src="~@/assets/images/logo.png" alt="" :class="{ 'mr-2': !collapsed }" />
    <h2 v-show="!collapsed" class="title">HG后台管理系统</h2>
  </div>
</template>

<script>
  export default {
    name: 'Index',
    props: {
      collapsed: {
        type: Boolean,
      },
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
