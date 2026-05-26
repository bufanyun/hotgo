<template>
  <n-alert :show-icon="false" title="说明">
    <n-p
      >这里主要演示多租户业务下，不同用户身份如何在同一页面下展示不同的表格功能和字段数据，以及添加/编辑购买订单时服务端如何自动维护多租户关系</n-p
    >
    <n-p style="font-weight: 600">不同身份的测试账号</n-p>
    <n-table :bordered="false" :single-line="false" size="small">
      <thead>
        <tr>
          <th class="table-center">身份</th>
          <th class="table-center">ID</th>
          <th class="table-center">账号</th>
          <th class="table-center">密码</th>
          <th>身份描述</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="account in accounts" :key="account.id">
          <td class="table-center">{{ account.type }}</td>
          <td class="table-center">{{ account.id }}</td>
          <td class="table-center">{{ account.username }}</td>
          <td class="table-center">{{ account.password }}</td>
          <td>{{ account.dc }}</td>
        </tr>
      </tbody>
    </n-table>
  </n-alert>
</template>

<script setup lang="ts">
  interface Account {
    type: string;
    id: number;
    username: string;
    password: string;
    dc: string;
  }

  const accounts: Account[] = [
    {
      type: '公司',
      id: 1,
      username: 'admin',
      password: '123456',
      dc: '可见全部数据。管理整个平台，包括商户和用户账户',
    },
    {
      type: '租户',
      id: 8,
      username: 'ameng',
      password: '123456',
      dc: '可见自己下面的商户和用户数据。多租户系统中顶层实体，有自己的多个商户、用户、产品、订单等',
    },
    {
      type: '商户',
      id: 11,
      username: 'abai',
      password: '123456',
      dc: '可见自己下面的用户数据。受租户的监管和管理，可以独立经营的实体，提供产品或服务，管理自己的业务，包括库存管理、订单处理、结算等',
    },
    {
      type: '用户',
      id: 12,
      username: 'asong',
      password: '123456',
      dc: '只能看到自己数据。真正购买产品或享受服务的人，与商户互动，管理个人信息等个性化功能',
    },
  ];
</script>

<style scoped lang="less">
  .table-center {
    text-align: center;
    min-width: 80px;
  }
</style>
