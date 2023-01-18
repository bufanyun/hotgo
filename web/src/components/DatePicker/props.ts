import type { PropType } from 'vue';
import { NDatePicker } from 'naive-ui';

export const basicProps = {
  ...NDatePicker.props,
  formValue: {
    type: String as PropType<string> | undefined | Date,
    default: () => '',
  },
  startValue: {
    type: String as PropType<string> | undefined | Date,
    default: () => '',
  },
  endValue: {
    type: String as PropType<string> | undefined | Date,
    default: () => '',
  },
};
