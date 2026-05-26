// @ts-ignore
import onerrorImg from '@/assets/images/onerror.png';
import { usePermission } from '@/hooks/web/usePermission';
import { ActionItem } from '@/components/Table';
import { isBoolean, isFunction } from '@/utils/is';
import { PermissionsEnum } from '@/enums/permissionsEnum';

/**
 * @deprecated 这个方法将来未来版本完全弃用，请统一使用useDictStore()来维护字典选项
 */
export interface Option {
  label: string;
  value: string | number;
  key: string | number;
  // type: string;
  listClass: 'default' | 'error' | 'primary' | 'info' | 'success' | 'warning';
}

/**
 * @deprecated 这个方法将来未来版本完全弃用，请统一使用useDictStore()来维护字典选项
 */
export interface Options {
  [name: string]: Option[];
}

/**
 * 获取选项名称
 * @deprecated 这个方法将来未来版本完全弃用，请统一使用useDictStore()来维护字典选项
 */
export function getOptionLabel(options: Option[], value) {
  if (options === undefined || options?.length === 0) {
    return `unknown`;
  }
  for (const item of options) {
    if (item.value == value) {
      return item.label;
    }
  }

  return `unknown`;
}

/**
 * 获取选项标签
 * @deprecated 这个方法将来未来版本完全弃用，请统一使用useDictStore()来维护字典选项
 */
export function getOptionTag(options: Option[], value) {
  if (options === undefined || options?.length === 0) {
    return 'default';
  }
  for (const item of options) {
    if (item.value == value) {
      return item.listClass;
    }
  }

  return 'default';
}

// 自适应模板宽度
export function adaModalWidth(def = 840) {
  const val = document.body.clientWidth;
  if (val <= def) {
    return '100%';
  } else {
    return def + 'px';
  }
}

export interface TableColumn {
  width?: number | string;
  auth?: PermissionsEnum | PermissionsEnum[] | string | string[];
  ifShow?: boolean | ((action: ActionItem) => boolean);
}

// 自适应表格组件横向滑动可见宽度
export function adaTableScrollX(columns: TableColumn[] = [], actionWidth: number) {
  const { hasPermission } = usePermission();

  let x = 50; // 勾选列宽度
  columns = columns.filter((column) => {
    return hasPermission(column.auth as string[]) && isIfShow(column);
  });
  for (const column of columns) {
    if (column.width && Number(column.width) >= 1) {
      x += Number(column.width);
    } else {
      x += 100; // 默认列宽度
    }
  }
  x += actionWidth;
  return x;
}

export function isIfShow(action: ActionItem): boolean {
  let isIfShow = true;
  const ifShow = action.ifShow;
  if (isBoolean(ifShow)) {
    isIfShow = ifShow;
  }
  if (isFunction(ifShow)) {
    isIfShow = ifShow(action);
  }
  return isIfShow;
}

// 图片加载失败显示自定义默认图片(缺省图) img标签使用
export function errorImg(e: any): void {
  e.target.src = onerrorImg;
  e.target.onerror = null;
}

// 图片加载失败显示自定义默认图片(缺省图) NImage组件使用
export function fallbackSrc(): string {
  return onerrorImg;
}

export function timeFix() {
  const time = new Date();
  const hour = time.getHours();
  return hour < 9
    ? '早上好'
    : hour <= 11
    ? '上午好'
    : hour <= 13
    ? '中午好'
    : hour < 20
    ? '下午好'
    : '晚上好';
}

// 随机浅色
export function rdmLightRgbColor(): string {
  const letters = '456789ABCDEF';
  let color = '#';
  for (let i = 0; i < 6; i++) {
    if (i === 0) {
      color += 'F'; // 确保第一个字符较亮
    } else {
      color += letters[Math.floor(Math.random() * letters.length)];
    }
  }
  return color;
}

// 将列表数据转为树形数据
export function convertListToTree(list: any[], idField = 'id', pidField = 'pid') {
  if (!list || list.length === 0) {
    return [];
  }
  const min = list.reduce((prev, current) => (prev[pidField] < current[pidField] ? prev : current));

  const map = list.reduce((acc, item) => {
    acc[item[idField]] = { ...item, children: [] };
    return acc;
  }, {});

  list.forEach((item) => {
    if (item[pidField] !== min[pidField]) {
      if (!map[item[pidField]]) {
        map[item[pidField]] = {};
      }

      if (!map[item[pidField]].children) {
        map[item[pidField]].children = [];
      }
      map[item[pidField]].children.push(map[item[idField]]);
    }
  });
  return list.filter((item) => item[pidField] === min[pidField]).map((item) => map[item[idField]]);
}

// 从树选项中获取所有key
export function getTreeKeys(data: any[], idField = 'id') {
  const keys: any = [];
  if (!data || data.length === 0) {
    return keys;
  }
  data.map((item: any) => {
    keys.push(item[idField]);
    if (item.children && item.children.length) {
      keys.push(...getTreeKeys(item.children));
    }
  });
  return keys;
}
