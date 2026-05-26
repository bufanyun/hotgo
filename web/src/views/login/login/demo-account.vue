<template>
  <template v-if="accounts && accounts.length > 0">
    <n-space :vertical="true">
      <n-divider>演示角色登录</n-divider>
      <n-space justify="center">
        <n-button
          v-for="item in accounts"
          :key="item.username"
          type="primary"
          @click="login(item.username, item.password)"
        >
          {{ item.name }}
        </n-button>
      </n-space>
      <n-space justify="center" class="mt-2">
        <n-text depth="3">SaaS系统多租户多应用设计</n-text>
      </n-space>
    </n-space>
  </template>
</template>

<script lang="ts" setup>
import { getDemoAccounts } from "@/debug/account";
interface Emits {
  (e: "login", param: { username: string; password: string }): void;
}

const emit = defineEmits<Emits>();

const accounts = getDemoAccounts();

function login(username: string, password: string) {
  emit("login", { username, password });
}
</script>

<style scoped></style>
