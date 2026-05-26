import { State } from '@/views/log/log/model';

export const basicProps = {
  state: {
    type: State,
    default: null,
  },
  column: {
    type: String,
    default: '',
  },
};
