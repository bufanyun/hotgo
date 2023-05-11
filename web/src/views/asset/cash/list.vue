<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        @keyup.enter="handleSubmit"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>

      <BasicTable
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        :scroll-x="1800"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable" class="min-left-space">
            <template #icon>
              <n-icon>
                <MoneyCollectOutlined />
              </n-icon>
            </template>
            申请提现
          </n-button>

          <n-button type="default" @click="setCash" class="min-left-space">
            <template #icon>
              <n-icon>
                <EditOutlined />
              </n-icon>
            </template>
            设置提现账户
          </n-button>
        </template>
      </BasicTable>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        preset="dialog"
        title="申请提现"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-alert type="info">
          <div v-html="config.cashTips"></div>
        </n-alert>
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="100"
          class="py-4"
        >
          <n-form-item label="可提现金额">
            <n-input v-model:value="newUserInfo.balance" disabled />
            <template #feedback
              ><p>{{ estimated }}</p></template
            >
          </n-form-item>
          <br />
          <n-form-item label="提现金额" path="money">
            <n-input-number v-model:value="formParams.money" :min="1" :max="newUserInfo.balance">
              <template #minus-icon>
                <n-icon :component="ArrowDownCircleOutline" />
              </template>
              <template #add-icon>
                <n-icon :component="ArrowUpCircleOutline" />
              </template>
            </n-input-number>
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>

      <n-modal
        v-model:show="showPaymentModal"
        :show-icon="false"
        preset="dialog"
        title="处理打款"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="paymentParams"
          :rules="rules"
          ref="PaymentRef"
          label-placement="left"
          :label-width="100"
          class="py-4"
        >
          <n-form-item label="最终到账金额">
            <n-input v-model:value="paymentParams.lastMoney" disabled />
          </n-form-item>
          <n-form-item label="收款信息">
            <n-input v-model:value="paymentParams.accountInfo" disabled />
          </n-form-item>

          <n-form-item label="收款码">
            <n-carousel draggable>
              <img style="width: 200px" class="carousel-img" :src="paymentParams.payeeCode" />
            </n-carousel>
          </n-form-item>

          <n-form-item label="提现状态" path="status">
            <n-radio-group v-model:value="paymentParams.status" name="status">
              <n-radio-button
                v-for="status in statusOptions"
                :key="status.value"
                :value="status.value"
                :label="status.label"
              />
            </n-radio-group>
          </n-form-item>

          <n-form-item label="处理结果">
            <n-input v-model:value="paymentParams.msg" />
            <template #feedback>不填默认显示提现状态</template>
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showPaymentModal = false)">取消</n-button>
            <n-button type="info" :loading="PaymentBtnLoading" @click="confirmPayment"
              >确定</n-button
            >
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Apply, List, Payment, View } from '@/api/cash';
  import { columns, statusOptions } from './columns';
  import { ArrowDownCircleOutline, ArrowUpCircleOutline } from '@vicons/ionicons5';
  import { MoneyCollectOutlined, EditOutlined } from '@vicons/antd';
  import { defRangeShortcuts, timestampToTime } from '@/utils/dateUtil';
  import { useRouter } from 'vue-router';
  import { getUserInfo } from '@/api/system/user';
  import { getCashConfig } from '@/api/sys/config';

  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '',
  });
  const router = useRouter();

  const params = ref<any>({
    pageSize: 10,
    title: '',
    content: '',
    status: null,
  });

  const rules = {};

  const estimated = ref(
    '本次提现预计将在 ' +
      timestampToTime(new Date().setTime(new Date().getTime() + 86400 * 4 * 1000) / 1000) +
      ' 前到账 (1-3个工作日，双休日和法定节假日顺延)'
  );

  const schemas: FormSchema[] = [
    {
      field: 'memberId',
      component: 'NInput',
      label: '管理员ID',
      componentProps: {
        placeholder: '请输入管理员ID',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入管理员ID', trigger: ['blur'] }],
    },
    {
      field: 'ip',
      component: 'NInput',
      label: '申请IP',
      componentProps: {
        placeholder: '请输入申请IP',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入申请IP', trigger: ['blur'] }],
    },
    {
      field: 'created_at',
      component: 'NDatePicker',
      label: '申请时间',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        shortcuts: defRangeShortcuts(),
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const newUserInfo = ref({ balance: 0 });
  const message = useMessage();
  const actionRef = ref();
  const showModal = ref(false);
  const showPaymentModal = ref(false);
  const PaymentRef = ref<any>({});
  const PaymentBtnLoading = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const config = ref<any>({
    cashMinFee: 3,
    cashMinFeeRatio: '0.03',
    cashMinMoney: 0,
    cashSwitch: false,
    cashTips: '',
  });

  const resetFormParams = {
    money: null,
    accountInfo: null,
  };
  let formParams = ref<any>(resetFormParams);

  const resetPaymentParams = {
    id: null,
    money: null,
  };
  let paymentParams = ref<any>(resetPaymentParams);

  const actionColumn = reactive({
    auth: ['/cash/payment'],
    width: 100,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '处理打款',
            onClick: handleEdit.bind(null, record),
          },
        ],
      });
    },
  });

  function setCash() {
    router.push({
      name: 'home_account',
      query: {
        type: 3,
      },
    });
  }

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  async function addTable() {
    showModal.value = true;
    formParams.value = resetFormParams;
    newUserInfo.value = await getUserInfo();

    if (newUserInfo.value.balance < config.value.cashMinMoney) {
      message.error('当前余额不满足提现条件，至少需要：' + config.value.cashMinMoney + '元');
    }
  }

  const loadDataTable = async (res) => {
    mapWidth();
    config.value = await getCashConfig();
    config.value = config.value.list;
    return await List({
      ...params.value,
      ...res,
      ...searchFormRef.value.formModel,
      ...{ status: props.type },
    });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  /**
   * 申请提现
   * @param e
   */
  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Apply({ money: formParams.value.money })
          .then((_res) => {
            message.success('操作成功');
            setTimeout(() => {
              showModal.value = false;
              reloadTable();
              formParams.value = ref(resetFormParams);
            });
          })
          .catch((_e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  /**
   * 处理打款
   * @param e
   */
  function confirmPayment(e) {
    e.preventDefault();
    PaymentBtnLoading.value = true;
    PaymentRef.value.validate((errors) => {
      if (!errors) {
        Payment({
          id: PaymentRef.value.model.id,
          status: PaymentRef.value.model.status,
          msg: PaymentRef.value.model.msg,
        })
          .then((_res) => {
            message.success('操作成功');
            setTimeout(() => {
              showPaymentModal.value = false;
              reloadTable();
              PaymentRef.value = ref(resetPaymentParams);
            });
          })
          .catch((_e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      } else {
        message.error('请填写完整信息');
      }
      PaymentBtnLoading.value = false;
    });
  }

  async function handleEdit(record: Recordable) {
    showPaymentModal.value = true;
    paymentParams.value = record;
    paymentParams.value = await View({ id: record.id });
    paymentParams.value.lastMoney = paymentParams.value.lastMoney.toFixed(2);
    paymentParams.value.accountInfo =
      paymentParams.value.name + ' - ' + paymentParams.value.account;
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    params.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  const dialogWidth = ref('50%');

  function mapWidth() {
    let val = document.body.clientWidth;
    const def = 720; // 默认宽度
    if (val < def) {
      dialogWidth.value = '100%';
    } else {
      dialogWidth.value = def + 'px';
    }

    return dialogWidth.value;
  }
</script>

<style lang="less" scoped></style>
