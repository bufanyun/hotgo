/**
 * 树形数组转平级
 * @date 2023-01-09
 * @returns {array}
 */
export const tree2FlatArray = (tree: any[], childrenKey: string = "children") => {
  return tree.reduce((arr, item) => {
    var children = item?.[childrenKey] ?? [];
    var obj = item;
    delete obj[childrenKey];
    //解构赋值+默认值
    return arr.concat([obj], tree2FlatArray(children)); //children部分进行递归
  }, []);
};
/**
 * 树形数组转对象
 * @date 2023-01-09
 * @returns {object}
 */
export const tree2Objects = (
  tree: any[],
  key: string = "id",
  childrenKey: string = "children"
) => {
  let rows = tree2FlatArray(tree, childrenKey);
  let objects = {};
  rows.forEach((item) => {
    objects[item[key]] = item;
  });
  return objects;
};

/**
 * 树形菜单数据输出
 * @date 2022-04-20
 * @param {object} data - 数据
 * @param {string} pid="parent_id" - 父类名称
 * @param {string} children="children" - 子类名称
 * @returns {object}
 */
export const treesBy = (
  data: any[],
  pid: string = "parent_id",
  children: string = "children"
) => {
  let map = {},
    val = <any>[];
  data.forEach((item) => {
    item[children] = [];
    map[item.id] = item;
  });
  data.forEach((item) => {
    const parent = map[item[pid]];
    if (!!parent) {
      parent[children].push(item);
    } else {
      val.push(item);
    }
  });
  return val;
};
/**
 * 格式化树形菜单,省市区
 * @param {object} org
 * @returns {object}
 */
export const mapTree = (org: any, keyId = "v", keyChild = "s", keyName = "n") => {
  const haveChildren =
    org[keyChild] && Array.isArray(org[keyChild]) && org[keyChild].length > 0;
  return {
    key: org[keyId],
    value: String(org[keyId]),
    label: org[keyName],
    isLeaf: true,
    children: haveChildren ? org[keyChild].map((i) => mapTree(i)) : [],
  };
};

/**
 * 格式化树形菜单,省市区
 * @param {object} org
 * @returns {object}
 */
export const toFlatArray = (
  tree: any[],
  parentId: string | number,
  keyId: string | number = "key"
) => {
  return tree.reduce((t, _) => {
    const child = _["children"];
    return [
      ...t,
      { [keyId]: _[keyId], parentId },
      ...(child && child.length ? toFlatArray(child, _[keyId]) : []),
    ];
  }, []);
};
