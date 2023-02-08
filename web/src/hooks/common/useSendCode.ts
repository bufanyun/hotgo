import { computed } from 'vue';
import useLoading from './useLoading';
import useCountDown from './useCountDown';

export default function useSmsCode() {
  const { loading, startLoading, endLoading } = useLoading();
  const { counts, start, isCounting } = useCountDown(60);
  const initLabel = '获取验证码';
  const countingLabel = (second: number) => `重新获取(${second})`;
  const sendLabel = computed(() => {
    let text = initLabel;
    if (loading.value) {
      text = '';
    }
    if (isCounting.value) {
      text = countingLabel(counts.value);
    }
    return text;
  });

  /**
   * 激活发送
   */
  function activateSend(request: Promise<any>) {
    startLoading();
    request
      .then((_res) => {
        window['$message']?.success('验证码发送成功！');
        start();
      })
      .finally(() => {
        endLoading();
      });
  }

  return {
    sendLabel,
    start,
    isCounting,
    activateSend,
    loading,
  };
}
