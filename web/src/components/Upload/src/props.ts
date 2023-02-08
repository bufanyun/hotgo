import type { PropType } from 'vue';
import { NUpload } from 'naive-ui';

export const basicProps = {
  ...NUpload.props,
  fileType: {
    type: String,
    default: 'image',
  },
  accept: {
    type: String,
    default: '.jpg,.png,.jpeg,.svg,.gif,.webp',
  },
  helpText: {
    type: String as PropType<string>,
    default: '',
  },
  maxSize: {
    type: Number as PropType<number>,
    default: 2,
  },
  maxNumber: {
    type: Number as PropType<number>,
    default: Infinity,
  },
  value: {
    type: String as PropType<string>,
    default: () => '',
  },
  values: {
    type: (Array as PropType<string[]>) || (Object as PropType<object>),
    default: () => [],
  },
  width: {
    type: Number as PropType<number>,
    default: 104,
  },
  height: {
    type: Number as PropType<number>,
    default: 104, //建议不小于这个尺寸 太小页面可能显示有异常
  },
};
