<template>
  <n-form-item class="default-color">
    <div class="flex view-account-other">
      <div class="flex-initial">
        <span>其它登录方式</span>
      </div>
      <div class="flex-initial mx-2">
        <a @click="handleLoginWechat">
          <n-icon size="24" color="rgb(24, 160, 88)">
            <LogoWechat />
          </n-icon>
        </a>
      </div>
      <div class="flex-initial mx-2">
        <a @click="handleLogoTiktok">
          <n-icon size="24" color="rgba(25, 28, 34, 0.88)">
            <LogoTiktok />
          </n-icon>
        </a>
      </div>
      <div class="flex-initial" style="margin-left: auto" v-if="userStore.loginConfig?.loginRegisterSwitch === 1">
        <a @click="updateActiveModule(moduleKey)">{{ tag }}</a>
      </div>
    </div>
  </n-form-item>
</template>

<script lang="ts" setup>
  import { LogoWechat, LogoTiktok } from '@vicons/ionicons5';

  import { useUserStore } from '@/store/modules/user';
  import {useMessage} from "naive-ui";

  const userStore = useUserStore();

  interface Props {
    moduleKey: string;
    tag: string;
  }

  withDefaults(defineProps<Props>(), {
    moduleKey: 'register',
    tag: '注册账号',
  });

  const message = useMessage();
  const emit = defineEmits(['updateActiveModule']);

  function updateActiveModule(key: string) {
    emit('updateActiveModule', key);
  }

  function handleLogoTiktok() {
    console.log('handleLogoTiktok...');
    message.info('暂未开放');
  }

  function handleLoginWechat() {
    console.log('handleLoginWechat...');
    message.info('暂未开放');
  }
</script>

<style scoped></style>
