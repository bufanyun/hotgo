<template>
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    @cancel="cancel"
    modalHeight="480"
    modalWidth="640">
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row class="form-row" :gutter="32">
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="上级部门" prop="pid">
            <select-dept
              v-model="form.pid"
              select-model="single"
              :select-scope="selectScope"
              @callBack="onSelectDept"
            />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="部门名称" prop="name">
            <a-input v-model="form.name" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="排序" prop="sort">
            <a-input-number v-model="form.sort" :min="0" style="width: 100%"/>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-collapse :bordered="false" expandIconPosition="left">
        <template #expandIcon="props">
          <a-icon type="caret-right" :rotate="props.isActive ? 90 : 0" />
        </template>
        <a-collapse-panel key="1" header="详细信息" :style="customStyle">
          <a-row :gutter="32">
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="负责人" prop="leader">
                <a-input v-model="form.leader" placeholder="请输入" />
              </a-form-model-item>
            </a-col>
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="办公电话" prop="phone">
                <a-input v-model="form.phone" placeholder="请输入" />
              </a-form-model-item>
            </a-col>
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="联系地址" prop="address">
                <a-input v-model="form.address" placeholder="请输入" />
              </a-form-model-item>
            </a-col>
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="电子邮箱" prop="email">
                <a-input v-model="form.email" placeholder="请输入" />
              </a-form-model-item>
            </a-col>
            <a-col :lg="24" :md="24" :sm="24" :span="24">
              <a-form-model-item label="备注" prop="remark">
                <a-input v-model="form.remark" placeholder="请输入" type="textarea" allow-clear />
              </a-form-model-item>
            </a-col>
          </a-row>
        </a-collapse-panel>
      </a-collapse>

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
  import sysDeptForm from './SysDeptForm'
  export default {
    ...sysDeptForm
      }
</script>
