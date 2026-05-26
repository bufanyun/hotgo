import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  id: number;
  title: string;
  type: number;
  tag: number = 1;
  content: string;
  receiver: number[];
  remark: string;
  sort: number;
  status: number = 1;
  createdBy: number;
  updatedBy: number;
  createdAt: string;
  updatedAt: string;
  deletedAt: string | null;
  readCount: number;
  receiverGroup: Receiver[];

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export class Receiver {
  name: string;
  src: string;
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

// 表单验证规则
export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入消息标题',
  },
};

// 表格搜索表单
export const schemas: FormSchema[] = [
  {
    field: 'type',
    component: 'NSelect',
    label: '消息类型',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择消息类型',
      options: dict.getOption('noticeTypeOptions'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'title',
    component: 'NInput',
    label: '消息标题',
    componentProps: {
      placeholder: '请输入消息标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: [{ message: '请输入消息标题', trigger: ['blur'] }],
  },
  {
    field: 'content',
    component: 'NInput',
    label: '消息内容',
    componentProps: {
      placeholder: '请输入消息内容关键词',
      showButton: false,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择类型',
      options: dict.getOption('sys_normal_disable'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}
