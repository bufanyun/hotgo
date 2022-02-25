<template>
  <a-drawer
    width="35%"
    :title="formTitle"
    :label-col="4"
    :wrapper-col="14"
    :visible="open"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    @close="cancel">
    <a-form-model ref="form" :model="form" :rules="rules">
      <a-form-model-item label="岗位名称" prop="name">
        <a-input v-model="form.name" placeholder="请输入" />
      </a-form-model-item>
      <a-form-model-item label="岗位编码" prop="code">
        <a-input v-model="form.code" placeholder="请输入" />
      </a-form-model-item>
      <a-form-model-item label="排序" prop="sort">
        <a-input-number v-model="form.sort" :min="0" style="width: 100%"/>
      </a-form-model-item>
      <a-form-model-item label="状态" prop="status">
        <a-radio-group v-model="form.status" button-style="solid">
          <a-radio-button v-for="(d, index) in statusOptions" :key="index" :value="d.value" >{{ d.label }}</a-radio-button>
        </a-radio-group>
      </a-form-model-item>
      <a-form-model-item label="备注" prop="remark">
        <a-input v-model="form.remark" placeholder="请输入" type="textarea" :rows="2" allow-clear />
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
import PostForm from './PostForm'
export default {
  ...PostForm
}
</script>
