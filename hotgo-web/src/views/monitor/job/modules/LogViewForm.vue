<template>
  <a-drawer width="750px" placement="right" :closable="false" :visible="openView" @close="onCloseView">
    <a-descriptions title="调度日志详细" layout="vertical" bordered :column="3">
      <a-descriptions-item label="日志序号">
        {{ form.jobLogId }}
      </a-descriptions-item>
      <a-descriptions-item label="任务名称">
        {{ form.jobName }}
      </a-descriptions-item>
      <a-descriptions-item label="任务分组">
        {{ jobGroupFormat(form) }}
      </a-descriptions-item>
      <a-descriptions-item label="执行状态">
        <a-badge v-if="form.status == 0" status="processing" text="正常" />
        <a-badge v-if="form.status == 1" status="warning" text="失败" />
      </a-descriptions-item>
      <a-descriptions-item label="执行时间" span="2">
        {{ form.createTime }}
      </a-descriptions-item>
      <a-descriptions-item label="调用方法" span="3">
        {{ form.invokeTarget }}
      </a-descriptions-item>
      <a-descriptions-item label="日志信息" span="3">
        {{ form.jobMessage }}
      </a-descriptions-item>
      <a-descriptions-item label="异常信息" span="3" v-if="form.status == 1">
        {{ form.exceptionInfo }}
      </a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script>

export default {
  name: 'LogViewForm',
  props: {
    jobGroupOptions: {
      type: Array,
      required: true
    }
  },
  data () {
    return {
      loading: false,
      // 表单参数
      form: {
        jobId: undefined,
        jobName: undefined,
        jobGroup: undefined,
        invokeTarget: undefined,
        cronExpression: undefined,
        misfirePolicy: '1',
        concurrent: '1',
        status: '0'
      },
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
    jobGroupFormat (row) {
      return this.selectDictLabel(this.jobGroupOptions, row.jobGroup)
    },
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
