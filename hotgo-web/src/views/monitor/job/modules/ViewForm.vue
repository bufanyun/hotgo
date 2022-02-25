<template>
  <a-drawer width="750px" placement="right" :closable="false" :visible="openView" @close="onCloseView">
    <a-descriptions title="任务详细" layout="vertical" bordered>
      <a-descriptions-item label="任务编号">
        {{ form.jobId }}
      </a-descriptions-item>
      <a-descriptions-item label="任务分组">
        {{ jobGroupFormat(form.jobGroup) }}
      </a-descriptions-item>
      <a-descriptions-item label="任务名称">
        {{ form.jobName }}
      </a-descriptions-item>
      <a-descriptions-item label="创建时间">
        {{ form.createTime }}
      </a-descriptions-item>
      <a-descriptions-item label="下次执行时间" span="2">
        {{ parseTime(form.nextValidTime) }}
      </a-descriptions-item>
      <a-descriptions-item label="调用目标方法">
        {{ form.invokeTarget }}
      </a-descriptions-item>
      <a-descriptions-item label="cron表达式" span="2">
        {{ form.cronExpression }}
      </a-descriptions-item>
      <a-descriptions-item label="任务状态">
        <a-badge v-if="form.status == 0" status="processing" text="运行中" />
        <a-badge v-if="form.status == 1" status="warning" text="暂停" />
      </a-descriptions-item>
      <a-descriptions-item label="是否并发">
        <a-badge v-if="form.concurrent == 0" status="processing" text="允许" />
        <a-badge v-if="form.concurrent == 1" status="warning" text="禁止" />
      </a-descriptions-item>
      <a-descriptions-item label="执行策略">
        <a-badge v-if="form.misfirePolicy == 0" status="default" text="默认策略" />
        <a-badge v-if="form.misfirePolicy == 1" status="Processing" text="立即执行" />
        <a-badge v-if="form.misfirePolicy == 2" color="purple" text="执行一次" />
        <a-badge v-if="form.misfirePolicy == 3" status="warning" text="放弃执行" />
      </a-descriptions-item>
    </a-descriptions>
  </a-drawer>
</template>

<script>

import { getJob } from '@/api/monitor/job'

export default {
  name: 'ViewForm',
  props: {
    jobGroupOptions: {
      type: Array,
      required: true
    }
  },
  data () {
    return {
      loading: false,
      formTitle: '',
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
    jobGroupFormat (jobGroup) {
      return this.selectDictLabel(this.jobGroupOptions, jobGroup)
    },
    handleView (row) {
      getJob(row.jobId).then(response => {
        this.form = response.data
        this.openView = true
      })
    },
    onCloseView () {
      this.openView = false
    }
  }
}
</script>
