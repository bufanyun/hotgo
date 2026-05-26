import { h, ref } from 'vue';
import { NTag, NButton } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { renderTooltip, renderIcon, renderOptionTag } from '@/utils';
import { HelpCircleOutline } from '@vicons/ionicons5';
import { TreeOption } from '@/api/org/dept';
import { useDictStore } from '@/store/modules/dict';
import type { FormRules } from 'naive-ui/es/form/src/interface';

const dict = useDictStore();

export class State {
  public id = 0; // 部门ID
  public pid = 0; // 父部门ID
  public name = ''; // 部门名称
  public code = ''; // 部门编码
  public type = 'company'; // 部门类型
  public leader = ''; // 负责人
  public phone = ''; // 联系电话
  public email = ''; // 邮箱
  public level = 0; // 关系树等级
  public tree = ''; // 关系树
  public sort = 0; // 排序
  public status = 1; // 部门状态
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 更新时间

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
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

// 表单验证规则
export const rules: FormRules = {
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '部门名称',
    componentProps: {
      placeholder: '请输入部门名称',
      onInput: (e: any) => {
        console.log(e);
      },
    },
    rules: [{ message: '请输入部门名称', trigger: ['blur'] }],
  },
  {
    field: 'code',
    component: 'NInput',
    label: '部门编码',
    componentProps: {
      placeholder: '请输入部门编码',
      showButton: false,
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'leader',
    component: 'NInput',
    label: '负责人',
    componentProps: {
      placeholder: '请输入负责人',
      showButton: false,
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
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
]);

export const filterIds = ref([]);

// 表格列
export const columns = [
  {
    title(_column) {
      return renderTooltip(
        h(
          NButton,
          {
            strong: true,
            size: 'small',
            text: true,
            iconPlacement: 'right',
          },
          { default: () => '部门名称', icon: renderIcon(HelpCircleOutline) }
        ),
        '支持上下级部门，点击列表中左侧 > 按钮可展开下级部门列表'
      );
    },
    key: 'name',
    render(row) {
      const filter = filterIds.value.includes(row.id as never);
      return h(
        NTag,
        {
          type: 'info',
          checkable: filter,
          checked: filter,
        },
        {
          default: () => row.name,
        }
      );
    },
    width: 200,
  },
  {
    title: '部门编码',
    key: 'code',
    width: 100,
  },
  {
    title: '部门类型',
    key: 'type',
    align: 'left',
    width: 100,
    render(row) {
      return renderOptionTag('deptType', row.type);
    },
  },
  {
    title: '负责人',
    key: 'leader',
    width: 100,
  },
  {
    title: '联系电话',
    key: 'phone',
    width: 150,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 80,
    render(row) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 150,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable', 'deptType']);
}

// 关系树选项
export const treeOption = ref([]);

// 加载关系树选项
export function loadTreeOption() {
  TreeOption().then((res) => {
    treeOption.value = res;
  });
}
