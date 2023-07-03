import { h, ref } from 'vue';
import { NTag, NButton } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isNullObject } from '@/utils/is';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { format } from 'date-fns';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
import { renderIcon, renderTooltip } from '@/utils';
import { HelpCircleOutline } from '@vicons/ionicons5';

export interface State {
  id: number;
  env: string;
  traceid: string;
  levelFormat: string;
  content: string;
  stack: any;
  line: string;
  triggerNs: number;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  env: '',
  traceid: '',
  levelFormat: '',
  content: '',
  stack: null,
  line: '',
  triggerNs: 0,
  status: 1,
  createdAt: '',
  updatedAt: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  sys_normal_disable: [],
  sys_log_type: [],
});

export const rules = {};

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
    field: 'levelFormat',
    component: 'NSelect',
    label: '日志级别',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择日志级别',
      options: [],
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
      if (isNullObject(row.levelFormat)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_log_type, row.levelFormat),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_log_type, row.levelFormat),
        }
      );
    },
    width: 120,
  },
  {
    title: '日志内容',
    key: 'content',
    width: 320,
  },
  {
    title: '调用行',
    key: 'line',
    width: 150,
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

async function loadOptions() {
  options.value = await Dicts({
    types: ['sys_normal_disable', 'sys_log_type'],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.sys_normal_disable;
        break;
      case 'levelFormat':
        item.componentProps.options = options.value.sys_log_type;
        break;
    }
  }
}

await loadOptions();
