import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderOptionTag, renderPopoverMemberSumma, MemberSumma } from '@/utils';
import { TreeOption } from '@/api/normalTreeDemo';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  public title = ''; // 标题
  public id = 0; // ID
  public pid = 0; // 上级
  public level = 1; // 关系树级别
  public tree = ''; // 关系树
  public categoryId = null; // 测试分类
  public description = ''; // 描述
  public sort = 0; // 排序
  public status = 1; // 状态
  public createdBy = 0; // 创建者
  public createdBySumma?: null | MemberSumma = null; // 创建者摘要信息
  public updatedBy = 0; // 更新者
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间
  public deletedAt = ''; // 删除时间

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
export const rules = {
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标题',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: {
      placeholder: '请输入标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'categoryId',
    component: 'NSelect',
    label: '测试分类',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择测试分类',
      options: dict.getOption('testCategoryOption'),
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
      placeholder: '请选择状态',
      options: dict.getOption('sys_normal_disable'),
      onUpdateValue: (e: any) => {
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

// 表格列
export const columns = [
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 200,
  },
  {
    title: '测试分类',
    key: 'categoryId',
    align: 'left',
    width: 100,
    render(row: State) {
      return renderOptionTag('testCategoryOption', row.categoryId);
    },
  },
  {
    title: '描述',
    key: 'description',
    align: 'left',
    width: 300,
  },
  {
    title: '状态',
    key: 'status',
    align: 'left',
    width: 150,
    render(row: State) {
      return renderOptionTag('sys_normal_disable', row.status);
    },
  },
  {
    title: '创建者',
    key: 'createdBy',
    align: 'left',
    width: 100,
    render(row: State) {
      return renderPopoverMemberSumma(row.createdBySumma);
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable', 'testCategoryOption']);
}

// 关系树选项
export const treeOption = ref([]);

// 加载关系树选项
export function loadTreeOption() {
  TreeOption().then((res) => {
    treeOption.value = res;
  });
}