<template>
  <div class="module" ref="portaletDiv">
    <div class="module-in">
      <ul class="information02">
        <li>
          <div class="sky">
            <!-- 此为图表盒子 -->
            <img src="@/views/demo/portLetCards/images-m/img01.png" alt="" />
          </div>
          <div class="information02-data">
            <span>50%</span>
            <p>压力仪表盘</p>
          </div>
        </li>
        <li>
          <div class="sky">
            <!-- 此为图表盒子 -->
            <img src="@/views/demo/portLetCards/images-m/img02.png" alt="" />
          </div>
          <div class="information02-data">
            <span>28%</span>
            <p>完成率</p>
          </div>
        </li>
      </ul>
    </div>
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
