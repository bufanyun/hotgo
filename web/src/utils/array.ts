
export function delNullProperty(obj) {
  for (const i in obj) {
    if (obj[i] === undefined || obj[i] === null || obj[i] === '') {
      delete obj[i];
    } else if (obj[i].constructor === Object) {
      if (Object.keys(obj[i]).length === 0) delete obj[i];
      delNullProperty(obj[i]);
    } else if (obj[i].constructor === Array) {
      if (Array.prototype.isPrototypeOf(obj[i]) && obj[i].length === 0) {
        delete obj[i];
      } else {
        for (let index = 0; index < obj[i].length; index++) {
          if (
            obj[i][index] === undefined ||
            obj[i][index] === null ||
            obj[i][index] === '' ||
            JSON.stringify(obj[i][index]) === '{}'
          ) {
            obj[i].splice(index, 1);
            index--;
          }
          if (obj[i][index] === undefined || obj[i][index].constructor !== undefined) {
            continue;
          }
          if (obj[i][index].constructor === Object || obj[i][index].constructor === Array) {
            delNullProperty(obj[i][index]);
          }
        }
      }
    }
  }
  return obj;
}

export function reverse(array) {
  if (array !== undefined && array !== null && array.length > 0) {
    return array.reverse();
  }
  return array;
}

export function encodeParams(obj) {
  const arr = [];
  for (const p in obj) {
    // @ts-ignore
    arr.push(encodeURIComponent(p) + '=' + encodeURIComponent(obj[p]));
  }
  return arr.join('&');
}

export function copyObj(obj2: any, obj1: any) {
  for (const key in obj1) {
    if (obj2[key] !== undefined) {
      obj2[key] = obj1[key];
    }
  }
  return obj2;
}

// 返回两个数组的差集
export function findArrayDifference(arr1: number[], arr2: number[]): number[] {
  return arr1.filter((num) => !arr2.includes(num));
}
