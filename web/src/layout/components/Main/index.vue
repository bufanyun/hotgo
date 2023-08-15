<template>
  <RouterView>
    <template #default="{ Component, route }">
      {{ retryKeepAlive(route) }}
      <template v-if="mode === 'production'">
        <transition :name="getTransitionName" appear mode="out-in">
          <keep-alive v-if="keepAliveComponents.length" :include="keepAliveComponents">
            <component :is="Component" :key="route.fullPath" />
          </keep-alive>
          <component :is="Component" v-else :key="route.fullPath" />
        </transition>
      </template>
      <template v-else>
        <keep-alive v-if="keepAliveComponents.length" :include="keepAliveComponents">
          <component :is="Component" :key="route.fullPath" />
        </keep-alive>
        <component :is="Component" v-else :key="route.fullPath" />
      </template>
    </template>
  </RouterView>
</template>

<script>
  import { computed, defineComponent, unref } from 'vue';
  import { useAsyncRouteStore } from '@/store/modules/asyncRoute';
  import { useProjectSetting } from '@/hooks/setting/useProjectSetting';
  import { useRouter } from 'vue-router';

  export default defineComponent({
    name: 'MainView',
    components: {},
    props: {
      notNeedKey: {
        type: Boolean,
        default: false,
      },
      animate: {
        type: Boolean,
        default: true,
      },
    },
    setup() {
      const router = useRouter();
      const { getIsPageAnimate, getPageAnimateType } = useProjectSetting();
      const asyncRouteStore = useAsyncRouteStore();
      // 需要缓存的路由组件
      const keepAliveComponents = computed(() => asyncRouteStore.keepAliveComponents);

      const getTransitionName = computed(() => {
        return unref(getIsPageAnimate) ? unref(getPageAnimateType) : '';
      });

      function getCurrentComponentName() {
        const currentMatched = router.currentRoute.value.matched;
        const currentComponent = currentMatched[currentMatched.length - 1].components.default;
        return currentComponent.name || currentComponent.__name;
      }

      function retryKeepAlive(route) {
        if (!route?.meta?.keepAlive) {
          return;
        }

        const currentName = getCurrentComponentName();
        if (currentName === undefined || route.name === undefined) {
          return;
        }

        const index = keepAliveComponents.value.findIndex((name) => name === route.name);
        if (index > -1 && currentName !== route.name) {
          const index2 = keepAliveComponents.value.findIndex((name) => name === currentName);
          if (index2 === -1) {
            console.warn(
              `the routing name configured on the backend is inconsistent with the component name in the. vue file. KeepAlive has been retried based on the actual component name, but this may cause unexpected issues.\n routeName:` +
                route.name +
                ',currentName:' +
                currentName
            );
            asyncRouteStore.keepAliveComponents.push(currentName);
          }
        }
      }
      const mode = import.meta.env.MODE;
      return {
        keepAliveComponents,
        getTransitionName,
        retryKeepAlive,
        mode,
      };
    },
  });
</script>

<style lang="less" scoped></style>
