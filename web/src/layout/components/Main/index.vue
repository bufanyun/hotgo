<template>
  <RouterView>
    <template #default="{ Component, route }">
      {{ setKeepAlive(route) }}
      <template v-if="mode === 'production'">
        <transition :name="getTransitionName" appear mode="out-in">
          <div>
            <keep-alive v-if="keepAliveComponents.length" :include="keepAliveComponents">
              <component :is="Component" :key="route.fullPath" />
            </keep-alive>
            <component :is="Component" v-else :key="route.fullPath" />
          </div>
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
      const mode = import.meta.env.MODE;
      const router = useRouter();
      const { getIsPageAnimate, getPageAnimateType } = useProjectSetting();
      const asyncRouteStore = useAsyncRouteStore();
      const keepAliveComponents = computed(() => asyncRouteStore.keepAliveComponents);
      const getTransitionName = computed(() => {
        return unref(getIsPageAnimate) ? unref(getPageAnimateType) : '';
      });

      function getCurrentComponentName() {
        const currentMatched = router.currentRoute.value.matched;
        const currentComponent = currentMatched[currentMatched.length - 1].components.default;
        return currentComponent.name || currentComponent.__name;
      }

      function setKeepAlive(route) {
        if (!route?.meta?.keepAlive) {
          return;
        }

        const currentName = getCurrentComponentName();
        if (currentName === undefined || route.name === undefined) {
          return;
        }
        if (!keepAliveComponents.value.includes(currentName)) {
          asyncRouteStore.keepAliveComponents.push(currentName);
        }
      }

      return {
        keepAliveComponents,
        getTransitionName,
        setKeepAlive,
        mode,
      };
    },
  });
</script>

<style lang="less" scoped></style>
