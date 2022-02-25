<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    modalWidth="800"
    @cancel="cancel"
    modalHeight="550">
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row :gutter="32">
        <a-col :span="12">
          <a-form-model-item label="模板名称" prop="templateName">
            <span slot="label">
              生成模块名
              <a-tooltip>
                <template slot="title">
                  可理解为子系统名，例如 system
                </template>
                <a-icon type="question-circle-o" />
              </a-tooltip>
            </span>
            <a-input v-model="form.templateName" placeholder="请输入模板名称" />
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row :gutter="32">
        <a-col :span="12">
          <a-form-model-item label="作者" prop="functionAuthor">
            <a-input v-model="form.functionAuthor" placeholder="请输入作者" />
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="邮箱" prop="functionAuthorEmail">
            <a-input v-model="form.functionAuthorEmail" placeholder="请输入邮箱" />
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="后端工作空间" prop="workspacePath">
            <a-input v-model="form.workspacePath" placeholder="示例：C:\cloud\AiDex-Sharp\ruoyi-system" />
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="前端工作空间" prop="webWorkspacePath">
            <a-input v-model="form.webWorkspacePath" placeholder="示例：C:\cloud\AidexSharp" />
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="模块名" prop="moduleName">
            <a-input v-model="form.moduleName" placeholder="示例：system" />
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="包路径" prop="packageName">
            <a-input v-model="form.packageName" placeholder="示例：com.aidex.system" />
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="排序" prop="sort">
            <a-input-number v-model="form.sort" placeholder="请输入排序" :min="0" style="width: 100%"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="是否默认" prop="templateDefault">
            <a-radio-group v-model="form.templateDefault" button-style="solid">
              <a-radio-button v-for="(d, index) in templateDefaultOptions" :key="index" :value="d.dictValue" >{{ d.dictLabel }}</a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="24">
          <a-form-model-item label="备注" prop="remark">
            <a-input v-model="form.remark" placeholder="请输入" type="textarea" :rows="4" allow-clear />
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
import GenConfigTemplateAddForm from './GenConfigTemplateForm'
export default {
  ...GenConfigTemplateAddForm
}
</script>
