<template>
  <div class="loading-body">
    <div class="spin-content">
      <a-spin />
    </div>
  </div>
</template>
<script>
import { Spin } from 'ant-design-vue'
export default {
  name: 'AsyncComponentLoading',
  components: {
    ASpin: Spin
  }
}
</script>
<style lang="less" scoped>
.loading-body {
  width: 100%;
  height: 100%;
  position: relative;

  .spin-content {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    width: 30px;
    height: 30px;
    margin: auto;
  }
}
</style>
