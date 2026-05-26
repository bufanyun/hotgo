<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-form :label-width="100" :model="formValue" :rules="rules" ref="formRef">
        <n-form-item label="申请提现开关" path="cashSwitch">
          <n-radio-group v-model:value="formValue.cashSwitch" name="cashSwitch">
            <n-space>
              <n-radio :value="1">开启</n-radio>
              <n-radio :value="2">关闭</n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="提现最低手续费（元）" path="cashMinFee">
          <n-input-number placeholder="" v-model:value="formValue.cashMinFee" style="width: 100%" />
        </n-form-item>

        <n-form-item label="提现最低手续费比率" path="cashMinFeeRatio">
          <n-input-number
            placeholder=""
            v-model:value="formValue.cashMinFeeRatio"
            style="width: 100%"
          />
        </n-form-item>

        <n-form-item label="提现最低金额" path="cashMinMoney">
          <n-input-number
            placeholder=""
            v-model:value="formValue.cashMinMoney"
            style="width: 100%"
          />
        </n-form-item>

        <n-form-item label="提现提示信息" path="cashTips">
          <Editor style="height: 320px" v-model:value="formValue.cashTips" />
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

  const group = ref('cash');
  const show = ref(false);
  const rules = {};
  const formRef: any = ref(null);
  const message = useMessage();
  const formValue = ref({
    cashSwitch: '',
    cashMinFee: 0,
    cashMinFeeRatio: 0,
    cashMinMoney: 0,
    cashTips: '',
  });

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

  onMounted(() => {
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
