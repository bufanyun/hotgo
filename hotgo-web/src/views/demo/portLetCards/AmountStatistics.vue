<template>
  <div
    class="module module01"
    ref="portaletDiv">
    <div class="card">
      <h3>项目金额总览</h3>
      <div class="card-inner">
        <span class="card-number">150032</span>
        <span class="card-tag card-blue">+15%</span>
      </div>
    </div>
    <a-icon type="transaction" class="card-icon" />
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
