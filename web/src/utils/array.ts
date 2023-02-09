export function arrayDelIndex(array: any, keyName: string, key: string): any {
  if (array === null || array === undefined || array.length === 0) {
    return array;
  }

  const newArray = [];
  for (let i = 0; i < array.length; i++) {
    if (array[i][keyName] !== undefined && array[i][keyName] === key) {
      continue;
    }
    // @ts-ignore
    newArray.push(array[i]);
  }

  return newArray;
}

export function arrayAddIndex(array: any, keyName: string, key: string, row: any): any {
  if (array === null || array === undefined) {
    return array;
  }
  const newArray = [];

  if (array.length === 0) {
    // @ts-ignore
    newArray.push(row);
  } else {
    let isFor = false;
    for (let i = 0; i < array.length; i++) {
      if (array[i][keyName] !== undefined && array[i][keyName] === key) {
        array[i] = row;
        isFor = true;
      }
      // @ts-ignore
      newArray.push(array[i]);
    }

    if (!isFor) {
      // @ts-ignore
      newArray.push(row);
    }
  }

  return newArray;
}

export function objDalEmpty(obj: object): object {
  for (const key in obj) {
    if (obj[key] === '' || obj[key] === undefined || obj[key] == null || obj[key].length === 0) {
      delete obj[key];
    }
  }
  return obj;
}

export function filterArray(condition, data) {
  return data.filter((item) => {
    return Object.keys(condition).every((key) => {
      return String(item[key]).toLowerCase().includes(String(condition[key]).trim().toLowerCase());
    });
  });
}

export function findIndex(value, arr) {
  for (let i = 0; i < arr.length; i++) {
    const item = arr[i];
    if (item.value == value) {
      return i;
    }
  }
  return false;
}

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

/**
 * 对象拷贝
 * @param obj2
 * @param obj1
 */
export function copyObj(obj2: any, obj1: any) {
  for (const key in obj1) {
    if (obj2[key] !== undefined) {
      obj2[key] = obj1[key];
    }
  }
  return obj2;
}
