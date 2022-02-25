<template>
  <a-drawer
    width="650"
    :title="formTitle"
    :label-col="4"
    :wrapper-col="14"
    :visible="open"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    @close="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules" layout="vertical">
      <a-form-model-item label="公告标题" prop="title">
        <a-input v-model="form.title" placeholder="请输入" />
      </a-form-model-item>
      <a-form-model-item label="公告类型" prop="type">
        <a-select placeholder="请选择" v-model="form.type">
          <a-select-option v-for="(d, index) in typeOptions" :key="index" :value="d.value" >{{ d.label }}</a-select-option>
        </a-select>
      </a-form-model-item>
      <a-form-model-item label="状态" prop="status">
        <a-radio-group v-model="form.status" button-style="solid">
          <a-radio-button v-for="(d, index) in statusOptions" :key="index" :value="d.value" >{{ d.label }}</a-radio-button>
        </a-radio-group>
      </a-form-model-item>
      <a-form-model-item label="内容" prop="content" class="addModel">
        <div id="vditorEdit" name="description" ></div>
      </a-form-model-item>
      <div class="bottom-control" style="z-index: 9999;">
        <a-space>
          <a-button type="primary" @click="submitForm">
            发布
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
import NoticeForm from './NoticeForm'
export default {
  ...NoticeForm
}
</script>
<style lang="less" scoped>
#vditor {
  margin: 0 5% 5% 5%;
}
.addModel{
    .vditor-toolbar--pin{
      top:auto!important;
    }
  }
</style>
