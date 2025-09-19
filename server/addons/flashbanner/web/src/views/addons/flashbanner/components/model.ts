import { h, render } from 'vue';
import { FormSchema, useForm } from '@/components/Form';
import { NImage } from 'naive-ui';
import { fallbackSrc } from '@/utils/hotgo';

// **********查询表单********
const detailSchemas: FormSchema[] = [
  {
    field: 'name',
    component: 'NInput',
    label: '图片名称',
    defaultValue: null,
    componentProps: {
      placeholder: '请输入图片名称',
    },
  },
];
export const [register, {}] = useForm({
  gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
  labelWidth: 80,
  schemas: detailSchemas,
});

// *********表格**********
export const defaultColumns = [
  {
    title: 'ID',
    key: 'id',
    width: 100,
  },
  {
    title: '图片名称',
    key: 'name',
  },
  {
    title: '图片',
    key: 'cover',
    render(row) {
      if (row.cover !== '') {
        return h(NImage, {
          width: 40,
          height: 40,
          src: row.cover,
          fallbackSrc: fallbackSrc(),
          style: {
            width: '40px',
            height: '40px',
            'max-width': '100%',
            'max-height': '100%',
          },
        });
      } else {
        return '暂无图片'
      }
    },
  },
  {
    title: '链接地址',
    key: 'link',
    render(row) {
      return h('a', { href: row.link, target: '_blank' }, row.link);
    }
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
  },
];

// *********编辑表单规则***********
export const rules = {
  name: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入图片名称',
  },
  cover: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请上传图片',
  },
};