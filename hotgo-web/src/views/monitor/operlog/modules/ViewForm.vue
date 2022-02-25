<template>
  <a-drawer width="750px" placement="right" :closable="false" :visible="openView" @close="onCloseView">
    <a-descriptions title="操作信息" layout="vertical" :column="3" style="word-break: break-all;word-wrap: break-word;" bordered>
      <a-descriptions-item label="操作模块">
        {{ form.module }}
      </a-descriptions-item>
      <a-descriptions-item label="登录信息">
        #{{ form.member_id }} / {{ form.member_name }} / {{ form.ip }} / {{ form.region }}
      </a-descriptions-item>
      <a-descriptions-item label="请求方式">
        {{ form.method }}
      </a-descriptions-item>
      <a-descriptions-item label="请求地址">
        {{ form.url }}
      </a-descriptions-item>
      <a-descriptions-item label="操作方法" :span="2">
        <div style="word-break: break-all;">{{ form.method }}</div>
      </a-descriptions-item>
    </a-descriptions>
    <a-divider/>
    <a-descriptions title="接口信息" layout="vertical" :column="3" style="word-break: break-all;word-wrap: break-word;" bordered>
      <a-descriptions-item label="请求头部" :span="3">
        {{ form.header_data }}
      </a-descriptions-item>
      <a-descriptions-item label="GET参数" :span="3">
        {{ form.get_data }}
      </a-descriptions-item>
      <a-descriptions-item label="POST参数" :span="3">
        {{ form.post_data }}
      </a-descriptions-item>
      <a-descriptions-item label="操作状态">
        <a-badge v-if="form.error_code === 0" status="processing" text="正常" />
        <a-badge v-if="form.error_code !== 0" status="error" text="失败" />
      </a-descriptions-item>
      <a-descriptions-item label="操作时间" >
        {{ parseTime(form.created_at) }}
      </a-descriptions-item>
      <a-descriptions-item label="耗时" >
        {{ form.take_up_time }}毫秒
      </a-descriptions-item>
      <a-descriptions-item label="回复信息">
        {{ form.error_msg }}
      </a-descriptions-item>
      <a-descriptions-item label="报错日志" v-if="form.error_code !== 0">
        {{ form.error_data }}
      </a-descriptions-item>
    </a-descriptions>

    <a-divider v-if="form.logContent !== null"/>
    <a-descriptions
      v-if="form.logContent !== null"
      title="变更记录"
      layout="vertical"
      :column="3"
      style="word-break: break-all;word-wrap: break-word;"
      bordered>
      <a-descriptions-item label="请求ID" :span="3">
        {{ form.req_id }}
      </a-descriptions-item>
      <a-descriptions-item label="变更情况" :span="3">
        无
      </a-descriptions-item>
    </a-descriptions>

  </a-drawer>
</template>

<script>

export default {
  name: 'ViewForm',
  props: {
  },
  data () {
    return {
      // 表单参数
      form: {},
      openView: false
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
  methods: {
    handleView (row) {
      this.openView = true
      this.form = row
    },
    onCloseView () {
      this.openView = false
    }
  }
}
</script>
