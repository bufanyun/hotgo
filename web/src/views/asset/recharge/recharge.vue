<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="在线充值"> 余额可用于购买付费产品或商城消费 </n-card>
    </div>
    <n-spin :show="loading">
      <n-grid cols="1 s:1 m:1 l:4 xl:4 2xl:4" responsive="screen" :x-gap="12">
        <n-gi span="4">
          <n-card :bordered="false" class="proCard">
            <n-thing>
              <template #description> <span class="title">充值金额</span> </template>
              <n-space>
                <n-button
                  type="primary"
                  ghost
                  v-for="item in amounts"
                  :key="item"
                  @click="SetAmount(item)"
                >
                  ￥{{ item }}

                  <n-icon
                    class="check-icon"
                    :size="18"
                    :component="CheckOutlined"
                    v-if="amount === item && amountType === 1"
                  />
                </n-button>

                <n-input-number v-model:value="amount" v-if="amountType === 2">
                  <template #prefix> ￥ </template>
                </n-input-number>
              </n-space>
              <template #footer> <span class="title">支付方式 </span></template>
              <template #action>
                <n-space>
                  <n-button
                    strong
                    secondary
                    :color="item.color"
                    v-for="item in payTypes"
                    :key="item"
                    @click="SetPayType(item.value)"
                  >
                    <template #icon>
                      <n-icon :component="item.icon" />
                    </template>
                    {{ item.label }}
                    <n-icon
                      class="check-icon"
                      :size="18"
                      :component="CheckOutlined"
                      v-if="payType === item.value"
                    />
                  </n-button>
                </n-space>

                <n-button
                  type="success"
                  class="create-order-button"
                  size="large"
                  @click="CreateOrder"
                >
                  立即充值
                </n-button>
              </template>
            </n-thing>
          </n-card>
        </n-gi>
      </n-grid>
    </n-spin>

    <n-modal v-model:show="showQrModal" :show-icon="false" preset="dialog" :title="qrParams.name">
      <n-form class="py-4">
        <div class="text-center">
          <qrcode-vue :value="qrParams.qrUrl" :size="220" class="canvas" style="margin: 0 auto" />
        </div>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showQrModal = false)">关闭</n-button>
        </n-space>
      </template>
    </n-modal>
    <RechargeLog class="mt-6" />
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, inject } from 'vue';
  import wx from 'weixin-js-sdk';
  import { WechatOutlined, AlipayOutlined, QqOutlined, CheckOutlined } from '@vicons/antd';
  import { useMessage } from 'naive-ui';
  import { Create } from '@/api/order';
  import QrcodeVue from 'qrcode.vue';
  import RechargeLog from '../rechargeLog/index.vue';
  import { SocketEnum } from '@/enums/socketEnum';
  import { addOnMessage } from '@/utils/websocket';

  const showQrModal = ref(false);
  const qrParams = ref({
    name: '',
    qrUrl: '',
  });
  const loading = ref(false);
  const message = useMessage();
  const amountType = ref(1);
  const amounts = ref([0.01, 10, 20, 30, 50, 100, '其他金额']);
  const payTypes = ref([
    { value: 'wxpay', label: '微信支付', icon: WechatOutlined, color: '#18a058' },
    { value: 'alipay', label: '支付宝', icon: AlipayOutlined, color: '#2d8cf0' },
    { value: 'qqpay', label: 'QQ支付', icon: QqOutlined, color: '#2d8cf0' },
  ]);

  const amount = ref<any>(null);
  const payType = ref<any>(null);

  onMounted(() => {});

  function SetPayType(type: string) {
    payType.value = type;
  }

  function SetAmount(a: number | string) {
    amount.value = a;
    if (a === '其他金额') {
      amountType.value = 2;
      amount.value = null;
    } else {
      amountType.value = 1;
    }
  }

  function CreateOrder() {
    if (amount.value === null || amount.value <= 0) {
      message.error('请选择充值金额');
      return;
    }

    if (payType.value === null || payType.value === '') {
      message.error('请选择支付方式');
      return;
    }

    loading.value = true;
    Create({
      orderType: 'balance',
      payType: payType.value,
      money: amount.value,
      returnUrl: window.location.href,
    })
      .then((res) => {
        if (res.order?.tradeType === undefined || res.order?.tradeType === '') {
          message.error('创建支付订单失败，没找到交易方式，请联系管理处理！');
          return;
        }

        if (res.order?.tradeType !== 'mp') {
          if (res.order?.payURL === undefined || res.order?.payURL === '') {
            message.error('创建支付订单失败，没找到支付地址，请联系管理处理！');
            return;
          }
        }

        switch (res.order?.tradeType) {
          case 'scan':
            showQr(res.order?.payURL, '打开微信【扫一扫】完成支付');
            break;
          case 'mp':
            if (res.order.jsApi === undefined) {
              message.error('支付失败请选择其他支付方式：JSAPI支付参数无效');
              return;
            }
            const jsApi = res.order.jsApi;

            // 配置微信JS SDK
            wx.config({
              // debug: true,
              appId: jsApi.config.app_id,
              timestamp: jsApi.config.timestamp,
              nonceStr: jsApi.config.nonce_str,
              signature: jsApi.config.signature,
              jsApiList: ['chooseWXPay'],
            });
            // 配置完成后返回一个resolve
            wx.ready(() => {
              wxJSPay({
                timestamp: jsApi.params.timeStamp,
                nonceStr: jsApi.params.nonceStr,
                package: jsApi.params.package,
                signType: jsApi.params.signType,
                paySign: jsApi.params.paySign,
              })
                .then((_res) => {
                  // ...
                })
                .catch((err) => {
                  message.success('支付失败：', err.message);
                });
            });

            break;
          case 'qqweb':
            showQr(res.order?.payURL, '打开QQ【扫一扫】完成支付');
            break;
          default:
            window.open(res.order?.payURL, '_blank');
        }
      })
      .finally(() => {
        loading.value = false;
      });
  }

  // 发起微信公众号支付
  function wxJSPay(params) {
    return new Promise((resolve, reject) => {
      // 调用微信支付
      wx.chooseWXPay({
        timestamp: params.timestamp,
        nonceStr: params.nonceStr,
        package: params.package,
        signType: params.signType,
        paySign: params.paySign,
        success: (res) => {
          // 支付成功时返回resolve
          resolve(res);
        },
        fail: (err) => {
          // 支付失败时返回reject
          reject(err);
        },
      });
    });
  }

  function showQr(url: string, name: string) {
    qrParams.value.qrUrl = url;
    qrParams.value.name = name;
    showQrModal.value = true;
  }

  const onMessageList = inject('onMessageList');

  const handleMessageList = (res) => {
    const data = JSON.parse(res.data);
    if (data.event === SocketEnum.EventAdminOrderNotify) {
      if (data.code == SocketEnum.CodeErr) {
        message.error('查询出错:' + data.event);
        return;
      }

      showQrModal.value = false;
      message.success('支付成功');

      location.reload();
      return;
    }
  };

  addOnMessage(onMessageList, handleMessageList);
</script>

<style lang="less" scoped>
  ::v-deep(.n-thing .n-thing-main .n-thing-main__footer:not(:first-child)) {
    margin-top: 36px;
    margin-bottom: 10px;
  }

  ::v-deep(.title) {
    font-weight: var(--n-title-font-weight);
    transition: color 0.3s var(--n-bezier);
    flex: 1;
    min-width: 0;
    color: var(--n-title-text-color);
    font-size: 18px;
  }

  ::v-deep(.check-icon) {
    margin-left: 3px;
  }

  ::v-deep(.create-order-button) {
    margin-top: 28px;
  }
</style>
