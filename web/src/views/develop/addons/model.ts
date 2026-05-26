import { cloneDeep } from 'lodash-es';
import { h, ref } from 'vue';
import { fallbackSrc } from '@/utils/hotgo';
import { isUrl } from '@/utils/is';
import { NIcon, NIconWrapper, NImage, NTag } from 'naive-ui';
import { getIconComponent } from '@/utils/icons';
import { FormSchema } from '@/components/Form';
import { useDictStore } from '@/store/modules/dict';

const dict = useDictStore();

export const genInfoObj = {
  label: '',
  name: '',
  group: 1,
  extend: ['resourcePublic', 'resourceTemplate'],
  version: 'v1.0.0',
  brief: '',
  description: '',
  author: '',
};

export function newState(state) {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(genInfoObj);
}

export const columns = [
  {
    title: '图标',
    key: 'logo',
    width: 80,
    render(row) {
      if (isUrl(row.logo)) {
        return h(NImage, {
          width: 48,
          height: 48,
          src: row.logo,
          fallbackSrc: fallbackSrc(),
          style: {
            width: '48px',
            height: '48px',
            'max-width': '100%',
            'max-height': '100%',
          },
        });
      } else {
        return h(
          NIconWrapper,
          {
            size: 48,
            borderRadius: 8,
          },
          {
            default: () =>
              h(
                NIcon,
                {
                  size: 36,
                  style: {
                    marginTop: '-8px',
                  },
                },
                {
                  default: () => h(getIconComponent(row.logo)),
                }
              ),
          }
        );
      }
    },
  },
  {
    title: '模块名称',
    key: 'name',
    width: 120,
    render(row) {
      return h('div', {
        innerHTML:
          '<div >' + row.label + '<br><span style="opacity: 0.8;">' + row.name + '</span></div>',
      });
    },
  },
  {
    title: '作者',
    key: 'author',
    width: 100,
  },
  {
    title: '分组',
    key: 'groupName',
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
          default: () => row.groupName,
        }
      );
    },
    width: 120,
  },
  {
    title: '简介',
    key: 'brief',
    render(row) {
      return row.brief;
    },
    width: 180,
  },
  {
    title: '详细描述',
    key: 'description',
    width: 300,
    render(row) {
      return h('p', { id: 'app' }, [
        h('div', {
          innerHTML: '<div style="white-space: pre-wrap">' + row.description + '</div>',
        }),
      ]);
    },
  },
  {
    title: '版本',
    key: 'version',
    width: 100,
  },
];

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NInput',
    label: '模块名称',
    componentProps: {
      placeholder: '请输入模块名称或标签',
      onInput: (e: any) => {
        console.log(e);
      },
    },
    rules: [{ trigger: ['blur'] }],
  },
  {
    field: 'group',
    component: 'NSelect',
    label: '分组',
    componentProps: {
      placeholder: '请选择分组',
      options: dict.getOption('addonsGroupOptions'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '安装状态',
    componentProps: {
      placeholder: '请选择状态',
      options: dict.getOption('addonsInstallStatus'),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export function loadOptions() {
  dict.loadOptions(['addonsGroupOptions', 'addonsInstallStatus', 'addonsExtend']);
}
