<template>
  <ant-modal
    modalWidth="800"
    modalHeight="550"
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    @cancel="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row :gutter="32">
        <a-col :span="12">
          <a-form-model-item label="名称" prop="name">
            <a-input v-model="form.name" placeholder="请输入名称"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="编码" prop="code">
            <a-input v-model="form.code" placeholder="请输入编码"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12" >
          <a-form-model-item label="应用范围" prop="applicationRange">
            <a-radio-group @change="rangeChange" v-model="form.applicationRange" button-style="solid">
              <a-radio-button
                v-for="(dict, index) in applicationRangeOptions"
                :key="index"
                :value="dict.dictValue"
              >
                {{ dict.dictLabel }}
              </a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="12" >
          <a-form-model-item label="是否默认" prop="isDefault">
            <a-radio-group v-model="form.isDefault" button-style="solid">
              <a-radio-button
                v-for="(dict, index) in isDefaultOptions"
                :key="index"
                :value="dict.dictValue"
              >
                {{ dict.dictLabel }}
              </a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="12" v-if="isShowResourceId">
          <a-form-model-item label="资源" prop="resourceId">
            <a-select v-if="form.applicationRange === 'R' " v-model="form.resourceId" placeholder="请选择" option-filter-prop="children">
              <a-select-option v-for="(d, index) in roleOptions" :key="index" :value="d.id">
                {{ d.roleName }}
              </a-select-option>
            </a-select>
            <select-user v-if="form.applicationRange === 'U' " v-model="form.resourceId"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="24">
          <a-form-model-item label="备注" prop="remark">
            <a-input v-model="form.remark" placeholder="请输入" type="textarea" allow-clear />
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
import SysPortalConfigAddForm from './SysPortalConfigForm'
export default {
  ...SysPortalConfigAddForm
}
</script>
