<template>
  <div class="module module01" ref="portaletDiv">
    <div class="card">
      <h3>访问量</h3>
      <div class="card-inner">
        <span class="card-number">8,846</span>
        <div class="card-chart">
          <div class="chart_sm">
            <img
              src="@/views/demo/portLetCards/images-m/chart_sm.png"
              width="100%"
            />
          </div>
        </div>
      </div>
    </div>
    <a-icon type="dashboard" class="card-icon" />
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
