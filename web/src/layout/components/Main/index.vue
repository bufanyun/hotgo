<template>
  <RouterView>
    <template #default="{ Component, route }">
      {{ retryKeepAlive(route) }}
      <transition :name="getTransitionName" mode="out-in" appear>
        <keep-alive v-if="keepAliveComponents" :include="keepAliveComponents">
          <component :is="Component" :key="route.fullPath" />
        </keep-alive>
        <component v-else :is="Component" :key="route.fullPath" />
      </transition>
    </template>
  </RouterView>
</template>

<script>
  import { defineComponent, computed, unref } from 'vue';
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

      return {
        keepAliveComponents,
        getTransitionName,
        retryKeepAlive,
      };
    },
  });
</script>

<style lang="less" scoped></style>
