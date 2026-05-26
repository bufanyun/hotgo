import { inject, provide } from 'vue';
import type { InjectionKey } from 'vue';

export default function useContext<T>(contextName = 'context') {
  const injectKey: InjectionKey<T> = Symbol(contextName);

  function useProvide(context: T) {
    provide(injectKey, context);
    return context;
  }

  function useInject() {
    return inject(injectKey) as T;
  }

  return {
    useProvide,
    useInject,
  };
}
