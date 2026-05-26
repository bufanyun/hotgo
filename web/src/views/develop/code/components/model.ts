import { cloneDeep } from 'lodash-es';
import { isJsonString } from '@/utils/is';

export const genFileObj = {
  meth: 1,
  content: '',
  path: '',
  required: true,
};

export interface joinAttr {
  uuid: string;
  linkTable: string;
  alias: string;
  linkMode: number;
  field: string;
  masterField: string;
  columns: any;
}

export const genInfoObj = {
  id: 0,
  genType: 10,
  genTemplate: null,
  varName: '',
  options: {
    headOps: ['add', 'batchDel', 'export'],
    columnOps: ['edit', 'del', 'view', 'status', 'check'],
    autoOps: ['genMenuPermissions', 'runDao', 'runService'],
    join: [],
    menu: {
      pid: 0,
      icon: 'MenuOutlined',
      sort: 0,
    },
    tree: {
      titleColumn: null,
      styleType: 1,
    },
    funcDict: {
      valueColumn: null,
      labelColumn: null,
    },
    presetStep: {
      formGridCols: 1,
    },
  },
  dbName: '',
  tableName: '',
  tableComment: '',
  daoName: '',
  masterColumns: [],
  addonName: null,
  status: 2,
  createdAt: '',
  updatedAt: '',
};

export const selectListObj = {
  db: [],
  genType: [],
  status: [],
  tables: [],
  formMode: [],
  formRole: [],
  dictMode: [],
  whereMode: [],
  buildMeth: [],
  tableAlign: [],
  treeStyleType: [],
};

export function newState(state) {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(genInfoObj);
}

export const formGridColsOptions = [
  {
    value: 1,
    label: '一行一列',
  },
  {
    value: 2,
    label: '一行两列',
  },
  {
    value: 3,
    label: '一行三列',
  },
  {
    value: 4,
    label: '一行四列',
  },
];

export const formGridSpanOptions = [
  {
    value: 1,
    label: '占一列位置',
  },
  {
    value: 2,
    label: '占两列位置',
  },
  {
    value: 3,
    label: '占三列位置',
  },
  {
    value: 4,
    label: '占四列位置',
  },
];

// 格式化列字段
export function formatColumns(columns: any) {
  if (columns === undefined || columns.length === 0) {
    columns = [];
  }

  if (isJsonString(columns)) {
    columns = JSON.parse(columns);
  }

  if (columns.length > 0) {
    for (let i = 0; i < columns.length; i++) {
      if (!columns[i].formGridSpan) {
        columns[i].formGridSpan = 1;
      }
      if (!columns[i].align) {
        columns[i].align = 'left';
      }
      if (!columns[i].width || columns[i].width < 1) {
        columns[i].width = null;
      }
    }
  }
  return columns;
}
