<template>
  <a-result title="该功能尚未开发!">
    <template #icon>
      <a-icon type="smile" theme="twoTone" />
    </template>
    <template #extra>
      <a-button type="primary" @click="returnIndex">
        返回工作台
      </a-button>
    </template>
  </a-result>
</template>

<script>

export default {
  name: 'Build',
  components: {
  },
  data () {
    return {
    }
  },
  filters: {
  },
  created () {
  },
  computed: {
  },
  watch: {
  },
  mounted: function () {
  },
  methods: {
    returnIndex () {
      this.$router.push({ path: '/index' })
    }
  }
}
</script>
