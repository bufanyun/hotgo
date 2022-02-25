<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    modalWidth="640"
    @cancel="cancel"
    modalHeight="420">
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row class="form-row" :gutter="32">
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="字典名称" prop="name">
            <a-input v-model="form.name" placeholder="请输入字典名称" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="字典类型" prop="type">
            <a-input v-model="form.type" placeholder="请输入字典类型" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="状态" prop="status">
            <a-select placeholder="请选择" v-model="form.status" allow-clear>
              <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{ d.label }}</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
        <a-col :lg="24" :md="24" :sm="24" :span="24">
          <a-form-model-item label="备注" prop="remark">
            <a-input v-model="form.remark" placeholder="请输入内容" type="textarea" allow-clear />
          </a-form-model-item>
        </a-col>
      </a-row>
    </a-form-model>
    <template slot="footer">
      <a-button @click="cancel">
        取消
      </a-button>
      <a-button type="primary" @click="submitForm">
        保存
      </a-button>
    </template>
  </ant-modal>
</template>

<script>
  import DictTypeForm from './DictTypeForm'
  export default {
    ...DictTypeForm
  }
</script>
