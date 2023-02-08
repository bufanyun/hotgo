import type { PropType } from 'vue';
import { NCascader } from 'naive-ui';

export const basicProps = {
  ...NCascader.props,
  defaultValue: {
    type: [Number, String, Array],
    default: null,
  },
  value: {
    type: [Number, String, Array],
    default: null,
  },
  dataType: {
    type: String as PropType<'p' | 'pc' | 'pca'>,
    default: 'pca',
  },
  checkStrategy: {
    type: String as PropType<'child' | 'all'>,
    default: 'child',
  },
};
