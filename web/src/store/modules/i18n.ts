import { defineStore } from 'pinia';
import { createStorage } from '@/utils/Storage';
import { CurrentLocale } from '@/store/mutation-types';
import i18n from '@/locale/index';
import { useUserStore } from '@/store/modules/user';

export const availableLocales = [
  {
    label: '简体中文',
    key: 'zh-CN',
  },
  {
    label: '繁體中文',
    key: 'zh-TW',
  },
  {
    label: 'English',
    key: 'en',
  },
];

export interface II18nStore {
  currentLocale: string;
}

const Storage = createStorage({ storage: localStorage });

export const useI18nStore = defineStore({
  id: 'I18nStore',
  state: (): II18nStore => ({
    currentLocale: Storage.get(CurrentLocale, 'zh-CN'),
  }),
  getters: {},
  actions: {
    getLocale(): string {
      return this.currentLocale;
    },

    setLocale(locale: string) {
      if (availableLocales.some((l) => l.key === locale)) {
        (i18n.global.locale as any).value = locale;
        this.currentLocale = locale;
        Storage.set(CurrentLocale, locale);
      }
    },

    initLocale(): void {
      const userStore = useUserStore();
      const defaultLanguage = userStore.loginConfig?.defaultLanguage || 'zh-CN';

      // 未开启国际化功能，仅设置语言不持久化
      if (!userStore.loginConfig?.i18nSwitch) {
        this.setLocale(defaultLanguage);
        return;
      }

      // 优先使用本地存储的语言设置
      const savedLocale = Storage.get(CurrentLocale, defaultLanguage);
      if (savedLocale && availableLocales.some((l) => l.key === savedLocale)) {
        this.setLocale(savedLocale);
        return;
      }

      // 使用系统配置的默认语言
      this.setLocale(defaultLanguage);
    },
  },
});
