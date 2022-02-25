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
          <a-form-model-item label="岗位编码" prop="code">
            <a-input v-model="form.code" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="岗位名称" prop="name">
            <a-input v-model="form.name" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="排序" prop="sort">
            <a-input-number v-model="form.sort" :min="0" style="width: 100%"/>
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="状态" prop="status">
            <a-radio-group v-model="form.status" button-style="solid">
              <a-radio-button v-for="(d, index) in statusOptions" :key="index" :value="d.value" >{{ d.label }}</a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :lg="24" :md="24" :sm="24">
          <a-form-model-item label="备注" prop="remark">
            <a-input v-model="form.remark" placeholder="请输入" type="textarea" allow-clear :rows="2" />
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
import PostForm from './PostForm'
export default {
  ...PostForm
}
</script>
