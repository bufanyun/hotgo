import { FormSchema } from '@/components/Form';
import { ref } from 'vue';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export class State {
  id = 0;
  reqId = '';
  appId = '';
  merchantId = 0;
  memberId = 0;
  method = '';
  module = '';
  url = '';
  getData: Record<string, any> | null = null;
  postData: Record<string, string[]> | null = null;
  headerData: Record<string, string[]> | null = null;
  ip = '';
  provinceId = 0;
  cityId = 0;
  errorCode = 0;
  errorMsg = '';
  errorData: string[] = [];
  userAgent = '';
  takeUpTime = 0;
  timestamp = 0;
  status = 0;
  createdAt = '';
  updatedAt = '';
  cityLabel = '';
  tags = '';
  summary = '';
  description = '';
}

export const schemas = ref<FormSchema[]>([
  {
    field: 'reqId',
    component: 'NInput',
    label: '链路ID',
    componentProps: {
      placeholder: '请输入链路ID',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'complexMemberId',
    component: 'ComplexMemberPicker',
    label: '操作人',
    componentProps: {
      placeholder: '请选择操作人',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'url',
    component: 'NInput',
    label: '接口路径',
    componentProps: {
      placeholder: '请输入接口路径',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'keyword',
    component: 'NInput',
    label: '关键词',
    componentProps: {
      placeholder: '请输入关键词',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'ip',
    component: 'NInput',
    label: '访问IP',
    componentProps: {
      placeholder: '请输入IP地址',
      onInput: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'method',
    component: 'NSelect',
    label: '请求方式',
    componentProps: {
      placeholder: '请选择请求方式',
      options: dict.getOption('HTTPMethod'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '访问时间',
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
    field: 'takeUpTime',
    component: 'NSelect',
    label: '请求耗时',
    componentProps: {
      placeholder: '请选择请求耗时',
      options: dict.getOption('HTTPHandlerTime'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'errorCode',
    component: 'NSelect',
    label: '状态码',
    labelMessage: '支持填入自定义状态码',
    componentProps: {
      placeholder: '请选择状态码',
      options: dict.getOption('HTTPApiCode'),
      filterable: true,
      tag: true,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['HTTPMethod', 'HTTPHandlerTime', 'HTTPApiCode']);
}
