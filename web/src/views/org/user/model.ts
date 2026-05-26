import { cloneDeep } from 'lodash-es';
import { ref } from 'vue';
import { getDeptOption } from '@/api/org/dept';
import { getRoleOption } from '@/api/system/role';
import { FormSchema, useForm } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore, Option } from '@/store/modules/dict';

// 增加余额/积分.

const dict = useDictStore();

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
      options: dict.getOption('sys_normal_disable'),
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

export const deptTreeOptions = ref([]);
export const roleTreeOptions = ref([]);

export async function loadOptions() {
  dict.loadOptions(['adminPostOption', 'sys_user_sex', 'sys_normal_disable']);

  getDeptOption().then((res) => {
    if (res.list) {
      deptTreeOptions.value = res.list;
    }
  });

  getRoleOption().then((res) => {
    if (res.list) {
      roleTreeOptions.value = res.list;
      registerRoleTabsOption();
    }
  });
}

// 将角色Tabs选项注册到字典
function registerRoleTabsOption() {
  const items = [{ id: -1, name: '全部' }];
  treeDataToCompressed(items, roleTreeOptions.value);

  const roleTabs: Option[] = [];
  for (const item of items) {
    roleTabs.push(dict.genOption(item.id, item.name));
  }
  dict.setOption('roleTabs', roleTabs);
}

function treeDataToCompressed(items: any[], source: any) {
  for (const i in source) {
    items.push(source[i]);
    source[i].children && source[i].children.length > 0
      ? treeDataToCompressed(items, source[i].children)
      : ''; // 子级递归
  }
  return items;
}
