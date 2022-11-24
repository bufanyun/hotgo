import { format } from 'date-fns';

const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm';
const DATE_FORMAT = 'YYYY-MM-DD ';

export function formatToDateTime(date: Date, formatStr = DATE_TIME_FORMAT): string {
  const date2 = new Date(date);
  return format(date2, formatStr);
}

export function formatToDate(date: Date, formatStr = DATE_FORMAT): string {
  return format(date, formatStr);
}

export function timestampToTime(timestamp) {
  const date = new Date(timestamp * 1000);
  const Y = date.getFullYear() + '-';
  const M = (date.getMonth() + 1 <= 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
  const D = (date.getDate() + 1 <= 10 ? '0' + date.getDate() : date.getDate()) + ' ';
  const h = (date.getHours() + 1 <= 10 ? '0' + date.getHours() : date.getHours()) + ':';
  const m = (date.getMinutes() + 1 <= 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
  const s = date.getSeconds() + 1 <= 10 ? '0' + date.getSeconds() : date.getSeconds();
  return Y + M + D + h + m + s;
}

export function timestampToTimeNF(timestamp) {
  const date = new Date(timestamp);
  const Y = date.getFullYear();
  const M = date.getMonth() + 1 <= 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1;
  const D = date.getDate() + 1 <= 10 ? '0' + date.getDate() : date.getDate();
  const h = date.getHours() + 1 <= 10 ? '0' + date.getHours() : date.getHours();
  const m = date.getMinutes() + 1 <= 10 ? '0' + date.getMinutes() : date.getMinutes();
  const s = date.getSeconds() + 1 <= 10 ? '0' + date.getSeconds() : date.getSeconds();
  return Y.toString() + M.toString() + D.toString() + h.toString() + m.toString() + s.toString();
}

export function timestampToDate(timestamp) {
  const date = new Date(timestamp);
  const Y = date.getFullYear() + '-';
  const M = (date.getMonth() + 1 <= 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
  const D = date.getDate() + 1 <= 10 ? '0' + date.getDate() : date.getDate();
  return Y + M + D;
}

export function getTime() {
  const myDate = new Date();

  const hour = myDate.getHours().toString().padStart(2, '0');

  const minutes = myDate.getMinutes().toString().padStart(2, '0');

  const seconed = myDate.getSeconds().toString().padStart(2, '0');

  return hour + ':' + minutes + ':' + seconed;
}

export function getDate() {
  const myDate = new Date();

  const month = (myDate.getMonth() + 1).toString().padStart(2, '0');

  const day = myDate.getDate().toString().padStart(2, '0');

  return myDate.getFullYear() + '-' + month + '-' + day;
}

export function defaultStatisticsTimeOptions() {
  return new Date().getTime() - 86400 * 1000;
}

export function formatBefore(oldDate) {
  //当前时间
  const newDate = new Date();
  const newDateTime1 = newDate.getTime(); //含有时分秒
  newDate.setHours(0);
  newDate.setMinutes(0);
  newDate.setSeconds(0);
  newDate.setMilliseconds(0);
  const newDateTime2 = newDate.getTime(); //当前时间,不含有时分秒

  //传递时间
  const oldDateTime1 = oldDate.getTime(); //含有时分秒
  oldDate.setHours(0);
  oldDate.setMinutes(0);
  oldDate.setSeconds(0);
  oldDate.setMilliseconds(0);
  const oldDateTime2 = oldDate.getTime(); //不含有时分秒

  const d1 = (newDateTime1 - oldDateTime1) / 1000;
  const d2 = (newDateTime2 - oldDateTime2) / 1000;

  let res = '';
  if (d2 > 0) {
    //是几天前
    const days = parseInt(d2 / 86400);
    if (days === 1) {
      res = '昨天';
    } else if (days === 2) {
      res = '前天';
    } else {
      res = days + '天前';
    }
  } else {
    //是今天
    const hours = parseInt(d1 / 3600);
    if (hours > 0) {
      res = hours + '小时前';
    } else {
      const minutes = parseInt(d1 / 60);
      if (minutes > 0) {
        res = minutes + '分钟前';
      } else {
        const seconds = parseInt(d1);
        if (seconds > 10) {
          res = seconds + '秒前';
        } else {
          res = '刚刚';
        }
      }
    }
  }
  return res;
}

// @ts-ignore
export function formatAfter(end): string {
  const start = new Date();
  let sjc = start.getTime() - end.getTime(); //时间差的毫秒数
  if (end.getTime() - start.getTime() > 0) {
    sjc = end.getTime() - start.getTime(); //时间差的毫秒数
  }

  const days = Math.floor(sjc / (24 * 3600 * 1000)); //计算出相差天数

  const leave1 = sjc % (24 * 3600 * 1000); //计算天数后剩余的毫秒数
  const hours = Math.floor(leave1 / (3600 * 1000)); //计算出小时数

  const leave2 = leave1 % (3600 * 1000); //计算小时数后剩余的毫秒数
  const minutes = Math.floor(leave2 / (60 * 1000)); //计算相差分钟数

  const leave3 = leave2 % (60 * 1000); //计算分钟数后剩余的毫秒数
  const seconds = Math.round(leave3 / 1000); //计算相差秒数
  if (days > 0) {
    return days + '天后';
  }
  if (hours > 0) {
    return hours + '小时后';
  }
  if (minutes > 0) {
    return minutes + '分钟后';
  }
  if (seconds > 0) {
    return seconds + '秒后';
  }

  return '刚刚';
}
