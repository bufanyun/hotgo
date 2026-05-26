import { NSelect } from 'naive-ui';

export const basicProps = {
  ...NSelect.props,
  defaultValue: {
    type: [Array],
    default: null,
  },
  value: {
    type: [Array],
    default: null,
  },
};
