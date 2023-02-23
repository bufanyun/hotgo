import { cloneDeep } from 'lodash-es';

export const genInfoObj = {
  label: '',
  name: '',
  group: 1,
  version: 'v1.0.0',
  brief: '',
  description: '',
  author: '',
};

export const selectListObj = {
  groupType: [],
  status: [],
};

export function newState(state) {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(genInfoObj);
}
