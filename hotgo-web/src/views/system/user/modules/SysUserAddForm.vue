<template>
  <!-- 增加修改 -->
  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    modalHeight="480"
    modalWidth="640"
    @cancel="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules" slot="content" layout="vertical">
      <a-row class="form-row" :gutter="32">
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="姓名" prop="realname">
            <a-input v-model="form.realname" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="登录名" prop="username" v-if="form.id === undefined">
            <a-input v-model="form.username" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="所属部门" prop="dept_id">
            <select-dept
              v-model="form.dept_id"
              select-model="single"
            />
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="岗位" prop="postIds">
            <a-select
              mode="multiple"
              v-model="form.postIds"
              placeholder="请选择"
              option-filter-prop="children">
              <a-select-option v-for="(d, index) in postOptions" :key="index" :value="d.id">
                {{ d.name }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
        <a-col :lg="12" :md="12" :sm="24">
          <a-form-model-item label="角色" prop="role">
            <a-select v-model="form.role" placeholder="请选择" option-filter-prop="children">
              <a-select-option v-for="(d, index) in roleOptions" :key="index" :value="d.id">
                {{ d.name }}
              </a-select-option>
            </a-select>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-collapse :bordered="false" expandIconPosition="left">
        <template #expandIcon="props">
          <a-icon type="caret-right" :rotate="props.isActive ? 90 : 0" />
        </template>
        <a-collapse-panel key="1" header="填写更多信息（可选)" :style="customStyle">
          <a-row gutter="32">
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="手机号" prop="mobile">
                <a-input v-model="form.mobile" placeholder="请输入" />
              </a-form-model-item>
            </a-col>
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="邮箱地址" prop="email">
                <a-input v-model="form.email" placeholder="请输入" />
              </a-form-model-item>
            </a-col>
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="状态" prop="status">
                <a-radio-group v-model="form.status" button-style="solid">
                  <a-radio-button v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{
                    d.label
                  }}</a-radio-button>
                </a-radio-group>
              </a-form-model-item>
            </a-col>

            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="性别" prop="sex">
                <a-radio-group v-model="form.sex" button-style="solid">
                  <a-radio-button v-for="(d, index) in sexOptions" :key="index" :value="d.value">{{
                    d.label
                  }}</a-radio-button>
                </a-radio-group>
              </a-form-model-item>
            </a-col>
            <a-col :lg="12" :md="12" :sm="24">
              <a-form-model-item label="密码" prop="password">
                <a-input-password v-model="form.password" placeholder="请输入" />
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
      <a-button type="primary" @click="submitForm">
        保存
      </a-button>
      <a-button @click="cancel">
        取消
      </a-button>
    </template>
  </ant-modal>
</template>
<script>
import sysUserForm from './SysUserForm'
export default {
        ...sysUserForm
    }
</script>
<style></style>
