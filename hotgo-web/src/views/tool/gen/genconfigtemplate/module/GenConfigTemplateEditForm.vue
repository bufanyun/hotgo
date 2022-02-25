<template>
  <a-drawer
    width="35%"
    :label-col="4"
    :wrapper-col="14"
    :visible="open"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    @close="onClose"
    :title="formTitle">
    <a-form-model ref="form" :model="form" :rules="rules" layout="vertical">
      <a-spin :spinning="spinning" :delay="delayTime" tip="Loading...">
        <a-row :gutter="32">
          <a-col :span="12">
            <a-form-model-item label="模板名称" prop="templateName">
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
            <a-form-model-item label="工作空间路径" prop="workspacePath">
              <a-input v-model="form.workspacePath" placeholder="请输入工作空间路径" />
            </a-form-model-item>
          </a-col>
          <a-col :span="12">
            <a-form-model-item label="前端工作空间" prop="webWorkspacePath">
              <a-input v-model="form.webWorkspacePath" placeholder="请输入前端工作空间" />
            </a-form-model-item>
          </a-col>
          <a-col :span="12">
            <a-form-model-item label="模块名" prop="moduleName">
              <a-input v-model="form.moduleName" placeholder="请输入模块名" />
            </a-form-model-item>
          </a-col>
          <a-col :span="12">
            <a-form-model-item label="包路径" prop="packageName">
              <a-input v-model="form.packageName" placeholder="请输入包路径" />
            </a-form-model-item>
          </a-col>
          <a-col :span="12">
            <a-form-model-item label="排序" prop="sort">
              <a-input-number v-model="form.sort" placeholder="请输入排序" :min="0" style="width: 100%" />
            </a-form-model-item>
          </a-col>
          <a-col :span="12">
            <a-form-model-item label="是否默认" prop="templateDefault">
              <a-radio-group v-model="form.templateDefault" button-style="solid">
                <a-radio-button v-for="(d, index) in templateDefaultOptions" :key="index" :value="d.dictValue">
                  {{ d.dictLabel }}</a-radio-button>
              </a-radio-group>
            </a-form-model-item>
          </a-col>
          <a-col :span="24">
            <a-form-model-item label="备注" prop="remark">
              <a-input v-model="form.remark" placeholder="请输入" type="textarea" :rows="4" allow-clear />
            </a-form-model-item>
          </a-col>
        </a-row>
      </a-spin>
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
  import NoticeForm from './GenConfigTemplateForm'
  export default {
    ...NoticeForm
  }
</script>
<style lang="less" scoped>
  #vditor {
    margin: 0 5% 5% 5%;
  }
</style>
