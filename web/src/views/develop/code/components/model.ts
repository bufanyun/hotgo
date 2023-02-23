import { cloneDeep } from 'lodash-es';

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
    columnOps: ['edit', 'del', 'view', 'status', 'switch', 'check'],
    autoOps: ['genMenuPermissions', 'runDao', 'runService'],
    join: [],
    menu: {
      pid: 0,
      icon: 'MenuOutlined',
      sort: 0,
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
};

export function newState(state) {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(genInfoObj);
}
