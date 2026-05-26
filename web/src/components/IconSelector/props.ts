import type { PropType } from 'vue';
import { NInput } from 'naive-ui';

export const basicProps = {
  ...NInput.props,
  option: {
    type: String as PropType<string>,
    default: 'antd', // ionicons5 | antd
  },
  value: {
    type: String as PropType<string>,
    default: () => '',
  },
};
