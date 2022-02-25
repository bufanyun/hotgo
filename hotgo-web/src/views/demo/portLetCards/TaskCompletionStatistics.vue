<template>
  <div class="module module01" ref="portaletDiv">
    <div class="card">
      <h3>当月任务完成情况总览</h3>
      <div class="card-inner">
        <!--        <span class="card-number">72</span>-->
        <div class="card-progress">
          <a-progress :percent="72" :stroke-width="12" status="active" />
        </div>
      </div>
    </div>
    <a-icon type="pie-chart" class="card-icon" />
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
