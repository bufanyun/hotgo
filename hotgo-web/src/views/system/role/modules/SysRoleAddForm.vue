<template>

  <ant-modal
    :visible="open"
    :modal-title="formTitle"
    :adjust-size="true"
    :modal-cut-height="108"
    @cancel="cancel"
  >
    <a-form-model ref="form" :model="form" :rules="rules" slot="content">
      <a-row :gutter="48" >
        <a-col :span="8" >
          <a-form-model-item label="角色名称" prop="roleName">
            <a-input v-model="form.roleName" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :span="8" >
          <a-form-model-item label="角色编码" prop="roleKey">
            <a-input v-model="form.roleKey" placeholder="请输入" />
          </a-form-model-item>
        </a-col>
        <a-col :span="8" >
          <a-form-model-item label="排序" prop="sort">
            <a-input-number placeholder="请输入" v-model="form.sort" :min="0" style="width: 100%"/>
          </a-form-model-item>
        </a-col>
      </a-row>
      <a-row :gutter="48" >
        <a-col :span="8" >
          <a-form-model-item label="状态" prop="status">
            <a-radio-group v-model="form.status" button-style="solid">
              <a-radio-button v-for="(d, index) in statusOptions" :key="index" :value="d.dictValue" >{{ d.dictLabel }}</a-radio-button>
            </a-radio-group>
          </a-form-model-item>
        </a-col>
        <a-col :span="16">
          <a-form-model-item label="菜单权限">
            <a-checkbox @change="handleCheckedTreeExpand($event)">
              展开/折叠
            </a-checkbox>
            <a-checkbox @change="handleCheckedTreeNodeAll($event)">
              全选/全不选
            </a-checkbox>
            <a-checkbox @change="handleCheckedTreeConnect($event)" :checked="form.menuCheckStrictly">
              父子联动
            </a-checkbox>
            <a-tree
              v-model="menuCheckedKeys"
              checkable
              :checkStrictly="!form.menuCheckStrictly"
              :expanded-keys="menuExpandedKeys"
              :auto-expand-parent="autoExpandParent"
              :tree-data="menuOptions"
              @check="onCheck"
              @expand="onExpandMenu"
              :replaceFields="defaultProps"
            />
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
      <a-button type="primary" @click="submitForm">
        保存
      </a-button>
      <a-button type="dashed" @click="cancel">
        取消
      </a-button>
    </template>
  </ant-modal>
</template>

<script>
  import sysRoleForm from './SysRoleForm'
  export default {
          ...sysRoleForm
      }
</script>
