<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="登录验证码开关" path="loginCaptchaSwitch">
          <n-radio-group v-model:value="formValue.loginCaptchaSwitch" name="loginCaptchaSwitch">
            <n-space>
              <n-radio :value="1">开启</n-radio>
              <n-radio :value="2">关闭</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="注册开关" path="loginRegisterSwitch">
          <n-radio-group v-model:value="formValue.loginRegisterSwitch" name="cashSwitch">
            <n-space>
              <n-radio :value="1">开启</n-radio>
              <n-radio :value="2">关闭</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="默认注册头像" path="loginAvatar">
          <UploadImage :maxNumber="1" v-model:value="formValue.loginAvatar" />
        </n-form-item>

        <n-form-item label="默认注册角色" path="loginRoleId">
          <n-tree-select
            key-field="id"
            :options="options.role"
            v-model:value="formValue.loginRoleId"
            :default-expand-all="true"
          />
        </n-form-item>

        <n-form-item label="默认注册部门" path="loginDeptId">
          <n-tree-select
            key-field="id"
            :options="options.dept"
            v-model:value="formValue.loginDeptId"
            :default-expand-all="true"
          />
        </n-form-item>

        <n-form-item label="默认注册岗位" path="loginPostIds">
          <n-select v-model:value="formValue.loginPostIds" multiple :options="options.post" />
        </n-form-item>

        <n-form-item label="用户协议" path="loginProtocol">
          <Editor
            style="height: 320px"
            v-model:value="formValue.loginProtocol"
            id="loginProtocol"
          />
        </n-form-item>

        <n-form-item label="隐私权政策" path="loginPolicy">
          <Editor style="height: 320px" v-model:value="formValue.loginPolicy" id="loginPolicy" />
        </n-form-item>

        <div>
          <n-space>
            <n-button type="primary" @click="formSubmit">保存更新</n-button>
          </n-space>
        </div>
      </n-form>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
  import { onMounted, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import Editor from '@/components/Editor/editor.vue';
  import { getDeptOption } from '@/api/org/dept';
  import { getRoleOption } from '@/api/system/role';
  import { getPostOption } from '@/api/org/post';
  import UploadImage from '@/components/Upload/uploadImage.vue';

  const group = ref('login');
  const show = ref(false);
  const rules = {};
  const formRef: any = ref(null);
  const message = useMessage();
  const formValue = ref({
    loginRegisterSwitch: true,
    loginCaptchaSwitch: true,
    loginAvatar: '',
    loginProtocol: '',
    loginPolicy: '',
    loginRoleId: null,
    loginDeptId: null,
    loginPostIds: [],
  });

  const options = ref<any>({
    role: [],
    roleTabs: [{ id: -1, name: '全部' }],
    dept: [],
    post: [],
  });

  async function loadOptions() {
    const dept = await getDeptOption();
    if (dept.list !== undefined) {
      options.value.dept = dept.list;
    }

    const role = await getRoleOption();
    if (role.list !== undefined) {
      options.value.role = role.list;
      treeDataToCompressed(role.list);
    }

    const post = await getPostOption();
    if (post.list !== undefined && post.list.length > 0) {
      for (let i = 0; i < post.list.length; i++) {
        post.list[i].label = post.list[i].name;
        post.list[i].value = post.list[i].id;
      }
      options.value.post = post.list;
    }
  }

  function treeDataToCompressed(source) {
    for (const i in source) {
      options.value.roleTabs.push(source[i]);
      source[i].children && source[i].children.length > 0
        ? treeDataToCompressed(source[i].children)
        : ''; // 子级递归
    }

    return options.value.roleTabs;
  }

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        updateConfig({ group: group.value, list: formValue.value })
          .then((_res) => {
            message.success('更新成功');
            load();
          })
          .catch((error) => {
            message.error(error.toString());
          });
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }

  onMounted(async () => {
    await loadOptions();
    load();
  });

  function load() {
    show.value = true;
    new Promise((_resolve, _reject) => {
      getConfig({ group: group.value })
        .then((res) => {
          show.value = false;
          formValue.value = res.list;
        })
        .catch((error) => {
          show.value = false;
          message.error(error.toString());
        });
    });
  }
</script>
