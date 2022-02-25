<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    modalWidth="640"
    @cancel="cancel"
    modalHeight="420"
  >
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row class="form-row" :gutter="32">
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="参数编码" prop="configKey">
            <a-input v-model="form.configKey" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="参数名称" prop="configName">
            <a-input v-model="form.configName" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="参数值" prop="configValue">
            <a-input v-model="form.configValue" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="是否内置参数" prop="configType">
            <a-radio-group v-model="form.configType" button-style="solid">
              <a-radio-button v-for="(d, index) in typeOptions" :key="index" :value="d.dictValue">{{
                d.dictLabel
              }}</a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :lg="24" :md="24" :sm="24" :span="24">
          <a-form-model-item label="备注" prop="remark">
            <a-input v-model="form.remark" placeholder="请输入备注" type="textarea" allow-clear />
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
import ConfigForm from './ConfigForm'
export default {
  ...ConfigForm
}
</script>
