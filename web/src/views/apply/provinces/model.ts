import { h } from 'vue';
import { NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { useDictStore } from '@/store/modules/dict';
import { renderOptionTag } from '@/utils';

const dict = useDictStore();

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
      return renderOptionTag('sys_normal_disable', row.status);
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

export function loadOptions() {
  dict.loadOptions(['sys_normal_disable']);
}
