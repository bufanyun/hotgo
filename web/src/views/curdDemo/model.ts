import { h, ref } from 'vue';
import { NSwitch } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { renderImage, renderFile, renderPopoverMemberSumma, MemberSumma } from '@/utils';
import { Switch } from '@/api/curdDemo';
import { useDictStore } from '@/store/modules/dict';
import { usePermission } from '@/hooks/web/usePermission';

const dict = useDictStore();
const { hasPermission } = usePermission();
const $message = window['$message'];

export class State {

  public id = 0;//ID
  public categoryId = null;//测试分类
  public title = '';//标题
  public description = '';//描述
  public content = '';//内容
  public image = '';//单图
  public attachfile = '';//附件
  public cityId = null;//所在城市
  public switch = 2;//显示开关
  public sort = 0;//排序
  public status = 1;//状态
  public createdBy = 0;//创建者
  public createdBySumma?: null | MemberSumma = null;//创建者摘要信息
  public updatedBy = 0;//更新者
  public updatedBySumma?: null | MemberSumma = null;//更新者摘要信息
  public deletedBy = 0;//删除者
  public deletedBySumma?: null | MemberSumma = null;//删除者摘要信息
  public createdAt = '';//创建时间
  public updatedAt = '';//修改时间
  public deletedAt = '';//删除时间
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
  description: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入描述',
  },
  content: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入内容',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInput',
    label: 'ID',
    componentProps: {
      placeholder: '请输入ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
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
    field: 'description',
    component: 'NInput',
    label: '描述',
    componentProps: {
      placeholder: '请输入描述',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdBy',
    component: 'NInput',
    label: '创建者',
    componentProps: {
      placeholder: '请输入ID|用户名|姓名|手机号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'deletedBy',
    component: 'NInput',
    label: '删除者',
    componentProps: {
      placeholder: '请输入ID|用户名|姓名|手机号',
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
  {
    field: 'testCategoryName',
    component: 'NInput',
    label: '关联分类',
    componentProps: {
      placeholder: '请输入关联分类',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: 'ID',
    key: 'id',
    align: 'left',
    width: 50,
  },
  {
    title: '标题',
    key: 'title',
    align: 'left',
    width: 150,
  },
  {
    title: '描述',
    key: 'description',
    align: 'left',
    width: 300,
  },
  {
    title: '单图',
    key: 'image',
    align: 'left',
    width: 100,
    render(row: State) {
      return renderImage(row.image);
    },
  },
  {
    title: '附件',
    key: 'attachfile',
    align: 'left',
    width: 100,
    render(row: State) {
      return renderFile(row.attachfile);
    },
  },
  {
    title: '显示开关',
    key: 'switch',
    align: 'left',
    width: 150,
    render(row: State) {
      return h(NSwitch, {
        value: row.switch === 1,
        checked: '开启',
        unchecked: '关闭',
        disabled: !hasPermission(['/curdDemo/switch']),
        onUpdateValue: function (e) {
          console.log('onUpdateValue e:' + JSON.stringify(e));
          row.switch = e ? 1 : 2;
          Switch({ id: row.id, key: 'switch', value: row.switch }).then((_res) => {
            $message.success('操作成功');
          });
        },
      });
    },
  },
  {
    title: '排序',
    key: 'sort',
    align: 'left',
    width: 100,
  },
  {
    title: '创建者',
    key: 'createdBy',
    align: 'left',
    width: 150,
    render(row: State) {
      return renderPopoverMemberSumma(row.createdBySumma);
    },
  },
  {
    title: '更新者',
    key: 'updatedBy',
    align: 'left',
    width: 150,
    render(row: State) {
      return renderPopoverMemberSumma(row.updatedBySumma);
    },
  },
  {
    title: '删除者',
    key: 'deletedBy',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderPopoverMemberSumma(row.deletedBySumma);
    },
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: 180,
  },
  {
    title: '修改时间',
    key: 'updatedAt',
    align: 'left',
    width: 180,
  },
  {
    title: '关联分类',
    key: 'testCategoryName',
    align: 'left',
    width: 100,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_normal_disable', 'testCategoryOption']);
}