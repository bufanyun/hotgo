import { cloneDeep } from 'lodash-es';
import { ref } from 'vue';
import { getDeptOption } from '@/api/org/dept';
import { getRoleOption } from '@/api/system/role';
import { getPostOption } from '@/api/org/post';
import { FormSchema, useForm } from '@/components/Form';
import { statusOptions } from '@/enums/optionsiEnum';
import { defRangeShortcuts } from '@/utils/dateUtil';

// 增加余额/积分.

export interface addState {
  id: number;
  username: string;
  realName: string;
  integral: number;
  balance: number;
  operateMode: number;
  num: number | null;
}

export const addDefaultState = {
  id: 0,
  realName: '',
  username: '',
  integral: 0,
  balance: 0,
  operateMode: 1,
  num: null,
};

export function addNewState(state: addState | null): addState {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(addDefaultState);
}

export const operateModes = [
  {
    value: 1,
    label: '加款',
  },
  {
    value: 2,
    label: '扣款',
  },
];

export const addRules = {};

// 用户列表.

export const defaultState = {
  id: 0,
  roleId: null,
  realName: '',
  username: '',
  password: '',
  deptId: null,
  postIds: null,
  mobile: '',
  email: '',
  sex: 1,
  leader: '',
  phone: '',
  sort: 0,
  status: 1,
  createdAt: '',
  updatedAt: '',
};

export interface State {
  id: number;
  roleId: number | null;
  realName: string;
  username: string;
  password: string;
  deptId: number | null;
  postIds: any;
  mobile: string;
  email: string;
  sex: number;
  leader: string;
  phone: string;
  sort: number;
  status: number;
  createdAt: string;
  updatedAt: string;
}

const schemas: FormSchema[] = [
  {
    field: 'username',
    component: 'NInput',
    label: '用户名',
    componentProps: {
      placeholder: '请输入用户名',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
    rules: [{ message: '请输入用户名', trigger: ['blur'] }],
  },
  {
    field: 'realName',
    component: 'NInput',
    label: '姓名',
    componentProps: {
      placeholder: '请输入姓名',
      showButton: false,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'mobile',
    component: 'NInputNumber',
    label: '手机号',
    componentProps: {
      placeholder: '请输入手机号码',
      showButton: false,
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'email',
    component: 'NInput',
    label: '邮箱',
    componentProps: {
      placeholder: '请输入邮箱地址',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'id',
    component: 'NInput',
    label: '用户ID',
    componentProps: {
      placeholder: '请输入用户ID',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'pid',
    component: 'NInput',
    label: '上级ID',
    componentProps: {
      placeholder: '请输入上级用户ID',
      onInput: (e: any) => {
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
      options: statusOptions,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'created_at',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
];

export const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas,
});

export const options = ref<any>({
  role: [],
  roleTabs: [{ id: -1, name: '全部' }],
  dept: [],
  post: [],
});

export async function loadOptions() {
  const dept = await getDeptOption();
  if (dept.list) {
    options.value.dept = dept.list;
  }

  const role = await getRoleOption();
  if (role.list) {
    options.value.role = role.list;
    options.value.roleTabs = [{ id: -1, name: '全部' }];
    treeDataToCompressed(role.list);
  }

  const post = await getPostOption();
  if (post.list && post.list.length > 0) {
    for (let i = 0; i < post.list.length; i++) {
      post.list[i].label = post.list[i].name;
      post.list[i].value = post.list[i].id;
    }
    options.value.post = post.list;
  }
}

function treeDataToCompressed(source) {
  for (const i in source) {
    options.value.roleTabs.push(source[i]);
    source[i].children && source[i].children.length > 0
      ? treeDataToCompressed(source[i].children)
      : ''; // 子级递归
  }
  return options.value.roleTabs;
}
