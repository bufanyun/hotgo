<template>
  <div class="double-column-menu">
    <!-- 统一的头部容器 -->
    <div class="menu-header" :class="[getMenuClass, { collapsed: _props.secondMenuCollapsed }]">
      <div class="header-content" :class="{ 'header-content--compact': isLongSiteName }">
        <!-- LOGO 区域 -->
        <div class="logo-area">
          <img src="~@/assets/images/logo.png" alt="logo" class="logo-img" />
        </div>
        <!-- 网站名称区域 -->
        <div v-if="!_props.secondMenuCollapsed" class="header-title">
          {{ userStore.loginConfig?.projectName }}
        </div>
      </div>
    </div>

    <!-- 菜单容器 -->
    <div class="menu-container">
      <!-- 一级菜单 -->
      <div class="first-level-menu" :class="getMenuClass">
        <div class="first-level-scroll">
          <div
            v-for="menu in firstLevelMenus"
            :key="menu.key"
            class="first-level-item"
            :class="{ active: activeFirstMenu === menu.key }"
            :style="
              activeFirstMenu === menu.key
                ? {
                    background: designStore.appTheme,
                    color: '#ffffff',
                  }
                : {}
            "
            @click="selectFirstMenu(menu)"
          >
            <n-icon
              v-if="menu.icon"
              :component="menu.icon"
              :size="20"
              :style="activeFirstMenu === menu.key ? { color: '#ffffff' } : {}"
            />
            <div
              class="menu-title"
              :style="activeFirstMenu === menu.key ? { color: '#ffffff' } : {}"
            >
              {{ menu.label }}
            </div>
          </div>
        </div>
      </div>

      <!-- 二级菜单 -->
      <div
        class="second-level-menu"
        :class="[getMenuClass, { collapsed: _props.secondMenuCollapsed }]"
      >
        <div class="second-level-content">
          <n-menu
            v-if="secondLevelMenus.length > 0 && !_props.secondMenuCollapsed"
            :options="secondLevelMenus"
            :inverted="inverted"
            :value="selectedKeys"
            :expanded-keys="openKeys"
            @update:value="clickMenuItem"
            @update:expanded-keys="openKeys = $event"
          />
          <!-- 收起状态显示的图标菜单 -->
          <div v-if="_props.secondMenuCollapsed" class="collapsed-menu">
            <n-tooltip
              v-for="menu in secondLevelMenus"
              :key="menu.key"
              placement="right"
              trigger="hover"
            >
              <template #trigger>
                <div
                  class="collapsed-menu-item"
                  :class="{ active: selectedKeys === menu.key }"
                  :style="
                    selectedKeys === menu.key
                      ? {
                          color: designStore.appTheme,
                          background: designStore.appTheme + '1a',
                        }
                      : {}
                  "
                  @click="clickMenuItem(menu.key)"
                >
                  <n-icon v-if="menu.icon" :component="menu.icon" :size="18" />
                  <n-icon v-else :size="18">
                    <MenuOutlined />
                  </n-icon>
                </div>
              </template>
              {{ menu.label }}
            </n-tooltip>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed, watch, onMounted } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useAsyncRouteStore } from '@/store/modules/asyncRoute';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useUserStore } from '@/store/modules/user';
  import { useDesignSettingStore } from '@/store/modules/designSetting';
  import { generatorMenu } from '@/utils';
  import { cloneDeep } from 'lodash-es';
  import { MenuOutlined } from '@vicons/antd';

  const _props = defineProps({
    secondMenuCollapsed: {
      type: Boolean,
      default: false,
    },
  });

  const emit = defineEmits(['clickMenuItem', 'update:secondMenuCollapsed']);
  const currentRoute = useRoute();
  const router = useRouter();
  const userStore = useUserStore();
  const asyncRouteStore = useAsyncRouteStore();
  const settingStore = useProjectSettingStore();
  const designStore = useDesignSettingStore();
  const activeFirstMenu = ref<string>('');
  const selectedKeys = ref<string>(currentRoute.name as string);
  const firstLevelMenus = ref<any[]>([]);
  const secondLevelMenus = ref<any[]>([]);
  const openKeys = ref<string[]>([]);
  const inverted = computed(() => {
    return ['dark', 'header-dark'].includes(settingStore.navTheme);
  });

  const isLongSiteName = computed(() => (userStore.loginConfig?.projectName?.length ?? 0) > 10);

  // 菜单样式类
  const getMenuClass = computed(() => {
    const navTheme = settingStore.navTheme;
    return {
      dark: navTheme === 'dark' || navTheme === 'header-dark',
      light: navTheme === 'light',
    };
  });

  // 获取当前路由名称（考虑 activeMenu）
  const getCurrentRouteName = computed(() => {
    const activeMenu: string = (currentRoute.meta?.activeMenu as string) || '';
    return activeMenu || (currentRoute.name as string);
  });

  // 获取所有菜单（缓存，避免重复调用）
  const getAllMenus = () => {
    return cloneDeep(generatorMenu(asyncRouteStore.getMenus));
  };

  // 更新 openKeys，根据当前路由查找父菜单项
  function updateOpenKeys() {
    const allMenus = getAllMenus();
    const activeMenuObj = allMenus.find((m) => m.key === activeFirstMenu.value);
    const currentRouteName = getCurrentRouteName.value;

    if (activeMenuObj?.children) {
      // 查找当前路由在哪个二级菜单项的子菜单中
      const parentMenu = activeMenuObj.children.find((child: any) => {
        if (child.key === currentRouteName) {
          return false;
        }
        // 检查子菜单的子菜单
        return child.children?.some((subChild: any) => subChild.key === currentRouteName);
      });
      openKeys.value = parentMenu?.key ? [parentMenu.key] : [];
    } else {
      openKeys.value = [];
    }
  }

  // 根据当前路由查找对应的一级菜单
  function findActiveFirstMenu(allMenus: any[]): string {
    if (firstLevelMenus.value.length === 0) {
      return '';
    }

    const targetRouteName = getCurrentRouteName.value;

    // 查找包含当前路由的一级菜单
    for (const menu of allMenus) {
      if (menu.key === targetRouteName) {
        return menu.key;
      }

      // 检查一级菜单的子菜单中是否包含当前路由
      if (menu.children && menu.children.length > 0) {
        const directChild = menu.children.find((child: any) => child.key === targetRouteName);
        if (directChild) {
          return menu.key;
        }

        // 检查子菜单的子菜单（三级菜单）
        for (const child of menu.children) {
          if (child.children && child.children.length > 0) {
            const grandChild = child.children.find(
              (subChild: any) => subChild.key === targetRouteName
            );
            if (grandChild) {
              return menu.key;
            }
          }
        }
      }
    }

    // 如果没找到，返回当前激活的一级菜单或第一个菜单
    return activeFirstMenu.value || firstLevelMenus.value[0]?.key || '';
  }

  // 更新菜单
  function updateMenu() {
    const allMenus = getAllMenus();
    firstLevelMenus.value = allMenus.map((menu) => ({
      key: menu.key,
      label: menu.label,
      icon: menu.icon,
    }));

    // 根据当前路由更新激活的一级菜单
    const foundFirstMenu = findActiveFirstMenu(allMenus);
    if (foundFirstMenu) {
      activeFirstMenu.value = foundFirstMenu;
    } else if (!activeFirstMenu.value && firstLevelMenus.value.length > 0) {
      // 如果没找到且当前没有激活的一级菜单，使用第一个
      activeFirstMenu.value = firstLevelMenus.value[0].key;
    }

    // 更新二级菜单
    const activeMenu = allMenus.find((m) => m.key === activeFirstMenu.value);
    secondLevelMenus.value = activeMenu?.children || [];
  }

  // 选择一级菜单
  function selectFirstMenu(menu: any) {
    activeFirstMenu.value = menu.key;
    const allMenus = getAllMenus();
    const activeMenu = allMenus.find((m) => m.key === menu.key);
    secondLevelMenus.value = activeMenu?.children || [];

    if (!secondLevelMenus.value || secondLevelMenus.value.length === 0) {
      clickMenuItem(menu.key);
      openKeys.value = [];
    } else {
      if (secondLevelMenus.value[0]?.children?.length > 0 && secondLevelMenus.value[0].key) {
        openKeys.value = [secondLevelMenus.value[0].key];
        if (secondLevelMenus.value[0].children[0]?.key) {
          clickMenuItem(secondLevelMenus.value[0].children[0].key);
        }
      } else if (secondLevelMenus.value[0]) {
        openKeys.value = [];
        clickMenuItem(secondLevelMenus.value[0].key);
      }
    }
  }

  function clickMenuItem(key: string) {
    if (/http(s)?:/.test(key)) {
      window.open(key);
    } else if (selectedKeys.value !== key) {
      router.push({ name: key });
    }
    selectedKeys.value = key;
    emit('clickMenuItem', key);
  }

  watch(
    () => currentRoute.fullPath,
    () => {
      selectedKeys.value = getCurrentRouteName.value;
      updateMenu();
      updateOpenKeys();
    }
  );

  onMounted(() => {
    updateMenu();
    selectedKeys.value = getCurrentRouteName.value;
    updateOpenKeys();
  });
</script>

<style lang="less" scoped>
  .double-column-menu {
    display: flex;
    flex-direction: column;
    height: 100vh;

    .menu-header {
      height: 64px;
      width: 240px;
      display: flex;
      align-items: center;
      justify-content: center;
      background: #001428;
      flex-shrink: 0;
      transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);
      overflow: hidden;
      position: relative;

      &::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 1px;
        background: linear-gradient(90deg, transparent, rgba(64, 158, 255, 0.3), transparent);
      }

      &.collapsed {
        width: 130px;
      }

      &.light {
        background: #fff;

        &::before {
          background: linear-gradient(90deg, transparent, rgba(64, 158, 255, 0.2), transparent);
        }
      }

      .header-content {
        display: flex;
        align-items: center;
        gap: 14px;
        max-width: 100%;
        padding: 0 18px;

        &--compact {
          gap: 8px;
          padding: 0 10px;
        }
      }

      .logo-area {
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;
        height: 100%;

        .logo-img {
          max-width: 32px;
          max-height: 32px;
          width: auto;
          height: auto;
          object-fit: contain;
          border-radius: 50%;
        }
      }

      .header-title {
        display: flex;
        align-items: center;
        font-size: 15px;
        font-weight: 700;
        letter-spacing: 0.3px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        line-height: 1.2;
        height: 100%;
        color: #ffffff;
      }
    }

    .light .header-title {
      color: #333639;
    }

    .menu-container {
      display: flex;
      flex: 1;
      overflow: hidden;
    }

    .first-level-menu {
      width: 80px;
      background: #001428;
      display: flex;
      flex-direction: column;
      flex-shrink: 0;
      height: 100%;
      border-right: 1px solid rgba(255, 255, 255, 0.08);

      &.light {
        background: #fff;
        border-right-color: #e5e7eb;

        .first-level-item {
          color: #4a5568;

          &:hover {
            background: rgba(64, 158, 255, 0.08);
            color: v-bind('designStore.appTheme');
          }

          &.active {
            font-weight: 600;
          }
        }
      }

      .first-level-scroll {
        flex: 1;
        overflow-y: auto;
        overflow-x: hidden;
        padding: 8px 5px;

        // 隐藏滚动条
        &::-webkit-scrollbar {
          width: 0;
          display: none;
        }

        -ms-overflow-style: none;
        scrollbar-width: none;
      }

      .first-level-item {
        width: 70px;
        height: 62px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        color: rgba(255, 255, 255, 0.7);
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        border-radius: 8px;
        margin-bottom: 4px;

        &:hover {
          background: rgba(255, 255, 255, 0.08);
          color: #fff;
          transform: translateY(-1px);
        }

        &.active {
          font-weight: 600;
        }

        .menu-title {
          font-size: 12px;
          margin-top: 5px;
          text-align: center;
          line-height: 1.3;
          max-width: 64px;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }
    }

    .second-level-menu {
      width: 160px;
      background: #001428;
      display: flex;
      flex-direction: column;
      flex-shrink: 0;
      height: 100%;
      transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1);

      &.light {
        background: #fff;

        .collapsed-menu .collapsed-menu-item {
          color: #4a5568;

          &:hover {
            background: rgba(64, 158, 255, 0.08);
            color: #409eff;
          }
        }
      }

      &.collapsed {
        width: 50px;
      }

      .second-level-content {
        flex: 1;
        overflow-y: auto;
        overflow-x: hidden;

        &::-webkit-scrollbar {
          width: 0;
          display: none;
        }

        -ms-overflow-style: none;
        scrollbar-width: none;
      }

      // 收起状态的图标菜单
      .collapsed-menu {
        padding: 8px 0;

        .collapsed-menu-item {
          width: 50px;
          height: 46px;
          display: flex;
          align-items: center;
          justify-content: center;
          cursor: pointer;
          color: rgba(255, 255, 255, 0.7);
          transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
          border-radius: 6px;
          margin: 2px 0;

          &:hover {
            background: rgba(64, 158, 255, 0.12);
            color: #fff;
          }
        }
      }

      :deep(.n-menu) {
        background: transparent;
        padding: 10px 8px;

        .n-menu-item {
          margin-top: 4px;
          border-radius: 6px;
          height: 44px;
          font-size: 14px;
          transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

          &:hover {
            background: rgba(64, 158, 255, 0.1);
            transform: translateX(2px);
          }
        }

        .n-menu-item-content {
          padding-left: 14px !important;
          border-radius: 6px;
          &::before {
            display: none !important;
          }
        }

        .n-menu-item-content.n-menu-item-content--selected::before {
          display: none !important;
        }

        .n-menu-item-content--selected {
          border-radius: 6px;
          transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
          background-color: var(--n-item-color-active) !important;
        }

        .n-menu-item--selected {
          background-color: transparent !important;

          .n-menu-item-content {
            background-color: v-bind('designStore.appTheme') !important;
            color: #ffffff !important;
            font-weight: 600 !important;

            &::before {
              display: none !important;
            }

            span,
            .n-menu-item-content__icon,
            * {
              color: #ffffff !important;
            }
          }
        }
      }

      &.light :deep(.n-menu) {
        .n-menu-item {
          &:hover {
            background: rgba(64, 158, 255, 0.06);
          }
        }
      }
    }
  }
</style>
