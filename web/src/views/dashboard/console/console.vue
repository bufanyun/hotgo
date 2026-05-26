<template>
  <div class="mt-2">
    <Company v-if="userStore.isCompanyDept" />
    <Tenant v-else-if="userStore.isTenantDept" />
    <Merchant v-else-if="userStore.isMerchantDept" />
    <User v-else-if="userStore.isUserDept" />
    <template v-else>
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="默认首页">
          部门类型未配置首页：{{ userStore.info?.deptType }}
        </n-card>
      </div>
    </template>
  </div>
</template>
<script lang="ts" setup>
  import Company from './console_company.vue';
  import Merchant from './console_merchant.vue';
  import User from './console_user.vue';
  import Tenant from './console_tenant.vue';
  import { useUserStore } from '@/store/modules/user';

  const userStore = useUserStore();
</script>
