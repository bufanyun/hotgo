import { h, ref } from 'vue';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { NTag } from 'naive-ui';
import { useDictStore } from '@/store/modules/dict';
import { renderOptionTag } from '@/utils';

const dict = useDictStore();

export const schemas = ref<FormSchema[]>([
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
  },
  {
    field: 'loginIp',
    component: 'NInput',
    label: 'IP地址',
    componentProps: {
      placeholder: '请输入IP地址',
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
      options: dict.getOption('sys_login_status'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'loginAt',
    component: 'NDatePicker',
    label: '登录时间',
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
    title: '记录ID',
    key: 'id',
    width: 80,
  },
  {
    title: '用户名',
    key: 'username',
    width: 120,
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'info',
          bordered: false,
        },
        {
          default: () => row.username,
        }
      );
    },
  },
  {
    title: '登录IP',
    key: 'loginIp',
    width: 160,
  },
  {
    title: 'IP归属地',
    key: 'cityLabel',
    width: 200,
  },
  {
    title: '浏览器',
    key: 'browser',
    width: 200,
  },
  {
    title: '操作系统',
    key: 'os',
    width: 150,
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      return renderOptionTag('sys_login_status', row.status);
    },
    width: 150,
  },
  {
    title: '提示信息',
    key: 'errMsg',
    render(row) {
      if (row.errMsg !== '') {
        return row.errMsg;
      }

      if (row.status === 1) {
        return '登录成功';
      }
      return ``;
    },
    width: 200,
  },
  {
    title: '登录时间',
    key: 'loginAt',
    width: 180,
  },
];

// 加载字典数据选项
export function loadOptions() {
  dict.loadOptions(['sys_login_status']);
}
