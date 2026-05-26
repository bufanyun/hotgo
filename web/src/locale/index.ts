import { createI18n } from 'vue-i18n';
import type { I18nOptions } from 'vue-i18n';
import en from './en.json';
import zhHant from './zh-Hant.json';
import zhHans from './zh-Hans.json';

const messages = {
  en,
  'zh-TW': zhHant,
  'zh-CN': zhHans,
};

// 创建 i18n 实例配置
const i18nConfig: I18nOptions = {
  legacy: false, // 使用 Composition API 模式
  locale: 'zh-CN', // 默认语言
  fallbackLocale: 'zh-CN', // 回退语言
  messages, // 语言包
  missingWarn: false, // 关闭找不到 key 的警告
  fallbackWarn: false, // 关闭回退警告
  // 当找不到翻译时，返回 key 本身
  missing: (_locale, key) => {
    return key;
  },
};

const i18n = createI18n(i18nConfig);

export default i18n;

// 导出 t 函数，支持多种参数形式
export const t = (key: string, named?: Record<string, unknown>, options?: any): string => {
  if (named !== undefined) {
    return i18n.global.t(key, named, options) as string;
  }
  return i18n.global.t(key) as string;
};
