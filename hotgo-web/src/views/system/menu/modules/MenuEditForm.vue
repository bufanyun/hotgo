<template>
  <a-drawer
    width="480"
    :title="formTitle"
    :label-col="4"
    :wrapper-col="14"
    :visible="open"
    :body-style="{height:'calc(100vh - 100px)',overflow:'auto'}"
    @close="cancel">
    <a-form-model ref="form" :model="form" :rules="rules" layout="vertical">
      <a-spin :spinning="spinning" :delay="delayTime" tip="Loading...">
        <a-form-model-item label="上级菜单" prop="pid">
          <a-tree-select
            v-model="form.pid"
            style="width: 100%"
            :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
            :tree-data="menuOptions"
            placeholder="请选择"
            :replaceFields="{ children: 'children', title: 'name', key: 'id', value: 'id' }"
            tree-default-expand-all
            @change="onMenuTreeChange"
          >
          </a-tree-select>
        </a-form-model-item>
        <a-form-model-item label="菜单类型" prop="type">
          <a-radio-group v-model="form.type" button-style="solid">
            <a-radio-button v-for="(d, index) in menuTypeOptions" :key="index" :disabled="menuTypeEnableValue.indexOf(d.menuTypeValue) === -1" :value="d.menuTypeValue">{{ d.menuTypeLabel }}</a-radio-button>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item label="图标" prop="icon">
          <a-space size="large" class="selectIconBox">
            <a-icon :component="allIcon[form.icon + 'Icon']" v-if="allIcon[form.icon + 'Icon']" class="selectIcon" />
            <a-icon :type="form.icon" v-if="!allIcon[form.icon + 'Icon']" />
            <a @click="selectIcon" class="selectup">
              <a-icon :type="SelectIcon" />
            </a>
          </a-space>
          <a-card :bordered="false" v-if="iconVisible">
            <icon-selector v-model="form.icon" @change="handleIconChange" :svgIcons="iconList" :allIcon="allIcon" />
          </a-card>
        </a-form-model-item>
        <a-form-model-item label="菜单编码" prop="code">
          <a-input v-model="form.code" placeholder="请输入" />
        </a-form-model-item>
        <a-form-model-item label="菜单名称" prop="name">
          <a-input v-model="form.name" placeholder="请输入" />
        </a-form-model-item>
        <a-form-model-item label="排序" prop="sort">
          <a-input-number v-model="form.sort" :min="0" style="width: 100%" />
        </a-form-model-item>
        <a-form-model-item label="是否外链" prop="is_frame" v-if="form.type !== 'F'">
          <a-radio-group v-model="form.is_frame" button-style="solid">
            <a-radio-button value="1">是</a-radio-button>
            <a-radio-button value="2">否</a-radio-button>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item label="路由地址" prop="path" v-if="form.type !== 'F'">
          <a-input v-model="form.path" placeholder="请输入" />
        </a-form-model-item>
        <a-form-model-item label="组件路径" prop="component" v-if="form.type === 'C'">
          <a-input v-model="form.component" placeholder="请输入" />
        </a-form-model-item>
        <a-form-model-item label="权限标识" prop="perms" v-if="form.type !== 'M'">
          <a-input v-model="form.perms" placeholder="请输入" />
        </a-form-model-item>
        <a-form-model-item label="是否显示" prop="is_visible" v-if="form.type !== 'F'">
          <a-radio-group v-model="form.is_visible" button-style="solid">
            <a-radio-button v-for="(d, index) in visibleOptions" :key="index" :value="d.value">{{
              d.label
            }}</a-radio-button>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item label="状态" prop="status" v-if="form.type !== 'F'">
          <a-radio-group v-model="form.status" button-style="solid">
            <a-radio-button v-for="(d, index) in statusOptions" :key="index" :value="d.value">{{
              d.label
            }}</a-radio-button>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item label="是否缓存" prop="is_cache" v-if="form.type === 'C'">
          <a-radio-group v-model="form.is_cache" button-style="solid">
            <a-radio-button value="1">缓存</a-radio-button>
            <a-radio-button value="2">不缓存</a-radio-button>
          </a-radio-group>
        </a-form-model-item>
        <a-form-model-item label="备注" prop="remark">
          <a-input v-model="form.remark" placeholder="请输入" type="textarea" allow-clear />
        </a-form-model-item>
      </a-spin>
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
import MenuForm from './MenuForm'
export default {
  ...MenuForm
}
</script>
