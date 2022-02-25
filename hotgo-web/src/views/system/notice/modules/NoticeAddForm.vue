<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    modalWidth="800"
    @cancel="cancel"
    modalHeight="550">
    <a-form-model
      ref="form"
      :model="form"
      :rules="rules"
      slot="content"
      layout="vertical">
      <a-row class="form-row" :gutter="32">
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="公告标题" prop="title">
            <a-input v-model="form.title" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="公告类型" prop="type">
            <a-select placeholder="请选择" v-model="form.type">
              <a-select-option v-for="(d, index) in typeOptions" :key="index" :value="d.value" >{{ d.label }}</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
        <a-col :lg="24" :md="24" :sm="24">
          <a-form-model-item label="内容" prop="content" class="addModel">
            <div id="vditor" name="description"></div>
          </a-form-model-item>
        </a-col>
      </a-row>
    </a-form-model>
    <template slot="footer">
      <a-button @click="cancel" :disabled="uploaderButtonStatus">
        取消
      </a-button>
      <a-button type="primary" @click="submitForm" :loading="uploaderButtonStatus">
        发布
      </a-button>
    </template>
  </ant-modal>
</template>
<script>
import NoticeForm from './NoticeForm'
export default {
  ...NoticeForm
}
</script>
<style lang="less">
  .addModel{
    .vditor-toolbar--pin{
      top:auto!important;
    }
  }
</style>
