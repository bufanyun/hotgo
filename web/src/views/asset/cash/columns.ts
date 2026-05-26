import { h } from 'vue';
import { NTag } from 'naive-ui';
import { renderPopoverMemberSumma } from '@/utils';

const msgMap = {
  1: '处理中',
  2: '提现成功',
  3: ' 提现异常',
};

export const statusOptions = [
  {
    value: 1,
    label: '处理中',
  },
  {
    value: 2,
    label: '提现成功',
  },
  {
    value: 3,
    label: '提现异常',
  },
];

export const columns = [
  {
    title: '提现ID',
    key: 'id',
    width: 100,
  },
  {
    title: '申请人',
    key: 'memberId',
    width: 100,
    render(row) {
      return renderPopoverMemberSumma(row.memberBySumma);
    },
  },
  {
    title: '提现金额',
    key: 'money',
    render(row) {
      return row.money.toFixed(2);
    },
    width: 100,
  },
  {
    title: '手续费',
    key: 'fee',
    render(row) {
      return row.fee.toFixed(2);
    },
    width: 100,
  },
  {
    title: '最终到账',
    key: 'lastMoney',
    render(row) {
      return row.lastMoney.toFixed(2);
    },
    width: 100,
  },
  {
    title: '处理结果',
    key: 'msg',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: row.status == 1 ? 'info' : row.status == 2 ? 'success' : 'warning',
          bordered: false,
        },
        {
          default: () => (row.msg == '' ? msgMap[row.status] : row.msg),
        }
      );
    },
    width: 200,
  },
  {
    title: '申请IP',
    key: 'ip',
    width: 180,
  },
  {
    title: '处理时间',
    key: 'handleAt',
    width: 180,
  },
  {
    title: '申请时间',
    key: 'createdAt',
    width: 180,
  },
];
