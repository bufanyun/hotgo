import { cloneDeep } from 'lodash-es';

export class State {
  id = 0;
  name = '';
  key = '';
  dataScope = 1;
  customDept: any[] = [];
  pid = 0;
  level = 1;
  tree = '';
  remark = '';
  sort = 0;
  status = 1;
  createdAt = '';
  updatedAt = '';
  label = '';
  value = 0;
  children: State[] | null = null;

  constructor(state?: Partial<State>) {
    if (state) {
      Object.keys(this).forEach((key) => {
        if (state[key] !== undefined) {
          this[key] = state[key];
        }
      });
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}
