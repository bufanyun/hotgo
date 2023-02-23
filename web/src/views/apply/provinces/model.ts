import { h, ref } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';
import { Dicts } from '@/api/dict/dict';
import { isNullObject } from '@/utils/is';

export const listColumns = [
  {
    title: '地区ID',
    key: 'id',
  },
  {
    title: '地区名称',
    key: 'title',
  },
  {
    title: '拼音',
    key: 'pinyin',
    render(row) {
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: 'success',
          bordered: false,
        },
        {
          default: () => row.pinyin,
        }
      );
    },
  },
  {
    title: '经度',
    key: 'lng',
  },
  {
    title: '维度',
    key: 'lat',
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.sys_normal_disable, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.sys_normal_disable, row.status),
        }
      );
    },
  },
];

export interface State {
  id: number | null;
  title: string;
  pinyin: string;
  lng: string;
  lat: string;
  pid: number;
  sort: number;
  status: number;
  oldId: number;
}

export const defaultState = {
  id: null,
  title: '',
  pinyin: '',
  lng: '',
  lat: '',
  pid: 0,
  sort: 0,
  status: 1,
  oldId: 0,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}

export const options = ref<Options>({
  sys_normal_disable: [],
});
