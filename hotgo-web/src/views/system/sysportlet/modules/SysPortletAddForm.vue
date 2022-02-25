<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    modalHeight="480"
    modalWidth="640"
    @cancel="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row :gutter="32">
        <a-col :span="12">
          <a-form-model-item label="小页名称" prop="name">
            <a-input v-model="form.name" placeholder="请输入小页名称"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="小页编码" prop="code">
            <a-input v-model="form.code" placeholder="请输入小页编码"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="小页URL" prop="url">
            <a-input v-model="form.url" placeholder="请输入小页URL"/>
          </a-form-model-item>
        </a-col>
        <!--   <a-col :span="12">
          <a-form-model-item label="刷新频率" prop="refreshRate">
            <a-input v-model="form.refreshRate" placeholder="请输入刷新频率"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12" >
          <a-form-model-item label="是否显示标题" prop="showTitle">
            <a-radio-group v-model="form.showTitle" button-style="solid">
              <a-radio-button
                v-for="(dict, index) in showTitleOptions"
                :key="index"
                :value="dict.dictValue"
              >
                {{ dict.dictLabel }}
              </a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col> -->
        <a-col :span="12" >
          <a-form-model-item label="是否允许拖拽" prop="isAllowDrag">
            <a-radio-group v-model="form.isAllowDrag" button-style="solid">
              <a-radio-button
                v-for="(dict, index) in isAllowDragOptions"
                :key="index"
                :value="dict.dictValue"
              >
                {{ dict.dictLabel }}
              </a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="横向比例" prop="xGridNumber">
            <a-select
              v-model="form.xGridNumber"
              :getPopupContainer="
                triggerNode => {
                  return triggerNode.parentNode || document.body
                }
              "
              style="width:100%"
            >
              <a-select-option value="3">1行</a-select-option>
              <a-select-option value="2">2/3行</a-select-option>
              <a-select-option value="1">1/3行</a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="高度" prop="yGridNumber">
            <a-input-number v-model="form.yGridNumber" :min="10" style="width: 80%"/>&nbsp;&nbsp;&nbsp;&nbsp;px
          </a-form-model-item>
        </a-col>
        <a-col :span="12">
          <a-form-model-item label="排序号" prop="sort">
            <a-input-number v-model="form.sort" :min="0" style="width: 100%"/>
          </a-form-model-item>
        </a-col>
        <a-col :span="12" >
          <a-form-model-item label="状态" prop="status">
            <a-radio-group v-model="form.status" button-style="solid">
              <a-radio-button
                v-for="(dict, index) in statusOptions"
                :key="index"
                :value="dict.dictValue"
              >
                {{ dict.dictLabel }}
              </a-radio-button>
            </a-radio-group>
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
import SysPortletAddForm from './SysPortletForm'
export default {
  ...SysPortletAddForm
}
</script>
