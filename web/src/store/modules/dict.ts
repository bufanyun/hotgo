import { computed, ComputedRef } from 'vue';
import { defineStore } from 'pinia';
import { createStorage, storage } from '@/utils/Storage';
import { store } from '@/store';
import { CURRENT_DICT } from '@/store/mutation-types';
import { Dicts } from '@/api/dict/dict';
import { SelectMixedOption } from 'naive-ui/es/select/src/interface';
import { localDict } from '@/enums/localDictEnum';

const Storage = createStorage({ storage: localStorage });

// 默认字典类型，通过后台字典管理添加的
export type DefaultDictType = 'sys_normal_disable' | 'sys_switch' | 'sys_user_sex';

// 枚举字典类型
export type EnumsDictType = 'creditType' | 'creditGroup' | 'deptType';

// 方法字典类型
export type FuncDictType = 'testCategoryOption';

// 可以把常用的字典类型加入进来，这样会有IDE提示
export type DictType = DefaultDictType | EnumsDictType | FuncDictType | string;

export type DictListClass = 'default' | 'error' | 'primary' | 'info' | 'success' | 'warning';

export interface Option {
  label: string;
  value: string | number;
  key: string | number;
  listClass: DictListClass;
  extra: any;
}

export class DictOptions {
  [name: DictType]: Option[];
}

export type IDictState = {
  options: DictOptions;
};

export const useDictStore = defineStore({
  id: 'app-dict',
  state: (): IDictState => ({
    options: Storage.get(CURRENT_DICT, new DictOptions()),
  }),
  getters: {},
  actions: {
    // 加载字典数据选项
    loadOptions(types: DictType[]) {
      Dicts({
        types: types,
      }).then((res) => {
        if (!res) {
          return;
        }
        Object.keys(res).forEach((key) => {
          if (res[key]) {
            this.options[key] = res[key] as unknown as Option[];
          }
        });
        storage.set(CURRENT_DICT, this.options);
      });
    },

    // 设置指定类型的字典选项
    setOption(type: DictType, opts: Option[]) {
      this.options[type] = opts;
      storage.set(CURRENT_DICT, this.options);
    },

    // 获取指定类型的字典选项
    getOption(type: DictType): ComputedRef<SelectMixedOption[]> {
      return computed(() => {
        const opts = this.checkOptionValue(type) ?? [];
        return opts as unknown as SelectMixedOption[];
      });
    },

    // 获取指定类型的字典选项
    getOptionUnRef(type: DictType): any[] {
      return this.checkOptionValue(type) ?? [];
    },

    // 是否存在选项，不存在返回null，存在就返回选项
    checkOptionValue(type: DictType): null | Option[] {
      // 本地字典
      const local = localDict[type] ?? [];
      if (local && local.length > 0) {
        return local;
      }
      const opts = this.options[type] ?? [];
      if (!opts || opts?.length === 0) {
        return null;
      }
      return opts;
    },

    // 获取选项名称
    getLabel(type: DictType, value: any) {
      if (value === null || value === undefined) {
        return ``;
      }
      const opts = this.checkOptionValue(type);
      if (!opts) {
        return ``;
      }
      for (const item of opts) {
        if (item.value.toString() === value.toString()) {
          return item.label;
        }
      }
      return ``;
    },

    // 获取选项标签类型
    getType(type: DictType, value: any): DictListClass {
      if (value === null || value === undefined) {
        return `default`;
      }
      const opts = this.checkOptionValue(type);
      if (!opts) {
        return `default`;
      }
      for (const item of opts) {
        if (item.value.toString() === value.toString()) {
          return item.listClass;
        }
      }
      return `default`;
    },

    // 获取选项额外的数据配置
    getExtra(type: DictType, value: any): any {
      if (value === null || value === undefined) {
        return null;
      }
      const opts = this.checkOptionValue(type);
      if (!opts) {
        return null;
      }
      for (const item of opts) {
        if (item.value.toString() === value.toString()) {
          return item?.extra;
        }
      }
      return null;
    },

    // 是否存在指定选项值
    hasValue(type: DictType, value: any): boolean {
      if (value === null || value === undefined) {
        return false;
      }
      const opts = this.checkOptionValue(type);
      if (!opts) {
        return false;
      }
      for (const item of opts) {
        if (item.value.toString() === value.toString()) {
          return true;
        }
      }
      return false;
    },

    // 获取指定类型的可选项数量
    getOptionLength(type: DictType): number {
      const opts = this.checkOptionValue(type);
      if (!opts) {
        return 0;
      }
      return opts?.length;
    },

    // 生成单个选项
    genOption(key: any, label: string, type = 'default', extra: any = null): Option {
      return <Option>{
        label: label,
        value: key,
        key: key,
        listClass: type,
        extra: extra,
      };
    },
  },
});

// Need to be used outside the setup
export function useDictStoreWidthOut() {
  return useDictStore(store);
}
