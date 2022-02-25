<template>
  <a-drawer
    width="35%"
    :label-col="4"
    :title="formTitle"
    :wrapper-col="14"
    :visible="open"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    @close="cancel">
    <a-form-model ref="form" :model="form" :rules="rules">
      <a-form-model-item label="字典名称" prop="name">
        <a-input v-model="form.name" placeholder="请输入字典名称" />
      </a-form-model-item>
      <a-form-model-item label="字典类型" prop="type">
        <a-input v-model="form.type" placeholder="请输入字典类型" />
      </a-form-model-item>
      <a-form-model-item label="状态" prop="status">
        <a-select placeholder="请选择" v-model="form.status">
          <a-select-option v-for="(d, index) in statusOptions" :key="index" :value="d.value" >{{ d.label }}</a-select-option>
        </a-select>
      </a-form-model-item>
      <a-form-model-item label="备注" prop="remark">
        <a-input v-model="form.remark" placeholder="请输入内容" type="textarea" allow-clear />
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
import DictTypeForm from './DictTypeForm'
export default {
  ...DictTypeForm
}
</script>
