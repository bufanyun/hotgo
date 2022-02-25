<template>
  <a-drawer
    width="45%"
    :title="formTitle"
    :label-col="4"
    :wrapper-col="14"
    :visible="open"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    @close="cancel">
    <a-form-model ref="form" :model="form" :rules="rules" layout="vertical">
      <a-form-model-item label="参数编码" prop="configKey">
        <a-input v-model="form.configKey" placeholder="请输入" />
      </a-form-model-item>
      <a-form-model-item label="参数名称" prop="configName">
        <a-input v-model="form.configName" placeholder="请输入" />
      </a-form-model-item>
      <a-form-model-item label="参数值" prop="configValue">
        <a-input v-model="form.configValue" placeholder="请输入" />
      </a-form-model-item>
      <a-form-model-item label="是否内置参数" prop="configType">
        <a-radio-group v-model="form.configType" button-style="solid">
          <a-radio-button v-for="(d, index) in typeOptions" :key="index" :value="d.dictValue">{{
            d.dictLabel
          }}</a-radio-button>
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
          <a-button @click="cancel">
            取消
          </a-button>
        </a-space>
      </div>
    </a-form-model>
  </a-drawer>
</template>
<script>
import ConfigForm from './ConfigForm'
export default {
  ...ConfigForm
}
</script>
