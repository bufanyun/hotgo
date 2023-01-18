/**
 * 随机生成字符串
 */
export function getRandomString(len = 12, isSmall = false) {
  const _charStr = 'abacdefghjklmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ0123456789';
  const _smallCharStr = 'abacdefghjklmnopqrstuvwxyz0123456789';
  const charStr = isSmall ? _smallCharStr : _charStr;
  const min = 0;
  const max = charStr.length - 1;
  let _str = '';
  //循环生成字符串
  for (let i = 0, index; i < len; i++) {
    index = (function (randomIndexFunc, i) {
      return randomIndexFunc(min, max, i, randomIndexFunc);
    })(function (min, max, i, _self) {
      const numStart = _charStr.length - 10;
      let indexTemp = Math.floor(Math.random() * (max - min + 1) + min);
      if (i == 0 && indexTemp >= numStart) {
        indexTemp = _self(min, max, i, _self);
      }
      return indexTemp;
    }, i);
    _str += _charStr[index];
  }
  return _str;
}

/**
 * 隐藏中间几位字符
 */
export function structure(array) {
  // 将字符串转化成数组
  const arrBox = [...array];
  const count = arrBox.length;
  if (count == 1) {
    return '*';
  }
  let min = 1;
  let max = count;
  // 两位姓名
  if (count == 2) {
    min = 0;
    max = count;
  }
  // 三位姓名
  if (count == 3) {
    min = 0;
    max = count - 1;
  }
  // if (count >= 2 && count <= 8) {
  //   min = 1;
  // }

  // 手机号
  if (count == 11) {
    min = 3;
    max = 7;
  }
  // 身份证号码
  if (count >= 15) {
    min = 9;
    max = count - 4;
  }
  // 2.将数组中的4-7位变成*
  let str = '';
  arrBox.map((res, index) => {
    if (index > min && index < max) {
      str += '*';
    } else {
      str += res;
    }
  });
  return str;
}
