<template>
  <a-drawer
    width="750px"
    :label-col="4"
    :title="formTitle"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    :wrapper-col="14"
    :visible="open"
    @close="onClose">
    <a-form-model ref="form" :model="form" :rules="rules">
      <a-form-model-item label="任务名称" prop="jobName">
        <a-input v-model="form.jobName" placeholder="请输入任务名称" />
      </a-form-model-item>
      <a-form-model-item label="任务分组" prop="jobGroup">
        <a-select placeholder="请选择" v-model="form.jobGroup">
          <a-select-option v-for="(d, index) in jobGroupOptions" :key="index" :value="d.dictValue" >{{ d.dictLabel }}</a-select-option>
        </a-select>
      </a-form-model-item>
      <a-form-model-item prop="invokeTarget">
        <span slot="label">
          调用方法&nbsp;
          <a-popover placement="topLeft">
            <template slot="content">
              <p>Class类调用示例：<code>com.aidex.quartz.task.RyTask.ryParams('ry')</code></p>
              <p>参数说明：支持字符串，布尔类型，长整型，浮点型，整型</p>
            </template>
            <span slot="title"> Bean调用示例：<code>ryTask.ryParams('ry')</code></span>
            <a-icon type="question-circle-o" />
          </a-popover>
        </span>
        <a-input v-model="form.invokeTarget" placeholder="请输入调用目标字符串" />
      </a-form-model-item>
      <a-form-model-item label="cron表达式" prop="cronExpression">
        <a-input v-model="form.cronExpression" placeholder="请输入cron执行表达式" />
      </a-form-model-item>
      <a-form-model-item label="是否并发" prop="concurrent">
        <a-radio-group v-model="form.concurrent" button-style="solid">
          <a-radio-button value="0">允许</a-radio-button>
          <a-radio-button value="1">禁止</a-radio-button>
        </a-radio-group>
      </a-form-model-item>
      <a-form-model-item label="错误策略" prop="misfirePolicy">
        <a-radio-group v-model="form.misfirePolicy" button-style="solid">
          <a-radio-button value="1">立即执行</a-radio-button>
          <a-radio-button value="2">执行一次</a-radio-button>
          <a-radio-button value="3">放弃执行</a-radio-button>
        </a-radio-group>
      </a-form-model-item>
      <a-form-model-item label="状态" prop="status">
        <a-radio-group v-model="form.status" button-style="solid">
          <a-radio v-for="(d, index) in statusOptions" :key="index" :value="d.dictValue">{{ d.dictLabel }}</a-radio>
        </a-radio-group>
      </a-form-model-item>
      <a-form-model-item label="备注" prop="remark">
        <a-input v-model="form.remark" placeholder="请输入备注" type="textarea" allow-clear />
      </a-form-model-item>
      <div class="bottom-control">
        <a-space>
          <a-button type="primary" @click="submitForm">
            保存
          </a-button>
          <a-button type="dashed" @click="cancel">
            取消
          </a-button>
        </a-space>
      </div>
    </a-form-model>
  </a-drawer>
</template>

<script>

import { getJob, addJob, updateJob } from '@/api/monitor/job'

export default {
  name: 'CreateForm',
  props: {
    statusOptions: {
      type: Array,
      required: true
    },
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
      open: false,
      openView: false,
      rules: {
        jobName: [{ required: true, message: '任务名称不能为空', trigger: 'blur' }],
        invokeTarget: [{ required: true, message: '调用目标字符串不能为空', trigger: 'blur' }],
        cronExpression: [{ required: true, message: 'cron执行表达式不能为空', trigger: 'blur' }]
      }
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
      getJob(row.jobId).then(response => {
        this.form = response.data
        this.openView = true
      })
    },
    onCloseView () {
      this.openView = false
    },
    onClose () {
      this.open = false
    },
    // 取消按钮
    cancel () {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset () {
      this.form = {
        jobId: undefined,
        jobName: undefined,
        jobGroup: undefined,
        invokeTarget: undefined,
        cronExpression: undefined,
        misfirePolicy: '1',
        concurrent: '1',
        status: '0'
      }
    },
     /** 新增按钮操作 */
    handleAdd () {
      this.reset()
      this.open = true
      this.formTitle = '添加任务'
    },
    /** 修改按钮操作 */
    handleUpdate (row, ids) {
      this.reset()
      const jobId = row ? row.jobId : ids
      getJob(jobId).then(response => {
        this.form = response.data
        this.open = true
        this.formTitle = '修改任务'
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs.form.validate(valid => {
        if (valid) {
          if (this.form.jobId !== undefined) {
            updateJob(this.form).then(response => {
              this.$message.success(
                '修改成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          } else {
            addJob(this.form).then(response => {
              this.$message.success(
                '新增成功',
                3
              )
              this.open = false
              this.$emit('ok')
            })
          }
        } else {
          return false
        }
      })
    }
  }
}
</script>
