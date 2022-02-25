<template>
  <div class="module module01" ref="portaletDiv">
    <div class="card">
      <h3>当月任务总数统计</h3>
      <div class="card-inner">
        <span class="card-number">150</span>
        <span class="card-tag card-green">-15%</span>
      </div>
    </div>
    <a-icon type="history" class="card-icon" />
  </div>
</template>

<script>
  export default {
    created () {
    },
    mounted () {
      this.$emit('setHeight', this.$refs.portaletDiv.offsetHeight)
    },
    watch: {
    },
    components: {
    },
    methods: {
    }
  }
</script>

<style lang="less" scoped>
@import './style/base.less';
</style>
