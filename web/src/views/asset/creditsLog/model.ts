import { ref } from 'vue';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';
import { renderOptionTag, renderPopoverMemberSumma } from '@/utils';

const dict = useDictStore();

export const schemas = ref<FormSchema[]>([
  {
    field: 'complexMemberId',
    component: 'ComplexMemberPicker',
    label: '选择用户',
    componentProps: {
      placeholder: '请选择用户',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'creditGroup',
    component: 'NSelect',
    label: '变动组别',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择变动组别',
      options: dict.getOption('creditGroup'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'remark',
    component: 'NInput',
    label: '备注',
    componentProps: {
      placeholder: '请输入备注',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'ip',
    component: 'NInput',
    label: '操作人IP',
    componentProps: {
      placeholder: '请输入操作人IP',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'id',
    component: 'NInput',
    label: '变动ID',
    componentProps: {
      placeholder: '请输入变动ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '变动时间',
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

export const columns = [
  {
    title: '变动ID',
    key: 'id',
    width: 100,
  },
  {
    title: '用户',
    key: 'memberId',
    width: 100,
    render(row) {
      return renderPopoverMemberSumma(row.memberBySumma);
    },
  },
  {
    title: '变动类型',
    key: 'creditType',
    render(row) {
      return renderOptionTag('creditType', row.creditType);
    },
    width: 150,
  },
  {
    title: '变动组别',
    key: 'creditGroup',
    render(row) {
      return renderOptionTag('creditGroup', row.creditGroup);
    },
    width: 150,
  },
  {
    title: '变动前',
    key: 'beforeNum',
    width: 100,
    render(row) {
      return Number(row.beforeNum).toFixed(2);
    },
  },
  {
    title: '变动数量',
    key: 'num',
    width: 100,
    render(row) {
      return Number(row.num).toFixed(2);
    },
  },
  {
    title: '变动后',
    key: 'afterNum',
    width: 100,
    render(row) {
      return Number(row.afterNum).toFixed(2);
    },
  },
  {
    title: '备注',
    key: 'remark',
    width: 200,
  },
  {
    title: '操作人IP',
    key: 'ip',
    width: 150,
  },
  {
    title: '关联ID',
    key: 'mapId',
    width: 100,
    render(row) {
      if (row.mapId === 0) {
        return '-';
      }
      return row.mapId;
    },
  },
  {
    title: '变动时间',
    key: 'createdAt',
    width: 180,
  },
];

export function loadOptions() {
  dict.loadOptions(['creditType', 'creditGroup']);
}
