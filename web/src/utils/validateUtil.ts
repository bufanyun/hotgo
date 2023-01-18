import { FormItemRule } from 'naive-ui';
/**
 * @description 表单验证封装
 */
export const validate = {
  ip(rule: FormItemRule, value: any, callback: Function) {
    // 支持通配符的ipv4正则
    const ipv4Regex =
      /^(?:[1-9]?[0-9]|1[0-9]{2}|2(?:[0-4][0-9]|5[0-5]))(?!.*?\.\*\.[*\d])(?:\.(?:(?:[1-9]?[0-9]|1[0-9]{2}|2(?:[0-4][0-9]|5[0-5]))|\*)){1,3}$/;
    //   Ipv6:
    const ipv6Regex =
      /^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$/;
    if (!value && !rule.required) {
      callback();
    }
    if (!value) {
      callback(new Error('请输入IP'));
    } else if (!ipv4Regex.test(value) && !ipv6Regex.test(value)) {
      callback(new Error('请输入正确的IP'));
    } else {
      callback();
    }
  },
  //0-100百分比验证
  percentage(rule: FormItemRule, value: any, callback: Function) {
    const reg = /^([1-9]{1,2}$)|(^[0-9]{1,2}\.[0-9]{1,2}$)|100$/;
    if (!value && !rule.required) {
      callback(new Error('请输入比例'));
    } else if (!reg.test(value)) {
      callback(new Error('请输入0-100的数字'));
    } else {
      callback();
    }
  },
  // 手机号 （eg:138********,159********）
  phone(rule: FormItemRule, value: any, callback: Function) {
    const regPhone = /^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\d{8}$/;
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入手机号码'));
    } else if (!regPhone.test(value)) {
      callback(new Error('手机号格式错误'));
    } else {
      callback();
    }
  },
  // 用户名 （eg:a123456）
  userName(rule: FormItemRule, value: any, callback: Function) {
    const regUserName = /^[0-9a-zA-Z]{6,16}$/;
    if (!value && !rule.required) {
      callback(new Error('请输入登录账号'));
    } else if (!regUserName.test(value)) {
      callback(new Error('请输入6-16位由字母和数字组成的登录账号'));
    } else {
      callback();
    }
  },
  // 账号
  account(rule: FormItemRule, value: any, callback: Function) {
    const regex = /^[\w_\d]{6,16}$/;
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入账号'));
    } else if (!regex.test(value)) {
      callback(new Error('请输入6-16位由字母、数字或下划线组成的账号'));
    } else {
      callback();
    }
  },
  // 密码
  password(rule: FormItemRule, value: any, callback: Function) {
    const regPassword = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$/;
    if (!value && !rule.required) {
      callback(new Error('请输入密码'));
    } else if (!regPassword.test(value)) {
      callback(new Error('密码格式错误！必须包含6-18为字母和数字'));
    } else {
      callback();
    }
  },
  // 邮箱
  email(rule: FormItemRule, value: any, callback: Function) {
    const regEmails = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
    // console.log('isRequired is: ', JSON.stringify(isRequired))
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入邮箱'));
    } else if (!regEmails.test(value)) {
      callback(new Error('邮箱格式错误'));
    } else {
      callback();
    }
  },
  // 金额验证
  amount(rule: FormItemRule, value: any, callback: Function) {
    const regAmount = /(^[0-9]{1,10}$)|(^[0-9]{1,10}[\.]{1}[0-9]{1,2}$)/;
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入金额'));
    } else if (!regAmount.test(value)) {
      callback(new Error('金额格式错误，最多允许输入10位整数及2位小数'));
    } else {
      callback();
    }
  },
  // 身份证验证
  idCard(rule: FormItemRule, value: any, callback: Function, isEnabled = true) {
    const regIdCard =
      /^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$|^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}([0-9]|X|x)$/;
    if (!isEnabled) {
      callback();
    } else if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入身份证号'));
    } else if (!regIdCard.test(value)) {
      callback(new Error('身份证号码格式错误'));
    } else {
      callback();
    }
  },
  // 银行卡验证
  bank(rule: FormItemRule, value: any, callback: Function) {
    const regBank = /^([1-9]{1})(\d{15}|\d{16}|\d{18})$/;
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback();
    } else if (!regBank.test(value)) {
      callback(new Error('银行卡号码格式错误'));
    } else {
      callback();
    }
  },
  // 非零正整数验证
  num(rule: FormItemRule, value: any, callback: Function) {
    const reg = /^\+?[1-9][0-9]*$/;
    if (!value && !rule.required) {
      callback(new Error('请填写非零正整数'));
    } else {
      if (!reg.test(value)) {
        callback(new Error('请输入非零的正整数'));
      } else {
        callback();
      }
    }
  },
  // 银行卡
  bankCard(rule: FormItemRule, value: any, callback: Function) {
    const regBankCard = /^(\d{16}|\d{19})$/;
    if (value == '' && !rule.required) {
      callback(new Error('请输入银行卡号'));
    } else {
      if (!regBankCard.test(value)) {
        callback(new Error('银行卡号格式错误'));
      } else {
        callback();
      }
    }
  },
  // 固话格式
  tel(rule: FormItemRule, value: any, callback: Function) {
    const regTel = /^(0\d{2,3}-?)?\d{7,8}$/;
    if (value == '' && !rule.required) {
      callback(new Error('请输入座机号码'));
    } else {
      if (!regTel.test(value)) {
        callback(new Error('座机号码格式错误'));
      } else {
        callback();
      }
    }
  },
  // QQ号码
  qq(rule: FormItemRule, value: any, callback: Function) {
    const regex = /^[1-9][0-9]{4,}$/;
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入QQ号码'));
    } else {
      if (!regex.test(value)) {
        callback(new Error('QQ号码格式错误'));
      } else {
        callback();
      }
    }
  },
  // weibo号
  weibo(rule: FormItemRule, value: any, callback: Function) {
    const regex = /^[0-9a-zA-Z\u4e00-\u9fa5_-]*$/;
    if (!value && !rule.required) {
      callback();
    } else if (!value) {
      callback(new Error('请输入微博账号'));
    } else {
      if (!regex.test(value)) {
        callback(new Error('微博号码格式错误'));
      } else {
        callback();
      }
    }
  },
  // 不验证
  none(_rule: FormItemRule, _value: any, callback: Function) {
    callback();
  },
};
