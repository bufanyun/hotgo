import { h, ref } from 'vue';
import { NTag, NButton } from 'naive-ui';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { format } from 'date-fns';
import { renderIcon, renderOptionTag, renderTooltip } from '@/utils';
import { HelpCircleOutline } from '@vicons/ionicons5';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export const schemas = ref<FormSchema[]>([
  {
    field: 'traceId',
    component: 'NInput',
    label: '链路ID',
    componentProps: {
      placeholder: '请输入链路ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'content',
    component: 'NInput',
    label: '日志内容',
    componentProps: {
      placeholder: '请输入日志内容关键词',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'levelFormat',
    component: 'NSelect',
    label: '日志级别',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择日志级别',
      options: dict.getOption('sys_log_type'),
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

export const columns = [
  {
    title: '日志ID',
    key: 'id',
    width: 80,
  },
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
          { default: () => '链路ID', icon: renderIcon(HelpCircleOutline) }
        ),
        'hotgo默认支持链路追踪，如果是web请求产生的日志则还可以关联对应的访问日志'
      );
    },
    key: 'traceId',
    width: 280,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'default',
          bordered: false,
          checkable: true,
        },
        {
          default: () => row.traceId,
        }
      );
    },
  },
  {
    title: '日志级别',
    key: 'levelFormat',
    render(row) {
      return renderOptionTag('sys_log_type', row.levelFormat);
    },
    width: 120,
  },
  {
    title: '日志内容',
    key: 'content',
    width: 320,
  },
  {
    title: '触发时间',
    key: 'triggerNs',
    width: 200,
    render(row) {
      if (row.triggerNs <= 0) {
        return '-';
      }
      return format(new Date(row.triggerNs / 1000000), 'yyyy-MM-dd HH:mm:ss.SSS');
    },
  },
  {
    title: '记录时间',
    key: 'createdAt',
    width: 150,
  },
];

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable', 'sys_log_type']);
}
